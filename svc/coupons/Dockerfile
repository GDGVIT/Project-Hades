FROM golang

RUN mkdir -p /go/src/github.com/GDGVIT/Project-Hades

ADD . /go/src/github.com/GDGVIT/Project-Hades
WORKDIR /go/src/github.com/GDGVIT/Project-Hades/coupons
#RUN go get  -t -v ./...
RUN GO111MODULE=on go get  github.com/canthefason/go-watcher
RUN GO111MODULE=on go install github.com/canthefason/go-watcher/cmd/watcher
EXPOSE 8082
ENTRYPOINT  watcher -run github.com/GDGVIT/Project-Hades/coupons/cmd  -watch github.com/GDGVIT/Project-Hades/coupons
