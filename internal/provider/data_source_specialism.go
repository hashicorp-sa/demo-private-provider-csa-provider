package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSpecialism() *schema.Resource {
	return &schema.Resource{
		Description: "Customer Success Example Data Source.",

		ReadContext: dataSourceSpecialismRead,

		Schema: map[string]*schema.Schema{
			"customer_success_architect": {
				Description: "Customer Success Architect.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"specialism": {
				Description: "Customer Success Architect Specialism.",
				Computed:    true,
				Type:        schema.TypeString,
			},
		},
	}
}

func dataSourceSpecialismRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	
	architect := d.Get("customer_success_architect").(string)
	client := csa_client{id: architect}

	d.SetId(architect)

	d.Set("specialism", client.GetSpecialism())

	return diags
}
