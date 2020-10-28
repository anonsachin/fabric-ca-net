package template

import (
	"fmt"
	"io/ioutil"
	"path"
	"strings"
)

//ConfigTX configtx.yaml
type ConfigTX struct {
	tempPath string // Path to configtx.yaml template
	org      string
	path     string // Path to MSP dir of the org
}

// NewConfigTX new context for configtx.yaml for a new org
func NewConfigTX(tempPath string, org string, path string) *ConfigTX {
	return &ConfigTX{
		tempPath: tempPath,
		org:      org,
		path:     path,
	}
}

//ConfigTXTemplate Generating configtx.yaml
func (c *ConfigTX) ConfigTXTemplate() {
	file, err := ioutil.ReadFile(c.tempPath)
	if err != nil {
		panic("The file error ==> " + err.Error())
	}

	fileAsString := string(file)

	outPath := path.Join(c.path, "/")

	outFile := strings.ReplaceAll(fileAsString, "ORG", c.org)
	outFile = strings.ReplaceAll(outFile, "PATH", outPath)
	outBytes := []byte(outFile)
	err = ioutil.WriteFile("configtx.yaml", outBytes, 0644)

	if err != nil {
		_ = fmt.Errorf("Did not create file configtx.yaml \n")
	}
}
