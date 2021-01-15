FROM golang:alpine AS build

RUN \
  apk add --no-cache git &&\
  mkdir /application

RUN \
  go get github.com/gregjones/httpcache/diskcache &&\
  go get github.com/xanzy/go-gitlab

ADD . /application/

RUN go build -o /delivery /application/src/github.com/gpenaud/delivery/*.go

# ---------------------------------------------------------------------------- #

FROM alpine:latest

RUN adduser --system --disabled-password --home /delivery delivery

WORKDIR /delivery
# USER delivery

COPY --from=build /delivery .
COPY --from=build /application/src/github.com/gpenaud/delivery/templates templates

ENV TEMPLATES_DIR /delivery/templates

EXPOSE 8080

CMD ["./delivery"]
