
FROM golang:alpine


ADD . .
RUN go mod init

RUN go build  -o /main .

CMD ["./main"]