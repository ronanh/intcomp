package intcomp

import (
	"unsafe"

	"github.com/klauspost/compress/zstd"
)

type BytesSlice struct {
	buf     []byte
	offsets []int
	values  [][]byte
}

func (bs *BytesSlice) Len() int {
	return len(bs.offsets)
}

func (bs *BytesSlice) Value(i int) []byte {
	if i == len(bs.offsets)-1 {
		return bs.buf[bs.offsets[i]:]
	}
	return bs.buf[bs.offsets[i]:bs.offsets[i+1]]
}

func (bs *BytesSlice) Values() [][]byte {
	if len(bs.offsets) == 0 {
		return nil
	}
	if cap(bs.values) < len(bs.offsets) {
		bs.values = make([][]byte, len(bs.offsets))
	} else {
		bs.values = bs.values[:len(bs.offsets)]
	}
	for i := range bs.values {
		if i == len(bs.values)-1 {
			bs.values[i] = bs.buf[bs.offsets[i]:]
		} else {
			bs.values[i] = bs.buf[bs.offsets[i]:bs.offsets[i+1]]
		}
	}
	return bs.values
}

func (bs *BytesSlice) Reset() {
	bs.buf = bs.buf[:0]
	bs.offsets = bs.offsets[:0]
}

type CompressedBytesSlice struct {
	// compressed buffer of concatenated bytes
	buf []byte
	// uncompressed tail of concatenated bytes
	tail []byte
	// compressed block offsets
	compressedBlockOffsets []int
	// offsets within concatenated bytes
	offsets CompressedSlice[int]
	// last offset for concatenated bytes
	lastOffset int
}

func (cs *CompressedBytesSlice) Len() int {
	return cs.offsets.Len()
}

func (cs *CompressedBytesSlice) CompressedSize() int {
	return len(cs.buf)*8 + len(cs.tail) + cs.offsets.CompressedSize()
}

func (cs *CompressedBytesSlice) MemSize() int {
	return int(unsafe.Sizeof(cs)) + cap(cs.buf)*8 + cap(cs.tail) + cap(cs.compressedBlockOffsets)*4 + cs.offsets.MemSize()
}

func (cp *CompressedBytesSlice) IsBlockCompressed(i int) bool {
	return cp.offsets.IsBlockCompressed(i)
}

func (cs *CompressedBytesSlice) BlockCount() int {
	return cs.offsets.BlockCount()
}

func (cp *CompressedBytesSlice) BlockLen(i int) int {
	return cp.offsets.BlockLen(i)
}

func (cs CompressedBytesSlice) Compress(src [][]byte) CompressedBytesSlice {
	newOffsets := make([]int, len(src))
	curOffset := cs.lastOffset
	for i, v := range src {
		newOffsets[i] = curOffset
		curOffset += len(v)
	}
	cs.lastOffset = curOffset
	firstCompressedBlock := cs.offsets.BlockCount() - 1
	if firstCompressedBlock >= 0 && !cs.offsets.IsBlockCompressed(firstCompressedBlock) {
		firstCompressedBlock--
	}
	firstCompressedBlock++

	cs.offsets = cs.offsets.Compress(newOffsets)
	enc := new(zstd.Encoder)
	originalTail := cs.tail
	// Use sliceOffsets new compressed blocks to add the corresponding data blocks
	blockCount := cs.offsets.BlockCount()
	for i := firstCompressedBlock; i < blockCount; i++ {
		if !cs.offsets.IsBlockCompressed(i) {
			break
		}
		start := cs.offsets.BlockFirstValue(i)
		end := curOffset
		if i+1 < blockCount {
			end = cs.offsets.BlockFirstValue(i + 1)
		}
		blockLen := end - start
		for len(cs.tail) < blockLen {
			cs.tail = append(cs.tail, src[0]...)
			src = src[1:]
		}
		if len(cs.tail) != blockLen {
			panic("invalid tail length")
		}
		cs.compressBlock(cs.tail, enc)
		if len(originalTail) > 0 && sameSlice(cs.tail, originalTail) {
			// we should not modify the original tail data
			cs.tail = nil
		} else {
			cs.tail = cs.tail[:0]
		}
	}
	// Add the remaining input to the tail
	for _, v := range src {
		cs.tail = append(cs.tail, v...)
	}

	return cs
}

func (cp *CompressedBytesSlice) compressBlock(block []byte, enc *zstd.Encoder) {
	// Compress offset block
	cp.compressedBlockOffsets = append(cp.compressedBlockOffsets, len(cp.buf))
	// compress data block
	cp.buf = enc.EncodeAll(block, cp.buf)
}

func (cs *CompressedBytesSlice) Decompress(dst BytesSlice) BytesSlice {
	dec, _ := zstd.NewReader(nil)
	defer dec.Close()

	blockCount := cs.BlockCount()
	for i := 0; i < blockCount; i++ {
		dst, _ = cs.decompressBlock(dst, i, dec)
	}
	return dst
}

func (cs *CompressedBytesSlice) DecompressBlock(dst BytesSlice, i int) (BytesSlice, int) {
	dec, _ := zstd.NewReader(nil)
	defer dec.Close()

	return cs.decompressBlock(dst, i, dec)
}

func (cs *CompressedBytesSlice) decompressBlock(dst BytesSlice, i int, dec *zstd.Decoder) (BytesSlice, int) {
	// Decompress offsets
	blockOffsetPos := len(dst.offsets)
	firstBlockOffset := len(dst.buf)
	dst.offsets, _ = cs.offsets.DecompressBlock(dst.offsets, i)
	// Fix offsets
	if delta := dst.offsets[blockOffsetPos] - firstBlockOffset; delta != 0 {
		for j := blockOffsetPos; j < len(dst.offsets); j++ {
			dst.offsets[j] -= delta
		}
	}

	// Decompress data
	if cs.IsBlockCompressed(i) {
		var blockBuf []byte
		if i+1 < len(cs.compressedBlockOffsets) {
			blockBuf = cs.buf[cs.compressedBlockOffsets[i]:cs.compressedBlockOffsets[i+1]]
		} else {
			blockBuf = cs.buf[cs.compressedBlockOffsets[i]:]
		}
		var err error
		dst.buf, err = dec.DecodeAll(blockBuf, dst.buf)
		if err != nil {
			panic(err)
		}
	} else {
		// last block is uncompressed
		dst.buf = append(dst.buf, cs.tail...)
	}

	return dst, cs.blockOffset(i)
}

func (cs *CompressedBytesSlice) blockOffset(i int) int {
	switch i {
	case 0:
		return 0
	case 1:
		return 64
	case 2:
		return 64 + 128
	default:
		return 64 + 128 + 192 + (i-3)*256
	}
}

func sameSlice(x, y []byte) bool {
	return len(x) == len(y) && &x[0] == &y[0]
}
