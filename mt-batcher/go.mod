module github.com/mantlenetworkio/mantle/mt-batcher

go 1.19

replace github.com/Layr-Labs/datalayr/common => ../datalayr-mantle/common

replace github.com/Layr-Labs/datalayr/lib/merkzg => ../datalayr-mantle/lib/merkzg

replace github.com/mantlenetworkio/mantle/l2geth v0.0.0 => ../l2geth

require (
<<<<<<< HEAD
	//github.com/Layr-Labs/datalayr/common v0.0.0-00010101000000-000000000000
	github.com/ethereum/go-ethereum v1.10.26
	github.com/mantlenetworkio/mantle/l2geth v0.0.0
=======
	github.com/Layr-Labs/datalayr/common v0.0.0-00010101000000-000000000000
	github.com/ethereum/go-ethereum v1.10.26
	github.com/go-resty/resty/v2 v2.7.0
	github.com/mantlenetworkio/mantle/l2geth v0.0.0
	github.com/mantlenetworkio/mantle/tss v0.0.0-20230313034302-2b6d5400976e
>>>>>>> feature/ctc-rollup-switch-to-eigenlayer
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.8.1
	github.com/syndtr/goleveldb v1.0.1-0.20210819022825-2ae1ddf74ef7
	github.com/urfave/cli v1.22.10
	google.golang.org/grpc v1.49.0
)

require (
	// github.com/Layr-Labs/datalayr/common v0.0.0-00010101000000-000000000000
	github.com/decred/dcrd/hdkeychain/v3 v3.0.0
	github.com/labstack/echo/v4 v4.9.0
	github.com/prometheus/client_golang v1.14.0
	github.com/tyler-smith/go-bip39 v1.1.0
)

//github.com/Layr-Labs/datalayr/common v0.0.0-00010101000000-000000000000
require github.com/shurcooL/graphql v0.0.0-20220606043923-3cf50f8a0a29

require github.com/Layr-Labs/datalayr/common v0.0.0-00010101000000-000000000000

require (
	github.com/Layr-Labs/datalayr/lib/merkzg v0.0.0-00010101000000-000000000000 // indirect
	github.com/VictoriaMetrics/fastcache v1.9.0 // indirect
	github.com/aristanetworks/goarista v0.0.0-20170210015632-ea17b1a17847 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/btcsuite/btcd v0.22.1 // indirect
	github.com/btcsuite/btcd/btcec/v2 v2.2.0 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/consensys/bavard v0.1.13 // indirect
	github.com/consensys/gnark-crypto v0.8.0 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/deckarep/golang-set v1.8.0 // indirect
	github.com/decred/base58 v1.0.3 // indirect
	github.com/decred/dcrd/crypto/blake256 v1.0.0 // indirect
	github.com/decred/dcrd/crypto/ripemd160 v1.0.1 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v3 v3.0.0 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.1.0 // indirect
	github.com/elastic/gosigar v0.14.2 // indirect
	github.com/fsnotify/fsnotify v1.5.4 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/hashicorp/golang-lru v0.5.5-0.20210104140557-80c98217689d // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/holiman/bloomfilter/v2 v2.0.3 // indirect
	github.com/holiman/uint256 v1.2.0 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/labstack/gommon v0.3.1 // indirect
	github.com/magiconair/properties v1.8.6 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mmcloughlin/addchain v0.4.0 // indirect
	github.com/olekukonko/tablewriter v0.0.5 // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_model v0.3.0 // indirect
	github.com/prometheus/common v0.37.0 // indirect
	github.com/prometheus/procfs v0.8.0 // indirect
	github.com/prometheus/tsdb v0.10.0 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/rjeczalik/notify v0.9.2 // indirect
	github.com/rs/cors v1.8.2 // indirect
	github.com/rs/zerolog v1.27.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/shirou/gopsutil v3.21.11+incompatible // indirect
	github.com/spf13/afero v1.8.2 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/cobra v1.5.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.10.1 // indirect
	github.com/steakknife/bloomfilter v0.0.0-20180922174646-6819c0d2a570 // indirect
	github.com/steakknife/hamming v0.0.0-20180906055917-c99c65617cd3 // indirect
<<<<<<< HEAD
=======
	github.com/subosito/gotenv v1.4.0 // indirect
>>>>>>> feature/ctc-rollup-switch-to-eigenlayer
	github.com/tklauser/go-sysconf v0.3.10 // indirect
	github.com/tklauser/numcpus v0.4.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.1 // indirect
	github.com/yusufpapurcu/wmi v1.2.2 // indirect
	golang.org/x/crypto v0.3.0 // indirect
	golang.org/x/net v0.3.0 // indirect
	golang.org/x/sys v0.3.0 // indirect
	golang.org/x/text v0.5.0 // indirect
	golang.org/x/time v0.0.0-20220922220347-f3bd1da661af // indirect
	google.golang.org/genproto v0.0.0-20211208223120-3a66f561d7aa // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/ini.v1 v1.66.6 // indirect
	gopkg.in/natefinch/npipe.v2 v2.0.0-20160621034901-c1b8fa8bdcce // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	rsc.io/tmplfunc v0.0.3 // indirect
)
