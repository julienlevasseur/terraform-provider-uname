package function

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/julienlevasseur/goInfo"
)

var (
	_ function.Function = &MachineFunction{}
)

func NewMachineFunction() function.Function {
	return &MachineFunction{}
}

type MachineFunction struct{}

func (f *MachineFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "machine"
}

func (f *MachineFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "print the machine hardware name",
		Description: "print the machine hardware name.",
		Parameters:  []function.Parameter{},
		Return:      function.StringReturn{},
	}
}

func (f *MachineFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	goInfo, _ := goInfo.GetInfo()

	resp.Error = function.ConcatFuncErrors(
		resp.Error,
		resp.Result.Set(ctx, goInfo.Platform),
	)
}
