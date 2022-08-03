FROM golang:1.16.7

LABEL maintainer="Eric Shanks <eshanks@vmware.com>"

RUN mkdir /loyalty

ADD . /loyalty

WORKDIR /loyalty

COPY ./rds-combined-ca-bundle.pem /etc/ssl/certs/rds-combined-ca-bundle.pem

RUN go get . && go build -o main .

EXPOSE 8080

CMD ["./main"]