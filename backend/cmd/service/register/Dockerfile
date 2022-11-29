FROM golang:1.19.3-alpine3.16

WORKDIR /build

# We need gcc installed for sqlite drivers
RUN apk add build-base

COPY go.mod go.sum ./

RUN go mod graph | awk '{if ($1 !~ "@") print $2}' | xargs go get

COPY . .

RUN go build -o /register_service /build/cmd/service/register/main.go

ENTRYPOINT [ "/register_service" ]
