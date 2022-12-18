package message

import (
	"fmt"
	"github.com/mantlenetworkio/mantle/l2geth/accounts/abi"
	"github.com/mantlenetworkio/mantle/l2geth/common"
	"math/big"
	"strings"
)

var (
	messageABI *abi.ABI
	sccABI     *abi.ABI
)

const (
	jsondata    = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1CrossDomainMessenger\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"msgHash\",\"type\":\"bytes32\"}],\"name\":\"FailedRelayedMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"msgHash\",\"type\":\"bytes32\"}],\"name\":\"RelayedMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"SentMessage\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"l1CrossDomainMessenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_messageNonce\",\"type\":\"uint256\"}],\"name\":\"relayMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"relayedMessages\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"_gasLimit\",\"type\":\"uint32\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"sentMessages\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"successfulMessages\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"xDomainMessageSender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"
	sccJsondata = "[{\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_shouldRollBack\",\"type\":\"uint256\"}],\"name\":\"rollBackMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
)

func messageAbi() *abi.ABI {
	if messageABI != nil {
		return messageABI
	}
	mABI, err := abi.JSON(strings.NewReader(jsondata))
	if err != nil {
		panic(err)
	}
	messageABI = &mABI
	return messageABI
}

func sccAbi() *abi.ABI {
	if sccABI != nil {
		return sccABI
	}
	sABI, err := abi.JSON(strings.NewReader(sccJsondata))
	if err != nil {
		panic(err)
	}
	sccABI = &sABI
	return sccABI
}

type Data struct {
	Target       common.Address `abi:"_target"`
	Sender       common.Address `abi:"_sender"`
	Message      []byte         `abi:"_message"`
	MessageNonce *big.Int       `abi:"_messageNonce"`
}

func (d *Data) UnPackData(bytesData []byte) error {
	if len(bytesData) < 4 {
		return fmt.Errorf("UnPacketData err:invalid bytesData lenght")
	}
	messageABI = messageAbi()
	return messageABI.Methods["relayMessage"].Inputs.Unpack(d, bytesData[4:])
}

type RollbackData struct {
	ShouldRollBack *big.Int `abi:"_shouldRollBack"`
}

func (rd *RollbackData) UnPackData(bytesData []byte) error {
	if len(bytesData) < 4 {
		return fmt.Errorf("UnPacketData err:invalid bytesData lenght")
	}
	sccABI = sccAbi()
	return sccABI.Methods["rollBackMessage"].Inputs.Unpack(rd, bytesData[4:])
}
