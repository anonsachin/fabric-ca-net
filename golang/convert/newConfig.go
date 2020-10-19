package convert

import (
	"os"
	"os/exec"
)

// NewConfig convert new config json to proto
func NewConfig() *exec.Cmd{
	// Getting configtxgen
	binary, lookErr := exec.LookPath("bash")
	if lookErr != nil {
		panic(lookErr)
	}

	// The command to run
	args := "configtxlator proto_encode --input modified_config.json --type common.Config --output modified_config.pb"
	cmd := exec.Command(binary, "-c", args)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd
}