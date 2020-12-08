TAG = 0.0.11

all: start

start:
	go run ./... > apiserver.log 2>&1 &

build-local:
	go build -o apiserver && chmod +x apiserver

build-arm:
	GOOS=linux GOARCH=arm64 go build -o apiserver && chmod +x apiserver

upload:
	docker build -t jensdriller/istio-test-server:v${TAG} .
	docker push jensdriller/istio-test-server:v${TAG}

.PHONY: start build