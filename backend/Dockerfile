
FROM golang:1.23.2-alpine AS build

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . .

RUN go build -o /app/main ./cmd/app

FROM alpine:latest

WORKDIR /root/

COPY --from=build /app/main .
COPY --from=build /app/configs ./configs

CMD ["./main", "--config", "./configs/local.yaml"]

