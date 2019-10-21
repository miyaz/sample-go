FROM golang:1.13-alpine3.10 AS build
ENV TZ Asia/Tokyo
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN apk update && apk add git alpine-sdk tzdata
WORKDIR /go/src/github.com/miyaz/sample-go
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN make deps clean bin/sample-go

FROM alpine:3.10
ENV TZ Asia/Tokyo
RUN apk --update --no-cache add ca-certificates tzdata
COPY --from=build /go/src/github.com/miyaz/sample-go/bin/sample-go /sample-go
CMD ["/sample-go"]
