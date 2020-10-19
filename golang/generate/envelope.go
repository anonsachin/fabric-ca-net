package generate

import (
	"addorg/env"
	"os"
	"os/exec"
)


// NewChannelUpdate to create the update proto with links to orgs in channel
func Envelope(c env.ENV) *exec.Cmd{
	// Getting configtxgen
	binary, lookErr := exec.LookPath("bash")
	if lookErr != nil {
		panic(lookErr)
	}

	c.SetEnv()
	// The command to run
	args := "echo '{\"payload\":{\"header\":{\"channel_header\":{\"channel_id\":\"'$CHANNEL_NAME'\", \"type\":2}},\"data\":{\"config_update\":'$(cat org3_update.json)'}}}' | jq . > org3_update_in_envelope.json"
	cmd := exec.Command(binary, "-c", args)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd
}