package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/ibrohimislam/terraform-provider-passphrase/internal/provider"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{ProviderFunc: provider.New})
}
