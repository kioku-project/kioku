FROM golang:alpine AS builder

RUN addgroup -S kioku \
    && adduser -S kioku -G kioku

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
COPY services/frontend/Makefile go.mod go.sum ./
RUN make init && go mod download 
COPY services/frontend/ services/frontend
COPY services/user/proto/user.* services/user/proto/
COPY services/carddeck/proto/carddeck.* services/carddeck/proto/
COPY services/collaboration/proto/collaboration.* services/collaboration/proto/
COPY services/srs/proto/srs.* services/srs/proto/
COPY services/notification/proto/notification.* services/notification/proto/
COPY store/ store/
COPY pkg/ pkg/
WORKDIR /go/src/kioku/services/frontend
RUN make proto tidy build

# Deployment container
FROM scratch
COPY --from=builder /etc/passwd /etc/passwd

COPY --from=builder /etc/ssl/certs /etc/ssl/certs
COPY --from=builder /go/src/kioku/services/frontend/frontend /frontend
ENTRYPOINT ["/frontend"]
CMD []
