package intcomp_test

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"testing"
	"unsafe"

	"github.com/ronanh/intcomp"
)

func TestCompressedIntBuffer(t *testing.T) {
	tests := []struct {
		name       string
		bufs       [][]int
		withMinMax bool
	}{
		{name: "empty buffer",
			bufs:       [][]int{},
			withMinMax: false,
		},
		{name: "one buffer 10",
			bufs:       [][]int{genBuf[int](0, 10)},
			withMinMax: false,
		},
		{name: "one buffer 64",
			bufs:       [][]int{genBuf[int](0, 64)},
			withMinMax: false,
		},
		{name: "one buffer 96",
			bufs:       [][]int{genBuf[int](0, 96)},
			withMinMax: false,
		},
		{name: "one buffer 128+64",
			bufs:       [][]int{genBuf[int](0, 128+64)},
			withMinMax: false,
		},
		{name: "one buffer 128+64+1",
			bufs:       [][]int{genBuf[int](0, 128+64+1)},
			withMinMax: false,
		},
		{name: "one buffer 256+128+64",
			bufs:       [][]int{genBuf[int](0, 256+128+64)},
			withMinMax: false,
		},
		{name: "one buffer 256+128+64+1",
			bufs:       [][]int{genBuf[int](0, 256+128+64+1)},
			withMinMax: false,
		},
		{name: "two buffers",
			bufs:       [][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, {11, 12, 13, 14, 15, 16, 17, 18, 19, 20}},
			withMinMax: false,
		},
		{name: "three buffers",
			bufs:       [][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, {11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64}},
			withMinMax: false,
		},
		{name: "empty buffer with minmax",
			bufs:       [][]int{},
			withMinMax: true,
		},
		{name: "one buffer with minmax",
			bufs:       [][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			withMinMax: true,
		},
		{name: "two buffers with minmax",
			bufs:       [][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, {11, 12, 13, 14, 15, 16, 17, 18, 19, 20}},
			withMinMax: true,
		},
		{name: "three buffers with minmax",
			bufs:       [][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, {11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, {21, 22, 23, 24, 25, 26, 27, 28, 29, 30}},
			withMinMax: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testCompressedBuffer(tt.bufs, tt.withMinMax, t)
		})
	}
}

func TestCompressedFloatBuffer(t *testing.T) {
	tests := []struct {
		name       string
		bufs       [][]float64
		withMinMax bool
	}{
		{name: "one buffer 64",
			bufs:       [][]float64{genBuf[float64](0, 64)},
			withMinMax: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testCompressedBuffer(tt.bufs, tt.withMinMax, t)
		})
	}
}

func testCompressedBuffer[T intcomp.PackType](bufs [][]T, withMinMax bool, t *testing.T) {
	// create a new compressed buffer with minmax enabled
	var cb intcomp.CompressedSlice[T]
	cb.WithMinMax(true)

	wantLen := 0
	wantDecompressed := make([]T, 0)
	// compress two buffers
	for _, buf := range bufs {
		wantLen += len(buf)
		wantDecompressed = append(wantDecompressed, buf...)
		cb = cb.Compress(buf)
	}

	// check Len
	if gotLen := cb.Len(); gotLen != wantLen {
		t.Errorf("cb.Len() = %d, want %d", gotLen, wantLen)
	}

	// check BlockCount
	wantBlockCount := 0
	firstValuePos := 0
	switch {
	case wantLen == 0:
		wantBlockCount = 0
		firstValuePos = -1
	case wantLen <= 64:
		wantBlockCount = 1
		firstValuePos = 0
	case wantLen <= 64+128:
		wantBlockCount = 2
		firstValuePos = 64
	default:
		wantBlockCount = 3 + (wantLen-128-64)/256
		if wantBlockCount == 3 {
			firstValuePos = 64 + 128
		} else {
			firstValuePos = 64 + 128 + 192 + (wantBlockCount-4)*256
		}
	}
	if gotBlockCount := cb.BlockCount(); gotBlockCount != wantBlockCount {
		t.Errorf("cb.BlockCount() = %d, want %d", gotBlockCount, wantBlockCount)
	}

	// check BlockFirstValue
	if firstValuePos >= 0 {
		wantFirstValue := wantDecompressed[firstValuePos]
		if gotFirstValue := cb.BlockFirstValue(cb.BlockCount() - 1); gotFirstValue != wantFirstValue {
			t.Errorf("cb.BlockFirstValue(0) = %v, want %v", gotFirstValue, wantFirstValue)
		}
	}

	// check Decompress
	if gotDecompressed := cb.Decompress(nil); !reflect.DeepEqual(gotDecompressed, wantDecompressed) {
		t.Errorf("cb.Decompress() = %v, want %v", gotDecompressed, wantDecompressed)
	}

	// check iteration
	var block []T
	var k int
	for i := 0; i < cb.BlockCount(); i++ {
		var blockPos int
		block, blockPos = cb.DecompressBlock(block[:0], i)
		for j, v := range block {
			if j+blockPos != k {
				t.Errorf("bad iteration")
			}
			if v != wantDecompressed[k] {
				t.Errorf("bad iteration")
			}
			k++
		}
	}
	if k != wantLen {
		t.Errorf("bad iteration len")
	}

	// check reverse iteration
	k = wantLen - 1
	for i := cb.BlockCount() - 1; i >= 0; i-- {
		block, blockPos := cb.DecompressBlock(block[:0], i)
		for j := len(block) - 1; j >= 0; j-- {
			if j+blockPos != k {
				t.Errorf("bad reverse iteration")
			}
			if block[j] != wantDecompressed[k] {
				t.Errorf("bad reverse iteration")
			}
			k--
		}
	}
}

func TestCompressBufferConstant(t *testing.T) {
	var testInput []int64
	for i := 0; i < 64; i++ {
		testInput = append(testInput, 123456789)
	}

	var cb intcomp.CompressedSlice[int64]
	cb = cb.Compress(testInput)
	if got := cb.CompressedSize(); got != 16 {
		t.Fatalf("expected 16 bytes, got: %d", got)
	}
}

func TestCompressBufferOneBit(t *testing.T) {
	var testInput []int64
	for i := 0; i < 64; i++ {
		testInput = append(testInput, 123456789+int64(i))
	}

	var cb intcomp.CompressedSlice[int64]
	cb = cb.Compress(testInput)
	if got := cb.CompressedSize(); got != 24 {
		t.Fatalf("expected 24 bytes, got: %d", got)
	}
}

func TestCompressBufferTwoBit(t *testing.T) {
	var testInput []int64
	for i := 0; i < 64; i++ {
		testInput = append(testInput, 123456789+2*int64(i))
	}

	var cb intcomp.CompressedSlice[int64]
	cb = cb.Compress(testInput)
	if got := cb.CompressedSize(); got != 24 {
		t.Fatalf("expected 24 bytes, got: %d", got)
	}
}

func TestCompressBufferNegative(t *testing.T) {
	var testInput []int64
	for i := 0; i < 64; i++ {
		testInput = append(testInput, 123456789-int64(i))
	}

	var cb intcomp.CompressedSlice[int64]
	cb = cb.Compress(testInput)
	if got := cb.CompressedSize(); got != 24 {
		t.Fatalf("expected 24 bytes, got: %d", got)
	}
}

func TestCompressBufferFull(t *testing.T) {
	rand.Seed(1) // nolint
	// iterate over all possible types of bit packing
	for _, typ := range []reflect.Kind{reflect.Int, reflect.Int64, reflect.Uint64, reflect.Int32, reflect.Uint32, reflect.Float32, reflect.Float64} {
		// call testCompressBufferFull with proper type
		switch typ {
		case reflect.Int:
			testCompressBufferIntFull[int](t)
		case reflect.Int64:
			testCompressBufferIntFull[int64](t)
		case reflect.Uint64:
			testCompressBufferIntFull[uint64](t)
		case reflect.Int32:
			testCompressBufferIntFull[int32](t)
		case reflect.Uint32:
			testCompressBufferIntFull[uint32](t)
		case reflect.Float32:
			testCompressBufferIntFull[float32](t)
		case reflect.Float64:
			testCompressBufferIntFull[float64](t)
		}
	}
}

func testCompressBufferIntFull[T intcomp.PackType](t *testing.T) {
	bufferSize := 64 + 128 + 192 + 256 + 256
	maxBits := int(unsafe.Sizeof(T(0))) * 8
	for nBits := 0; nBits <= maxBits; nBits++ {
		switch any(T(0)).(type) {
		case float32:
			// exponent is 8 bits, mantissa is 23 bits, sign is 1 bit
			// exponent should not be 0
			if nBits < 32-8-1 {
				continue
			}
		case float64:
			// exponent is 11 bits, mantissa is 52 bits, sign is 1 bit
			// exponent should not be 0
			if nBits < 64-11-1 {
				continue
			}
		}
		for ntz := 0; ntz < nBits; ntz++ {
			if nBits > maxBits {
				continue
			}
			if nBits > 0 && ntz >= nBits {
				continue
			}
			var testInput []T
			prev := rand.Uint64()
			testInput = append(testInput, *(*T)(unsafe.Pointer(&prev)))
			for i := 1; i < bufferSize; i++ {
				var udiff uint64
				if nBits-ntz == maxBits {
					udiff = rand.Uint64()
				} else if nBits-ntz > 0 {
					udiff = rand.Uint64()
					udiff %= 1 << (nBits - ntz)
					udiff <<= ntz
				}
				switch any(T(0)).(type) {
				case float32, float64:
					prev = uint64(prev ^ udiff)
				default:
					prev = uint64((-(int64(udiff) & 1)) ^ (int64(udiff) >> 1))
				}
				testInput = append(testInput, *(*T)(unsafe.Pointer(&prev)))
			}

			t.Run(fmt.Sprintf("type=%T nBits=%d ntz=%d", testInput[0], nBits, ntz), func(t *testing.T) {
				var cb intcomp.CompressedSlice[T]
				cb = cb.Compress(testInput)
				// decompress
				got := cb.Decompress(nil)

				// check decompressed
				if !checkEqual(got, testInput, 1) {
					t.Errorf("cb.Decompress() = %v\n--- Want %v", got, testInput)
				}
			})
		}
	}
}

func TestCompressFloat(t *testing.T) {
	rand.Seed(1) // nolint
	testCompressFloat(t, "noise", func(i int) float32 { return float32(rand.NormFloat64()) }, 1, 0)
	testCompressFloat(t, "noise", func(i int) float64 { return rand.NormFloat64() }, 1, 0)

	testCompressFloat(t, "constant", func(i int) float32 { return -1234 }, 1, 0)
	testCompressFloat(t, "constant", func(i int) float64 { return -1235 }, 1, 0)

	testCompressFloat(t, "ints", func(i int) float32 { return float32(i) }, 1, 0)
	testCompressFloat(t, "ints", func(i int) float64 { return float64(i) }, 1, 0)

	testCompressFloat(t, "fractions", func(i int) float32 { return float32(i) / 10 }, 1, 0)
	testCompressFloat(t, "fractions", func(i int) float64 { return float64(i) / 10 }, 1, 0)

	testCompressFloat(t, "sampleData", func(i int) float32 { return float32(sampleTimeSerieData[i]) }, 1, 0)
	testCompressFloat(t, "sampleData", func(i int) float64 { return sampleTimeSerieData[i] }, 1, 0)

	testCompressFloat(t, "noise lossy", func(i int) float32 { return float32(rand.NormFloat64()) }, 0.9999, 14)
	testCompressFloat(t, "noise lossy", func(i int) float64 { return rand.NormFloat64() }, 0.9999, 14)

	testCompressFloat(t, "positive noise lossy", func(i int) float32 { return float32(rand.ExpFloat64()) }, 0.9999, 14)
	testCompressFloat(t, "positive noise lossy", func(i int) float64 { return rand.NormFloat64() }, 0.9999, 14)
}

func testCompressFloat[T float32 | float64](t *testing.T, name string, gen func(int) T, precision float64, fractionBits int) {
	var testInput []T
	for i := 0; i < 64+128+192+256+256+13; i++ {
		testInput = append(testInput, gen(i))
	}

	t.Run(fmt.Sprintf("name=%s type=%T precision=%f fractionBits=%d", name, testInput[0], precision, fractionBits), func(t *testing.T) {
		var cb intcomp.CompressedSlice[T]
		if precision < 1 {
			cb = cb.CompressLossy(testInput, fractionBits)
		} else {
			cb = cb.Compress(testInput)
		}
		got := cb.Decompress(nil)
		t.Logf("bits per value: %f", float64(cb.CompressedSize())*8/float64(len(testInput)))

		// check decompressed
		if !checkEqual(got, testInput, precision) {
			t.Errorf("cb.Decompress() = %v\n--- Want %v", got, testInput)
		}
	})
}

func checkEqual[T intcomp.PackType](got, want []T, precision float64) bool {
	if len(got) != len(want) {
		return false
	}
	for i := range got {
		if got[i] == want[i] {
			continue
		}

		switch v := any(want[i]).(type) {
		case float32:
			if math.IsNaN(float64(v)) && math.IsNaN(float64(got[i])) || math.IsInf(float64(v), 0) && math.IsInf(float64(got[i]), 0) {
				continue
			}
		case float64:
			if math.IsNaN(v) && math.IsNaN(float64(got[i])) || math.IsInf(v, 0) && math.IsInf(float64(got[i]), 0) {
				continue
			}
		}
		ref := want[i]
		if ref == 0 {
			ref = got[i]
		}
		if diffPrecision := 1 - math.Abs(float64(got[i]-want[i]))/math.Abs(float64(ref)); diffPrecision < precision {
			return false
		}
	}
	return true
}

func genBuf[T intcomp.PackType](start, size int) []T {
	buf := make([]T, size)
	for i := 0; i < size; i++ {
		buf[i] = T(start + i)
	}
	return buf
}

var (
	sampleTimeSerieData []float64 = []float64{0.0491, 0.0491, 0.0491, 0.0491, 0.0491, 0.0491, 0.0458, 0.0421, 0.0367, 0.0349, 0.0334, 0.0334, 0.0334, 0.0334,
		0.0334, 0.0334, 0.0334, 0.0334, 0.0334, 0.0334, 0.0334, 0.0334, 0.0334, 0.0334, 0.0334, 0.0334, 0.0295, 0.0264, 0.0228, 0.0197, 0.0162, 0.0162,
		0.0162, 0.0162, 0.0162, 0.0162, 0.0162, 0.0162, 0.0162, 0.0162, 0.0162, 0.0162, 0.0162, 0.0162, 0.0162, 0.0162, 0.0165, 0.0168, 0.0142, 0.0121,
		0.0131, 0.0131, 0.0131, 0.0131, 0.0131, 0.0131, 0.0131, 0.0131, 0.0131, 0.0131, 0.0131, 0.0131, 0.0131, 0.0131, 0.0131, 0.0131, 0.0162, 0.0206,
		0.0262, 0.0282, 0.0292, 0.0292, 0.0292, 0.0292, 0.0292, 0.0292, 0.0292, 0.0292, 0.0292, 0.0292, 0.0292, 0.0292, 0.0292, 0.0292, 0.0292, 0.0292,
		0.0323, 0.0351, 0.0391, 0.0415, 0.0425, 0.0425, 0.0425, 0.0425, 0.0425, 0.0425, 0.0425, 0.0425, 0.0425, 0.0425, 0.0425, 0.0425, 0.0425, 0.0425,
		0.0425, 0.0425, 0.0423, 0.0419, 0.0445, 0.0460, 0.0455, 0.0455, 0.0455, 0.0455, 0.0455, 0.0455, 0.0455, 0.0455, 0.0455, 0.0455, 0.0455, 0.0455,
		0.0455, 0.0455, 0.0455, 0.0455, 0.0444, 0.0428, 0.0402, 0.0402, 0.0402, 0.0402, 0.0402, 0.0402, 0.0402, 0.0402, 0.0402, 0.0402, 0.0402, 0.0403,
		0.0403, 0.0403, 0.0403, 0.0403, 0.0403, 0.0403, 0.0401, 0.0408, 0.0435, 0.0430, 0.0433, 0.0433, 0.0433, 0.0433, 0.0433, 0.0433, 0.0433, 0.0433,
		0.0433, 0.0433, 0.0433, 0.0433, 0.0433, 0.0433, 0.0433, 0.0433, 0.0435, 0.0440, 0.0420, 0.0415, 0.0402, 0.0402, 0.0402, 0.0402, 0.0402, 0.0402,
		0.0402, 0.0402, 0.0402, 0.0402, 0.0402, 0.0402, 0.0402, 0.0402, 0.0402, 0.0402, 0.0381, 0.0352, 0.0322, 0.0303, 0.0293, 0.0293, 0.0293, 0.0293,
		0.0295, 0.0295, 0.0295, 0.0295, 0.0295, 0.0294, 0.0294, 0.0294, 0.0294, 0.0294, 0.0294, 0.0294, 0.0272, 0.0274, 0.0256, 0.0257, 0.0253, 0.0253,
		0.0253, 0.0253, 0.0253, 0.0253, 0.0253, 0.0253, 0.0253, 0.0253, 0.0253, 0.0253, 0.0253, 0.0253, 0.0253, 0.0253, 0.0254, 0.0220, 0.0182, 0.0161,
		0.0154, 0.0154, 0.0154, 0.0154, 0.0156, 0.0156, 0.0156, 0.0157, 0.0157, 0.0157, 0.0157, 0.0157, 0.0157, 0.0158, 0.0158, 0.0158, 0.0174, 0.0187,
		0.0217, 0.0261, 0.0281, 0.0292, 0.0292, 0.0292, 0.0290, 0.0290, 0.0290, 0.0290, 0.0290, 0.0290, 0.0290, 0.0290, 0.0290, 0.0290, 0.0290, 0.0290,
		0.0301, 0.0263, 0.0214, 0.0194, 0.0185, 0.0185, 0.0185, 0.0185, 0.0185, 0.0185, 0.0185, 0.0185, 0.0185, 0.0185, 0.0185, 0.0185, 0.0185, 0.0185,
		0.0185, 0.0186, 0.0157, 0.0162, 0.0193, 0.0231, 0.0247, 0.0252, 0.0252, 0.0252, 0.0250, 0.0250, 0.0250, 0.0250, 0.0249, 0.0249, 0.0249, 0.0250,
		0.0249, 0.0248, 0.0248, 0.0248, 0.0248, 0.0236, 0.0205, 0.0162, 0.0141, 0.0131, 0.0131, 0.0131, 0.0131, 0.0131, 0.0131, 0.0131, 0.0131, 0.0131,
		0.0131, 0.0131, 0.0131, 0.0131, 0.0130, 0.0130, 0.0112, 0.0127, 0.0153, 0.0153, 0.0153, 0.0154, 0.0154, 0.0154, 0.0154, 0.0154, 0.0154, 0.0154,
		0.0154, 0.0154, 0.0154, 0.0157, 0.0157, 0.0157, 0.0157, 0.0156, 0.0156, 0.0151, 0.0120, 0.00817, 0.00660, 0.00609, 0.00610, 0.00610, 0.00610,
		0.00610, 0.00610, 0.00680, 0.00680, 0.00680, 0.00680, 0.00678, 0.00678, 0.00678, 0.00678, 0.00678, 0.00518, 0.00518, 0.00518, 0.00518, 0.00518,
		0.00518, 0.00518, 0.00518, 0.00517, 0.00517, 0.00517, 0.00517, 0.00517, 0.00518, 0.00522, 0.00525, 0.00527, 0.00535, 0.00536, 0.00536, 0.00538,
		0.00392, 0.00128, 0.00128, 0.00127, 0.00127, 0.00126, 0.00126, 0.00126, 0.00123, 0.00123, 0.00123, 0.00124, 0.00124, 0.00125, 0.000947, 0.000947,
		0.000947, 0.000947, 0.000986, 0.000986, 0.000986, 0.000986, 0.000997, 0.000997, 0.00101, 0.000997, 0.000997, 0.000997, 0.000997, 0.000997, 0.000292,
		0.000299, 0.000299, 0.000299, 0.000299, 0.000299, 0.000299, 0.000299, 0.000299, 0.00396, 0.00773, 0.0131, 0.0153, 0.0178, 0.0178, 0.0178, 0.0178,
		0.0178, 0.0178, 0.0178, 0.0178, 0.0178, 0.0178, 0.0178, 0.0177, 0.0177, 0.0177, 0.0177, 0.0177, 0.0211, 0.0244, 0.0297, 0.0317, 0.0337, 0.0337,
		0.0337, 0.0337, 0.0337, 0.0337, 0.0337, 0.0337, 0.0337, 0.0337, 0.0337, 0.0337, 0.0337, 0.0337, 0.0337, 0.0337, 0.0344, 0.0373, 0.0421, 0.0443,
		0.0484, 0.0484, 0.0484, 0.0484, 0.0484, 0.0484, 0.0484, 0.0484, 0.0484, 0.0484, 0.0484, 0.0484, 0.0484, 0.0484, 0.0484, 0.0484, 0.0447, 0.0410,
		0.0356, 0.0334, 0.0309, 0.0309, 0.0309, 0.0309, 0.0309, 0.0309, 0.0309, 0.0309, 0.0309, 0.0309, 0.0309, 0.0309, 0.0309, 0.0310, 0.0310, 0.0311,
		0.0276, 0.0243, 0.0191, 0.0171, 0.0150, 0.0150, 0.0150, 0.0150, 0.0150, 0.0150, 0.0150, 0.0150, 0.0150, 0.0150, 0.0150, 0.0150, 0.0150, 0.0151,
		0.0151, 0.0151, 0.0143, 0.0114, 0.00657, 0.00443, 0.000316, 0.000316, 0.000316, 0.000474, 0.000477, 0.000477, 0.000477, 0.000493, 0.000493, 0.000510,
		0.000510, 0.000510, 0.000510, 0.000510, 0.000518, 0.000518, 0.000518, 0.000501, 0.000501, 0.000501, 0.000483, 0.000483, 0.000485, 0.000493, 0.000493,
		0.000493, 0.000493, 0.000493, 0.000500, 0.000500, 0.000500, 0.000481, 0.000448, 0.000398, 0.000345, 0.000293, 0.00257, 0.00544, 0.0102, 0.0117,
		0.0138, 0.0138, 0.0138, 0.0138, 0.0138, 0.0138, 0.0138, 0.0138, 0.0138, 0.0138, 0.0138, 0.0138, 0.0138, 0.0138, 0.0138, 0.0138, 0.0180, 0.0218,
		0.0268, 0.0300, 0.0331, 0.0331, 0.0331, 0.0330, 0.0330, 0.0330, 0.0330, 0.0330, 0.0330, 0.0330, 0.0330, 0.0331, 0.0331, 0.0331, 0.0331, 0.0331,
		0.0359, 0.0383, 0.0417, 0.0459, 0.0476, 0.0476, 0.0476, 0.0476, 0.0476, 0.0476, 0.0476, 0.0476, 0.0476, 0.0476, 0.0476, 0.0476, 0.0476, 0.0476,
		0.0476, 0.0476, 0.0453, 0.0424, 0.0376, 0.0361, 0.0340, 0.0342, 0.0342, 0.0342, 0.0342, 0.0342, 0.0342, 0.0344, 0.0344, 0.0344, 0.0344, 0.0344,
		0.0344, 0.0344, 0.0344, 0.0344, 0.0303, 0.0277, 0.0227, 0.0196, 0.0164, 0.0164, 0.0164, 0.0164, 0.0164, 0.0164, 0.0164, 0.0164, 0.0164, 0.0164,
		0.0164, 0.0162, 0.0162, 0.0162, 0.0162, 0.0162, 0.0134, 0.0111, 0.00763, 0.00343, 0.00171, 0.00171, 0.00171, 0.00171, 0.00171, 0.00171, 0.00171,
		0.00171, 0.00171, 0.00170, 0.00170, 0.00171, 0.00172, 0.00172, 0.00172, 0.00172, 0.00172, 0.00172, 0.00171, 0.00171, 0.00171, 0.00156, 0.00156,
		0.00153, 0.00152, 0.00152, 0.00154, 0.00138, 0.00138, 0.00138, 0.00138, 0.00138, 0.00138, 0.00137, 0.00137, 0.00137, 0.00408, 0.00565, 0.0112, 0.0130,
		0.0155, 0.0155, 0.0155, 0.0155, 0.0155, 0.0155, 0.0155, 0.0155, 0.0155, 0.0155, 0.0155, 0.0156, 0.0156, 0.0156, 0.0156, 0.0156, 0.0189, 0.0216, 0.0261,
		0.0279, 0.0305, 0.0305, 0.0305, 0.0305, 0.0305, 0.0305, 0.0305, 0.0305, 0.0305, 0.0305, 0.0305, 0.0305, 0.0305, 0.0305, 0.0305, 0.0305, 0.0305,
		0.0305, 0.0305, 0.0305, 0.0305, 0.0305, 0.0305, 0.0305, 0.0305, 0.0305, 0.0305, 0.0305, 0.0305, 0.0305, 0.0305, 0.0305, 0.0305, 0.0305, 0.0305,
		0.0305, 0.0305, 0.0309, 0.0306, 0.0302, 0.0290, 0.0290, 0.0290, 0.0290, 0.0290, 0.0290, 0.0290, 0.0290, 0.0290, 0.0290, 0.0290, 0.0290, 0.0290,
		0.0290, 0.0290, 0.0290, 0.0257, 0.0230, 0.0184, 0.0167, 0.0141, 0.0141, 0.0141, 0.0141, 0.0141, 0.0141, 0.0141, 0.0141, 0.0141, 0.0141, 0.0141,
		0.0141, 0.0141, 0.0141, 0.0141, 0.0141, 0.0141, 0.0141, 0.0141, 0.0141, 0.0141, 0.0141, 0.0141, 0.0141, 0.0141, 0.0141, 0.0141, 0.0141, 0.0141,
		0.0141, 0.0141, 0.0141, 0.0141, 0.0141, 0.0141, 0.0141, 0.0113, 0.00809, 0.00281, 0.00143, 0.0000451, 0.0000771, 0.0000771, 0.0000903, 0.0000962,
		0.0000962, 0.0000962, 0.0000962, 0.0000962, 0.0000962, 0.0000962, 0.0000962, 0.0000962, 0.0000890, 0.0000890, 0.0000766, 0.0000766, 0.0000766,
		0.0000766, 0.0000766, 0.0000766, 0.0000766, 0.0000766, 0.0000766, 0.0000766, 0.0000766, 0.0000766, 0.0000766, 0.0000766, 0.0000766, 0.0000766,
		0.000169, 0.000169, 0.000169, 0.000161, 0.000161, 0.00371, 0.00768, 0.0148, 0.0177, 0.0183, 0.0184, 0.0184, 0.0184, 0.0184, 0.0184, 0.0184, 0.0184,
		0.0184, 0.0184, 0.0184, 0.0184, 0.0184, 0.0184, 0.0184, 0.0184, 0.0184, 0.0184, 0.0184, 0.0184, 0.0184, 0.0184, 0.0184, 0.0183, 0.0183, 0.0183,
		0.0183, 0.0183, 0.0183, 0.0183, 0.0183, 0.0183, 0.0183, 0.0183, 0.0184, 0.0184, 0.0184, 0.0184, 0.0184, 0.0184, 0.0184, 0.0184, 0.0184, 0.0184,
		0.0184, 0.0184, 0.0184, 0.0184, 0.0184, 0.0184, 0.0184, 0.0183, 0.0183, 0.0183, 0.0183, 0.0183, 0.0147, 0.0107, 0.00362, 0.000764, 0.0000762,
		0.0000628, 0.0000526, 0.0000526, 0.0000526, 0.0000526, 0.0000360, 0.0000360, 0.0000360, 0.0000269, 0.0000269, 0.0000269, 0.0000269, 0.0000269,
		0.0000269, 0.0000269, 0.00420, 0.00747, 0.0138, 0.0147, 0.0147, 0.0147, 0.0147, 0.0147, 0.0147, 0.0147, 0.0147, 0.0147, 0.0147, 0.0147, 0.0147,
		0.0147, 0.0147, 0.0147, 0.0147, 0.0147, 0.0147, 0.0147, 0.0147, 0.0150, 0.0165, 0.0165, 0.0165, 0.0165, 0.0165, 0.0165, 0.0165, 0.0165, 0.0165,
		0.0165, 0.0165, 0.0165, 0.0165, 0.0165, 0.0165, 0.0165, 0.0165, 0.0165, 0.0165, 0.0165, 0.0165, 0.0165, 0.0165, 0.0165, 0.0165, 0.0165,
		0.0166, 0.0166, 0.0166, 0.0166, 0.0166, 0.0166, 0.0166, 0.0166, 0.0166, 0.0166, 0.0133, 0.0100, 0.00366, 0.00277, 0.00277, 0.00277, 0.00277,
		0.00275, 0.00275, 0.00275, 0.00275, 0.00275, 0.00276, 0.00276, 0.00276, 0.00276, 0.00276, 0.00276, 0.00276, 0.00278, 0.00516, 0.00926,
		0.0140, 0.0163, 0.0156, 0.0156, 0.0156, 0.0156, 0.0156, 0.0156, 0.0156, 0.0156, 0.0156, 0.0156, 0.0156, 0.0156, 0.0156, 0.0156, 0.0156,
		0.0156, 0.0186, 0.0228, 0.0272, 0.0294, 0.0302, 0.0302, 0.0302, 0.0302, 0.0302, 0.0302, 0.0302, 0.0302, 0.0302, 0.0302, 0.0302, 0.0302,
		0.0302, 0.0302, 0.0302, 0.0302, 0.0324, 0.0353, 0.0400, 0.0420, 0.0457, 0.0457, 0.0457, 0.0457, 0.0457, 0.0457, 0.0457, 0.0457, 0.0457,
		0.0457, 0.0457, 0.0457, 0.0457, 0.0457, 0.0457, 0.0456, 0.0433, 0.0392, 0.0345, 0.0320, 0.0311, 0.0311, 0.0311, 0.0311, 0.0310, 0.0310,
		0.0310, 0.0310, 0.0310, 0.0311, 0.0311, 0.0311, 0.0311, 0.0311, 0.0311, 0.0311, 0.0282, 0.0240, 0.0196, 0.0174, 0.0166, 0.0166, 0.0166,
		0.0166, 0.0166, 0.0166, 0.0166, 0.0166, 0.0166, 0.0166, 0.0166, 0.0166, 0.0166, 0.0166, 0.0166, 0.0166, 0.0134, 0.0105, 0.00581, 0.00387,
		0.000225, 0.000225, 0.000225, 0.000234, 0.000234, 0.000234, 0.000234, 0.000234, 0.000234, 0.000234, 0.000234, 0.000234, 0.000234, 0.000234,
		0.000234, 0.000234, 0.00356, 0.00657, 0.0108, 0.0132, 0.0138, 0.0138, 0.0138, 0.0138, 0.0138, 0.0138, 0.0138, 0.0138, 0.0138, 0.0138,
		0.0138, 0.0138, 0.0138, 0.0138, 0.0138, 0.0138, 0.0138, 0.0138, 0.0138, 0.0138, 0.0138, 0.0138, 0.0138, 0.0138, 0.0138, 0.0138, 0.0138,
		0.0138, 0.0138, 0.0138, 0.0138, 0.0138, 0.0138, 0.0138, 0.0138, 0.0138, 0.0138, 0.0138, 0.0138, 0.0137, 0.0137, 0.0137, 0.0137, 0.0137,
		0.0137, 0.0137, 0.0137, 0.0137, 0.0137, 0.0137, 0.0137, 0.0137, 0.0137, 0.0137, 0.0137, 0.0137, 0.0137, 0.0147, 0.0141, 0.0142, 0.0153,
		0.0153, 0.0153, 0.0153, 0.0153, 0.0153, 0.0153, 0.0153, 0.0153, 0.0153, 0.0153, 0.0153, 0.0153, 0.0153, 0.0153, 0.0153, 0.0153, 0.0153,
		0.0153, 0.0153, 0.0153, 0.0153, 0.0153, 0.0153, 0.0153, 0.0153, 0.0153, 0.0153, 0.0153, 0.0153, 0.0153, 0.0153, 0.0153, 0.0153, 0.0153,
		0.0153, 0.0153, 0.0153, 0.0153, 0.0153, 0.0153, 0.0153, 0.0153, 0.0153, 0.0153, 0.0153, 0.0153, 0.0153, 0.0153, 0.0153, 0.0153, 0.0153,
		0.0153, 0.0153, 0.0153, 0.0153, 0.0120, 0.00793, 0.00433, 0.00179, 0.000159, 0.000159, 0.000159, 0.000159, 0.000159, 0.000159,
		0.000165, 0.000165, 0.000165, 0.000165, 0.000165, 0.000139, 0.0000858, 0.0000745, 0.0000745, 0.0000745, 0.0000745, 0.0000745,
		0.0000745, 0.0000745, 0.00227, 0.00227, 0.00227, 0.00233, 0.00233, 0.00233, 0.00238, 0.00238, 0.00238, 0.00238, 0.00238, 0.00238,
		0.00238, 0.00238, 0.00238, 0.00238, 0.00238, 0.00238, 0.00238, 0.00238, 0.00520, 0.00520, 0.00520, 0.00520, 0.00521, 0.00522,
		0.00522, 0.00522, 0.00522, 0.00522, 0.00522, 0.00524, 0.00525, 0.00526, 0.00526, 0.00526, 0.0100, 0.0120, 0.0120, 0.0120, 0.0119,
		0.0119, 0.0119, 0.0119, 0.0119, 0.0119, 0.0119, 0.0119, 0.0119, 0.0119, 0.0120, 0.0120, 0.0120, 0.0120, 0.0120, 0.0120, 0.0120,
		0.0120, 0.0120, 0.0120, 0.00984, 0.00984, 0.00984, 0.00979, 0.00979, 0.00979, 0.00973, 0.00973, 0.00973, 0.00973, 0.00973, 0.00973,
		0.00973, 0.00973, 0.00973, 0.00973, 0.00973, 0.00973, 0.00973, 0.00973, 0.00691, 0.00691, 0.00694, 0.00706, 0.00705, 0.00705,
		0.00705, 0.00705, 0.00705, 0.00705, 0.00705, 0.00703, 0.00702, 0.00701, 0.00701, 0.00701, 0.00224, 0.000261, 0.000261, 0.000261,
		0.000261, 0.000261, 0.000261, 0.000261, 0.000261, 0.000261, 0.000261, 0.000261, 0.000261, 0.000261, 0.000183, 0.000183, 0.000183,
		0.000202, 0.000202, 0.000202, 0.00411, 0.00770, 0.0128, 0.0146, 0.0173, 0.0173, 0.0173, 0.0173, 0.0173, 0.0173, 0.0173, 0.0174,
		0.0174, 0.0174, 0.0174, 0.0174, 0.0174, 0.0174, 0.0174, 0.0174, 0.0175, 0.0175, 0.0175, 0.0180, 0.0211, 0.0220, 0.0219, 0.0218,
		0.0218, 0.0218, 0.0218, 0.0218, 0.0218, 0.0218, 0.0220}
)
