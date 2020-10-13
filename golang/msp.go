package main

import (
	"syscall"
	"os/exec"
	"os"
	"path/filepath"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func getCaURL(host string, org string, tls bool) string {
	//Getting the base
	u, _ := url.ParseRequestURI(host)
	//Assigning the path
	if tls {
		u.Path = "/v1/" + org + "TLSCA/ca/pem"
	} else {
		u.Path = "/v1/" + org + "CA/ca/pem"
	}
	// Return the url
	return u.String()
}

func getCaCert(host string, org string) ([]byte, []byte) {
	//MSP CA
	mspReq, err := http.Get(getCaURL(host, org, false))
	if err != nil {
		_ = fmt.Errorf("The request Failed becaues %s", err.Error())
	}
	fmt.Println("the Status is ", mspReq.Status)

	if mspReq.StatusCode != 200 {
		_ = fmt.Errorf("The did'nt get the right response")
	}

	defer mspReq.Body.Close()

	body, err := ioutil.ReadAll(mspReq.Body)
	if err != nil {
		_ = fmt.Errorf("The bady parsing failed becaues %s", err.Error())
	}

	//TLS CA
	tlsReq, err := http.Get(getCaURL(host, org, true))
	if err != nil {
		_ = fmt.Errorf("The request Failed becaues %s", err.Error())
	}
	fmt.Println("the Status is ", tlsReq.Status)

	if tlsReq.StatusCode != 200 {
		_ = fmt.Errorf("The did'nt get the right response")
	}

	defer tlsReq.Body.Close()

	bodytls, err := ioutil.ReadAll(tlsReq.Body)
	if err != nil {
		_ = fmt.Errorf("The bady parsing failed becaues %s", err.Error())
	}

	return body, bodytls
}

// CreateMSP gets the ca's and stores them in the right place
func CreateMSP(vaultHost string, newOrg string, outDir string){
		//Getting the MSP's
		mspCa, tlsCa := getCaCert(vaultHost, newOrg)
		//Creating the Dirctory of MSP CA
		path := filepath.Join(outDir, "msp/cacerts")
		os.MkdirAll(path, os.ModePerm)
		//Writing to file
		path = filepath.Join(path, "ca.pem")
		err := ioutil.WriteFile(path, mspCa, 0644)
		if err != nil {
			_ = fmt.Errorf("Did not file %s", path)
		}
		//Creating the Dirctory of TLS CA
		path = filepath.Join(outDir, "msp/tlscacerts")
		os.MkdirAll(path, os.ModePerm)
		//Writing to file
		path = filepath.Join(path, "ca.pem")
		err = ioutil.WriteFile(path, tlsCa, 0644)
		if err != nil {
			_ = fmt.Errorf("Did not file %s", path)
		}
}

// GenCerts to genrate the certs
func GenCerts(consulPath string){
	// Getting configtxgen
	binary, lookErr := exec.LookPath("bash")
    if lookErr != nil {
        panic(lookErr)
	}
	// ENV for exec
	env := os.Environ()

	// The command to run

	command  := "consul-template -config "+consulPath+" -once"

	args := []string{"bash","-c", command}

	execErr := syscall.Exec(binary, args, env)
    if execErr != nil {
        panic(execErr)
    }
}