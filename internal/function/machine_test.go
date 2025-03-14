package function

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestMachineFunctionFactory(t *testing.T) {
	t.Parallel()

	p := NewMachineFunction()
	if p == nil {
		t.Errorf("NewMachineFunction() did not return a Function")
	}
}

func TestMachineFunctionRun(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		request  function.RunRequest
		expected function.RunResponse
	}{
		"value-valid": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData(
					[]attr.Value{types.StringValue("")},
				),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(types.StringValue("x86_64")),
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := function.RunResponse{
				Result: function.NewResultData(types.StringUnknown()),
			}

			var f MachineFunction
			f.Run(context.Background(), testCase.request, &got)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
