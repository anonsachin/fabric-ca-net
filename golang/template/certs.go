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
	tempFile string //Path to template file for MSP certs
	tlsTempFile string //Path to template file for TLS certs
	outDir string
	role string
	org string
	tempPath string //Path to consul-template config file
	vaultHost string //Vault url like http://127.0.0.1:8200/
	basePath string //Base path for the consul-template file for the full path to the template file like /home/sachin/ca-net/golang/NewOrg/admin/msp/cacerts/ca.cert.tpl
	baseDomain string //Base domain for the network
}

// NewConsul constructor
func NewConsul(tempFile string, 
	tlsTempFile string,
	outDir string,
	role string,
	org string,
	tempPath string,
	vaultHost string,
	basePath string,
	baseDomain string) *ConsulTemp{
		return &ConsulTemp{
			tempFile: tempFile,
			tlsTempFile: tlsTempFile,
			outDir: outDir,
			role: role,
			org: org,
			tempPath: tempPath,
			vaultHost: vaultHost,
			basePath: basePath,
			baseDomain: baseDomain,
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
		fmt.Printf("The MSP output ==> \n %v:%v \n", path, templateGen(temp, certs[i], c.role, c.org, c.baseDomain))
		//Writing the output to MSP
		tempBytes := []byte(templateGen(temp, certs[i], c.role, c.org, c.baseDomain))
		path = filepath.Join(path, destFile[i])
		err := ioutil.WriteFile(path, tempBytes, 0644)
		if err != nil {
			_ = fmt.Errorf("Did not file %s", path)
		}
		//The Output
		fmt.Printf("The TLS output ==> \n %v:%v \n", path, templateGen(tlsTemp, certs[i], c.role, c.org, c.baseDomain))
		//Writing the output to TLS
		tempBytes = []byte(templateGen(tlsTemp, certs[i], c.role, c.org, c.baseDomain))
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
	outFile = strings.ReplaceAll(outFile, "VAULTHOST", c.vaultHost)

	outBytes := []byte(outFile)

	dest := "consul-template-" + c.role + ".hcl"
	err = ioutil.WriteFile(dest, outBytes, 0644)

	if err != nil {
		_ = fmt.Errorf("Did not create file configtx.yaml \n")
	}
}

func templateGen(a string, cert string, role string, org string, baseDomain string) string {
	a = strings.Replace(a, "ORG", org, 1)
	a = strings.Replace(a, "ROLE", role, 1)
	var cname string
	if role == "admin"{
		cname = role
	} else {
		cname = role+"."+baseDomain
	}
	a = strings.Replace(a, "CNAME", cname, 2)
	a = strings.Replace(a, "TTL", "24h", 1)
	a = strings.Replace(a, "CERT", cert, 1)
	return a
}