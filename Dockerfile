FROM golang:1.23-alpine

WORKDIR /app

COPY . .

RUN go build -o health-check .

EXPOSE 8080

CMD ["./health-check"]
