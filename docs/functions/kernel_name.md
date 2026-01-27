---
layout: "uname"
page_title: "Uname: uname_functions"
sidebar_current: "docs-uname-functions-kernel_name"
description: |-
   Provides details about uname kernel_name function
---

# Function: kernel_name

Provides host's Kernel name.

## Example Usage

```hcl
output "kernel_name" {
    value = provider::uname::kernel_name()
}
```

## Signature

```hcl
kernel_name()
```
