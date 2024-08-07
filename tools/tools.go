//go:build tools

package tools

import (
	_ "github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs"
)

//go:generate go install github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs
