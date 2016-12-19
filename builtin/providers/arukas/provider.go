package arukas

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"os"
)

const (
	JSONTokenParamName  = "ARUKAS_JSON_API_TOKEN"
	JSONSecretParamName = "ARUKAS_JSON_API_SECRET"
	JSONUrlParamName    = "ARUKAS_JSON_API_URL"
	JSONDebugParamName  = "ARUKAS_DEBUG"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ARUKAS_JSON_API_TOKEN", nil),
				Description: "your Arukas APIKey(token)",
			},
			"secret": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ARUKAS_JSON_API_SECRET", nil),
				Description: "your Arukas APIKey(secret)",
			},
			"api_url": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ARUKAS_JSON_API_URL", "https://app.arukas.io/api/"),
				Description: "default Arukas API url",
			},
			"trace": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ARUKAS_DEBUG", ""),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"arukas_container": resourceArukasContainer(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {

	os.Setenv(JSONTokenParamName, d.Get("token").(string))
	os.Setenv(JSONSecretParamName, d.Get("secret").(string))
	os.Setenv(JSONUrlParamName, d.Get("api_url").(string))
	os.Setenv(JSONDebugParamName, d.Get("trace").(string))

	config := Config{}

	return config.NewClient()
}
