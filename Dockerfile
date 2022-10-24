FROM golang:1.16-alpine as builder
WORKDIR /project
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o main main.go

FROM alpine:latest as dns

WORKDIR /dns
COPY --from=builder /project/main .
RUN chmod +x ./main
CMD ["/dns/main"]
EXPOSE 3000