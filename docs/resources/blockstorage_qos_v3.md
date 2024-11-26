---
subcategory: "Block Storage / Cinder"
layout: "openstack"
page_title: "ViettelIdc: viettelidc_blockstorage_qos_v3"
sidebar_current: "docs-openstack-resource-blockstorage-qos-v3"
description: |-
  Manages a V3 Quality-Of-Servirce (qos) resource within ViettelIdc.
---

# viettelidc\_blockstorage\_qos\_v3

Manages a V3 block storage Quality-Of-Servirce (qos) resource within ViettelIdc.

~> **Note:** This usually requires admin privileges.


## Example Usage

```hcl
resource "viettelidc_blockstorage_qos_v3" "qos" {
  name = "foo"
  consumer = "back-end"
  specs = {
		read_iops_sec  = "40000"
		write_iops_sec = "40000"
	}
}

```

## Argument Reference

The following arguments are supported:

* `region` - (Optional) The region in which to create the qos. If omitted,
    the `region` argument of the provider is used. Changing this creates
    a new qos.

* `name` - (Required) Name of the qos.  Changing this creates a new qos.

* `consumer` - (Optional) The consumer of qos. Can be one of `front-end`,
    `back-end` or `both`. Changing this updates the `consumer` of an
    existing qos.

* `specs` - (Optional) Key/Value pairs of specs for the qos.

## Attributes Reference

The following attributes are exported:

* `region` - See Argument Reference above.
* `name` - See Argument Reference above.
* `consumer` - See Argument Reference above.
* `specs` - See Argument Reference above.

## Import

Qos can be imported using the `qos_id`, e.g.

```
$ terraform import viettelidc_blockstorage_qos_v3.qos 941793f0-0a34-4bc4-b72e-a6326ae58283
```
