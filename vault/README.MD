VAULT PKI
=========

* To bring up vault run `make vault`
* To Create the *PKI* engine run `make pki`
* To bring down vault run `make vault-down`

Use the `vault-setup Simple Network.postman_collection.json` to setup PKI for *orderer* and one org named *org* 

* Move the Vault Address to `:8200` to help with golang binary to run inside the cli container and still keep the current values.