FROM golang:1.15.6

WORKDIR /go/src/Cryptopals-API

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 8081

CMD ["src"]
