package provider

import (
	"errors"
	"strings"

	"github.com/santhosh-tekuri/jsonschema/v3"
	"gopkg.in/yaml.v3"
)

func ValidateYaml(yamlText string, schemaText string) error {
	var m interface{}
	if err := yaml.Unmarshal([]byte(yamlText), &m); err != nil {
		return err
	}

	m, err := toStringKeys(m)
	if err != nil {
		return err
	}

	compiler := jsonschema.NewCompiler()
	if err := compiler.AddResource("schema.json", strings.NewReader(schemaText)); err != nil {
		return err
	}

	schema, err := compiler.Compile("schema.json")
	if err != nil {
		return err
	}

	if err := schema.ValidateInterface(m); err != nil {
		return err
	}

	return nil
}

func toStringKeys(val interface{}) (interface{}, error) {
	switch val := val.(type) {
	case map[interface{}]interface{}:
		m := make(map[string]interface{})
		for k, v := range val {
			k, ok := k.(string)
			if !ok {
				return nil, errors.New("found non-string key")
			}
			m[k] = v
		}
		return m, nil
	case []interface{}:
		var err error
		var l = make([]interface{}, len(val))
		for i, v := range l {
			l[i], err = toStringKeys(v)
			if err != nil {
				return nil, err
			}
		}
		return l, nil
	default:
		return val, nil
	}
}
