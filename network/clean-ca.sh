#!/bin/bash

if [ -d certs ]; then
	rm -r certs
fi

if [ -d channel-artifacts ]; then
	sudo rm -r channel-artifacts
fi

pushd ./scripts/
 ./cleanup.sh
popd