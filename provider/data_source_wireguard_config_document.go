package provider

import (
	"bytes"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
	"text/template"
)

func dataSourceWireguardConfigDocument() *schema.Resource {
	return &schema.Resource{
		Description: "A WireGuard configuration document.",
		Schema: map[string]*schema.Schema{
			"conf": {
				Description: "The rendered config document.",
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
			},

			"private_key": {
				Description: "The base64 private key for this peer's interface.",
				Type:        schema.TypeString,
				Required:    true,
			},

			"listen_port": {
				Description: ".",
				Type:        schema.TypeInt,
				Optional:    true,
			},

			"firewall_mark": {
				Description: "A 32-bit fwmark for outgoing packets.",
				Type:        schema.TypeString,
				Optional:    true,
			},

			"addresses": {
				Description: "IPs (or CIDR) to be assigned to the interface. (`wg-quick`/apps only.)",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"dns": {
				Description: "IPs or hostnames of Domain Name Servers to set as the interface's DNS search domains. (`wg-quick`/apps only.)",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"mtu": {
				Description: "Manual MTU to override automatic discovery. (`wg-quick`/apps only.)",
				Type:        schema.TypeInt,
				Optional:    true,
			},

			"routing_table": {
				Description: "Controls the routing table (or \"off\") to which routes are added. (`wg-quick`/apps only.)",
				Type:        schema.TypeString,
				Optional:    true,
			},

			"pre_up": {
				Description: "Script to run before setting up the interface. (`wg-quick`/apps only.)",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"post_up": {
				Description: "Scripts to run after setting up the interface. (`wg-quick`/apps only.)",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"pre_down": {
				Description: "Scripts to run before tearing down the interface. (`wg-quick`/apps only.)",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"post_down": {
				Description: "Scripts to run before tearing down the interface. (`wg-quick`/apps only.)",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"peer": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"public_key": {
							Description: "The base64 public key for this peer.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"preshared_key": {
							Description: "A base64 preshared key from this peer.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"allowed_ips": {
							Description: "IPs (or CIDR) allowed for traffic to/from this peer.",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"endpoint": {
							Description: "An endpoint IP:port or hostname:port at which this peer can be reached initially.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"persistent_keepalive": {
							Description: "Period in seconds (or \"off\") after which to ping the peer to keep a stateful firewall or NAT mapping valid.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
					},
				},
			},
		},

		Read: dataSourceWireguardConfigDocumentRead,
	}
}

const wgTemplateStr = `
[Interface]
PrivateKey = {{ .PrivateKey }}

{{- if .ListenPort }}
ListenPort = {{ .ListenPort }}
{{- end }}

{{- if .FirewallMark }}
FwMark = {{ .FirewallMark }}
{{- end }}

{{- range .Addresses }}
Address = {{ . }}
{{- end }}

{{- if .DNS }}
DNS = {{ range $i, $d := .DNS }}{{ if $i }},{{ end }}{{ $d }}{{ end }}
{{- end }}

{{- if .MTU }}
MTU = {{ .MTU }}
{{- end }}

{{- if .RoutingTable }}
Table = {{ .RoutingTable }}
{{- end }}

{{- range .PreUp }}
PreUp = {{ . }}
{{- end }}

{{- range .PostUp }}
PostUp = {{ . }}
{{- end }}

{{- range .PreDown }}
PreDown = {{ . }}
{{- end }}

{{- range .PostDown }}
PostDown = {{ . }}
{{- end }}

{{- range .Peers }}

[Peer]
PublicKey = {{ .PublicKey }}

{{- if .PresharedKey }}
PresharedKey = {{ .PresharedKey }}
{{- end }}

{{- range .AllowedIPs }}
AllowedIPs = {{ . }}
{{- end }}

{{- if .Endpoint }}
Endpoint = {{ .Endpoint }}
{{- end }}

{{- if .PersistentKeepalive }}
PersistentKeepalive = {{ .PersistentKeepalive }}
{{- end }}

{{- end }}{{/* peers */}}
`

var wgTemplate = template.Must(template.New("wg").Parse(wgTemplateStr))

type WgPeerConfig struct {
	PublicKey           string
	PresharedKey        *string
	AllowedIPs          []string
	Endpoint            *string
	PersistentKeepalive *int
}

type WgQuickConfig struct {
	Addresses    []string
	DNS          []string
	MTU          *int
	RoutingTable *string
	PreUp        []string
	PostUp       []string
	PreDown      []string
	PostDown     []string
}

type WgConfig struct {
	PrivateKey   string
	ListenPort   *int
	FirewallMark *string
	WgQuickConfig
	Peers []WgPeerConfig
}

func dataSourceWireguardConfigDocumentRead(d *schema.ResourceData, m interface{}) error {
	cfg := WgConfig{
		PrivateKey: d.Get("private_key").(string),
	}

	if v, set := d.GetOk("listen_port"); set {
		p := v.(int)
		cfg.ListenPort = &p
	}

	if v, set := d.GetOk("firewall_mark"); set {
		fw := v.(string)
		cfg.FirewallMark = &fw
	}

	if vals := d.Get("addresses").(*schema.Set).List(); len(vals) > 0 {
		cfg.Addresses = make([]string, len(vals))
		for i, v := range vals {
			cfg.Addresses[i] = v.(string)
		}
	}

	if vals := d.Get("dns").(*schema.Set).List(); len(vals) > 0 {
		cfg.DNS = make([]string, len(vals))
		for i, v := range vals {
			cfg.DNS[i] = v.(string)
		}
	}

	if v, set := d.GetOk("mtu"); set {
		mtu := v.(int)
		cfg.MTU = &mtu
	}

	if v, set := d.GetOk("routing_table"); set {
		rt := v.(string)
		cfg.RoutingTable = &rt
	}

	if vals := d.Get("pre_up").([]interface{}); len(vals) > 0 {
		cfg.PreUp = make([]string, len(vals))
		for i, v := range vals {
			cfg.PreUp[i] = v.(string)
		}
	}

	if vals := d.Get("post_up").([]interface{}); len(vals) > 0 {
		cfg.PostUp = make([]string, len(vals))
		for i, v := range vals {
			cfg.PostUp[i] = v.(string)
		}
	}

	if vals := d.Get("pre_down").([]interface{}); len(vals) > 0 {
		cfg.PreDown = make([]string, len(vals))
		for i, v := range vals {
			cfg.PreDown[i] = v.(string)
		}
	}

	if vals := d.Get("post_down").([]interface{}); len(vals) > 0 {
		cfg.PostDown = make([]string, len(vals))
		for i, v := range vals {
			cfg.PostDown[i] = v.(string)
		}
	}

	if v, set := d.GetOk("peer"); set {
		peers := v.([]interface{})

		for _, peer := range peers {
			peer := peer.(map[string]interface{})

			peerCfg := WgPeerConfig{
				PublicKey: peer["public_key"].(string),
			}

			if v := peer["preshared_key"]; v != "" {
				psk := v.(string)
				peerCfg.PresharedKey = &psk
			}

			if vals := peer["allowed_ips"].([]interface{}); len(vals) > 0 {
				peerCfg.AllowedIPs = make([]string, len(vals))
				for i, v := range vals {
					peerCfg.AllowedIPs[i] = v.(string)
				}
			}

			if v := peer["endpoint"]; v != "" {
				ep := v.(string)
				peerCfg.Endpoint = &ep
			}

			if v := peer["persistent_keepalive"]; v != 0 {
				ka := v.(int)
				peerCfg.PersistentKeepalive = &ka
			}

			cfg.Peers = append(cfg.Peers, peerCfg)
		}
	}

	var buf bytes.Buffer
	if err := wgTemplate.Execute(&buf, cfg); err != nil {
		return err
	}

	key, err := wgtypes.ParseKey(cfg.PrivateKey)
	if err != nil {
		return err
	}

	d.SetId(key.PublicKey().String())
	d.Set("conf", buf.String())
	return nil
}
