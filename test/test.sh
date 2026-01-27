#!/usr/bin/env bash

set -e

TF_VAR_kernel_name="$(uname -s)"
TF_VAR_node_name="$(uname -n)"
TF_VAR_kernel_release="$(uname -r)"
TF_VAR_machine="$(uname -m)"
TF_VAR_operating_system="$(uname -o)"

export TF_VAR_kernel_name
export TF_VAR_node_name
export TF_VAR_kernel_release
export TF_VAR_machine
export TF_VAR_operating_system

printf "\n"
printf "testing with terraform \n"
printf "\n"
if command -v terraform &> /dev/null; then
  terraform test
fi

printf "\n"
printf "testing with opentofu \n"
printf "\n"
if command -v tofu &> /dev/null; then
  tofu test
fi
