package convert

import (
	"os"
	"os/exec"
)

// ChannelConfig convert channel config block to json and extract needed values
func ChannelConfig(configBlock string) *exec.Cmd{
	// Getting configtxgen
	binary, lookErr := exec.LookPath("bash")
	if lookErr != nil {
		panic(lookErr)
	}

	// The command to run
	args := "configtxlator proto_decode --input "+configBlock+" --type common.Block | jq .data.data[0].payload.data.config > config.json"
	cmd := exec.Command(binary, "-c", args)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd
}