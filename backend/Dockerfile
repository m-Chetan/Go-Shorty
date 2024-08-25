FROM golang:1.22.5-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o go-shorty .

EXPOSE 8080

CMD ["./go-shorty"]