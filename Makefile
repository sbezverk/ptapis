.PHONY: all generate-go 

all: generate-go 

generate-go:
	$(MAKE) -C ./indexer

