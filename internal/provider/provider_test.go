package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/julienlevasseur/goInfo"
)

func TestProviderFactory(t *testing.T) {
	t.Parallel()

	mockVersion := "0.0.0"

	p := New(mockVersion)
	if p == nil {
		t.Errorf("New() did not return a providerFunc")
	}
}

func TestProvider(t *testing.T) {
	t.Parallel()

	reference, _ := goInfo.GetInfo()

	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"uname": providerserver.NewProtocol6WithError(&UnameProvider{}),
		},
		Steps: []resource.TestStep{
			{
				Config: `
				data "uname" "localhost" {}

                output "uname_kernel_name" {
                    value = data.uname.localhost.kernel_name
                }`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue(
						"uname_kernel_name",
						knownvalue.StringExact(reference.Kernel),
					),
				},
			},
			{
				Config: `
				data "uname" "localhost" {}

                output "uname_node_name" {
                    value = data.uname.localhost.node_name
                }`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue(
						"uname_node_name",
						knownvalue.StringExact(reference.Hostname),
					),
				},
			},
			{
				Config: `
				data "uname" "localhost" {}

                output "uname_kernel_release" {
                    value = data.uname.localhost.kernel_release
                }`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue(
						"uname_kernel_release",
						knownvalue.StringExact(reference.Core),
					),
				},
			},
			{
				Config: `
				data "uname" "localhost" {}

                output "uname_machine" {
                    value = data.uname.localhost.machine
                }`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue(
						"uname_machine",
						knownvalue.StringExact(reference.Platform),
					),
				},
			},
			{
				Config: `
				data "uname" "localhost" {}

                output "uname_operating_system" {
                    value = data.uname.localhost.operating_system
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
