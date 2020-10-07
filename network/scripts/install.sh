#!/bin/bash

. $PWD/channel.sh

installChannel
anchorPeerUpdate orgMSP

. $PWD/chaincode.sh

prepareChaincode "consumption:7054"
installAndApprove
commitChanicode