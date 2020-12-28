#/bin/bash
set -e

port=5000

jinja2 deployment.yml -D service=$1 -D port=$port -D tag=$tag |
  kubectl apply -f -

jinja2 service.yml -D service=$1 -D port=$port |
  kubectl apply -f -
