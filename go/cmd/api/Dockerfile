FROM golang:1.17 AS builder
WORKDIR /workspace
COPY go.* /workspace/
RUN go mod download -x
COPY . /workspace/
RUN GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o /go/bin/app /workspace/cmd/api

FROM alpine:3.12.1
WORKDIR /go/bin
COPY --from=builder /go/bin/app /go/bin/
CMD [ "/go/bin/app" ]
