variable "env" {
  type = string
}

variable "disk" {
  type = map
  default = {
    dev = "1",
    prod = "3"
  }
}


resource "null_resource" "one" {
provisioner "local-exec" {
  command = "echo 'disk: ${lookup(var.disk,var.env)}'"
 }
}
