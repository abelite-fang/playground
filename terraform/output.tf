# generate inventory file for Ansible
resource "local_file" "ansible-inventory" {
  content = templatefile("${path.module}/inventory.tmpl",
    {
      test_instance = aws_instance.elb-testing.*.private_ip
    }
  )
  filename = "../ansible/inventory"
}
