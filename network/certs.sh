#!/bin/bash

if [ ! -d certs ]; then
	mkdir -p certs/admin/tls-ca/
	mkdir -p certs/admin/org-ca/
	mkdir -p certs/admin/ord-ca/
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
export FABRIC_CA_CLIENT_HOME=$PWD/certs/admin/org-ca/
set -x
fabric-ca-client enroll -d -u https://admin:adminpw@0.0.0.0:7055
fabric-ca-client register -d --id.name peer-org --id.secret peerpw --id.type peer -u https://0.0.0.0:7055
fabric-ca-client register -d --id.name admin-org --id.secret adminpw --id.type admin -u https://0.0.0.0:7055
fabric-ca-client register -d --id.name user-org --id.secret clientpw --id.type client -u https://0.0.0.0:7055
set +x

echo'#peer
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
export FABRIC_CA_CLIENT_HOME=$PWD/certs/admin/ord-ca/
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
