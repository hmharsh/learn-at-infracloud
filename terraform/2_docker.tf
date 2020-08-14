# Configure the Docker provider
provider "docker" {}

# Create a container
resource "docker_container" "nginx" {
  image = "nginx"
  name  = "nginx"
}
