module github.com/mantlenetworkio/mantle/l2geth

go 1.18

replace github.com/mantlenetworkio/mantle/fraud-proof => ../fraud-proof

replace github.com/mantlenetworkio/mantle/bss-core => ../bss-core

replace github.com/mantlenetworkio/mantle/metrics => ../metrics

require (
	github.com/Azure/azure-storage-blob-go v0.7.0
	github.com/VictoriaMetrics/fastcache v1.9.0
	github.com/aristanetworks/goarista v0.0.0-20170210015632-ea17b1a17847
	github.com/aws/aws-sdk-go v1.42.6
	github.com/btcsuite/btcd v0.22.1
	github.com/cespare/cp v0.1.0
	github.com/cloudflare/cloudflare-go v0.14.0
	github.com/davecgh/go-spew v1.1.1
	github.com/deckarep/golang-set v1.8.0
	github.com/docker/docker v20.10.10+incompatible
	github.com/edsrzf/mmap-go v1.1.0
	github.com/elastic/gosigar v0.14.2
	github.com/ethereum/go-ethereum v1.10.26
	github.com/fatih/color v1.13.0
	github.com/fjl/memsize v0.0.1
	github.com/gballet/go-libpcsclite v0.0.0-20191108122812-4678299bea08
	github.com/go-resty/resty/v2 v2.7.0
	github.com/go-stack/stack v1.8.1
	github.com/golang/protobuf v1.5.3
	github.com/golang/snappy v0.0.4
	github.com/gorilla/websocket v1.5.0
	github.com/graph-gophers/graphql-go v1.3.0
	github.com/hashicorp/golang-lru v0.5.5-0.20210104140557-80c98217689d
	github.com/huin/goupnp v1.0.3
	github.com/influxdata/influxdb v1.8.3
	github.com/jackpal/go-nat-pmp v1.0.2
	github.com/jarcoal/httpmock v1.0.8
	github.com/julienschmidt/httprouter v1.3.0
	github.com/karalabe/usb v0.0.2
	github.com/mantlenetworkio/mantle/fraud-proof v0.0.0
	github.com/mattn/go-colorable v0.1.13
	github.com/mattn/go-isatty v0.0.16
	github.com/naoina/toml v0.1.2-0.20170918210437-9fafd6967416
	github.com/olekukonko/tablewriter v0.0.5
	github.com/pborman/uuid v1.2.0
	github.com/peterh/liner v1.1.1-0.20190123174540-a2c9a5303de7
	github.com/pkg/errors v0.9.1
	github.com/prometheus/tsdb v0.10.0
	github.com/rjeczalik/notify v0.9.2
	github.com/robertkrimen/otto v0.0.0-20191219234010-c382bd3c16ff
	github.com/rs/cors v1.8.2
	github.com/status-im/keycard-go v0.2.0
	github.com/steakknife/bloomfilter v0.0.0-20180922174646-6819c0d2a570
	github.com/stretchr/testify v1.8.4
	github.com/syndtr/goleveldb v1.0.1-0.20210819022825-2ae1ddf74ef7
	github.com/tyler-smith/go-bip39 v1.1.0
	github.com/wsddn/go-ecdh v0.0.0-20161211032359-48726bab9208
	golang.org/x/crypto v0.9.0
	golang.org/x/sync v0.2.0
	golang.org/x/sys v0.8.0
	golang.org/x/text v0.9.0
	golang.org/x/time v0.3.0
	gopkg.in/natefinch/npipe.v2 v2.0.0-20160621034901-c1b8fa8bdcce
	gopkg.in/olebedev/go-duktape.v3 v3.0.0-20200619000410-60c24ae608a6
	gopkg.in/urfave/cli.v1 v1.20.0
)

require (
	cloud.google.com/go/compute v1.19.3 // indirect
	cloud.google.com/go/compute/metadata v0.2.3 // indirect
	cloud.google.com/go/iam v1.0.1 // indirect
	cloud.google.com/go/kms v1.11.0 // indirect
	github.com/Azure/azure-pipeline-go v0.2.1 // indirect
	github.com/Azure/go-autorest/autorest/adal v0.9.23 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/btcsuite/btcd/btcec/v2 v2.2.0 // indirect
	github.com/cbergoon/merkletree v0.2.0 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/decred/base58 v1.0.3 // indirect
	github.com/decred/dcrd/crypto/blake256 v1.0.0 // indirect
	github.com/decred/dcrd/crypto/ripemd160 v1.0.1 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v3 v3.0.0 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.1.0 // indirect
	github.com/decred/dcrd/hdkeychain/v3 v3.0.0 // indirect
	github.com/getsentry/sentry-go v0.12.0 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/s2a-go v0.1.4 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.2.3 // indirect
	github.com/googleapis/gax-go/v2 v2.10.0 // indirect
	github.com/holiman/bloomfilter/v2 v2.0.3 // indirect
	github.com/holiman/uint256 v1.2.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/mantlenetworkio/mantle/bss-core v0.0.0 // indirect
	github.com/mantlenetworkio/mantle/metrics v0.0.0 // indirect
	github.com/mattn/go-ieproxy v0.0.0-20190610004146-91bb50d98149 // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/naoina/go-stringutil v0.1.0 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_golang v1.15.1 // indirect
	github.com/prometheus/client_model v0.3.0 // indirect
	github.com/prometheus/common v0.42.0 // indirect
	github.com/prometheus/procfs v0.9.0 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/shirou/gopsutil v3.21.11+incompatible // indirect
	github.com/specularl2/specular/clients/geth/specular v0.0.0-20221120100224-5e02437e9455 // indirect
	github.com/steakknife/hamming v0.0.0-20180906055917-c99c65617cd3 // indirect
	github.com/tklauser/go-sysconf v0.3.10 // indirect
	github.com/tklauser/numcpus v0.4.0 // indirect
	github.com/yusufpapurcu/wmi v1.2.2 // indirect
	go.opencensus.io v0.24.0 // indirect
	golang.org/x/net v0.10.0 // indirect
	golang.org/x/oauth2 v0.8.0 // indirect
	golang.org/x/term v0.8.0 // indirect
	google.golang.org/api v0.126.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20230530153820-e85fd2cbaebc // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20230530153820-e85fd2cbaebc // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230530153820-e85fd2cbaebc // indirect
	google.golang.org/grpc v1.55.0 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
	gopkg.in/sourcemap.v1 v1.0.5 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gotest.tools/v3 v3.4.0 // indirect
)
