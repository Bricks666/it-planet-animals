FROM golang:alpine3.17 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o app .
EXPOSE 5000
CMD ["./app"]
