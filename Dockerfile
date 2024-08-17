# Stage 1: Build our Go App
FROM golang:1.22-alpine AS builder

ENV APP=famchat
ENV CMD_PATH=./cmd/server/main.go

# Set the working directory in the container
WORKDIR /${APP}

# Copy our go.mod and go.sum files, then download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy all the other stuff
COPY . .

# Build it and they will come
RUN go build -o famchat-server ${CMD_PATH} 


#--------------Deploy stage------------#
FROM alpine

WORKDIR /root/

COPY --from=builder /famchat/.env ./

COPY --from=builder /famchat/famchat-server ./

# Expose those ports
EXPOSE 8080

# We cruise'n now
ENTRYPOINT ["./famchat-server"]

# Health check to ensure the service is running
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
  CMD curl -f http://localhost:8080/health_check || exit 1
