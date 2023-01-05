module github.com/mantlenetworkio/mantle/mt-challenger

go 1.19

replace github.com/Layr-Labs/datalayr/common => ../datalayr-mantle/common
replace github.com/Layr-Labs/datalayr/lib/merkzg => ../datalayr-mantle/lib/merkzg

require (
	github.com/Layr-Labs/datalayr/common v0.0.0-00010101000000-000000000000
	github.com/urfave/cli v1.22.10
)

require (
	github.com/cpuguy83/go-md2man/v2 v2.0.0-20190314233015-f79a8a8ca69d // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/rs/zerolog v1.27.0 // indirect
	github.com/russross/blackfriday/v2 v2.0.1 // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	golang.org/x/sys v0.0.0-20220727055044-e65921a090b8 // indirect
)
