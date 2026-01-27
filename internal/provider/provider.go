package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	unameFunction "github.com/julienlevasseur/terraform-provider-uname/internal/function"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	// _ provider.Provider = (*UnameProvider)(nil)
	_ provider.Provider              = &UnameProvider{}
	_ provider.ProviderWithFunctions = &UnameProvider{}
)

// UnameProvider is the provider implementation.
type UnameProvider struct {
	version string
}

// New is a helper function to simplify provider server and testing implementation.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &UnameProvider{
			version: version,
		}
	}
}

// Metadata returns the provider type name.
func (p *UnameProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "uname"
}

// Schema defines the provider-level schema for configuration data.
func (p *UnameProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{}
}

// Configure prepares a Uname API client for data sources and resources.
func (p *UnameProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
}

// DataSources defines the data sources implemented in the provider.
func (p *UnameProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewUnameDataSource,
	}
}

// Resources defines the resources implemented in the provider.
func (p *UnameProvider) Resources(ctx context.Context) []func() resource.Resource {
	// return []func() resource.Resource{}
	return nil
}

func (p *UnameProvider) Functions(_ context.Context) []func() function.Function {
	return []func() function.Function{
		unameFunction.NewKernelNameFunction,
		unameFunction.NewNodeNameFunction,
		unameFunction.NewKernelReleaseFunction,
		unameFunction.NewMachineFunction,
		unameFunction.NewOperatingSystemFunction,
	}
}
