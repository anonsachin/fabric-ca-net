package main

import (
	"fmt"
	"os"
	"os/exec"
)



func generateOrgConfig(org string) {
	defer Recov()
	// Getting configtxgen
	binary, lookErr := exec.LookPath("bash")
	if lookErr != nil {
		panic(lookErr)
	}

	//Setting the PATH
	os.Setenv("FABRIC_CFG_PATH", os.Getenv("PWD"))

	// The command to run
	args := "configtxgen -printOrg " + org + " > " + org + ".json"
	cmd := exec.Command(binary, "-c", args)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run()

	if err != nil {
		panic(err)
	}

	// env := os.Environ()

	// // The command to run
	// command := "configtxgen -printOrg " + org + " > " + org + ".json"
	// fmt.Println("Exec ==> ",command)

	// args := []string{"bash", "-c", command}

	// execErr := syscall.Exec(binary, args, env)
	// fmt.Println("Success exec ==> ",command)
	// if execErr != nil {
	// 	panic(execErr)
	// }
}

func Recov() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
	}
}
