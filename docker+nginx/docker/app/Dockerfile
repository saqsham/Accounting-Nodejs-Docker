FROM golang:1.13.5

WORKDIR /go/src/app

COPY ../app-main .

RUN go get -v github.com/go-sql-driver/mysql

# EXPOSE 3000

CMD ["go", "run", "server.go"]
