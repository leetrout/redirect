FROM golang:alpine

RUN adduser -D -u 1000 foo

RUN mkdir /app

RUN chown foo /app

WORKDIR /app

USER foo

COPY redirect.go .

RUN go build redirect.go

CMD ./redirect
