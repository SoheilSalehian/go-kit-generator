all:
	mkdir ${U_SERVICE_DIR}; ./configure ${CONFIG_FILE} $(U_SERVICE_DIR)
	cd $(U_SERVICE_DIR); go generate && gofmt

clean:
	cd ${U_SERVICE_DIR}; rm *_gen.go; rm gen.go
	
	
