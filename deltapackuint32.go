// This file is generated, do not modify directly
// Use 'go generate' to regenerate.

package intcomp

import "unsafe"

// AppendGroup_uint32 Binary packing of one block of `in`, starting from `initoffset`
// to out. Differential coding is applied first.
// Caller must give the proper `bitlen` of the block
func appendGroup_uint32[T uint32 | int32](dst []uint32, in *[32]T, initoffset T, bitlen int) []uint32 {
	switch bitlen - 1 {
	case 1 - 1:
		return appendGroup32_1(dst, in, initoffset)
	case 2 - 1:
		return appendGroup32_2(dst, in, initoffset)
	case 3 - 1:
		return appendGroup32_3(dst, in, initoffset)
	case 4 - 1:
		return appendGroup32_4(dst, in, initoffset)
	case 5 - 1:
		return appendGroup32_5(dst, in, initoffset)
	case 6 - 1:
		return appendGroup32_6(dst, in, initoffset)
	case 7 - 1:
		return appendGroup32_7(dst, in, initoffset)
	case 8 - 1:
		return appendGroup32_8(dst, in, initoffset)
	case 9 - 1:
		return appendGroup32_9(dst, in, initoffset)
	case 10 - 1:
		return appendGroup32_10(dst, in, initoffset)
	case 11 - 1:
		return appendGroup32_11(dst, in, initoffset)
	case 12 - 1:
		return appendGroup32_12(dst, in, initoffset)
	case 13 - 1:
		return appendGroup32_13(dst, in, initoffset)
	case 14 - 1:
		return appendGroup32_14(dst, in, initoffset)
	case 15 - 1:
		return appendGroup32_15(dst, in, initoffset)
	case 16 - 1:
		return appendGroup32_16(dst, in, initoffset)
	case 17 - 1:
		return appendGroup32_17(dst, in, initoffset)
	case 18 - 1:
		return appendGroup32_18(dst, in, initoffset)
	case 19 - 1:
		return appendGroup32_19(dst, in, initoffset)
	case 20 - 1:
		return appendGroup32_20(dst, in, initoffset)
	case 21 - 1:
		return appendGroup32_21(dst, in, initoffset)
	case 22 - 1:
		return appendGroup32_22(dst, in, initoffset)
	case 23 - 1:
		return appendGroup32_23(dst, in, initoffset)
	case 24 - 1:
		return appendGroup32_24(dst, in, initoffset)
	case 25 - 1:
		return appendGroup32_25(dst, in, initoffset)
	case 26 - 1:
		return appendGroup32_26(dst, in, initoffset)
	case 27 - 1:
		return appendGroup32_27(dst, in, initoffset)
	case 28 - 1:
		return appendGroup32_28(dst, in, initoffset)
	case 29 - 1:
		return appendGroup32_29(dst, in, initoffset)
	case 30 - 1:
		return appendGroup32_30(dst, in, initoffset)
	case 31 - 1:
		return appendGroup32_31(dst, in, initoffset)
	case 32 - 1:
		same := (*[32]uint32)(unsafe.Pointer(in))
		return append(dst, same[:]...)
	}
	return dst
}

// deltaUnpack_uint32 Decoding operation for DeltaPack_uint32
func deltaUnpack_uint32[T uint32 | int32](initoffset T, in []uint32, out []T, bitlen int) {
	switch bitlen {
	case 0:
		deltaunpack32_0(initoffset, (*[0]uint32)(in), (*[32]T)(out))
	case 1:
		deltaunpack32_1(initoffset, (*[1]uint32)(in), (*[32]T)(out))
	case 2:
		deltaunpack32_2(initoffset, (*[2]uint32)(in), (*[32]T)(out))
	case 3:
		deltaunpack32_3(initoffset, (*[3]uint32)(in), (*[32]T)(out))
	case 4:
		deltaunpack32_4(initoffset, (*[4]uint32)(in), (*[32]T)(out))
	case 5:
		deltaunpack32_5(initoffset, (*[5]uint32)(in), (*[32]T)(out))
	case 6:
		deltaunpack32_6(initoffset, (*[6]uint32)(in), (*[32]T)(out))
	case 7:
		deltaunpack32_7(initoffset, (*[7]uint32)(in), (*[32]T)(out))
	case 8:
		deltaunpack32_8(initoffset, (*[8]uint32)(in), (*[32]T)(out))
	case 9:
		deltaunpack32_9(initoffset, (*[9]uint32)(in), (*[32]T)(out))
	case 10:
		deltaunpack32_10(initoffset, (*[10]uint32)(in), (*[32]T)(out))
	case 11:
		deltaunpack32_11(initoffset, (*[11]uint32)(in), (*[32]T)(out))
	case 12:
		deltaunpack32_12(initoffset, (*[12]uint32)(in), (*[32]T)(out))
	case 13:
		deltaunpack32_13(initoffset, (*[13]uint32)(in), (*[32]T)(out))
	case 14:
		deltaunpack32_14(initoffset, (*[14]uint32)(in), (*[32]T)(out))
	case 15:
		deltaunpack32_15(initoffset, (*[15]uint32)(in), (*[32]T)(out))
	case 16:
		deltaunpack32_16(initoffset, (*[16]uint32)(in), (*[32]T)(out))
	case 17:
		deltaunpack32_17(initoffset, (*[17]uint32)(in), (*[32]T)(out))
	case 18:
		deltaunpack32_18(initoffset, (*[18]uint32)(in), (*[32]T)(out))
	case 19:
		deltaunpack32_19(initoffset, (*[19]uint32)(in), (*[32]T)(out))
	case 20:
		deltaunpack32_20(initoffset, (*[20]uint32)(in), (*[32]T)(out))
	case 21:
		deltaunpack32_21(initoffset, (*[21]uint32)(in), (*[32]T)(out))
	case 22:
		deltaunpack32_22(initoffset, (*[22]uint32)(in), (*[32]T)(out))
	case 23:
		deltaunpack32_23(initoffset, (*[23]uint32)(in), (*[32]T)(out))
	case 24:
		deltaunpack32_24(initoffset, (*[24]uint32)(in), (*[32]T)(out))
	case 25:
		deltaunpack32_25(initoffset, (*[25]uint32)(in), (*[32]T)(out))
	case 26:
		deltaunpack32_26(initoffset, (*[26]uint32)(in), (*[32]T)(out))
	case 27:
		deltaunpack32_27(initoffset, (*[27]uint32)(in), (*[32]T)(out))
	case 28:
		deltaunpack32_28(initoffset, (*[28]uint32)(in), (*[32]T)(out))
	case 29:
		deltaunpack32_29(initoffset, (*[29]uint32)(in), (*[32]T)(out))
	case 30:
		deltaunpack32_30(initoffset, (*[30]uint32)(in), (*[32]T)(out))
	case 31:
		deltaunpack32_31(initoffset, (*[31]uint32)(in), (*[32]T)(out))
	case 32:
		*(*[32]T)(out) = *(*[32]T)(unsafe.Pointer((*[32]uint32)(in)))
	default:
		panic("unsupported bitlen")
	}
}

// --- zigzag

// AppendGroupZigZag_uint32 Binary packing of one block of `in`, starting from `initoffset`
// to out. Differential coding is applied first, the difference is zigzag encoded.
//
//	Caller must give the proper `bitlen` of the block
func appendGroupZigZag_uint32(dst []uint32, in *[32]uint32, initoffset uint32, bitlen int) []uint32 {
	switch bitlen - 1 {
	case 1 - 1:
		return appendGroupZigZag32_1(dst, in, initoffset)
	case 2 - 1:
		return appendGroupZigZag32_2(dst, in, initoffset)
	case 3 - 1:
		return appendGroupZigZag32_3(dst, in, initoffset)
	case 4 - 1:
		return appendGroupZigZag32_4(dst, in, initoffset)
	case 5 - 1:
		return appendGroupZigZag32_5(dst, in, initoffset)
	case 6 - 1:
		return appendGroupZigZag32_6(dst, in, initoffset)
	case 7 - 1:
		return appendGroupZigZag32_7(dst, in, initoffset)
	case 8 - 1:
		return appendGroupZigZag32_8(dst, in, initoffset)
	case 9 - 1:
		return appendGroupZigZag32_9(dst, in, initoffset)
	case 10 - 1:
		return appendGroupZigZag32_10(dst, in, initoffset)
	case 11 - 1:
		return appendGroupZigZag32_11(dst, in, initoffset)
	case 12 - 1:
		return appendGroupZigZag32_12(dst, in, initoffset)
	case 13 - 1:
		return appendGroupZigZag32_13(dst, in, initoffset)
	case 14 - 1:
		return appendGroupZigZag32_14(dst, in, initoffset)
	case 15 - 1:
		return appendGroupZigZag32_15(dst, in, initoffset)
	case 16 - 1:
		return appendGroupZigZag32_16(dst, in, initoffset)
	case 17 - 1:
		return appendGroupZigZag32_17(dst, in, initoffset)
	case 18 - 1:
		return appendGroupZigZag32_18(dst, in, initoffset)
	case 19 - 1:
		return appendGroupZigZag32_19(dst, in, initoffset)
	case 20 - 1:
		return appendGroupZigZag32_20(dst, in, initoffset)
	case 21 - 1:
		return appendGroupZigZag32_21(dst, in, initoffset)
	case 22 - 1:
		return appendGroupZigZag32_22(dst, in, initoffset)
	case 23 - 1:
		return appendGroupZigZag32_23(dst, in, initoffset)
	case 24 - 1:
		return appendGroupZigZag32_24(dst, in, initoffset)
	case 25 - 1:
		return appendGroupZigZag32_25(dst, in, initoffset)
	case 26 - 1:
		return appendGroupZigZag32_26(dst, in, initoffset)
	case 27 - 1:
		return appendGroupZigZag32_27(dst, in, initoffset)
	case 28 - 1:
		return appendGroupZigZag32_28(dst, in, initoffset)
	case 29 - 1:
		return appendGroupZigZag32_29(dst, in, initoffset)
	case 30 - 1:
		return appendGroupZigZag32_30(dst, in, initoffset)
	case 31 - 1:
		return appendGroupZigZag32_31(dst, in, initoffset)
	case 32 - 1:
		same := (*[32]uint32)(unsafe.Pointer(in))
		return append(dst, same[:]...)
	}
	return dst
}

// deltaUnpackZigzag_uint32 Decoding operation for DeltaPackZigzag_uint32
func deltaUnpackZigzag_uint32(initoffset uint32, in []uint32, out []uint32, bitlen int) {
	switch bitlen {
	case 0:
		deltaunpackzigzag32_0(initoffset, (*[0]uint32)(in), (*[32]uint32)(out))
	case 1:
		deltaunpackzigzag32_1(initoffset, (*[1]uint32)(in), (*[32]uint32)(out))
	case 2:
		deltaunpackzigzag32_2(initoffset, (*[2]uint32)(in), (*[32]uint32)(out))
	case 3:
		deltaunpackzigzag32_3(initoffset, (*[3]uint32)(in), (*[32]uint32)(out))
	case 4:
		deltaunpackzigzag32_4(initoffset, (*[4]uint32)(in), (*[32]uint32)(out))
	case 5:
		deltaunpackzigzag32_5(initoffset, (*[5]uint32)(in), (*[32]uint32)(out))
	case 6:
		deltaunpackzigzag32_6(initoffset, (*[6]uint32)(in), (*[32]uint32)(out))
	case 7:
		deltaunpackzigzag32_7(initoffset, (*[7]uint32)(in), (*[32]uint32)(out))
	case 8:
		deltaunpackzigzag32_8(initoffset, (*[8]uint32)(in), (*[32]uint32)(out))
	case 9:
		deltaunpackzigzag32_9(initoffset, (*[9]uint32)(in), (*[32]uint32)(out))
	case 10:
		deltaunpackzigzag32_10(initoffset, (*[10]uint32)(in), (*[32]uint32)(out))
	case 11:
		deltaunpackzigzag32_11(initoffset, (*[11]uint32)(in), (*[32]uint32)(out))
	case 12:
		deltaunpackzigzag32_12(initoffset, (*[12]uint32)(in), (*[32]uint32)(out))
	case 13:
		deltaunpackzigzag32_13(initoffset, (*[13]uint32)(in), (*[32]uint32)(out))
	case 14:
		deltaunpackzigzag32_14(initoffset, (*[14]uint32)(in), (*[32]uint32)(out))
	case 15:
		deltaunpackzigzag32_15(initoffset, (*[15]uint32)(in), (*[32]uint32)(out))
	case 16:
		deltaunpackzigzag32_16(initoffset, (*[16]uint32)(in), (*[32]uint32)(out))
	case 17:
		deltaunpackzigzag32_17(initoffset, (*[17]uint32)(in), (*[32]uint32)(out))
	case 18:
		deltaunpackzigzag32_18(initoffset, (*[18]uint32)(in), (*[32]uint32)(out))
	case 19:
		deltaunpackzigzag32_19(initoffset, (*[19]uint32)(in), (*[32]uint32)(out))
	case 20:
		deltaunpackzigzag32_20(initoffset, (*[20]uint32)(in), (*[32]uint32)(out))
	case 21:
		deltaunpackzigzag32_21(initoffset, (*[21]uint32)(in), (*[32]uint32)(out))
	case 22:
		deltaunpackzigzag32_22(initoffset, (*[22]uint32)(in), (*[32]uint32)(out))
	case 23:
		deltaunpackzigzag32_23(initoffset, (*[23]uint32)(in), (*[32]uint32)(out))
	case 24:
		deltaunpackzigzag32_24(initoffset, (*[24]uint32)(in), (*[32]uint32)(out))
	case 25:
		deltaunpackzigzag32_25(initoffset, (*[25]uint32)(in), (*[32]uint32)(out))
	case 26:
		deltaunpackzigzag32_26(initoffset, (*[26]uint32)(in), (*[32]uint32)(out))
	case 27:
		deltaunpackzigzag32_27(initoffset, (*[27]uint32)(in), (*[32]uint32)(out))
	case 28:
		deltaunpackzigzag32_28(initoffset, (*[28]uint32)(in), (*[32]uint32)(out))
	case 29:
		deltaunpackzigzag32_29(initoffset, (*[29]uint32)(in), (*[32]uint32)(out))
	case 30:
		deltaunpackzigzag32_30(initoffset, (*[30]uint32)(in), (*[32]uint32)(out))
	case 31:
		deltaunpackzigzag32_31(initoffset, (*[31]uint32)(in), (*[32]uint32)(out))
	case 32:
		*(*[32]uint32)(out) = *(*[32]uint32)(unsafe.Pointer((*[32]uint32)(in)))
	default:
		panic("unsupported bitlen")
	}
}
