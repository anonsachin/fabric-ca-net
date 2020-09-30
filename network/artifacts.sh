#!/bin/bash

export FABRIC_CFG_PATH=${PWD}
export CHANNEL_NAME="testchannel"

if [ ! -d channel-artifacts ]; then
       mkdir channel-artifacts
fi       

function genesis() {
	echo "#########  Generating Orderer Genesis block ##############"

	set -x
	configtxgen -profile OrdererGenesis -channelID system-channel -outputBlock ./channel-artifacts/genesis.block
	set +x
}

function channel() {
    echo "#########  Generating Channel Creation Tx ##############"
	set -x
 	configtxgen -profile OrgChannel -outputCreateChannelTx ./channel-artifacts/${CHANNEL_NAME}.tx -channelID $CHANNEL_NAME
	set +x
}

function anchorPeer() {
	ORG=$1
	echo "#########  Generating Anchor Peer Update Tx ##############"
	set -x
 	configtxgen -profile OrgChannel -outputAnchorPeersUpdate ./channel-artifacts/${ORG}.tx -channelID $CHANNEL_NAME -asOrg $ORG
	set +x
}

genesis
channel
anchorPeer orgMSP
