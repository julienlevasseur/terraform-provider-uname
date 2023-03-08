package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/matishsiao/goInfo"
)

var (
	_ datasource.DataSource = (*unameDataSource)(nil)
)

func NewUnameDataSource() datasource.DataSource {
	return &unameDataSource{}
}

type unameDataSource struct{}

func (u *unameDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Gets Uname informations from the system.",
		Attributes: map[string]schema.Attribute{
			"kernel_name": schema.StringAttribute{
				Description: "kernel name",
				Computed:    true,
			},
			"node_name": schema.StringAttribute{
				Description: "network node hostname",
				Computed:    true,
			},
			"kernel_release": schema.StringAttribute{
				Description: "kernel release",
				Computed:    true,
			},
			"machine": schema.StringAttribute{
				Description: "machine hardware name",
				Computed:    true,
			},
			"operating_system": schema.StringAttribute{
				Description: "operating system",
				Computed:    true,
			},
		},
	}
}

func (n *unameDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName
}

func (n *unameDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

	var config unameDataSourceModel
	diags := req.Config.Get(ctx, &config)

	goInfo, _ := goInfo.GetInfo()

	state := unameDataSourceModel{
		KernelName:      types.StringValue(goInfo.Kernel),
		NodeName:        types.StringValue(goInfo.Hostname),
		KernelRelease:   types.StringValue(goInfo.Core),
		Machine:         types.StringValue(goInfo.Platform),
		OperatingSystem: types.StringValue(goInfo.OS),
	}
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

type unameDataSourceModel struct {
	KernelName      types.String `tfsdk:"kernel_name"`
	NodeName        types.String `tfsdk:"node_name"`
	KernelRelease   types.String `tfsdk:"kernel_release"`
	Machine         types.String `tfsdk:"machine"`
	OperatingSystem types.String `tfsdk:"operating_system"`
}
