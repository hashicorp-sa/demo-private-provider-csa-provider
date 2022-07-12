package acceptanceTesting

import (
	"os"
	"testing"

	"github.com/hashicorp-csa/terraform-provider-csa/internal/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestAccPreCheck(t *testing.T) {
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

var ProviderFactories = map[string]func() (*schema.Provider, error){
	"demo": func() (*schema.Provider, error) {
		return provider.New("dev")(), nil
	},
}
