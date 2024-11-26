---
subcategory: "Identity / Keystone"
layout: "openstack"
page_title: "ViettelIdc: viettelidc_identity_role_assignment_v3"
sidebar_current: "docs-openstack-resource-identity-role-assignment-v3"
description: |-
  Manages a V3 Role assignment within ViettelIdc Keystone.
---

# viettelidc\_identity\_role\_assignment\_v3

Manages a V3 Role assignment within ViettelIdc Keystone.

~> **Note:** You _must_ have admin privileges in your ViettelIdc cloud to use
this resource.

## Example Usage

```hcl
resource "viettelidc_identity_project_v3" "project_1" {
  name = "project_1"
}

resource "viettelidc_identity_user_v3" "user_1" {
  name               = "user_1"
  default_project_id = viettelidc_identity_project_v3.project_1.id
}

resource "viettelidc_identity_role_v3" "role_1" {
  name = "role_1"
}

resource "viettelidc_identity_role_assignment_v3" "role_assignment_1" {
  user_id    = viettelidc_identity_user_v3.user_1.id
  project_id = viettelidc_identity_project_v3.project_1.id
  role_id    = viettelidc_identity_role_v3.role_1.id
}
```

## Argument Reference

The following arguments are supported:

* `domain_id` - (Optional; Required if `project_id` is empty) The domain to assign the role in.

* `group_id` - (Optional; Required if `user_id` is empty) The group to assign the role to.

* `project_id` - (Optional; Required if `domain_id` is empty) The project to assign the role in.

* `user_id` - (Optional; Required if `group_id` is empty) The user to assign the role to.

* `role_id` - (Required) The role to assign.

## Attributes Reference

The following attributes are exported:

* `domain_id` - See Argument Reference above.
* `project_id` - See Argument Reference above.
* `group_id` - See Argument Reference above.
* `user_id` - See Argument Reference above.
* `role_id` - See Argument Reference above.

## Import

Role assignments can be imported using a constructed id. The id should have the form of
`domainID/projectID/groupID/userID/roleID`. When something is not used then leave blank.

For example this will import the role assignment for: 
projectID: 014395cd-89fc-4c9b-96b7-13d1ee79dad2,
userID: 4142e64b-1b35-44a0-9b1e-5affc7af1106,
roleID: ea257959-eeb1-4c10-8d33-26f0409a755d
( domainID and groupID are left blank)

```
$ terraform import viettelidc_identity_role_assignment_v3.role_assignment_1 /014395cd-89fc-4c9b-96b7-13d1ee79dad2//4142e64b-1b35-44a0-9b1e-5affc7af1106/ea257959-eeb1-4c10-8d33-26f0409a755d
```
