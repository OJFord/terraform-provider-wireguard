package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"wireguard_config_document": dataSourceWireguardConfigDocument(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"wireguard_asymmetric_key": resourceWireguardAsymmetricKey(),
		},
	}
}
