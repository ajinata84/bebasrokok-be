FROM golang:1.22.2

WORKDIR /app

COPY go.mod go.sum ./

# Download and cache dependencies
RUN go mod download

COPY . . 

# Verify module and dependencies
RUN go mod tidy

# Debugging step: List files to ensure all required files are present
RUN ls -la

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping .

EXPOSE 8080

CMD ["/docker-gs-ping"]
