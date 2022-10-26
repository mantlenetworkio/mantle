package types

import (
	"context"
	"time"
)

type TssCommitteeInfo struct {
	ElectionId    uint64   `json:"election_id"`
	ClusterPubKey string   `json:"cluster_pub_key"`
	TssMembers    []string `json:"tss_members"`
	Threshold     int      `json:"threshold"`
}

type CpkData struct {
	Cpk          string    `json:"cpk"`
	ElectionId   uint64    `json:"election_id"`
	CreationTime time.Time `json:"creation_time"`
}

// Context ---------------------------------------------
type Context struct {
	ctx            context.Context
	requestId      string
	tssInfo        TssCommitteeInfo
	availableNodes []string
	approvers      []string
	electionId     uint64
	stateBatchRoot [32]byte
}

func NewContext() Context {
	return Context{
		ctx: context.Background(),
	}
}

func (c Context) RequestId() string {
	return c.requestId
}
func (c Context) TssInfos() TssCommitteeInfo {
	return c.tssInfo
}
func (c Context) AvailableNodes() []string {
	return c.availableNodes
}
func (c Context) Approvers() []string {
	return c.approvers
}
func (c Context) ElectionId() uint64 {
	return c.electionId
}
func (c Context) StateBatchRoot() [32]byte {
	return c.stateBatchRoot
}

func (c Context) WithRequestId(requestId string) Context {
	c.requestId = requestId
	return c
}

func (c Context) WithTssInfo(tssInfos TssCommitteeInfo) Context {
	c.tssInfo = tssInfos
	return c
}

func (c Context) WithAvailableNodes(nodes []string) Context {
	c.availableNodes = nodes
	return c
}

func (c Context) WithApprovers(nodes []string) Context {
	c.approvers = nodes
	return c
}

func (c Context) WithElectionId(election uint64) Context {
	c.electionId = election
	return c
}

func (c Context) WithStateBatchRoot(stateBatchRoot [32]byte) Context {
	c.stateBatchRoot = stateBatchRoot
	return c
}
