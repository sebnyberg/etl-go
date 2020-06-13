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