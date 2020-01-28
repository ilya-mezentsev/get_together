FROM golang:1.13.5-alpine3.10 AS build
# Support CGO and SSL
RUN apk --no-cache add gcc g++ make
RUN apk add git
WORKDIR ${GOPATH}
RUN go get github.com/dgrijalva/jwt-go
RUN go get github.com/jmoiron/sqlx
RUN go get github.com/lib/pq

COPY . .
EXPOSE 8080
