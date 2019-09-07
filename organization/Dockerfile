FROM golang

RUN mkdir -p /go/src/github.com/GDGVIT/Project-Hades

ADD . /go/src/github.com/GDGVIT/Project-Hades

WORKDIR /go/src/github.com/GDGVIT/Project-Hades/organization

RUN GO111MODULE=on go get  github.com/canthefason/go-watcher

RUN GO111MODULE=on go install github.com/canthefason/go-watcher/cmd/watcher

EXPOSE 8087

ENTRYPOINT  watcher -run github.com/GDGVIT/Project-Hades/organization/cmd  -watch github.com/GDGVIT/Project-Hades/organization


