package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var _ function.Function = YamlValidateFunction{}

func NewYamlValidateFunction() function.Function {
	return &YamlValidateFunction{}
}

type YamlValidateFunction struct{}

func (r YamlValidateFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "yaml_validate"
}

func (r YamlValidateFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Validate a yaml string against a schema",
		MarkdownDescription: "This function validates a yaml string against a schema.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:                "input",
				MarkdownDescription: "A YAML string to be validated.",
			},
			function.StringParameter{
				Name:                "schema",
				MarkdownDescription: "The JSON schema to validate the yaml against.",
			},
		},
		Return: function.StringReturn{},
	}
}

func (r YamlValidateFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var input string
	var schemaText string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &input, &schemaText))

	if resp.Error != nil {
		return
	}

	if err := ValidateYaml(input, schemaText); err != nil {
		resp.Error = function.ConcatFuncErrors(function.NewFuncError(err.Error()))
		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, input))
}
