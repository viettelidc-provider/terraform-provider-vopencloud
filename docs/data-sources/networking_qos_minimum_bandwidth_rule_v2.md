---
subcategory: "Networking / Neutron"
layout: "openstack"
page_title: "ViettelIdc: viettelidc_networking_qos_minimum_bandwidth_rule_v2"
sidebar_current: "docs-openstack-datasource-networking-qos-minimum-bandwidth-rule-v2"
description: |-
  Get information on an ViettelIdc QoS minimum bandwidth rule.
---

# viettelidc\_networking\_qos\_minimum\_bandwidth\_rule\_v2

Use this data source to get the ID of an available ViettelIdc QoS minimum bandwidth rule.

## Example Usage

```hcl
data "viettelidc_networking_qos_minimum_bandwidth_rule_v2" "qos_min_bw_rule_1" {
  min_kbps = 2000
}
```

## Argument Reference

* `region` - (Optional) The region in which to obtain the V2 Networking client.
    A Networking client is needed to create a Neutron QoS minimum bandwidth rule. If omitted, the
    `region` argument of the provider is used.

* `qos_policy_id` - (Required) The QoS policy reference.

* `min_kbps` - (Optional) The value of a minimum kbps bandwidth.


## Attributes Reference

`id` is set to the `qos_policy_id/minimum_bandwidth_rule_id` format of the found QoS minimum bandwidth rule.
In addition, the following attributes are exported:

* `region` - See Argument Reference above.
* `qos_policy_id` - See Argument Reference above.
* `min_kbps` - See Argument Reference above.