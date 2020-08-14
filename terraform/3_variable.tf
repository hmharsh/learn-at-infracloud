# declared variable can also be supplied by, terraform apply -var 'name=harshit'
variable "name" {
  type = string
  default = "harshit mahajan"  // will not ask for value of default is there at time of terraform apply 
}

resource "null_resource" "one" {
provisioner "local-exec" {
  command = "echo 'Hello ${var.name}'"
}
}
