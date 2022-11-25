# rpc-proxy

This tool implements `proxyd`, an RPC request router and proxy. It does the following things:

1. Whitelists RPC methods.
1. Routes RPC methods to groups of backend services.
1. Automatically retries failed backend requests.
1. Rate limit by IP(X-Forwarded-For header) and RPC method.
1. Cache for method `eth_chainId` `eth_getBlockByNumber` and so on.
1. Provides metrics the measure request latency, error rates, and the like.

## Usage

Run `make proxyd` to build the binary. No additional dependencies are necessary.

To configure `proxyd` for use, you'll need to create a configuration file to define your proxy backends and routing rules.  Check out [example.config.toml](./example.config.toml) for how to do this alongside a full list of all options with commentary.

Once you have a config file, start the daemon via `proxyd <path-to-config>.toml`.

## Use Case

### WS Whitelist Method
```toml
ws_method_whitelist = [
  "eth_subscribe",
  "eth_call",
  "eth_chainId"
]
```

### Enable Cache

```toml
[cache]
enabled = true
```
Method list which could be cached are predefined.
Cache are stored in redis if redis enabled, or, in memory.


### Enable Redis

1. store cache data
2. store rate limit data

```toml
[redis]
url = "redis://localhost:6379"
```

### Enable Rate limit
```toml
[rate_limit]
use_redis = true
base_rate = 10000
base_interval = "2s"
enable_backend_rate_limiter = true

# override for single rpc method
[rate_limit.method_overrides.eth_chainId]
limit = 100000
interval = "5s"
```

### Enable Metrics Exporter(Prometheus Format)
```toml
[metrics]
enabled = true
host = "0.0.0.0"
port = 9761
```

### Route Based on RPC Method
```toml
[rpc_method_mappings]
eth_sendTransaction = "l2geth"
eth_sendRawTransaction = "l2geth"

net_version = "verifier"
eth_chainId = "verifier"
...
```


## Metrics

See `metrics.go` for a list of all available metrics.

The metrics port is configurable via the `metrics.port` and `metrics.host` keys in the config.

## Adding Backend SSL Certificates in Docker

The Docker image runs on Alpine Linux. If you get SSL errors when connecting to a backend within Docker, you may need to add additional certificates to Alpine's certificate store. To do this, bind mount the certificate bundle into a file in `/usr/local/share/ca-certificates`. The `entrypoint.sh` script will then update the store with whatever is in the `ca-certificates` directory prior to starting `proxyd`.

