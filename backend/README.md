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
        make init update tidy

2. Make adjustments
    1. Update the service definitions in `proto/<name-of-new-service>.proto`
    2. Generate proto files by running `make proto`
    3. Update handler in `handler/<name-of-new-service>.go`

3. Update docker-compose
    
    Add a new service in `docker-compose.yml` for the created service.

4. Adjust the environment for local development
    
    You need to create a protobuf file for all combined services in the landscape. This can be achieved with the following commands:

        cd backend/services
        protoc --proto_path=. -I../googleapis --include_imports --descriptor_set_out=combined.pb */proto/*.proto

    For local development, envoy is used as a proxy for the gRPC services and is also handling JSON to gRPC transcoding for the frontend. After adding a new service, you can update the `envoy.yaml` config accordingly. You need to add it to the services in the `http_filters` section, and you are also required to create a new cluster and route match so that envoy knows where to route specific requests.