'use strict';

const gateway = require("./gateway.js")

async function main(org,channelName,contractName,cpuName){
    try {

        //Getting the contract
        let contract = await gateway.getContract(org,channelName,contractName);

        //Submitting the TX
        console.log("Submitting the TX");
        let buffer = await contract.submitTransaction('AddCpu',cpuName);

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

main('org','testchannel','cpu','cpu-123').then(()=> console.log("TX submitted!!")).catch(e => console.error(e));