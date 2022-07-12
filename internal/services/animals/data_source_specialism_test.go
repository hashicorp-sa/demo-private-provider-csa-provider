package animal_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp-csa/terraform-provider-csa/internal/testing"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceAnimal(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acceptanceTesting.TestAccPreCheck(t) },
		ProviderFactories: acceptanceTesting.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceAnimal,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("data.demo_animal.foo", "class", regexp.MustCompile("^Bird")),
					resource.TestMatchResourceAttr("data.demo_animal.foo", "animal", regexp.MustCompile("^Peregrine Falcon")),
				),
			},
		},
	})
}

const testAccDataSourceAnimal = `
data "demo_animal" "foo" {
  class = "Bird"
}
`
