#!/usr/bin/env bash

RED='\x1b[0;31m'
GREEN='\x1b[0;32m'
NO_COLOR='\x1b[0m'

set -E
cd node_modules/web3/dist

if [ command -v go-bindata >/dev/null 2>&1 ]; then
    echo -e "${RED}go-bindata is required to build bindata.go . Aborting...${NO_COLOR}" >&2
else
    go-bindata -pkg="abi" web3.min.js
    mv bindata.go ../../..
fi

