#!/bin/bash

export FABRIC_CFG_PATH=$PWD

ORG=$1

configtxgen -printOrg $ORG > "${ORG}.json"