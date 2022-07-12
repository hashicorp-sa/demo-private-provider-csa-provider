package provider

import (
	"context"

	"github.com/hashicorp-csa/terraform-provider-csa/client/animals"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	schema.DescriptionKind = schema.StringMarkdown
}

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			DataSourcesMap: map[string]*schema.Resource{
				"demo_animal": dataSourceAnimal(),
			},
			ResourcesMap: map[string]*schema.Resource{
				"demo_animal": resourceAnimals(),
			},
			Schema: map[string]*schema.Schema{
				"url": {
					Type:        schema.TypeString,
					Required:    true,
					DefaultFunc: schema.EnvDefaultFunc("ANIMALS_URL", nil),
				},
				"token": {
					Type:        schema.TypeString,
					Required:    true,
					Sensitive:   true,
					DefaultFunc: schema.EnvDefaultFunc("ANIMALS_TOKEN", nil),
				},
			},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(ctx context.Context, r *schema.ResourceData) (interface{}, diag.Diagnostics) {
		url := r.Get("url").(string)
		token := r.Get("token").(string)

		var diags diag.Diagnostics

		if (url != "") && (token != "") {
			client, err := animals.New(url, token)
			if err != nil {
				return nil, diag.FromErr(err)
			}

			return client, diags
		}

		return nil, diag.Errorf("url and token are required")
	}
}
