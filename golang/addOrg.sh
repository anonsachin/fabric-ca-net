#!/bin/bash

export FABRIC_CFG_PATH=$PWD

ORG=$1

configtxgen -printOrg $ORG > "${ORG}.json"

export CORE_PEER_LOCALMSPID="orgMSP"
export CORE_PEER_TLS_ROOTCERT_FILE=/home/sachin/ca-net/network/certs/peerorg/admin/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=/home/sachin/ca-net/network/certs/peerorg/admin/msp
export CORE_PEER_ADDRESS=127.0.0.1:7051