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

func TestNodeNameFunctionFactory(t *testing.T) {
	t.Parallel()

	p := NewNodeNameFunction()
	if p == nil {
		t.Errorf("NewNodeNameFunction() did not return a Function")
	}
}

func getHostname(t testing.TB) string {
	t.Helper()

	path := filepath.Join("/etc/hostname")
	data, err := os.ReadFile(path)
	require.NoError(t, err)
	h := strings.Split(string(data), " ")[0]
	return strings.Trim(h, "\n")
}

func TestNodeNameFunctionRun(t *testing.T) {
	t.Parallel()

	hostname := getHostname(t)

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
				Result: function.NewResultData(types.StringValue(hostname)),
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

			var f NodeNameFunction
			f.Run(context.Background(), testCase.request, &got)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
