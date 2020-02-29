package wireguard

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

func resourceWireguardAsymmetricKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceWireguardAsymmetricKeyCreate,
		Read:   resourceWireguardAsymmetricKeyRead,
		Delete: resourceWireguardAsymmetricKeyDelete,

		Schema: map[string]*schema.Schema{
			"bind": &schema.Schema{
				Default:  "",
				ForceNew: true,
				Optional: true,
				Type:     schema.TypeString,
			},
			"private_key": &schema.Schema{
				Computed:  true,
				ForceNew:  true,
				Optional:  true,
				Sensitive: true,
				Type:      schema.TypeString,
			},
			"public_key": &schema.Schema{
				Computed: true,
				Type:     schema.TypeString,
			},
		},
	}
}

func resourceWireguardAsymmetricKeyCreate(d *schema.ResourceData, m interface{}) error {
	var key wgtypes.Key
	var err error

	private_key := d.Get("private_key")
	if private_key == "" {
		key, err = wgtypes.GeneratePrivateKey()
		d.Set("private_key", key.String())
	} else {
		key, err = wgtypes.ParseKey(private_key.(string))
	}

	if err != nil {
		return err
	}

	d.Set("public_key", key.PublicKey().String())
	d.SetId(key.PublicKey().String())

	return nil
}

func resourceWireguardAsymmetricKeyRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceWireguardAsymmetricKeyDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
