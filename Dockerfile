FROM golang:1.23.3

WORKDIR /app

COPY .env .env
COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o /receipt-processor ./cmd

EXPOSE 8080

CMD [ "/receipt-processor" ]
