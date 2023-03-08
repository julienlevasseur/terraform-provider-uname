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

	goInfo, err := goInfo.GetInfo()
	if err != nil {
		return err
	}

	err = d.Set("kernel_name", goInfo.Kernel)
	if err != nil {
		return err
	}

	err = d.Set("nodename", goInfo.Hostname)
	if err != nil {
		return err
	}

	err = d.Set("kernel_release", goInfo.Core)
	if err != nil {
		return err
	}

	err = d.Set("machine", goInfo.Platform)
	if err != nil {
		return err
	}

	err = d.Set("operating_system", goInfo.OS)

	return nil

}
