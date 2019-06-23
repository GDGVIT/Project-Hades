.PHONY: build
build: 
		GO111MODULE=on GOOS=linux CGO_ENABLED=0 go build -v -a -installsuffix cgo -o ./bin/analytics ./analytics/cmd/main.go; \
		CGO_ENABLED=0 go build -v -a -installsuffix cgo -o ./bin/organization ./organization/cmd/main.go; \
		GO111MODULE=on GOOS=linux CGO_ENABLED=0 go build -v -a -installsuffix cgo -o ./bin/coupons ./coupons/cmd/main.go; \
		CGO_ENABLED=0 go build -v -a -installsuffix cgo -o ./bin/events ./events/cmd/main.go; \
		GO111MODULE=on GOOS=linux CGO_ENABLED=0 go build -v -a -installsuffix cgo -o ./bin/exporter ./exporter/cmd/main.go; \
		GO111MODULE=on GOOS=linux CGO_ENABLED=0 go build -v -a -installsuffix cgo -o ./bin/guests ./guests/cmd/main.go; \
		GO111MODULE=on GOOS=linux CGO_ENABLED=0 go build -v -a -installsuffix cgo -o ./bin/simple_projection ./simple_projection/cmd/main.go; \
		GO111MODULE=on GOOS=linux CGO_ENABLED=0 go build -v -a -installsuffix cgo -o ./bin/participants ./participants/cmd/main.go 

.PHONY: docs
docs:  
		apidoc -i analytics \
			 -i organization \
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

