FROM golang:1.21.5

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project to the working directory
COPY . .

# Install the package
RUN go install ./...

# Run bash
CMD ["/bin/bash"]
