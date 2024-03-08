#!/bin/bash
# This script allows to deploy kioku to a Kubernetes cluster
# Assumptions:
#   - KUBECONFIG already targets the cluster
VALUE_STRING=""
if ! [ -z "$1" ]; then
  VALUE_STRING="-f $1"
  echo "$1"
fi

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
helm upgrade --install kioku helm/kioku $VALUE_STRING