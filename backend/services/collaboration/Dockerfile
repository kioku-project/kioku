FROM golang:alpine AS builder

RUN addgroup -S kioku \
    && adduser -S kioku -G kioku

# Set Go env
ENV CGO_ENABLED=0 GOOS=linux
WORKDIR /go/src/kioku

# Install dependencies
RUN apk --update --no-cache add ca-certificates gcc libtool make musl-dev protoc git protobuf-dev

# Build Go binary
COPY services/collaboration/Makefile go.mod go.sum ./
RUN make init && go mod download 
COPY services/user/proto/user.* services/user/proto/
COPY services/carddeck/proto/carddeck.* services/carddeck/proto/
COPY services/srs/proto/srs.* services/srs/proto/
COPY services/collaboration/ services/collaboration
COPY store/ store/
COPY pkg/ pkg/
WORKDIR /go/src/kioku/services/collaboration
RUN make proto build

# Deployment container
FROM scratch
COPY --from=builder /etc/passwd /etc/passwd
USER kioku

COPY --from=builder /etc/ssl/certs /etc/ssl/certs
COPY --from=builder /go/src/kioku/services/collaboration/collaboration /collaboration
ENTRYPOINT ["/collaboration"]
CMD []
