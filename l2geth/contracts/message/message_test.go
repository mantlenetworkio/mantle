package message

import (
	"fmt"
	"log"
	"testing"

	"github.com/mantlenetworkio/mantle/l2geth/common/hexutil"
)

func TestUnpacket(t *testing.T) {
	hexData := "0xcbd4ece9000000000000000000000000deaddeaddeaddeaddeaddeaddeaddeaddead2222000000000000000000000000d9e2f450525079e1e29fb23bc7caca6f61f8fd4a0000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000000e0000000000000000000000000000000000000000000000000000000000000024f523f40d00000000000000000000000000000000000000000000000000000000000003e800000000000000000000000000000000000000000000000000000000"
	data := &Data{}
	decodedData, err := hexutil.Decode(hexData)
	if err != nil {
		log.Fatal(err)
	}
	if err := data.UnPackData(decodedData); err != nil {
		log.Fatal(err)
	}
}

func TestRollbackData_UnPackData(t *testing.T) {
	hexData := "0xf523f40d000000000000000000000000000000000000000000000000000000000000000e"
	data := &RollbackData{}
	decodedData, err := hexutil.Decode(hexData)
	if err != nil {
		log.Fatal(err)
	}
	if err := data.UnPackData(decodedData); err != nil {
		log.Fatal(err)
	}
}

func TestData_UnPackData(t *testing.T) {
	ch := make(chan int, 100)
	//ch <- 1

	for tx := range ch {
		if len(ch) == 0 {
			break
		}
		fmt.Println(tx)
	}
}
