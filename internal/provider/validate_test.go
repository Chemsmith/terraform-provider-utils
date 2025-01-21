package provider

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateYaml(t *testing.T) {
	tests := []struct {
		name         string
		yamlText     string
		schemaText   string
		expectErr    bool
		expectedYaml string
	}{
		{
			name: "valid yaml",
			yamlText: `name: john
age: 30`,
			schemaText: `{"type":"object","required":["name","age"],"properties":{"name":{"type":"string"},"age":{"type":"integer"}}}`,
			expectErr:  false,
		},
		{
			name: "invalid yaml",
			yamlText: `name john
			age 30`,
			schemaText: `{"type":"object","required":["name","age"],"properties":{"name":{"type":"string"},"age":{"type":"integer"}}}`,
			expectErr:  true,
		},
		{
			name: "non-schema compliant yaml",
			yamlText: `name: john
age: 30`,
			schemaText: `{"type":"object","required":["name","age"],"properties":{"name":{"type":"string"},"age":{"type":"string"}}}`,
			expectErr:  true,
		},
		{
			name:       "missing required field",
			yamlText:   `name: john`,
			schemaText: `{"type":"object","required":["name","age"],"properties":{"name":{"type":"string"},"age":{"type":"string"}}}`,
			expectErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateYaml(tt.yamlText, tt.schemaText)
			if tt.expectErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
		})
	}
}
