FROM golang:1.23-alpine

WORKDIR /app

COPY ../go.mod ../go.sum ./
RUN go mod download
COPY ./shipping-service ./shipping-service
COPY ./shared ./shared
WORKDIR /app/shipping-service/cmd

RUN go build -o /app/server main.go

EXPOSE 50051

CMD ["/app/server"]