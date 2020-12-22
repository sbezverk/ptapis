.PHONY: all generate-go 

all: generate-go 

generate-go:
	$(MAKE) -C ./id
	$(MAKE) -C ./registration
	$(MAKE) -C ./configuration
	$(MAKE) -C ./pathtrace
