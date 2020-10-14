package main

import (
	"flag"
)

func getFlags() (*string, *string, *string, *string, *string, *string, *string, *string, *string, *string, *bool, *bool, *string) {
	//Setting up the flags
	consulTempPath := flag.String("consul_template_path", "consul-template-peer.hcl", "Used to get the path for consul-template")
	basePath := flag.String("base_path", "/home/sachin/ca-net/golang", "Used to get the base path for consul-template")
	consulTemp := flag.String("consul_template", "templates/consul-template.hcl", "Used to get the template file for consul-template")
	tmpFile := flag.String("template", "templates/template.tpl", "Used to get the template file for MSP")
	tlsTmpFile := flag.String("tls_template", "templates/tlstemplate.tpl", "Used to get the template file for TLS")
	outDir := flag.String("out_dir", "./NewOrg/", "Used to get the directory for certs of the new org")
	newOrg := flag.String("new_org", "NewOrg", "Used to specify the new org name")
	vaultHost := flag.String("vault_host", "http://127.0.0.1:8200", "Used to specify the vautl http ednpoint, Default = http://127.0.0.1:8200")
	role := flag.String("role", "peer", "Used to specify the role of the certs")
	configtxFile := flag.String("configtx", "templates/configtx-templat.yaml", "To Generate the config tx for new org")
	msp := flag.Bool("msp", true, "To generate msp")
	configtxReq := flag.Bool("configtx_req", false, "To Create the configtx.yaml")
	baseDomain := flag.String("base_domain", "service.consul", "Used to get the base path for consul-template")
	//Getting there values
	flag.Parse()

	return consulTempPath, basePath, consulTemp, tmpFile, tlsTmpFile, outDir, newOrg, vaultHost, role, configtxFile, msp, configtxReq, baseDomain
}