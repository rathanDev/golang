FROM golang:1.18-alpine

WORKDIR /app

RUN apk update && apk add libc-dev && apk add gcc && apk add make

COPY ./go.mod ./go.sum ./

RUN go mod download && go mod verify

RUN apk update && apk add mysql-client

# RUN go build -o app .

# EXPOSE 8080

# CMD ["./app"]

COPY wait-for-mysql.sh /wait-for-mysql.sh
RUN chmod +x /wait-for-mysql.sh
COPY . .

CMD /wait-for-mysql.sh db-hostname:3306 -- go run main.go
