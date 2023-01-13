package mt_meso_service

import (
	"github.com/mantlenetworkio/mantle/mt-meso-service/flags"
	"github.com/urfave/cli"
	"time"
)

type Config struct {
	EthRpc               string
	EigenDaHttpPort      int
	EigenContractAddress string
	GraphProvider        string
	RetrieverSocket      string
	Timeout              time.Duration
	DisableHTTP2         bool
	EchoDebug            bool
}

func NewConfig(ctx *cli.Context) (Config, error) {
	cfg := Config{
		EthRpc:               ctx.GlobalString(flags.EthRpcFlag.Name),
		EigenDaHttpPort:      ctx.GlobalInt(flags.EigenDaHttpPortFlag.Name),
		EigenContractAddress: ctx.GlobalString(flags.EigenContractAddressFlag.Name),
		RetrieverSocket:      ctx.GlobalString(flags.RetrieverSocketFlag.Name),
		DisableHTTP2:         ctx.GlobalBool(flags.HTTP2DisableFlag.Name),
		EchoDebug:            ctx.GlobalBool(flags.EchoDebugFlag.Name),
	}
	return cfg, nil
}
