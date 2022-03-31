FROM golang:1.18

WORKDIR /go/stocks

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN go  build -o bin/stocks

EXPOSE 8083

CMD [ "bin/stocks" ]