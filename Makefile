.PHONY: all generate-go 

all: generate-go 

generate-go:
	$(MAKE) -C ./id

