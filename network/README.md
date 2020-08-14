Docker Network With Fabric CA's for tls and identities
======================================================

simple network with one ca for tls and two ca's for orderer and peer (using external peer -[peer-external-builder](https://github.com/sachin-ngpws/peer-external-builder.git) )

commands for channel creation and join
--------------------------------------
* creating channel block
```
export CHANNEL_NAME="testchannel"
peer channel create -o orderer:7050 -c $CHANNEL_NAME --ordererTLSHostnameOverride orderer -f ./channel-artifacts/${CHANNEL_NAME}.tx --outputBlock ./channel-artifacts/${CHANNEL_NAME}.block --tls --cafile $CORE_PEER_TLS_ROOTCERT_FILE
```
* Joining the channel
```
peer channel join -b ./channel-artifacts/$CHANNEL_NAME.block
```

* Connection info
```
jq -n '{"address":"192.168.1.102:7054","dial_timeout": "10s","tls_required": false}' > connection.json
jq -n '{"path":"","type":"external","label":"mycc"}' > metadata.json
```

* Install chaincode
```
tar cfz code.tar.gz connection.json
tar cfz cpu.tgz metadata.json code.tar.gz
peer lifecycle chaincode install ./cpu.tgz
```
* Run Chaincode container
this is using the external chaincode -- [cpu-shim](https://github.com/sachin-ngpws/cpu-shim.git)
```
docker run -it --rm --name cpu-shim --hostname cpu-shim --env-file chaincode.env --network host ngp/cpu-shim
```

* Approving and commiting chaincode
without tls commands for more detail look at - [https://hyperledger-fabric.readthedocs.io/en/release-2.2/deploy_chaincode.html](https://hyperledger-fabric.readthedocs.io/en/release-2.2/deploy_chaincode.html)
```
peer lifecycle chaincode queryinstalled
export CC_PACKAGE_ID=<PACKAGE ID>
peer lifecycle chaincode approveformyorg -o orderer:7050 --channelID $CHANNEL_NAME --name cpu --version 1.0 --package-id $CC_PACKAGE_ID --sequence 1
peer lifecycle chaincode checkcommitreadiness --channelID $CHANNEL_NAME --name cpu --version 1.0 --sequence 1 --output json
peer lifecycle chaincode commit -o orderer:7050 --channelID $CHANNEL_NAME --name cpu --version 1.0 --sequence 1
peer lifecycle chaincode querycommitted --channelID $CHANNEL_NAME --name cpu
peer chaincode invoke -o orderer:7050 -C $CHANNEL_NAME -n cpu -c '{"function":"init","Args":[]}'
```