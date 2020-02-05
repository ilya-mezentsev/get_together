FROM golang:1.13.6-alpine3.10 AS build
RUN apk --no-cache add gcc g++ make
RUN apk add git

WORKDIR ${GOPATH}
COPY . .
RUN go get -v -d ./...

EXPOSE 8080
