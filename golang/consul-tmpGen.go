
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
	a = strings.Replace(a, "CNAME", role+".service.consul", 2)
	a = strings.Replace(a, "TTL", "24h", 1)
	a = strings.Replace(a, "CERT", cert, 1)
	return a
}

//ConsulTempGen Generate template
func ConsulTempGen(tempFile string, tlsTempFile string, outDir string, role string, org string) {
	//Reading the Template
	fmt.Printf("The template file from %s \n", tempFile)
	tempBytes, err := ioutil.ReadFile(tempFile)
	if err != nil {
		panic("The file error ==> " + err.Error())
	}
	// Conserting to string
	temp := string(tempBytes)
	//Reading the TLS Template
	fmt.Printf("The template file from %s \n", tlsTempFile)
	tlsTempBytes, err := ioutil.ReadFile(tlsTempFile)
	if err != nil {
		panic("The file error ==> " + err.Error())
	}
	// Conserting to string
	tlsTemp := string(tlsTempBytes)
	//Showing the template
	fmt.Printf("The MSP Template ==> \n %s \n", temp)
	fmt.Printf("The TLS Template ==> \n %s \n", tlsTemp)
	//The Requirements
	certs := []string{"issuing_ca", "private_key", "certificate"}
	destinations := []string{"msp/cacerts", "msp/keystore", "msp/signcerts"}
	tlsDestinations := "tls"
	destFile := []string{"ca.cert.tpl", "agent.key.tpl", "agent.crt.tpl"}
	var path string
	//Creating the structure
	// TLS folder
	path = filepath.Join(outDir, role, tlsDestinations)
	os.MkdirAll(path, os.ModePerm)
	//MSP structure and templates
	for i := range certs {
		//Creating the Dirctory
		path = filepath.Join(outDir, role, destinations[i])
		os.MkdirAll(path, os.ModePerm)
		//The Output
		fmt.Printf("The MSP output ==> \n %v:%v \n", path, TemplateGen(temp, certs[i], role, org))
		//Writing the output to MSP
		tempBytes := []byte(TemplateGen(temp, certs[i], role, org))
		path = filepath.Join(path, destFile[i])
		err := ioutil.WriteFile(path, tempBytes, 0644)
		if err != nil {
			_ = fmt.Errorf("Did not file %s", path)
		}
		//The Output
		fmt.Printf("The TLS output ==> \n %v:%v \n", path, TemplateGen(tlsTemp, certs[i], role, org))
		//Writing the output to TLS
		tempBytes = []byte(TemplateGen(tlsTemp, certs[i], role, org))
		path = filepath.Join(outDir, role, tlsDestinations)
		path = filepath.Join(path, destFile[i])
		err = ioutil.WriteFile(path, tempBytes, 0644)
		if err != nil {
			_ = fmt.Errorf("Did not create file at %s", path)
		}
	}
}