FROM golang:1.22.1
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
VOLUME ["/app"]
