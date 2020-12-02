---
layout: "uname"
page_title: "Uname: uname_data_source"
sidebar_current: "docs-uname-data-source-uname"
description: |-
   Provides details about uname data source
---

# uname

`uname` DataSource provide system informations.

## Example Usage

```hcl
data "uname" "localhost" {}

output "localhost_kernel_name" {
    value = data.uname.system.kernel_name
}

output "localhost_nodename" {
    value = data.uname.system.nodename
}

output "localhost_kernel_release" {
    value = data.uname.system.kernel_release
}

output "localhost_machine" {
    value = data.uname.system.machine
}

output "localhost_operating_system" {
    value = data.uname.system.operating_system
}
```

## Attribute Reference

* `kernel_name` - provide the kernel name
* `nodename` - provide the network node hostname
* `kernel_release` - provide the kernel release
* `machine` - provide the machine hardware name
* `operating_system` - provide the operating system