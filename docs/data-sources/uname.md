---
layout: "uname"
page_title: "Uname: uname_data_source"
sidebar_current: "docs-uname-data-source-uname"
description: |-
   Provides details about uname data source
---

# uname

`uname` DataSource provide system information.

## Example Usage

```hcl
data "uname" "localhost" {}

output "kernel_name" {
    value = data.uname.localhost.kernel_name
}

output "nodename" {
    value = data.uname.localhost.node_name
}

output "kernel_release" {
    value = data.uname.localhost.kernel_release
}

output "machine" {
    value = data.uname.localhost.machine
}

output "operating_system" {
    value = data.uname.localhost.operating_system
}
```

## Attribute Reference

* `kernel_name` - provide the kernel name
* `node_name` - provide the network node hostname
* `kernel_release` - provide the kernel release
* `machine` - provide the machine hardware name
* `operating_system` - provide the operating system

## Functions

From version 0.3.0 the `uname` provider offer support of built-in functions.

## Example Usage

```hcl
output "kernel_name" {
    value = provider::uname::kernel_name()
}

output "node_name" {
    value = provider::uname::node_name()
}

output "kernel_release" {
    value = provider::uname::kernel_release()
}

output "machine" {
    value = provider::uname::machine()
}

output "operating_system" {
    value = provider::uname::operating_system()
}
```
