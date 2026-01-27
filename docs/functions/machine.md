---
layout: "uname"
page_title: "Uname: uname_functions"
sidebar_current: "docs-uname-functions-machine"
description: |-
   Provides details about uname machine function
---

# Function: machine

Provides host's machine architecture type.

## Example Usage

```hcl
output "machine" {
    value = provider::uname::machine()
}
```

## Signature

```hcl
machine()
```
