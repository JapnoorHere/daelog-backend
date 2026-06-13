FROM --platform=$BUILDPLATFORM golang:1.26-alpine AS builder
ARG TARGETARCH
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN GOARCH=${TARGETARCH} go build -o daelog-backend ./cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/daelog-backend .
EXPOSE 8080
CMD ["./daelog-backend"]
