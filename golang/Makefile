certs:
	go build . && ./main && ./main -role=admin

clean:
	rm -r main NewOrg/

pki:
	docker run --rm --network host -v /home/sachin/ca-net/golang/collection/:/etc/postman/ \
	 -t postman/newman run /etc/postman/vault-setup-New-Org.postman_collection.json