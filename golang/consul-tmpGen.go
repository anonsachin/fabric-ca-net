package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

//TemplateGen edits the consul template
func TemplateGen(a string, cert string, role string, org string) string {
	a = strings.Replace(a, "ORG", org, 1)
	a = strings.Replace(a, "ROLE", role, 1)
	a = strings.Replace(a, "CNAME", "peer.service.consul", 2)
	a = strings.Replace(a, "TTL", "24h", 1)
	a = strings.Replace(a, "CERT", cert, 1)
	return a
}

//ConsulTempGen Generate template
func ConsulTempGen(a string, outDir string, role string, org string) {
	//Showing the template
	fmt.Printf("The BASE Versin ==> \n %s \n", a)
	//The Requirements
	certs := []string{"issuing_ca", "private_key", "certificate"}
	destinations := []string{"msp/cacerts", "msp/keystore", "msp/signcerts"}
	destFile := []string{"ca.cert.tpl", "agent.key.tpl", "agent.crt.tpl"}
	var path string
	//Creating the structure
	for i := range certs {
		//Creating the Dirctory
		path = filepath.Join(outDir, role, destinations[i])
		os.MkdirAll(path, os.ModePerm)
		//The Output
		fmt.Printf("The NEW Versin ==> \n %v:%v \n", path, TemplateGen(a, certs[i], role, org))
		//Writing the output
		tempBytes := []byte(TemplateGen(a, certs[i], role, org))
		path = filepath.Join(path, destFile[i])
		err := ioutil.WriteFile(path, tempBytes, 0644)
		if err != nil {
			_ = fmt.Errorf("Did not file %s", path)
		}
	}
}
