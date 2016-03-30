FROM golang:1.6
MAINTAINER Joe Lanford <joe.lanford@gmail.com>

ADD . /go/src/github.com/joelanford/wink_exporter
RUN go install github.com/joelanford/wink_exporter

ENTRYPOINT [ "/go/bin/wink_exporter" ]

EXPOSE 9200