package main

import (
	"fmt"
	"io/ioutil"
	"flag"
)

func main() {

	tmpFile := flag.String("template","template.tpl","Used to get the template file")
	outDir := flag.String("out_dir","./NewOrg/","Used to get the directory for certs of the new org")
	role := flag.String("role","peer","Used to specify the role of the certs")
	flag.Parse()
	fmt.Printf("The template file from %s \n",*tmpFile)
	file, err := ioutil.ReadFile(*tmpFile)
	if err != nil {
		panic("The file error ==> "+err.Error())
	}
	ConsulTempGen(string(file),*outDir,*role)
}