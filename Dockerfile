# Build stage: use builder variant to install dependencies and compile
FROM quay.io/hummingbird/go:latest AS builder
RUN mkdir /src
COPY main.go /src
WORKDIR /src
RUN go build -o /app main.go

# Runtime stage: use minimal base image for the compiled binary
FROM quay.io/hummingbird/core-runtime:latest
COPY --from=builder /app /app
ENTRYPOINT ["/app"]
