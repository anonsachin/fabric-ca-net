'use strict';

const fs = require('fs');
const {Wallets,X509Identity} = require('fabric-network');

async function main(org){
    try {
        const filePath = process.env.WALLET_PATH;
        // checking for wallet path
        errorHandler(filePath,"File path not defined");
        const certPath = process.env.CERT_PATH;
        // checking cert path
        errorHandler(certPath,"Certificate path not defined");
        const keyPath = process.env.KEY_PATH;
        // checking key path
        errorHandler(keyPath,"Key path not defined");

        //Getting cert
        const cert = fs.readFileSync(certPath).toString();

        //Getting key
        const key = fs.readFileSync(keyPath).toString();

        //Creating wallet
        const wallet = await Wallets.newFileSystemWallet(filePath);
        console.log("Plain wallet created!!")

        //Creating identity

        const identity = {
            credentials:{
                certificate: cert,
                privateKey: key,
            },
            mspId: org+"MSP",
            type: 'X.509',
        };
        // storing the identity
        await wallet.put(org,identity)

        console.log("Identity Created!!")

    } catch (error) {
        console.error(error)
        throw new Error(error)
    }
}

function errorHandler(value,message){
    if (value == undefined ){
        throw new Error(message);
    }
}

main("org").then(()=> console.log("Done!")).catch(e => console.error(e));