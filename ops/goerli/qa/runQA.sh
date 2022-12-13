#!/bin/bash


cp -r Makefile.goerli.qa ../../
cp -r docker-compose.goerli.qa.yml ../../

cd ../../
rm -rf .env
make -f Makefile.goerli.qa clean
make -f Makefile.goerli.qa up

rm -rf Makefile.goerli.qa
rm -rf docker-compose.goerli.qa.yml
