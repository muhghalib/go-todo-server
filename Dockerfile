FROM golang:1.23.3-alpine AS deps

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

FROM deps AS builder

WORKDIR /app

COPY --from=deps /app /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main ./cmd/web

FROM builder as runner

WORKDIR /app

ENV APP_ENV production

COPY --from=builder /app /app

EXPOSE ${SERVER_PORT}

CMD ["/app/main"]



