// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure Provider satisfies various provider interfaces.
var _ provider.Provider = &WaitForProvider{}
var _ provider.ProviderWithFunctions = &WaitForProvider{}
var _ provider.ProviderWithEphemeralResources = &WaitForProvider{}
var _ provider.ProviderWithActions = &WaitForProvider{}

// WaitForProvider defines the provider implementation.
type WaitForProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// WaitForProviderModel describes the provider data model.
type WaitForProviderModel struct {
}

func (p *WaitForProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "wait"
	resp.Version = p.version
}

func (p *WaitForProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{}
}

func (p *WaitForProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
}

func (p *WaitForProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}

func (p *WaitForProvider) EphemeralResources(ctx context.Context) []func() ephemeral.EphemeralResource {
	return []func() ephemeral.EphemeralResource{}
}

func (p *WaitForProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func (p *WaitForProvider) Functions(ctx context.Context) []func() function.Function {
	return []func() function.Function{}
}

// Actions implements provider.ProviderWithActions.
func (p *WaitForProvider) Actions(context.Context) []func() action.Action {
	return []func() action.Action{
		NewWaitForPortAction,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &WaitForProvider{
			version: version,
		}
	}
}
