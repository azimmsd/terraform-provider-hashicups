package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = &HashicupsCoffeesDataSource{}

func NewHashicupsCoffeesDataSource() datasource.DataSource {
	return &HashicupsCoffeesDataSource{}
}

type HashicupsCoffeesDataSource struct{}

type HashicupsCoffeesDataSourceModel struct {
	Id types.String `tfsdk:"id"`
}

func (d *HashicupsCoffeesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_coffees"
}

func (d *HashicupsCoffeesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "List the available coffee products.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Identifier for the coffee list.",
				Computed:            true,
			},
		},
	}
}

func (d *HashicupsCoffeesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
}

func (d *HashicupsCoffeesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data HashicupsCoffeesDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = types.StringValue("hashicups-coffees")
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
