output "floating_ip" {
  value = vopencloud_networking_floatingip_v2.fip.address
}

output "volume_devices" {
  value = vopencloud_compute_volume_attach_v2.attached.device
}
