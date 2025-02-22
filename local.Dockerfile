FROM golang:1.24-alpine

WORKDIR /app

COPY . .


RUN go mod download
RUN go build -o /app/main cmd/main.go

RUN ls -la /app

EXPOSE 8080

CMD ["/app/main"]