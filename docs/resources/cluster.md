---
page_title: "minikube_cluster Resource - terraform-provider-minikube"
subcategory: ""
description: |-
  Sample resource in the Terraform provider Minikube.
---

# Resource `minikube_cluster`

Sample resource in the Terraform provider Minikube.

## Example Usage

```terraform
resource "minikube_cluster" "example" {
  sample_attribute = "foo"
}
```

## Schema

### Optional

- **id** (String) The ID of this resource.
- **sample_attribute** (String) Sample attribute.


