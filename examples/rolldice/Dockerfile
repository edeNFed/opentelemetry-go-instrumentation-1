FROM golang:1.22.3
WORKDIR /app
COPY ./*.go .
RUN go mod init main
RUN go mod tidy
RUN go build -o main
ENTRYPOINT ["/app/main"]
