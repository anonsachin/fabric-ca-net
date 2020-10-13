package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func configTemplate(tempPath string, org string, path string) {
	file, err := ioutil.ReadFile(tempPath)
	if err != nil {
		panic("The file error ==> " + err.Error())
	}

	fileAsString := string(file)

	outFile := strings.ReplaceAll(fileAsString, "ORG", org)
	outFile = strings.ReplaceAll(outFile, "PATH", path)
	outBytes := []byte(outFile)
	err = ioutil.WriteFile("configtx.yaml", outBytes, 0644)

	if err != nil {
		_ = fmt.Errorf("Did not create file configtx.yaml \n")
	}
}

func generateOrgConfig(org string) {
	defer Recov()
	// Getting configtxgen
	binary, lookErr := exec.LookPath("bash")
	if lookErr != nil {
		panic(lookErr)
	}

	//Setting the PATH
	os.Setenv("FABRIC_CFG_PATH", os.Getenv("PWD"))

	env := os.Environ()

	// The command to run
	command := "configtxgen -printOrg " + org + " > " + org + ".json"
	fmt.Println("Exec ==> ",command)

	args := []string{"bash", "-c", command}

	execErr := syscall.Exec(binary, args, env)
	fmt.Println("Success exec ==> ",command)
	if execErr != nil {
		panic(execErr)
	}
}

func configConsulTemplate(tempPath string, vault string, path string, org string, role string) {
	file, err := ioutil.ReadFile(tempPath)
	if err != nil {
		panic("The file error ==> " + err.Error())
	}

	fileAsString := string(file)

	outFile := strings.ReplaceAll(fileAsString, "ORG", org)
	outFile = strings.ReplaceAll(outFile, "BASEPATH", path)
	outFile = strings.ReplaceAll(outFile, "ROLE", role)
	outFile = strings.ReplaceAll(outFile, "VAULTHOST", vault)

	outBytes := []byte(outFile)

	dest := "consul-template-" + role + ".hcl"
	err = ioutil.WriteFile(dest, outBytes, 0644)

	if err != nil {
		_ = fmt.Errorf("Did not create file configtx.yaml \n")
	}
}


func Recov(){
	if r := recover(); r!= nil {
        fmt.Println("recovered from ", r)
    }
}