#!/bin/bash

if [ ! -d certs ]; then
	mkdir -p certs/peerorg/admin
	mkdir -p certs/peerorg/peer
	mkdir -p certs/ordererorg/orderer
fi

function msp(){
export ORG=$1
export ORG_PATH=$PWD/certs/$ORG

echo "##### GENERATING CA CERT ######"
  mkdir -p $ORG_PATH/msp/cacerts/
  touch $ORG_PATH/msp/cacerts/ca.pem
  curl \
    $( echo "http://127.0.0.1:8200/v1/${ROLE}CA/ca/pem") > $ORG_PATH/msp/cacerts/ca.pem

echo "##### GENERATING TLS CA CERT ######"
  mkdir -p $ORG_PATH/msp/tlscacerts/
  touch $ORG_PATH/msp/tlscacerts/ca.pem
  curl \
    $( echo "http://127.0.0.1:8200/v1/${ROLE}TLSCA/ca/pem") > $ORG_PATH/msp/tlscacerts/ca.pem

echo 'NodeOUs:
  Enable: true
  ClientOUIdentifier:
    Certificate: cacerts/ca.pem
    OrganizationalUnitIdentifier: client
  PeerOUIdentifier:
    Certificate: cacerts/ca.pem
    OrganizationalUnitIdentifier: peer
  AdminOUIdentifier:
    Certificate: cacerts/ca.pem
    OrganizationalUnitIdentifier: admin
  OrdererOUIdentifier:
    Certificate: cacerts/ca.pem
    OrganizationalUnitIdentifier: orderer' > $ORG_PATH/msp/config.yaml
}

function role() {
export ORG=$1
export ROLE=$2
export CNAME=$3
export TTL=$4
echo "#### CREATING THE TEMPLATES FOR MSP ####"

export ORG_PATH=$PWD/certs/$ORG/$ROLE

mkdir -p $ORG_PATH/msp/signcerts/
touch $ORG_PATH/msp/signcerts/cert.tpl
mkdir -p $ORG_PATH/msp/keystore/
touch $ORG_PATH/msp/keystore/key.tpl
mkdir -p $ORG_PATH/msp/cacerts/
touch $ORG_PATH/msp/cacerts/ca.tpl

echo "{{ with secret \"${ORG}CA/issue/${ROLE}\" \"common_name=${CNAME}\" \"ttl=${TTL}\" \"alt_names=localhost,${CNAME}\" \"ip_sans=127.0.0.1\"}}
{{ .Data.certificate }}
{{ end }}" > $ORG_PATH/msp/signcerts/cert.tpl

echo "{{ with secret \"${ORG}CA/issue/${ROLE}\" \"common_name=${CNAME}\" \"ttl=${TTL}\" \"alt_names=localhost,${CNAME}\" \"ip_sans=127.0.0.1\"}}
{{ .Data.private_key }}
{{ end }}" > $ORG_PATH/msp/keystore/key.tpl

echo "{{ with secret \"${ORG}CA/issue/${ROLE}\" \"common_name=${CNAME}\" \"ttl=${TTL}\" \"alt_names=localhost,${CNAME}\" \"ip_sans=127.0.0.1\"}}
{{ .Data.issuing_ca }}
{{ end }}" > $ORG_PATH/msp/cacerts/ca.tpl

echo 'NodeOUs:
  Enable: true
  ClientOUIdentifier:
    Certificate: cacerts/ca.pem
    OrganizationalUnitIdentifier: client
  PeerOUIdentifier:
    Certificate: cacerts/ca.pem
    OrganizationalUnitIdentifier: peer
  AdminOUIdentifier:
    Certificate: cacerts/ca.pem
    OrganizationalUnitIdentifier: admin
  OrdererOUIdentifier:
    Certificate: cacerts/ca.pem
    OrganizationalUnitIdentifier: orderer' > $ORG_PATH/msp/config.yaml

echo "#### CREATING THE TEMPLATES FOR TLS ####"

mkdir -p $ORG_PATH/tls/
touch $ORG_PATH/tls/cert.tpl
touch $ORG_PATH/tls/key.tpl
touch $ORG_PATH/tls/ca.tpl
  
echo "{{ with secret \"${ORG}CA/issue/${ROLE}\" \"common_name=${CNAME}\" \"ttl=${TTL}\" \"alt_names=localhost,${CNAME}\" \"ip_sans=127.0.0.1\"}}
{{ .Data.certificate }}
{{ end }}" > $ORG_PATH/tls/cert.tpl

echo "{{ with secret \"${ORG}CA/issue/${ROLE}\" \"common_name=${CNAME}\" \"ttl=${TTL}\" \"alt_names=localhost,${CNAME}\" \"ip_sans=127.0.0.1\"}}
{{ .Data.private_key }}
{{ end }}" > $ORG_PATH/tls/key.tpl

echo "{{ with secret \"${ORG}CA/issue/${ROLE}\" \"common_name=${CNAME}\" \"ttl=${TTL}\" \"alt_names=localhost,${CNAME}\" \"ip_sans=127.0.0.1\"}}
{{ .Data.issuing_ca }}
{{ end }}" > $ORG_PATH/tls/ca.tpl

}


msp ordererorg
msp peerorg

role ordererorg orderer orderer.testnetwork.com 2400h

role peerorg admin admin.testnetwork.com 2400h

role peerorg peer peer.testnetwork.com 2400h