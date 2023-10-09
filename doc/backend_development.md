# Backend development

## Index
1. [Infrastructure overview](#infrastructure-overview)
2. [Getting started](#getting-started-with-backend-development)
3. [Create a new service](#create-a-new-service)

## Infrastructure overview
![Architecture-Design](https://github.com/kioku-project/kioku/assets/60541979/4aedb2be-f1ff-41bc-ac7f-a45662bd819a)

The backend is designed with modern microservice philosophy in mind. The communication between the microservices themselves is done using gRPC which allows for fast and streamlined communication. All services can only be accessed from outside through the `frontend_proxy` service which translates REST API calls to the corresponding gRPC methods. The `frontend_proxy` is also responsible for the authentication of the user, allowing the services to assume the provided userID as safe to use. The services are written in the [go](https://go.dev/) programming language and use [go-micro](https://github.com/go-micro/go-micro) as a framework.

## Getting started with backend development
Before getting started, ensure that you have a recent version of [Go](https://golang.org) installed. Version 1.16 or higher is required.
In order to create services that are in line with our design vision at Kioku, we wrote our own fork of the go-micro cli.
You will need to install this version to be able to create new services and re-compile existing ones.
```
go install github.com/kioku-project/go-micro-cli/cmd/go-micro@latest
```
Additionally, you need to set up the `protoc` compiler in your environment. See [here](https://grpc.io/docs/protoc-installation/) for installation instructions according to your system.
And finally, you need to install `make` on your system. Mac users can install it alongside common developer packages via `xcode-select --install`, Windows users are able to obtain it via [Chocolatey](https://chocolatey.org/install) by using the command `choco install make`

## Create a new service
1. Set up a new service

        cd backend/services/
        go-micro new service --health --kubernetes github.com/kioku-project/kioku/services/<name-of-new-service>
        cd <name-of-new-service>
        make init proto update tidy

2. Make adjustments
    1. Update the service definitions in `proto/<name-of-new-service>.proto`
    2. Generate proto files by running `make proto`
    3. Update handler in `handler/<name-of-new-service>.go`

3. Update docker-compose
    
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

4. Adjust the proxy rules in the frontend service to be able to serve the new service if needed
    1. Create a new handler in `backend/services/frontend/handler/frontend.go` for a new api endpoint
    2. In `backend/services/frontend/main.go`, add the new handler with the desired route
    3. Adjust all the relevant Dockerfiles to integrate the new proto files of the new service
