FROM golang:1.24-alpine AS builder
WORKDIR /src

COPY server/main.go ./main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags="-s -w" -o /out/app ./main.go

FROM gcr.io/distroless/static-debian12:nonroot
WORKDIR /app
COPY --from=builder /out/app /app/app

EXPOSE 1337
ENTRYPOINT ["/app/app"]
