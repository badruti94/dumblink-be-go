FROM golang:1.18-alpine as builder
WORKDIR /app/
COPY . .
RUN go build -o main main.go

FROM alpine:3
WORKDIR /app/
COPY --from=builder /app/main /app/
EXPOSE 8080
CMD ./main