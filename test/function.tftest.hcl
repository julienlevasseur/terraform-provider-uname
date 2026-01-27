run "valid_function" {
    command = plan

    assert {
        condition = output.func_kernel_name == var.kernel_name
        error_message = "unexpected kernel name"
    }

    assert {
        condition = output.func_node_name == var.node_name
        error_message = "unexpected node name"
    }

    assert {
        condition = output.func_kernel_release == var.kernel_release
        error_message = "unexpected kernel release"
    }

    assert {
        condition = output.func_machine ==var.machine
        error_message = "unexpected machine"
    }

    assert {
        condition = output.func_operating_system ==var.operating_system
        error_message = "unexpected operating system"
    }
}
