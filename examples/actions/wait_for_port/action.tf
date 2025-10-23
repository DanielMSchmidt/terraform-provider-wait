action "wait_for_port" "ansible_host" {
  config {
    host = aws_instance.web.public_ip
    port = 22
  }
}


resource "aws_instance" "web" {
  // Example usage of the action

  lifecycle {
    action_trigger {
      events  = [after_create, after_update]
      actions = [action.wait_for_port.ansible_host, action.ansible_playbook.ansible]
    }
  }
}
