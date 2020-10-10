package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	//Setting up the flags
	tmpFile := flag.String("template", "template.tpl", "Used to get the template file")
	tlsTmpFile := flag.String("tls_template", "tlstemplate.tpl", "Used to get the template file")
	outDir := flag.String("out_dir", "./NewOrg/", "Used to get the directory for certs of the new org")
	newOrg := flag.String("new_org", "NewOrg", "Used to specify the new org name")
	vaultHost := flag.String("vault_host", "http://127.0.0.1:8200", "Used to specify the vautl http ednpoint, Default = http://127.0.0.1:8200")
	role := flag.String("role", "peer", "Used to specify the role of the certs")
	//Getting there values
	flag.Parse()
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
