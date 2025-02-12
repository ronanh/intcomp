package main

import (
	"fmt"
)

type deltapack struct {
}

func (dp *deltapack) Dpn() []dpN {
	dpn := make([]dpN, 64)
	for i := range dpn {
		dpn[i] = dpN{dp, i}
	}

	return dpn
}

type dpN struct {
	dp *deltapack
	N  int
}

func (dpn *dpN) DVar() string {
	return fmt.Sprintf("d%02d", dpn.N)
}

// func (dpn *dpN) DVal() string {
// 	// if dpn.N == 0 {
// 	// 	return "int64(in[0] - initoffset)"
// 	// }
// 	// return fmt.Sprintf("int64(in[%d]-in[%d])", dpn.N, dpn.N-1)

// 	if dpn.N == 0 {
// 		return "int64(*(*IT)(unsafe.Pointer(&in[0])) - *(*IT)(unsafe.Pointer(&initoffset)))"
// 	}
// 	return fmt.Sprintf("int64(*(*IT)(unsafe.Pointer(&in[%d])) - *(*IT)(unsafe.Pointer(&in[%d])))", dpn.N, dpn.N-1)
// }

func (dpn *dpN) DVal32() string {
	return dpn.dVal32(dpn.N)
}

func (dpn *dpN) dVal32(n int) string {
	if n == 0 {
		return "int64(*(*int32)(unsafe.Pointer(&in[0])) - *(*int32)(unsafe.Pointer(&initvalue)))"
	}
	return fmt.Sprintf("int64(*(*int32)(unsafe.Pointer(&in[%d])) - *(*int32)(unsafe.Pointer(&in[%d])))", n, n-1)
}

func (dpn *dpN) DVal64() string {
	return dpn.dVal64(dpn.N)
}

func (dpn *dpN) dVal64(n int) string {
	if n == 0 {
		return "*(*int64)(unsafe.Pointer(&in[0])) - *(*int64)(unsafe.Pointer(&initvalue))"
	}
	return fmt.Sprintf("*(*int64)(unsafe.Pointer(&in[%d])) - *(*int64)(unsafe.Pointer(&in[%d]))", n, n-1)
}

func (dpn *dpN) ZZDVar() string {
	return fmt.Sprintf("zzd%02d", dpn.N)
}

// func (dpn *dpN) ZZDVal() string {
// 	return fmt.Sprintf("uint64((%s << 1) ^ (%s >> 63))", dpn.DVal(), dpn.DVal())
// }

func (dpn *dpN) ZZDVal32() string {
	return dpn.zzdVal32(dpn.N)
}

func (dpn *dpN) zzdVal32(n int) string {
	return fmt.Sprintf("uint64(((%s) << 1) ^ ((%s) >> 63))", dpn.dVal32(n), dpn.dVal32(n))
}

func (dpn *dpN) ZZDVal64() string {
	return dpn.zzdVal64(dpn.N)
}

func (dpn *dpN) zzdVal64(n int) string {
	return fmt.Sprintf("uint64(((%s) << 1) ^ ((%s) >> 63))", dpn.dVal64(n), dpn.dVal64(n))
}

func (dpn *dpN) XORVar() string {
	return fmt.Sprintf("xor%02d", dpn.N)
}

func (dpn *dpN) XORVal32() string {
	return dpn.xorVal32(dpn.N)
}

func (dpn *dpN) xorVal32(n int) string {
	if n == 0 {
		return "uint64(uint32(*(*int32)(unsafe.Pointer(&initvalue)) ^ *(*int32)(unsafe.Pointer(&in[0]))))"
	}
	return fmt.Sprintf("uint64(uint32(*(*int32)(unsafe.Pointer(&in[%d])) ^ *(*int32)(unsafe.Pointer(&in[%d]))))", n-1, n)
}

func (dpn *dpN) XORVal64() string {
	return dpn.xorVal64(dpn.N)
}

func (dpn *dpN) xorVal64(n int) string {
	if n == 0 {
		return "uint64(*(*int64)(unsafe.Pointer(&initvalue)) ^ *(*int64)(unsafe.Pointer(&in[0])))"
	}
	return fmt.Sprintf("uint64(*(*int64)(unsafe.Pointer(&in[%d])) ^ *(*int64)(unsafe.Pointer(&in[%d])))", n-1, n)
}

func (dpn *dpN) Comma() string {
	if dpn.N == 0 {
		return ""
	}
	return ","
}

func (dpn *dpN) Pipe() string {
	if dpn.N == 0 {
		return ""
	}
	return "|"
}

func (dpn *dpN) NbBytes() int {
	return dpn.N
}

func (dpn *dpN) Dpnb() []dpNByte {
	dpnb := make([]dpNByte, dpn.N)
	for i := range dpnb {
		dpnb[i] = dpNByte{dpn, i}
	}
	return dpnb
}

func (dpn *dpN) Dunb() []duNByte {
	dunb := make([]duNByte, 64)
	for i := range dunb {
		dunb[i] = duNByte{dpn, i}
	}
	return dunb
}

type dpNByte struct {
	dpn *dpN
	I   int
}

func (dpnb *dpNByte) PackLines() []string {
	return dpnb.packLines(false, false)
}

func (dpnb *dpNByte) PackLinesZigZag() []string {
	return dpnb.packLines(true, false)
}

func (dpnb *dpNByte) PackLinesNtz() []string {
	return dpnb.packLines(false, true)
}

func (dpnb *dpNByte) PackLinesZigZagNtz() []string {
	return dpnb.packLines(true, true)
}

func (dpnb *dpNByte) packLines(zigzag, ntz bool) []string {
	var lines []string
	nbBytes := dpnb.dpn.NbBytes()
	if nbBytes == 0 {
		return nil
	}
	var overlapping bool
	minOffset := 64 * dpnb.I
	maxOffset := 64 * (dpnb.I + 1)
	if minOffset%nbBytes != 0 {
		minOffset -= nbBytes
		overlapping = true
	}
	for i := 0; i < 64; i++ {
		offset := i * nbBytes
		if offset < minOffset || offset >= maxOffset {
			continue
		}

		ishift := offset - minOffset
		if overlapping {
			ishift -= nbBytes
		}

		// v1 := fmt.Sprintf("in[%d]", i)
		// v2 := "initoffset"
		// if i > 0 {
		// 	v2 = fmt.Sprintf("in[%d]", i-1)
		// }

		// diff := fmt.Sprintf("*(*IT)(unsafe.Pointer(&%s)) - *(*IT)(unsafe.Pointer(&%s))", v1, v2)
		// // force signed 64-bit
		//diff = fmt.Sprintf("int64(%s)", diff)

		diff := fmt.Sprintf("d%02d", i)

		if zigzag {
			// // force signed
			// diff = fmt.Sprintf("int64(%s)", diff)
			// zigzag encoding of v: (v << 1) ^ (v >> 31)
			//diff = fmt.Sprintf("((%s) << 1) ^ ((%s) >> 63)", diff, diff)
			//diff = fmt.Sprintf("uint64((%s << 1) ^ (%s >> 63))", diff, diff)
			diff = fmt.Sprintf("zzd%02d", i)
			//diff = dpnb.dpn.ZZDVar()
		}
		//diff = fmt.Sprintf("uint64(%s)", diff)

		var line string
		switch {
		case ishift < 0:
			if ntz {
				//line = fmt.Sprintf("(%s) >> ntz >> %d", diff, -ishift)
				line = fmt.Sprintf("%s >> ntz >> %d", diff, -ishift)
			} else {
				//line = fmt.Sprintf("(%s) >> %d", diff, -ishift)
				line = fmt.Sprintf("%s >> %d", diff, -ishift)
			}
		case ishift > 0:
			if ntz {
				//line = fmt.Sprintf("((%s) >> ntz << %d)", diff, ishift)
				line = fmt.Sprintf("%s >> ntz << %d", diff, ishift)
			} else {
				//line = fmt.Sprintf("((%s) << %d)", diff, ishift)
				line = fmt.Sprintf("%s << %d", diff, ishift)
			}
		default:
			if ntz {
				// line = fmt.Sprintf("(%s) >> ntz", diff)
				line = fmt.Sprintf("%s >> ntz", diff)
			} else {
				line = diff
			}
		}
		if offset+nbBytes < maxOffset {
			line += " | "
		}

		lines = append(lines, line)
	}
	return lines
}

// func (dpnb *dpNByte) PackLinesZigZagNtz32() []string {
// 	return dpnb.packLines2(true)
// }

// func (dpnb *dpNByte) PackLinesZigZagNtz64() []string {
// 	return dpnb.packLines2(false)
// }

// func (dpnb *dpNByte) packLines2(bits32 bool) []string {
// 	var lines []string
// 	nbBytes := dpnb.dpn.NbBytes()
// 	if nbBytes == 0 {
// 		return nil
// 	}
// 	var overlapping bool
// 	minOffset := 64 * dpnb.I
// 	maxOffset := 64 * (dpnb.I + 1)
// 	if minOffset%nbBytes != 0 {
// 		minOffset -= nbBytes
// 		overlapping = true
// 	}
// 	for i := 0; i < 64; i++ {
// 		offset := i * nbBytes
// 		if offset < minOffset || offset >= maxOffset {
// 			continue
// 		}

// 		ishift := offset - minOffset
// 		if overlapping {
// 			ishift -= nbBytes
// 		}

// 		// v1 := fmt.Sprintf("in[%d]", i)
// 		// v2 := "initoffset"
// 		// if i > 0 {
// 		// 	v2 = fmt.Sprintf("in[%d]", i-1)
// 		// }

// 		// diff := fmt.Sprintf("*(*IT)(unsafe.Pointer(&%s)) - *(*IT)(unsafe.Pointer(&%s))", v1, v2)
// 		// // force signed 64-bit
// 		//diff = fmt.Sprintf("int64(%s)", diff)

// 		// // force signed
// 		// diff = fmt.Sprintf("int64(%s)", diff)
// 		// zigzag encoding of v: (v << 1) ^ (v >> 31)
// 		//diff = fmt.Sprintf("((%s) << 1) ^ ((%s) >> 63)", diff, diff)
// 		//diff = fmt.Sprintf("uint64((%s << 1) ^ (%s >> 63))", diff, diff)
// 		diff := dpnb.dpn.zzdVal32(i)

// 		//diff = dpnb.dpn.ZZDVar()
// 		//diff = fmt.Sprintf("uint64(%s)", diff)

// 		var line string
// 		switch {
// 		case ishift < 0:
// 			line = fmt.Sprintf("%s >> ntz >> %d", diff, -ishift)
// 		case ishift > 0:
// 			line = fmt.Sprintf("%s >> ntz << %d", diff, ishift)
// 		default:
// 			line = fmt.Sprintf("%s >> ntz", diff)
// 		}
// 		if offset+nbBytes < maxOffset {
// 			line += " | "
// 		}

// 		lines = append(lines, line)
// 	}
// 	return lines
// }

func (dpnb *dpNByte) PackLinesXor() []string {
	var lines []string
	nbBytes := dpnb.dpn.NbBytes()
	if nbBytes == 0 {
		return nil
	}
	var overlapping bool
	minOffset := 64 * dpnb.I
	maxOffset := 64 * (dpnb.I + 1)
	if minOffset%nbBytes != 0 {
		minOffset -= nbBytes
		overlapping = true
	}
	for i := 0; i < 64; i++ {
		offset := i * nbBytes
		if offset < minOffset || offset >= maxOffset {
			continue
		}

		ishift := offset - minOffset
		if overlapping {
			ishift -= nbBytes
		}

		xor := fmt.Sprintf("xor%02d", i)

		var line string
		switch {
		case ishift < 0:
			line = fmt.Sprintf("%s >> ntz >> %d", xor, -ishift)
		case ishift > 0:
			line = fmt.Sprintf("%s >> ntz << %d", xor, ishift)
		default:
			line = fmt.Sprintf("%s >> ntz", xor)
		}
		if offset+nbBytes < maxOffset {
			line += " | "
		}

		lines = append(lines, line)
	}
	return lines
}

type duNByte struct {
	dpn *dpN
	I   int
}

// func (dunb *duNByte) UnpackLine() string {
// 	return dunb.unpackLine(false, false)
// }

// func (dunb *duNByte) UnpackLineZigZag() string {
// 	return dunb.unpackLine(true, false)
// }

// func (dunb *duNByte) UnpackLineNtz() string {
// 	return dunb.unpackLine(false, true)
// }

// func (dunb *duNByte) UnpackLineZigZagNtz() string {
// 	return dunb.unpackLine(true, true)
// }

// func (dunb *duNByte) UnpackLineZigZagNtz32() string {
// 	return dunb.unpackLine2(true, true)
// }

// func (dunb *duNByte) UnpackLineZigZagNtz64() string {
// 	return dunb.unpackLine2(false, true)
// }

// func (dunb *duNByte) UnpackLineZigZag32() string {
// 	return dunb.unpackLine2(true, false)
// }

// func (dunb *duNByte) UnpackLineZigZag64() string {
// 	return dunb.unpackLine2(false, false)
// }

// func (dunb *duNByte) unpackLine(zigzag, ntz bool) string {
// 	if dunb.dpn.N == 0 {
// 		return "initoffset"
// 	}
// 	nbBytes := dunb.dpn.NbBytes()
// 	startByte := dunb.I * nbBytes / 64
// 	startBit := dunb.I * nbBytes % 64
// 	endByte := (dunb.I + 1) * nbBytes / 64
// 	endBit := (dunb.I + 1) * nbBytes % 64

// 	var startMask, endMask int
// 	in1 := fmt.Sprintf("in[%d]", startByte)
// 	in2 := fmt.Sprintf("in[%d]", endByte)
// 	out := "initoffset"
// 	if dunb.I > 0 {
// 		out = fmt.Sprintf("out[%d]", dunb.I-1)
// 	}
// 	var val string
// 	if startByte == endByte {
// 		for i := startBit; i < endBit; i++ {
// 			startMask <<= 1
// 			startMask |= 1
// 		}
// 		// val = fmt.Sprintf("(int%d(uint%d(%s) >> %d)) & 0x%X", dunb.dpn.dp.Bits, dunb.dpn.dp.Bits, in1, startBit, startMask)
// 		// if dunb.dpn.dp.Unsigned {
// 		val = fmt.Sprintf("(%s >> %d) & 0x%X", in1, startBit, startMask)
// 		// }
// 	} else {
// 		for i := startBit; i < 64; i++ {
// 			startMask <<= 1
// 			startMask |= 1
// 		}
// 		for i := 0; i < endBit; i++ {
// 			endMask <<= 1
// 			endMask |= 1
// 		}
// 		// val = fmt.Sprintf("int%d(uint%d(%s) >> %d)", dunb.dpn.dp.Bits, dunb.dpn.dp.Bits, in1, startBit)
// 		// if dunb.dpn.dp.Unsigned {
// 		val = fmt.Sprintf("(%s >> %d)", in1, startBit)
// 		// }
// 		if endBit > 0 {
// 			val = fmt.Sprintf("(%s | ((%s & 0x%X) << %d))", val, in2, endMask, nbBytes-endBit)
// 		}
// 	}
// 	if zigzag {
// 		// force signed
// 		val = fmt.Sprintf("int%d(%s)", 64, val)
// 		// zigzag decoding of val: (-(val & 1))^(val>>1))
// 		val = fmt.Sprintf("((-((%s) & 1))^((%s)>>1))", val, val)
// 	}
// 	if ntz {
// 		return fmt.Sprintf("T(%s) << ntz + %s", val, out)
// 	} else {
// 		return fmt.Sprintf("T(%s) + %s", val, out)
// 	}
// }

// func (dunb *duNByte) unpackLine2(bits32, ntz bool) string {
// 	if dunb.dpn.N == 0 {
// 		return "initoffset"
// 	}
// 	nbBytes := dunb.dpn.NbBytes()
// 	startByte := dunb.I * nbBytes / 64
// 	startBit := dunb.I * nbBytes % 64
// 	endByte := (dunb.I + 1) * nbBytes / 64
// 	endBit := (dunb.I + 1) * nbBytes % 64

// 	var startMask, endMask int
// 	in1 := fmt.Sprintf("in2[%d]", startByte)
// 	in2 := fmt.Sprintf("in2[%d]", endByte)
// 	out := "*(*int64)(unsafe.Pointer(&initoffset))"
// 	if bits32 {
// 		out = "*(*int32)(unsafe.Pointer(&initoffset))"
// 	}
// 	if dunb.I > 0 {
// 		//out = fmt.Sprintf("out[%d]", dunb.I-1)
// 		out = "vout"
// 	}
// 	var val string
// 	if startByte == endByte {
// 		for i := startBit; i < endBit; i++ {
// 			startMask <<= 1
// 			startMask |= 1
// 		}
// 		// val = fmt.Sprintf("(int%d(uint%d(%s) >> %d)) & 0x%X", dunb.dpn.dp.Bits, dunb.dpn.dp.Bits, in1, startBit, startMask)
// 		// if dunb.dpn.dp.Unsigned {
// 		val = fmt.Sprintf("(%s >> %d) & 0x%X", in1, startBit, startMask)
// 		// }
// 	} else {
// 		for i := startBit; i < 64; i++ {
// 			startMask <<= 1
// 			startMask |= 1
// 		}
// 		for i := 0; i < endBit; i++ {
// 			endMask <<= 1
// 			endMask |= 1
// 		}
// 		// val = fmt.Sprintf("int%d(uint%d(%s) >> %d)", dunb.dpn.dp.Bits, dunb.dpn.dp.Bits, in1, startBit)
// 		// if dunb.dpn.dp.Unsigned {
// 		val = fmt.Sprintf("(%s >> %d)", in1, startBit)
// 		// }
// 		if endBit > 0 {
// 			val = fmt.Sprintf("(%s | ((%s & 0x%X) << %d))", val, in2, endMask, nbBytes-endBit)
// 		}
// 	}
// 	// force signed
// 	val = fmt.Sprintf("int%d(%s)", 64, val)
// 	// zigzag decoding of val: (-(val & 1))^(val>>1))
// 	val = fmt.Sprintf("((-(%s & 1))^(%s>>1))", val, val)
// 	if bits32 {
// 		if ntz {
// 			return fmt.Sprintf("int32(%s) << ntz + %s", val, out)
// 		} else {
// 			return fmt.Sprintf("int32(%s) + %s", val, out)
// 		}
// 	} else {
// 		if ntz {
// 			return fmt.Sprintf("int64(%s) << ntz + %s", val, out)
// 		} else {
// 			return fmt.Sprintf("int64(%s) + %s", val, out)
// 		}
// 	}
// }

func (dunb *duNByte) UnpackLineZigZagNtz() string {
	return dunb.unpackLine3(true)
}

func (dunb *duNByte) UnpackLineZigZag() string {
	return dunb.unpackLine3(false)
}

func (dunb *duNByte) unpackLine3(ntz bool) string {
	if dunb.dpn.N == 0 {
		return "initvalue"
	}
	nbBytes := dunb.dpn.NbBytes()
	startByte := dunb.I * nbBytes / 64
	startBit := dunb.I * nbBytes % 64
	endByte := (dunb.I + 1) * nbBytes / 64
	endBit := (dunb.I + 1) * nbBytes % 64

	var startMask, endMask int
	in1 := fmt.Sprintf("in2[%d]", startByte)
	in2 := fmt.Sprintf("in2[%d]", endByte)
	// out := "*(*IT)(unsafe.Pointer(&initoffset))"
	// if dunb.I > 0 {
	// 	//out = fmt.Sprintf("out[%d]", dunb.I-1)
	// 	out = "vout"
	// }
	var val string
	if startByte == endByte {
		for i := startBit; i < endBit; i++ {
			startMask <<= 1
			startMask |= 1
		}
		// val = fmt.Sprintf("(int%d(uint%d(%s) >> %d)) & 0x%X", dunb.dpn.dp.Bits, dunb.dpn.dp.Bits, in1, startBit, startMask)
		// if dunb.dpn.dp.Unsigned {
		val = fmt.Sprintf("((%s >> %d) & 0x%X)", in1, startBit, startMask)
		// }
	} else {
		for i := startBit; i < 64; i++ {
			startMask <<= 1
			startMask |= 1
		}
		for i := 0; i < endBit; i++ {
			endMask <<= 1
			endMask |= 1
		}
		// val = fmt.Sprintf("int%d(uint%d(%s) >> %d)", dunb.dpn.dp.Bits, dunb.dpn.dp.Bits, in1, startBit)
		// if dunb.dpn.dp.Unsigned {
		val = fmt.Sprintf("(%s >> %d)", in1, startBit)
		// }
		if endBit > 0 {
			val = fmt.Sprintf("(%s | ((%s & 0x%X) << %d))", val, in2, endMask, nbBytes-endBit)
		}
	}
	if ntz {
		val = fmt.Sprintf("(%s << ntz)", val)
	}
	// force signed
	//val = fmt.Sprintf("int64(%s)", val)
	// zigzag decoding of val: (-(val & 1))^(val>>1))
	// val = fmt.Sprintf("((-(%s & 1))^(%s>>1))", val, val)
	// if ntz {
	// 	return fmt.Sprintf("IT(%s) << ntz", val)
	// } else {
	// 	return fmt.Sprintf("IT(%s)", val)
	// }
	return fmt.Sprintf("IT((-(%s & 1))^(%s>>1))", val, val)
}

func (dunb *duNByte) UnpackLineXor() string {
	if dunb.dpn.N == 0 {
		return "initvalue"
	}
	nbBytes := dunb.dpn.NbBytes()
	startByte := dunb.I * nbBytes / 64
	startBit := dunb.I * nbBytes % 64
	endByte := (dunb.I + 1) * nbBytes / 64
	endBit := (dunb.I + 1) * nbBytes % 64

	var startMask, endMask int
	in1 := fmt.Sprintf("in2[%d]", startByte)
	in2 := fmt.Sprintf("in2[%d]", endByte)

	var val string
	if startByte == endByte {
		for i := startBit; i < endBit; i++ {
			startMask <<= 1
			startMask |= 1
		}
		val = fmt.Sprintf("((%s >> %d) & 0x%X)", in1, startBit, startMask)
		// }
	} else {
		for i := startBit; i < 64; i++ {
			startMask <<= 1
			startMask |= 1
		}
		for i := 0; i < endBit; i++ {
			endMask <<= 1
			endMask |= 1
		}
		// val = fmt.Sprintf("int%d(uint%d(%s) >> %d)", dunb.dpn.dp.Bits, dunb.dpn.dp.Bits, in1, startBit)
		// if dunb.dpn.dp.Unsigned {
		val = fmt.Sprintf("(%s >> %d)", in1, startBit)
		// }
		if endBit > 0 {
			val = fmt.Sprintf("(%s | ((%s & 0x%X) << %d))", val, in2, endMask, nbBytes-endBit)
		}
	}
	return fmt.Sprintf("IT(%s << ntz)", val)
}
