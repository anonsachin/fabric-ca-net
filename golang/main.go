package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
	//Setting up the flags
	tmpFile := flag.String("template", "template.tpl", "Used to get the template file")
	outDir := flag.String("out_dir", "./NewOrg/", "Used to get the directory for certs of the new org")
	newOrg := flag.String("new_org", "NewOrg", "Used to specify the new org name")
	role := flag.String("role", "peer", "Used to specify the role of the certs")
	//Getting there values
	flag.Parse()
	//Reading the Template
	fmt.Printf("The template file from %s \n", *tmpFile)
	file, err := ioutil.ReadFile(*tmpFile)
	if err != nil {
		panic("The file error ==> " + err.Error())
	}
	//Genrating the folder structure and templates
	ConsulTempGen(string(file), *outDir, *role, *newOrg)
}
