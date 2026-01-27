package function

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/require"
)

func TestKernelReleaseFunctionFactory(t *testing.T) {
	t.Parallel()

	p := NewKernelReleaseFunction()
	if p == nil {
		t.Errorf("NewKernelReleaseFunction() did not return a Function")
	}
}

func getKernelVersion(t testing.TB) string {
	t.Helper()

	path := filepath.Join("/proc/version")
	data, err := os.ReadFile(path)
	require.NoError(t, err)
	return strings.Split(string(data), " ")[2]
}

func TestKernelReleaseFunctionRun(t *testing.T) {
	t.Parallel()

	kernelVersion := getKernelVersion(t)

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
				// Result: function.NewResultData(types.StringValue("6.8.0-39-generic")),
				Result: function.NewResultData(types.StringValue(kernelVersion)),
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

			var f KernelReleaseFunction
			f.Run(context.Background(), testCase.request, &got)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
