// This file is generated, do not modify directly
// Use 'go generate' to regenerate.

package intcomp

import "unsafe"

// AppendGroup_uint32 Binary packing of one block of `in`, starting from `initoffset`
// to out. Differential coding is applied first.
// Caller must give the proper `bitlen` of the block
func appendGroup_uint32[T uint32 | int32](dst []uint32, in *[32]T, initoffset T, bitlen int) []uint32 {
	switch bitlen {
	case 0:
		return deltapack32_0(initoffset, in, dst)
	case 1:
		return deltapack32_1(initoffset, in, dst)
	case 2:
		return deltapack32_2(initoffset, in, dst)
	case 3:
		return deltapack32_3(initoffset, in, dst)
	case 4:
		return deltapack32_4(initoffset, in, dst)
	case 5:
		return deltapack32_5(initoffset, in, dst)
	case 6:
		return deltapack32_6(initoffset, in, dst)
	case 7:
		return deltapack32_7(initoffset, in, dst)
	case 8:
		return deltapack32_8(initoffset, in, dst)
	case 9:
		return deltapack32_9(initoffset, in, dst)
	case 10:
		return deltapack32_10(initoffset, in, dst)
	case 11:
		return deltapack32_11(initoffset, in, dst)
	case 12:
		return deltapack32_12(initoffset, in, dst)
	case 13:
		return deltapack32_13(initoffset, in, dst)
	case 14:
		return deltapack32_14(initoffset, in, dst)
	case 15:
		return deltapack32_15(initoffset, in, dst)
	case 16:
		return deltapack32_16(initoffset, in, dst)
	case 17:
		return deltapack32_17(initoffset, in, dst)
	case 18:
		return deltapack32_18(initoffset, in, dst)
	case 19:
		return deltapack32_19(initoffset, in, dst)
	case 20:
		return deltapack32_20(initoffset, in, dst)
	case 21:
		return deltapack32_21(initoffset, in, dst)
	case 22:
		return deltapack32_22(initoffset, in, dst)
	case 23:
		return deltapack32_23(initoffset, in, dst)
	case 24:
		return deltapack32_24(initoffset, in, dst)
	case 25:
		return deltapack32_25(initoffset, in, dst)
	case 26:
		return deltapack32_26(initoffset, in, dst)
	case 27:
		return deltapack32_27(initoffset, in, dst)
	case 28:
		return deltapack32_28(initoffset, in, dst)
	case 29:
		return deltapack32_29(initoffset, in, dst)
	case 30:
		return deltapack32_30(initoffset, in, dst)
	case 31:
		return deltapack32_31(initoffset, in, dst)
	case 32:
		same := (*[32]uint32)(unsafe.Pointer(in))
		return append(dst, same[:]...)
	default:
		panic("unsupported bitlen")
	}
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
	switch bitlen {
	case 0:
		return deltapackzigzag32_0(initoffset, (*[32]uint32)(in), dst)
	case 1:
		return deltapackzigzag32_1(initoffset, (*[32]uint32)(in), dst)
	case 2:
		return deltapackzigzag32_2(initoffset, (*[32]uint32)(in), dst)
	case 3:
		return deltapackzigzag32_3(initoffset, (*[32]uint32)(in), dst)
	case 4:
		return deltapackzigzag32_4(initoffset, (*[32]uint32)(in), dst)
	case 5:
		return deltapackzigzag32_5(initoffset, (*[32]uint32)(in), dst)
	case 6:
		return deltapackzigzag32_6(initoffset, (*[32]uint32)(in), dst)
	case 7:
		return deltapackzigzag32_7(initoffset, (*[32]uint32)(in), dst)
	case 8:
		return deltapackzigzag32_8(initoffset, (*[32]uint32)(in), dst)
	case 9:
		return deltapackzigzag32_9(initoffset, (*[32]uint32)(in), dst)
	case 10:
		return deltapackzigzag32_10(initoffset, (*[32]uint32)(in), dst)
	case 11:
		return deltapackzigzag32_11(initoffset, (*[32]uint32)(in), dst)
	case 12:
		return deltapackzigzag32_12(initoffset, (*[32]uint32)(in), dst)
	case 13:
		return deltapackzigzag32_13(initoffset, (*[32]uint32)(in), dst)
	case 14:
		return deltapackzigzag32_14(initoffset, (*[32]uint32)(in), dst)
	case 15:
		return deltapackzigzag32_15(initoffset, (*[32]uint32)(in), dst)
	case 16:
		return deltapackzigzag32_16(initoffset, (*[32]uint32)(in), dst)
	case 17:
		return deltapackzigzag32_17(initoffset, (*[32]uint32)(in), dst)
	case 18:
		return deltapackzigzag32_18(initoffset, (*[32]uint32)(in), dst)
	case 19:
		return deltapackzigzag32_19(initoffset, (*[32]uint32)(in), dst)
	case 20:
		return deltapackzigzag32_20(initoffset, (*[32]uint32)(in), dst)
	case 21:
		return deltapackzigzag32_21(initoffset, (*[32]uint32)(in), dst)
	case 22:
		return deltapackzigzag32_22(initoffset, (*[32]uint32)(in), dst)
	case 23:
		return deltapackzigzag32_23(initoffset, (*[32]uint32)(in), dst)
	case 24:
		return deltapackzigzag32_24(initoffset, (*[32]uint32)(in), dst)
	case 25:
		return deltapackzigzag32_25(initoffset, (*[32]uint32)(in), dst)
	case 26:
		return deltapackzigzag32_26(initoffset, (*[32]uint32)(in), dst)
	case 27:
		return deltapackzigzag32_27(initoffset, (*[32]uint32)(in), dst)
	case 28:
		return deltapackzigzag32_28(initoffset, (*[32]uint32)(in), dst)
	case 29:
		return deltapackzigzag32_29(initoffset, (*[32]uint32)(in), dst)
	case 30:
		return deltapackzigzag32_30(initoffset, (*[32]uint32)(in), dst)
	case 31:
		return deltapackzigzag32_31(initoffset, (*[32]uint32)(in), dst)
	case 32:
		same := (*[32]uint32)(unsafe.Pointer(in))
		return append(dst, same[:]...)
	default:
		panic("unsupported bitlen")
	}
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
