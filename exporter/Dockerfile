FROM golang

RUN mkdir -p /go/src/github.com/GDGVIT/Project-Hades

ADD . /go/src/github.com/GDGVIT/Project-Hades

WORKDIR /go/src/github.com/GDGVIT/Project-Hades/exporter

RUN go get -t -v ./...

RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

EXPOSE 8084

ENTRYPOINT  watcher -run github.com/GDGVIT/Project-Hades/exporter/cmd  -watch github.com/GDGVIT/Project-Hades/exporter
