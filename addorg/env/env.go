package env

import (
	"os"
)

type PeerContext struct {
	MSPID string
	MSPCONFIGPATH string
	TLSROOTCERT string
	PEERADDRESS string
	ORDERERTLSCA string
	CHANNELNAME string
}

// DefaultOrgEnv Points to default org peer of the network
func DefaultOrgEnv() *PeerContext{
	return &PeerContext{
		MSPID: "orgMSP",
		MSPCONFIGPATH: "/etc/hyperledger/fabric/msp",
		TLSROOTCERT: "/etc/hyperledger/fabric/tls/ca.crt",
		PEERADDRESS: "peer.testnetwork.com:7051",
		ORDERERTLSCA: "/etc/hyperledger/fabric/orderer-ca.cert",
		CHANNELNAME: "testchannel",
	}
}

// SetEnv to set the context
func (p *PeerContext) SetEnv(){
	os.Setenv("CORE_PEER_LOCALMSPID", p.MSPID)
	os.Setenv("CORE_PEER_MSPCONFIGPATH", p.MSPCONFIGPATH)
	os.Setenv("CORE_PEER_TLS_ROOTCERT_FILE", p.TLSROOTCERT)
	os.Setenv("CORE_PEER_ADDRESS",p.PEERADDRESS)
	os.Setenv("ORDERER_CA_CERT",p.ORDERERTLSCA)
	os.Setenv("CHANNEL_NAME",p.CHANNELNAME)
}