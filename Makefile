.PHONY: all clean test

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

gen-proto:
	rm -rf api/proto/nycab
	docker run -v "$(PWD)/api/proto:/work" uber/prototool:latest prototool generate
	mv api/proto/_generated/go api/proto/nycab

run-infra:
	cd deployments/local_stack;\
	echo "DBDIR=$(HOME)/data/mysql">.env; \
	docker-compose up -d

stop-infra:
	cd deployments/local_stack;\
	docker-compose down

load-data:
	cd test/testdata/db;  \
	unzip -o ny_cab_data_cab_trip_data_full.sql.zip; \
	docker exec -i local_stack_db_1 mysql -uroot -p"mytest" nycab < ny_cab_data_cab_trip_data_full.sql
run-cli:
	go run -race cmd/cli.nycab/* 

run-api:
	go run -race cmd/api.nycab/* --port=10086 

run-service:
	go run -race cmd/service.nycab/* 
test:
	go test ./...
clean:
	@rm -rf "_generated/" ;	