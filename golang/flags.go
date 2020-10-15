package main

import (
	"flag"
)

func getFlags() (*bool, *string, *string, *string, *string, *string, *string, *string, *string, *string, *string, *bool, *bool, *string) {
	//Setting up the flags
	// consulTempPath := flag.String("consul_template_path", "consul-template-peer.hcl", "Used to get the path for consul-template")
	// basePath := flag.String("base_path", "/home/sachin/ca-net/golang", "Used to get the base path for consul-template")
	// consulTemp := flag.String("consul_template", "templates/consul-template.hcl", "Used to get the template file for consul-template")
	// tmpFile := flag.String("template", "templates/template.tpl", "Used to get the template file for MSP")
	// tlsTmpFile := flag.String("tls_template", "templates/tlstemplate.tpl", "Used to get the template file for TLS")
	// outDir := flag.String("out_dir", "./NewOrg/", "Used to get the directory for certs of the new org")
	// newOrg := flag.String("new_org", "NewOrg", "Used to specify the new org name")
	// vaultHost := flag.String("vault_host", "http://127.0.0.1:8200", "Used to specify the vautl http ednpoint, Default = http://127.0.0.1:8200")
	// role := flag.String("role", "peer", "Used to specify the role of the certs")
	// configtxFile := flag.String("configtx", "templates/configtx-templat.yaml", "To Generate the config tx for new org")
	// msp := flag.Bool("msp", true, "To generate msp")
	// configtxReq := flag.Bool("configtx_req", false, "To Create the configtx.yaml")
	// baseDomain := flag.String("base_domain", "service.consul", "Used to get the base domain for consul-template")

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

	return certs, consulTempOutPath, basePath, consulTemp, tmpFile, tlsTmpFile, outDir, newOrg, vaultHost, role, configtxFile, msp, configtx, baseDomain
}

// All others are dependent on these flags may not be all flags
func baseFlags() (*string, *string, *string){
	outDir := flag.String("out_dir", "./NewOrg/", "Used to get the directory for certs of the new org")
	newOrg := flag.String("new_org", "NewOrg", "Used to specify the new org name")
	vaultHost := flag.String("vault_host", "http://127.0.0.1:8200", "Used to specify the vautl http ednpoint, Default = http://127.0.0.1:8200")
	//Getting there values
	flag.Parse()

	return outDir, newOrg, vaultHost
}

func mspFlags() ( *bool){
	msp := flag.Bool("msp", true, "To generate msp")
	//Getting there values
	flag.Parse()

	return msp
}

// Dependent on mspFlags
func orgConfigFlags() (*bool, *string){
	// To generate the config
	configtx := flag.Bool("configtx_req", false, "To generate org config from msp and configtx.yaml")
	// To template the configtx.yaml
	configtxFile := flag.String("configtx", "templates/configtx-templat.yaml", "To Generate the config tx for new org")
	//Getting there values
	flag.Parse()

	return configtx, configtxFile
}

func certsFlags() (*string, *string, *string, *string, *string, *string, *string){
	// For the tls and msp certs template
	tmpFile := flag.String("template", "templates/template.tpl", "Used to get the template file for MSP")
	tlsTmpFile := flag.String("tls_template", "templates/tlstemplate.tpl", "Used to get the template file for TLS")
	baseDomain := flag.String("base_domain", "service.consul", "Used to get the base domain for consul-template")
	// For consul-template
	consulTemp := flag.String("consul_template", "templates/consul-template.hcl", "Used to get the template file for consul-template")
	basePath := flag.String("base_path", "/home/sachin/ca-net/golang", "Used to get the base path for consul-template")
	consulTempOutPath := flag.String("consul_template_path", "consul-template-peer.hcl", "Used to get the path for consul-template output of templating")
	// For all
	role := flag.String("role", "peer", "Used to specify the role of the certs")
	
	//Getting there values
	flag.Parse()
	
	return tmpFile, tlsTmpFile, baseDomain, role, consulTemp, basePath, consulTempOutPath
}