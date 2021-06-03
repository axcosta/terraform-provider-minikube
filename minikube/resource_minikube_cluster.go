package minikube

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/minikube/pkg/minikube/cluster"
	"k8s.io/minikube/pkg/minikube/config"
	"k8s.io/minikube/pkg/minikube/machine"
)

func resourceCluster() *schema.Resource {
	return &schema.Resource{
		Description: "Cluster resource in the Terraform Minikube provider.",

		CreateContext: resourceClusterCreate,
		ReadContext:   resourceClusterRead,
		UpdateContext: resourceClusterUpdate,
		DeleteContext: resourceClusterDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Description: "Cluster name.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func resourceClusterCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	log.Println("=================== Creating Minikube Cluster ==================")
	// name := d.Get("name").(string)
	// idFromAPI := "my-id"
	// d.SetId(idFromAPI)

	api, err := machine.NewAPIClient()
	if err != nil {
		log.Printf("Error getting client: %s\n", err)
		return err
	}
	defer api.Close()

	config := config.MachineConfig{
		// MinikubeISO:         isoURL,
		// Memory:              memory,
		// CPUs:                cpus,
		// DiskSize:            diskSizeMB,
		// VMDriver:            vmDriver,
		// XhyveDiskDriver:     xhyveDiskDriver,
		// DockerEnv:           dockerEnv.([]string),
		// DockerOpt:           dockerOpt.([]string),
		// InsecureRegistry:    insecureRegistry.([]string),
		// RegistryMirror:      registryMirror.([]string),
		// HostOnlyCIDR:        hostOnlyCIDR,
		// HypervVirtualSwitch: hypervVirtualSwitch,
		// KvmNetwork:          kvmNetwork,
		// Downloader:          pkgutil.DefaultDownloader{},
		// DisableDriverMounts: disableDriverMounts,
	}

	host, err := cluster.StartHost(api, config)
	if err != nil {
		log.Printf("Error starting host: %v", err)
		return err
	}

	return diags
}

func resourceClusterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}

func resourceClusterUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}

func resourceClusterDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}
