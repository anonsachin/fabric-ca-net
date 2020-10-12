package main

import (
	"syscall"
	"os/exec"
	"os"
	"fmt"
	"strings"
	"io/ioutil"
)

func configTemplate(tempPath string, org string, path string){
	file, err := ioutil.ReadFile(tempPath)
	if err != nil {
		panic("The file error ==> " + err.Error())
	}

	fileAsString := string(file)

	outFile := strings.ReplaceAll(fileAsString,"ORG",org)
	outFile = strings.ReplaceAll(outFile,"PATH",path)
	outBytes := []byte(outFile)
	err = ioutil.WriteFile("configtx.yaml",outBytes,0644)

	if err != nil {
		_ = fmt.Errorf("Did not create file configtx.yaml \n")
	}
}

func generateOrgConfig(org string){
	// Getting configtxgen
	binary, lookErr := exec.LookPath("bash")
    if lookErr != nil {
        panic(lookErr)
	}

	env := os.Environ()

	args := []string{"bash","-c","./addOrg.sh "+org}
	fmt.Printf("The command is %s",args)
	execErr := syscall.Exec(binary, args, env)
    if execErr != nil {
        panic(execErr)
    }
}