FROM golang:1.22.4 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o gateway ./cmd/

FROM alpine:latest
WORKDIR /app
COPY .env .

COPY --from=builder /app/api/casbin/model.conf ./api/casbin/
COPY --from=builder /app/api/casbin/policy.csv ./api/casbin/

EXPOSE 3333

CMD ["./gateway"]