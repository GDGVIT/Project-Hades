.PHONY: compile
compile: 
		mkdir bin
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


.PHONY: build
build:
	docker image build -f ./coupons/Dockerfile_performant -t angadsharma1016/hades-coupons .
	docker image build -t angadsharma1016/hades-event -f ./events/Dockerfile_performant .
	docker image build -t angadsharma1016/hades-participants -f ./participants/Dockerfile_performant .
	docker image build -t angadsharma1016/hades-simple_projection -f ./simple_projection/Dockerfile_performant .
	docker image build -t angadsharma1016/hades-exporter -f ./exporter/Dockerfile_performant .
	docker image build -t angadsharma1016/hades-analytics -f ./analytics/Dockerfile_performant .
	docker image build -t angadsharma1016/hades-guests -f ./guests/Dockerfile_performant .
	docker image build -t angadsharma1016/hades-nginx -f ./Web/Dockerfile .

.PHONY: push
push:
	docker image push angadsharma1016/hades-coupons:latest
	docker image push angadsharma1016/hades-event:latest
	docker image push angadsharma1016/hades-participants:latest
	docker image push angadsharma1016/hades-simple_projection:latest
	docker image push angadsharma1016/hades-exporter:latest
	docker image push angadsharma1016/hades-analytics:latest
	docker image push angadsharma1016/hades-guests:latest
	docker image push angadsharma1016/hades-nginx:latest
