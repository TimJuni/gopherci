install:
	go install
	gopherci

test:
	@echo Running unit tests
	go test $(shell go list ./... | grep -v '/vendor/')

test-integration:
	@echo Running integration tests
	go install
	go test -tags integration -v -run TestGCPPubSubQueue ./internal/queue/
	go test -tags integration -v

test-all: test test-integration
