'use strict';

const gateway = require("./gateway.js")

async function main(org,channelName,contractName,cpuName,value1,value2){
    try {

        //Getting the contract
        let contract = await gateway.getContract(org,channelName,contractName);

        //Submitting the TX
        console.log("Submitting the TX");
        let buffer = await contract.submitTransaction('AddUsage',cpuName,value1,value2);

        //Converting to json
        let asset =JSON.parse(buffer.toString());
        gateway.errorHandler(asset,"Empty asset")

        console.log(asset);
        
        return asset;
        
    } catch (error) {
        console.error(error)
        throw new Error(error);
    }finally{
        await gateway.disconnect();
    }
}

module.exports.run = main;