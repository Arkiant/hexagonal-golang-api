FROM golang:latest AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo /app/cmd/api

FROM alpine:latest 

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/.env .

WORKDIR /

ENV ENV=""

EXPOSE 8080

COPY --from=builder /app/api .

CMD ["./api"]