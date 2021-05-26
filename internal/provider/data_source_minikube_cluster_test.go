package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceMinikubeCluster(t *testing.T) {
	t.Skip("data source not yet implemented, remove this once you add your own code")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceMinikubeCluster,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"data.minikube_cluster.foo", "sample_attribute", regexp.MustCompile("^ba")),
				),
			},
		},
	})
}

const testAccDataSourceMinikubeCluster = `
data "minikube_cluster" "foo" {
	sample_attribute = "bar"
}
`
