FROM golang:1.16
WORKDIR /go/src/app
COPY main.go .
COPY main_test.go .
COPY go.mod .
RUN go build main.go
RUN go test ./...
CMD ["/go/src/app/main"]