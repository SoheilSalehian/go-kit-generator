# Makefile to generate the essential skeleton of a microservice 
all:
	mkdir ${U_SERVICE_DIR}; go-bindata templates/...; go build && cd ${U_SERVICE_DIR}; go generate && go fmt

clean: 
	rm bindata.go
	cd ${U_SERVICE_DIR} && rm *_gen.go 

