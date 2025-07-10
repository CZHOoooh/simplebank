#Build stage
FROM golang:1.23.10-alpine3.22 AS builder
WORKDIR /app
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.io,direct
COPY . .
RUN go build -o main main.go

# Run stage
FROM alpine:3.22
WORKDIR /app
COPY --from=builder /app/main .
COPY migrate .
COPY app.env .
COPY start.sh .
COPY db/migration ./migration

EXPOSE 8080
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]