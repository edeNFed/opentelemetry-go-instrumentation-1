FROM golang:1.21.5-bullseye as builder

WORKDIR /app

RUN apt-get update && apt-get install -y curl clang gcc llvm make libbpf-dev

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading
# them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN make build

FROM gcr.io/distroless/base-debian12@sha256:8a0bb635ccf4cfa88b2bae707c82e633d0867b77827054ed146ad5e13e35ce79
COPY --from=builder /app/otel-go-instrumentation /
CMD ["/otel-go-instrumentation"]
