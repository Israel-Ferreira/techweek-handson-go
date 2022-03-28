FROM golang:1.18

WORKDIR /go/products

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN go  build -o bin/products

EXPOSE 8082

CMD [ "bin/products" ]