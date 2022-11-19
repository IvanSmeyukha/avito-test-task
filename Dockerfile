FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o avito-test-task ./main.go

CMD ["./avito-test-task"]