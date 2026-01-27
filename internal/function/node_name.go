package function

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/julienlevasseur/goInfo"
)

var (
	_ function.Function = &NodeNameFunction{}
)

func NewNodeNameFunction() function.Function {
	return &NodeNameFunction{}
}

type NodeNameFunction struct{}

func (f *NodeNameFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "node_name"
}

func (f *NodeNameFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "print the network node hostname",
		Description: "print the network node hostname.",
		Parameters:  []function.Parameter{},
		Return:      function.StringReturn{},
	}
}

func (f *NodeNameFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	goInfo, _ := goInfo.GetInfo()

	resp.Error = function.ConcatFuncErrors(
		resp.Error,
		resp.Result.Set(ctx, goInfo.Hostname),
	)
}
