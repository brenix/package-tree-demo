FROM golang:1.12 AS builder

# Pull dependencies
WORKDIR /src

# Build
COPY . .
RUN GOARCH=amd64 CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o package-tree .

# Copy binary to scratch container
FROM scratch
COPY --from=builder /src/package-tree /
EXPOSE 8080
ENTRYPOINT ["/package-tree"]