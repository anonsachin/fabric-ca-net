.PHONY: newOrg
newOrg:
	protoc --go_out=. newOrg.proto

clean: 
	rm -r server/

newOrg-grpc:
	protoc --go-grpc_out=. newOrg.proto