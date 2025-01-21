---
page_title: "yaml_validate function - terraform-provider-utils"
subcategory: ""
description: |-
  Validate a YAML string against a JSON schema.
---

# function: yaml_validate

Validate a list of a yaml string against a JSON schema. Returns the YAML string if it fits the schema.

## Example Usage

```terraform
locals {
  schema = <<-EOT
{
  "$id": "pets-schema.json",
  "$schema": "http://json-schema.org/draft-03/schema",
  "type": "object",
  "properties": {
    "pets": {
      "type": "array",
      "items": {
        "type": "object",
        "properties": {
          "name": {
            "type": "string"
          },
          "species": {
            "type": "string",
            "enum": ["cat", "dog", "bird"]
          }
        }
      }
    }
  }
}
EOT

  yaml = <<-EOT
pets:
  - name: Fluffy
    species: cat
  - name: Fido
    species: dog
  EOT
}

output "output" {
  value = provider::utils::yaml_validate(local.yaml, local.schema)
}

/* 
output = <<-EOT
pets:
  - name: Fluffy
    species: cat
  - name: Fido
    species: dog
  EOT
*/
```

## Signature

```text
yaml_validate(input, schema) string
```

## Arguments

1. `input` (String) A YAML string to be validated.
2. `schema` (String) The JSON schema

## Return Type

The return type of `yaml_validate` is a string containing the validated YAML.
