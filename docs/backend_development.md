# Backend development

## Table of Contents

- [Backend development](#backend-development)
- [Table of Contents](#table-of-contents)
- [Infrastructure overview](#infrastructure-overview)
- [Getting started with backend development](#getting-started-with-backend-development)
- [Create a new service](#create-a-new-service)

## Infrastructure overview

![Architecture-Design](https://github.com/kioku-project/kioku/assets/60541979/4aedb2be-f1ff-41bc-ac7f-a45662bd819a)

The backend is designed with modern microservice philosophy in mind. The communication between the microservices themselves is done using gRPC which allows for fast and streamlined communication. All services can only be accessed from outside through the `frontend_proxy` service which translates REST API calls to the corresponding gRPC methods. The `frontend_proxy` is also responsible for the authentication of the user, allowing the services to assume the provided userID as safe to use. The services are written in the [Go](https://go.dev/) programming language and use [go-micro](https://github.com/go-micro/go-micro) as a framework.

## Getting started with backend development

Before getting started, ensure that you have a recent version of [Go](https://go.dev) installed. Version 1.16 or higher is required.
In order to create services that are in line with our design vision at Kioku, we wrote our own fork of the go-micro cli.
You will need to install this version to be able to create new services and re-compile existing ones.

```
go install github.com/kioku-project/go-micro-cli/cmd/go-micro@latest
```

Additionally, you need to set up the `protoc` compiler in your environment. See [here](https://grpc.io/docs/protoc-installation/) for installation instructions according to your system.
And finally, you need to install `make` on your system. Mac users can install it alongside common developer packages via `xcode-select --install`, Windows users are able to obtain it via [Chocolatey](https://chocolatey.org/install) by using the command `choco install make`

## Create a new service

1.  Set up a new service

```
cd backend/services/
go-micro new service --health --kubernetes github.com/kioku-project/kioku/services/<name-of-new-service>
cd <name-of-new-service>
make init proto update tidy
```

2.  Make adjustments

    1. Update the service definitions in `proto/<name-of-new-service>.proto`
    2. Generate proto files by running `make proto`
    3. Update handler in `handler/<name-of-new-service>.go`
    4. Instrument the service in `services/<name-of-new-service>/main.go`
    ```go
    tp, err := helper.SetupTracing(context.TODO(), service)
    if err != nil {
      logger.Fatal("Error setting up tracer: %v", err)
    }
    defer func() {
      if err := tp.Shutdown(context.Background()); err != nil {
        logger.Error("Error shutting down tracer provider: %v", err)
      }
    }()
    ```

3.  Update docker-compose
    Add a new service in `docker-compose.yml` for the created service.

```yaml
<name>_service:
  build:
    context: backend
    dockerfile: services/<name>/Dockerfile
  container_name: kioku-<name>_service
  restart: always
  env_file:
    - ./backend/.env
  depends_on:
    - db
```

4. Add Service to GitHub Workflows in `.github/workflows/<name>_service.yaml`

```yaml
name: <name> Service

on:
  pull_request:
    branches: [main]
    paths:
      - "backend/services/<name>/**"
      - "backend/store/**"

jobs:
  build-carddeck:
    uses: ./.github/workflows/build_service.yml
    with:
      image-name: kioku_<name>
      path: ./backend/services/<name>
      context: ./backend
```

5. Add the service to the Kubernetes deployment

   1. Add the Kubernetes template in `helm/kioku/templates/<name>.yaml`

   ```yaml
   ---
   apiVersion: apps/v1
   kind: Deployment
   metadata:
     name: "{{ .Values.<name>.name }}-deployment"
     labels:
       {{- include "kioku.<name>.labels" . | nindent 4 }}
   spec:
     replicas: 1
     selector:
       matchLabels:
         {{- include "kioku.<name>.labels" . | nindent 6 }}
     template:
       metadata:
         labels:
           {{- include "kioku.<name>.labels" . | nindent 8 }}
       spec:
         serviceAccountName: go-micro
         containers:
         - name: {{ .Values.<name>.name }}
           image: "{{ .Values.<name>.image }}:{{ .Values.<name>.tag }}"
         {{ if eq .Values.mode "production" }}
           imagePullPolicy: Always
         {{ else }}
           imagePullPolicy: Never
         {{ end }}
           ports:
             - containerPort: 8080
           resources:
             limits:
               cpu: 500m
               memory: 500M
             requests:
               cpu: 200m
               memory: 200M
           env:
             - name: HOSTNAME
               valueFrom:
                 fieldRef:
                   fieldPath: metadata.name
             - name: PORT
               value: "8080"
             - name: POSTGRES_PASSWORD
               valueFrom:
                 secretKeyRef:
                   name: {{ print "postgres." .Values.database.databaseName ".credentials.postgresql.acid.zalan.do" }}
                   key: password

           envFrom:
             - secretRef:
                 name: {{ .Values.database.secret.name }}
             - configMapRef:
                 name: service-env
   ---
   apiVersion: v1
   kind: Service
   metadata:
     name: "{{ .Values.<name>.name }}-service"
   spec:
     selector:
       {{- include "kioku.<name>.labels" . | nindent 4 }}
     ports:
       - port: 8080
         targetPort: 8080
   ---
   apiVersion: autoscaling/v1
   kind: HorizontalPodAutoscaler
   metadata:
   name: "hpa-{{ .Values.<name>.name }}-deployment"
   spec:
   scaleTargetRef:
     apiVersion: apps/v1
     kind: Deployment
     name: "{{ .Values.<name>.name }}-deployment"
   minReplicas: {{ .Values.<name>.autoscaler.min }}
   maxReplicas: {{ .Values.<name>.autoscaler.max }}
   targetCPUUtilizationPercentage: {{ .Values.<name>.autoscaler.targetCPUUtilizationPercentage }}
   ```

   2. Add configuration values in `helm/kioku/values.yaml`

   ```yaml
   carddeck:
   name: kioku-<name>
   image: ghcr.io/kioku-project/kioku_<name>
   tag: prod

   autoscaler:
     min: 1
     max: 10
     targetCPUUtilizationPercentage: 50
   ```

   3. Add labels `helm/kioku/templates/_helpers.tpl`

   ```yaml
   {{/*
   <name> labels
   */}}
   {{- define "kioku.<name>.labels" -}}
   app.kubernetes.io/name: {{ .Values.<name>.name }}
   {{- end }}
   ```

6. Adjust the proxy rules in the frontend service to be able to serve the new service if needed
   1. Create a new handler in `backend/services/frontend/handler/frontend.go` for a new API endpoint
   2. In `backend/services/frontend/main.go`, add the new handler with the desired route
   3. Adjust all the relevant Dockerfiles to integrate the new proto files of the new service
