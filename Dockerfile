# Build the application from source
FROM golang:1.23.3 AS build-stage

RUN mkdir /app

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o file-service ./cmd/api/

# Deploy the application binary into a lean image
FROM alpine:latest AS build-release-stage

WORKDIR /

COPY --from=build-stage /app/file-service /file-service

ENTRYPOINT ["/file-service"]
