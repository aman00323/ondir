export GOPATH=$(PWD)
export GOBIN=$(GOPATH)/bin

all: 
	go get 
	go build
	go install


