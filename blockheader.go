package intcomp

// BlockHeader is a uint64 used to store the number of bits in a group (7 bits) and the number of trailing zeros (6 bits).
// up to 4 groups can be stored in a block header.
// the number of groups are stored in the first 4 bits of the header.
type BlockHeader uint64

const MaxGroups = 4

// GroupCount returns the number of groups in the block header.
func (bh BlockHeader) GroupCount() int {
	return int(bh >> 60)
}

// AddGroup adds a group of bits to the block header.
// bitlen is the number of bits in the group, and ntz is the number of trailing zeros.
// bitlen and ntz are stored in the header as 6-bit values.
// adding a group increases the number of groups in the header by 1.
func (bh BlockHeader) AddGroup(bitlen, ntz int) BlockHeader {
	groupNum := bh >> 60
	// check that the number of groups is less than the maximum
	if groupNum >= MaxGroups {
		panic("too many groups")
	}
	// increment the number of groups and
	return (groupNum+1)<<60 | bh&0x0fffffffffffffff | BlockHeader(bitlen&0x7f<<6|ntz&0x3f)<<(groupNum*13)
}

// GetGroup returns the bitlen and ntz of the group at index i.
func (bh BlockHeader) GetGroup(i int) (int, int) {
	return int((bh >> (i*13 + 6)) & 0x7f), int((bh >> (i * 13)) & 0x3f)
}

// BlockLen returns the number of compressed uint64 in the block
// not including the header.
func (bh BlockHeader) BlockLen() int {
	var res int
	for i := 0; i < bh.GroupCount(); i++ {
		bitlen, ntz := bh.GetGroup(i)
		res += (bitlen - ntz)
	}
	return res
}
