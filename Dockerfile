FROM golang:1.17

WORKDIR /go/src/ecomm

COPY . .

RUN go mod download
RUN GOOS=linux GOARCH=amd64 go build -o main ./cmd/routes/main.go

EXPOSE 80
CMD ["./main"]