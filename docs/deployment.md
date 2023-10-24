# Deploy Kioku to Kubernetes cluster

## Prerequisites

- Kubernetes cluster
  - Ingress controller
- [`helm`](https://helm.sh/)

## Installation

If you wish to deploy Kioku onto a Kubernetes cluster, please ensure that an ingress controller is provisioned beforehand.
Furthermore, it is generally recommended to install Kioku to a new Kubernetes cluster, because there are some dependent deployments, which are all being deployed using the `deploy.sh` script:

- [`postgres-operator`](https://github.com/zalando/postgres-operator)

From the root-directory of the repository execute the following:

```bash
./deploy.sh
```

## Tracing
The performance and metrics of the deployment such as dependencies between services can be viewed using [`jaeger`](https://www.jaegertracing.io/). In order to access the jaeger frontend, you will have to expose the port `16686` of the `jaeger-service` in the `monitoring` namespace.