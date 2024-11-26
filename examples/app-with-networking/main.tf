resource "vopencloud_compute_keypair_v2" "terraform" {
  name       = "terraform"
  public_key = file("${var.ssh_key_file}.pub")
}

resource "vopencloud_networking_network_v2" "terraform" {
  name           = "terraform"
  admin_state_up = "true"
}

resource "vopencloud_networking_subnet_v2" "terraform" {
  name            = "terraform"
  network_id      = vopencloud_networking_network_v2.terraform.id
  cidr            = "10.0.0.0/24"
  ip_version      = 4
  dns_nameservers = ["8.8.8.8", "8.8.4.4"]
}

resource "vopencloud_networking_router_v2" "terraform" {
  name                = "terraform"
  admin_state_up      = "true"
  external_network_id = data.vopencloud_networking_network_v2.terraform.id
}

resource "vopencloud_networking_router_interface_v2" "terraform" {
  router_id = vopencloud_networking_router_v2.terraform.id
  subnet_id = vopencloud_networking_subnet_v2.terraform.id
}

resource "vopencloud_networking_secgroup_v2" "terraform" {
  name        = "terraform"
  description = "Security group for the Terraform example instances"
}

resource "vopencloud_networking_secgroup_rule_v2" "terraform_22" {
  direction         = "ingress"
  ethertype         = "IPv4"
  protocol          = "tcp"
  port_range_min    = 22
  port_range_max    = 22
  remote_ip_prefix  = "0.0.0.0/0"
  security_group_id = vopencloud_networking_secgroup_v2.terraform.id
}

resource "vopencloud_networking_secgroup_rule_v2" "terraform_80" {
  direction         = "ingress"
  ethertype         = "IPv4"
  protocol          = "tcp"
  port_range_min    = 80
  port_range_max    = 80
  remote_ip_prefix  = "0.0.0.0/0"
  security_group_id = vopencloud_networking_secgroup_v2.terraform.id
}

resource "vopencloud_networking_secgroup_rule_v2" "terraform" {
  direction         = "ingress"
  ethertype         = "IPv4"
  protocol          = "icmp"
  remote_ip_prefix  = "0.0.0.0/0"
  security_group_id = vopencloud_networking_secgroup_v2.terraform.id
}

resource "vopencloud_networking_floatingip_v2" "terraform" {
  pool = var.pool
}

resource "vopencloud_compute_instance_v2" "terraform" {
  name            = "terraform"
  image_name      = var.image
  flavor_name     = var.flavor
  key_pair        = vopencloud_compute_keypair_v2.terraform.name
  security_groups = ["${vopencloud_networking_secgroup_v2.terraform.name}"]

  network {
    uuid = vopencloud_networking_network_v2.terraform.id
  }
}

resource "vopencloud_compute_floatingip_associate_v2" "terraform" {
  floating_ip = vopencloud_networking_floatingip_v2.terraform.address
  instance_id = vopencloud_compute_instance_v2.terraform.id

  provisioner "remote-exec" {
    connection {
      host        = vopencloud_networking_floatingip_v2.terraform.address
      user        = var.ssh_user_name
      private_key = file(var.ssh_key_file)
    }

    inline = [
      "sudo apt-get -y update",
      "sudo apt-get -y install nginx",
      "sudo service nginx start",
    ]
  }
}
