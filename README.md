# Terraform Provider for Minikube

## Overview

The Terraform Provider for Minikube enables Terraform to provision local Minikube clusters. This is based on [Jeff Gensler's Minikube provider project](https://github.com/jgensler8/terraform-provider-minikube).

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 0.13.x
- [Go](https://golang.org/doc/install) >= 1.15

## Organization

- Code is at `minikube/` folder,
- Examples (`examples/`) and generated documentation (`docs/`),
- Miscellaneous meta files at `/` and `tools/`

## Building the provider

1. Clone the repository
2. Enter the repository directory
3. Build the provider using the Go `install` command:

```sh
go install
```

## Adding Dependencies

This provider uses [Go modules](https://github.com/golang/go/wiki/Modules). Please see the Go documentation for the most up to date information about using Go modules.

To add a new dependency `github.com/author/dependency` to your Terraform provider:

```sh
go get github.com/author/dependency
go mod tidy
```

Then commit the changes to `go.mod` and `go.sum`.

## Using the provider

See documentation at `docs/`

## Upgrading the provider

The provider doesn't upgrade automatically once you've started using it. After a new release you can run

```bash
terraform init -upgrade
```

to upgrade to the latest stable version. See the [Terraform website](https://www.terraform.io/docs/configuration/providers.html#provider-versions)
for more information on provider upgrades, and how to set version constraints on your provider.

## Developing the provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (see [Requirements](#requirements) above).

To compile the provider, run `go install`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

To generate or update documentation, run `go generate`.

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
make testacc
```

## Publishing the provider

[Publish it on the Terraform Registry](https://www.terraform.io/docs/registry/providers/publishing.html) so that others can use it.
