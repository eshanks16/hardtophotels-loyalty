FROM golang:latest

LABEL maintainer="Eric Shanks <eshanks@vmware.com>"

RUN mkdir /htweb

ADD . /htweb

WORKDIR /htweb

RUN go get . && go build -o main .

EXPOSE 8080

CMD ["./main"]