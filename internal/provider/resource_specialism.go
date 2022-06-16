package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSpecialism() *schema.Resource {
	return &schema.Resource{
		Description: "Customer Success Example Resource.",

		CreateContext: resourceSpecialismCreate,
		ReadContext:   resourceSpecialismRead,
		UpdateContext: resourceSpecialismUpdate,
		DeleteContext: resourceSpecialismDelete,

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
			"date_configured": {
				Description: "Customer Success Architect Setup Date.",
				Computed:    true,
				Type:        schema.TypeString,
			},
		},
	}
}

func resourceSpecialismCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	architect := d.Get("customer_success_architect").(string)
	client := csa_client{id: architect}

	d.SetId(architect)
	resourceSpecialismRead(ctx, d, meta)
	d.Set("date_configured", client.GetSetupDate())

	tflog.Trace(ctx, "created a resource")

	return diags
}

func resourceSpecialismRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := csa_client{id: d.Id()}

	d.Set("specialism", client.GetSpecialism())

	tflog.Trace(ctx, "read a resource")

	return diags
}

func resourceSpecialismUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := csa_client{id: d.Id()}

	resourceSpecialismRead(ctx, d, meta)
	d.Set("date_configured", client.GetSetupDate())

	tflog.Trace(ctx, "updated a resource")

	return diags
}

func resourceSpecialismDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	d.SetId("")
	resourceSpecialismRead(ctx, d, meta)
	d.Set("date_configured", "")

	tflog.Trace(ctx, "deleted a resource")

	return diags
}
