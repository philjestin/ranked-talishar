# Official Golang image (You shouldn't use the `latest` version in production but I'm a bad guy)
FROM golang:latest

# Working directory
WORKDIR /app

# Copy everything at /app
COPY . /app

# Install dependencies
RUN go mod download

# Build the go app
RUN go build -o main .

# Expose port
EXPOSE 8000

# Define the command to run the app
CMD ["./main"]