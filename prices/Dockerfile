FROM golang:1.18

WORKDIR /go/prices

COPY go.mod .

COPY go.sum .

RUN go clean --modcache

RUN go mod download

COPY . .

RUN go  build -o bin/prices

EXPOSE 8084

CMD [ "bin/prices" ]