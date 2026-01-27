---
layout: "uname"
page_title: "Uname: uname_functions"
sidebar_current: "docs-uname-functions-kernel_release"
description: |-
   Provides details about uname kernel_release function
---

# Function: kernel_release

Provides host's Kernel release.

## Example Usage

```hcl
output "kernel_release" {
    value = provider::uname::kernel_release()
}
```

## Signature

```hcl
kernel_release()
```
