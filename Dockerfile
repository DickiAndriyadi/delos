FROM golang:latest

WORKDIR /app/delos

COPY go.mod .

COPY go.sum .

COPY .env .

RUN go mod download

COPY . .

# ENV PORT 8778

EXPOSE 8080

RUN go build

CMD ["./delos"]
