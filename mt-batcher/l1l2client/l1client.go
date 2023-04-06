package l1l2client

import (
	"context"
	"crypto/tls"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
)

func L1EthClientWithTimeout(ctx context.Context, url string, disableHTTP2 bool) (
	*ethclient.Client, error) {
	ctxt, cancel := context.WithTimeout(ctx, DefaultTimeout)
	defer cancel()
	if strings.HasPrefix(url, "http") {
		httpClient := new(http.Client)
		if disableHTTP2 {
			log.Info("Disabled HTTP/2 support in L1 eth client")
			httpClient.Transport = &http.Transport{
				TLSNextProto: make(map[string]func(authority string, c *tls.Conn) http.RoundTripper),
			}
		}
		rpcClient, err := rpc.DialHTTPWithClient(url, httpClient)
		if err != nil {
			return nil, err
		}
		return ethclient.NewClient(rpcClient), nil
	}
	return ethclient.DialContext(ctxt, url)
}
