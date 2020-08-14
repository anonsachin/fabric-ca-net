'use strict';

const fs = require('fs');
const {Wallets,Gateway} = require('fabric-network');
let gateway;

async function getChannel(org,channelName){
    try {
        const filePath = process.env.WALLET_PATH;
        //Getting wallet path
        errorHandler(filePath,"File path not defined");

        //Getting wallet
        console.log("Getting wallet")
        const wallet = await Wallets.newFileSystemWallet(filePath);
        //Getting user
        const identity = await wallet.get(org)
        errorHandler(identity,"IDENTITY not present")

        //Getting ccp
        console.log("Getting ccp")
        const ccpPath = process.env.CCP_PATH
        errorHandler(ccpPath,"CCP path not defined")
        let ccp = JSON.parse(fs.readFileSync(ccpPath, 'utf8'));
        errorHandler(ccp,"CCP is not defined")

        //Getting gateway
        console.log("Getting gateway")
        gateway = new Gateway();
        await gateway.connect(ccp, { wallet, identity: org, discovery: { enabled: false, asLocalhost: true } });

        //Getting the channel
        console.log("Getting channel")
        const network = await gateway.getNetwork(channelName);

        return network;

    } catch (error) {
        console.error(error);
        return error;
    }
}

async function getContract(org,channelName,contractName){
    try {
        const filePath = process.env.WALLET_PATH;
        //Getting wallet path
        errorHandler(filePath,"File path not defined");

        //Getting wallet
        console.log("Getting wallet")
        const wallet = await Wallets.newFileSystemWallet(filePath);
        //Getting user
        const identity = await wallet.get(org)
        errorHandler(identity,"IDENTITY not present")
        console.log(identity)

        //Getting ccp
        console.log("Getting ccp")
        const ccpPath = process.env.CCP_PATH
        errorHandler(ccpPath,"CCP path not defined")
        let ccp = JSON.parse(fs.readFileSync(ccpPath, 'utf8'));
        errorHandler(ccp,"CCP is not defined")

        //Getting gateway
        console.log("Getting gateway")
        gateway = new Gateway();
        let connectionOption = {
            identity: identity,
            discovery: { enabled: false, asLocalhost: true }
        }
        console.log("Connecting...")
        await gateway.connect(ccp,connectionOption);

        //Getting the channel
        console.log("Getting channel")
        const network = await gateway.getNetwork(channelName);

        //Getting the contract
        console.log("Getting contract")
        const contract = network.getContract(contractName);

        return contract;

    } catch (error) {
        console.error(error);
        return error;
    }
}

function errorHandler(value,message){
    if (value == undefined ){
        throw new Error(message);
    }
}

async function disconnect(){
    console.log("Disconnecting...");
    await gateway.disconnect();
}

module.exports.getContract = getContract;
module.exports.getChannel = getChannel;
module.exports.disconnect = disconnect;
module.exports.errorHandler = errorHandler;