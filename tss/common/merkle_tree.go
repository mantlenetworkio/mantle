package common

import (
	"errors"
	"github.com/mantlenetworkio/mantle/l2geth/common/hexutil"
	"github.com/mantlenetworkio/mantle/l2geth/crypto"
)

var defaults = []string{
	"0x290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563",
	"0x633dc4d7da7256660a892f8f1604a44b5432649cc8ec5cb3ced4c4e6ac94dd1d",
	"0x890740a8eb06ce9be422cb8da5cdafc2b58c0a5e24036c578de2a433c828ff7d",
	"0x3b8ec09e026fdc305365dfc94e189a81b38c7597b3d941c279f042e8206e0bd8",
	"0xecd50eee38e386bd62be9bedb990706951b65fe053bd9d8a521af753d139e2da",
	"0xdefff6d330bb5403f63b14f33b578274160de3a50df4efecf0e0db73bcdd3da5",
	"0x617bdd11f7c0a11f49db22f629387a12da7596f9d1704d7465177c63d88ec7d7",
	"0x292c23a9aa1d8bea7e2435e555a4a60e379a5a35f3f452bae60121073fb6eead",
	"0xe1cea92ed99acdcb045a6726b2f87107e8a61620a232cf4d7d5b5766b3952e10",
	"0x7ad66c0a68c72cb89e4fb4303841966e4062a76ab97451e3b9fb526a5ceb7f82",
	"0xe026cc5a4aed3c22a58cbd3d2ac754c9352c5436f638042dca99034e83636516",
	"0x3d04cffd8b46a874edf5cfae63077de85f849a660426697b06a829c70dd1409c",
	"0xad676aa337a485e4728a0b240d92b3ef7b3c372d06d189322bfd5f61f1e7203e",
	"0xa2fca4a49658f9fab7aa63289c91b7c7b6c832a6d0e69334ff5b0a3483d09dab",
	"0x4ebfd9cd7bca2505f7bef59cc1c12ecc708fff26ae4af19abe852afe9e20c862",
	"0x2def10d13dd169f550f578bda343d9717a138562e0093b380a1120789d53cf10",
}

// GetMerkleRoot
// Calculates a merkle root for a list of 32-byte leaf hashes.  WARNING: If the number
// of leaves passed in is not a power of two, it pads out the tree with zero hashes.
// If you do not know the original length of elements for the tree you are verifying, then
// this may allow empty leaves past elements.length to pass a verification check down the line.
// Note that the _elements argument is modified, therefore it must not be used again afterwards
// @param elements Array of hashes from which to generate a merkle root.
// @return Merkle root of the leaves, with zero hashes for non-powers-of-two (see above).
func GetMerkleRoot(elements [][32]byte) ([32]byte, error) {
	if len(elements) == 0 {
		return [32]byte{}, errors.New("Must provide at least one element. ")
	}
	if len(elements) == 1 {
		return elements[0], nil
	}

	// We'll need to keep track of left and right siblings.
	var leftSibling [32]byte
	var rightSibling [32]byte

	// Number of non-empty nodes at the current depth.
	rowSize := len(elements)
	// Current depth, counting from 0 at the leaves
	depth := 0

	// Common sub-expressions
	var halfRowSize int   // rowSize / 2
	var rowSizeIsOdd bool // rowSize % 2 == 1

	buf := make([]byte, 64)
	for rowSize > 1 {
		halfRowSize = rowSize / 2
		rowSizeIsOdd = rowSize%2 == 1

		for i := 0; i < halfRowSize; i++ {
			leftSibling = elements[(2 * i)]
			rightSibling = elements[(2*i)+1]

			copy(buf[0:32], leftSibling[:])
			copy(buf[32:64], rightSibling[:])
			elements[i] = crypto.Keccak256Hash(buf)
		}

		if rowSizeIsOdd {
			leftSibling = elements[rowSize-1]
			bz, _ := hexutil.Decode(defaults[depth])

			copy(buf[0:32], leftSibling[:])
			copy(buf[32:64], bz)
			elements[halfRowSize] = crypto.Keccak256Hash(buf)
		}

		rowSize = halfRowSize
		if rowSizeIsOdd {
			rowSize += 1
		}
		depth++
	}
	return elements[0], nil
}
