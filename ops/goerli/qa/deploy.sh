#!/bin/bash

cp -r Makefile.goerli.qa.deploy ../../
cp -r docker-compose.goerli.qa.deploy.yml ../../

cd ../../
rm -rf .env
make -f Makefile.goerli.qa.deploy clean
make -f Makefile.goerli.qa.deploy up

rm -rf Makefile.goerli.qa.deploy
rm -rf docker-compose.goerli.qa.deploy.yml
