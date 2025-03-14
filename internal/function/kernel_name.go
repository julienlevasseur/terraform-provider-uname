package function

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/julienlevasseur/goInfo"
)

var (
	_ function.Function = &KernelNameFunction{}
)

func NewKernelNameFunction() function.Function {
	return &KernelNameFunction{}
}

type KernelNameFunction struct{}

func (f *KernelNameFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "kernel_name"
}

func (f *KernelNameFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "print the kernel name",
		Description: "print the kernel name.",
		Parameters:  []function.Parameter{},
		Return:      function.StringReturn{},
	}
}

func (f *KernelNameFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	goInfo, err := goInfo.GetInfo()
	if err != nil {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError(err.Error()))
	}

	resp.Error = function.ConcatFuncErrors(
		resp.Error,
		resp.Result.Set(ctx, goInfo.Kernel),
	)
}
