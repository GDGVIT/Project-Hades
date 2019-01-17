FROM golang

RUN mkdir -p /go/src/angadsharma1016/omega_microservices

ADD . /go/src/angadsharma1016/omega_microservices

WORKDIR /go/src/angadsharma1016/omega_microservices/events

RUN go get -t -v ./...
RUN go get github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

ENTRYPOINT  watcher -run angadsharma1016/omega_microservices/events/cmd  -watch angadsharma1016/omega_microservices/events
