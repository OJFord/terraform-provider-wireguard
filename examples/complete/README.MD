# Example
* `terraform init`
* `terraform plan --var-file=wgvars.tfvars`
* `terraform apply --var-file=wgvars.tfvars`

# Expected results
* Two peer configuration files created in the outputs directory which can then be used with `qrencode -r <FILE> -ansiutf8`
* `terraform.tfstate` is more readable where the resource can easily be identified with description persisted for troubleshooting wireguard issues such as incorrect keys etc

