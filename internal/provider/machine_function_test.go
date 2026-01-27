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

func TestMachineFunction_Valid(t *testing.T) {
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
                output "uname_machine" {
                    value = provider::uname::machine()
                }`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue(
						"uname_machine",
						knownvalue.StringExact(reference.Platform),
					),
				},
			},
		},
	})
}

func TestMachineFunction_Invalid(t *testing.T) {
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
                output "machine" {
                    value = provider::uname::machine("")
                }`,
				ExpectError: regexp.MustCompile(`Too many function arguments`),
			},
		},
	})
}

func TestMachineFunction_Null(t *testing.T) {
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
                output "uname_machine" {
                    value = provider::uname::machine(null)
                }`,
				ExpectError: regexp.MustCompile(`Too many function arguments`),
			},
		},
	})
}
