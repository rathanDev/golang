FROM golang:1.18-alpine

WORKDIR /app

RUN apk update && apk add libc-dev && apk add gcc && apk add make

COPY ./go.mod ./go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -o app .

EXPOSE 8080

CMD ["./app"]