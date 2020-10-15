package flags

import (
	"flag"
)

// Flags are all the inputs reuired
type Flags struct {
	outDir    *string
	newOrg    *string
	vaultHost *string

	msp *bool

	certs             *bool
	tmpFile           *string
	tlsTmpFile        *string
	baseDomain        *string
	consulTemp        *string
	basePath          *string
	consulTempOutPath *string
	role              *string

	configtx     *bool
	configtxFile *string
}

func GetFlags() *Flags {
	// Base Flags
	outDir := flag.String("out_dir", "./NewOrg/", "Used to get the directory for certs of the new org")
	newOrg := flag.String("new_org", "NewOrg", "Used to specify the new org name")
	vaultHost := flag.String("vault_host", "http://127.0.0.1:8200", "Used to specify the vautl http ednpoint, Default = http://127.0.0.1:8200")
	// MSP
	msp := flag.Bool("msp", true, "To generate msp")
	//Certs
	// For the tls and msp certs template
	certs := flag.Bool("certs", false, "To generate certs for a role")
	tmpFile := flag.String("template", "templates/template.tpl", "Used to get the template file for MSP")
	tlsTmpFile := flag.String("tls_template", "templates/tlstemplate.tpl", "Used to get the template file for TLS")
	baseDomain := flag.String("base_domain", "service.consul", "Used to get the base domain for consul-template")
	// For consul-template
	consulTemp := flag.String("consul_template", "templates/consul-template.hcl", "Used to get the template file for consul-template")
	basePath := flag.String("base_path", "/home/sachin/ca-net/golang", "Used to get the base path for consul-template")
	consulTempOutPath := flag.String("consul_template_path", "consul-template-peer.hcl", "Used to get the path for consul-template output of templating")
	// For all certs
	role := flag.String("role", "peer", "Used to specify the role of the certs")

	// For configtx.yaml
	// To generate the config
	configtx := flag.Bool("configtx_req", false, "To generate org config from msp and configtx.yaml")
	// To template the configtx.yaml
	configtxFile := flag.String("configtx", "templates/configtx-templat.yaml", "To Generate the config tx for new org")

	//Getting there values
	flag.Parse()

	return &Flags{
		outDir:    outDir,
		newOrg:    newOrg,
		vaultHost: vaultHost,

		msp: msp,

		certs:             certs,
		tmpFile:           tmpFile,
		tlsTmpFile:        tlsTmpFile,
		baseDomain:        baseDomain,
		consulTemp:        consulTemp,
		basePath:          basePath,
		consulTempOutPath: consulTempOutPath,
		role:              role,

		configtx:     configtx,
		configtxFile: configtxFile,
	}
}

// All others are dependent on these flags may not be all flags
func (f *Flags) BaseFlags() (*string, *string, *string) {

	return f.outDir, f.newOrg, f.vaultHost
}

func (f *Flags) MSPFlags() *bool {

	return f.msp
}

// Dependent on mspFlags
func (f *Flags) OrgConfigFlags() (*bool, *string) {

	return f.configtx, f.configtxFile
}

func (f *Flags) CertsFlags() (*string, *string, *string, *string, *string, *string, *string, *bool) {

	return f.tmpFile, f.tlsTmpFile, f.baseDomain, f.role, f.consulTemp, f.basePath, f.consulTempOutPath, f.certs
}