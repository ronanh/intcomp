package intcomp

import (
	"unsafe"
)

const (
	groupSize = 64
)

type CompressedSlice[T PackType] struct {
	// compressed buffer
	buf []uint64
	// uncompressed tail
	tail []T
	// block offsets
	blockOffsets []int
	// optional min and max values
	minMax []minMax[T]
}

func (cs *CompressedSlice[T]) WithMinMax(v bool) {
	if v {
		cs.minMax = make([]minMax[T], 0)
	} else {
		cs.minMax = nil
	}
}

func (cs *CompressedSlice[T]) Len() int {
	if cs.BlockCount() == 0 {
		return 0
	}
	if len(cs.tail) > 0 {
		return len(cs.tail) + cs.blockOffset(cs.BlockCount()-1)
	}
	return cs.blockOffset(cs.BlockCount())
}

func (cs *CompressedSlice[T]) CompressedSize() int {
	return len(cs.buf)*8 + len(cs.tail)*int(unsafe.Sizeof(T(0)))
}

func (cs *CompressedSlice[T]) MemSize() int {
	return int(unsafe.Sizeof(cs)) + cap(cs.buf)*8 + cap(cs.tail)*int(unsafe.Sizeof(T(0))) + cap(cs.blockOffsets)*4 + cap(cs.minMax)*int(unsafe.Sizeof(minMax[T]{}))
}

func (cs *CompressedSlice[T]) IsBlockCompressed(i int) bool {
	return i < len(cs.blockOffsets)
}

func (cs *CompressedSlice[T]) BlockCount() int {
	if len(cs.tail) > 0 {
		return len(cs.blockOffsets) + 1
	}
	return len(cs.blockOffsets)
}

func (cs *CompressedSlice[T]) BlockLen(i int) int {
	if i < len(cs.blockOffsets) {
		if i < MaxGroups {
			return (i + 1) * groupSize
		}
		return MaxGroups * groupSize
	}
	if i == len(cs.blockOffsets) && len(cs.tail) > 0 {
		return len(cs.tail)
	}
	panic("invalid block index")
}

func (cs *CompressedSlice[T]) BlockFirstValue(i int) T {
	if i < len(cs.blockOffsets) {
		return *(*T)(unsafe.Pointer(&cs.buf[cs.blockOffsets[i]]))
	}
	if i == len(cs.blockOffsets) && len(cs.tail) > 0 {
		return cs.tail[0]
	}
	panic("invalid block index")
}

func (cs *CompressedSlice[T]) BlockMinMax(i int) (T, T) {
	if cs.minMax == nil {
		panic("minmax not enabled")
	}
	if i < len(cs.minMax) {
		return cs.minMax[i].min, cs.minMax[i].max
	}
	if i == len(cs.minMax) && len(cs.tail) > 0 {
		// find min and max of tail
		min, max := cs.tail[0], cs.tail[0]
		for _, v := range cs.tail {
			if v < min {
				min = v
			}
			if v > max {
				max = v
			}
		}
		return min, max
	}
	panic("invalid block index")
}

func (cs CompressedSlice[T]) Compress(src []T) CompressedSlice[T] {
	return cs.compress(src, 0)
}

func (cs CompressedSlice[T]) CompressLossy(src []T, maxBits int) CompressedSlice[T] {
	minNtz := cs.fractionSize() - maxBits
	if minNtz < 1 {
		panic("maxBits too large")
	}
	return cs.compress(src, minNtz)
}

func (cs CompressedSlice[T]) compress(src []T, minNtz int) CompressedSlice[T] {
	// Compute nbGroups:
	// first block should have 1 group
	// second block should have 2 groups
	// third block should have 3 groups
	// later blocks should have 4 groups
	nbGroups := len(cs.blockOffsets)
	for len(src) > 0 {
		nbGroups++
		if nbGroups > MaxGroups {
			nbGroups = MaxGroups
		}
		if len(cs.tail) > 0 {
			appendTailCount := groupSize*nbGroups - len(cs.tail)
			if appendTailCount > len(src) {
				appendTailCount = len(src)
			}
			cs.tail = append(cs.tail, src[:appendTailCount]...)
			src = src[appendTailCount:]
			if len(cs.tail) == groupSize*nbGroups {
				// tail is full, compress it
				cs.compressBlock(cs.tail, minNtz)
				cs.tail = nil
			}
		} else if len(src) >= groupSize*nbGroups {
			// compress a full block
			cs.compressBlock(src[:groupSize*nbGroups], minNtz)
			src = src[groupSize*nbGroups:]
		} else {
			// append to tail
			cs.tail = append(cs.tail, src...)
			src = nil
		}
	}
	return cs
}

func (cs *CompressedSlice[T]) compressBlock(block []T, minNtz int) {
	initvalue := block[0]
	var bh BlockHeader
	BlockHeaderPos := len(cs.buf)
	// append block header (initvalue + block header)
	cs.buf = append(cs.buf, *(*uint64)(unsafe.Pointer(&initvalue)), uint64(bh))
	for len(block) > 0 {
		var bitlen, ntz int
		group := (*[groupSize]T)(block)
		switch any(T(0)).(type) {
		case float32, float64:
			cs.buf, bitlen, ntz, _ = compressGroupXorAppend(cs.buf, group, initvalue, minNtz)
		default:
			cs.buf, bitlen, ntz, _ = compressGroupDeltaAppend(cs.buf, group, initvalue)
		}
		initvalue = group[groupSize-1]
		bh = bh.AddGroup(bitlen, ntz)
		block = block[groupSize:]
	}
	cs.buf[BlockHeaderPos+1] = uint64(bh)
	cs.blockOffsets = append(cs.blockOffsets, BlockHeaderPos)
}

func (cs *CompressedSlice[T]) Decompress(dst []T) []T {
	if dst == nil {
		dst = make([]T, 0, MaxGroups*groupSize)
	}
	blockCount := cs.BlockCount()
	for i := 0; i < blockCount; i++ {
		dst, _ = cs.DecompressBlock(dst, i)
	}
	return dst
}

func (cs *CompressedSlice[T]) DecompressBlock(dst []T, i int) ([]T, int) {
	if dst == nil {
		dst = make([]T, 0, MaxGroups*groupSize)
	}
	if i < len(cs.blockOffsets) {
		blockOffset := cs.blockOffsets[i]
		initvalue := *(*T)(unsafe.Pointer(&cs.buf[blockOffset]))
		bh := BlockHeader(cs.buf[blockOffset+1])
		in := cs.buf[cs.blockOffsets[i]+2:]
		nbGroups := bh.GroupCount()
		for i := 0; i < nbGroups; i++ {
			bitlen, ntz := bh.GetGroup(i)
			switch any(T(0)).(type) {
			case float32, float64:
				dst = decompressGroupXorAppend(dst, in, initvalue, bitlen, ntz)
			default:
				dst = decompressGroupDeltaAppend(dst, in, initvalue, bitlen, ntz)
			}
			initvalue = dst[len(dst)-1]
			in = in[bitlen-ntz:]
		}
		return dst, cs.blockOffset(i)
	}
	if i == len(cs.blockOffsets) && len(cs.tail) > 0 {
		return append(dst, cs.tail...), cs.blockOffset(i)
	}
	panic("invalid block index")
}

func (cs *CompressedSlice[T]) blockOffset(i int) int {
	switch i {
	case 0:
		return 0
	case 1:
		return groupSize
	case 2:
		return (1 + 2) * groupSize
	default:
		return (1+2+3)*groupSize + (i-3)*(MaxGroups*groupSize)
	}
}

func (cs CompressedSlice[T]) fractionSize() int {
	switch any(T(0)).(type) {
	case float32:
		// size of float32 fraction is 23 bits
		// https://en.wikip
		return 23
	case float64:
		// size of float64 fraction is 52 bits
		// https://en.wikipedia.org/wiki/Double-precision_floating-point_format
		return 52
	default:
		panic("fractionSize applies only to float types")
	}
}

type minMax[T PackType] struct {
	min T
	max T
}
