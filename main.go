package main

import (
	"github.com/OJFord/terraform-provider-wireguard/wireguard"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return wireguard.Provider()
		},
	})
}
