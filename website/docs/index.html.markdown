---
layout: "wireguard"
page_title: "Provider: wireguard"
sidebar_current: "docs-wireguard-index"
description: |-
  The WireGuard provider is used to handle WireGuard metadata.
---

# WireGuard Provider

The WireGuard provider is used to handle WireGuard metadata.

## Example Usage

```hcl
# Configure the WireGuard Provider
provider "wireguard" {
}

# Create a WireGuard key
resource "wireguard_asymmetric_key" "server" {
  # ...
}
```

## Argument Reference

The following arguments are supported:
