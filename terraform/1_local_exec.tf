
// Provisioners can be used to model specific actions on the local machine or on a remote machine in order to prepare servers or other infrastructure objects for service.

# If depends_on not used than all resource will run in parallel
resource "null_resource" "first" {
    provisioner "local-exec" {
        command = "echo 'first'"
    }
}

resource "null_resource" "second" {
    depends_on = ["null_resource.first"]
    provisioner "local-exec" {
        command = "echo 'second'"
    }
}

resource "null_resource" "third" {
    depends_on = ["null_resource.second"]
    provisioner "local-exec" {
        command = "echo 'third'"
    }
}


# one resource can also contain more than one provisioner also

/*
  
resource "aws_instance" "web" {
  # ...

  provisioner "local-exec" {
    command    = "echo The server's IP address is ${self.private_ip}"
    on_failure = continue  // or 'fail'
  }
}

*/
