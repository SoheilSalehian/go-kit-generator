# Makefile to generate the essential skeleton of a microservice 
SUBDIRS := $(wildcard u*/.)
all: $(SUBDIRS) 
$(SUBDIRS):
	/Users/Soheil/Projects/OL/go/bin/go-bindata -o resources.go transport.tmpl && go build && cd $@;rm transport.go; go generate

.PHONY: all $(SUBDIRS)
