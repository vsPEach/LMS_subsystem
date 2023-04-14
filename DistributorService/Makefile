
TAG := "gateway:0.0.1"

build-img:
	docker build -t $(TAG) .
run-img: build-img
	docker run --name app -p 5005:5005 $(TAG)
build:
	go build cmd/main.go
run: build
	./main