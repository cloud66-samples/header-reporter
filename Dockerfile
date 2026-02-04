# build stage - compile the go binary
FROM golang:1.25-alpine AS builder

WORKDIR /app

# copy module files first for better layer caching
COPY go.mod ./
RUN go mod download

# copy source and build
COPY main.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o header-reporter .

# runtime stage - minimal image
FROM alpine:latest

WORKDIR /app

# copy the binary from builder
COPY --from=builder /app/header-reporter .

# expose port 80
EXPOSE 80

# run the application
CMD ["./header-reporter"]
