package intcomp

import (
	"math/bits"
	"unsafe"
)

const (
	BitPackingBlockSize32 = 128
	BitPackingBlockSize64 = 256
)

// CompressDeltaBinPackInt32 compress blocks of 128 integers from `in`
// and append to `out`. `out` slice will be resized if necessary.
// The function returns the values from `in` that were not compressed (could
// not fit into a block), and the updated `out` slice.
//
// Compression logic:
//  1. Difference of consecutive inputs is computed (differential coding)
//  2. ZigZag encoding is applied if a block contains at least one negative delta value
//  3. The result is bit packed into the optimal number of bits for the block
func CompressDeltaBinPackInt32(in []int32, out []uint32) ([]int32, []uint32) {
	if len(in) < BitPackingBlockSize32 {
		return in, out
	}
	initoffset := in[0]

	if out == nil {
		out = make([]uint32, 0, len(in)/2)
	}
	headerpos := len(out)
	// header completed at end
	out = append(out, 0, 0, uint32(initoffset))

	var inpos int
	for ; ; inpos += BitPackingBlockSize32 {
		// babysit Go bounds check
		if inpos < 0 || inpos >= len(in) {
			break
		}
		block := in[inpos:]
		if len(block) < BitPackingBlockSize32 {
			break
		}

		const groupSize = BitPackingBlockSize32 / 4
		group1 := (*[groupSize]int32)(block[0*groupSize : 1*groupSize])
		group2 := (*[groupSize]int32)(block[1*groupSize : 2*groupSize])
		group3 := (*[groupSize]int32)(block[2*groupSize : 3*groupSize])
		group4 := (*[groupSize]int32)(block[3*groupSize : 4*groupSize])

		bitlen1, sign1 := deltaBitLenAndSignInt32(group1, initoffset)
		bitlen2, sign2 := deltaBitLenAndSignInt32(group2, group1[31])
		bitlen3, sign3 := deltaBitLenAndSignInt32(group3, group2[31])
		bitlen4, sign4 := deltaBitLenAndSignInt32(group4, group3[31])

		// write block header
		out = append(out, uint32((sign1<<31)|(bitlen1<<24)|
			(sign2<<23)|(bitlen2<<16)|
			(sign3<<15)|(bitlen3<<8)|
			(sign4<<7)|bitlen4))

		// write groups (4 x 32 packed inputs)
		if sign1 == 0 {
			out = appendGroup_int32(out, group1, initoffset, bitlen1)
		} else {
			out = appendGroupZigZag_int32(out, group1, initoffset, bitlen1)
		}

		if sign2 == 0 {
			out = appendGroup_int32(out, group2, group1[31], bitlen2)
		} else {
			out = appendGroupZigZag_int32(out, group2, group1[31], bitlen2)
		}

		if sign3 == 0 {
			out = appendGroup_int32(out, group3, group2[31], bitlen3)
		} else {
			out = appendGroupZigZag_int32(out, group3, group2[31], bitlen3)
		}

		if sign4 == 0 {
			out = appendGroup_int32(out, group4, group3[31], bitlen4)
		} else {
			out = appendGroupZigZag_int32(out, group4, group3[31], bitlen4)
		}

		initoffset = group4[31]
	}

	// write header
	var header []uint32
	// babysit Go bounds check
	if headerpos >= 0 && headerpos <= len(out) {
		header = out[headerpos:]
	}
	if len(header) > 1 {
		header[0] = uint32(inpos)
		header[1] = uint32(len(out) - headerpos)
	}

	var remain []int32
	// babysit Go bounds check
	if inpos >= 0 && inpos < len(in) {
		remain = in[inpos:]
	}
	return remain, out
}

func deltaBitLenAndSignInt32(in *[32]int32, pass int32) (int, int) {
	mask := deltaMask32(in, pass)
	sign := int(mask & 1)
	// remove sign in zigzag encoding
	mask >>= 1
	return bits.Len32(uint32(mask)) + sign, sign
}

// UncompressDeltaBinPackInt32 uncompress one ore more blocks of 128 integers from `in`
// and append the result to `out`. `out` slice will be resized if necessary.
// The function returns the values from `in` that were not uncompressed, and the updated `out` slice.
func UncompressDeltaBinPackInt32(in []uint32, out []int32) ([]uint32, []int32) {
	if len(in) == 0 {
		return in, out
	}
	// read header
	initoffset := int32(in[2])
	inpos := 3

	// output fix
	outpos := len(out)
	if l := int(uint(in[0])); l <= cap(out)-len(out) {
		out = out[:len(out)+l]
	} else {
		grow := make([]int32, len(out)+l)
		copy(grow, out)
		out = grow
	}

	for ; outpos < len(out); outpos += BitPackingBlockSize32 {
		const groupSize = BitPackingBlockSize32 / 4

		tmp := uint32(in[inpos])
		sign1 := int(tmp>>31) & 0x1
		sign2 := int(tmp>>23) & 0x1
		sign3 := int(tmp>>15) & 0x1
		sign4 := int(tmp>>7) & 0x1
		bitlen1 := int(tmp>>24) & 0x7F
		bitlen2 := int(tmp>>16) & 0x7F
		bitlen3 := int(tmp>>8) & 0x7F
		bitlen4 := int(tmp & 0x7F)
		inpos++

		if sign1 == 0 {
			deltaUnpack_int32(initoffset, in[inpos:], out[outpos:], bitlen1)
		} else {
			deltaUnpackZigzag_int32(initoffset, in[inpos:], out[outpos:], bitlen1)
		}
		inpos += int(bitlen1)
		initoffset = out[outpos+groupSize-1]

		if sign2 == 0 {
			deltaUnpack_int32(initoffset, in[inpos:], out[outpos+groupSize:], bitlen2)
		} else {
			deltaUnpackZigzag_int32(initoffset, in[inpos:], out[outpos+groupSize:], bitlen2)
		}
		inpos += int(bitlen2)
		initoffset = out[outpos+2*groupSize-1]

		if sign3 == 0 {
			deltaUnpack_int32(initoffset, in[inpos:], out[outpos+2*groupSize:], bitlen3)
		} else {
			deltaUnpackZigzag_int32(initoffset, in[inpos:], out[outpos+2*groupSize:], bitlen3)
		}
		inpos += int(bitlen3)
		initoffset = out[outpos+3*groupSize-1]

		if sign4 == 0 {
			deltaUnpack_int32(initoffset, in[inpos:], out[outpos+3*groupSize:], bitlen4)
		} else {
			deltaUnpackZigzag_int32(initoffset, in[inpos:], out[outpos+3*groupSize:], bitlen4)
		}
		inpos += int(bitlen4)
		initoffset = out[outpos+4*groupSize-1]
	}

	return in[inpos:], out
}

// CompressDeltaBinPackUint32 compress blocks of 128 integers from `in`
// and append to `out`. `out` slice will be resized if necessary.
// The function returns the values from `in` that were not compressed (could
// not fit into a block), and the updated `out` slice.
//
// Compression logic:
//  1. Difference of consecutive inputs is computed (differential coding)
//  2. ZigZag encoding is applied if a block contains at least one negative delta value
//  3. The result is bit packed into the optimal number of bits for the block
func CompressDeltaBinPackUint32(in, out []uint32) ([]uint32, []uint32) {
	if len(in) < BitPackingBlockSize32 {
		return in, out
	}
	initoffset := in[0]

	if out == nil {
		out = make([]uint32, 0, len(in)/2)
	}
	headerpos := len(out)
	// header completed at end
	out = append(out, 0, 0, initoffset)

	var inpos int
	for ; ; inpos += BitPackingBlockSize32 {
		// babysit Go bounds check
		if inpos < 0 || inpos >= len(in) {
			break
		}
		block := in[inpos:]
		if len(block) < BitPackingBlockSize32 {
			break
		}

		const groupSize = BitPackingBlockSize32 / 4
		group1 := (*[groupSize]uint32)(block[0*groupSize : 1*groupSize])
		group2 := (*[groupSize]uint32)(block[1*groupSize : 2*groupSize])
		group3 := (*[groupSize]uint32)(block[2*groupSize : 3*groupSize])
		group4 := (*[groupSize]uint32)(block[3*groupSize : 4*groupSize])

		bitlen1, sign1 := deltaBitLenAndSignUint32(group1, initoffset)
		bitlen2, sign2 := deltaBitLenAndSignUint32(group2, group1[31])
		bitlen3, sign3 := deltaBitLenAndSignUint32(group3, group2[31])
		bitlen4, sign4 := deltaBitLenAndSignUint32(group4, group3[31])

		// write block header
		out = append(out, uint32((sign1<<31)|(bitlen1<<24)|
			(sign2<<23)|(bitlen2<<16)|
			(sign3<<15)|(bitlen3<<8)|
			(sign4<<7)|bitlen4))

		// write groups (4 x 32 packed inputs)
		if sign1 == 0 {
			out = appendGroup_uint32(out, group1, initoffset, bitlen1)
		} else {
			out = appendGroupZigZag_uint32(out, group1, initoffset, bitlen1)
		}

		if sign2 == 0 {
			out = appendGroup_uint32(out, group2, group1[31], bitlen2)
		} else {
			out = appendGroupZigZag_uint32(out, group2, group1[31], bitlen2)
		}

		if sign3 == 0 {
			out = appendGroup_uint32(out, group3, group2[31], bitlen3)
		} else {
			out = appendGroupZigZag_uint32(out, group3, group2[31], bitlen3)
		}

		if sign4 == 0 {
			out = appendGroup_uint32(out, group4, group3[31], bitlen4)
		} else {
			out = appendGroupZigZag_uint32(out, group4, group3[31], bitlen4)
		}

		initoffset = group4[31]
	}

	// write header
	var header []uint32
	// babysit Go bounds check
	if headerpos >= 0 && headerpos <= len(out) {
		header = out[headerpos:]
	}
	if len(header) > 1 {
		header[0] = uint32(inpos)
		header[1] = uint32(len(out) - headerpos)
	}

	var remain []uint32
	// babysit Go bounds check
	if inpos >= 0 && inpos < len(in) {
		remain = in[inpos:]
	}
	return remain, out
}

func deltaBitLenAndSignUint32(in *[32]uint32, pass uint32) (int, int) {
	mask := deltaMask32(in, pass)
	sign := int(mask & 1)
	// remove sign in zigzag encoding
	mask >>= 1
	return bits.Len32(uint32(mask)) + sign, sign
}

// UncompressDeltaBinPackUint32 uncompress one ore more blocks of 128 integers from `in`
// and append the result to `out`. `out` slice will be resized if necessary.
// The function returns the values from `in` that were not uncompressed, and the updated `out` slice.
func UncompressDeltaBinPackUint32(in, out []uint32) ([]uint32, []uint32) {
	if len(in) == 0 {
		return in, out
	}
	// read header
	initoffset := in[2]
	inpos := 3

	// output fix
	outpos := len(out)
	if l := int(uint(in[0])); l <= cap(out)-len(out) {
		out = out[:len(out)+l]
	} else {
		grow := make([]uint32, len(out)+l)
		copy(grow, out)
		out = grow
	}

	for ; outpos < len(out); outpos += BitPackingBlockSize32 {
		const groupSize = BitPackingBlockSize32 / 4

		tmp := uint32(in[inpos])
		sign1 := int(tmp>>31) & 0x1
		sign2 := int(tmp>>23) & 0x1
		sign3 := int(tmp>>15) & 0x1
		sign4 := int(tmp>>7) & 0x1
		bitlen1 := int(tmp>>24) & 0x7F
		bitlen2 := int(tmp>>16) & 0x7F
		bitlen3 := int(tmp>>8) & 0x7F
		bitlen4 := int(tmp & 0x7F)
		inpos++

		if sign1 == 0 {
			deltaUnpack_uint32(initoffset, in[inpos:], out[outpos:], bitlen1)
		} else {
			deltaUnpackZigzag_uint32(initoffset, in[inpos:], out[outpos:], bitlen1)
		}
		inpos += bitlen1
		initoffset = out[outpos+groupSize-1]

		if sign2 == 0 {
			deltaUnpack_uint32(initoffset, in[inpos:], out[outpos+groupSize:], bitlen2)
		} else {
			deltaUnpackZigzag_uint32(initoffset, in[inpos:], out[outpos+groupSize:], bitlen2)
		}
		inpos += bitlen2
		initoffset = out[outpos+2*groupSize-1]

		if sign3 == 0 {
			deltaUnpack_uint32(initoffset, in[inpos:], out[outpos+2*groupSize:], bitlen3)
		} else {
			deltaUnpackZigzag_uint32(initoffset, in[inpos:], out[outpos+2*groupSize:], bitlen3)
		}
		inpos += bitlen3
		initoffset = out[outpos+3*groupSize-1]

		if sign4 == 0 {
			deltaUnpack_uint32(initoffset, in[inpos:], out[outpos+3*groupSize:], bitlen4)
		} else {
			deltaUnpackZigzag_uint32(initoffset, in[inpos:], out[outpos+3*groupSize:], bitlen4)
		}
		inpos += bitlen4
		initoffset = out[outpos+4*groupSize-1]
	}

	return in[inpos:], out
}

// CompressDeltaBinPackInt64 compress blocks of 256 integers from `in`
// and append to `out`. `out` slice will be resized if necessary.
// The function returns the values from `in` that were not compressed (could
// not fit into a block), and the updated `out` slice.
//
// Compression logic:
//  1. Difference of consecutive inputs is computed (differential coding)
//  2. ZigZag encoding is applied if a block contains at least one negative delta value
//  3. The result is bit packed into the optimal number of bits for the block
func CompressDeltaBinPackInt64(in []int64, out []uint64) ([]int64, []uint64) {
	return CompressDeltaBinPack(in, out)
}

func CompressDeltaBinPack[T PackType](in []T, out []uint64) ([]T, []uint64) {
	if len(in) < BitPackingBlockSize64 {
		return in, out
	}
	initoffset := in[0]

	if out == nil {
		out = make([]uint64, 0, len(in)/2)
	}
	headerpos := len(out)
	// header completed at end
	out = append(out, 0, uint64(initoffset))

	const groupSize = BitPackingBlockSize64 / 4
	var nBlocks int
	for len(in) >= BitPackingBlockSize64 {
		blockHeaderpos := len(out)
		var bh BlockHeader
		out = append(out, uint64(bh)) // placeholder for header

		for i := 0; i < MaxGroups; i++ {
			group := (*[groupSize]T)(in)
			var bitlen, ntz int

			out, bitlen, ntz, _ = compressGroupDeltaAppend(out, group, initoffset)
			bh = bh.AddGroup(bitlen, ntz)
			initoffset = group[63]
			in = in[groupSize:]
		}
		out[blockHeaderpos] = uint64(bh)
		nBlocks++
	}
	out[headerpos] = uint64(nBlocks*BitPackingBlockSize64) | uint64(len(out)-headerpos)<<32
	return in, out
}

func deltaBitTzAndLenAndSignInt64(in *[64]int64, pass int64) (int, int, int) {
	mask := deltaMask64(in, pass)
	sign := int(mask & 1)
	// remove sign in zigzag encoding
	mask >>= 1
	var ntz int
	if mask != 0 {
		ntz = bits.TrailingZeros64(mask)
	}
	return ntz, bits.Len64(uint64(mask)) + sign, sign
}

// UncompressDeltaBinPackInt64 uncompress one ore more blocks of 256 integers from `in`
// and append the result to `out`. `out` slice will be resized if necessary.
// The function returns the values from `in` that were not uncompressed, and the updated `out` slice.
func UncompressDeltaBinPackInt64Old(in []uint64, out []int64) ([]uint64, []int64) {
	if len(in) == 0 {
		return in, out
	}
	// read header
	initoffset := int64(in[1])
	inpos := 2

	// output fix
	outpos := len(out)
	if l := int(uint32(in[0])); l <= cap(out)-len(out) {
		out = out[:len(out)+l]
	} else {
		grow := make([]int64, len(out)+l)
		copy(grow, out)
		out = grow
	}
	if len(in) == 0 {
		return in, out
	}

	for ; outpos < len(out); outpos += BitPackingBlockSize64 {
		const groupSize = BitPackingBlockSize64 / 4

		tmp := uint64(in[inpos])
		ntz1 := int(tmp>>56) & 0xFF
		ntz2 := int(tmp>>48) & 0xFF
		ntz3 := int(tmp>>40) & 0xFF
		ntz4 := int(tmp>>32) & 0xFF
		sign1 := int(tmp>>31) & 1
		sign2 := int(tmp>>23) & 1
		sign3 := int(tmp>>15) & 1
		sign4 := int(tmp>>7) & 1
		bitlen1 := int(tmp>>24) & 0x7F
		bitlen2 := int(tmp>>16) & 0x7F
		bitlen3 := int(tmp>>8) & 0x7F
		bitlen4 := int(tmp & 0x7F)
		inpos++

		if sign1 == 0 {
			deltaUnpack_int64(initoffset, in[inpos:], out[outpos:], ntz1, bitlen1)
		} else {
			deltaUnpackZigzag_int64(initoffset, in[inpos:], out[outpos:], ntz1, bitlen1)
		}
		inpos += int(bitlen1 - ntz1)
		initoffset = out[outpos+groupSize-1]

		if sign2 == 0 {
			deltaUnpack_int64(initoffset, in[inpos:], out[outpos+groupSize:], ntz2, bitlen2)
		} else {
			deltaUnpackZigzag_int64(initoffset, in[inpos:], out[outpos+groupSize:], ntz2, bitlen2)
		}
		inpos += int(bitlen2 - ntz2)
		initoffset = out[outpos+2*groupSize-1]

		if sign3 == 0 {
			deltaUnpack_int64(initoffset, in[inpos:], out[outpos+2*groupSize:], ntz3, bitlen3)
		} else {
			deltaUnpackZigzag_int64(initoffset, in[inpos:], out[outpos+2*groupSize:], ntz3, bitlen3)
		}
		inpos += int(bitlen3 - ntz3)
		initoffset = out[outpos+3*groupSize-1]

		if sign4 == 0 {
			deltaUnpack_int64(initoffset, in[inpos:], out[outpos+3*groupSize:], ntz4, bitlen4)
		} else {
			deltaUnpackZigzag_int64(initoffset, in[inpos:], out[outpos+3*groupSize:], ntz4, bitlen4)
		}
		inpos += int(bitlen4 - ntz4)
		initoffset = out[outpos+4*groupSize-1]
	}

	return in[inpos:], out
}

func UncompressDeltaBinPackInt64(in []uint64, out []int64) ([]uint64, []int64) {
	return UncompressDeltaBinPack(in, out)
}

func UncompressDeltaBinPack[T PackType](in []uint64, out []T) ([]uint64, []T) {
	if len(in) == 0 {
		return in, out
	}

	// Ensure output has enough capacity
	outpos := len(out)
	finalLen := int(uint32(in[0])) + len(out)
	if finalLen > cap(out) {
		grow := make([]T, len(out), finalLen)
		copy(grow, out)
		out = grow
	}

	// read header
	initoffset := *(*T)(unsafe.Pointer(&in[1]))
	in = in[2:]
	for ; outpos < finalLen; outpos += BitPackingBlockSize64 {
		bh := BlockHeader(in[0])
		in = in[1:]

		nbGroups := bh.GroupCount()
		for i := 0; i < nbGroups; i++ {
			bitlen, ntz := bh.GetGroup(i)
			out = decompressGroupDeltaAppend(out, in, initoffset, bitlen, ntz)
			in = in[bitlen-ntz:]
			initoffset = out[len(out)-1]
		}
	}

	return in, out
}

// CompressDeltaBinPackUint64 compress blocks of 256 integers from `in`
// and append to `out`. `out` slice will be resized if necessary.
// The function returns the values from `in` that were not compressed (could
// not fit into a block), and the updated `out` slice.
//
// Compression logic:
//  1. Difference of consecutive inputs is computed (differential coding)
//  2. ZigZag encoding is applied if a block contains at least one negative delta value
//  3. The result is bit packed into the optimal number of bits for the block
func CompressDeltaBinPackUint64(in, out []uint64) ([]uint64, []uint64) {
	if len(in) < BitPackingBlockSize64 {
		return in, out
	}
	initoffset := in[0]

	if out == nil {
		out = make([]uint64, 0, len(in)/2)
	}
	headerpos := len(out)
	// header completed at end
	out = append(out, 0, uint64(initoffset))

	var inpos int
	for ; ; inpos += BitPackingBlockSize64 {
		// babysit Go bounds check
		if inpos < 0 || inpos >= len(in) {
			break
		}
		block := in[inpos:]
		if len(block) < BitPackingBlockSize64 {
			break
		}
		const groupSize = BitPackingBlockSize64 / 4
		group1 := (*[groupSize]uint64)(block[0*groupSize : 1*groupSize])
		group2 := (*[groupSize]uint64)(block[1*groupSize : 2*groupSize])
		group3 := (*[groupSize]uint64)(block[2*groupSize : 3*groupSize])
		group4 := (*[groupSize]uint64)(block[3*groupSize : 4*groupSize])

		bitlen1, sign1 := deltaBitLenAndSignUint64(group1, initoffset)
		bitlen2, sign2 := deltaBitLenAndSignUint64(group2, group1[63])
		bitlen3, sign3 := deltaBitLenAndSignUint64(group2, group2[63])
		bitlen4, sign4 := deltaBitLenAndSignUint64(group4, group3[63])

		// write block header (min/max bits)
		out = append(out, uint64(
			(sign1<<31)|(bitlen1<<24)|
				(sign2<<23)|(bitlen2<<16)|
				(sign3<<15)|(bitlen3<<8)|
				(sign4<<7)|bitlen4))

		// write groups (4 x 64 packed inputs)
		if sign1 == 0 {
			out = appendGroup_uint64(out, group1, initoffset, bitlen1)
		} else {
			out = appendGroupZigZag_uint64(out, group1, initoffset, bitlen1)
		}

		if sign2 == 0 {
			out = appendGroup_uint64(out, group2, group1[63], bitlen2)
		} else {
			out = appendGroupZigZag_uint64(out, group2, group1[63], bitlen2)
		}

		if sign3 == 0 {
			out = appendGroup_uint64(out, group3, group2[63], bitlen3)
		} else {
			out = appendGroupZigZag_uint64(out, group3, group2[63], bitlen3)
		}

		if sign4 == 0 {
			out = appendGroup_uint64(out, group4, group3[63], bitlen4)
		} else {
			out = appendGroupZigZag_uint64(out, group4, group3[63], bitlen4)
		}

		initoffset = group4[63]
	}

	// write header
	// babysit Go bounds check
	if headerpos >= 0 && headerpos < len(out) {
		out[headerpos] = uint64(inpos) | uint64(len(out)-headerpos)<<32
	}

	var remain []uint64
	// babysit Go bounds check
	if inpos >= 0 && inpos < len(in) {
		remain = in[inpos:]
	}
	return remain, out
}

func deltaBitLenAndSignUint64(in *[64]uint64, pass uint64) (int, int) {
	mask := deltaMask64(in, pass)
	sign := int(mask & 1)
	// remove sign in zigzag encoding
	mask >>= 1
	return bits.Len64(uint64(mask)) + sign, sign
}

// UncompressDeltaBinPackUint64 uncompress one ore more blocks of 256 integers from `in`
// and append the result to `out`. `out` slice will be resized if necessary.
// The function returns the values from `in` that were not uncompressed, and the updated `out` slice.
func UncompressDeltaBinPackUint64(in, out []uint64) ([]uint64, []uint64) {
	if len(in) == 0 {
		return in, out
	}
	// read header
	initoffset := in[1]
	inpos := 2

	// output fix
	outpos := len(out)
	if l := int(uint32(in[0])); l <= cap(out)-len(out) {
		out = out[:len(out)+l]
	} else {
		grow := make([]uint64, len(out)+l)
		copy(grow, out)
		out = grow
	}

	for ; outpos < len(out); outpos += BitPackingBlockSize64 {
		const groupSize = BitPackingBlockSize64 / 4

		tmp := uint64(in[inpos])
		sign1 := int(tmp>>31) & 1
		sign2 := int(tmp>>23) & 1
		sign3 := int(tmp>>15) & 1
		sign4 := int(tmp>>7) & 1
		bitlen1 := int(tmp>>24) & 0x7F
		bitlen2 := int(tmp>>16) & 0x7F
		bitlen3 := int(tmp>>8) & 0x7F
		bitlen4 := int(tmp & 0x7F)
		inpos++

		if sign1 == 0 {
			deltaUnpack_uint64(initoffset, in[inpos:], out[outpos:], bitlen1)
		} else {
			deltaUnpackZigzag_uint64(initoffset, in[inpos:], out[outpos:], bitlen1)
		}
		inpos += bitlen1
		initoffset = out[outpos+groupSize-1]

		if sign2 == 0 {
			deltaUnpack_uint64(initoffset, in[inpos:], out[outpos+groupSize:], bitlen2)
		} else {
			deltaUnpackZigzag_uint64(initoffset, in[inpos:], out[outpos+groupSize:], bitlen2)
		}
		inpos += bitlen2
		initoffset = out[outpos+2*groupSize-1]

		if sign3 == 0 {
			deltaUnpack_uint64(initoffset, in[inpos:], out[outpos+2*groupSize:], bitlen3)
		} else {
			deltaUnpackZigzag_uint64(initoffset, in[inpos:], out[outpos+2*groupSize:], bitlen3)
		}
		inpos += bitlen3
		initoffset = out[outpos+3*groupSize-1]

		if sign4 == 0 {
			deltaUnpack_uint64(initoffset, in[inpos:], out[outpos+3*groupSize:], bitlen4)
		} else {
			deltaUnpackZigzag_uint64(initoffset, in[inpos:], out[outpos+3*groupSize:], bitlen4)
		}
		inpos += bitlen4
		initoffset = out[outpos+4*groupSize-1]
	}

	return in[inpos:], out
}

// DeltaMask32 returns the bits in use for all deltas on in. Pass is the last
// value from the previous in or in[0] when first in line.
// Note that deltas are zig-zag encoded to optimise the leading-zero count.
func deltaMask32[T uint32 | int32](in *[32]T, pass T) uint32 {
	var mask uint32
	{
		const base = 0
		d0 := int32(in[base] - pass)
		d1 := int32(in[base+1] - in[base])
		d2 := int32(in[base+2] - in[base+1])
		d3 := int32(in[base+3] - in[base+2])
		mask |= uint32((d0<<1)^(d0>>31)) | uint32((d1<<1)^(d1>>31))
		mask |= uint32((d2<<1)^(d2>>31)) | uint32((d3<<1)^(d3>>31))
	}
	{
		const base = 4
		d0 := int32(in[base] - in[base-1])
		d1 := int32(in[base+1] - in[base])
		d2 := int32(in[base+2] - in[base+1])
		d3 := int32(in[base+3] - in[base+2])
		mask |= uint32((d0<<1)^(d0>>31)) | uint32((d1<<1)^(d1>>31))
		mask |= uint32((d2<<1)^(d2>>31)) | uint32((d3<<1)^(d3>>31))
	}
	{
		const base = 8
		d0 := int32(in[base] - in[base-1])
		d1 := int32(in[base+1] - in[base])
		d2 := int32(in[base+2] - in[base+1])
		d3 := int32(in[base+3] - in[base+2])
		mask |= uint32((d0<<1)^(d0>>31)) | uint32((d1<<1)^(d1>>31))
		mask |= uint32((d2<<1)^(d2>>31)) | uint32((d3<<1)^(d3>>31))
	}
	{
		const base = 12
		d0 := int32(in[base] - in[base-1])
		d1 := int32(in[base+1] - in[base])
		d2 := int32(in[base+2] - in[base+1])
		d3 := int32(in[base+3] - in[base+2])
		mask |= uint32((d0<<1)^(d0>>31)) | uint32((d1<<1)^(d1>>31))
		mask |= uint32((d2<<1)^(d2>>31)) | uint32((d3<<1)^(d3>>31))
	}
	{
		const base = 16
		d0 := int32(in[base] - in[base-1])
		d1 := int32(in[base+1] - in[base])
		d2 := int32(in[base+2] - in[base+1])
		d3 := int32(in[base+3] - in[base+2])
		mask |= uint32((d0<<1)^(d0>>31)) | uint32((d1<<1)^(d1>>31))
		mask |= uint32((d2<<1)^(d2>>31)) | uint32((d3<<1)^(d3>>31))
	}
	{
		const base = 20
		d0 := int32(in[base] - in[base-1])
		d1 := int32(in[base+1] - in[base])
		d2 := int32(in[base+2] - in[base+1])
		d3 := int32(in[base+3] - in[base+2])
		mask |= uint32((d0<<1)^(d0>>31)) | uint32((d1<<1)^(d1>>31))
		mask |= uint32((d2<<1)^(d2>>31)) | uint32((d3<<1)^(d3>>31))
	}
	{
		const base = 24
		d0 := int32(in[base] - in[base-1])
		d1 := int32(in[base+1] - in[base])
		d2 := int32(in[base+2] - in[base+1])
		d3 := int32(in[base+3] - in[base+2])
		mask |= uint32((d0<<1)^(d0>>31)) | uint32((d1<<1)^(d1>>31))
		mask |= uint32((d2<<1)^(d2>>31)) | uint32((d3<<1)^(d3>>31))
	}
	{
		const base = 28
		d0 := int32(in[base] - in[base-1])
		d1 := int32(in[base+1] - in[base])
		d2 := int32(in[base+2] - in[base+1])
		d3 := int32(in[base+3] - in[base+2])
		mask |= uint32((d0<<1)^(d0>>31)) | uint32((d1<<1)^(d1>>31))
		mask |= uint32((d2<<1)^(d2>>31)) | uint32((d3<<1)^(d3>>31))
	}
	return mask
}

// DeltaMask64 returns the bits in use for all deltas on in. Pass is the last
// value from the previous in or in[0] when first in line.
// Note that deltas are zig-zag encoded to optimise the leading-zero count.
func deltaMask64[T uint64 | int64](in *[64]T, pass T) uint64 {
	var mask uint64
	for _, batch := range []*[32]T{(*[32]T)(in[:32]), (*[32]T)(in[32:64])} {
		{
			const base = 0
			d0 := int64(batch[base] - pass)
			d1 := int64(batch[base+1] - batch[base])
			d2 := int64(batch[base+2] - batch[base+1])
			d3 := int64(batch[base+3] - batch[base+2])
			mask |= uint64((d0<<1)^(d0>>63)) | uint64((d1<<1)^(d1>>63))
			mask |= uint64((d2<<1)^(d2>>63)) | uint64((d3<<1)^(d3>>63))
		}
		{
			const base = 4
			d0 := int64(batch[base] - batch[base-1])
			d1 := int64(batch[base+1] - batch[base])
			d2 := int64(batch[base+2] - batch[base+1])
			d3 := int64(batch[base+3] - batch[base+2])
			mask |= uint64((d0<<1)^(d0>>63)) | uint64((d1<<1)^(d1>>63))
			mask |= uint64((d2<<1)^(d2>>63)) | uint64((d3<<1)^(d3>>63))
		}
		{
			const base = 8
			d0 := int64(batch[base] - batch[base-1])
			d1 := int64(batch[base+1] - batch[base])
			d2 := int64(batch[base+2] - batch[base+1])
			d3 := int64(batch[base+3] - batch[base+2])
			mask |= uint64((d0<<1)^(d0>>63)) | uint64((d1<<1)^(d1>>63))
			mask |= uint64((d2<<1)^(d2>>63)) | uint64((d3<<1)^(d3>>63))
		}
		{
			const base = 12
			d0 := int64(batch[base] - batch[base-1])
			d1 := int64(batch[base+1] - batch[base])
			d2 := int64(batch[base+2] - batch[base+1])
			d3 := int64(batch[base+3] - batch[base+2])
			mask |= uint64((d0<<1)^(d0>>63)) | uint64((d1<<1)^(d1>>63))
			mask |= uint64((d2<<1)^(d2>>63)) | uint64((d3<<1)^(d3>>63))
		}
		{
			const base = 16
			d0 := int64(batch[base] - batch[base-1])
			d1 := int64(batch[base+1] - batch[base])
			d2 := int64(batch[base+2] - batch[base+1])
			d3 := int64(batch[base+3] - batch[base+2])
			mask |= uint64((d0<<1)^(d0>>63)) | uint64((d1<<1)^(d1>>63))
			mask |= uint64((d2<<1)^(d2>>63)) | uint64((d3<<1)^(d3>>63))
		}
		{
			const base = 20
			d0 := int64(batch[base] - batch[base-1])
			d1 := int64(batch[base+1] - batch[base])
			d2 := int64(batch[base+2] - batch[base+1])
			d3 := int64(batch[base+3] - batch[base+2])
			mask |= uint64((d0<<1)^(d0>>63)) | uint64((d1<<1)^(d1>>63))
			mask |= uint64((d2<<1)^(d2>>63)) | uint64((d3<<1)^(d3>>63))
		}
		{
			const base = 24
			d0 := int64(batch[base] - batch[base-1])
			d1 := int64(batch[base+1] - batch[base])
			d2 := int64(batch[base+2] - batch[base+1])
			d3 := int64(batch[base+3] - batch[base+2])
			mask |= uint64((d0<<1)^(d0>>63)) | uint64((d1<<1)^(d1>>63))
			mask |= uint64((d2<<1)^(d2>>63)) | uint64((d3<<1)^(d3>>63))
		}
		{
			const base = 28
			d0 := int64(batch[base] - batch[base-1])
			d1 := int64(batch[base+1] - batch[base])
			d2 := int64(batch[base+2] - batch[base+1])
			d3 := int64(batch[base+3] - batch[base+2])
			mask |= uint64((d0<<1)^(d0>>63)) | uint64((d1<<1)^(d1>>63))
			mask |= uint64((d2<<1)^(d2>>63)) | uint64((d3<<1)^(d3>>63))
			pass = batch[base+3]
		}
	}
	return mask
}
