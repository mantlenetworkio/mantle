<!--
parent:
  order: false
-->

<div align="center">
  <h1> DA Challenger </h1>
</div>

mt-challenger is a project about the fraud proof for EigenDA of Mantle Network

**Tips**: need [Go 1.18+](https://golang.org/dl/)

## Install

### 1.Install dependencies
```bash
go mod tidy
```

### 2.build binary
```bash
make datalayr-mantle
make binding
make mt-challenger
```

### 3.build docker
```bash
make datalayr-mantle
docker build -t challenger:latest -f ./Dockerfile .
```

### 4.start binary

Config .env, You can refer to .env_example, if config finished, you can exec following command

```bash
./mt-challenger
```

