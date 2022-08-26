package vm

import (
	"os"
)

// UsingBVM is used to enable or disable functionality necessary for the bvm.
var UsingBVM bool

func init() {
	UsingBVM = os.Getenv("USING_bvm") == "true"
}
