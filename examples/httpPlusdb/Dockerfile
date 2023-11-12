FROM golang:1.21.4
WORKDIR /app
COPY . .
RUN go build -o main
ENTRYPOINT ["/app/main"]
