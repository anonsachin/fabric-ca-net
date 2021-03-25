Fabric network with Vault PKI engine for PKI
============================================

Requirements
------------
consul-template version 0.19.5 is used choose for operating system of your choice - https://releases.hashicorp.com/consul-template/0.19.5/

Three segments 
--------------

* vault dir - for vault docker container and pki setup.
* network dir - for the fabric network setup and shell scripts.
* connection dir - common connection profile to connect to network if needed.


To Run
------

* First cd into `vault/` directory and run `make vault` to start vault and then `make pki` to setup the *PKI*
* Then cd into `network/` directory and run `make up` to create the cert, the config, bring up the netwrok and run the base setup transactions to setup the blockchain.
* Then in `network/` run `sudo make clean` to clean up the network, the artifcats and certs.
* Then in `vault/` run `make vault-down` to bring down the vault server container.

For more details explore the Makefile's and the shell scripts