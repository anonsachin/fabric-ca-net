package generate

import (
	"fmt"
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

	// The command to run
	args := fmt.Sprintf("configtxgen -printOrg %s > %s/%s.json", org,org,org) //"configtxgen -printOrg " + org + " > " + org + ".json"
	cmd := exec.Command(binary, "-c", args)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	//Setting the PATH
	path := fmt.Sprintf("PATH=%s",os.Getenv("PATH"))
	fabricPath := fmt.Sprintf("FABRIC_CFG_PATH=%s",os.Getenv("PWD"))

	cmd.Env = []string{path,fabricPath}

	return cmd

}