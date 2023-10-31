package main

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/julienlevasseur/terraform-provider-uname/internal/provider"
)

var (
	version string = "dev"
)

func main() {
	err := providerserver.Serve(context.Background(), provider.New(version), providerserver.ServeOpts{
		Address:         "registry.terraform.io/julienlevasseur/uname",
		ProtocolVersion: 5,
	})
	if err != nil {
		log.Fatal(err)
	}
}
