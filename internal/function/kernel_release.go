package function

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/julienlevasseur/goInfo"
)

var (
	_ function.Function = &KernelReleaseFunction{}
)

func NewKernelReleaseFunction() function.Function {
	return &KernelReleaseFunction{}
}

type KernelReleaseFunction struct{}

func (f *KernelReleaseFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "kernel_release"
}

func (f *KernelReleaseFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "print the kernel release",
		Description: "print the kernel release.",
		Parameters:  []function.Parameter{},
		Return:      function.StringReturn{},
	}
}

func (f *KernelReleaseFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	goInfo, _ := goInfo.GetInfo()

	resp.Error = function.ConcatFuncErrors(
		resp.Error,
		resp.Result.Set(ctx, goInfo.Core),
	)
}
