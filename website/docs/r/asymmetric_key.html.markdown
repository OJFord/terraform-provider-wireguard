---
layout: "wireguard"
page_title: "WireGuard: asymmetric key"
sidebar_current: "docs-wireguard-resource-asymmetric-key"
description: |-
  Provides a WireGuard asymmetric key resource. This can be used to create, read, and delete WireGuard keys in terraform state.
---

# wireguard_asymmetric_key

Provides a WireGuard asymmetric key resource. This can be used to create, read, and delete WireGuard keys in terraform state.

## Example Usage

Create a new key:

```hcl
resource "wireguard_asymmetric_key" "example" {
}
```

## Argument Reference

The following arguments are supported:

* `bind` - (Optional) A string to tie the lifecycle to, e.g. the terraform ID of another resource.
* `private_key` - (Optional) A supplied WireGuard private key. By default one is generated.

## Attributes Reference

In addition to the arguments listed above, the following attributes are exported:

* `public_key` - The public key corresponding to the (given or generated) private key.
