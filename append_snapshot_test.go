package intcomp

import (
	"reflect"
	"testing"
)

// A compressed prefix may still be consumed while another caller appends.
// Appending must preserve the entire original slice,including its trailer.
func checkAppendSnapshot[T ~int32 | ~uint32 | ~int64 | ~uint64, U ~uint32 | ~uint64](t *testing.T, compress func([]T, []U) []U, uncompress func([]U, []T) []T) {
	t.Helper()
	for _, size := range []int{3, 131, 259} {
		input := make([]T, size)
		for i := range input {
			input[i] = T(2*i + 1)
		}
		for _, extraSize := range []int{3, 129, 257} {
			extra := make([]T, extraSize)
			for i := range extra {
				extra[i] = T(2*(size+i) + 1)
			}
			original := compress(input, make([]U, 0, 4096))
			snapshot := append([]U(nil), original...)
			combined := compress(extra, original)
			if !reflect.DeepEqual(original, snapshot) {
				t.Errorf("append rewrote retained compressed slice:prefix=%d append=%d", size, extraSize)
				continue
			}
			if got := uncompress(original, nil); !reflect.DeepEqual(got, input) {
				t.Errorf("retained slice decoded incorrectly:prefix=%d append=%d", size, extraSize)
			}
			want := append(append([]T(nil), input...), extra...)
			if got := uncompress(combined, nil); !reflect.DeepEqual(got, want) {
				t.Errorf("combined slice decoded incorrectly:prefix=%d append=%d", size, extraSize)
			}
		}
	}
}
func TestAppendPreservesCompressedInt32(t *testing.T) {
	checkAppendSnapshot(t, CompressInt32, UncompressInt32)
}
func TestAppendPreservesCompressedUint32(t *testing.T) {
	checkAppendSnapshot(t, CompressUint32, UncompressUint32)
}
func TestAppendPreservesCompressedInt64(t *testing.T) {
	checkAppendSnapshot(t, CompressInt64, UncompressInt64)
}
func TestAppendPreservesCompressedUint64(t *testing.T) {
	checkAppendSnapshot(t, CompressUint64, UncompressUint64)
}
