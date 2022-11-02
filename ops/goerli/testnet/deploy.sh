#!/bin/bash

cp -r Makefile.goerli.testnet.deploy ../../
cp -r docker-compose.goerli.testnet.deploy.yml ../../

cd ../../
rm -rf .env
make -f Makefile.goerli.testnet.deploy clean
make -f Makefile.goerli.testnet.deploy up

rm -rf Makefile.goerli.testnet.deploy
rm -rf docker-compose.goerli.testnet.deploy.yml
