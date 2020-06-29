#!/bin/bash

if [ -d certs ]; then
	rm -r certs
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
