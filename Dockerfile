FROM golang:1.20 AS builder
WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download
ADD . .
RUN make ttvbot
CMD ["/usr/src/app/ttvbot"]
