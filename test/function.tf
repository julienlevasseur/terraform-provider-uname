output "func_kernel_name" {
    value = provider::uname::kernel_name()
}

output "func_node_name" {
    value = provider::uname::node_name()
}

output "func_kernel_release" {
    value = provider::uname::kernel_release()
}

output "func_machine" {
    value = provider::uname::machine()
}

output "func_operating_system" {
    value = provider::uname::operating_system()
}
