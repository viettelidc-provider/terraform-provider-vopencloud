---
subcategory: "Identity / Keystone"
layout: "openstack"
page_title: "ViettelIdc: viettelidc_identity_user_membership_v3"
sidebar_current: "docs-openstack-resource-identity-user-membership-v3"
description: |-
  Manages a user membership to group V3 resource within ViettelIdc.
---

# viettelidc\_identity\_user\_membership\_v3

Manages a user membership to group V3 resource within ViettelIdc.

~> **Note:** You _must_ have admin privileges in your ViettelIdc cloud to use
this resource.

---

## Example Usage

```hcl
resource "viettelidc_identity_project_v3" "project_1" {
  name = "project_1"
}

resource "viettelidc_identity_user_v3" "user_1" {
  name               = "user_1"
  default_project_id = viettelidc_identity_project_v3.project_1.id
}

resource "viettelidc_identity_group_v3" "group_1" {
  name        = "group_1"
  description = "group 1"
}

resource "viettelidc_identity_role_v3" "role_1" {
  name = "role_1"
}

resource "viettelidc_identity_user_membership_v3" "user_membership_1" {
  user_id  = viettelidc_identity_user_v3.user_1.id
  group_id = viettelidc_identity_group_v3.group_1.id
}

resource "viettelidc_identity_role_assignment_v3" "role_assignment_1" {
  group_id   = viettelidc_identity_group_v3.group_1.id
  project_id = viettelidc_identity_project_v3.project_1.id
  role_id    = viettelidc_identity_role_v3.role_1.id
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional) The region in which to obtain the V3 Identity client.
  If omitted, the `region` argument of the provider is used.
  Changing this creates a new user membership.

* `user_id` - (Required) The UUID of user to use. Changing this creates a new user membership.

* `group_id` - (Required) The UUID of group to which the user will be added.
  Changing this creates a new user membership.

## Attributes Reference

The following attributes are exported:

* `region` - See Argument Reference above.
* `user_id` - See Argument Reference above.
* `group_id` - See Argument Reference above.

## Import

This resource can be imported by specifying all two arguments, separated
by a forward slash:

```
$ terraform import viettelidc_identity_user_membership_v3.user_membership_1 user_id/group_id
```
