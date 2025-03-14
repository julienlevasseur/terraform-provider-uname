data "uname" "localhost" {}

output "kernel_name" {
    value = data.uname.localhost.kernel_name
}

output "node_name" {
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
