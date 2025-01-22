locals {
  schema = <<-EOT
{
  "$id": "pets-schema.json",
  "$schema": "http://json-schema.org/draft-07/schema",
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
