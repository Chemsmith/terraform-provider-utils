// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure UtilsProvider satisfies various provider interfaces.
var _ provider.Provider = &UtilsProvider{}
var _ provider.ProviderWithFunctions = &UtilsProvider{}
var _ provider.ProviderWithEphemeralResources = &UtilsProvider{}

// UtilsProvider defines the provider implementation.
type UtilsProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// UtilsProviderModel describes the provider data model.
type UtilsProviderModel struct {
	Endpoint types.String `tfsdk:"endpoint"`
}

func (p *UtilsProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "utils"
	resp.Version = p.version
}

func (p *UtilsProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{}
}

func (p *UtilsProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
}

func (p *UtilsProvider) Resources(ctx context.Context) []func() resource.Resource {
	return nil

}

func (p *UtilsProvider) EphemeralResources(ctx context.Context) []func() ephemeral.EphemeralResource {
	return nil

}

func (p *UtilsProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return nil
}

func (p *UtilsProvider) Functions(ctx context.Context) []func() function.Function {
	return []func() function.Function{
		NewYamlValidateFunction,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &UtilsProvider{
			version: version,
		}
	}
}
