# Kioku
Welcome to the kioku-project!

## Deploy Kioku to Kubernetes cluster

### Prerequisites
- Kubernetes cluster
    - Ingress controller
- [`helm`](https://helm.sh/)

### Installation

It is generally recommended to install Kioku to a fresh Kubernetes cluster, because there are some dependent deployments, which are all being deployed using the `deploy.sh`:
- [`postgres-operator`](https://github.com/zalando/postgres-operator)

From the root-directory of the repository execute the following:

```bash
./deploy.sh
```

## Setup Storybook and Chromatic
 
To use [Storybook](https://storybook.js.org/) locally run `npm run storybook` and open [localhost:6006](http://localhost:6006) if it does not open automatically.
 
New to Storybook? Learn how to write stories [here](https://storybook.js.org/docs/react/writing-stories/introduction).
 
To use [Chromatic](https://www.chromatic.com/) you have to add the [chromatic-project-token](https://www.chromatic.com/manage?appId=63e354941aa15501d3467f88&view=configure) to your .env file. 

 ```
 CHROMATIC_PROJECT_TOKEN=
 ```
 
To publish storybook to Chromatic run `npm run chromatic`.
