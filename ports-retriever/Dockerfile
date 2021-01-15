FROM golang:alpine AS build

RUN \
  apk add --no-cache git &&\
  go get gopkg.in/yaml.v2

ADD . /

RUN go build -o /ports-retriever /src/*.go

# ---------------------------------------------------------------------------- #

FROM alpine:latest

RUN \
  mkdir -p /application-binary &&\
  adduser --system --disabled-password --home /ports-retriever ports-retriever

COPY --from=build /ports-retriever .
COPY --from=build /ports.yml .

USER ports-retriever

ENV PORTS_RETRIEVER_FILE_PATH "./ports.yml"

ENTRYPOINT ["./ports-retriever"]
