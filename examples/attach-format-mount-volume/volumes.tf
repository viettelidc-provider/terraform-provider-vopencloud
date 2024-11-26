resource "vopencloud_blockstorage_volume_v2" "volume_1" {
  name = "volume_1"
  size = var.volume_size
}

resource "vopencloud_compute_volume_attach_v2" "attached" {
  instance_id = vopencloud_compute_instance_v2.my_instance.id
  volume_id   = vopencloud_blockstorage_volume_v2.volume_1.id
  # Prevent re-creation
  #   lifecycle {
  #     ignore_changes = ["volume_id", "instance_id"]
  #   }
}
