resource "null_resource" "one" {
 provisioner "local-exec" {
   command = "echo 'one one'"
 }
 provisioner "local-exec" {
   command = "echo 'one two'"
 }
}

resource "null_resource" "two" {
 depends_on = ["null_resource.one"]
 provisioner "local-exec" {
   command = "echo 'two'"
 }
}

resource "null_resource" "three" {
 depends_on = ["null_resource.two"]
 provisioner "local-exec" {
   command = "echo 'three'"
 }
}

output "outcome_is" {
value = null_resource.one
description = "will show if of specific resource run"
}
