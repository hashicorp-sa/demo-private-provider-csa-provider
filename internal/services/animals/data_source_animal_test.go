package animal_test

import (
	"regexp"
	"testing"

	acceptanceTesting "github.com/hashicorp-sa/terraform-provider-csa/internal/testing"
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
					resource.TestMatchResourceAttr("data.demo_animal.foo", "animal_id", regexp.MustCompile("^12345")),
					resource.TestMatchResourceAttr("data.demo_animal.foo", "class", regexp.MustCompile("^Bird")),
					resource.TestMatchResourceAttr("data.demo_animal.foo", "animal", regexp.MustCompile("^Peregrine Falcon")),
				),
			},
		},
	})
}

const testAccDataSourceAnimal = `
data "demo_animal" "foo" {
  animal_id = "12345"
  class     = "Bird"
}
`
