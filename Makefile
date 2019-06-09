.PHONY: build
build: 
		GO111MODULE=on GOOS=linux GOARCH=amd64 CGOENABLED=0 \
		go build -ldflags="-w -s" -o ./bin/analytics ./analytics/cmd/main.go &&\
		go build -ldflags="-w -s" -o ./bin/auth ./auth/cmd/main.go  &&\
		go build -ldflags="-w -s" -o ./bin/coupons ./coupons/cmd/main.go &&\
		go build -ldflags="-w -s" -o ./bin/events ./events/cmd/main.go  &&\
		go build -ldflags="-w -s" -o ./bin/exporter ./exporter/cmd/main.go  &&\
		go build -ldflags="-w -s" -o ./bin/guests ./guests/cmd/main.go  &&\
		go build -ldflags="-w -s" -o ./bin/simple_projection ./simple_projection/cmd/main.go &&\
		go build -ldflags="-w -s" -o ./bin/participants ./participants/cmd/main.go 

.PHONY: docs
docs:  
		apidoc -i analytics \
			 -i auth \
			 -i events \
			 -i exporter \
			 -i guests \
			 -i participants \
			 -i simple_projection \
			 -i events \
			 -i coupons \
			 -o docs

.PHONY: lint
lint:
		eslint ./docs/*/*.js --fix

.PHONY: test
test:
	go test -v ./test/...

.PHONY: exec
exec:
		./exec

.PHONY: dep
dep:
	GO111MODULE=on
	go mod vendor
	go mod verify

