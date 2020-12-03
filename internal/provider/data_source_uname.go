package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/matishsiao/goInfo"
)

func dataSourceUname() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceUnameRead,

		Schema: map[string]*schema.Schema{
			"kernel_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"nodename": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"kernel_release": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"machine": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"operating_system": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceUnameRead(d *schema.ResourceData, _ interface{}) error {
	//

	d.Set("kernel_name", goInfo.GetInfo().Kernel)
	d.Set("nodename", goInfo.GetInfo().Hostname)
	d.Set("kernel_release", goInfo.GetInfo().Core)
	d.Set("machine", goInfo.GetInfo().Platform)
	d.Set("operating_system", goInfo.GetInfo().OS)

	return nil

}
