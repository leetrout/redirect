FROM golang:alpine

COPY redirect.go .

RUN go build redirect.go

CMD ./redirect
