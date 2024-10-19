FROM golang:1.23.2-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum main.go ./

RUN go mod tidy

COPY . .

RUN go build -o /app

# EXPOSE 3000

CMD ["air", "./main.go", "-b", "0.0.0.0"]
