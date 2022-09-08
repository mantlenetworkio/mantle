#!/bin/sh


## install base-dep
sudo apt-get update -y && sudo apt-get install curl make jq git software-properties-common docker.io -y
sudo add-apt-repository -y ppa:ethereum/ethereum
sudo apt-get update
sudo apt-get install ethereum -y

##
echo $PASSWORD > password.txt
geth account new --password password.txt
geth account new --password password.txt
geth account new --password password.txt
mkdir -p /root/.ethereum/key

## clone the ethkey-parse
git clone https://github.com/mantlenetworkio/ethkey-parse

## docker build ethkey-parse binary
docker run -it -v "$PWD/ethkey-parse":/usr/src/ethkey-parse -w /usr/src/ethkey-parse golang go build -v

## upload key
## the private keys are in the /root/.ethereum/key



## parse ethkey
cd ethkey-parse
./ethkey-parse
cd ..

## clear
rm -rf password.txt
