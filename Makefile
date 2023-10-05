.PHONY: avro

avro:
	rm -rf avro
	mkdir avro
	$$GOPATH/bin/gogen-avro \
		--containers \
		--short-unions \
		--package avro \
		avro/ purchase.avsc 

install:
	go get github.com/actgardner/gogen-avro

memprof:
	go tool pprof -web mem.pprof

cpuprof:
	go tool pprof -web cpu.pprof

build:
	go build -o main main.go

run: build
	./main gen -n 1000000