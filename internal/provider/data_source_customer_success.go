package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCustomerSuccess() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Sample data source in the Terraform provider scaffolding.",

		ReadContext: dataSourceCustomerSuccessRead,

		Schema: map[string]*schema.Schema{
			"customer_success_architect": {
				Description: "Customer Success Architect.",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func dataSourceCustomerSuccessRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics	
	client := csa_client{id: d.Get("customer_success_architect").(string)}

	d.Set("specialism", client.GetSpecialism())

	return diags
}
