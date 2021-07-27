FROM golang:1.16-alpine

WORKDIR /build
RUN mkdir -p /out

ADD ./api/ /build

EXPOSE 4000
CMD go run main.go
