#!/bin/bash

if [ -d certs ]; then
	rm -r certs
fi

if [ -d channel-artifacts ]; then
	sudo rm -r channel-artifacts
fi
# This is for chaincode
# pushd ./scripts/
#  ./cleanup.sh
# popd