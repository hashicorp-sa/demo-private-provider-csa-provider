package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// providerFactories are used to instantiate a provider during acceptance testing.
// The factory function will be invoked for every Terraform CLI command executed
// to create a provider server to which the CLI can reattach.
var providerFactories = map[string]func() (*schema.Provider, error){
	"demo": func() (*schema.Provider, error) {
		return New("dev")(), nil
	},
}

func TestProvider(t *testing.T) {
	if err := New("dev")().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func testAccPreCheck(t *testing.T) {
	//For testing purposes only
	os.Setenv("ANIMALS_URL", "http://localhost:8080")
	os.Setenv("ANIMALS_TOKEN", "12345")

	if v := os.Getenv("ANIMALS_URL"); v == "" {
		t.Fatal("ANIMALS_URL must be set for acceptance tests")
	}
	if v := os.Getenv("ANIMALS_TOKEN"); v == "" {
		t.Fatal("ANIMALS_TOKEN must be set for acceptance tests")
	}
}
