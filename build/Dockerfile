FROM golang:1.21-alpine as builder
WORKDIR /workspace
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o tiny-url-server cmd/main.go

FROM  gcr.io/distroless/static:latest
WORKDIR /workspace
COPY --from=builder /workspace/tiny-url-server .
COPY build build

ENTRYPOINT ["/workspace/tiny-url-server"]