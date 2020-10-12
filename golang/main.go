package main

import (
	"syscall"
	"os/exec"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	//Setting up the flags
	tmpFile, tlsTmpFile, outDir, newOrg, vaultHost, role, configtxFile, msp, configtxReq := getFlags()

	//Reading the Template
	fmt.Printf("The template file from %s \n", *tmpFile)
	file, err := ioutil.ReadFile(*tmpFile)
	if err != nil {
		panic("The file error ==> " + err.Error())
	}
	//Reading the TLS Template
	fmt.Printf("The template file from %s \n", *tmpFile)
	tlsFile, err := ioutil.ReadFile(*tlsTmpFile)
	if err != nil {
		panic("The file error ==> " + err.Error())
	}
	//Genrating the folder structure and templates
	ConsulTempGen(string(file), string(tlsFile), *outDir, *role, *newOrg)

	if *msp {
		//Getting the MSP's
		mspCa, tlsCa := getCaCert(*vaultHost, *newOrg)
		//Creating the Dirctory of MSP CA
		path := filepath.Join(*outDir, "msp/cacerts")
		os.MkdirAll(path, os.ModePerm)
		//Writing to file
		path = filepath.Join(path, "ca.pem")
		err = ioutil.WriteFile(path, mspCa, 0644)
		if err != nil {
			_ = fmt.Errorf("Did not file %s", path)
		}
		//Creating the Dirctory of TLS CA
		path = filepath.Join(*outDir, "msp/tlscacerts")
		os.MkdirAll(path, os.ModePerm)
		//Writing to file
		path = filepath.Join(path, "ca.pem")
		err = ioutil.WriteFile(path, tlsCa, 0644)
		if err != nil {
			_ = fmt.Errorf("Did not file %s", path)
		}
	}

	if *configtxReq {
		configTemplate(*configtxFile,*newOrg)
	}

	dirStatus(*outDir)

}

func dirStatus(outDir string){
	binary, lookErr := exec.LookPath("tree")
    if lookErr != nil {
        panic(lookErr)
	}
	
	args := []string{"tree", outDir}

	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
    if execErr != nil {
        panic(execErr)
    }
}

func getFlags() (*string, *string, *string, *string, *string, *string, *string, *bool, *bool){
	//Setting up the flags
	tmpFile := flag.String("template", "templates/template.tpl", "Used to get the template file")
	tlsTmpFile := flag.String("tls_template", "templates/tlstemplate.tpl", "Used to get the template file")
	outDir := flag.String("out_dir", "./NewOrg/", "Used to get the directory for certs of the new org")
	newOrg := flag.String("new_org", "NewOrg", "Used to specify the new org name")
	vaultHost := flag.String("vault_host", "http://127.0.0.1:8200", "Used to specify the vautl http ednpoint, Default = http://127.0.0.1:8200")
	role := flag.String("role", "peer", "Used to specify the role of the certs")
	configtxFile := flag.String("configtx","templates/configtx-templat.yaml","To Generate the config tx for new org")
	msp := flag.Bool("msp",true,"To generate msp")
	configtxReq := flag.Bool("configtx_req", false, "To Create the configtx.yaml")
	//Getting there values
	flag.Parse()

	return tmpFile, tlsTmpFile, outDir, newOrg, vaultHost, role, configtxFile, msp, configtxReq
}