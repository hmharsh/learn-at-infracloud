resource "random_id" "age" {
 byte_length = 2
}

output "age" {
 value = random_id.age.dec
}
