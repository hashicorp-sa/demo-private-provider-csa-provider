package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceSpecialism(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceScaffolding,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("csa_specialism.foo", "customer_success_architect", regexp.MustCompile("^Jared Holgate")),
					resource.TestMatchResourceAttr("csa_specialism.foo", "specialism", regexp.MustCompile("^Terraform")),
				),
			},
		},
	})
}

const testAccResourceScaffolding = `
resource "csa_specialism" "foo" {
  customer_success_architect = "Jared Holgate"
}
`
