FROM golang:1.22.2

WORKDIR /app

COPY go.mod go.sum ./

# Download and cache dependencies
RUN go mod download

COPY *.go ./

# Verify module and dependencies
RUN go mod tidy

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

EXPOSE 8080

CMD ["/docker-gs-ping"]
