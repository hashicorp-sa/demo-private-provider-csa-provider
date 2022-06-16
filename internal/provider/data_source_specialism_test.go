package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceSpecialism(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSpecialism,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("data.csa_specialism.foo", "customer_success_architect", regexp.MustCompile("^Jared Holgate")),
					resource.TestMatchResourceAttr("data.csa_specialism.foo", "specialism", regexp.MustCompile("^Terraform")),
				),
			},
		},
	})
}

const testAccDataSourceSpecialism = `
data "csa_specialism" "foo" {
  customer_success_architect = "Jared Holgate"
}
`
