.ONESHELL:

default: build

build: fmt
	go build -o terraform-provider-wireguard

fmt:
	go fmt

terraformrc:
	cat <<-EOC > $@
		provider_installation {
		  dev_overrides {
		    "OJFord/wireguard" = "$$PWD"
		  }
		  direct {}
		}
	EOC

examples: build terraformrc
	shopt -s globstar
	for d in examples/**/versions.tf; do
		d="$$(dirname $$d)"
		echo "Applying example $$d"
		TF_CLI_CONFIG_FILE="$$PWD/terraformrc" terraform -chdir="$$d" init
		TF_CLI_CONFIG_FILE="$$PWD/terraformrc" terraform -chdir="$$d" apply -auto-approve
	done
