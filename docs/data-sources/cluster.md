---
page_title: "minikube_cluster Data Source - terraform-provider-minikube"
subcategory: ""
description: |-
  Sample data source in the Terraform provider Minikube.
---

# Data Source `minikube_cluster`

Sample data source in the Terraform provider Minikube.

## Example Usage

```terraform
data "minikube_cluster" "example" {
  sample_attribute = "foo"
}
```

## Schema

### Required

- **sample_attribute** (String) Sample attribute.

### Optional

- **id** (String) The ID of this resource.


