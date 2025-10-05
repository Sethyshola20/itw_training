FROM golang:1.25-alpine

# Install tools for Air + build
RUN apk add --no-cache git curl build-base

# Set working directory
WORKDIR /app

# Install Air for hot reload
RUN go install github.com/air-verse/air@latest
ENV PATH=$PATH:/go/bin

# Copy go.mod first and download deps (cache)
COPY go.mod ./
RUN go mod download

# Copy the rest of your code
COPY . .

# Expose the billing service port
EXPOSE 8002

# Run Air (auto reload on file change)
CMD ["air", "-c", ".air.toml"]
