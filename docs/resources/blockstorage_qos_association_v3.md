---
subcategory: "Block Storage / Cinder"
layout: "openstack"
page_title: "ViettelIdc: viettelidc_blockstorage_qos_association_v3"
sidebar_current: "docs-openstack-resource-blockstorage-qos-association-v3"
description: |-
  Manages a V3 Qos association resource within ViettelIdc.
---

# viettelidc\_blockstorage\_qos\_association\_v3

Manages a V3 block storage Qos Association resource within ViettelIdc.

~> **Note:** This usually requires admin privileges.


## Example Usage

```hcl
resource "viettelidc_blockstorage_qos_v3" "qos" {
  name = "%s"
  consumer = "front-end"
  specs = {
	  read_iops_sec = "20000"
  }
}

resource "viettelidc_blockstorage_volume_type_v3" "volume_type" {
  name = "%s"
}

resource "viettelidc_blockstorage_qos_association_v3" "qos_association" {
  qos_id         = viettelidc_blockstorage_qos_v3.qos.id
  volume_type_id = viettelidc_blockstorage_volume_type_v3.volume_type.id
}

```

## Argument Reference

The following arguments are supported:

* `region` - (Optional) The region in which to create the qos association.
    If omitted, the `region` argument of the provider is used. Changing
    this creates a new qos association.

* `qos_id` - (Required) ID of the qos to associate. Changing this creates
    a new qos association.

* `volume_type_id` - (Required) ID of the volume_type to associate.
    Changing this creates a new qos association.

## Attributes Reference

The following attributes are exported:

* `region` - See Argument Reference above.
* `qos_id` - See Argument Reference above.
* `volume_type_id` - See Argument Reference above.

## Import

Qos association can be imported using the `qos_id/volume_type_id`, e.g.

```
$ terraform import viettelidc_blockstorage_qos_association_v3.qos_association 941793f0-0a34-4bc4-b72e-a6326ae58283/ea257959-eeb1-4c10-8d33-26f0409a755d
```
