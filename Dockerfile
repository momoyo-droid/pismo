FROM golang:1.26.1

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app ./api/cmd/app
# Fix permission denied issue
RUN chmod +x app

EXPOSE 3000

CMD ["./app"]