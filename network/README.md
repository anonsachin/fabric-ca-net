Docker Network With Fabric CA's for tls and identities
======================================================

simple network with one ca for tls and two ca's for orderer and peer

commands for channel creation and join
--------------------------------------
* creating channel block
`` 
export CHANNEL_NAME="testchannel"
peer channel create -o orderer:7050 -c $CHANNEL_NAME --ordererTLSHostnameOverride orderer -f ./channel-artifacts/${CHANNEL_NAME}.tx --outputBlock ./channel-artifacts/${CHANNEL_NAME}.block --tls --cafile $CORE_PEER_TLS_ROOTCERT_FILE
``
* joining the channel
`` 
peer channel join -b ./channel-artifacts/$CHANNEL_NAME.block
``
