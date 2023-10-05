# Deploy Kioku to Kubernetes cluster
### Prerequisites
- Kubernetes cluster
    - Ingress controller
- [`helm`](https://helm.sh/)

### Installation
If you wish to deploy kioku onto a Kubernetes cluster, please ensure that an ingress controller is provisioned beforehand.
Furthermore, it is generally recommended to install Kioku to a fresh Kubernetes cluster, because there are some dependent deployments, which are all being deployed using the `deploy.sh` script:
- [`postgres-operator`](https://github.com/zalando/postgres-operator)

From the root-directory of the repository execute the following:

```bash
./deploy.sh
```