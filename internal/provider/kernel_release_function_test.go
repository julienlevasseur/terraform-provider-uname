package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
	"github.com/julienlevasseur/goInfo"
)

func TestKernelReleaseFunction_Valid(t *testing.T) {
	t.Parallel()

	reference, _ := goInfo.GetInfo()

	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"uname": providerserver.NewProtocol6WithError(&UnameProvider{}),
		},
		Steps: []resource.TestStep{
			{
				Config: `
                output "uname_kernel_release" {
                    value = provider::uname::kernel_release()
                }`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue(
						"uname_kernel_release",
						knownvalue.StringExact(reference.Core),
					),
				},
			},
		},
	})
}

func TestKernelReleaseFunction_Invalid(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"uname": providerserver.NewProtocol6WithError(&UnameProvider{}),
		},
		Steps: []resource.TestStep{
			{
				Config: `
                output "kernel_release" {
                    value = provider::uname::kernel_release("")
                }`,
				ExpectError: regexp.MustCompile(`Too many function arguments`),
			},
		},
	})
}

func TestKernelReleaseFunction_Null(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"uname": providerserver.NewProtocol6WithError(&UnameProvider{}),
		},
		Steps: []resource.TestStep{
			{
				Config: `
                output "uname_kernel_release" {
                    value = provider::uname::kernel_name(null)
                }`,
				ExpectError: regexp.MustCompile(`Too many function arguments`),
			},
		},
	})
}
