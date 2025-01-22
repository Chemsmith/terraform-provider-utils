# `utils` Terraform Provider

A small utility provider for Terraform 1.8+ to provide functions not available out of the box.

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 1.8

## Using the provider

```terraform
terraform {
  required_providers {
    utils = {
      source = "chemsmith/utils"
    }
  }
}

provider "utils" {}
```

Additional documentation, including available resources and their arguments/attributes can be found on the [Terraform documentation website](https://registry.terraform.io/providers/chemsmith/utils/latest/docs).

## Included functions
- [`yaml_validate`](docs/functions/yaml_validate.md) check a YAML string against a JSON schema for compliance.


## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (see [Requirements](#requirements) above).

To compile the provider, run `go install`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
make test
```
