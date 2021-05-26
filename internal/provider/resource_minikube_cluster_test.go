package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceMinikubeCluster(t *testing.T) {
	t.Skip("resource not yet implemented, remove this once you add your own code")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceMinikubeCluster,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"minikube_cluster.foo", "sample_attribute", regexp.MustCompile("^ba")),
				),
			},
		},
	})
}

const testAccResourceMinikubeCluster = `
resource "minikube_cluster" "foo" {
	sample_attribute = "bar"
}
`
