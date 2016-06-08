package main

import (
	//"github.com/hashicorp/terraform/builtin/providers/sakuracloud"
	"github.com/hashicorp/terraform/plugin"
	arukas "github.com/yamamoto-febc/terraform-provider-arukas"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: arukas.Provider,
	})
}
