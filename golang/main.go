package main

import (
	"fmt"
	"addorg/template"
	"addorg/generate"
)

func main() {
	//Setting up the flags
	// Base flags
	// outDir, newOrg, vaultHost :=  baseFlags()

	// // msp flags
	// msp := mspFlags()

	// // certs flags
	// tmpFile, tlsTmpFile, baseDomain, role, consulTemp, basePath, consulTempOutPath := certsFlags()

	// // OrgConfig flags
	// configtx, configtxFile := orgConfigFlags()
	
	certs, consulTempOutPath, basePath, consulTemp, tmpFile, tlsTmpFile, outDir, newOrg, vaultHost, role, configtxFile, msp, configtx, baseDomain := getFlags()

	// Getting New consul-template)
	ct := template.NewConsul(*tmpFile, *tlsTmpFile, *outDir, *role, *newOrg, *consulTemp, *vaultHost, *basePath, *baseDomain)
	// Getting New configtx.yaml
	conf  := template.NewConfigTX(*configtxFile,*newOrg,*outDir)

	// Getting New MSP
	mspNew := generate.NewMSP(*vaultHost, *newOrg, *outDir)

	if *certs {
		// //Genrating the folder structure and templates
	ct.ConsulTempGen()

	// //  Genrating consul template
	// configConsulTemplate(*consulTemp, *vaultHost, *basePath, *newOrg, *role)
	ct.ConfigConsulTemplate()

	fmt.Println("Generating template ==> ", *consulTempOutPath)
		// Generate Certs from Template
		err:= execute(generate.Certs(*consulTempOutPath))
		if err != nil {
			panic(err)
		}
	}

	// Generate The MSP
	if *msp {
		mspNew.CreateMSP()
	}

	// Running the consul-template seperatly ass it stops execution after completion
	// sep := make(chan string)

	// go func() {
	// 	fmt.Println("Generating template ==> ", *consulTempOutPath)
	// 	// Generate Certs from Template
	// 	err:= execute(generate.Certs(*consulTempOutPath))
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	sep <- "done"
	// }()

	// Generate the config JSON
	fmt.Println("the Config status", *configtx)

	if *configtx {
		fmt.Println("Generating the Configs")
		// configTemplate(*configtxFile, *newOrg, *outDir)
		conf.ConfigTXTemplate()
		err:= execute(generate.OrgConfig(*newOrg))
		if err != nil {
			panic(err)
		}
		fmt.Println("Completed generating the Configs ")
	}

	// status := <-sep

	// fmt.Printf("The completion Status of consul-template %v \n", status)
}