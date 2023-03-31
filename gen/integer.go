package main

import "fmt"

type deltapackint struct {
	BitLen  int
	WithNtz bool
}

func (dp *deltapackint) Dpn() []deltapackN {
	dpn := make([]deltapackN, dp.BitLen)
	for i := range dpn {
		dpn[i] = deltapackN{dp, i}
	}
	return dpn
}

type deltapackN struct {
	dp   *deltapackint
	BitN int
}

func (dpn *deltapackN) Dpnb() []deltapackNByte {
	dpnb := make([]deltapackNByte, dpn.BitN)
	for i := range dpnb {
		dpnb[i] = deltapackNByte{dpn, i}
	}
	return dpnb
}

func (dpn *deltapackN) Dunb() []deltaunpackNByte {
	dunb := make([]deltaunpackNByte, dpn.dp.BitLen)
	for i := range dunb {
		dunb[i] = deltaunpackNByte{dpn, i}
	}
	return dunb
}

type deltapackNByte struct {
	dpn  *deltapackN
	I int
}

func (dpnb *deltapackNByte) PackLines() []string {
	return dpnb.packLines(false, false)
}

func (dpnb *deltapackNByte) PackLinesZigZag() []string {
	return dpnb.packLines(true, false)
}

func (dpnb *deltapackNByte) PackLinesNtz() []string {
	return dpnb.packLines(false, true)
}

func (dpnb *deltapackNByte) PackLinesZigZagNtz() []string {
	return dpnb.packLines(true, true)
}

func (dpnb *deltapackNByte) packLines(zigzag, ntz bool) []string {
	var lines []string
	bitN := dpnb.dpn.BitN
	if bitN == 0 {
		return nil
	}
	var overlapping bool
	minOffset := dpnb.dpn.dp.BitLen * dpnb.I
	maxOffset := dpnb.dpn.dp.BitLen * (dpnb.I + 1)
	if minOffset%bitN != 0 {
		minOffset -= bitN
		overlapping = true
	}
	for i := 0; i < dpnb.dpn.dp.BitLen; i++ {
		offset := i * bitN
		if offset < minOffset || offset >= maxOffset {
			continue
		}

		ishift := offset - minOffset
		if overlapping {
			ishift -= bitN
		}

		v1 := fmt.Sprintf("in[%d]", i)
		v2 := "initoffset"
		if i > 0 {
			v2 = fmt.Sprintf("in[%d]", i-1)
		}
		diff := fmt.Sprintf("%s - %s", v1, v2)
		if zigzag {
			// force signed
			diff = fmt.Sprintf("int%d(%s)", dpnb.dpn.dp.BitLen, diff)
			// zigzag encoding of v: (v << 1) ^ (v >> 31)
			diff = fmt.Sprintf("((%s) << 1) ^ ((%s) >> %d)", diff, diff, dpnb.dpn.dp.BitLen-1)
		}
		var line string
		switch {
		case ishift < 0:
			if ntz {
				line = fmt.Sprintf("(%s) >> ((%d+ntz)&%d)", diff, -ishift, dpnb.dpn.dp.BitLen-1)
			} else {
				line = fmt.Sprintf("(%s) >> %d", diff, -ishift)
			}
		case ishift > 0:
			if ntz {
				line = fmt.Sprintf("((%s) >> ntz << %d)", diff, ishift)
			} else {
				line = fmt.Sprintf("((%s) << %d)", diff, ishift)
			}
		default:
			if ntz {
				line = fmt.Sprintf("(%s) >> ntz", diff)
			} else {
				line = diff
			}
		}
		if offset+bitN < maxOffset {
			line += " | "
		}

		lines = append(lines, line)
	}
	return lines
}

type deltaunpackNByte struct {
	dpn  *deltapackN
	I int
}

func (dunb *deltaunpackNByte) UnpackLine() string {
	return dunb.unpackLine(false, false)
}

func (dunb *deltaunpackNByte) UnpackLineZigZag() string {
	return dunb.unpackLine(true, false)
}

func (dunb *deltaunpackNByte) UnpackLineNtz() string {
	return dunb.unpackLine(false, true)
}

func (dunb *deltaunpackNByte) UnpackLineZigZagNtz() string {
	return dunb.unpackLine(true, true)
}

func (dunb *deltaunpackNByte) unpackLine(zigzag, ntz bool) string {
	bitN := dunb.dpn.BitN
	if bitN == 0 {
		return "initoffset"
	}
	startByte := dunb.I * bitN / dunb.dpn.dp.BitLen
	startBit := dunb.I * bitN % dunb.dpn.dp.BitLen
	endByte := (dunb.I + 1) * bitN / dunb.dpn.dp.BitLen
	endBit := (dunb.I + 1) * bitN % dunb.dpn.dp.BitLen

	var startMask, endMask int
	in1 := fmt.Sprintf("in[%d]", startByte)
	in2 := fmt.Sprintf("in[%d]", endByte)
	out := "initoffset"
	if dunb.I > 0 {
		out = fmt.Sprintf("out[%d]", dunb.I-1)
	}
	var val string
	if startByte == endByte {
		for i := startBit; i < endBit; i++ {
			startMask <<= 1
			startMask |= 1
		}
		// val = fmt.Sprintf("(int%d(uint%d(%s) >> %d)) & 0x%X", dunb.dpn.dp.BitLen, dunb.dpn.dp.BitLen, in1, startBit, startMask)
		val = fmt.Sprintf("(%s >> %d) & 0x%X", in1, startBit, startMask)
	} else {
		for i := startBit; i < dunb.dpn.dp.BitLen; i++ {
			startMask <<= 1
			startMask |= 1
		}
		for i := 0; i < endBit; i++ {
			endMask <<= 1
			endMask |= 1
		}
		// val = fmt.Sprintf("int%d(uint%d(%s) >> %d)", dunb.dpn.dp.BitLen, dunb.dpn.dp.BitLen, in1, startBit)
		val = fmt.Sprintf("(%s >> %d)", in1, startBit)
		if endBit > 0 {
			val = fmt.Sprintf("(%s | ((%s & 0x%X) << %d))", val, in2, endMask, bitN-endBit)
		}
	}
	if zigzag {
		// force signed
		val = fmt.Sprintf("int%d(%s)", dunb.dpn.dp.BitLen, val)
		// zigzag decoding of val: (-(val & 1))^(val>>1))
		val = fmt.Sprintf("((-((%s) & 1))^((%s)>>1))", val, val)
	}
	if ntz {
		return fmt.Sprintf("T(%s) << ntz + %s", val, out)
	} else {
		return fmt.Sprintf("T(%s) + %s", val, out)
	}
}
