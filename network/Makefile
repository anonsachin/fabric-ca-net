net-up:
	docker-compose up -d

certs:
	./certs.sh

artifacts:
	./artifacts.sh

net-down:
	docker-compose down

net-down-volume:
	docker-compose down -v

clean-certs/artifacts:
	./clean-ca.sh

install-channel:
	docker exec cli ./install.sh

clean: net-down-volume clean-certs/artifacts

up: certs artifacts net-up install-channel