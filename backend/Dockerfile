FROM golang:1.22-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build Arguments
ARG GIT_SHA
ARG VERSION
ARG BUILD_TIME

# Inject Metadata
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-w -s -X 'main.GitSHA=${GIT_SHA}' -X 'main.Version=${VERSION}' -X 'main.BuildTime=${BUILD_TIME}'" \
    -o /app/api ./cmd/api

FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /app/api .
USER 65532:65532

EXPOSE 8080
ENTRYPOINT ["/api"]
