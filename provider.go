package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	twilioc "github.com/tulip/twiliogo"
)

func provider() *schema.Provider {

	provider := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"account_sid": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("TWILIO_SID", nil),
				Description: "The SID (application ID) for the the Twilio API",
			},
			"auth_token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("TWILIO_AUTH_TOKEN", nil),
				Description: "Oauth token for the the Twilio API",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"twilio_phonenumber": resourcePhonenumber(),
		},
		ConfigureFunc: providerConfigure,
	}

	return provider
}

type twilioMeta struct {
	Client *twilioc.TwilioClient
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	return &twilioMeta{
		Client: twilioc.NewClient(
			d.Get("account_sid").(string),
			d.Get("auth_token").(string),
		),
	}, nil
}
