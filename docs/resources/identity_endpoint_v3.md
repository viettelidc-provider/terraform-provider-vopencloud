---
subcategory: "Identity / Keystone"
layout: "openstack"
page_title: "VOpenCloud: vopencloud_identity_endpoint_v3"
sidebar_current: "docs-openstack-resource-identity-endpoint-v3"
description: |-
  Manages a V3 Endpoint resource within VOpenCloud Keystone.
---

# vopencloud\_identity\_endpoint\_v3

Manages a V3 Endpoint resource within VOpenCloud Keystone.

~> **Note:** This usually requires admin privileges.

## Example Usage

```hcl
resource "vopencloud_identity_service_v3" "service_1" {
  name = "my-service"
  type = "my-service-type"
}

resource "vopencloud_identity_endpoint_v3" "endpoint_1" {
  name            = "my-endpoint"
  service_id      = vopencloud_identity_service_v3.service_1.id
  endpoint_region = vopencloud_identity_service_v3.service_1.region
  url             = "http://my-endpoint"
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional) The region in which to obtain the V3 Keystone client.
  If omitted, the `region` argument of the provider is used.

* `name` - (Optional) The endpoint name.

* `endpoint_region` - (Required) The endpoint region. The `region` and
  `endpoint_region` can be different.

* `url` - (Required) The endpoint url.

* `interface` - (Optional) The endpoint interface. Valid values are `public`,
  `internal` and `admin`. Default value is `public`

* `service_id` - (Required) The endpoint service ID.

## Attributes Reference

`id` is set to the ID of the endpoint. In addition, the following attributes are
exported:

* `region` - See Argument Reference above.
* `name` - See Argument Reference above.
* `endpoint_region` - See Argument Reference above.
* `url` - See Argument Reference above.
* `interface` - See Argument Reference above.
* `service_id` - See Argument Reference above.
* `service_name` - The service name of the endpoint.
* `service_type` - The service type of the endpoint.

## Import

Endpoints can be imported using the `id`, e.g.

```
$ terraform import vopencloud_identity_endpoint_v3.endpoint_1 5392472b-106a-4845-90c6-7c8445f18770
```
