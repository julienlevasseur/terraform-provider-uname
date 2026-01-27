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

func TestOperatingSystemFunction_Valid(t *testing.T) {
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
                output "uname_operating_system" {
                    value = provider::uname::operating_system()
                }`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue(
						"uname_operating_system",
						knownvalue.StringExact(reference.OS),
					),
				},
			},
		},
	})
}

func TestOperatingSystemFunction_Invalid(t *testing.T) {
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
                output "operating_system" {
                    value = provider::uname::operating_system("")
                }`,
				ExpectError: regexp.MustCompile(`Too many function arguments`),
			},
		},
	})
}

func TestOperatingSystemFunction_Null(t *testing.T) {
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
                output "uname_operating_system" {
                    value = provider::uname::operating_system(null)
                }`,
				ExpectError: regexp.MustCompile(`Too many function arguments`),
			},
		},
	})
}
