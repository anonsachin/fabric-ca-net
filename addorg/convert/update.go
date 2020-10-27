package convert

import (
	"os"
	"os/exec"
)

// UpdateConfig convert update proto to json
func UpdateConfig() *exec.Cmd{
	// Getting configtxgen
	binary, lookErr := exec.LookPath("bash")
	if lookErr != nil {
		panic(lookErr)
	}

	// The command to run
	args := "configtxlator proto_decode --input org3_update.pb --type common.ConfigUpdate | jq . > org3_update.json"
	cmd := exec.Command(binary, "-c", args)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd
}

