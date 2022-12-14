package messages

import (
	"fmt"

	tss "github.com/binance-chain/tss-lib/tss"
	"github.com/libp2p/go-libp2p/core/peer"
)

type TSSMessageTpe uint8

const (
	TSSKeyGenMsg TSSMessageTpe = iota
	TSSKeySignMsg
	TSSTaskDone
	Unknown
)

func (msgType TSSMessageTpe) String() string {
	switch msgType {
	case TSSKeyGenMsg:
		return "TSSKeyGenMsg"
	case TSSKeySignMsg:
		return "TSSKeySignMsg"
	default:
		return "Unknown"

	}
}

type WrappedMessage struct {
	MessageType TSSMessageTpe `json:"message_type"`
	MsgID       string        `json:"message_id"`
	Payload     []byte        `json:"payload"`
}

type BroadcastMsgChan struct {
	WrappedMessage WrappedMessage
	PeersID        []peer.ID
}

type WireMessage struct {
	Routing   *tss.MessageRouting `json:"routing"`
	RoundInfo string              `json:"round_info"`
	Message   []byte              `json:"message"`
	Sig       []byte              `json:"signature"`
}

func (m *WireMessage) GetCacheKey() string {
	return fmt.Sprintf("%s-%s", m.Routing.From.Id, m.RoundInfo)
}

type TssTaskNotifier struct {
	TaskDone bool `json:"task_done"`
}
