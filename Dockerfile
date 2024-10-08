# syntax=docker/dockerfile:1

FROM golang:1.18

# Set destination for COPY
WORKDIR /forum

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
COPY . ./
RUN ls -la /forum

# Build
RUN CGO_ENABLED=1 GOOS=linux go build

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 2345

# Run
CMD ["./forum_reboot"]