---
subcategory: "Identity / Keystone"
layout: "openstack"
page_title: "VOpenCloud: vopencloud_identity_role_v3"
sidebar_current: "docs-openstack-datasource-identity-role-v3"
description: |-
  Get information on an VOpenCloud Role.
---

# vopencloud\_identity\_role\_v3

Use this data source to get the ID of an VOpenCloud role.

## Example Usage

```hcl
data "vopencloud_identity_role_v3" "admin" {
  name = "admin"
}
```

## Argument Reference

* `name` - The name of the role.

* `domain_id` - (Optional) The domain the role belongs to.

* `region` - (Optional) The region in which to obtain the V3 Keystone client.
    If omitted, the `region` argument of the provider is used.


## Attributes Reference

`id` is set to the ID of the found role. In addition, the following attributes
are exported:

* `name` - See Argument Reference above.
* `domain_id` - See Argument Reference above.
* `region` - See Argument Reference above.
