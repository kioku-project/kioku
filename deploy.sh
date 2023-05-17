#!/bin/bash
# This script allows to deploy kioku to an external Kubernetes cluster
# Assumptions:
#   - KUBECONFIG already targets the cluster
#   - Script is executed from kioku-repo root folder

# Get dependent repositories
helm repo add jaegertracing https://jaegertracing.github.io/helm-charts
helm repo add jetstack https://charts.jetstack.io

helm repo update

# Install cert-manager
helm install \
  cert-manager jetstack/cert-manager \
  --namespace cert-manager \
  --create-namespace \
  --version v1.11.0 \
  --set installCRDs=true

# Install Jaeger Operator
helm upgrade -i jaeger-operator jaegertracing/jaeger-operator \
    --namespace observability \
    --create-namespace

# Install Kioku
helm install kioku helm/kioku