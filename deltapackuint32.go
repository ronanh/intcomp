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
		deltapack32_0((*[0]uint32)(out), in, initoffset)
	case 1:
		deltapack32_1((*[1]uint32)(out), in, initoffset)
	case 2:
		deltapack32_2((*[2]uint32)(out), in, initoffset)
	case 3:
		deltapack32_3((*[3]uint32)(out), in, initoffset)
	case 4:
		deltapack32_4((*[4]uint32)(out), in, initoffset)
	case 5:
		deltapack32_5((*[5]uint32)(out), in, initoffset)
	case 6:
		deltapack32_6((*[6]uint32)(out), in, initoffset)
	case 7:
		deltapack32_7((*[7]uint32)(out), in, initoffset)
	case 8:
		deltapack32_8((*[8]uint32)(out), in, initoffset)
	case 9:
		deltapack32_9((*[9]uint32)(out), in, initoffset)
	case 10:
		deltapack32_10((*[10]uint32)(out), in, initoffset)
	case 11:
		deltapack32_11((*[11]uint32)(out), in, initoffset)
	case 12:
		deltapack32_12((*[12]uint32)(out), in, initoffset)
	case 13:
		deltapack32_13((*[13]uint32)(out), in, initoffset)
	case 14:
		deltapack32_14((*[14]uint32)(out), in, initoffset)
	case 15:
		deltapack32_15((*[15]uint32)(out), in, initoffset)
	case 16:
		deltapack32_16((*[16]uint32)(out), in, initoffset)
	case 17:
		deltapack32_17((*[17]uint32)(out), in, initoffset)
	case 18:
		deltapack32_18((*[18]uint32)(out), in, initoffset)
	case 19:
		deltapack32_19((*[19]uint32)(out), in, initoffset)
	case 20:
		deltapack32_20((*[20]uint32)(out), in, initoffset)
	case 21:
		deltapack32_21((*[21]uint32)(out), in, initoffset)
	case 22:
		deltapack32_22((*[22]uint32)(out), in, initoffset)
	case 23:
		deltapack32_23((*[23]uint32)(out), in, initoffset)
	case 24:
		deltapack32_24((*[24]uint32)(out), in, initoffset)
	case 25:
		deltapack32_25((*[25]uint32)(out), in, initoffset)
	case 26:
		deltapack32_26((*[26]uint32)(out), in, initoffset)
	case 27:
		deltapack32_27((*[27]uint32)(out), in, initoffset)
	case 28:
		deltapack32_28((*[28]uint32)(out), in, initoffset)
	case 29:
		deltapack32_29((*[29]uint32)(out), in, initoffset)
	case 30:
		deltapack32_30((*[30]uint32)(out), in, initoffset)
	case 31:
		deltapack32_31((*[31]uint32)(out), in, initoffset)
	case 32:
		*(*[32]uint32)(out) = *((*[32]uint32)(unsafe.Pointer(in)))
	default:
		panic("unsupported bitlen")
	}
}

// deltaUnpack_uint32 Decoding operation for DeltaPack_uint32
func deltaUnpack_uint32[T uint32 | int32](out *[32]T, in []uint32, initoffset T, bitlen int) {
	switch bitlen {
	case 0:
		deltaunpack32_0(out, (*[0]uint32)(in), initoffset)
	case 1:
		deltaunpack32_1(out, (*[1]uint32)(in), initoffset)
	case 2:
		deltaunpack32_2(out, (*[2]uint32)(in), initoffset)
	case 3:
		deltaunpack32_3(out, (*[3]uint32)(in), initoffset)
	case 4:
		deltaunpack32_4(out, (*[4]uint32)(in), initoffset)
	case 5:
		deltaunpack32_5(out, (*[5]uint32)(in), initoffset)
	case 6:
		deltaunpack32_6(out, (*[6]uint32)(in), initoffset)
	case 7:
		deltaunpack32_7(out, (*[7]uint32)(in), initoffset)
	case 8:
		deltaunpack32_8(out, (*[8]uint32)(in), initoffset)
	case 9:
		deltaunpack32_9(out, (*[9]uint32)(in), initoffset)
	case 10:
		deltaunpack32_10(out, (*[10]uint32)(in), initoffset)
	case 11:
		deltaunpack32_11(out, (*[11]uint32)(in), initoffset)
	case 12:
		deltaunpack32_12(out, (*[12]uint32)(in), initoffset)
	case 13:
		deltaunpack32_13(out, (*[13]uint32)(in), initoffset)
	case 14:
		deltaunpack32_14(out, (*[14]uint32)(in), initoffset)
	case 15:
		deltaunpack32_15(out, (*[15]uint32)(in), initoffset)
	case 16:
		deltaunpack32_16(out, (*[16]uint32)(in), initoffset)
	case 17:
		deltaunpack32_17(out, (*[17]uint32)(in), initoffset)
	case 18:
		deltaunpack32_18(out, (*[18]uint32)(in), initoffset)
	case 19:
		deltaunpack32_19(out, (*[19]uint32)(in), initoffset)
	case 20:
		deltaunpack32_20(out, (*[20]uint32)(in), initoffset)
	case 21:
		deltaunpack32_21(out, (*[21]uint32)(in), initoffset)
	case 22:
		deltaunpack32_22(out, (*[22]uint32)(in), initoffset)
	case 23:
		deltaunpack32_23(out, (*[23]uint32)(in), initoffset)
	case 24:
		deltaunpack32_24(out, (*[24]uint32)(in), initoffset)
	case 25:
		deltaunpack32_25(out, (*[25]uint32)(in), initoffset)
	case 26:
		deltaunpack32_26(out, (*[26]uint32)(in), initoffset)
	case 27:
		deltaunpack32_27(out, (*[27]uint32)(in), initoffset)
	case 28:
		deltaunpack32_28(out, (*[28]uint32)(in), initoffset)
	case 29:
		deltaunpack32_29(out, (*[29]uint32)(in), initoffset)
	case 30:
		deltaunpack32_30(out, (*[30]uint32)(in), initoffset)
	case 31:
		deltaunpack32_31(out, (*[31]uint32)(in), initoffset)
	case 32:
		*out = *(*[32]T)(unsafe.Pointer((*[32]uint32)(in)))
	default:
		panic("unsupported bitlen")
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
		deltapackzigzag32_0((*[0]uint32)(out), in, initoffset)
	case 1:
		deltapackzigzag32_1((*[1]uint32)(out), in, initoffset)
	case 2:
		deltapackzigzag32_2((*[2]uint32)(out), in, initoffset)
	case 3:
		deltapackzigzag32_3((*[3]uint32)(out), in, initoffset)
	case 4:
		deltapackzigzag32_4((*[4]uint32)(out), in, initoffset)
	case 5:
		deltapackzigzag32_5((*[5]uint32)(out), in, initoffset)
	case 6:
		deltapackzigzag32_6((*[6]uint32)(out), in, initoffset)
	case 7:
		deltapackzigzag32_7((*[7]uint32)(out), in, initoffset)
	case 8:
		deltapackzigzag32_8((*[8]uint32)(out), in, initoffset)
	case 9:
		deltapackzigzag32_9((*[9]uint32)(out), in, initoffset)
	case 10:
		deltapackzigzag32_10((*[10]uint32)(out), in, initoffset)
	case 11:
		deltapackzigzag32_11((*[11]uint32)(out), in, initoffset)
	case 12:
		deltapackzigzag32_12((*[12]uint32)(out), in, initoffset)
	case 13:
		deltapackzigzag32_13((*[13]uint32)(out), in, initoffset)
	case 14:
		deltapackzigzag32_14((*[14]uint32)(out), in, initoffset)
	case 15:
		deltapackzigzag32_15((*[15]uint32)(out), in, initoffset)
	case 16:
		deltapackzigzag32_16((*[16]uint32)(out), in, initoffset)
	case 17:
		deltapackzigzag32_17((*[17]uint32)(out), in, initoffset)
	case 18:
		deltapackzigzag32_18((*[18]uint32)(out), in, initoffset)
	case 19:
		deltapackzigzag32_19((*[19]uint32)(out), in, initoffset)
	case 20:
		deltapackzigzag32_20((*[20]uint32)(out), in, initoffset)
	case 21:
		deltapackzigzag32_21((*[21]uint32)(out), in, initoffset)
	case 22:
		deltapackzigzag32_22((*[22]uint32)(out), in, initoffset)
	case 23:
		deltapackzigzag32_23((*[23]uint32)(out), in, initoffset)
	case 24:
		deltapackzigzag32_24((*[24]uint32)(out), in, initoffset)
	case 25:
		deltapackzigzag32_25((*[25]uint32)(out), in, initoffset)
	case 26:
		deltapackzigzag32_26((*[26]uint32)(out), in, initoffset)
	case 27:
		deltapackzigzag32_27((*[27]uint32)(out), in, initoffset)
	case 28:
		deltapackzigzag32_28((*[28]uint32)(out), in, initoffset)
	case 29:
		deltapackzigzag32_29((*[29]uint32)(out), in, initoffset)
	case 30:
		deltapackzigzag32_30((*[30]uint32)(out), in, initoffset)
	case 31:
		deltapackzigzag32_31((*[31]uint32)(out), in, initoffset)
	case 32:
		*(*[32]uint32)(out) = *((*[32]uint32)(unsafe.Pointer(in)))
	default:
		panic("unsupported bitlen")
	}
}

// deltaUnpackZigzag_uint32 Decoding operation for DeltaPackZigzag_uint32
func deltaUnpackZigzag_uint32(out *[32]uint32, in []uint32, initoffset uint32, bitlen int) {
	switch bitlen {
	case 0:
		deltaunpackzigzag32_0(out, (*[0]uint32)(in), initoffset)
	case 1:
		deltaunpackzigzag32_1(out, (*[1]uint32)(in), initoffset)
	case 2:
		deltaunpackzigzag32_2(out, (*[2]uint32)(in), initoffset)
	case 3:
		deltaunpackzigzag32_3(out, (*[3]uint32)(in), initoffset)
	case 4:
		deltaunpackzigzag32_4(out, (*[4]uint32)(in), initoffset)
	case 5:
		deltaunpackzigzag32_5(out, (*[5]uint32)(in), initoffset)
	case 6:
		deltaunpackzigzag32_6(out, (*[6]uint32)(in), initoffset)
	case 7:
		deltaunpackzigzag32_7(out, (*[7]uint32)(in), initoffset)
	case 8:
		deltaunpackzigzag32_8(out, (*[8]uint32)(in), initoffset)
	case 9:
		deltaunpackzigzag32_9(out, (*[9]uint32)(in), initoffset)
	case 10:
		deltaunpackzigzag32_10(out, (*[10]uint32)(in), initoffset)
	case 11:
		deltaunpackzigzag32_11(out, (*[11]uint32)(in), initoffset)
	case 12:
		deltaunpackzigzag32_12(out, (*[12]uint32)(in), initoffset)
	case 13:
		deltaunpackzigzag32_13(out, (*[13]uint32)(in), initoffset)
	case 14:
		deltaunpackzigzag32_14(out, (*[14]uint32)(in), initoffset)
	case 15:
		deltaunpackzigzag32_15(out, (*[15]uint32)(in), initoffset)
	case 16:
		deltaunpackzigzag32_16(out, (*[16]uint32)(in), initoffset)
	case 17:
		deltaunpackzigzag32_17(out, (*[17]uint32)(in), initoffset)
	case 18:
		deltaunpackzigzag32_18(out, (*[18]uint32)(in), initoffset)
	case 19:
		deltaunpackzigzag32_19(out, (*[19]uint32)(in), initoffset)
	case 20:
		deltaunpackzigzag32_20(out, (*[20]uint32)(in), initoffset)
	case 21:
		deltaunpackzigzag32_21(out, (*[21]uint32)(in), initoffset)
	case 22:
		deltaunpackzigzag32_22(out, (*[22]uint32)(in), initoffset)
	case 23:
		deltaunpackzigzag32_23(out, (*[23]uint32)(in), initoffset)
	case 24:
		deltaunpackzigzag32_24(out, (*[24]uint32)(in), initoffset)
	case 25:
		deltaunpackzigzag32_25(out, (*[25]uint32)(in), initoffset)
	case 26:
		deltaunpackzigzag32_26(out, (*[26]uint32)(in), initoffset)
	case 27:
		deltaunpackzigzag32_27(out, (*[27]uint32)(in), initoffset)
	case 28:
		deltaunpackzigzag32_28(out, (*[28]uint32)(in), initoffset)
	case 29:
		deltaunpackzigzag32_29(out, (*[29]uint32)(in), initoffset)
	case 30:
		deltaunpackzigzag32_30(out, (*[30]uint32)(in), initoffset)
	case 31:
		deltaunpackzigzag32_31(out, (*[31]uint32)(in), initoffset)
	case 32:
		*out = *(*[32]uint32)(unsafe.Pointer((*[32]uint32)(in)))
	default:
		panic("unsupported bitlen")
	}
}
