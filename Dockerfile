FROM golang:1.13.5-alpine AS builder
WORKDIR /app
COPY . .
ENV GOPATH="$PWD"
RUN CGO_ENABLED=0 go build -o grpc-app

FROM scratch
COPY --from=builder /app/grpc-app /app/grpc-app
EXPOSE 3000
ENTRYPOINT ["/app/grpc-app"]