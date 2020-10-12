package main

import (
	"fmt"
	"strings"
	"io/ioutil"
)

func configTemplate(tempPath string, org string){
	file, err := ioutil.ReadFile(tempPath)
	if err != nil {
		panic("The file error ==> " + err.Error())
	}

	fileAsString := string(file)

	outFile := strings.ReplaceAll(fileAsString,"ORG",org)
	outBytes := []byte(outFile)
	err = ioutil.WriteFile("configtx.yaml",outBytes,0644)

	if err != nil {
		_ = fmt.Errorf("Did not create file configtx.yaml \n")
	}
}