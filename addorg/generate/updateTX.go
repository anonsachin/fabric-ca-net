package generate

import (
	"addorg/env"
	"os"
	"os/exec"
)

// ChannelUpdateTX to submit the update
func ChannelUpdateTX(c env.ENV) *exec.Cmd{
	// Getting configtxgen
	binary, lookErr := exec.LookPath("bash")
	if lookErr != nil {
		panic(lookErr)
	}

	//Setting the PATH
	c.SetEnv()
	// The command to run
	args := "peer channel update -f org3_update_in_envelope.pb -c $CHANNEL_NAME -o orderer.testnetwork.com:7050 --tls --cafile $ORDERER_CA_CERT"
	cmd := exec.Command(binary, "-c", args)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd
}