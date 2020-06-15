FROM golang:1.14.4-buster


COPY . /src/app

WORKDIR /src/app

RUN go mod vendor && go build && cp docker-events /go/bin/


CMD ["/go/bin/docker-events"]
