package animal

import (
	"context"

	"github.com/hashicorp-sa/terraform-provider-csa/client/animals"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceAnimal() *schema.Resource {
	return &schema.Resource{
		Description: "Animal Example Data Source.",

		ReadContext: dataSourceAnimalsRead,

		Schema: map[string]*schema.Schema{
			"animal_id": {
				Description: "Id of animal.",
				Type:        schema.TypeString,
				Required:    true,
			},
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
		},
	}
}

func dataSourceAnimalsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := meta.(animals.Client)

	var model animals.AnimalReadModel
	model.Id = d.Get("animal_id").(string)
	model.Class = d.Get("class").(string) // This is a cheat for our stateless example.

	animal, err := client.Read(model)
	if err != nil {
		return diag.Errorf("error reading animal: %s", err)
	}

	d.Set("animal", animal.Animal)
	d.SetId(animal.Id)

	tflog.Trace(ctx, "read an animal")

	return diags
}
