FROM golang:alpine AS build

ADD . /

RUN go build -o /release-version-parser /src/*.go

# ---------------------------------------------------------------------------- #

FROM alpine:latest

RUN \
  mkdir -p /application-binary &&\
  adduser --system --disabled-password --home /release-version-parser release-version-parser

COPY --from=build /release-version-parser .

USER release-version-parser

ENV RELEASE_VERSION_FILE_PATH "productVersions.json"

ENTRYPOINT ["./release-version-parser"]
