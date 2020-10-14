package main

import (
	"os"
	"os/exec"
)



// GenCerts to genrate the certs
func GenCerts(consulPath string) {
	defer Recov()
	// Getting configtxgen
	binary, lookErr := exec.LookPath("bash")
	if lookErr != nil {
		panic(lookErr)
	}
	// The command to run
	args := "consul-template -config " + consulPath + " -once"
	cmd := exec.Command(binary, "-c", args)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run()

	if err != nil {
		panic(err)
	}
}
