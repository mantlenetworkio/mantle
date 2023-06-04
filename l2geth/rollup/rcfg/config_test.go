package rcfg

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/mantlenetworkio/mantle/l2geth/common"
)

func TestMantleTokenSymbol(t *testing.T) {
	value := "MNT"
	// string	->	slot
	// MNT 		-> 	[77 78 84 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 6]
	fmt.Println(common.BytesToHash(packStringForSlot(value)))
	fmt.Println(packStringForSlot(value))

	targetResult := []byte{77, 78, 84, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 6}
	if !bytes.Equal(packStringForSlot(value), targetResult) {
		t.Fatal("packStringForSlot error")
	}
}
