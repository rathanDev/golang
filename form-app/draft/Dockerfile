FROM golang:1.18

ENV APP_HOME /go/src/form3-client
RUN mkdir -p "$APP_HOME"
WORKDIR "$APP_HOME"

COPY . .

RUN go mod download

RUN go install -v ./...

EXPOSE 8085

CMD ["go", "test", "-v", "./operation/", "form3-client"]