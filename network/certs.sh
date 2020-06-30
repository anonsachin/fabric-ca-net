#!/bin/bash

if [ ! -d certs ]; then
	mkdir -p certs/admin/tls-ca/
	mkdir -p certs/peerorg/
	mkdir -p certs/ordererorg/
	mkdir -p certs/peerorg/admin
	mkdir -p certs/peerorg/peer
	mkdir -p certs/peerorg/client
	mkdir -p certs/ordererorg/orderer
fi

export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/ca/tls-ca/ca-cert.pem
export FABRIC_CA_CLIENT_HOME=$PWD/certs/admin/tls-ca/
set -x
fabric-ca-client enroll -d -u https://admin:adminpw@0.0.0.0:7054
fabric-ca-client register -d --id.name peer-org --id.secret peerpw --id.type peer -u https://0.0.0.0:7054
fabric-ca-client register -d --id.name admin-org --id.secret adminpw --id.type admin -u https://0.0.0.0:7054
fabric-ca-client register -d --id.name user-org --id.secret clientpw --id.type client -u https://0.0.0.0:7054
fabric-ca-client register -d --id.name orderer --id.secret ordererpw --id.type orderer -u https://0.0.0.0:7054
set +x

echo '#tls 
NodeOUs:
  Enable: true
  ClientOUIdentifier:
    Certificate: cacerts/0-0-0-0-7054.pem
    OrganizationalUnitIdentifier: client
  PeerOUIdentifier:
    Certificate: cacerts/0-0-0-0-7054.pem
    OrganizationalUnitIdentifier: peer
  AdminOUIdentifier:
    Certificate: cacerts/0-0-0-0-7054.pem
    OrganizationalUnitIdentifier: admin
  OrdererOUIdentifier:
    Certificate: cacerts/0-0-0-0-7054.pem
    OrganizationalUnitIdentifier: orderer'> $FABRIC_CA_CLIENT_HOME/msp/config.yaml

export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/ca/org-ca/ca-cert.pem
export FABRIC_CA_CLIENT_HOME=$PWD/certs/peerorg/
set -x
fabric-ca-client enroll -d -u https://admin:adminpw@0.0.0.0:7055
fabric-ca-client register -d --id.name peer-org --id.secret peerpw --id.type peer -u https://0.0.0.0:7055
fabric-ca-client register -d --id.name admin-org --id.secret adminpw --id.type admin -u https://0.0.0.0:7055
fabric-ca-client register -d --id.name user-org --id.secret clientpw --id.type client -u https://0.0.0.0:7055
set +x

echo '#peer
NodeOUs:
  Enable: true
  ClientOUIdentifier:
    Certificate: cacerts/0-0-0-0-7055.pem
    OrganizationalUnitIdentifier: client
  PeerOUIdentifier:
    Certificate: cacerts/0-0-0-0-7055.pem
    OrganizationalUnitIdentifier: peer
  AdminOUIdentifier:
    Certificate: cacerts/0-0-0-0-7055.pem
    OrganizationalUnitIdentifier: admin
  OrdererOUIdentifier:
    Certificate: cacerts/0-0-0-0-7055.pem
    OrganizationalUnitIdentifier: orderer' >  $FABRIC_CA_CLIENT_HOME/msp/config.yaml

export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/ca/ord-ca/ca-cert.pem
export FABRIC_CA_CLIENT_HOME=$PWD/certs/ordererorg/
set -x
fabric-ca-client enroll -d -u https://admin:adminpw@0.0.0.0:7056
fabric-ca-client register -d --id.name orderer --id.secret ordererpw --id.type orderer -u https://0.0.0.0:7056
set +x

echo '#orderer 
NodeOUs:
  Enable: true
  ClientOUIdentifier:
    Certificate: cacerts/0-0-0-0-7056.pem
    OrganizationalUnitIdentifier: client
  PeerOUIdentifier:
    Certificate: cacerts/0-0-0-0-7056.pem
    OrganizationalUnitIdentifier: peer
  AdminOUIdentifier:
    Certificate: cacerts/0-0-0-0-7056.pem
    OrganizationalUnitIdentifier: admin
  OrdererOUIdentifier:
    Certificate: cacerts/0-0-0-0-7056.pem
    OrganizationalUnitIdentifier: orderer' > $FABRIC_CA_CLIENT_HOME/msp/config.yaml

echo "####### ENROLLING ORDERER #########"

export FABRIC_CA_CLIENT_HOME=$PWD/certs/ordererorg/orderer
set -x
fabric-ca-client enroll -d -u https://orderer:ordererpw@0.0.0.0:7056
set +x

export FABRIC_CA_CLIENT_MSPDIR=msp
echo "moving the config file"
cp $PWD/certs/ordererorg/msp/config.yaml $FABRIC_CA_CLIENT_HOME/msp/config.yaml

echo "changing the root ca"
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/ca/tls-ca/ca-cert.pem

export FABRIC_CA_CLIENT_MSPDIR=tls
set -x
fabric-ca-client enroll -d -u https://orderer:ordererpw@0.0.0.0:7054 --enrollment.profile tls --csr.hosts orderer --csr.hosts localhost
set +x

cp  $FABRIC_CA_CLIENT_HOME/tls/tlscacerts/* $FABRIC_CA_CLIENT_HOME/tls/ca.crt
cp  $FABRIC_CA_CLIENT_HOME/tls/signcerts/* $FABRIC_CA_CLIENT_HOME/tls/server.crt
cp  $FABRIC_CA_CLIENT_HOME/tls/keystore/* $FABRIC_CA_CLIENT_HOME/tls/server.key

echo "editing the ord-ca admin tlscerts "

mkdir $PWD/certs/ordererorg/msp/tlscacerts
cp $FABRIC_CA_CLIENT_HOME/tls/ca.crt $PWD/certs/ordererorg/msp/tlscacerts/ca.cert

echo "editing the orderer root tlsca directory"

mkdir $PWD/certs/ordererorg/tlsca
cp $FABRIC_CA_CLIENT_HOME/tls/ca.crt $PWD/certs/ordererorg/tlsca/tls-ca.crt

echo "editing root ca directory"

mkdir $PWD/certs/ordererorg/ca
cp $FABRIC_CA_CLIENT_HOME/msp/cacerts/* $PWD/certs/ordererorg/ca/orderer-ca.cert

echo "####### ENROLLING PEER #########"

export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/ca/org-ca/ca-cert.pem
export FABRIC_CA_CLIENT_HOME=$PWD/certs/peerorg/peer
export FABRIC_CA_CLIENT_MSPDIR=msp
set -x
fabric-ca-client enroll -d -u https://peer-org:peerpw@0.0.0.0:7055
set +x

echo "moving the config file"
cp $PWD/certs/peerorg/msp/config.yaml $FABRIC_CA_CLIENT_HOME/msp/config.yaml

echo "changing the root ca"
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/ca/tls-ca/ca-cert.pem

export FABRIC_CA_CLIENT_MSPDIR=tls
set -x
fabric-ca-client enroll -d -u https://peer-org:peerpw@0.0.0.0:7054 --enrollment.profile tls --csr.hosts peer-org --csr.hosts localhost
set +x

cp  $FABRIC_CA_CLIENT_HOME/tls/tlscacerts/* $FABRIC_CA_CLIENT_HOME/tls/ca.crt
cp  $FABRIC_CA_CLIENT_HOME/tls/signcerts/* $FABRIC_CA_CLIENT_HOME/tls/server.crt
cp  $FABRIC_CA_CLIENT_HOME/tls/keystore/* $FABRIC_CA_CLIENT_HOME/tls/server.key

echo "editing the ord-ca admin tlscerts "

mkdir $PWD/certs/peerorg/msp/tlscacerts
cp $FABRIC_CA_CLIENT_HOME/tls/ca.crt $PWD/certs/peerorg/msp/tlscacerts/ca.cert

echo "editing the orderer root tlsca directory"

mkdir $PWD/certs/peerorg/tlsca
cp $FABRIC_CA_CLIENT_HOME/tls/ca.crt $PWD/certs/peerorg/tlsca/tls-ca.crt

echo "editing root ca directory"
mkdir $PWD/certs/peerorg/ca
cp $FABRIC_CA_CLIENT_HOME/msp/cacerts/* $PWD/certs/peerorg/ca/peerorg-ca.cert

echo "####### ENROLLING PEER ADMIN #########"

export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/ca/org-ca/ca-cert.pem
export FABRIC_CA_CLIENT_HOME=$PWD/certs/peerorg/admin
export FABRIC_CA_CLIENT_MSPDIR=msp
set -x
fabric-ca-client enroll -d -u https://admin-org:adminpw@0.0.0.0:7055
set +x

echo "moving the config file"
cp $PWD/certs/peerorg/msp/config.yaml $FABRIC_CA_CLIENT_HOME/msp/config.yaml

echo "changing the root ca"
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/ca/tls-ca/ca-cert.pem

export FABRIC_CA_CLIENT_MSPDIR=tls
set -x
fabric-ca-client enroll -d -u https://admin-org:adminpw@0.0.0.0:7054 --enrollment.profile tls --csr.hosts admin-org --csr.hosts localhost
set +x

cp  $FABRIC_CA_CLIENT_HOME/tls/tlscacerts/* $FABRIC_CA_CLIENT_HOME/tls/ca.crt
cp  $FABRIC_CA_CLIENT_HOME/tls/signcerts/* $FABRIC_CA_CLIENT_HOME/tls/server.crt
cp  $FABRIC_CA_CLIENT_HOME/tls/keystore/* $FABRIC_CA_CLIENT_HOME/tls/server.key

echo "####### ENROLLING PEER CLIENT #########"

export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/ca/org-ca/ca-cert.pem
export FABRIC_CA_CLIENT_HOME=$PWD/certs/peerorg/client
export FABRIC_CA_CLIENT_MSPDIR=msp
set -x
fabric-ca-client enroll -d -u https://user-org:clientpw@0.0.0.0:7055
set +x

echo "moving the config file"
cp $PWD/certs/peerorg/msp/config.yaml $FABRIC_CA_CLIENT_HOME/msp/config.yaml

echo "changing the root ca"
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/ca/tls-ca/ca-cert.pem

export FABRIC_CA_CLIENT_MSPDIR=tls
set -x
fabric-ca-client enroll -d -u https://user-org:clientpw@0.0.0.0:7054 --enrollment.profile tls --csr.hosts client-org --csr.hosts localhost
set +x

cp  $FABRIC_CA_CLIENT_HOME/tls/tlscacerts/* $FABRIC_CA_CLIENT_HOME/tls/ca.crt
cp  $FABRIC_CA_CLIENT_HOME/tls/signcerts/* $FABRIC_CA_CLIENT_HOME/tls/server.crt
cp  $FABRIC_CA_CLIENT_HOME/tls/keystore/* $FABRIC_CA_CLIENT_HOME/tls/server.key