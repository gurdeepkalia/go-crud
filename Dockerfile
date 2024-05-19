# syntax=docker/dockerfile:1

FROM golang:1.21 AS build-stage
# Set destination for COPY
WORKDIR /app
# Download Go modules
COPY go.mod go.sum ./
RUN go mod download
# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
COPY . ./
# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-crud

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...


# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage
WORKDIR /app
COPY .env .
COPY --from=build-stage /go-crud .
# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 8000
USER nonroot:nonroot
ENTRYPOINT ["/app/go-crud"]