package main

import(
	"net/url"
	"net/http"
	"io/ioutil"
	"fmt"
)

func getCaUrl(host string, org string, tls bool) string{
	//Getting the base
	u, _ := url.ParseRequestURI(host)
	//Assigning the path
	if tls {
		u.Path = "/v1/"+org+"TLSCA/ca/pem"
	} else {
		u.Path = "/v1/"+org+"CA/ca/pem"
	}
	// Return the url
	return u.String()
}

func getCaCert(host string, org string) ([]byte,[]byte){
	//MSP CA
	mspReq, err := http.Get(getCaUrl(host,org,false))
	if err != nil {
		_ = fmt.Errorf("The request Failed becaues %s",err.Error())
	}
	fmt.Println("the Status is ",mspReq.Status)

	if mspReq.StatusCode != 200 {
		_ = fmt.Errorf("The did'nt get the right response")
	}

	defer mspReq.Body.Close()

	body, err := ioutil.ReadAll(mspReq.Body)
	if err != nil {
		_ = fmt.Errorf("The bady parsing failed becaues %s",err.Error())
	}
	
	//TLS CA
	tlsReq, err := http.Get(getCaUrl(host,org,true))
	if err != nil {
		_ = fmt.Errorf("The request Failed becaues %s",err.Error())
	}
	fmt.Println("the Status is ",tlsReq.Status)

	if tlsReq.StatusCode != 200 {
		_ = fmt.Errorf("The did'nt get the right response")
	}

	defer tlsReq.Body.Close()

	bodytls, err := ioutil.ReadAll(tlsReq.Body)
	if err != nil {
		_ = fmt.Errorf("The bady parsing failed becaues %s",err.Error())
	}

	return body, bodytls
}