package types

import (
	"context"
	"time"
)

type TssInfos struct {
	ClusterPubKey string   `json:"cluster_pub_key"`
	PartyPubKeys  []string `json:"party_pub_keys"`
	Threshold     int      `json:"threshold"`
}

type CpkData struct {
	Cpk          string    `json:"cpk"`
	ElectionId   uint64    `json:"election_id"`
	CreationTime time.Time `json:"creation_time"`
}

type Context struct {
	ctx            context.Context
	requestId      string
	tssInfo        TssInfos
	availableNodes []string
	approvers      []string
}

func NewContext() Context {
	return Context{
		ctx: context.Background(),
	}
}

func (c Context) RequestId() string {
	return c.requestId
}
func (c Context) TssInfos() TssInfos {
	return c.tssInfo
}
func (c Context) AvailableNodes() []string {
	return c.availableNodes
}
func (c Context) Approvers() []string {
	return c.approvers
}

func (c Context) WithRequestId(requestId string) Context {
	c.requestId = requestId
	return c
}

func (c Context) WithTssInfo(tssInfos TssInfos) Context {
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
