FROM golang:1.23-alpine

WORKDIR /app

# Install git for buf and mockery support
RUN apk add --no-cache git curl

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o report-service ./backend/main.go

EXPOSE 8080

CMD ["./report-service"]
