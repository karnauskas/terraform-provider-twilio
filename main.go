package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/terraform-providers/terraform-provider-twilio/twilio"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: twilio.Provider,
	})
}
