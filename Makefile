EXAMPLES := $(wildcard examples/*/*)

default: build

build: fmt
	go build -o terraform-provider-wireguard

fmt:
	go fmt

examples: build $(EXAMPLES)
	for d in $(EXAMPLES); do \
		echo "Applying example $$d" ;\
		terraform -chdir="$$d" init;\
		terraform -chdir="$$d" apply -auto-approve;\
	done
