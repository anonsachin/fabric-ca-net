package main

import (
	"strings"
	"fmt"
	"path/filepath"
	"os"
)

//TemplateGen edits the consul template
func TemplateGen(a string,cert string, role string) string{
	a = strings.Replace(a,"ORG","peer",1)
	a = strings.Replace(a,"ROLE",role,1)
	a = strings.Replace(a,"CNAME","peer.service.consul",2)
	a = strings.Replace(a,"TTL","24h",1)
	a = strings.Replace(a,"CERT",cert,1)
	return a
}
//ConsulTempGen Generate template
func ConsulTempGen(a string,outDir string, role string) {
	// a := "{{ with secret \"ORGCA/issue/ROLE\" \"common_name=CNAME\" \"ttl=TTL\" \"alt_names=localhost,CNAME\" \"ip_sans=127.0.0.1\"}}\n{{ .Data.CERT }}\n{{ end }}"
	fmt.Printf("The BASE Versin ==> \n %s \n",a)
	certs := []string{"issuing_ca","private_key","certificate"}
	destinations := []string{"msp/cacerts","msp/keystore","msp/signcerts"}
	var path string
	for i := range certs {
		path = filepath.Join(outDir,role,destinations[i])
		os.MkdirAll(path,os.ModePerm)
		fmt.Printf("The NEW Versin ==> \n %v:%v \n",path,TemplateGen(a,certs[i],role))
	}
}