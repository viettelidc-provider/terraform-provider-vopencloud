---
subcategory: "Compute / Nova"
layout: "openstack"
page_title: "ViettelIdc: viettelidc_compute_aggregate_v2"
sidebar_current: "docs-openstack-datasource-compute-aggregate-v2"
description: |-
  Get information on Openstack Host Aggregate
---

# viettelidc\_compute\_aggregate\_v2

Use this data source to get information about host aggregates
by name.

## Example Usage

```hcl
data "viettelidc_compute_aggregate_v2" "test" {
  name = "test"
}
```

## Argument Reference

* `name` - The name of the host aggregate

## Attributes Reference

`id` is set to the ID of the found Host Aggregate. In addition, the
following attributes are exported:

* `name` - See Argument Reference above.
* `zone` - Availability zone of the Host Aggregate
* `metadata` - Metadata of the Host Aggregate
* `hosts` - List of Hypervisors contained in the Host Aggregate
