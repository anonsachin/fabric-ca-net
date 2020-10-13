certs:
	go build . && ./main && ./main -consul_template_path=consul-template-admin.hcl -role=admin -msp=false -configtx_req=true

clean:
	rm -r main configtx.yaml NewOrg.json consul-template-*.hcl NewOrg/

pki:
	docker run --rm --network host -v /home/sachin/ca-net/golang/collection/:/etc/postman/ \
	 -t postman/newman run /etc/postman/vault-setup-New-Org.postman_collection.json

cli-certs:
	go build . && ./main -base_path=/opt/gopath/src/addorg -vault_host=http://192.168.1.102:8200 && ./main -base_path=/opt/gopath/src/addorg -vault_host=http://192.168.1.102:8200 -role=admin -msp=false -configtx_req=true