FROM golang:alpine3.19

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o myapp

CMD ["go", "run", "."]
