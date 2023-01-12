package mt_meso_service

import (
	"github.com/mantlenetworkio/mantle/mt-challenger/flags"
	"github.com/urfave/cli"
)

type Config struct {
	EigenContractAddress string
	GraphProvider        string
	RetrieverSocket      string
	DisableHTTP2         bool
}

func NewConfig(ctx *cli.Context) (Config, error) {
	cfg := Config{
		GraphProvider:        ctx.GlobalString(flags.GraphProviderFlag.Name),
		EigenContractAddress: ctx.GlobalString(flags.EigenContractAddressFlag.Name),
		RetrieverSocket:      ctx.GlobalString(flags.RetrieverSocketFlag.Name),
		DisableHTTP2:         ctx.GlobalBool(flags.HTTP2DisableFlag.Name),
	}
	return cfg, nil
}
