package main

import (
	"strings"
	"fmt"
)

//TemplateGen edits the consul template
func TemplateGen(a string,cert string) string{
	a = strings.Replace(a,"ORG","peer",1)
	a = strings.Replace(a,"ROLE","peer",1)
	a = strings.Replace(a,"CNAME","peer.service.consul",2)
	a = strings.Replace(a,"TTL","24h",1)
	a = strings.Replace(a,"CERT",cert,1)
	return a
}
//ConsulTempGen Generate template
func ConsulTempGen() {
	a := "{{ with secret \"ORGCA/issue/ROLE\" \"common_name=CNAME\" \"ttl=TTL\" \"alt_names=localhost,CNAME\" \"ip_sans=127.0.0.1\"}}\n{{ .Data.CERT }}\n{{ end }}"
	fmt.Printf("The BASE Versin ==> %s \n",a)
	certs := []string{"issuing_ca","private_key","certificate"}
	for _,cert := range certs {
		fmt.Printf("The NEW Versin ==> %v \n",TemplateGen(a,cert))
	}
}