pki:
	docker run --rm --network host -v /home/sachin/ca-net/vault/collection:/etc/postman/ -t postman/newman run /etc/postman/vault-setup-Simple-Network.postman_collection.json

vault:
	docker-compose up -d

vault-down:
	docker-compose down -v