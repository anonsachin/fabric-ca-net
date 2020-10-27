package generate

import (
	"addorg/env"
	"os"
	"os/exec"
)


// NewChannelUpdate to create the update proto with links to orgs in channel
func NewChannelUpdate(c env.ENV) *exec.Cmd{
	// Getting configtxgen
	binary, lookErr := exec.LookPath("bash")
	if lookErr != nil {
		panic(lookErr)
	}

	c.SetEnv()
	// The command to run
	args := "configtxlator compute_update --channel_id $CHANNEL_NAME --original config.pb --updated modified_config.pb --output org3_update.pb"
	cmd := exec.Command(binary, "-c", args)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd
}