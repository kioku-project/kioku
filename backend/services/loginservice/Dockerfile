FROM golang:alpine AS builder

# Set Go env
ENV CGO_ENABLED=0 GOOS=linux
WORKDIR /go/src/kioku

# Install dependencies
RUN apk --update --no-cache add ca-certificates gcc libtool make musl-dev protoc git protobuf-dev

# Download grpc_health_probe
RUN GRPC_HEALTH_PROBE_VERSION=v0.4.11 && \
wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
chmod +x /bin/grpc_health_probe

# Build Go binary
COPY services/loginservice/Makefile go.mod go.sum ./
RUN make init && go mod download 
COPY services/loginservice/ services/loginservice
COPY store/ store/
COPY pkg/ pkg/
COPY googleapis/ googleapis/
WORKDIR /go/src/kioku/services/loginservice
RUN make proto tidy build

# Deployment container
FROM scratch

COPY --from=builder /etc/ssl/certs /etc/ssl/certs
COPY --from=builder /bin/grpc_health_probe /bin/
COPY --from=builder /go/src/kioku/services/loginservice/loginservice /loginservice
ENTRYPOINT ["/loginservice"]
CMD []
