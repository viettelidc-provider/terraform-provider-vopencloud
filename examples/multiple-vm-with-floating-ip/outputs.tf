output "floating_ip_addresses" {
  value = vopencloud_networking_floatingip_v2.fip.*.address
}
