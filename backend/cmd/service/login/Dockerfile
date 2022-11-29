FROM golang:1.19.3-alpine3.16

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod graph | awk '{if ($1 !~ "@") print $2}' | xargs go get

COPY . .

RUN go build -o /login_service /build/cmd/service/login/main.go

ENTRYPOINT [ "/login_service" ]
