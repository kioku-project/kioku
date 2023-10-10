#!/bin/bash
# This script allows to deploy kioku to an external Kubernetes cluster
# Assumptions:
#   - KUBECONFIG already targets the cluster
#   - Script is executed from kioku-repo root folder

# Get dependent repositories
helm repo add postgres-operator-charts https://opensource.zalando.com/postgres-operator/charts/postgres-operator

# Update local Helm chart repository cache
helm repo update

# Install dependent Charts
helm install postgres-operator postgres-operator-charts/postgres-operator
# TODO: When migrating to Jaeger production environment, Cassandra should be used to persist span traces
# helm install cassandra oci://registry-1.docker.io/bitnamicharts/cassandra \
#   -f helm/kioku/cassandra_values.yaml \
#   -n monitoring \
#   --create-namespace

# Install Kioku
helm install kioku helm/kioku