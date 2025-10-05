FROM golang:1.25-alpine

# Install dependencies for Air + Go
RUN apk add --no-cache git curl build-base

# Set working directory
WORKDIR /app

# Install Air (hot reload tool)
RUN go install github.com/air-verse/air@latest
ENV PATH=$PATH:/go/bin

# Copy go.mod and download dependencies first (cache layer)
COPY go.mod ./
RUN go mod download

# Copy the source code
COPY . .

# Expose service port
EXPOSE 8080

# Run Air (auto reloads on file change)
CMD ["air", "-c", ".air.toml"]
