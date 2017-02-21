#!/bin/bash

set -e

SCRIPTDIR="$(cd $(dirname "$0")/ && pwd)"

secrets/genssl.sh nginx

kubectl create secret generic edge-nginx-secrets \
 --from-file=./secrets/nginx.conf \
 --from-file=./secrets/openwhisk-cert.pem \
 --from-file=./secrets/openwhisk-key.pem
