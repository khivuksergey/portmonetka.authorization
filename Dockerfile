# syntax=docker/dockerfile:1

FROM golang:1.22

# Set destination for COPY
WORKDIR /app

COPY go.mod go.sum wait-for-postgres.sh ./

# Install psql
RUN apt-get update
RUN apt-get -y install postgresql-client

# Make wait-for-postgres.sh executable
RUN chmod +x wait-for-postgres.sh

# Download Go modules
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY *.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /portmonetka.authorization

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 8080

# Run
CMD ["/portmonetka.authorization"]