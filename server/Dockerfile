FROM golang:1.13

WORKDIR /app
ADD main.go go.mod /app/
RUN go build -o app

ENTRYPOINT ["./app"]
