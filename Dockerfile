FROM golang:latest AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo /app/ddd-golang-framework

FROM alpine:latest 

RUN apk --no-cache add ca-certificates

WORKDIR /

ENV ENV=""

COPY --from=builder /app/ddd-golang-framework .

CMD ["./ddd-golang-framework"]