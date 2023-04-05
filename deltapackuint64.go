// This file is generated, do not modify directly
// Use 'go generate' to regenerate.

package intcomp

import "unsafe"

// deltaPack_uint64 Binary packing of one block of `in`, starting from `initoffset`
// to out. Differential coding is applied first.
// Caller must give the proper `bitlen` of the block
func deltaPack_uint64[T uint64 | int64](out []uint64, in *[64]T, initoffset T, bitlen int) {
	switch bitlen {
	case 0:
		deltapack64_0((*[0]uint64)(out), in, initoffset)
	case 1:
		deltapack64_1((*[1]uint64)(out), in, initoffset)
	case 2:
		deltapack64_2((*[2]uint64)(out), in, initoffset)
	case 3:
		deltapack64_3((*[3]uint64)(out), in, initoffset)
	case 4:
		deltapack64_4((*[4]uint64)(out), in, initoffset)
	case 5:
		deltapack64_5((*[5]uint64)(out), in, initoffset)
	case 6:
		deltapack64_6((*[6]uint64)(out), in, initoffset)
	case 7:
		deltapack64_7((*[7]uint64)(out), in, initoffset)
	case 8:
		deltapack64_8((*[8]uint64)(out), in, initoffset)
	case 9:
		deltapack64_9((*[9]uint64)(out), in, initoffset)
	case 10:
		deltapack64_10((*[10]uint64)(out), in, initoffset)
	case 11:
		deltapack64_11((*[11]uint64)(out), in, initoffset)
	case 12:
		deltapack64_12((*[12]uint64)(out), in, initoffset)
	case 13:
		deltapack64_13((*[13]uint64)(out), in, initoffset)
	case 14:
		deltapack64_14((*[14]uint64)(out), in, initoffset)
	case 15:
		deltapack64_15((*[15]uint64)(out), in, initoffset)
	case 16:
		deltapack64_16((*[16]uint64)(out), in, initoffset)
	case 17:
		deltapack64_17((*[17]uint64)(out), in, initoffset)
	case 18:
		deltapack64_18((*[18]uint64)(out), in, initoffset)
	case 19:
		deltapack64_19((*[19]uint64)(out), in, initoffset)
	case 20:
		deltapack64_20((*[20]uint64)(out), in, initoffset)
	case 21:
		deltapack64_21((*[21]uint64)(out), in, initoffset)
	case 22:
		deltapack64_22((*[22]uint64)(out), in, initoffset)
	case 23:
		deltapack64_23((*[23]uint64)(out), in, initoffset)
	case 24:
		deltapack64_24((*[24]uint64)(out), in, initoffset)
	case 25:
		deltapack64_25((*[25]uint64)(out), in, initoffset)
	case 26:
		deltapack64_26((*[26]uint64)(out), in, initoffset)
	case 27:
		deltapack64_27((*[27]uint64)(out), in, initoffset)
	case 28:
		deltapack64_28((*[28]uint64)(out), in, initoffset)
	case 29:
		deltapack64_29((*[29]uint64)(out), in, initoffset)
	case 30:
		deltapack64_30((*[30]uint64)(out), in, initoffset)
	case 31:
		deltapack64_31((*[31]uint64)(out), in, initoffset)
	case 32:
		deltapack64_32((*[32]uint64)(out), in, initoffset)
	case 33:
		deltapack64_33((*[33]uint64)(out), in, initoffset)
	case 34:
		deltapack64_34((*[34]uint64)(out), in, initoffset)
	case 35:
		deltapack64_35((*[35]uint64)(out), in, initoffset)
	case 36:
		deltapack64_36((*[36]uint64)(out), in, initoffset)
	case 37:
		deltapack64_37((*[37]uint64)(out), in, initoffset)
	case 38:
		deltapack64_38((*[38]uint64)(out), in, initoffset)
	case 39:
		deltapack64_39((*[39]uint64)(out), in, initoffset)
	case 40:
		deltapack64_40((*[40]uint64)(out), in, initoffset)
	case 41:
		deltapack64_41((*[41]uint64)(out), in, initoffset)
	case 42:
		deltapack64_42((*[42]uint64)(out), in, initoffset)
	case 43:
		deltapack64_43((*[43]uint64)(out), in, initoffset)
	case 44:
		deltapack64_44((*[44]uint64)(out), in, initoffset)
	case 45:
		deltapack64_45((*[45]uint64)(out), in, initoffset)
	case 46:
		deltapack64_46((*[46]uint64)(out), in, initoffset)
	case 47:
		deltapack64_47((*[47]uint64)(out), in, initoffset)
	case 48:
		deltapack64_48((*[48]uint64)(out), in, initoffset)
	case 49:
		deltapack64_49((*[49]uint64)(out), in, initoffset)
	case 50:
		deltapack64_50((*[50]uint64)(out), in, initoffset)
	case 51:
		deltapack64_51((*[51]uint64)(out), in, initoffset)
	case 52:
		deltapack64_52((*[52]uint64)(out), in, initoffset)
	case 53:
		deltapack64_53((*[53]uint64)(out), in, initoffset)
	case 54:
		deltapack64_54((*[54]uint64)(out), in, initoffset)
	case 55:
		deltapack64_55((*[55]uint64)(out), in, initoffset)
	case 56:
		deltapack64_56((*[56]uint64)(out), in, initoffset)
	case 57:
		deltapack64_57((*[57]uint64)(out), in, initoffset)
	case 58:
		deltapack64_58((*[58]uint64)(out), in, initoffset)
	case 59:
		deltapack64_59((*[59]uint64)(out), in, initoffset)
	case 60:
		deltapack64_60((*[60]uint64)(out), in, initoffset)
	case 61:
		deltapack64_61((*[61]uint64)(out), in, initoffset)
	case 62:
		deltapack64_62((*[62]uint64)(out), in, initoffset)
	case 63:
		deltapack64_63((*[63]uint64)(out), in, initoffset)
	case 64:
		*(*[64]uint64)(out) = *((*[64]uint64)(unsafe.Pointer(in)))
	}
}

// deltaUnpack_uint64 Decoding operation for DeltaPack_uint64
func deltaUnpack_uint64[T uint64 | int64](out *[64]T, in []uint64, initoffset T, bitlen int) {
	switch bitlen {
	case 0:
		deltaunpack64_0(out, (*[0]uint64)(in), initoffset)
	case 1:
		deltaunpack64_1(out, (*[1]uint64)(in), initoffset)
	case 2:
		deltaunpack64_2(out, (*[2]uint64)(in), initoffset)
	case 3:
		deltaunpack64_3(out, (*[3]uint64)(in), initoffset)
	case 4:
		deltaunpack64_4(out, (*[4]uint64)(in), initoffset)
	case 5:
		deltaunpack64_5(out, (*[5]uint64)(in), initoffset)
	case 6:
		deltaunpack64_6(out, (*[6]uint64)(in), initoffset)
	case 7:
		deltaunpack64_7(out, (*[7]uint64)(in), initoffset)
	case 8:
		deltaunpack64_8(out, (*[8]uint64)(in), initoffset)
	case 9:
		deltaunpack64_9(out, (*[9]uint64)(in), initoffset)
	case 10:
		deltaunpack64_10(out, (*[10]uint64)(in), initoffset)
	case 11:
		deltaunpack64_11(out, (*[11]uint64)(in), initoffset)
	case 12:
		deltaunpack64_12(out, (*[12]uint64)(in), initoffset)
	case 13:
		deltaunpack64_13(out, (*[13]uint64)(in), initoffset)
	case 14:
		deltaunpack64_14(out, (*[14]uint64)(in), initoffset)
	case 15:
		deltaunpack64_15(out, (*[15]uint64)(in), initoffset)
	case 16:
		deltaunpack64_16(out, (*[16]uint64)(in), initoffset)
	case 17:
		deltaunpack64_17(out, (*[17]uint64)(in), initoffset)
	case 18:
		deltaunpack64_18(out, (*[18]uint64)(in), initoffset)
	case 19:
		deltaunpack64_19(out, (*[19]uint64)(in), initoffset)
	case 20:
		deltaunpack64_20(out, (*[20]uint64)(in), initoffset)
	case 21:
		deltaunpack64_21(out, (*[21]uint64)(in), initoffset)
	case 22:
		deltaunpack64_22(out, (*[22]uint64)(in), initoffset)
	case 23:
		deltaunpack64_23(out, (*[23]uint64)(in), initoffset)
	case 24:
		deltaunpack64_24(out, (*[24]uint64)(in), initoffset)
	case 25:
		deltaunpack64_25(out, (*[25]uint64)(in), initoffset)
	case 26:
		deltaunpack64_26(out, (*[26]uint64)(in), initoffset)
	case 27:
		deltaunpack64_27(out, (*[27]uint64)(in), initoffset)
	case 28:
		deltaunpack64_28(out, (*[28]uint64)(in), initoffset)
	case 29:
		deltaunpack64_29(out, (*[29]uint64)(in), initoffset)
	case 30:
		deltaunpack64_30(out, (*[30]uint64)(in), initoffset)
	case 31:
		deltaunpack64_31(out, (*[31]uint64)(in), initoffset)
	case 32:
		deltaunpack64_32(out, (*[32]uint64)(in), initoffset)
	case 33:
		deltaunpack64_33(out, (*[33]uint64)(in), initoffset)
	case 34:
		deltaunpack64_34(out, (*[34]uint64)(in), initoffset)
	case 35:
		deltaunpack64_35(out, (*[35]uint64)(in), initoffset)
	case 36:
		deltaunpack64_36(out, (*[36]uint64)(in), initoffset)
	case 37:
		deltaunpack64_37(out, (*[37]uint64)(in), initoffset)
	case 38:
		deltaunpack64_38(out, (*[38]uint64)(in), initoffset)
	case 39:
		deltaunpack64_39(out, (*[39]uint64)(in), initoffset)
	case 40:
		deltaunpack64_40(out, (*[40]uint64)(in), initoffset)
	case 41:
		deltaunpack64_41(out, (*[41]uint64)(in), initoffset)
	case 42:
		deltaunpack64_42(out, (*[42]uint64)(in), initoffset)
	case 43:
		deltaunpack64_43(out, (*[43]uint64)(in), initoffset)
	case 44:
		deltaunpack64_44(out, (*[44]uint64)(in), initoffset)
	case 45:
		deltaunpack64_45(out, (*[45]uint64)(in), initoffset)
	case 46:
		deltaunpack64_46(out, (*[46]uint64)(in), initoffset)
	case 47:
		deltaunpack64_47(out, (*[47]uint64)(in), initoffset)
	case 48:
		deltaunpack64_48(out, (*[48]uint64)(in), initoffset)
	case 49:
		deltaunpack64_49(out, (*[49]uint64)(in), initoffset)
	case 50:
		deltaunpack64_50(out, (*[50]uint64)(in), initoffset)
	case 51:
		deltaunpack64_51(out, (*[51]uint64)(in), initoffset)
	case 52:
		deltaunpack64_52(out, (*[52]uint64)(in), initoffset)
	case 53:
		deltaunpack64_53(out, (*[53]uint64)(in), initoffset)
	case 54:
		deltaunpack64_54(out, (*[54]uint64)(in), initoffset)
	case 55:
		deltaunpack64_55(out, (*[55]uint64)(in), initoffset)
	case 56:
		deltaunpack64_56(out, (*[56]uint64)(in), initoffset)
	case 57:
		deltaunpack64_57(out, (*[57]uint64)(in), initoffset)
	case 58:
		deltaunpack64_58(out, (*[58]uint64)(in), initoffset)
	case 59:
		deltaunpack64_59(out, (*[59]uint64)(in), initoffset)
	case 60:
		deltaunpack64_60(out, (*[60]uint64)(in), initoffset)
	case 61:
		deltaunpack64_61(out, (*[61]uint64)(in), initoffset)
	case 62:
		deltaunpack64_62(out, (*[62]uint64)(in), initoffset)
	case 63:
		deltaunpack64_63(out, (*[63]uint64)(in), initoffset)
	case 64:
		*out = *(*[64]T)(unsafe.Pointer((*[64]uint64)(in)))
	}
}

// --- zigzag

// deltaPackZigzag_uint64 Binary packing of one block of `in`, starting from `initoffset`
// to out. Differential coding is applied first, the difference is zigzag encoded.
//
//	Caller must give the proper `bitlen` of the block
func deltaPackZigzag_uint64(out []uint64, in *[64]uint64, initoffset uint64, bitlen int) {
	switch bitlen {
	case 0:
		deltapackzigzag64_0((*[0]uint64)(out), in, initoffset)
	case 1:
		deltapackzigzag64_1((*[1]uint64)(out), in, initoffset)
	case 2:
		deltapackzigzag64_2((*[2]uint64)(out), in, initoffset)
	case 3:
		deltapackzigzag64_3((*[3]uint64)(out), in, initoffset)
	case 4:
		deltapackzigzag64_4((*[4]uint64)(out), in, initoffset)
	case 5:
		deltapackzigzag64_5((*[5]uint64)(out), in, initoffset)
	case 6:
		deltapackzigzag64_6((*[6]uint64)(out), in, initoffset)
	case 7:
		deltapackzigzag64_7((*[7]uint64)(out), in, initoffset)
	case 8:
		deltapackzigzag64_8((*[8]uint64)(out), in, initoffset)
	case 9:
		deltapackzigzag64_9((*[9]uint64)(out), in, initoffset)
	case 10:
		deltapackzigzag64_10((*[10]uint64)(out), in, initoffset)
	case 11:
		deltapackzigzag64_11((*[11]uint64)(out), in, initoffset)
	case 12:
		deltapackzigzag64_12((*[12]uint64)(out), in, initoffset)
	case 13:
		deltapackzigzag64_13((*[13]uint64)(out), in, initoffset)
	case 14:
		deltapackzigzag64_14((*[14]uint64)(out), in, initoffset)
	case 15:
		deltapackzigzag64_15((*[15]uint64)(out), in, initoffset)
	case 16:
		deltapackzigzag64_16((*[16]uint64)(out), in, initoffset)
	case 17:
		deltapackzigzag64_17((*[17]uint64)(out), in, initoffset)
	case 18:
		deltapackzigzag64_18((*[18]uint64)(out), in, initoffset)
	case 19:
		deltapackzigzag64_19((*[19]uint64)(out), in, initoffset)
	case 20:
		deltapackzigzag64_20((*[20]uint64)(out), in, initoffset)
	case 21:
		deltapackzigzag64_21((*[21]uint64)(out), in, initoffset)
	case 22:
		deltapackzigzag64_22((*[22]uint64)(out), in, initoffset)
	case 23:
		deltapackzigzag64_23((*[23]uint64)(out), in, initoffset)
	case 24:
		deltapackzigzag64_24((*[24]uint64)(out), in, initoffset)
	case 25:
		deltapackzigzag64_25((*[25]uint64)(out), in, initoffset)
	case 26:
		deltapackzigzag64_26((*[26]uint64)(out), in, initoffset)
	case 27:
		deltapackzigzag64_27((*[27]uint64)(out), in, initoffset)
	case 28:
		deltapackzigzag64_28((*[28]uint64)(out), in, initoffset)
	case 29:
		deltapackzigzag64_29((*[29]uint64)(out), in, initoffset)
	case 30:
		deltapackzigzag64_30((*[30]uint64)(out), in, initoffset)
	case 31:
		deltapackzigzag64_31((*[31]uint64)(out), in, initoffset)
	case 32:
		deltapackzigzag64_32((*[32]uint64)(out), in, initoffset)
	case 33:
		deltapackzigzag64_33((*[33]uint64)(out), in, initoffset)
	case 34:
		deltapackzigzag64_34((*[34]uint64)(out), in, initoffset)
	case 35:
		deltapackzigzag64_35((*[35]uint64)(out), in, initoffset)
	case 36:
		deltapackzigzag64_36((*[36]uint64)(out), in, initoffset)
	case 37:
		deltapackzigzag64_37((*[37]uint64)(out), in, initoffset)
	case 38:
		deltapackzigzag64_38((*[38]uint64)(out), in, initoffset)
	case 39:
		deltapackzigzag64_39((*[39]uint64)(out), in, initoffset)
	case 40:
		deltapackzigzag64_40((*[40]uint64)(out), in, initoffset)
	case 41:
		deltapackzigzag64_41((*[41]uint64)(out), in, initoffset)
	case 42:
		deltapackzigzag64_42((*[42]uint64)(out), in, initoffset)
	case 43:
		deltapackzigzag64_43((*[43]uint64)(out), in, initoffset)
	case 44:
		deltapackzigzag64_44((*[44]uint64)(out), in, initoffset)
	case 45:
		deltapackzigzag64_45((*[45]uint64)(out), in, initoffset)
	case 46:
		deltapackzigzag64_46((*[46]uint64)(out), in, initoffset)
	case 47:
		deltapackzigzag64_47((*[47]uint64)(out), in, initoffset)
	case 48:
		deltapackzigzag64_48((*[48]uint64)(out), in, initoffset)
	case 49:
		deltapackzigzag64_49((*[49]uint64)(out), in, initoffset)
	case 50:
		deltapackzigzag64_50((*[50]uint64)(out), in, initoffset)
	case 51:
		deltapackzigzag64_51((*[51]uint64)(out), in, initoffset)
	case 52:
		deltapackzigzag64_52((*[52]uint64)(out), in, initoffset)
	case 53:
		deltapackzigzag64_53((*[53]uint64)(out), in, initoffset)
	case 54:
		deltapackzigzag64_54((*[54]uint64)(out), in, initoffset)
	case 55:
		deltapackzigzag64_55((*[55]uint64)(out), in, initoffset)
	case 56:
		deltapackzigzag64_56((*[56]uint64)(out), in, initoffset)
	case 57:
		deltapackzigzag64_57((*[57]uint64)(out), in, initoffset)
	case 58:
		deltapackzigzag64_58((*[58]uint64)(out), in, initoffset)
	case 59:
		deltapackzigzag64_59((*[59]uint64)(out), in, initoffset)
	case 60:
		deltapackzigzag64_60((*[60]uint64)(out), in, initoffset)
	case 61:
		deltapackzigzag64_61((*[61]uint64)(out), in, initoffset)
	case 62:
		deltapackzigzag64_62((*[62]uint64)(out), in, initoffset)
	case 63:
		deltapackzigzag64_63((*[63]uint64)(out), in, initoffset)
	case 64:
		*(*[64]uint64)(out) = *((*[64]uint64)(unsafe.Pointer(in)))
	}
}

// deltaUnpackZigzag_uint64 Decoding operation for DeltaPackZigzag_uint64
func deltaUnpackZigzag_uint64(out *[64]uint64, in []uint64, initoffset uint64, bitlen int) {
	switch bitlen {
	case 0:
		deltaunpackzigzag64_0(out, (*[0]uint64)(in), initoffset)
	case 1:
		deltaunpackzigzag64_1(out, (*[1]uint64)(in), initoffset)
	case 2:
		deltaunpackzigzag64_2(out, (*[2]uint64)(in), initoffset)
	case 3:
		deltaunpackzigzag64_3(out, (*[3]uint64)(in), initoffset)
	case 4:
		deltaunpackzigzag64_4(out, (*[4]uint64)(in), initoffset)
	case 5:
		deltaunpackzigzag64_5(out, (*[5]uint64)(in), initoffset)
	case 6:
		deltaunpackzigzag64_6(out, (*[6]uint64)(in), initoffset)
	case 7:
		deltaunpackzigzag64_7(out, (*[7]uint64)(in), initoffset)
	case 8:
		deltaunpackzigzag64_8(out, (*[8]uint64)(in), initoffset)
	case 9:
		deltaunpackzigzag64_9(out, (*[9]uint64)(in), initoffset)
	case 10:
		deltaunpackzigzag64_10(out, (*[10]uint64)(in), initoffset)
	case 11:
		deltaunpackzigzag64_11(out, (*[11]uint64)(in), initoffset)
	case 12:
		deltaunpackzigzag64_12(out, (*[12]uint64)(in), initoffset)
	case 13:
		deltaunpackzigzag64_13(out, (*[13]uint64)(in), initoffset)
	case 14:
		deltaunpackzigzag64_14(out, (*[14]uint64)(in), initoffset)
	case 15:
		deltaunpackzigzag64_15(out, (*[15]uint64)(in), initoffset)
	case 16:
		deltaunpackzigzag64_16(out, (*[16]uint64)(in), initoffset)
	case 17:
		deltaunpackzigzag64_17(out, (*[17]uint64)(in), initoffset)
	case 18:
		deltaunpackzigzag64_18(out, (*[18]uint64)(in), initoffset)
	case 19:
		deltaunpackzigzag64_19(out, (*[19]uint64)(in), initoffset)
	case 20:
		deltaunpackzigzag64_20(out, (*[20]uint64)(in), initoffset)
	case 21:
		deltaunpackzigzag64_21(out, (*[21]uint64)(in), initoffset)
	case 22:
		deltaunpackzigzag64_22(out, (*[22]uint64)(in), initoffset)
	case 23:
		deltaunpackzigzag64_23(out, (*[23]uint64)(in), initoffset)
	case 24:
		deltaunpackzigzag64_24(out, (*[24]uint64)(in), initoffset)
	case 25:
		deltaunpackzigzag64_25(out, (*[25]uint64)(in), initoffset)
	case 26:
		deltaunpackzigzag64_26(out, (*[26]uint64)(in), initoffset)
	case 27:
		deltaunpackzigzag64_27(out, (*[27]uint64)(in), initoffset)
	case 28:
		deltaunpackzigzag64_28(out, (*[28]uint64)(in), initoffset)
	case 29:
		deltaunpackzigzag64_29(out, (*[29]uint64)(in), initoffset)
	case 30:
		deltaunpackzigzag64_30(out, (*[30]uint64)(in), initoffset)
	case 31:
		deltaunpackzigzag64_31(out, (*[31]uint64)(in), initoffset)
	case 32:
		deltaunpackzigzag64_32(out, (*[32]uint64)(in), initoffset)
	case 33:
		deltaunpackzigzag64_33(out, (*[33]uint64)(in), initoffset)
	case 34:
		deltaunpackzigzag64_34(out, (*[34]uint64)(in), initoffset)
	case 35:
		deltaunpackzigzag64_35(out, (*[35]uint64)(in), initoffset)
	case 36:
		deltaunpackzigzag64_36(out, (*[36]uint64)(in), initoffset)
	case 37:
		deltaunpackzigzag64_37(out, (*[37]uint64)(in), initoffset)
	case 38:
		deltaunpackzigzag64_38(out, (*[38]uint64)(in), initoffset)
	case 39:
		deltaunpackzigzag64_39(out, (*[39]uint64)(in), initoffset)
	case 40:
		deltaunpackzigzag64_40(out, (*[40]uint64)(in), initoffset)
	case 41:
		deltaunpackzigzag64_41(out, (*[41]uint64)(in), initoffset)
	case 42:
		deltaunpackzigzag64_42(out, (*[42]uint64)(in), initoffset)
	case 43:
		deltaunpackzigzag64_43(out, (*[43]uint64)(in), initoffset)
	case 44:
		deltaunpackzigzag64_44(out, (*[44]uint64)(in), initoffset)
	case 45:
		deltaunpackzigzag64_45(out, (*[45]uint64)(in), initoffset)
	case 46:
		deltaunpackzigzag64_46(out, (*[46]uint64)(in), initoffset)
	case 47:
		deltaunpackzigzag64_47(out, (*[47]uint64)(in), initoffset)
	case 48:
		deltaunpackzigzag64_48(out, (*[48]uint64)(in), initoffset)
	case 49:
		deltaunpackzigzag64_49(out, (*[49]uint64)(in), initoffset)
	case 50:
		deltaunpackzigzag64_50(out, (*[50]uint64)(in), initoffset)
	case 51:
		deltaunpackzigzag64_51(out, (*[51]uint64)(in), initoffset)
	case 52:
		deltaunpackzigzag64_52(out, (*[52]uint64)(in), initoffset)
	case 53:
		deltaunpackzigzag64_53(out, (*[53]uint64)(in), initoffset)
	case 54:
		deltaunpackzigzag64_54(out, (*[54]uint64)(in), initoffset)
	case 55:
		deltaunpackzigzag64_55(out, (*[55]uint64)(in), initoffset)
	case 56:
		deltaunpackzigzag64_56(out, (*[56]uint64)(in), initoffset)
	case 57:
		deltaunpackzigzag64_57(out, (*[57]uint64)(in), initoffset)
	case 58:
		deltaunpackzigzag64_58(out, (*[58]uint64)(in), initoffset)
	case 59:
		deltaunpackzigzag64_59(out, (*[59]uint64)(in), initoffset)
	case 60:
		deltaunpackzigzag64_60(out, (*[60]uint64)(in), initoffset)
	case 61:
		deltaunpackzigzag64_61(out, (*[61]uint64)(in), initoffset)
	case 62:
		deltaunpackzigzag64_62(out, (*[62]uint64)(in), initoffset)
	case 63:
		deltaunpackzigzag64_63(out, (*[63]uint64)(in), initoffset)
	case 64:
		*out = *(*[64]uint64)(unsafe.Pointer((*[64]uint64)(in)))
	}
}
