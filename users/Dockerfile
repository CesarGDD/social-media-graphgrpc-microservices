FROM golang:alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

EXPOSE 50051


CMD ["go", "run", "main.go"]