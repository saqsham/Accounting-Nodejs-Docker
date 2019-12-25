FROM golang:1.13.5

# WORKDIR /go/src/app
# COPY . .
EXPOSE 3000

RUN go get -v github.com/go-sql-driver/mysql
# no -d flag
# RUN go install -v ./...

# CMD ["app"]
# same as
# CMD ["go", "run", "server.go"]

