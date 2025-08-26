FROM golang:latest AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /url-shortener ./cmd/url-shortener

FROM alpine:latest
WORKDIR /root/

COPY --from=build /url-shortener .

COPY --from=build /app/config ./config

ENV CONFIG_PATH=/root/config/local.yaml
CMD ["./url-shortener"]
