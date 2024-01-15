FROM golang:alpine AS builder
WORKDIR /app/api
COPY go.mod go.sum ./
COPY . .
RUN go install github.com/swaggo/swag/cmd/swag@v1.7.8
RUN swag init
RUN CGO_ENABLED=0 GOOS=linux go build -o api

FROM alpine:latest
WORKDIR /app/api
COPY --from=builder /app/api .
ENTRYPOINT [ "./api" ]
