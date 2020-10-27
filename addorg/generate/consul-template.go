package generate

import (
	"os"
	"os/exec"
)

// Certs invokes consul template to create the certs
func Certs(consulPath string) *exec.Cmd{
	binary, lookErr := exec.LookPath("bash")
	if lookErr != nil {
		panic(lookErr)
	}
	// The command to run
	args := "consul-template -config " + consulPath + " -once"
	cmd := exec.Command(binary, "-c", args)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd
}