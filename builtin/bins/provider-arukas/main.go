package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/yamamoto-febc/terraform-provider-arukas/builtin/providers/arukas"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: arukas.Provider,
	})
}
