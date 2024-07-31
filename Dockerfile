FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o main .

FROM scratch

COPY --from=builder /app/main /app/main

EXPOSE 3000
CMD ["./app/main"]
