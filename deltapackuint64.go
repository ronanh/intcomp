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
		if len(out) >= 0 { // eliminate bounds check
			deltapack64_0((*[0]uint64)(out), in, initoffset)
		}
	case 1:
		if len(out) >= 1 { // eliminate bounds check
			deltapack64_1((*[1]uint64)(out), in, initoffset)
		}
	case 2:
		if len(out) >= 2 { // eliminate bounds check
			deltapack64_2((*[2]uint64)(out), in, initoffset)
		}
	case 3:
		if len(out) >= 3 { // eliminate bounds check
			deltapack64_3((*[3]uint64)(out), in, initoffset)
		}
	case 4:
		if len(out) >= 4 { // eliminate bounds check
			deltapack64_4((*[4]uint64)(out), in, initoffset)
		}
	case 5:
		if len(out) >= 5 { // eliminate bounds check
			deltapack64_5((*[5]uint64)(out), in, initoffset)
		}
	case 6:
		if len(out) >= 6 { // eliminate bounds check
			deltapack64_6((*[6]uint64)(out), in, initoffset)
		}
	case 7:
		if len(out) >= 7 { // eliminate bounds check
			deltapack64_7((*[7]uint64)(out), in, initoffset)
		}
	case 8:
		if len(out) >= 8 { // eliminate bounds check
			deltapack64_8((*[8]uint64)(out), in, initoffset)
		}
	case 9:
		if len(out) >= 9 { // eliminate bounds check
			deltapack64_9((*[9]uint64)(out), in, initoffset)
		}
	case 10:
		if len(out) >= 10 { // eliminate bounds check
			deltapack64_10((*[10]uint64)(out), in, initoffset)
		}
	case 11:
		if len(out) >= 11 { // eliminate bounds check
			deltapack64_11((*[11]uint64)(out), in, initoffset)
		}
	case 12:
		if len(out) >= 12 { // eliminate bounds check
			deltapack64_12((*[12]uint64)(out), in, initoffset)
		}
	case 13:
		if len(out) >= 13 { // eliminate bounds check
			deltapack64_13((*[13]uint64)(out), in, initoffset)
		}
	case 14:
		if len(out) >= 14 { // eliminate bounds check
			deltapack64_14((*[14]uint64)(out), in, initoffset)
		}
	case 15:
		if len(out) >= 15 { // eliminate bounds check
			deltapack64_15((*[15]uint64)(out), in, initoffset)
		}
	case 16:
		if len(out) >= 16 { // eliminate bounds check
			deltapack64_16((*[16]uint64)(out), in, initoffset)
		}
	case 17:
		if len(out) >= 17 { // eliminate bounds check
			deltapack64_17((*[17]uint64)(out), in, initoffset)
		}
	case 18:
		if len(out) >= 18 { // eliminate bounds check
			deltapack64_18((*[18]uint64)(out), in, initoffset)
		}
	case 19:
		if len(out) >= 19 { // eliminate bounds check
			deltapack64_19((*[19]uint64)(out), in, initoffset)
		}
	case 20:
		if len(out) >= 20 { // eliminate bounds check
			deltapack64_20((*[20]uint64)(out), in, initoffset)
		}
	case 21:
		if len(out) >= 21 { // eliminate bounds check
			deltapack64_21((*[21]uint64)(out), in, initoffset)
		}
	case 22:
		if len(out) >= 22 { // eliminate bounds check
			deltapack64_22((*[22]uint64)(out), in, initoffset)
		}
	case 23:
		if len(out) >= 23 { // eliminate bounds check
			deltapack64_23((*[23]uint64)(out), in, initoffset)
		}
	case 24:
		if len(out) >= 24 { // eliminate bounds check
			deltapack64_24((*[24]uint64)(out), in, initoffset)
		}
	case 25:
		if len(out) >= 25 { // eliminate bounds check
			deltapack64_25((*[25]uint64)(out), in, initoffset)
		}
	case 26:
		if len(out) >= 26 { // eliminate bounds check
			deltapack64_26((*[26]uint64)(out), in, initoffset)
		}
	case 27:
		if len(out) >= 27 { // eliminate bounds check
			deltapack64_27((*[27]uint64)(out), in, initoffset)
		}
	case 28:
		if len(out) >= 28 { // eliminate bounds check
			deltapack64_28((*[28]uint64)(out), in, initoffset)
		}
	case 29:
		if len(out) >= 29 { // eliminate bounds check
			deltapack64_29((*[29]uint64)(out), in, initoffset)
		}
	case 30:
		if len(out) >= 30 { // eliminate bounds check
			deltapack64_30((*[30]uint64)(out), in, initoffset)
		}
	case 31:
		if len(out) >= 31 { // eliminate bounds check
			deltapack64_31((*[31]uint64)(out), in, initoffset)
		}
	case 32:
		if len(out) >= 32 { // eliminate bounds check
			deltapack64_32((*[32]uint64)(out), in, initoffset)
		}
	case 33:
		if len(out) >= 33 { // eliminate bounds check
			deltapack64_33((*[33]uint64)(out), in, initoffset)
		}
	case 34:
		if len(out) >= 34 { // eliminate bounds check
			deltapack64_34((*[34]uint64)(out), in, initoffset)
		}
	case 35:
		if len(out) >= 35 { // eliminate bounds check
			deltapack64_35((*[35]uint64)(out), in, initoffset)
		}
	case 36:
		if len(out) >= 36 { // eliminate bounds check
			deltapack64_36((*[36]uint64)(out), in, initoffset)
		}
	case 37:
		if len(out) >= 37 { // eliminate bounds check
			deltapack64_37((*[37]uint64)(out), in, initoffset)
		}
	case 38:
		if len(out) >= 38 { // eliminate bounds check
			deltapack64_38((*[38]uint64)(out), in, initoffset)
		}
	case 39:
		if len(out) >= 39 { // eliminate bounds check
			deltapack64_39((*[39]uint64)(out), in, initoffset)
		}
	case 40:
		if len(out) >= 40 { // eliminate bounds check
			deltapack64_40((*[40]uint64)(out), in, initoffset)
		}
	case 41:
		if len(out) >= 41 { // eliminate bounds check
			deltapack64_41((*[41]uint64)(out), in, initoffset)
		}
	case 42:
		if len(out) >= 42 { // eliminate bounds check
			deltapack64_42((*[42]uint64)(out), in, initoffset)
		}
	case 43:
		if len(out) >= 43 { // eliminate bounds check
			deltapack64_43((*[43]uint64)(out), in, initoffset)
		}
	case 44:
		if len(out) >= 44 { // eliminate bounds check
			deltapack64_44((*[44]uint64)(out), in, initoffset)
		}
	case 45:
		if len(out) >= 45 { // eliminate bounds check
			deltapack64_45((*[45]uint64)(out), in, initoffset)
		}
	case 46:
		if len(out) >= 46 { // eliminate bounds check
			deltapack64_46((*[46]uint64)(out), in, initoffset)
		}
	case 47:
		if len(out) >= 47 { // eliminate bounds check
			deltapack64_47((*[47]uint64)(out), in, initoffset)
		}
	case 48:
		if len(out) >= 48 { // eliminate bounds check
			deltapack64_48((*[48]uint64)(out), in, initoffset)
		}
	case 49:
		if len(out) >= 49 { // eliminate bounds check
			deltapack64_49((*[49]uint64)(out), in, initoffset)
		}
	case 50:
		if len(out) >= 50 { // eliminate bounds check
			deltapack64_50((*[50]uint64)(out), in, initoffset)
		}
	case 51:
		if len(out) >= 51 { // eliminate bounds check
			deltapack64_51((*[51]uint64)(out), in, initoffset)
		}
	case 52:
		if len(out) >= 52 { // eliminate bounds check
			deltapack64_52((*[52]uint64)(out), in, initoffset)
		}
	case 53:
		if len(out) >= 53 { // eliminate bounds check
			deltapack64_53((*[53]uint64)(out), in, initoffset)
		}
	case 54:
		if len(out) >= 54 { // eliminate bounds check
			deltapack64_54((*[54]uint64)(out), in, initoffset)
		}
	case 55:
		if len(out) >= 55 { // eliminate bounds check
			deltapack64_55((*[55]uint64)(out), in, initoffset)
		}
	case 56:
		if len(out) >= 56 { // eliminate bounds check
			deltapack64_56((*[56]uint64)(out), in, initoffset)
		}
	case 57:
		if len(out) >= 57 { // eliminate bounds check
			deltapack64_57((*[57]uint64)(out), in, initoffset)
		}
	case 58:
		if len(out) >= 58 { // eliminate bounds check
			deltapack64_58((*[58]uint64)(out), in, initoffset)
		}
	case 59:
		if len(out) >= 59 { // eliminate bounds check
			deltapack64_59((*[59]uint64)(out), in, initoffset)
		}
	case 60:
		if len(out) >= 60 { // eliminate bounds check
			deltapack64_60((*[60]uint64)(out), in, initoffset)
		}
	case 61:
		if len(out) >= 61 { // eliminate bounds check
			deltapack64_61((*[61]uint64)(out), in, initoffset)
		}
	case 62:
		if len(out) >= 62 { // eliminate bounds check
			deltapack64_62((*[62]uint64)(out), in, initoffset)
		}
	case 63:
		if len(out) >= 63 { // eliminate bounds check
			deltapack64_63((*[63]uint64)(out), in, initoffset)
		}
	case 64:
		if len(out) >= 64 { // eliminate bounds check
			*(*[64]uint64)(out) = *((*[64]uint64)(unsafe.Pointer(in)))
		}
	}
}

// deltaUnpack_uint64 Decoding operation for DeltaPack_uint64
func deltaUnpack_uint64[T uint64 | int64](out *[64]T, in []uint64, initoffset T, bitlen int) {
	switch bitlen {
	case 0:
		if len(in) >= 0 { // eliminate bounds check
			deltaunpack64_0(out, (*[0]uint64)(in), initoffset)
		}
	case 1:
		if len(in) >= 1 { // eliminate bounds check
			deltaunpack64_1(out, (*[1]uint64)(in), initoffset)
		}
	case 2:
		if len(in) >= 2 { // eliminate bounds check
			deltaunpack64_2(out, (*[2]uint64)(in), initoffset)
		}
	case 3:
		if len(in) >= 3 { // eliminate bounds check
			deltaunpack64_3(out, (*[3]uint64)(in), initoffset)
		}
	case 4:
		if len(in) >= 4 { // eliminate bounds check
			deltaunpack64_4(out, (*[4]uint64)(in), initoffset)
		}
	case 5:
		if len(in) >= 5 { // eliminate bounds check
			deltaunpack64_5(out, (*[5]uint64)(in), initoffset)
		}
	case 6:
		if len(in) >= 6 { // eliminate bounds check
			deltaunpack64_6(out, (*[6]uint64)(in), initoffset)
		}
	case 7:
		if len(in) >= 7 { // eliminate bounds check
			deltaunpack64_7(out, (*[7]uint64)(in), initoffset)
		}
	case 8:
		if len(in) >= 8 { // eliminate bounds check
			deltaunpack64_8(out, (*[8]uint64)(in), initoffset)
		}
	case 9:
		if len(in) >= 9 { // eliminate bounds check
			deltaunpack64_9(out, (*[9]uint64)(in), initoffset)
		}
	case 10:
		if len(in) >= 10 { // eliminate bounds check
			deltaunpack64_10(out, (*[10]uint64)(in), initoffset)
		}
	case 11:
		if len(in) >= 11 { // eliminate bounds check
			deltaunpack64_11(out, (*[11]uint64)(in), initoffset)
		}
	case 12:
		if len(in) >= 12 { // eliminate bounds check
			deltaunpack64_12(out, (*[12]uint64)(in), initoffset)
		}
	case 13:
		if len(in) >= 13 { // eliminate bounds check
			deltaunpack64_13(out, (*[13]uint64)(in), initoffset)
		}
	case 14:
		if len(in) >= 14 { // eliminate bounds check
			deltaunpack64_14(out, (*[14]uint64)(in), initoffset)
		}
	case 15:
		if len(in) >= 15 { // eliminate bounds check
			deltaunpack64_15(out, (*[15]uint64)(in), initoffset)
		}
	case 16:
		if len(in) >= 16 { // eliminate bounds check
			deltaunpack64_16(out, (*[16]uint64)(in), initoffset)
		}
	case 17:
		if len(in) >= 17 { // eliminate bounds check
			deltaunpack64_17(out, (*[17]uint64)(in), initoffset)
		}
	case 18:
		if len(in) >= 18 { // eliminate bounds check
			deltaunpack64_18(out, (*[18]uint64)(in), initoffset)
		}
	case 19:
		if len(in) >= 19 { // eliminate bounds check
			deltaunpack64_19(out, (*[19]uint64)(in), initoffset)
		}
	case 20:
		if len(in) >= 20 { // eliminate bounds check
			deltaunpack64_20(out, (*[20]uint64)(in), initoffset)
		}
	case 21:
		if len(in) >= 21 { // eliminate bounds check
			deltaunpack64_21(out, (*[21]uint64)(in), initoffset)
		}
	case 22:
		if len(in) >= 22 { // eliminate bounds check
			deltaunpack64_22(out, (*[22]uint64)(in), initoffset)
		}
	case 23:
		if len(in) >= 23 { // eliminate bounds check
			deltaunpack64_23(out, (*[23]uint64)(in), initoffset)
		}
	case 24:
		if len(in) >= 24 { // eliminate bounds check
			deltaunpack64_24(out, (*[24]uint64)(in), initoffset)
		}
	case 25:
		if len(in) >= 25 { // eliminate bounds check
			deltaunpack64_25(out, (*[25]uint64)(in), initoffset)
		}
	case 26:
		if len(in) >= 26 { // eliminate bounds check
			deltaunpack64_26(out, (*[26]uint64)(in), initoffset)
		}
	case 27:
		if len(in) >= 27 { // eliminate bounds check
			deltaunpack64_27(out, (*[27]uint64)(in), initoffset)
		}
	case 28:
		if len(in) >= 28 { // eliminate bounds check
			deltaunpack64_28(out, (*[28]uint64)(in), initoffset)
		}
	case 29:
		if len(in) >= 29 { // eliminate bounds check
			deltaunpack64_29(out, (*[29]uint64)(in), initoffset)
		}
	case 30:
		if len(in) >= 30 { // eliminate bounds check
			deltaunpack64_30(out, (*[30]uint64)(in), initoffset)
		}
	case 31:
		if len(in) >= 31 { // eliminate bounds check
			deltaunpack64_31(out, (*[31]uint64)(in), initoffset)
		}
	case 32:
		if len(in) >= 32 { // eliminate bounds check
			deltaunpack64_32(out, (*[32]uint64)(in), initoffset)
		}
	case 33:
		if len(in) >= 33 { // eliminate bounds check
			deltaunpack64_33(out, (*[33]uint64)(in), initoffset)
		}
	case 34:
		if len(in) >= 34 { // eliminate bounds check
			deltaunpack64_34(out, (*[34]uint64)(in), initoffset)
		}
	case 35:
		if len(in) >= 35 { // eliminate bounds check
			deltaunpack64_35(out, (*[35]uint64)(in), initoffset)
		}
	case 36:
		if len(in) >= 36 { // eliminate bounds check
			deltaunpack64_36(out, (*[36]uint64)(in), initoffset)
		}
	case 37:
		if len(in) >= 37 { // eliminate bounds check
			deltaunpack64_37(out, (*[37]uint64)(in), initoffset)
		}
	case 38:
		if len(in) >= 38 { // eliminate bounds check
			deltaunpack64_38(out, (*[38]uint64)(in), initoffset)
		}
	case 39:
		if len(in) >= 39 { // eliminate bounds check
			deltaunpack64_39(out, (*[39]uint64)(in), initoffset)
		}
	case 40:
		if len(in) >= 40 { // eliminate bounds check
			deltaunpack64_40(out, (*[40]uint64)(in), initoffset)
		}
	case 41:
		if len(in) >= 41 { // eliminate bounds check
			deltaunpack64_41(out, (*[41]uint64)(in), initoffset)
		}
	case 42:
		if len(in) >= 42 { // eliminate bounds check
			deltaunpack64_42(out, (*[42]uint64)(in), initoffset)
		}
	case 43:
		if len(in) >= 43 { // eliminate bounds check
			deltaunpack64_43(out, (*[43]uint64)(in), initoffset)
		}
	case 44:
		if len(in) >= 44 { // eliminate bounds check
			deltaunpack64_44(out, (*[44]uint64)(in), initoffset)
		}
	case 45:
		if len(in) >= 45 { // eliminate bounds check
			deltaunpack64_45(out, (*[45]uint64)(in), initoffset)
		}
	case 46:
		if len(in) >= 46 { // eliminate bounds check
			deltaunpack64_46(out, (*[46]uint64)(in), initoffset)
		}
	case 47:
		if len(in) >= 47 { // eliminate bounds check
			deltaunpack64_47(out, (*[47]uint64)(in), initoffset)
		}
	case 48:
		if len(in) >= 48 { // eliminate bounds check
			deltaunpack64_48(out, (*[48]uint64)(in), initoffset)
		}
	case 49:
		if len(in) >= 49 { // eliminate bounds check
			deltaunpack64_49(out, (*[49]uint64)(in), initoffset)
		}
	case 50:
		if len(in) >= 50 { // eliminate bounds check
			deltaunpack64_50(out, (*[50]uint64)(in), initoffset)
		}
	case 51:
		if len(in) >= 51 { // eliminate bounds check
			deltaunpack64_51(out, (*[51]uint64)(in), initoffset)
		}
	case 52:
		if len(in) >= 52 { // eliminate bounds check
			deltaunpack64_52(out, (*[52]uint64)(in), initoffset)
		}
	case 53:
		if len(in) >= 53 { // eliminate bounds check
			deltaunpack64_53(out, (*[53]uint64)(in), initoffset)
		}
	case 54:
		if len(in) >= 54 { // eliminate bounds check
			deltaunpack64_54(out, (*[54]uint64)(in), initoffset)
		}
	case 55:
		if len(in) >= 55 { // eliminate bounds check
			deltaunpack64_55(out, (*[55]uint64)(in), initoffset)
		}
	case 56:
		if len(in) >= 56 { // eliminate bounds check
			deltaunpack64_56(out, (*[56]uint64)(in), initoffset)
		}
	case 57:
		if len(in) >= 57 { // eliminate bounds check
			deltaunpack64_57(out, (*[57]uint64)(in), initoffset)
		}
	case 58:
		if len(in) >= 58 { // eliminate bounds check
			deltaunpack64_58(out, (*[58]uint64)(in), initoffset)
		}
	case 59:
		if len(in) >= 59 { // eliminate bounds check
			deltaunpack64_59(out, (*[59]uint64)(in), initoffset)
		}
	case 60:
		if len(in) >= 60 { // eliminate bounds check
			deltaunpack64_60(out, (*[60]uint64)(in), initoffset)
		}
	case 61:
		if len(in) >= 61 { // eliminate bounds check
			deltaunpack64_61(out, (*[61]uint64)(in), initoffset)
		}
	case 62:
		if len(in) >= 62 { // eliminate bounds check
			deltaunpack64_62(out, (*[62]uint64)(in), initoffset)
		}
	case 63:
		if len(in) >= 63 { // eliminate bounds check
			deltaunpack64_63(out, (*[63]uint64)(in), initoffset)
		}
	case 64:
		if len(in) >= 64 { // eliminate bounds check
			*out = *(*[64]T)(unsafe.Pointer((*[64]uint64)(in)))
		}
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
		if len(out) >= 0 { // eliminate bounds check
			deltapackzigzag64_0((*[0]uint64)(out), in, initoffset)
		}
	case 1:
		if len(out) >= 1 { // eliminate bounds check
			deltapackzigzag64_1((*[1]uint64)(out), in, initoffset)
		}
	case 2:
		if len(out) >= 2 { // eliminate bounds check
			deltapackzigzag64_2((*[2]uint64)(out), in, initoffset)
		}
	case 3:
		if len(out) >= 3 { // eliminate bounds check
			deltapackzigzag64_3((*[3]uint64)(out), in, initoffset)
		}
	case 4:
		if len(out) >= 4 { // eliminate bounds check
			deltapackzigzag64_4((*[4]uint64)(out), in, initoffset)
		}
	case 5:
		if len(out) >= 5 { // eliminate bounds check
			deltapackzigzag64_5((*[5]uint64)(out), in, initoffset)
		}
	case 6:
		if len(out) >= 6 { // eliminate bounds check
			deltapackzigzag64_6((*[6]uint64)(out), in, initoffset)
		}
	case 7:
		if len(out) >= 7 { // eliminate bounds check
			deltapackzigzag64_7((*[7]uint64)(out), in, initoffset)
		}
	case 8:
		if len(out) >= 8 { // eliminate bounds check
			deltapackzigzag64_8((*[8]uint64)(out), in, initoffset)
		}
	case 9:
		if len(out) >= 9 { // eliminate bounds check
			deltapackzigzag64_9((*[9]uint64)(out), in, initoffset)
		}
	case 10:
		if len(out) >= 10 { // eliminate bounds check
			deltapackzigzag64_10((*[10]uint64)(out), in, initoffset)
		}
	case 11:
		if len(out) >= 11 { // eliminate bounds check
			deltapackzigzag64_11((*[11]uint64)(out), in, initoffset)
		}
	case 12:
		if len(out) >= 12 { // eliminate bounds check
			deltapackzigzag64_12((*[12]uint64)(out), in, initoffset)
		}
	case 13:
		if len(out) >= 13 { // eliminate bounds check
			deltapackzigzag64_13((*[13]uint64)(out), in, initoffset)
		}
	case 14:
		if len(out) >= 14 { // eliminate bounds check
			deltapackzigzag64_14((*[14]uint64)(out), in, initoffset)
		}
	case 15:
		if len(out) >= 15 { // eliminate bounds check
			deltapackzigzag64_15((*[15]uint64)(out), in, initoffset)
		}
	case 16:
		if len(out) >= 16 { // eliminate bounds check
			deltapackzigzag64_16((*[16]uint64)(out), in, initoffset)
		}
	case 17:
		if len(out) >= 17 { // eliminate bounds check
			deltapackzigzag64_17((*[17]uint64)(out), in, initoffset)
		}
	case 18:
		if len(out) >= 18 { // eliminate bounds check
			deltapackzigzag64_18((*[18]uint64)(out), in, initoffset)
		}
	case 19:
		if len(out) >= 19 { // eliminate bounds check
			deltapackzigzag64_19((*[19]uint64)(out), in, initoffset)
		}
	case 20:
		if len(out) >= 20 { // eliminate bounds check
			deltapackzigzag64_20((*[20]uint64)(out), in, initoffset)
		}
	case 21:
		if len(out) >= 21 { // eliminate bounds check
			deltapackzigzag64_21((*[21]uint64)(out), in, initoffset)
		}
	case 22:
		if len(out) >= 22 { // eliminate bounds check
			deltapackzigzag64_22((*[22]uint64)(out), in, initoffset)
		}
	case 23:
		if len(out) >= 23 { // eliminate bounds check
			deltapackzigzag64_23((*[23]uint64)(out), in, initoffset)
		}
	case 24:
		if len(out) >= 24 { // eliminate bounds check
			deltapackzigzag64_24((*[24]uint64)(out), in, initoffset)
		}
	case 25:
		if len(out) >= 25 { // eliminate bounds check
			deltapackzigzag64_25((*[25]uint64)(out), in, initoffset)
		}
	case 26:
		if len(out) >= 26 { // eliminate bounds check
			deltapackzigzag64_26((*[26]uint64)(out), in, initoffset)
		}
	case 27:
		if len(out) >= 27 { // eliminate bounds check
			deltapackzigzag64_27((*[27]uint64)(out), in, initoffset)
		}
	case 28:
		if len(out) >= 28 { // eliminate bounds check
			deltapackzigzag64_28((*[28]uint64)(out), in, initoffset)
		}
	case 29:
		if len(out) >= 29 { // eliminate bounds check
			deltapackzigzag64_29((*[29]uint64)(out), in, initoffset)
		}
	case 30:
		if len(out) >= 30 { // eliminate bounds check
			deltapackzigzag64_30((*[30]uint64)(out), in, initoffset)
		}
	case 31:
		if len(out) >= 31 { // eliminate bounds check
			deltapackzigzag64_31((*[31]uint64)(out), in, initoffset)
		}
	case 32:
		if len(out) >= 32 { // eliminate bounds check
			deltapackzigzag64_32((*[32]uint64)(out), in, initoffset)
		}
	case 33:
		if len(out) >= 33 { // eliminate bounds check
			deltapackzigzag64_33((*[33]uint64)(out), in, initoffset)
		}
	case 34:
		if len(out) >= 34 { // eliminate bounds check
			deltapackzigzag64_34((*[34]uint64)(out), in, initoffset)
		}
	case 35:
		if len(out) >= 35 { // eliminate bounds check
			deltapackzigzag64_35((*[35]uint64)(out), in, initoffset)
		}
	case 36:
		if len(out) >= 36 { // eliminate bounds check
			deltapackzigzag64_36((*[36]uint64)(out), in, initoffset)
		}
	case 37:
		if len(out) >= 37 { // eliminate bounds check
			deltapackzigzag64_37((*[37]uint64)(out), in, initoffset)
		}
	case 38:
		if len(out) >= 38 { // eliminate bounds check
			deltapackzigzag64_38((*[38]uint64)(out), in, initoffset)
		}
	case 39:
		if len(out) >= 39 { // eliminate bounds check
			deltapackzigzag64_39((*[39]uint64)(out), in, initoffset)
		}
	case 40:
		if len(out) >= 40 { // eliminate bounds check
			deltapackzigzag64_40((*[40]uint64)(out), in, initoffset)
		}
	case 41:
		if len(out) >= 41 { // eliminate bounds check
			deltapackzigzag64_41((*[41]uint64)(out), in, initoffset)
		}
	case 42:
		if len(out) >= 42 { // eliminate bounds check
			deltapackzigzag64_42((*[42]uint64)(out), in, initoffset)
		}
	case 43:
		if len(out) >= 43 { // eliminate bounds check
			deltapackzigzag64_43((*[43]uint64)(out), in, initoffset)
		}
	case 44:
		if len(out) >= 44 { // eliminate bounds check
			deltapackzigzag64_44((*[44]uint64)(out), in, initoffset)
		}
	case 45:
		if len(out) >= 45 { // eliminate bounds check
			deltapackzigzag64_45((*[45]uint64)(out), in, initoffset)
		}
	case 46:
		if len(out) >= 46 { // eliminate bounds check
			deltapackzigzag64_46((*[46]uint64)(out), in, initoffset)
		}
	case 47:
		if len(out) >= 47 { // eliminate bounds check
			deltapackzigzag64_47((*[47]uint64)(out), in, initoffset)
		}
	case 48:
		if len(out) >= 48 { // eliminate bounds check
			deltapackzigzag64_48((*[48]uint64)(out), in, initoffset)
		}
	case 49:
		if len(out) >= 49 { // eliminate bounds check
			deltapackzigzag64_49((*[49]uint64)(out), in, initoffset)
		}
	case 50:
		if len(out) >= 50 { // eliminate bounds check
			deltapackzigzag64_50((*[50]uint64)(out), in, initoffset)
		}
	case 51:
		if len(out) >= 51 { // eliminate bounds check
			deltapackzigzag64_51((*[51]uint64)(out), in, initoffset)
		}
	case 52:
		if len(out) >= 52 { // eliminate bounds check
			deltapackzigzag64_52((*[52]uint64)(out), in, initoffset)
		}
	case 53:
		if len(out) >= 53 { // eliminate bounds check
			deltapackzigzag64_53((*[53]uint64)(out), in, initoffset)
		}
	case 54:
		if len(out) >= 54 { // eliminate bounds check
			deltapackzigzag64_54((*[54]uint64)(out), in, initoffset)
		}
	case 55:
		if len(out) >= 55 { // eliminate bounds check
			deltapackzigzag64_55((*[55]uint64)(out), in, initoffset)
		}
	case 56:
		if len(out) >= 56 { // eliminate bounds check
			deltapackzigzag64_56((*[56]uint64)(out), in, initoffset)
		}
	case 57:
		if len(out) >= 57 { // eliminate bounds check
			deltapackzigzag64_57((*[57]uint64)(out), in, initoffset)
		}
	case 58:
		if len(out) >= 58 { // eliminate bounds check
			deltapackzigzag64_58((*[58]uint64)(out), in, initoffset)
		}
	case 59:
		if len(out) >= 59 { // eliminate bounds check
			deltapackzigzag64_59((*[59]uint64)(out), in, initoffset)
		}
	case 60:
		if len(out) >= 60 { // eliminate bounds check
			deltapackzigzag64_60((*[60]uint64)(out), in, initoffset)
		}
	case 61:
		if len(out) >= 61 { // eliminate bounds check
			deltapackzigzag64_61((*[61]uint64)(out), in, initoffset)
		}
	case 62:
		if len(out) >= 62 { // eliminate bounds check
			deltapackzigzag64_62((*[62]uint64)(out), in, initoffset)
		}
	case 63:
		if len(out) >= 63 { // eliminate bounds check
			deltapackzigzag64_63((*[63]uint64)(out), in, initoffset)
		}
	case 64:
		if len(out) >= 64 { // eliminate bounds check
			*(*[64]uint64)(out) = *((*[64]uint64)(unsafe.Pointer(in)))
		}
	}
}

// deltaUnpackZigzag_uint64 Decoding operation for DeltaPackZigzag_uint64
func deltaUnpackZigzag_uint64(out *[64]uint64, in []uint64, initoffset uint64, bitlen int) {
	switch bitlen {
	case 0:
		if len(in) >= 0 { // eliminate bounds check
			deltaunpackzigzag64_0(out, (*[0]uint64)(in), initoffset)
		}
	case 1:
		if len(in) >= 1 { // eliminate bounds check
			deltaunpackzigzag64_1(out, (*[1]uint64)(in), initoffset)
		}
	case 2:
		if len(in) >= 2 { // eliminate bounds check
			deltaunpackzigzag64_2(out, (*[2]uint64)(in), initoffset)
		}
	case 3:
		if len(in) >= 3 { // eliminate bounds check
			deltaunpackzigzag64_3(out, (*[3]uint64)(in), initoffset)
		}
	case 4:
		if len(in) >= 4 { // eliminate bounds check
			deltaunpackzigzag64_4(out, (*[4]uint64)(in), initoffset)
		}
	case 5:
		if len(in) >= 5 { // eliminate bounds check
			deltaunpackzigzag64_5(out, (*[5]uint64)(in), initoffset)
		}
	case 6:
		if len(in) >= 6 { // eliminate bounds check
			deltaunpackzigzag64_6(out, (*[6]uint64)(in), initoffset)
		}
	case 7:
		if len(in) >= 7 { // eliminate bounds check
			deltaunpackzigzag64_7(out, (*[7]uint64)(in), initoffset)
		}
	case 8:
		if len(in) >= 8 { // eliminate bounds check
			deltaunpackzigzag64_8(out, (*[8]uint64)(in), initoffset)
		}
	case 9:
		if len(in) >= 9 { // eliminate bounds check
			deltaunpackzigzag64_9(out, (*[9]uint64)(in), initoffset)
		}
	case 10:
		if len(in) >= 10 { // eliminate bounds check
			deltaunpackzigzag64_10(out, (*[10]uint64)(in), initoffset)
		}
	case 11:
		if len(in) >= 11 { // eliminate bounds check
			deltaunpackzigzag64_11(out, (*[11]uint64)(in), initoffset)
		}
	case 12:
		if len(in) >= 12 { // eliminate bounds check
			deltaunpackzigzag64_12(out, (*[12]uint64)(in), initoffset)
		}
	case 13:
		if len(in) >= 13 { // eliminate bounds check
			deltaunpackzigzag64_13(out, (*[13]uint64)(in), initoffset)
		}
	case 14:
		if len(in) >= 14 { // eliminate bounds check
			deltaunpackzigzag64_14(out, (*[14]uint64)(in), initoffset)
		}
	case 15:
		if len(in) >= 15 { // eliminate bounds check
			deltaunpackzigzag64_15(out, (*[15]uint64)(in), initoffset)
		}
	case 16:
		if len(in) >= 16 { // eliminate bounds check
			deltaunpackzigzag64_16(out, (*[16]uint64)(in), initoffset)
		}
	case 17:
		if len(in) >= 17 { // eliminate bounds check
			deltaunpackzigzag64_17(out, (*[17]uint64)(in), initoffset)
		}
	case 18:
		if len(in) >= 18 { // eliminate bounds check
			deltaunpackzigzag64_18(out, (*[18]uint64)(in), initoffset)
		}
	case 19:
		if len(in) >= 19 { // eliminate bounds check
			deltaunpackzigzag64_19(out, (*[19]uint64)(in), initoffset)
		}
	case 20:
		if len(in) >= 20 { // eliminate bounds check
			deltaunpackzigzag64_20(out, (*[20]uint64)(in), initoffset)
		}
	case 21:
		if len(in) >= 21 { // eliminate bounds check
			deltaunpackzigzag64_21(out, (*[21]uint64)(in), initoffset)
		}
	case 22:
		if len(in) >= 22 { // eliminate bounds check
			deltaunpackzigzag64_22(out, (*[22]uint64)(in), initoffset)
		}
	case 23:
		if len(in) >= 23 { // eliminate bounds check
			deltaunpackzigzag64_23(out, (*[23]uint64)(in), initoffset)
		}
	case 24:
		if len(in) >= 24 { // eliminate bounds check
			deltaunpackzigzag64_24(out, (*[24]uint64)(in), initoffset)
		}
	case 25:
		if len(in) >= 25 { // eliminate bounds check
			deltaunpackzigzag64_25(out, (*[25]uint64)(in), initoffset)
		}
	case 26:
		if len(in) >= 26 { // eliminate bounds check
			deltaunpackzigzag64_26(out, (*[26]uint64)(in), initoffset)
		}
	case 27:
		if len(in) >= 27 { // eliminate bounds check
			deltaunpackzigzag64_27(out, (*[27]uint64)(in), initoffset)
		}
	case 28:
		if len(in) >= 28 { // eliminate bounds check
			deltaunpackzigzag64_28(out, (*[28]uint64)(in), initoffset)
		}
	case 29:
		if len(in) >= 29 { // eliminate bounds check
			deltaunpackzigzag64_29(out, (*[29]uint64)(in), initoffset)
		}
	case 30:
		if len(in) >= 30 { // eliminate bounds check
			deltaunpackzigzag64_30(out, (*[30]uint64)(in), initoffset)
		}
	case 31:
		if len(in) >= 31 { // eliminate bounds check
			deltaunpackzigzag64_31(out, (*[31]uint64)(in), initoffset)
		}
	case 32:
		if len(in) >= 32 { // eliminate bounds check
			deltaunpackzigzag64_32(out, (*[32]uint64)(in), initoffset)
		}
	case 33:
		if len(in) >= 33 { // eliminate bounds check
			deltaunpackzigzag64_33(out, (*[33]uint64)(in), initoffset)
		}
	case 34:
		if len(in) >= 34 { // eliminate bounds check
			deltaunpackzigzag64_34(out, (*[34]uint64)(in), initoffset)
		}
	case 35:
		if len(in) >= 35 { // eliminate bounds check
			deltaunpackzigzag64_35(out, (*[35]uint64)(in), initoffset)
		}
	case 36:
		if len(in) >= 36 { // eliminate bounds check
			deltaunpackzigzag64_36(out, (*[36]uint64)(in), initoffset)
		}
	case 37:
		if len(in) >= 37 { // eliminate bounds check
			deltaunpackzigzag64_37(out, (*[37]uint64)(in), initoffset)
		}
	case 38:
		if len(in) >= 38 { // eliminate bounds check
			deltaunpackzigzag64_38(out, (*[38]uint64)(in), initoffset)
		}
	case 39:
		if len(in) >= 39 { // eliminate bounds check
			deltaunpackzigzag64_39(out, (*[39]uint64)(in), initoffset)
		}
	case 40:
		if len(in) >= 40 { // eliminate bounds check
			deltaunpackzigzag64_40(out, (*[40]uint64)(in), initoffset)
		}
	case 41:
		if len(in) >= 41 { // eliminate bounds check
			deltaunpackzigzag64_41(out, (*[41]uint64)(in), initoffset)
		}
	case 42:
		if len(in) >= 42 { // eliminate bounds check
			deltaunpackzigzag64_42(out, (*[42]uint64)(in), initoffset)
		}
	case 43:
		if len(in) >= 43 { // eliminate bounds check
			deltaunpackzigzag64_43(out, (*[43]uint64)(in), initoffset)
		}
	case 44:
		if len(in) >= 44 { // eliminate bounds check
			deltaunpackzigzag64_44(out, (*[44]uint64)(in), initoffset)
		}
	case 45:
		if len(in) >= 45 { // eliminate bounds check
			deltaunpackzigzag64_45(out, (*[45]uint64)(in), initoffset)
		}
	case 46:
		if len(in) >= 46 { // eliminate bounds check
			deltaunpackzigzag64_46(out, (*[46]uint64)(in), initoffset)
		}
	case 47:
		if len(in) >= 47 { // eliminate bounds check
			deltaunpackzigzag64_47(out, (*[47]uint64)(in), initoffset)
		}
	case 48:
		if len(in) >= 48 { // eliminate bounds check
			deltaunpackzigzag64_48(out, (*[48]uint64)(in), initoffset)
		}
	case 49:
		if len(in) >= 49 { // eliminate bounds check
			deltaunpackzigzag64_49(out, (*[49]uint64)(in), initoffset)
		}
	case 50:
		if len(in) >= 50 { // eliminate bounds check
			deltaunpackzigzag64_50(out, (*[50]uint64)(in), initoffset)
		}
	case 51:
		if len(in) >= 51 { // eliminate bounds check
			deltaunpackzigzag64_51(out, (*[51]uint64)(in), initoffset)
		}
	case 52:
		if len(in) >= 52 { // eliminate bounds check
			deltaunpackzigzag64_52(out, (*[52]uint64)(in), initoffset)
		}
	case 53:
		if len(in) >= 53 { // eliminate bounds check
			deltaunpackzigzag64_53(out, (*[53]uint64)(in), initoffset)
		}
	case 54:
		if len(in) >= 54 { // eliminate bounds check
			deltaunpackzigzag64_54(out, (*[54]uint64)(in), initoffset)
		}
	case 55:
		if len(in) >= 55 { // eliminate bounds check
			deltaunpackzigzag64_55(out, (*[55]uint64)(in), initoffset)
		}
	case 56:
		if len(in) >= 56 { // eliminate bounds check
			deltaunpackzigzag64_56(out, (*[56]uint64)(in), initoffset)
		}
	case 57:
		if len(in) >= 57 { // eliminate bounds check
			deltaunpackzigzag64_57(out, (*[57]uint64)(in), initoffset)
		}
	case 58:
		if len(in) >= 58 { // eliminate bounds check
			deltaunpackzigzag64_58(out, (*[58]uint64)(in), initoffset)
		}
	case 59:
		if len(in) >= 59 { // eliminate bounds check
			deltaunpackzigzag64_59(out, (*[59]uint64)(in), initoffset)
		}
	case 60:
		if len(in) >= 60 { // eliminate bounds check
			deltaunpackzigzag64_60(out, (*[60]uint64)(in), initoffset)
		}
	case 61:
		if len(in) >= 61 { // eliminate bounds check
			deltaunpackzigzag64_61(out, (*[61]uint64)(in), initoffset)
		}
	case 62:
		if len(in) >= 62 { // eliminate bounds check
			deltaunpackzigzag64_62(out, (*[62]uint64)(in), initoffset)
		}
	case 63:
		if len(in) >= 63 { // eliminate bounds check
			deltaunpackzigzag64_63(out, (*[63]uint64)(in), initoffset)
		}
	case 64:
		if len(in) >= 64 { // eliminate bounds check
			*out = *(*[64]uint64)(unsafe.Pointer((*[64]uint64)(in)))
		}
	}
}
