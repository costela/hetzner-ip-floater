FROM golang:1.12-alpine AS build

RUN apk add --update git

WORKDIR /app

COPY . .

RUN go build


FROM alpine

RUN apk add --no-cache ca-certificates

COPY --from=build /app/hetzner-ip-floater /

CMD [ "/hetzner-ip-floater" ]