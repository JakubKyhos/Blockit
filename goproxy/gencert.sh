#!/bin/bash
set -ex
# generate CA's  key and cert
openssl req -x509 -newkey rsa:4096 \
  -keyout key.pem \
  -out ca.pem \
  -sha256 \
  -days 3650 \
  -nodes
