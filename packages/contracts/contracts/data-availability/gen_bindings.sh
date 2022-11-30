#!/bin/bash

# to run this cd in here
mkdir -p data

dirname="EigenDataLayrRollup"
for name in $dirname; do

    mkdir -p ../common/contracts/bindings/${name}

    mkdir -p ./bindings/${name}/
    touch ./bindings/${name}/binding.go

    dl_json="./out/${name}.sol/${name}.json"
    solc_abi=$(cat ${dl_json} | jq -r '.abi')
    solc_bin=$(cat ${dl_json} | jq -r '.bytecode.object')
    echo ${solc_abi} > data/DL.abi
    echo ${solc_bin} > data/DL.bin

    rm -f ../common/contracts/bindings/${name}/binding.go

    abigen --bin=data/DL.bin --abi=data/DL.abi --pkg=contract${name} --out=./bindings/${name}/binding.go
    echo "generated bindings for ${name}"
done





