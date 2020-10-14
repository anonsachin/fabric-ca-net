package template

import (
	"os"
	"path/filepath"
	"io/ioutil"
	"fmt"
	"strings"
)

//ConsulTemp consul-template
type ConsulTemp struct{
	tempFile string
	tlsTempFile string
	outDir string
	role string
	org string
	tempPath string
	vault string
	basePath string
}

// NewConsul constructor
func NewConsul(tempFile string, 
	tlsTempFile string,
	outDir string,
	role string,
	org string,
	tempPath string,
	vault string,
	basePath string) *ConsulTemp{
		return &ConsulTemp{
			tempFile: tempFile,
			tlsTempFile: tlsTempFile,
			outDir: outDir,
			role: role,
			org: org,
			tempPath: tempPath,
			vault: vault,
			basePath: basePath,
		}
}

//ConsulTempGen Generate consul-templates for vault pki certs
func (c *ConsulTemp) ConsulTempGen( ) {
	//Reading the Template
	fmt.Printf("The template file from %s \n", c.tempFile)
	tempBytes, err := ioutil.ReadFile(c.tempFile)
	if err != nil {
		panic("The file error ==> " + err.Error())
	}
	// Conserting to string
	temp := string(tempBytes)
	//Reading the TLS Template
	fmt.Printf("The template file from %s \n", c.tlsTempFile)
	tlsTempBytes, err := ioutil.ReadFile(c.tlsTempFile)
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
	path = filepath.Join(c.outDir, c.role, tlsDestinations)
	os.MkdirAll(path, os.ModePerm)
	//MSP structure and templates
	for i := range certs {
		//Creating the Dirctory
		path = filepath.Join(c.outDir, c.role, destinations[i])
		os.MkdirAll(path, os.ModePerm)
		//The Output
		fmt.Printf("The MSP output ==> \n %v:%v \n", path, templateGen(temp, certs[i], c.role, c.org))
		//Writing the output to MSP
		tempBytes := []byte(templateGen(temp, certs[i], c.role, c.org))
		path = filepath.Join(path, destFile[i])
		err := ioutil.WriteFile(path, tempBytes, 0644)
		if err != nil {
			_ = fmt.Errorf("Did not file %s", path)
		}
		//The Output
		fmt.Printf("The TLS output ==> \n %v:%v \n", path, templateGen(tlsTemp, certs[i], c.role, c.org))
		//Writing the output to TLS
		tempBytes = []byte(templateGen(tlsTemp, certs[i], c.role, c.org))
		path = filepath.Join(c.outDir, c.role, tlsDestinations)
		path = filepath.Join(path, destFile[i])
		err = ioutil.WriteFile(path, tempBytes, 0644)
		if err != nil {
			_ = fmt.Errorf("Did not create file at %s", path)
		}
	}
}

//ConfigConsulTemplate genreting Consul template for a role
func (c *ConsulTemp) ConfigConsulTemplate() {
	file, err := ioutil.ReadFile(c.tempPath)
	if err != nil {
		panic("The file error ==> " + err.Error())
	}

	fileAsString := string(file)

	outFile := strings.ReplaceAll(fileAsString, "ORG", c.org)
	outFile = strings.ReplaceAll(outFile, "BASEPATH", c.basePath)
	outFile = strings.ReplaceAll(outFile, "ROLE", c.role)
	outFile = strings.ReplaceAll(outFile, "VAULTHOST", c.vault)

	outBytes := []byte(outFile)

	dest := "consul-template-" + c.role + ".hcl"
	err = ioutil.WriteFile(dest, outBytes, 0644)

	if err != nil {
		_ = fmt.Errorf("Did not create file configtx.yaml \n")
	}
}

func templateGen(a string, cert string, role string, org string) string {
	a = strings.Replace(a, "ORG", org, 1)
	a = strings.Replace(a, "ROLE", role, 1)
	a = strings.Replace(a, "CNAME", role+".service.consul", 2)
	a = strings.Replace(a, "TTL", "24h", 1)
	a = strings.Replace(a, "CERT", cert, 1)
	return a
}