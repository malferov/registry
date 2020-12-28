#/bin/bash
set -e

cd ../terraform
jinja2 ../manifest/secret.yml -D crt=$(terraform output crt) -D key=$(terraform output key) \
  > ../manifest/secret.rendered

cd ../manifest
kubectl apply -f secret.rendered

jinja2 ingress.yml -D app=$app -D domain=$domain -D service=$service \
  | kubectl apply -f -
