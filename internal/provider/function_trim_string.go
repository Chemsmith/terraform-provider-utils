package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ function.Function = TrimStringFunction{}

func NewTrimStringFunction() function.Function {
	return &TrimStringFunction{}
}

type TrimStringFunction struct{}

func (r TrimStringFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "string_filter_trim"
}

func (r TrimStringFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Trim a string to at most the given length, first removing one or more given prefixes/suffixes.",
		MarkdownDescription: "Trim a string to at most the given length, first removing one or more given prefixes/suffixes. The substrings will be removed in alphabetical order, and the function will preserve the suffix on the string when truncating if `preserve_suffix` is set to true.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:                "input",
				MarkdownDescription: "A string to be trimmed.",
			},
			function.Int32Parameter{
				Name:                "length",
				MarkdownDescription: "The maximum length of the string after trimming.",
			},
			function.ListParameter{
				Name:                "substrings",
				MarkdownDescription: "A list of prefixes/suffixes to remove from the string, these will be removed in alphabetical order.",
				ElementType:         types.StringType,
			},
			function.BoolParameter{
				Name:                "preserve_suffix",
				MarkdownDescription: "If true, the function will preserve the suffix on the string when truncating.",
			},
		},
		Return: function.StringReturn{},
	}
}

func (r TrimStringFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var input string
	var length int32
	var substrings []string
	var preserveSuffix bool
	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &input, &length, &substrings, &preserveSuffix))

	if resp.Error != nil {
		return
	}

	trimmed := TrimString(input, int(length), substrings, preserveSuffix)
	if trimmed == "" {
		resp.Error = function.ConcatFuncErrors(function.NewFuncError("Trimmed string is empty"))
		return
	}
	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, trimmed))
}
