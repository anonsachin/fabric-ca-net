package generate

import (
	"os/exec"
	"os"
)

// OrgConfig is to create the org config from the configtx and the msp
func OrgConfig(org string) *exec.Cmd{
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

	return cmd

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