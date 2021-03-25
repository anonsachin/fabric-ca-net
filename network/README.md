Docker Network With Fabric CA's for tls and identities
======================================================

* To bring up the network run `make up`
* to bring down `sudo make clean`
* Look into the Makefile for more details, it's just an easier way to run all the scripts and commands.

commands for channel creation and join that are in the shell script
-------------------------------------------------------------------
* creating channel block
```
    export CHANNEL_NAME="testchannel"
    peer channel create -o orderer:7050 -c $CHANNEL_NAME --ordererTLSHostnameOverride orderer.testnetwork.com -f ./channel-artifacts/${CHANNEL_NAME}.tx --outputBlock ./channel-artifacts/${CHANNEL_NAME}.block --tls --cafile $CORE_PEER_TLS_ROOTCERT_FILE
```
The TLS CA is same for all participants hence you can use `$CORE_PEER_TLS_ROOTCERT_FILE` or `$ORDERER_CA_CERT` as they are the same cert. `ORDERER_CA_CERT` defined by me in the docker-compose.yml file


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
    export CC_PACKAGE_ID=<PACKAGE ID> (export CC_PACKAGE_ID=$(peer lifecycle chaincode queryinstalled | grep Package | cut -d \  -f 3 | cut -d , -f 1) => possible automation)
    peer lifecycle chaincode approveformyorg -o orderer:7050 --channelID $CHANNEL_NAME --name cpu --version 1.0 --package-id $CC_PACKAGE_ID --sequence 1
    peer lifecycle chaincode checkcommitreadiness --channelID $CHANNEL_NAME --name cpu --version 1.0 --sequence 1 --output json
    peer lifecycle chaincode commit -o orderer:7050 --channelID $CHANNEL_NAME --name cpu --version 1.0 --sequence 1
    peer lifecycle chaincode querycommitted --channelID $CHANNEL_NAME --name cpu
    peer chaincode invoke -o orderer:7050 -C $CHANNEL_NAME -n cpu -c '{"function":"init","Args":[]}'
```

* For tls add
```
    --tls --cafile $<root ca cert file for orderer>
    optional: -ordererTLSHostnameOverride <orderer name> <example => orderer.testnetwork.com>
```