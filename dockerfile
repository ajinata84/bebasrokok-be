FROM golang:1.22.2

WORKDIR /app

COPY go.mod go.sum ./
COPY *.go ./

RUN go get


RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

EXPOSE 8080

CMD ["/docker-gs-ping"]