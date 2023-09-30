#!/bin/bash
# This script allows to deploy kioku to an external Kubernetes cluster
# Assumptions:
#   - KUBECONFIG already targets the cluster
#   - Script is executed from kioku-repo root folder

# Get dependent repositories
helm repo add postgres-operator-charts https://opensource.zalando.com/postgres-operator/charts/postgres-operator
helm repo add jaegertracing https://jaegertracing.github.io/helm-charts
helm repo add jetstack https://charts.jetstack.io

# Update local Helm chart repository cache
helm repo update

# Install dependent Charts
helm install postgres-operator postgres-operator-charts/postgres-operator

helm install \
  cert-manager jetstack/cert-manager \
  --namespace cert-manager \
  --create-namespace \
  --version v1.13.1 \
  --set installCRDs=true

helm install jaeger-operator jaegertracing/jaeger-operator -n observability --create-namespace

# Install Kioku
helm install kioku helm/kioku