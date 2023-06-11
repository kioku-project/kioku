#!/bin/bash
# This script allows to deploy kioku to an external Kubernetes cluster
# Assumptions:
#   - KUBECONFIG already targets the cluster
#   - Script is executed from kioku-repo root folder

# Get dependent repositories
helm repo add postgres-operator-charts https://opensource.zalando.com/postgres-operator/charts/postgres-operator

# Install dependent Charts
helm install postgres-operator postgres-operator-charts/postgres-operator

# Install Kioku
helm install kioku helm/kioku