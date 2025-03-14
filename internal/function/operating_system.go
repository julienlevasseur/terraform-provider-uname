package function

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/julienlevasseur/goInfo"
)

var (
	_ function.Function = &OperatingSystemFunction{}
)

func NewOperatingSystemFunction() function.Function {
	return &OperatingSystemFunction{}
}

type OperatingSystemFunction struct{}

func (f *OperatingSystemFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "operating_system"
}

func (f *OperatingSystemFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "print the operating system",
		Description: "print the operating system.",
		Parameters:  []function.Parameter{},
		Return:      function.StringReturn{},
	}
}

func (f *OperatingSystemFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	goInfo, _ := goInfo.GetInfo()

	resp.Error = function.ConcatFuncErrors(
		resp.Error,
		resp.Result.Set(ctx, goInfo.OS),
	)
}
