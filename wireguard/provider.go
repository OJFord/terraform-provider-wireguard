package wireguard

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"wireguard_asymmetric_key": resourceWireguardAsymmetricKey(),
		},
	}
}
