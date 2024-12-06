# Tahap build
FROM golang:1.23.0-alpine3.20 AS builder

# Set GOPRIVATE untuk repositori privat 
ENV go env -w GOPRIVATE=github.com/AsrofunNiam/*

# Install git
RUN apk add --no-cache git

# Build argument for GitHub token
ARG GITHUB_TOKEN
ARG GITHUB_USERNAME

# Ensure git uses the GitHub token for private repo authentication 

RUN git config --global url."https://${GITHUB_USERNAME}:${GITHUB_TOKEN}@github.com/".insteadOf "https://github.com/"

WORKDIR /app

# Copy the go mod and sum files
COPY go.mod go.sum ./

# Pass the GitHub token as an environment variable for go mod download
ENV GITHUB_TOKEN=${GITHUB_TOKEN}

# Download Go modules, termasuk dari repositori privat
RUN --mount=type=secret,id=github_token go mod download

# Copy the project files
COPY . .

# Build the Go binary
RUN go build -o main .

# Tahap runtime
FROM alpine:3.20

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/main .

# Copy .env file if your application requires it
COPY ./configuration/.env ./configuration/.env

EXPOSE 8080

CMD ["/app/main"]
