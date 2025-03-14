---
layout: "uname"
page_title: "Uname: uname_functions"
sidebar_current: "docs-uname-functions-node-name"
description: |-
   Provides details about uname node_name function
---

# Function: node_name

Provides host's network node hostname.

## Example Usage

```hcl
output "node_name" {
    value = provider::uname::node_name()
}
```

## Signature

```hcl
node_name()
```
