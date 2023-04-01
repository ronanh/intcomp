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
		return appendGroup64_0(dst, in, initoffset)
	case 1:
		return appendGroup64_1(dst, in, initoffset)
	case 2:
		return appendGroup64_2(dst, in, initoffset)
	case 3:
		return appendGroup64_3(dst, in, initoffset)
	case 4:
		return appendGroup64_4(dst, in, initoffset)
	case 5:
		return appendGroup64_5(dst, in, initoffset)
	case 6:
		return appendGroup64_6(dst, in, initoffset)
	case 7:
		return appendGroup64_7(dst, in, initoffset)
	case 8:
		return appendGroup64_8(dst, in, initoffset)
	case 9:
		return appendGroup64_9(dst, in, initoffset)
	case 10:
		return appendGroup64_10(dst, in, initoffset)
	case 11:
		return appendGroup64_11(dst, in, initoffset)
	case 12:
		return appendGroup64_12(dst, in, initoffset)
	case 13:
		return appendGroup64_13(dst, in, initoffset)
	case 14:
		return appendGroup64_14(dst, in, initoffset)
	case 15:
		return appendGroup64_15(dst, in, initoffset)
	case 16:
		return appendGroup64_16(dst, in, initoffset)
	case 17:
		return appendGroup64_17(dst, in, initoffset)
	case 18:
		return appendGroup64_18(dst, in, initoffset)
	case 19:
		return appendGroup64_19(dst, in, initoffset)
	case 20:
		return appendGroup64_20(dst, in, initoffset)
	case 21:
		return appendGroup64_21(dst, in, initoffset)
	case 22:
		return appendGroup64_22(dst, in, initoffset)
	case 23:
		return appendGroup64_23(dst, in, initoffset)
	case 24:
		return appendGroup64_24(dst, in, initoffset)
	case 25:
		return appendGroup64_25(dst, in, initoffset)
	case 26:
		return appendGroup64_26(dst, in, initoffset)
	case 27:
		return appendGroup64_27(dst, in, initoffset)
	case 28:
		return appendGroup64_28(dst, in, initoffset)
	case 29:
		return appendGroup64_29(dst, in, initoffset)
	case 30:
		return appendGroup64_30(dst, in, initoffset)
	case 31:
		return appendGroup64_31(dst, in, initoffset)
	case 32:
		return appendGroup64_32(dst, in, initoffset)
	case 33:
		return appendGroup64_33(dst, in, initoffset)
	case 34:
		return appendGroup64_34(dst, in, initoffset)
	case 35:
		return appendGroup64_35(dst, in, initoffset)
	case 36:
		return appendGroup64_36(dst, in, initoffset)
	case 37:
		return appendGroup64_37(dst, in, initoffset)
	case 38:
		return appendGroup64_38(dst, in, initoffset)
	case 39:
		return appendGroup64_39(dst, in, initoffset)
	case 40:
		return appendGroup64_40(dst, in, initoffset)
	case 41:
		return appendGroup64_41(dst, in, initoffset)
	case 42:
		return appendGroup64_42(dst, in, initoffset)
	case 43:
		return appendGroup64_43(dst, in, initoffset)
	case 44:
		return appendGroup64_44(dst, in, initoffset)
	case 45:
		return appendGroup64_45(dst, in, initoffset)
	case 46:
		return appendGroup64_46(dst, in, initoffset)
	case 47:
		return appendGroup64_47(dst, in, initoffset)
	case 48:
		return appendGroup64_48(dst, in, initoffset)
	case 49:
		return appendGroup64_49(dst, in, initoffset)
	case 50:
		return appendGroup64_50(dst, in, initoffset)
	case 51:
		return appendGroup64_51(dst, in, initoffset)
	case 52:
		return appendGroup64_52(dst, in, initoffset)
	case 53:
		return appendGroup64_53(dst, in, initoffset)
	case 54:
		return appendGroup64_54(dst, in, initoffset)
	case 55:
		return appendGroup64_55(dst, in, initoffset)
	case 56:
		return appendGroup64_56(dst, in, initoffset)
	case 57:
		return appendGroup64_57(dst, in, initoffset)
	case 58:
		return appendGroup64_58(dst, in, initoffset)
	case 59:
		return appendGroup64_59(dst, in, initoffset)
	case 60:
		return appendGroup64_60(dst, in, initoffset)
	case 61:
		return appendGroup64_61(dst, in, initoffset)
	case 62:
		return appendGroup64_62(dst, in, initoffset)
	case 63:
		return appendGroup64_63(dst, in, initoffset)
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
		return appendGroupZigZag64_0(dst, in, initoffset)
	case 1:
		return appendGroupZigZag64_1(dst, in, initoffset)
	case 2:
		return appendGroupZigZag64_2(dst, in, initoffset)
	case 3:
		return appendGroupZigZag64_3(dst, in, initoffset)
	case 4:
		return appendGroupZigZag64_4(dst, in, initoffset)
	case 5:
		return appendGroupZigZag64_5(dst, in, initoffset)
	case 6:
		return appendGroupZigZag64_6(dst, in, initoffset)
	case 7:
		return appendGroupZigZag64_7(dst, in, initoffset)
	case 8:
		return appendGroupZigZag64_8(dst, in, initoffset)
	case 9:
		return appendGroupZigZag64_9(dst, in, initoffset)
	case 10:
		return appendGroupZigZag64_10(dst, in, initoffset)
	case 11:
		return appendGroupZigZag64_11(dst, in, initoffset)
	case 12:
		return appendGroupZigZag64_12(dst, in, initoffset)
	case 13:
		return appendGroupZigZag64_13(dst, in, initoffset)
	case 14:
		return appendGroupZigZag64_14(dst, in, initoffset)
	case 15:
		return appendGroupZigZag64_15(dst, in, initoffset)
	case 16:
		return appendGroupZigZag64_16(dst, in, initoffset)
	case 17:
		return appendGroupZigZag64_17(dst, in, initoffset)
	case 18:
		return appendGroupZigZag64_18(dst, in, initoffset)
	case 19:
		return appendGroupZigZag64_19(dst, in, initoffset)
	case 20:
		return appendGroupZigZag64_20(dst, in, initoffset)
	case 21:
		return appendGroupZigZag64_21(dst, in, initoffset)
	case 22:
		return appendGroupZigZag64_22(dst, in, initoffset)
	case 23:
		return appendGroupZigZag64_23(dst, in, initoffset)
	case 24:
		return appendGroupZigZag64_24(dst, in, initoffset)
	case 25:
		return appendGroupZigZag64_25(dst, in, initoffset)
	case 26:
		return appendGroupZigZag64_26(dst, in, initoffset)
	case 27:
		return appendGroupZigZag64_27(dst, in, initoffset)
	case 28:
		return appendGroupZigZag64_28(dst, in, initoffset)
	case 29:
		return appendGroupZigZag64_29(dst, in, initoffset)
	case 30:
		return appendGroupZigZag64_30(dst, in, initoffset)
	case 31:
		return appendGroupZigZag64_31(dst, in, initoffset)
	case 32:
		return appendGroupZigZag64_32(dst, in, initoffset)
	case 33:
		return appendGroupZigZag64_33(dst, in, initoffset)
	case 34:
		return appendGroupZigZag64_34(dst, in, initoffset)
	case 35:
		return appendGroupZigZag64_35(dst, in, initoffset)
	case 36:
		return appendGroupZigZag64_36(dst, in, initoffset)
	case 37:
		return appendGroupZigZag64_37(dst, in, initoffset)
	case 38:
		return appendGroupZigZag64_38(dst, in, initoffset)
	case 39:
		return appendGroupZigZag64_39(dst, in, initoffset)
	case 40:
		return appendGroupZigZag64_40(dst, in, initoffset)
	case 41:
		return appendGroupZigZag64_41(dst, in, initoffset)
	case 42:
		return appendGroupZigZag64_42(dst, in, initoffset)
	case 43:
		return appendGroupZigZag64_43(dst, in, initoffset)
	case 44:
		return appendGroupZigZag64_44(dst, in, initoffset)
	case 45:
		return appendGroupZigZag64_45(dst, in, initoffset)
	case 46:
		return appendGroupZigZag64_46(dst, in, initoffset)
	case 47:
		return appendGroupZigZag64_47(dst, in, initoffset)
	case 48:
		return appendGroupZigZag64_48(dst, in, initoffset)
	case 49:
		return appendGroupZigZag64_49(dst, in, initoffset)
	case 50:
		return appendGroupZigZag64_50(dst, in, initoffset)
	case 51:
		return appendGroupZigZag64_51(dst, in, initoffset)
	case 52:
		return appendGroupZigZag64_52(dst, in, initoffset)
	case 53:
		return appendGroupZigZag64_53(dst, in, initoffset)
	case 54:
		return appendGroupZigZag64_54(dst, in, initoffset)
	case 55:
		return appendGroupZigZag64_55(dst, in, initoffset)
	case 56:
		return appendGroupZigZag64_56(dst, in, initoffset)
	case 57:
		return appendGroupZigZag64_57(dst, in, initoffset)
	case 58:
		return appendGroupZigZag64_58(dst, in, initoffset)
	case 59:
		return appendGroupZigZag64_59(dst, in, initoffset)
	case 60:
		return appendGroupZigZag64_60(dst, in, initoffset)
	case 61:
		return appendGroupZigZag64_61(dst, in, initoffset)
	case 62:
		return appendGroupZigZag64_62(dst, in, initoffset)
	case 63:
		return appendGroupZigZag64_63(dst, in, initoffset)
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
