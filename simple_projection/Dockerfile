FROM golang

RUN mkdir -p /go/src/github.com/GDGVIT/Project-Hades

ADD . /go/src/github.com/GDGVIT/Project-Hades
WORKDIR /go/src/github.com/GDGVIT/Project-Hades/simple_projection
RUN go get  -t -v ./...
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher
EXPOSE 8083
ENTRYPOINT  watcher -run github.com/GDGVIT/Project-Hades/simple_projection/cmd  -watch github.com/GDGVIT/Project-Hades/simple_projection
