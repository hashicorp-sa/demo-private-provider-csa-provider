package animal_test

import (
	"fmt"
	"regexp"
	"testing"

	acceptanceTesting "github.com/hashicorp-sa/terraform-provider-csa/internal/testing"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceAnimal(t *testing.T) {
	resourceName := "demo_animal.foo"

	var resourceId1, resourceId2 string

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acceptanceTesting.TestAccPreCheck(t) },
		ProviderFactories: acceptanceTesting.ProviderFactories,
		CheckDestroy:      testDoesNotExistsInState(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceAnimalCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(resourceName, "class", regexp.MustCompile("^Bird")),
					resource.TestMatchResourceAttr(resourceName, "animal", regexp.MustCompile("^Peregrine Falcon")),
					testCheckInstanceExists(resourceName, &resourceId1),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: false,
			},
			{
				Config: testAccResourceAnimalUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(resourceName, "class", regexp.MustCompile("^Mammal")),
					resource.TestMatchResourceAttr(resourceName, "animal", regexp.MustCompile("^Horse")),
					testCheckInstanceExists(resourceName, &resourceId2),
					testCheckInstanceSame(&resourceId1, &resourceId2),
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

func testCheckInstanceExists(resourceName string, resourceId *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		resource, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Found: %s", resourceName)
		}

		*resourceId = resource.Primary.ID

		return nil
	}
}

func testCheckInstanceSame(before, after *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if *before != *after {
			return fmt.Errorf("The resource ID has changed from %s to %s", *before, *after)
		}

		return nil
	}
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
