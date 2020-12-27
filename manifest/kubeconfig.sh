#!/bin/bash
set -e
kubectl config set-cluster kube --server=https://$K8S_SERVER:6443
kubectl config set clusters.kube.certificate-authority-data $K8S_CA
kubectl config set-context ci@kube --cluster=kube --user=ci --namespace=$app
kubectl config use-context ci@kube
kubectl config set users.ci.client-certificate-data $K8S_CRT
kubectl config set users.ci.client-key-data $K8S_KEY
