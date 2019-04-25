FROM golang:1.12 AS builder
WORKDIR /go/src/github.com/jongschneider/wait-for-elasticsearch
COPY . .
RUN CGO_ENABLED=0 go install ./...

FROM alpine
WORKDIR /usr/bin
COPY --from=builder /go/bin/wait-for-elasticsearch /usr/bin/wait-for-elasticsearch
ENTRYPOINT ["wait-for-elasticsearch"]
