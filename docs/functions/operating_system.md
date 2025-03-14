---
layout: "uname"
page_title: "Uname: uname_functions"
sidebar_current: "docs-uname-functions-operating_system"
description: |-
   Provides details about uname operating_system function
---

# Function: operating_system

Provides host's operating system name.

## Example Usage

```hcl
output "operating_system" {
    value = provider::uname::operating_system()
}
```

## Signature

```hcl
operating_system()
```
