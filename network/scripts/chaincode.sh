#!/bin/bash

function prepareChaincode(){

    export CONNECT=$1

    echo "######### Creating the Chaincode JSON's ##############"
    set -x
    jq -n '{"address":env.CONNECT,"dial_timeout": "10s","tls_required": false}' > connection.json
    jq -n '{"path":"","type":"external","label":"consumption"}' > metadata.json
    set +x

    echo "######### Creating the Chaincode TAR's ##############"
    set -x
    tar cfz code.tar.gz connection.json
    tar cfz consumption.tgz metadata.json code.tar.gz
    set +x
}

function installAndApprove(){

    echo "######### Install the Chaincode ##############"
    set -x
    peer lifecycle chaincode install ./consumption.tgz
    set +x

    echo "######### Approve the Chaincode ##############"
    set -x
    peer lifecycle chaincode queryinstalled
    export CC_PACKAGE_ID=$(peer lifecycle chaincode queryinstalled | grep Package | cut -d \  -f 3 | cut -d , -f 1)
    peer lifecycle chaincode approveformyorg -o orderer:7050 --ordererTLSHostnameOverride orderer --channelID $CHANNEL_NAME --name cpu --version 1.0 --package-id $CC_PACKAGE_ID --sequence 1 \
--tls --cafile $CORE_PEER_TLS_ROOTCERT_FILE
    set +x
}

function commitChanicode(){
    echo "######### Commiting the Chaincode ##############"

    set -x
    peer lifecycle chaincode checkcommitreadiness --channelID $CHANNEL_NAME --name cpu --version 1.0 --sequence 1 --output json
    peer lifecycle chaincode commit -o orderer:7050 --ordererTLSHostnameOverride orderer --channelID $CHANNEL_NAME --name cpu --version 1.0 --sequence 1 \
--tls --cafile $CORE_PEER_TLS_ROOTCERT_FILE
    set +x
}

# prepareChaincode "consumption:7054"