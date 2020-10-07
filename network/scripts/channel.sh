#!/bin/bash

function installChannel(){
    echo "######### Creating the Channel Block ##############"
    set -x
    export CHANNEL_NAME="testchannel"
    peer channel create -o orderer:7050 -c $CHANNEL_NAME --ordererTLSHostnameOverride orderer -f ../channel-artifacts/${CHANNEL_NAME}.tx --outputBlock ../channel-artifacts/${CHANNEL_NAME}.block --tls --cafile $CORE_PEER_TLS_ROOTCERT_FILE
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
    peer channel update -o orderer:7050 -c $CHANNEL_NAME --ordererTLSHostnameOverride orderer -f ../channel-artifacts/${ORG}.tx  --tls --cafile $CORE_PEER_TLS_ROOTCERT_FILE
    set +x
}