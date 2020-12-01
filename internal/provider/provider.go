package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func New() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Resource{
			DataSourcesMap: map[string]*schema.Resource{
				"uname": dataSourceUname(),
			},
		},
	}
}
