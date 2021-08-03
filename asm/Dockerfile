FROM golang:1.15 as builder

WORKDIR /root/go/src/github.com/mox692/gooo/asm

# Install.

# Add files.

ENV GOBIN=/root/go/bin

ENV GOPATH=/root/go

ENV GO111MODULE=on

ENV GOROOT=/usr/local/go

COPY ./go.mod ./

RUN go mod download

COPY ./ ./

# RUN env GOOS=linux GOARCH=amd64 go build -o /asm ./cmd/main.go 
# RUN chmod 777 ./server/server.go

CMD ["bash"]