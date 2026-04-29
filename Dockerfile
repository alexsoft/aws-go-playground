FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server .

FROM scratch
LABEL org.opencontainers.image.source=https://github.com/alexsoft/aws-go-playground
COPY --from=builder /app/server /server
EXPOSE 8080
ENTRYPOINT ["/server"]
