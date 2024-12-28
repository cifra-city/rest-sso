FROM golang:1.23

WORKDIR /app
COPY . .
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz && \
    mv migrate /usr/local/bin/
RUN go mod tidy
RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
