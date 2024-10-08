FROM golang:1.23-alpine

WORKDIR /app

COPY ../go.mod go.sum ./
RUN go mod download
COPY .. .

RUN go build -o server cmd/shippings/server/main.go cmd/shippings/server/route.go

EXPOSE 50051

CMD ["./server"]