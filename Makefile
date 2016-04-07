# Makefile to generate the essential skeleton of a microservice 
SUBDIRS := $(wildcard u*/.)
all: $(SUBDIRS) 
$(SUBDIRS):
	go-bindata templates/... ; go build && cd $@; go generate && go fmt

.PHONY: all $(SUBDIRS)
