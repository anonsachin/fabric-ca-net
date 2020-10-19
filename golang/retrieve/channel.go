package retrieve

import (
	"addorg/env"
	"os"
	"os/exec"
)

// ChannelConfig to get the config off the current channel
func ChannelConfig(configBlock string, c env.ENV) *exec.Cmd{
	// Getting configtxgen
	binary, lookErr := exec.LookPath("bash")
	if lookErr != nil {
		panic(lookErr)
	}

	//Setting the PATH
	c.SetEnv()
	// The command to run
	args := "peer channel fetch config "+configBlock+" -o orderer.testnetwork.com:7050 -c $CHANNEL_NAME --tls --cafile $ORDERER_CA_CERT"
	cmd := exec.Command(binary, "-c", args)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd
}

// func defaultOrgEnv(){
// 	os.Setenv("CORE_PEER_LOCALMSPID", "orgMSP")
// 	os.Setenv("CORE_PEER_MSPCONFIGPATH", "/etc/hyperledger/fabric/msp")
// 	os.Setenv("CORE_PEER_TLS_ROOTCERT_FILE", "/etc/hyperledger/fabric/tls/ca.crt")
// 	os.Setenv("CORE_PEER_ADDRESS","peer.testnetwork.com:7051")
// 	os.Setenv("ORDERER_CA_CERT","/etc/hyperledger/fabric/orderer-ca.cert")
// 	os.Setenv("CHANNEL_NAME","testchannel")
// }