# Official Golang image (You shouldn't use the `latest` version in production but I'm a bad guy)
FROM golang:latest

# Working directory
WORKDIR /backend

# Copy everything at /app
COPY . /backend

# Build the go app
RUN go build -o main .

# Expose port
EXPOSE 8080

# Define the command to run the app
CMD ["./"]