package convert

import (
	"os"
	"os/exec"
)

// OldConfig convert old config json to proto
func OldConfig() *exec.Cmd{
	// Getting configtxgen
	binary, lookErr := exec.LookPath("bash")
	if lookErr != nil {
		panic(lookErr)
	}

	// The command to run
	args := "configtxlator proto_encode --input config.json --type common.Config --output config.pb"
	cmd := exec.Command(binary, "-c", args)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd
}