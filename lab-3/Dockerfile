FROM golang:1.17 AS builder
WORKDIR /app
COPY . .
RUN go build 
CMD ["./char-encoding-lab"]
