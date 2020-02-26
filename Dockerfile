FROM golang:1.14.0-alpine as build
WORKDIR /app
COPY duckdns_updater.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o duckdns_updater

FROM alpine:3.11
WORKDIR /app
COPY --from=build /app/duckdns_updater .
CMD ["./duckdns_updater"]