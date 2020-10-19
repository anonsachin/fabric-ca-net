package main

import (
	"addorg/convert"
	"addorg/env"
	"addorg/flags"
	"addorg/generate"
	"addorg/retrieve"
	"addorg/template"
	"fmt"
)

func main() {
	//Setting up the flags
	f := flags.GetFlags()
	// Base flags
	outDir, newOrg, vaultHost := f.BaseFlags()

	// msp flags
	msp := f.MSPFlags()

	// certs flags
	tmpFile, tlsTmpFile, baseDomain, role, consulTemp, basePath, consulTempOutPath, certs := f.CertsFlags()

	// OrgConfig flags
	configtx, configtxFile := f.OrgConfigFlags()

	// Channel flgs
	channel, configBlock := f.ChannelFlags()

	if *certs {
		// Getting New consul-template)
		ct := template.NewConsul(*tmpFile, *tlsTmpFile, *outDir, *role, *newOrg, *consulTemp, *vaultHost, *basePath, *baseDomain)
		//Genrating the folder structure and templates
		ct.ConsulTempGen()

		//  Genrating consul template
		ct.ConfigConsulTemplate()

		fmt.Println("Generating template ==> ", *consulTempOutPath)
		// Generate Certs from Template
		err := execute(generate.Certs(*consulTempOutPath))
		if err != nil {
			panic(err)
		}
	}

	// Generate The MSP
	if *msp {
		// Getting New MSP
		mspNew := generate.NewMSP(*vaultHost, *newOrg, *outDir)

		mspNew.CreateMSP()
	}

	// Generate the config JSON
	fmt.Println("the Config status", *configtx)

	if *configtx {
		// Getting New configtx.yaml
		conf := template.NewConfigTX(*configtxFile, *newOrg, *outDir)

		fmt.Println("Generating the Configs")
		// configTemplate(*configtxFile, *newOrg, *outDir)
		conf.ConfigTXTemplate()
		err := execute(generate.OrgConfig(*newOrg))
		if err != nil {
			panic(err)
		}
		fmt.Println("Completed generating the Configs ")
	}

	// Operate on Channel
	if *channel {
		fmt.Println("Setting the peer context")
		context := env.DefaultOrgEnv()
		fmt.Println("Getting the Channel Configs ==> ", *configBlock)
		err := execute(retrieve.ChannelConfig(*configBlock, context))
		if err != nil {
			panic(err)
		}
		fmt.Println("Converting the Channel Configs block to json and extracting required values")
		err = execute(convert.ChannelConfig(*configBlock))
		if err != nil {
			panic(err)
		}
		fmt.Println("Creating the new channel configs")
		err = execute(generate.NewChannelConfig(*newOrg))
		if err != nil {
			panic(err)
		}
		fmt.Println("Converting old config json to proto")
		err = execute(convert.OldConfig())
		if err != nil {
			panic(err)
		}
		fmt.Println("Converting new config json to proto")
		err = execute(convert.NewConfig())
		if err != nil {
			panic(err)
		}
	}
}
