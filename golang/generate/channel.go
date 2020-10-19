package generate

import (
	"os"
	"os/exec"
)


// NewChannelConfig Updating the channel config
func NewChannelConfig(newOrg string) *exec.Cmd{
	// Getting configtxgen
	binary, lookErr := exec.LookPath("bash")
	if lookErr != nil {
		panic(lookErr)
	}

	// The command to run
	args := "jq -s '.[0] * {\"channel_group\":{\"groups\":{\"Application\":{\"groups\": {\""+newOrg+"\":.[1]}}}}}' config.json "+newOrg+".json > modified_config.json"
	cmd := exec.Command(binary, "-c", args)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd
}