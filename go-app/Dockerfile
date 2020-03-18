FROM golang:1.13-alpine AS ci
WORKDIR /app
ADD . /app
ENV CGO_ENABLED=0
RUN go test -v -mod=vendor
# RUN go get -v && go install -v go-app

FROM ci AS build
RUN go install -v -mod=vendor /app/main.go

FROM alpine:latest AS release
WORKDIR /app
COPY --from=build /go/bin/main /app/
CMD ["/app/main"]
