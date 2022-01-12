package provider

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
	"io"
)

func resourceWireguardPresharedKey() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a WireGuard asymmetric key resource. This can be used to create, read, and delete WireGuard keys in terraform state.",

		Create: resourceWireguardPresharedKeyCreate,
		Read:   resourceWireguardPresharedKeyRead,
		Delete: resourceWireguardPresharedKeyDelete,

		Schema: map[string]*schema.Schema{
			"preshared_key": {
				Description: "Additional layer of symmetric-key cryptography to be mixed into the already existing public-key cryptography, for post-quantum resistance.",
				Computed:    true,
				Type:        schema.TypeString,
			},
		},
	}
}

func resourceWireguardPresharedKeyCreate(d *schema.ResourceData, m interface{}) error {
	var key wgtypes.Key
	var err error

	key, err = wgtypes.GenerateKey()
	err = d.Set("preshared_key", key.String())
	if err != nil {
		return err
	}
	h := md5.New()
	_, err = io.WriteString(h, key.String())
	if err != nil {
		return err
	}
	d.SetId(hex.EncodeToString(h.Sum(nil)))

	return nil
}

func resourceWireguardPresharedKeyRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceWireguardPresharedKeyDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
