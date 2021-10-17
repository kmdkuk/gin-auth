FROM golang:1.17.1-buster as builder

WORKDIR /workspace

COPY . .

RUN make build

FROM ubuntu:20.04

COPY --from=builder /workspace/bin/gin-auth /usr/local/bin

CMD ["/usr/local/bin/gin-auth"]
