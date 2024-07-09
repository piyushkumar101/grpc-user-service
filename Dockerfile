FROM golang:1.21.3

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o grpc-user-service .

EXPOSE 50051

CMD ["./grpc-user-service"]
