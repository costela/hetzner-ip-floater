FROM golang:1.10-alpine AS build

RUN apk add --update git

WORKDIR /go/src/app

RUN go get -d github.com/hetznercloud/hcloud-go/hcloud

COPY . .

RUN go get -d -v ./...
RUN go build -v ./...


FROM alpine

RUN apk add --no-cache ca-certificates

WORKDIR /app
COPY --from=build /go/src/app/app /app

CMD [ "./app" ]