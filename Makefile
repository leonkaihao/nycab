build:
	rm -rf out > /dev/null
	mkdir out
	go build -o out/cli.nycab cmd/cli.nycab/*
	go build -o out/api.nycab cmd/api.nycab/*
	go build -o out/service.nycab cmd/service.nycab/*
tidy:
	go fmt ./...
	go mod tidy

gen-rest:
	cd api/swagger && \
	swagger generate server -f nycab.yaml --principal models.Principal --exclude-main

run-cli:
	go run -race cmd/cli.nycab/* 

run-api:
	go run -race cmd/api.nycab/* --port=10086 

run-service:
	go run -race cmd/service.nycab/* 