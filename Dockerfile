FROM golang:1.10-alpine

RUN apk add --update git

WORKDIR /go/src/app

RUN go get -d github.com/hetznercloud/hcloud-go/hcloud

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]