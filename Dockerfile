# syntax=docker/dockerfile:1

FROM golang:1.22

# Set destination for COPY
WORKDIR /app

ENV POSTGRES_PASSWORD=^+h9Cd~D/8JAHOB7

#COPY wait-for-postgres.sh ./
# Install psql
RUN apt-get update && apt-get -y install postgresql-client

# Make wait-for-postgres.sh executable
#RUN chmod +x wait-for-postgres.sh

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN chmod +x wait-for-postgres.sh

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /portmonetka.authorization

EXPOSE 8080

# Run
CMD ["/portmonetka.authorization"]