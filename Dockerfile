FROM golang:latest

WORKDIR /app

COPY . /app

RUN go mod tidy
RUN go build -o myapp

EXPOSE 8080

CMD ["./myapp"]
