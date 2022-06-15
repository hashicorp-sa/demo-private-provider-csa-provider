package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceScaffolding(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceScaffolding,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"customer_success.foo", "customer_success_architect", regexp.MustCompile("^Jared Holgate")),
				),
			},
		},
	})
}

const testAccResourceScaffolding = `
resource "customer_success" "foo" {
  customer_success_architect = "Jared Holgate"
}
`
