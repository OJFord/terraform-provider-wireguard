package provider

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

func resourceWireguardPresharedKey() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a WireGuard key resource. This can be used to create, read, and delete WireGuard preshared keys in terraform state.",

		Create: resourceWireguardPresharedKeyCreate,
		Read:   resourceWireguardPresharedKeyRead,
		Delete: resourceWireguardPresharedKeyDelete,

		Schema: map[string]*schema.Schema{
			"key": {
				Description: "Additional layer of symmetric-key cryptography to be mixed into the already existing public-key cryptography, for post-quantum resistance.",
				Computed:    true,
				Sensitive:   true,
				Type:        schema.TypeString,
			},
		},
	}
}

func resourceWireguardPresharedKeyCreate(d *schema.ResourceData, m interface{}) error {
	var key wgtypes.Key
	var err error

	key, err = wgtypes.GenerateKey()
	err = d.Set("key", key.String())
	if err != nil {
		return err
	}
	hash := sha256.Sum256([]byte(key.String()))
	d.SetId(hex.EncodeToString(hash[:]))

	return nil
}

func resourceWireguardPresharedKeyRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceWireguardPresharedKeyDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
