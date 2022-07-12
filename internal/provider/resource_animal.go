package provider

import (
	"context"
	"github.com/hashicorp-csa/terraform-provider-csa/client/animals"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAnimals() *schema.Resource {
	return &schema.Resource{
		Description: "Animals Example Resource.",

		CreateContext: resourceAnimalsCreate,
		ReadContext:   resourceAnimalsRead,
		UpdateContext: resourceAnimalsUpdate,
		DeleteContext: resourceAnimalsDelete,

		Schema: map[string]*schema.Schema{
			"class": {
				Description: "Class of animal.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"animal": {
				Description: "Example animal of the specified class.",
				Computed:    true,
				Type:        schema.TypeString,
			},
			"date_configured": {
				Description: "Setup Date.",
				Computed:    true,
				Type:        schema.TypeString,
			},
		},
	}
}

func resourceAnimalsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	class := d.Get("class").(string)
	client := meta.(animals.Client)

	d.SetId(class)
	resourceAnimalsRead(ctx, d, meta)
	d.Set("date_configured", client.GetSetupDate())

	tflog.Trace(ctx, "created a resource")

	return diags
}

func resourceAnimalsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := meta.(animals.Client)

	d.Set("animal", client.GetAnimalFromClass(d.Id()))

	tflog.Trace(ctx, "read a resource")

	return diags
}

func resourceAnimalsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := meta.(animals.Client)

	resourceAnimalsRead(ctx, d, meta)
	d.Set("date_configured", client.GetSetupDate())

	tflog.Trace(ctx, "updated a resource")

	return diags
}

func resourceAnimalsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	d.SetId("")
	resourceAnimalsRead(ctx, d, meta)
	d.Set("date_configured", "")

	tflog.Trace(ctx, "deleted a resource")

	return diags
}
