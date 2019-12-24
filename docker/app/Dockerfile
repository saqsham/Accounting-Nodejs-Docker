FROM golang:1.13.5

# COPY ../wait-for-it.sh /go/src/app

# WORKDIR /go/src/app
# COPY . .
EXPOSE 3000


# RUN apt-get update && apt-get install -y mysql-client
RUN go get -v github.com/go-sql-driver/mysql
# no -d flag
# RUN go install -v ./...

# RUN chmod +x wait-for-it.sh

# RUN mv wait-for-it.sh /

# CMD [ "sh", "-c", "./wait-fot-it", "mysql:3306", "--", "go", "run", "server.go" ]

# CMD ["app"]
# same as
# CMD ["go", "run", "server.go"]

