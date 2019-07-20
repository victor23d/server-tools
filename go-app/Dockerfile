FROM golang:1.11-alpine
ADD . /go/src/go-app
RUN go install go-app

FROM alpine:latest
WORKDIR /app
COPY --from=0 /go/bin/go-app /app/
COPY --from=0 /go/src/go-app/scripts /app/scripts
ENV PORT 8080
CMD ["/app/go-app"]
