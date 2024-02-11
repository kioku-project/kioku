FROM golang:alpine AS builder

RUN addgroup -S kioku \
    && adduser -S kioku -G kioku

# Set Go env
ENV CGO_ENABLED=0 GOOS=linux
WORKDIR /go/src/kioku

# Install dependencies
RUN apk --update --no-cache add ca-certificates gcc libtool make musl-dev protoc git protobuf-dev

# Build Go binary
COPY services/notification/Makefile go.mod go.sum ./
RUN make init && go mod download 
COPY services/srs/ services/srs
COPY services/carddeck/proto/carddeck.* services/carddeck/proto/
COPY services/user/proto/user.* services/user/proto/
COPY services/collaboration/proto/collaboration.* services/collaboration/proto/
COPY services/notification/ services/notification
COPY store/ store/
COPY pkg/ pkg/
WORKDIR /go/src/kioku/services/notification
RUN make proto build

# Deployment container
FROM scratch
COPY --from=builder /etc/passwd /etc/passwd
USER kioku

COPY --from=builder /etc/ssl/certs /etc/ssl/certs
COPY --from=builder /bin/grpc_health_probe /bin/
COPY --from=builder /go/src/kioku/services/notification/notification /notification
ENTRYPOINT ["/notification"]
CMD []
