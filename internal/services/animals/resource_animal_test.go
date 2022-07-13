package animal_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp-csa/terraform-provider-csa/internal/testing"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceAnimal(t *testing.T) {
	resourceName := "demo_animal.foo"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acceptanceTesting.TestAccPreCheck(t) },
		ProviderFactories: acceptanceTesting.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceAnimalCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(resourceName, "class", regexp.MustCompile("^Bird")),
					resource.TestMatchResourceAttr(resourceName, "animal", regexp.MustCompile("^Peregrine Falcon")),
				),
			},
			{
				ResourceName: resourceName,
				ImportState: true,
				ImportStateVerify: false,
			},
			{
				Config: testAccResourceAnimalUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(resourceName, "class", regexp.MustCompile("^Mammal")),
					resource.TestMatchResourceAttr(resourceName, "animal", regexp.MustCompile("^Horse")),
				),
			},
			{
				Config: testAccResourceAnimalDelete,
				Check: resource.ComposeTestCheckFunc(
					testDoesNotExistsInState(resourceName),
				),
			},
		},
	})
}

func testDoesNotExistsInState(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		_, ok := s.RootModule().Resources[resourceName]
		if ok {
			return fmt.Errorf("Found: %s", resourceName)
		}

		return nil
	}
}

const testAccResourceAnimalCreate = `
resource "demo_animal" "foo" {
  class = "Bird"
}
`

const testAccResourceAnimalUpdate = `
resource "demo_animal" "foo" {
  class = "Mammal"
}
`

const testAccResourceAnimalDelete = `

`
