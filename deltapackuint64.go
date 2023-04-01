// This file is generated, do not modify directly
// Use 'go generate' to regenerate.

package intcomp

import "unsafe"

// AppendGroup_uint64 Binary packing of one block of `in`, starting from `initoffset`
// to out. Differential coding is applied first.
// Caller must give the proper `bitlen` of the block
func appendGroup_uint64[T uint64 | int64](dst []uint64, in *[64]T, initoffset T, bitlen int) []uint64 {
	switch bitlen {
	case 0:
		return deltapack64_0(initoffset, in, dst)
	case 1:
		return deltapack64_1(initoffset, in, dst)
	case 2:
		return deltapack64_2(initoffset, in, dst)
	case 3:
		return deltapack64_3(initoffset, in, dst)
	case 4:
		return deltapack64_4(initoffset, in, dst)
	case 5:
		return deltapack64_5(initoffset, in, dst)
	case 6:
		return deltapack64_6(initoffset, in, dst)
	case 7:
		return deltapack64_7(initoffset, in, dst)
	case 8:
		return deltapack64_8(initoffset, in, dst)
	case 9:
		return deltapack64_9(initoffset, in, dst)
	case 10:
		return deltapack64_10(initoffset, in, dst)
	case 11:
		return deltapack64_11(initoffset, in, dst)
	case 12:
		return deltapack64_12(initoffset, in, dst)
	case 13:
		return deltapack64_13(initoffset, in, dst)
	case 14:
		return deltapack64_14(initoffset, in, dst)
	case 15:
		return deltapack64_15(initoffset, in, dst)
	case 16:
		return deltapack64_16(initoffset, in, dst)
	case 17:
		return deltapack64_17(initoffset, in, dst)
	case 18:
		return deltapack64_18(initoffset, in, dst)
	case 19:
		return deltapack64_19(initoffset, in, dst)
	case 20:
		return deltapack64_20(initoffset, in, dst)
	case 21:
		return deltapack64_21(initoffset, in, dst)
	case 22:
		return deltapack64_22(initoffset, in, dst)
	case 23:
		return deltapack64_23(initoffset, in, dst)
	case 24:
		return deltapack64_24(initoffset, in, dst)
	case 25:
		return deltapack64_25(initoffset, in, dst)
	case 26:
		return deltapack64_26(initoffset, in, dst)
	case 27:
		return deltapack64_27(initoffset, in, dst)
	case 28:
		return deltapack64_28(initoffset, in, dst)
	case 29:
		return deltapack64_29(initoffset, in, dst)
	case 30:
		return deltapack64_30(initoffset, in, dst)
	case 31:
		return deltapack64_31(initoffset, in, dst)
	case 32:
		return deltapack64_32(initoffset, in, dst)
	case 33:
		return deltapack64_33(initoffset, in, dst)
	case 34:
		return deltapack64_34(initoffset, in, dst)
	case 35:
		return deltapack64_35(initoffset, in, dst)
	case 36:
		return deltapack64_36(initoffset, in, dst)
	case 37:
		return deltapack64_37(initoffset, in, dst)
	case 38:
		return deltapack64_38(initoffset, in, dst)
	case 39:
		return deltapack64_39(initoffset, in, dst)
	case 40:
		return deltapack64_40(initoffset, in, dst)
	case 41:
		return deltapack64_41(initoffset, in, dst)
	case 42:
		return deltapack64_42(initoffset, in, dst)
	case 43:
		return deltapack64_43(initoffset, in, dst)
	case 44:
		return deltapack64_44(initoffset, in, dst)
	case 45:
		return deltapack64_45(initoffset, in, dst)
	case 46:
		return deltapack64_46(initoffset, in, dst)
	case 47:
		return deltapack64_47(initoffset, in, dst)
	case 48:
		return deltapack64_48(initoffset, in, dst)
	case 49:
		return deltapack64_49(initoffset, in, dst)
	case 50:
		return deltapack64_50(initoffset, in, dst)
	case 51:
		return deltapack64_51(initoffset, in, dst)
	case 52:
		return deltapack64_52(initoffset, in, dst)
	case 53:
		return deltapack64_53(initoffset, in, dst)
	case 54:
		return deltapack64_54(initoffset, in, dst)
	case 55:
		return deltapack64_55(initoffset, in, dst)
	case 56:
		return deltapack64_56(initoffset, in, dst)
	case 57:
		return deltapack64_57(initoffset, in, dst)
	case 58:
		return deltapack64_58(initoffset, in, dst)
	case 59:
		return deltapack64_59(initoffset, in, dst)
	case 60:
		return deltapack64_60(initoffset, in, dst)
	case 61:
		return deltapack64_61(initoffset, in, dst)
	case 62:
		return deltapack64_62(initoffset, in, dst)
	case 63:
		return deltapack64_63(initoffset, in, dst)
	case 64:
		same := (*[64]uint64)(unsafe.Pointer(in))
		return append(dst, same[:]...)
	default:
		panic("unsupported bitlen")
	}
}

// deltaUnpack_uint64 Decoding operation for DeltaPack_uint64
func deltaUnpack_uint64[T uint64 | int64](initoffset T, in []uint64, out []T, bitlen int) {
	switch bitlen {
	case 0:
		deltaunpack64_0(initoffset, (*[0]uint64)(in), (*[64]T)(out))
	case 1:
		deltaunpack64_1(initoffset, (*[1]uint64)(in), (*[64]T)(out))
	case 2:
		deltaunpack64_2(initoffset, (*[2]uint64)(in), (*[64]T)(out))
	case 3:
		deltaunpack64_3(initoffset, (*[3]uint64)(in), (*[64]T)(out))
	case 4:
		deltaunpack64_4(initoffset, (*[4]uint64)(in), (*[64]T)(out))
	case 5:
		deltaunpack64_5(initoffset, (*[5]uint64)(in), (*[64]T)(out))
	case 6:
		deltaunpack64_6(initoffset, (*[6]uint64)(in), (*[64]T)(out))
	case 7:
		deltaunpack64_7(initoffset, (*[7]uint64)(in), (*[64]T)(out))
	case 8:
		deltaunpack64_8(initoffset, (*[8]uint64)(in), (*[64]T)(out))
	case 9:
		deltaunpack64_9(initoffset, (*[9]uint64)(in), (*[64]T)(out))
	case 10:
		deltaunpack64_10(initoffset, (*[10]uint64)(in), (*[64]T)(out))
	case 11:
		deltaunpack64_11(initoffset, (*[11]uint64)(in), (*[64]T)(out))
	case 12:
		deltaunpack64_12(initoffset, (*[12]uint64)(in), (*[64]T)(out))
	case 13:
		deltaunpack64_13(initoffset, (*[13]uint64)(in), (*[64]T)(out))
	case 14:
		deltaunpack64_14(initoffset, (*[14]uint64)(in), (*[64]T)(out))
	case 15:
		deltaunpack64_15(initoffset, (*[15]uint64)(in), (*[64]T)(out))
	case 16:
		deltaunpack64_16(initoffset, (*[16]uint64)(in), (*[64]T)(out))
	case 17:
		deltaunpack64_17(initoffset, (*[17]uint64)(in), (*[64]T)(out))
	case 18:
		deltaunpack64_18(initoffset, (*[18]uint64)(in), (*[64]T)(out))
	case 19:
		deltaunpack64_19(initoffset, (*[19]uint64)(in), (*[64]T)(out))
	case 20:
		deltaunpack64_20(initoffset, (*[20]uint64)(in), (*[64]T)(out))
	case 21:
		deltaunpack64_21(initoffset, (*[21]uint64)(in), (*[64]T)(out))
	case 22:
		deltaunpack64_22(initoffset, (*[22]uint64)(in), (*[64]T)(out))
	case 23:
		deltaunpack64_23(initoffset, (*[23]uint64)(in), (*[64]T)(out))
	case 24:
		deltaunpack64_24(initoffset, (*[24]uint64)(in), (*[64]T)(out))
	case 25:
		deltaunpack64_25(initoffset, (*[25]uint64)(in), (*[64]T)(out))
	case 26:
		deltaunpack64_26(initoffset, (*[26]uint64)(in), (*[64]T)(out))
	case 27:
		deltaunpack64_27(initoffset, (*[27]uint64)(in), (*[64]T)(out))
	case 28:
		deltaunpack64_28(initoffset, (*[28]uint64)(in), (*[64]T)(out))
	case 29:
		deltaunpack64_29(initoffset, (*[29]uint64)(in), (*[64]T)(out))
	case 30:
		deltaunpack64_30(initoffset, (*[30]uint64)(in), (*[64]T)(out))
	case 31:
		deltaunpack64_31(initoffset, (*[31]uint64)(in), (*[64]T)(out))
	case 32:
		deltaunpack64_32(initoffset, (*[32]uint64)(in), (*[64]T)(out))
	case 33:
		deltaunpack64_33(initoffset, (*[33]uint64)(in), (*[64]T)(out))
	case 34:
		deltaunpack64_34(initoffset, (*[34]uint64)(in), (*[64]T)(out))
	case 35:
		deltaunpack64_35(initoffset, (*[35]uint64)(in), (*[64]T)(out))
	case 36:
		deltaunpack64_36(initoffset, (*[36]uint64)(in), (*[64]T)(out))
	case 37:
		deltaunpack64_37(initoffset, (*[37]uint64)(in), (*[64]T)(out))
	case 38:
		deltaunpack64_38(initoffset, (*[38]uint64)(in), (*[64]T)(out))
	case 39:
		deltaunpack64_39(initoffset, (*[39]uint64)(in), (*[64]T)(out))
	case 40:
		deltaunpack64_40(initoffset, (*[40]uint64)(in), (*[64]T)(out))
	case 41:
		deltaunpack64_41(initoffset, (*[41]uint64)(in), (*[64]T)(out))
	case 42:
		deltaunpack64_42(initoffset, (*[42]uint64)(in), (*[64]T)(out))
	case 43:
		deltaunpack64_43(initoffset, (*[43]uint64)(in), (*[64]T)(out))
	case 44:
		deltaunpack64_44(initoffset, (*[44]uint64)(in), (*[64]T)(out))
	case 45:
		deltaunpack64_45(initoffset, (*[45]uint64)(in), (*[64]T)(out))
	case 46:
		deltaunpack64_46(initoffset, (*[46]uint64)(in), (*[64]T)(out))
	case 47:
		deltaunpack64_47(initoffset, (*[47]uint64)(in), (*[64]T)(out))
	case 48:
		deltaunpack64_48(initoffset, (*[48]uint64)(in), (*[64]T)(out))
	case 49:
		deltaunpack64_49(initoffset, (*[49]uint64)(in), (*[64]T)(out))
	case 50:
		deltaunpack64_50(initoffset, (*[50]uint64)(in), (*[64]T)(out))
	case 51:
		deltaunpack64_51(initoffset, (*[51]uint64)(in), (*[64]T)(out))
	case 52:
		deltaunpack64_52(initoffset, (*[52]uint64)(in), (*[64]T)(out))
	case 53:
		deltaunpack64_53(initoffset, (*[53]uint64)(in), (*[64]T)(out))
	case 54:
		deltaunpack64_54(initoffset, (*[54]uint64)(in), (*[64]T)(out))
	case 55:
		deltaunpack64_55(initoffset, (*[55]uint64)(in), (*[64]T)(out))
	case 56:
		deltaunpack64_56(initoffset, (*[56]uint64)(in), (*[64]T)(out))
	case 57:
		deltaunpack64_57(initoffset, (*[57]uint64)(in), (*[64]T)(out))
	case 58:
		deltaunpack64_58(initoffset, (*[58]uint64)(in), (*[64]T)(out))
	case 59:
		deltaunpack64_59(initoffset, (*[59]uint64)(in), (*[64]T)(out))
	case 60:
		deltaunpack64_60(initoffset, (*[60]uint64)(in), (*[64]T)(out))
	case 61:
		deltaunpack64_61(initoffset, (*[61]uint64)(in), (*[64]T)(out))
	case 62:
		deltaunpack64_62(initoffset, (*[62]uint64)(in), (*[64]T)(out))
	case 63:
		deltaunpack64_63(initoffset, (*[63]uint64)(in), (*[64]T)(out))
	case 64:
		*(*[64]T)(out) = *(*[64]T)(unsafe.Pointer((*[64]uint64)(in)))
	default:
		panic("unsupported bitlen")
	}
}

// --- zigzag

// AppendGroupZigZag_uint64 Binary packing of one block of `in`, starting from `initoffset`
// to out. Differential coding is applied first, the difference is zigzag encoded.
//
//	Caller must give the proper `bitlen` of the block
func appendGroupZigZag_uint64(dst []uint64, in *[64]uint64, initoffset uint64, bitlen int) []uint64 {
	switch bitlen {
	case 0:
		return deltapackzigzag64_0(initoffset, (*[64]uint64)(in), dst)
	case 1:
		return deltapackzigzag64_1(initoffset, (*[64]uint64)(in), dst)
	case 2:
		return deltapackzigzag64_2(initoffset, (*[64]uint64)(in), dst)
	case 3:
		return deltapackzigzag64_3(initoffset, (*[64]uint64)(in), dst)
	case 4:
		return deltapackzigzag64_4(initoffset, (*[64]uint64)(in), dst)
	case 5:
		return deltapackzigzag64_5(initoffset, (*[64]uint64)(in), dst)
	case 6:
		return deltapackzigzag64_6(initoffset, (*[64]uint64)(in), dst)
	case 7:
		return deltapackzigzag64_7(initoffset, (*[64]uint64)(in), dst)
	case 8:
		return deltapackzigzag64_8(initoffset, (*[64]uint64)(in), dst)
	case 9:
		return deltapackzigzag64_9(initoffset, (*[64]uint64)(in), dst)
	case 10:
		return deltapackzigzag64_10(initoffset, (*[64]uint64)(in), dst)
	case 11:
		return deltapackzigzag64_11(initoffset, (*[64]uint64)(in), dst)
	case 12:
		return deltapackzigzag64_12(initoffset, (*[64]uint64)(in), dst)
	case 13:
		return deltapackzigzag64_13(initoffset, (*[64]uint64)(in), dst)
	case 14:
		return deltapackzigzag64_14(initoffset, (*[64]uint64)(in), dst)
	case 15:
		return deltapackzigzag64_15(initoffset, (*[64]uint64)(in), dst)
	case 16:
		return deltapackzigzag64_16(initoffset, (*[64]uint64)(in), dst)
	case 17:
		return deltapackzigzag64_17(initoffset, (*[64]uint64)(in), dst)
	case 18:
		return deltapackzigzag64_18(initoffset, (*[64]uint64)(in), dst)
	case 19:
		return deltapackzigzag64_19(initoffset, (*[64]uint64)(in), dst)
	case 20:
		return deltapackzigzag64_20(initoffset, (*[64]uint64)(in), dst)
	case 21:
		return deltapackzigzag64_21(initoffset, (*[64]uint64)(in), dst)
	case 22:
		return deltapackzigzag64_22(initoffset, (*[64]uint64)(in), dst)
	case 23:
		return deltapackzigzag64_23(initoffset, (*[64]uint64)(in), dst)
	case 24:
		return deltapackzigzag64_24(initoffset, (*[64]uint64)(in), dst)
	case 25:
		return deltapackzigzag64_25(initoffset, (*[64]uint64)(in), dst)
	case 26:
		return deltapackzigzag64_26(initoffset, (*[64]uint64)(in), dst)
	case 27:
		return deltapackzigzag64_27(initoffset, (*[64]uint64)(in), dst)
	case 28:
		return deltapackzigzag64_28(initoffset, (*[64]uint64)(in), dst)
	case 29:
		return deltapackzigzag64_29(initoffset, (*[64]uint64)(in), dst)
	case 30:
		return deltapackzigzag64_30(initoffset, (*[64]uint64)(in), dst)
	case 31:
		return deltapackzigzag64_31(initoffset, (*[64]uint64)(in), dst)
	case 32:
		return deltapackzigzag64_32(initoffset, (*[64]uint64)(in), dst)
	case 33:
		return deltapackzigzag64_33(initoffset, (*[64]uint64)(in), dst)
	case 34:
		return deltapackzigzag64_34(initoffset, (*[64]uint64)(in), dst)
	case 35:
		return deltapackzigzag64_35(initoffset, (*[64]uint64)(in), dst)
	case 36:
		return deltapackzigzag64_36(initoffset, (*[64]uint64)(in), dst)
	case 37:
		return deltapackzigzag64_37(initoffset, (*[64]uint64)(in), dst)
	case 38:
		return deltapackzigzag64_38(initoffset, (*[64]uint64)(in), dst)
	case 39:
		return deltapackzigzag64_39(initoffset, (*[64]uint64)(in), dst)
	case 40:
		return deltapackzigzag64_40(initoffset, (*[64]uint64)(in), dst)
	case 41:
		return deltapackzigzag64_41(initoffset, (*[64]uint64)(in), dst)
	case 42:
		return deltapackzigzag64_42(initoffset, (*[64]uint64)(in), dst)
	case 43:
		return deltapackzigzag64_43(initoffset, (*[64]uint64)(in), dst)
	case 44:
		return deltapackzigzag64_44(initoffset, (*[64]uint64)(in), dst)
	case 45:
		return deltapackzigzag64_45(initoffset, (*[64]uint64)(in), dst)
	case 46:
		return deltapackzigzag64_46(initoffset, (*[64]uint64)(in), dst)
	case 47:
		return deltapackzigzag64_47(initoffset, (*[64]uint64)(in), dst)
	case 48:
		return deltapackzigzag64_48(initoffset, (*[64]uint64)(in), dst)
	case 49:
		return deltapackzigzag64_49(initoffset, (*[64]uint64)(in), dst)
	case 50:
		return deltapackzigzag64_50(initoffset, (*[64]uint64)(in), dst)
	case 51:
		return deltapackzigzag64_51(initoffset, (*[64]uint64)(in), dst)
	case 52:
		return deltapackzigzag64_52(initoffset, (*[64]uint64)(in), dst)
	case 53:
		return deltapackzigzag64_53(initoffset, (*[64]uint64)(in), dst)
	case 54:
		return deltapackzigzag64_54(initoffset, (*[64]uint64)(in), dst)
	case 55:
		return deltapackzigzag64_55(initoffset, (*[64]uint64)(in), dst)
	case 56:
		return deltapackzigzag64_56(initoffset, (*[64]uint64)(in), dst)
	case 57:
		return deltapackzigzag64_57(initoffset, (*[64]uint64)(in), dst)
	case 58:
		return deltapackzigzag64_58(initoffset, (*[64]uint64)(in), dst)
	case 59:
		return deltapackzigzag64_59(initoffset, (*[64]uint64)(in), dst)
	case 60:
		return deltapackzigzag64_60(initoffset, (*[64]uint64)(in), dst)
	case 61:
		return deltapackzigzag64_61(initoffset, (*[64]uint64)(in), dst)
	case 62:
		return deltapackzigzag64_62(initoffset, (*[64]uint64)(in), dst)
	case 63:
		return deltapackzigzag64_63(initoffset, (*[64]uint64)(in), dst)
	case 64:
		same := (*[64]uint64)(unsafe.Pointer(in))
		return append(dst, same[:]...)
	default:
		panic("unsupported bitlen")
	}
}

// deltaUnpackZigzag_uint64 Decoding operation for DeltaPackZigzag_uint64
func deltaUnpackZigzag_uint64(initoffset uint64, in []uint64, out []uint64, bitlen int) {
	switch bitlen {
	case 0:
		deltaunpackzigzag64_0(initoffset, (*[0]uint64)(in), (*[64]uint64)(out))
	case 1:
		deltaunpackzigzag64_1(initoffset, (*[1]uint64)(in), (*[64]uint64)(out))
	case 2:
		deltaunpackzigzag64_2(initoffset, (*[2]uint64)(in), (*[64]uint64)(out))
	case 3:
		deltaunpackzigzag64_3(initoffset, (*[3]uint64)(in), (*[64]uint64)(out))
	case 4:
		deltaunpackzigzag64_4(initoffset, (*[4]uint64)(in), (*[64]uint64)(out))
	case 5:
		deltaunpackzigzag64_5(initoffset, (*[5]uint64)(in), (*[64]uint64)(out))
	case 6:
		deltaunpackzigzag64_6(initoffset, (*[6]uint64)(in), (*[64]uint64)(out))
	case 7:
		deltaunpackzigzag64_7(initoffset, (*[7]uint64)(in), (*[64]uint64)(out))
	case 8:
		deltaunpackzigzag64_8(initoffset, (*[8]uint64)(in), (*[64]uint64)(out))
	case 9:
		deltaunpackzigzag64_9(initoffset, (*[9]uint64)(in), (*[64]uint64)(out))
	case 10:
		deltaunpackzigzag64_10(initoffset, (*[10]uint64)(in), (*[64]uint64)(out))
	case 11:
		deltaunpackzigzag64_11(initoffset, (*[11]uint64)(in), (*[64]uint64)(out))
	case 12:
		deltaunpackzigzag64_12(initoffset, (*[12]uint64)(in), (*[64]uint64)(out))
	case 13:
		deltaunpackzigzag64_13(initoffset, (*[13]uint64)(in), (*[64]uint64)(out))
	case 14:
		deltaunpackzigzag64_14(initoffset, (*[14]uint64)(in), (*[64]uint64)(out))
	case 15:
		deltaunpackzigzag64_15(initoffset, (*[15]uint64)(in), (*[64]uint64)(out))
	case 16:
		deltaunpackzigzag64_16(initoffset, (*[16]uint64)(in), (*[64]uint64)(out))
	case 17:
		deltaunpackzigzag64_17(initoffset, (*[17]uint64)(in), (*[64]uint64)(out))
	case 18:
		deltaunpackzigzag64_18(initoffset, (*[18]uint64)(in), (*[64]uint64)(out))
	case 19:
		deltaunpackzigzag64_19(initoffset, (*[19]uint64)(in), (*[64]uint64)(out))
	case 20:
		deltaunpackzigzag64_20(initoffset, (*[20]uint64)(in), (*[64]uint64)(out))
	case 21:
		deltaunpackzigzag64_21(initoffset, (*[21]uint64)(in), (*[64]uint64)(out))
	case 22:
		deltaunpackzigzag64_22(initoffset, (*[22]uint64)(in), (*[64]uint64)(out))
	case 23:
		deltaunpackzigzag64_23(initoffset, (*[23]uint64)(in), (*[64]uint64)(out))
	case 24:
		deltaunpackzigzag64_24(initoffset, (*[24]uint64)(in), (*[64]uint64)(out))
	case 25:
		deltaunpackzigzag64_25(initoffset, (*[25]uint64)(in), (*[64]uint64)(out))
	case 26:
		deltaunpackzigzag64_26(initoffset, (*[26]uint64)(in), (*[64]uint64)(out))
	case 27:
		deltaunpackzigzag64_27(initoffset, (*[27]uint64)(in), (*[64]uint64)(out))
	case 28:
		deltaunpackzigzag64_28(initoffset, (*[28]uint64)(in), (*[64]uint64)(out))
	case 29:
		deltaunpackzigzag64_29(initoffset, (*[29]uint64)(in), (*[64]uint64)(out))
	case 30:
		deltaunpackzigzag64_30(initoffset, (*[30]uint64)(in), (*[64]uint64)(out))
	case 31:
		deltaunpackzigzag64_31(initoffset, (*[31]uint64)(in), (*[64]uint64)(out))
	case 32:
		deltaunpackzigzag64_32(initoffset, (*[32]uint64)(in), (*[64]uint64)(out))
	case 33:
		deltaunpackzigzag64_33(initoffset, (*[33]uint64)(in), (*[64]uint64)(out))
	case 34:
		deltaunpackzigzag64_34(initoffset, (*[34]uint64)(in), (*[64]uint64)(out))
	case 35:
		deltaunpackzigzag64_35(initoffset, (*[35]uint64)(in), (*[64]uint64)(out))
	case 36:
		deltaunpackzigzag64_36(initoffset, (*[36]uint64)(in), (*[64]uint64)(out))
	case 37:
		deltaunpackzigzag64_37(initoffset, (*[37]uint64)(in), (*[64]uint64)(out))
	case 38:
		deltaunpackzigzag64_38(initoffset, (*[38]uint64)(in), (*[64]uint64)(out))
	case 39:
		deltaunpackzigzag64_39(initoffset, (*[39]uint64)(in), (*[64]uint64)(out))
	case 40:
		deltaunpackzigzag64_40(initoffset, (*[40]uint64)(in), (*[64]uint64)(out))
	case 41:
		deltaunpackzigzag64_41(initoffset, (*[41]uint64)(in), (*[64]uint64)(out))
	case 42:
		deltaunpackzigzag64_42(initoffset, (*[42]uint64)(in), (*[64]uint64)(out))
	case 43:
		deltaunpackzigzag64_43(initoffset, (*[43]uint64)(in), (*[64]uint64)(out))
	case 44:
		deltaunpackzigzag64_44(initoffset, (*[44]uint64)(in), (*[64]uint64)(out))
	case 45:
		deltaunpackzigzag64_45(initoffset, (*[45]uint64)(in), (*[64]uint64)(out))
	case 46:
		deltaunpackzigzag64_46(initoffset, (*[46]uint64)(in), (*[64]uint64)(out))
	case 47:
		deltaunpackzigzag64_47(initoffset, (*[47]uint64)(in), (*[64]uint64)(out))
	case 48:
		deltaunpackzigzag64_48(initoffset, (*[48]uint64)(in), (*[64]uint64)(out))
	case 49:
		deltaunpackzigzag64_49(initoffset, (*[49]uint64)(in), (*[64]uint64)(out))
	case 50:
		deltaunpackzigzag64_50(initoffset, (*[50]uint64)(in), (*[64]uint64)(out))
	case 51:
		deltaunpackzigzag64_51(initoffset, (*[51]uint64)(in), (*[64]uint64)(out))
	case 52:
		deltaunpackzigzag64_52(initoffset, (*[52]uint64)(in), (*[64]uint64)(out))
	case 53:
		deltaunpackzigzag64_53(initoffset, (*[53]uint64)(in), (*[64]uint64)(out))
	case 54:
		deltaunpackzigzag64_54(initoffset, (*[54]uint64)(in), (*[64]uint64)(out))
	case 55:
		deltaunpackzigzag64_55(initoffset, (*[55]uint64)(in), (*[64]uint64)(out))
	case 56:
		deltaunpackzigzag64_56(initoffset, (*[56]uint64)(in), (*[64]uint64)(out))
	case 57:
		deltaunpackzigzag64_57(initoffset, (*[57]uint64)(in), (*[64]uint64)(out))
	case 58:
		deltaunpackzigzag64_58(initoffset, (*[58]uint64)(in), (*[64]uint64)(out))
	case 59:
		deltaunpackzigzag64_59(initoffset, (*[59]uint64)(in), (*[64]uint64)(out))
	case 60:
		deltaunpackzigzag64_60(initoffset, (*[60]uint64)(in), (*[64]uint64)(out))
	case 61:
		deltaunpackzigzag64_61(initoffset, (*[61]uint64)(in), (*[64]uint64)(out))
	case 62:
		deltaunpackzigzag64_62(initoffset, (*[62]uint64)(in), (*[64]uint64)(out))
	case 63:
		deltaunpackzigzag64_63(initoffset, (*[63]uint64)(in), (*[64]uint64)(out))
	case 64:
		*(*[64]uint64)(out) = *(*[64]uint64)(unsafe.Pointer((*[64]uint64)(in)))
	default:
		panic("unsupported bitlen")
	}
}
