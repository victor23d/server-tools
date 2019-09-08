FROM golang:1.13-alpine
WORKDIR /app
ADD . /app
# RUN  go install -v -mod=vendor go-app
RUN go get -v && go install -v go-app

FROM alpine:latest
WORKDIR /app
COPY --from=0 /go/bin/go-app /app/
COPY --from=0 /app/scripts /app/scripts/
ENV PORT 8080
CMD ["/app/go-app"]
