# terraform-uname-provider test

This `test` provider is provided (no pun intended) for development purpose.
It allows testing the provider to ensure any modification behave as expected.

The test must be ran from a provided script to allow predictive outcome from the provider.

## Usage

```bash
bash test.sh
```

> **Note**
>
> For development testing, you will have to configure this provider for local registry override: [https://developer.hashicorp.com/terraform/tutorials/providers-plugin-framework/providers-plugin-framework-provider#prepare-terraform-for-local-provider-install](https://developer.hashicorp.com/terraform/tutorials/providers-plugin-framework/providers-plugin-framework-provider#prepare-terraform-for-local-provider-install)
>
> ~/.terraformrc example:
>
> ```hcl
> provider_installation {
>     dev_overrides {
>         "julienlevasseur/uname" = "/${path_to_repository}/terraform-provider-uname"
>     }
>
>     direct {}
> }
> ```