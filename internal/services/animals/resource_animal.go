package animal

import (
	"context"

	"github.com/hashicorp-csa/terraform-provider-csa/client/animals"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceAnimal() *schema.Resource {
	return &schema.Resource{
		Description: "Animals Example Resource.",

		CreateContext: resourceAnimalsCreate,
		ReadContext:   resourceAnimalsRead,
		UpdateContext: resourceAnimalsUpdate,
		DeleteContext: resourceAnimalsDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

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
	client := meta.(animals.Client)

	var model animals.AnimalCreateModel
	model.Class = d.Get("class").(string)

	animal, err := client.Create(model)
	if err != nil {
		return diag.Errorf("error creating animal: %s", err)
	}

	d.SetId(animal.Id)
	d.Set("animal", animal.Animal)
	d.Set("date_configured", animal.Created)

	tflog.Trace(ctx, "created an animal")

	return diags
}

func resourceAnimalsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := meta.(animals.Client)

	var model animals.AnimalReadModel
	model.Id = d.Id()
	model.Class = d.Get("class").(string) // This is a cheat for our stateless example.
	if d.Get("date_configured") != nil {
		model.Created = d.Get("date_configured").(string) // This is a cheat for our stateless example.
	}

	animal, err := client.Read(model)
	if err != nil {
		return diag.Errorf("error reading animal: %s", err)
	}

	d.Set("animal", animal.Animal)

	tflog.Trace(ctx, "read an animal")

	return diags
}

func resourceAnimalsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := meta.(animals.Client)

	var model animals.AnimalUpdateModel

	model.Id = d.Id()
	model.Class = d.Get("class").(string)

	animal, err := client.Update(model)
	if err != nil {
		return diag.Errorf("error updating animal: %s", err)
	}
	d.Set("animal", animal.Animal)
	d.Set("date_configured", animal.Created)

	tflog.Trace(ctx, "updated an animal")

	return diags
}

func resourceAnimalsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := meta.(animals.Client)

	var model animals.AnimalDeleteModel
	model.Id = d.Id()

	err := client.Delete(model)
	if err != nil {
		return diag.Errorf("error deleting animal: %s", err)
	}

	d.SetId("")
	d.Set("date_configured", "")

	tflog.Trace(ctx, "deleted a resource")

	return diags
}
