FROM golang:latest

ENV GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64

RUN mkdir /build
WORKDIR /build

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 3000

CMD [ "go-restful-example" ]