# Kioku
Welcome to the kioku-backend! The go-micro framework is used for the creation of services in the context of the application.

## Set-up
Set up the specialized go-micro CLI for kioku by running

    go install github.com/kioku-project/go-micro-cli/cmd/go-micro@latest

Additionally, you need to set up the `protoc` compiler in your environment. See [here](https://grpc.io/docs/protoc-installation/) for installation instructions according to your system.

And finally, you need to install `make` on your system. Mac users can install it alongside common developer packages via `xcode-select --install`, Windows users are able to obtain it via [Chocolatey](https://chocolatey.org/install) by using the command `choco install make`.

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

4. Adjust the proxy rules in the frontend service to be able to serve the new service if needed
    1. Create a new handler in `backend/services/frontend/handler/frontend.go` for a new api endpoint
    2. In `backend/services/frontend/main.go`, add the new handler with the desired route
    3. Adjust all the relevant Dockerfiles to integrate the new proto files of the new service