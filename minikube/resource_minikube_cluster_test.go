package minikube

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceCluster(t *testing.T) {
	// t.Skip("resource not yet implemented, remove this once you add your own code")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceCluster,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"minikube_cluster.test", "name", "minikube"),
				),
			},
		},
	})
}

const testAccResourceCluster = `
resource "minikube_cluster" "test" {	
}
`
