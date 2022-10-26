#!/bin/bash

cp -r Makefile.goerli.mantle ../../
cp -r docker-compose.goerli.mantle.yml ../../

cd ../../
rm -rf .env
make -f Makefile.goerli.mantle clean
make -f Makefile.goerli.mantle up

rm -rf Makefile.goerli.mantle
rm -rf docker-compose.goerli.mantle.yml
