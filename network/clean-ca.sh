#!/bin/bash

if [ -d certs ]; then
	rm -r certs
fi

if [ -d channel-artifacts ]; then
	rm -r channel-artifacts
fi

pushd ./ca/tls-ca/
 ./clean-up.sh
 popd
pushd ./ca/ord-ca/
 ./clean-up.sh
popd
pushd ./ca/org-ca/
 ./clean-up.sh
popd
