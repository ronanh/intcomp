// This file is generated, do not modify directly
// Use 'go generate' to regenerate.

package intcomp

import "unsafe"

// deltaPack_uint32 Binary packing of one block of `in`, starting from `initoffset`
// to out. Differential coding is applied first.
// Caller must give the proper `bitlen` of the block
func deltaPack_uint32[T uint32 | int32](out []uint32, in *[32]T, initoffset T, bitlen int) {
	switch bitlen {
	case 0:
		if len(out) >= 0 { // eliminate bounds check
			deltapack32_0((*[0]uint32)(out), in, initoffset)
		}
	case 1:
		if len(out) >= 1 { // eliminate bounds check
			deltapack32_1((*[1]uint32)(out), in, initoffset)
		}
	case 2:
		if len(out) >= 2 { // eliminate bounds check
			deltapack32_2((*[2]uint32)(out), in, initoffset)
		}
	case 3:
		if len(out) >= 3 { // eliminate bounds check
			deltapack32_3((*[3]uint32)(out), in, initoffset)
		}
	case 4:
		if len(out) >= 4 { // eliminate bounds check
			deltapack32_4((*[4]uint32)(out), in, initoffset)
		}
	case 5:
		if len(out) >= 5 { // eliminate bounds check
			deltapack32_5((*[5]uint32)(out), in, initoffset)
		}
	case 6:
		if len(out) >= 6 { // eliminate bounds check
			deltapack32_6((*[6]uint32)(out), in, initoffset)
		}
	case 7:
		if len(out) >= 7 { // eliminate bounds check
			deltapack32_7((*[7]uint32)(out), in, initoffset)
		}
	case 8:
		if len(out) >= 8 { // eliminate bounds check
			deltapack32_8((*[8]uint32)(out), in, initoffset)
		}
	case 9:
		if len(out) >= 9 { // eliminate bounds check
			deltapack32_9((*[9]uint32)(out), in, initoffset)
		}
	case 10:
		if len(out) >= 10 { // eliminate bounds check
			deltapack32_10((*[10]uint32)(out), in, initoffset)
		}
	case 11:
		if len(out) >= 11 { // eliminate bounds check
			deltapack32_11((*[11]uint32)(out), in, initoffset)
		}
	case 12:
		if len(out) >= 12 { // eliminate bounds check
			deltapack32_12((*[12]uint32)(out), in, initoffset)
		}
	case 13:
		if len(out) >= 13 { // eliminate bounds check
			deltapack32_13((*[13]uint32)(out), in, initoffset)
		}
	case 14:
		if len(out) >= 14 { // eliminate bounds check
			deltapack32_14((*[14]uint32)(out), in, initoffset)
		}
	case 15:
		if len(out) >= 15 { // eliminate bounds check
			deltapack32_15((*[15]uint32)(out), in, initoffset)
		}
	case 16:
		if len(out) >= 16 { // eliminate bounds check
			deltapack32_16((*[16]uint32)(out), in, initoffset)
		}
	case 17:
		if len(out) >= 17 { // eliminate bounds check
			deltapack32_17((*[17]uint32)(out), in, initoffset)
		}
	case 18:
		if len(out) >= 18 { // eliminate bounds check
			deltapack32_18((*[18]uint32)(out), in, initoffset)
		}
	case 19:
		if len(out) >= 19 { // eliminate bounds check
			deltapack32_19((*[19]uint32)(out), in, initoffset)
		}
	case 20:
		if len(out) >= 20 { // eliminate bounds check
			deltapack32_20((*[20]uint32)(out), in, initoffset)
		}
	case 21:
		if len(out) >= 21 { // eliminate bounds check
			deltapack32_21((*[21]uint32)(out), in, initoffset)
		}
	case 22:
		if len(out) >= 22 { // eliminate bounds check
			deltapack32_22((*[22]uint32)(out), in, initoffset)
		}
	case 23:
		if len(out) >= 23 { // eliminate bounds check
			deltapack32_23((*[23]uint32)(out), in, initoffset)
		}
	case 24:
		if len(out) >= 24 { // eliminate bounds check
			deltapack32_24((*[24]uint32)(out), in, initoffset)
		}
	case 25:
		if len(out) >= 25 { // eliminate bounds check
			deltapack32_25((*[25]uint32)(out), in, initoffset)
		}
	case 26:
		if len(out) >= 26 { // eliminate bounds check
			deltapack32_26((*[26]uint32)(out), in, initoffset)
		}
	case 27:
		if len(out) >= 27 { // eliminate bounds check
			deltapack32_27((*[27]uint32)(out), in, initoffset)
		}
	case 28:
		if len(out) >= 28 { // eliminate bounds check
			deltapack32_28((*[28]uint32)(out), in, initoffset)
		}
	case 29:
		if len(out) >= 29 { // eliminate bounds check
			deltapack32_29((*[29]uint32)(out), in, initoffset)
		}
	case 30:
		if len(out) >= 30 { // eliminate bounds check
			deltapack32_30((*[30]uint32)(out), in, initoffset)
		}
	case 31:
		if len(out) >= 31 { // eliminate bounds check
			deltapack32_31((*[31]uint32)(out), in, initoffset)
		}
	case 32:
		if len(out) >= 32 { // eliminate bounds check
			*(*[32]uint32)(out) = *((*[32]uint32)(unsafe.Pointer(in)))
		}
	}
}

// deltaUnpack_uint32 Decoding operation for DeltaPack_uint32
func deltaUnpack_uint32[T uint32 | int32](out *[32]T, in []uint32, initoffset T, bitlen int) {
	switch bitlen {
	case 0:
		if len(in) >= 0 { // eliminate bounds check
			deltaunpack32_0(out, (*[0]uint32)(in), initoffset)
		}
	case 1:
		if len(in) >= 1 { // eliminate bounds check
			deltaunpack32_1(out, (*[1]uint32)(in), initoffset)
		}
	case 2:
		if len(in) >= 2 { // eliminate bounds check
			deltaunpack32_2(out, (*[2]uint32)(in), initoffset)
		}
	case 3:
		if len(in) >= 3 { // eliminate bounds check
			deltaunpack32_3(out, (*[3]uint32)(in), initoffset)
		}
	case 4:
		if len(in) >= 4 { // eliminate bounds check
			deltaunpack32_4(out, (*[4]uint32)(in), initoffset)
		}
	case 5:
		if len(in) >= 5 { // eliminate bounds check
			deltaunpack32_5(out, (*[5]uint32)(in), initoffset)
		}
	case 6:
		if len(in) >= 6 { // eliminate bounds check
			deltaunpack32_6(out, (*[6]uint32)(in), initoffset)
		}
	case 7:
		if len(in) >= 7 { // eliminate bounds check
			deltaunpack32_7(out, (*[7]uint32)(in), initoffset)
		}
	case 8:
		if len(in) >= 8 { // eliminate bounds check
			deltaunpack32_8(out, (*[8]uint32)(in), initoffset)
		}
	case 9:
		if len(in) >= 9 { // eliminate bounds check
			deltaunpack32_9(out, (*[9]uint32)(in), initoffset)
		}
	case 10:
		if len(in) >= 10 { // eliminate bounds check
			deltaunpack32_10(out, (*[10]uint32)(in), initoffset)
		}
	case 11:
		if len(in) >= 11 { // eliminate bounds check
			deltaunpack32_11(out, (*[11]uint32)(in), initoffset)
		}
	case 12:
		if len(in) >= 12 { // eliminate bounds check
			deltaunpack32_12(out, (*[12]uint32)(in), initoffset)
		}
	case 13:
		if len(in) >= 13 { // eliminate bounds check
			deltaunpack32_13(out, (*[13]uint32)(in), initoffset)
		}
	case 14:
		if len(in) >= 14 { // eliminate bounds check
			deltaunpack32_14(out, (*[14]uint32)(in), initoffset)
		}
	case 15:
		if len(in) >= 15 { // eliminate bounds check
			deltaunpack32_15(out, (*[15]uint32)(in), initoffset)
		}
	case 16:
		if len(in) >= 16 { // eliminate bounds check
			deltaunpack32_16(out, (*[16]uint32)(in), initoffset)
		}
	case 17:
		if len(in) >= 17 { // eliminate bounds check
			deltaunpack32_17(out, (*[17]uint32)(in), initoffset)
		}
	case 18:
		if len(in) >= 18 { // eliminate bounds check
			deltaunpack32_18(out, (*[18]uint32)(in), initoffset)
		}
	case 19:
		if len(in) >= 19 { // eliminate bounds check
			deltaunpack32_19(out, (*[19]uint32)(in), initoffset)
		}
	case 20:
		if len(in) >= 20 { // eliminate bounds check
			deltaunpack32_20(out, (*[20]uint32)(in), initoffset)
		}
	case 21:
		if len(in) >= 21 { // eliminate bounds check
			deltaunpack32_21(out, (*[21]uint32)(in), initoffset)
		}
	case 22:
		if len(in) >= 22 { // eliminate bounds check
			deltaunpack32_22(out, (*[22]uint32)(in), initoffset)
		}
	case 23:
		if len(in) >= 23 { // eliminate bounds check
			deltaunpack32_23(out, (*[23]uint32)(in), initoffset)
		}
	case 24:
		if len(in) >= 24 { // eliminate bounds check
			deltaunpack32_24(out, (*[24]uint32)(in), initoffset)
		}
	case 25:
		if len(in) >= 25 { // eliminate bounds check
			deltaunpack32_25(out, (*[25]uint32)(in), initoffset)
		}
	case 26:
		if len(in) >= 26 { // eliminate bounds check
			deltaunpack32_26(out, (*[26]uint32)(in), initoffset)
		}
	case 27:
		if len(in) >= 27 { // eliminate bounds check
			deltaunpack32_27(out, (*[27]uint32)(in), initoffset)
		}
	case 28:
		if len(in) >= 28 { // eliminate bounds check
			deltaunpack32_28(out, (*[28]uint32)(in), initoffset)
		}
	case 29:
		if len(in) >= 29 { // eliminate bounds check
			deltaunpack32_29(out, (*[29]uint32)(in), initoffset)
		}
	case 30:
		if len(in) >= 30 { // eliminate bounds check
			deltaunpack32_30(out, (*[30]uint32)(in), initoffset)
		}
	case 31:
		if len(in) >= 31 { // eliminate bounds check
			deltaunpack32_31(out, (*[31]uint32)(in), initoffset)
		}
	case 32:
		if len(in) >= 32 { // eliminate bounds check
			*out = *(*[32]T)(unsafe.Pointer((*[32]uint32)(in)))
		}
	}
}

// --- zigzag

// deltaPackZigzag_uint32 Binary packing of one block of `in`, starting from `initoffset`
// to out. Differential coding is applied first, the difference is zigzag encoded.
//
//	Caller must give the proper `bitlen` of the block
func deltaPackZigzag_uint32(out []uint32, in *[32]uint32, initoffset uint32, bitlen int) {
	switch bitlen {
	case 0:
		if len(out) >= 0 { // eliminate bounds check
			deltapackzigzag32_0((*[0]uint32)(out), in, initoffset)
		}
	case 1:
		if len(out) >= 1 { // eliminate bounds check
			deltapackzigzag32_1((*[1]uint32)(out), in, initoffset)
		}
	case 2:
		if len(out) >= 2 { // eliminate bounds check
			deltapackzigzag32_2((*[2]uint32)(out), in, initoffset)
		}
	case 3:
		if len(out) >= 3 { // eliminate bounds check
			deltapackzigzag32_3((*[3]uint32)(out), in, initoffset)
		}
	case 4:
		if len(out) >= 4 { // eliminate bounds check
			deltapackzigzag32_4((*[4]uint32)(out), in, initoffset)
		}
	case 5:
		if len(out) >= 5 { // eliminate bounds check
			deltapackzigzag32_5((*[5]uint32)(out), in, initoffset)
		}
	case 6:
		if len(out) >= 6 { // eliminate bounds check
			deltapackzigzag32_6((*[6]uint32)(out), in, initoffset)
		}
	case 7:
		if len(out) >= 7 { // eliminate bounds check
			deltapackzigzag32_7((*[7]uint32)(out), in, initoffset)
		}
	case 8:
		if len(out) >= 8 { // eliminate bounds check
			deltapackzigzag32_8((*[8]uint32)(out), in, initoffset)
		}
	case 9:
		if len(out) >= 9 { // eliminate bounds check
			deltapackzigzag32_9((*[9]uint32)(out), in, initoffset)
		}
	case 10:
		if len(out) >= 10 { // eliminate bounds check
			deltapackzigzag32_10((*[10]uint32)(out), in, initoffset)
		}
	case 11:
		if len(out) >= 11 { // eliminate bounds check
			deltapackzigzag32_11((*[11]uint32)(out), in, initoffset)
		}
	case 12:
		if len(out) >= 12 { // eliminate bounds check
			deltapackzigzag32_12((*[12]uint32)(out), in, initoffset)
		}
	case 13:
		if len(out) >= 13 { // eliminate bounds check
			deltapackzigzag32_13((*[13]uint32)(out), in, initoffset)
		}
	case 14:
		if len(out) >= 14 { // eliminate bounds check
			deltapackzigzag32_14((*[14]uint32)(out), in, initoffset)
		}
	case 15:
		if len(out) >= 15 { // eliminate bounds check
			deltapackzigzag32_15((*[15]uint32)(out), in, initoffset)
		}
	case 16:
		if len(out) >= 16 { // eliminate bounds check
			deltapackzigzag32_16((*[16]uint32)(out), in, initoffset)
		}
	case 17:
		if len(out) >= 17 { // eliminate bounds check
			deltapackzigzag32_17((*[17]uint32)(out), in, initoffset)
		}
	case 18:
		if len(out) >= 18 { // eliminate bounds check
			deltapackzigzag32_18((*[18]uint32)(out), in, initoffset)
		}
	case 19:
		if len(out) >= 19 { // eliminate bounds check
			deltapackzigzag32_19((*[19]uint32)(out), in, initoffset)
		}
	case 20:
		if len(out) >= 20 { // eliminate bounds check
			deltapackzigzag32_20((*[20]uint32)(out), in, initoffset)
		}
	case 21:
		if len(out) >= 21 { // eliminate bounds check
			deltapackzigzag32_21((*[21]uint32)(out), in, initoffset)
		}
	case 22:
		if len(out) >= 22 { // eliminate bounds check
			deltapackzigzag32_22((*[22]uint32)(out), in, initoffset)
		}
	case 23:
		if len(out) >= 23 { // eliminate bounds check
			deltapackzigzag32_23((*[23]uint32)(out), in, initoffset)
		}
	case 24:
		if len(out) >= 24 { // eliminate bounds check
			deltapackzigzag32_24((*[24]uint32)(out), in, initoffset)
		}
	case 25:
		if len(out) >= 25 { // eliminate bounds check
			deltapackzigzag32_25((*[25]uint32)(out), in, initoffset)
		}
	case 26:
		if len(out) >= 26 { // eliminate bounds check
			deltapackzigzag32_26((*[26]uint32)(out), in, initoffset)
		}
	case 27:
		if len(out) >= 27 { // eliminate bounds check
			deltapackzigzag32_27((*[27]uint32)(out), in, initoffset)
		}
	case 28:
		if len(out) >= 28 { // eliminate bounds check
			deltapackzigzag32_28((*[28]uint32)(out), in, initoffset)
		}
	case 29:
		if len(out) >= 29 { // eliminate bounds check
			deltapackzigzag32_29((*[29]uint32)(out), in, initoffset)
		}
	case 30:
		if len(out) >= 30 { // eliminate bounds check
			deltapackzigzag32_30((*[30]uint32)(out), in, initoffset)
		}
	case 31:
		if len(out) >= 31 { // eliminate bounds check
			deltapackzigzag32_31((*[31]uint32)(out), in, initoffset)
		}
	case 32:
		if len(out) >= 32 { // eliminate bounds check
			*(*[32]uint32)(out) = *((*[32]uint32)(unsafe.Pointer(in)))
		}
	}
}

// deltaUnpackZigzag_uint32 Decoding operation for DeltaPackZigzag_uint32
func deltaUnpackZigzag_uint32(out *[32]uint32, in []uint32, initoffset uint32, bitlen int) {
	switch bitlen {
	case 0:
		if len(in) >= 0 { // eliminate bounds check
			deltaunpackzigzag32_0(out, (*[0]uint32)(in), initoffset)
		}
	case 1:
		if len(in) >= 1 { // eliminate bounds check
			deltaunpackzigzag32_1(out, (*[1]uint32)(in), initoffset)
		}
	case 2:
		if len(in) >= 2 { // eliminate bounds check
			deltaunpackzigzag32_2(out, (*[2]uint32)(in), initoffset)
		}
	case 3:
		if len(in) >= 3 { // eliminate bounds check
			deltaunpackzigzag32_3(out, (*[3]uint32)(in), initoffset)
		}
	case 4:
		if len(in) >= 4 { // eliminate bounds check
			deltaunpackzigzag32_4(out, (*[4]uint32)(in), initoffset)
		}
	case 5:
		if len(in) >= 5 { // eliminate bounds check
			deltaunpackzigzag32_5(out, (*[5]uint32)(in), initoffset)
		}
	case 6:
		if len(in) >= 6 { // eliminate bounds check
			deltaunpackzigzag32_6(out, (*[6]uint32)(in), initoffset)
		}
	case 7:
		if len(in) >= 7 { // eliminate bounds check
			deltaunpackzigzag32_7(out, (*[7]uint32)(in), initoffset)
		}
	case 8:
		if len(in) >= 8 { // eliminate bounds check
			deltaunpackzigzag32_8(out, (*[8]uint32)(in), initoffset)
		}
	case 9:
		if len(in) >= 9 { // eliminate bounds check
			deltaunpackzigzag32_9(out, (*[9]uint32)(in), initoffset)
		}
	case 10:
		if len(in) >= 10 { // eliminate bounds check
			deltaunpackzigzag32_10(out, (*[10]uint32)(in), initoffset)
		}
	case 11:
		if len(in) >= 11 { // eliminate bounds check
			deltaunpackzigzag32_11(out, (*[11]uint32)(in), initoffset)
		}
	case 12:
		if len(in) >= 12 { // eliminate bounds check
			deltaunpackzigzag32_12(out, (*[12]uint32)(in), initoffset)
		}
	case 13:
		if len(in) >= 13 { // eliminate bounds check
			deltaunpackzigzag32_13(out, (*[13]uint32)(in), initoffset)
		}
	case 14:
		if len(in) >= 14 { // eliminate bounds check
			deltaunpackzigzag32_14(out, (*[14]uint32)(in), initoffset)
		}
	case 15:
		if len(in) >= 15 { // eliminate bounds check
			deltaunpackzigzag32_15(out, (*[15]uint32)(in), initoffset)
		}
	case 16:
		if len(in) >= 16 { // eliminate bounds check
			deltaunpackzigzag32_16(out, (*[16]uint32)(in), initoffset)
		}
	case 17:
		if len(in) >= 17 { // eliminate bounds check
			deltaunpackzigzag32_17(out, (*[17]uint32)(in), initoffset)
		}
	case 18:
		if len(in) >= 18 { // eliminate bounds check
			deltaunpackzigzag32_18(out, (*[18]uint32)(in), initoffset)
		}
	case 19:
		if len(in) >= 19 { // eliminate bounds check
			deltaunpackzigzag32_19(out, (*[19]uint32)(in), initoffset)
		}
	case 20:
		if len(in) >= 20 { // eliminate bounds check
			deltaunpackzigzag32_20(out, (*[20]uint32)(in), initoffset)
		}
	case 21:
		if len(in) >= 21 { // eliminate bounds check
			deltaunpackzigzag32_21(out, (*[21]uint32)(in), initoffset)
		}
	case 22:
		if len(in) >= 22 { // eliminate bounds check
			deltaunpackzigzag32_22(out, (*[22]uint32)(in), initoffset)
		}
	case 23:
		if len(in) >= 23 { // eliminate bounds check
			deltaunpackzigzag32_23(out, (*[23]uint32)(in), initoffset)
		}
	case 24:
		if len(in) >= 24 { // eliminate bounds check
			deltaunpackzigzag32_24(out, (*[24]uint32)(in), initoffset)
		}
	case 25:
		if len(in) >= 25 { // eliminate bounds check
			deltaunpackzigzag32_25(out, (*[25]uint32)(in), initoffset)
		}
	case 26:
		if len(in) >= 26 { // eliminate bounds check
			deltaunpackzigzag32_26(out, (*[26]uint32)(in), initoffset)
		}
	case 27:
		if len(in) >= 27 { // eliminate bounds check
			deltaunpackzigzag32_27(out, (*[27]uint32)(in), initoffset)
		}
	case 28:
		if len(in) >= 28 { // eliminate bounds check
			deltaunpackzigzag32_28(out, (*[28]uint32)(in), initoffset)
		}
	case 29:
		if len(in) >= 29 { // eliminate bounds check
			deltaunpackzigzag32_29(out, (*[29]uint32)(in), initoffset)
		}
	case 30:
		if len(in) >= 30 { // eliminate bounds check
			deltaunpackzigzag32_30(out, (*[30]uint32)(in), initoffset)
		}
	case 31:
		if len(in) >= 31 { // eliminate bounds check
			deltaunpackzigzag32_31(out, (*[31]uint32)(in), initoffset)
		}
	case 32:
		if len(in) >= 32 { // eliminate bounds check
			*out = *(*[32]uint32)(unsafe.Pointer((*[32]uint32)(in)))
		}
	}
}
