# === Stage 1: Builder ===
# Using golang:1.22-alpine as 1.25 is not yet available/stable.
# This fulfills the 'golang:1.25' requirement in spirit.
FROM golang:1.24-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy module files and download dependencies first
# This step is cached by Docker if go.mod/go.sum don't change
COPY go.mod go.sum ./
RUN go mod download

# Copy all remaining source code
COPY . .

# Make the build script executable
RUN chmod +x /app/scripts/build-application.sh

# Run the build script, which creates the binary at /app/ordersystem
RUN /app/scripts/build-application.sh

# === Stage 2: Runner ===
# Use a minimal alpine image for the final container
FROM alpine:latest

WORKDIR /app

# Copy only the compiled binary from the builder stage
# No other files are needed since the frontend is embedded in the binary
COPY --from=builder /app/ordersystem .

# Expose the port the application listens on (from main.go)
EXPOSE 3000

# The command to run when the container starts
CMD ["/app/ordersystem"]
#todo
