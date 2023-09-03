FROM golang:1.18.8
ENV APP_HOME /go/src/toxic-rest-api
WORKDIR "$APP_HOME"
COPY app/ .
RUN go mod download
RUN go build -v ./cmd/toxic-rest-api.go
EXPOSE 8080
CMD ["./toxic-rest-api"]
