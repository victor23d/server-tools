FROM golang:1.13-alpine AS builder
WORKDIR /app/
COPY . /app/
ENV CGO_ENABLED=0
RUN go test -v -mod=vendor
# RUN go get -d -v && go install -v go-app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest AS release
RUN apk --no-cache add ca-certificates curl
WORKDIR /root/
COPY --from=builder /app/app .
CMD ["./app"]
