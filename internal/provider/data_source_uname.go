package provider

import (
	"strings"
	"syscall"

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

func get_operating_system() string {
	return goInfo.GetInfo().OS
}

func arrayToString(x [65]int8) string {
	var buf [65]byte
	for i, b := range x {
		buf[i] = byte(b)
	}
	str := string(buf[:])
	if i := strings.Index(str, "\x00"); i != -1 {
		str = str[:i]
	}
	return str
}

func dataSourceUnameRead(d *schema.ResourceData, _ interface{}) error {
	// syscall.Uname is not supported by MacOS, so, for Darwin OS, we leverage
	// goInfo package to build Uname informations:
	if get_operating_system() == "Darwin" {

		d.Set("kernel_name", goInfo.GetInfo().Kernel)
		d.Set("nodename", goInfo.GetInfo().Hostname)
		d.Set("kernel_release", goInfo.GetInfo().Core)
		d.Set("machine", goInfo.GetInfo().Platform)
		d.Set("operating_system", goInfo.GetInfo().OS)

		return nil

	} else {

		var uname syscall.Utsname
		if err := syscall.Uname(&uname); err != nil {
			return err
		}

		d.Set("kernel_name", arrayToString(uname.Release))
		d.Set("nodename", arrayToString(uname.Nodename))
		d.Set("kernel_release", arrayToString(uname.Version))
		d.Set("machine", arrayToString(uname.Machine))
		d.Set("operating_system", arrayToString(uname.Nodename))

		return nil
	}

}
