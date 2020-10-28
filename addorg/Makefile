install:
	GOBIN=${PWD} go install ./cmd/addorg

install-server:
	GOBIN=${PWD} go install ./cmd/server

msp-config:
	./addorg -configtx_req=true

admin-certs:
	./addorg -consul_template_path=consul-template-admin.hcl -role=admin -msp=false -certs=true

peer-certs:
	./addorg -msp=false -certs=true

certs:
	./addorg -msp=false -certs=true && ./addorg -consul_template_path=consul-template-admin.hcl -role=admin -msp=false -certs=true

all: build msp-config certs

channel:
	./addorg -channel=true -msp=false

clean:
	rm -r addorg configtx.yaml config_block.pb NewOrg.json consul-template-*.hcl config.json modified_config.json NewOrg/

pki:
	docker run --rm --network host -v /home/sachin/ca-net/addorg/collection/:/etc/postman/ \
	 -t postman/newman run /etc/postman/vault-setup-New-Org.postman_collection.json

cli-msp:
	./addorg -configtx_req=true -base_path=/opt/gopath/src/addorg -vault_host=http://192.168.1.102:8200

cli: build cli-msp channel

cli-certs:
	./addorg -base_path=/opt/gopath/src/addorg -vault_host=http://192.168.1.102:8200 -certs=true  && ./addorg -base_path=/opt/gopath/src/addorg -consul_template_path="./consul-template-admin.hcl" -vault_host=http://192.168.1.102:8200 -role=admin -msp=false -certs=true