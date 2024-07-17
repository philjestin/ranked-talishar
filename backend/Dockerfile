FROM golang:1.22-alpine3.19 AS builder

# Working directory
WORKDIR /app

# Copy everything at /app
COPY . /app

# Install dependencies
RUN go mod download

# Build the go app
RUN go build -o main .
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/migrate.linux-amd64 ./migrate
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./migration


# Expose port
EXPOSE 8000

# Define the command to run the app
CMD ["/app/main"]
ENTRYPOINT [ "/app/start.sh" ]