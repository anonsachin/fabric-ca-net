#!/bin/bash

function installChannel(){
    echo "######### Creating the Channel Block ##############"
    set -x
    export CHANNEL_NAME="testchannel"
    peer channel create -o orderer.testnetwork.com:7050 -c $CHANNEL_NAME --ordererTLSHostnameOverride orderer.testnetwork.com -f ../channel-artifacts/${CHANNEL_NAME}.tx --outputBlock ../channel-artifacts/${CHANNEL_NAME}.block --tls --cafile $ORDERER_CA_CERT
    set +x

    echo "######### Joining the Channel ##############"

    set -x
    peer channel join -b ../channel-artifacts/$CHANNEL_NAME.block
    set +x

}

function anchorPeerUpdate(){
    ORG=$1
    # DEPENDENT ON CHANNEL_NAME
    echo "######### Updating the Anchor Peer ##############"
    set -x
    peer channel update -o orderer.testnetwork.com:7050 -c $CHANNEL_NAME --ordererTLSHostnameOverride orderer.testnetwork.com -f ../channel-artifacts/${ORG}.tx  --tls --cafile $ORDERER_CA_CERT
    set +x
}