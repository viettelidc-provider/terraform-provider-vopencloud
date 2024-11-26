---
subcategory: "Networking / Neutron"
layout: "openstack"
page_title: "ViettelIdc: viettelidc_networking_router_route_v2"
sidebar_current: "docs-openstack-resource-networking-router-route-v2"
description: |-
  Creates a routing entry on a ViettelIdc V2 router.
---

# viettelidc\_networking\_router\_route\_v2

Creates a routing entry on a ViettelIdc V2 router.

## Example Usage

```hcl
resource "viettelidc_networking_router_v2" "router_1" {
  name           = "router_1"
  admin_state_up = "true"
}

resource "viettelidc_networking_network_v2" "network_1" {
  name           = "network_1"
  admin_state_up = "true"
}

resource "viettelidc_networking_subnet_v2" "subnet_1" {
  network_id = viettelidc_networking_network_v2.network_1.id
  cidr       = "192.168.199.0/24"
  ip_version = 4
}

resource "viettelidc_networking_router_interface_v2" "int_1" {
  router_id = viettelidc_networking_router_v2.router_1.id
  subnet_id = viettelidc_networking_subnet_v2.subnet_1.id
}

resource "viettelidc_networking_router_route_v2" "router_route_1" {
  depends_on       = ["viettelidc_networking_router_interface_v2.int_1"]
  router_id        = viettelidc_networking_router_v2.router_1.id
  destination_cidr = "10.0.1.0/24"
  next_hop         = "192.168.199.254"
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional) The region in which to obtain the V2 networking client.
    A networking client is needed to configure a routing entry on a router. If omitted, the
    `region` argument of the provider is used. Changing this creates a new
    routing entry.

* `router_id` - (Required) ID of the router this routing entry belongs to. Changing
    this creates a new routing entry.

* `destination_cidr` - (Required) CIDR block to match on the packet’s destination IP. Changing
    this creates a new routing entry.

* `next_hop` - (Required) IP address of the next hop gateway.  Changing
    this creates a new routing entry.

## Attributes Reference

The following attributes are exported:

* `region` - See Argument Reference above.
* `router_id` - See Argument Reference above.
* `destination_cidr` - See Argument Reference above.
* `next_hop` - See Argument Reference above.

## Notes

The `next_hop` IP address must be directly reachable from the router at the ``viettelidc_networking_router_route_v2``
resource creation time.  You can ensure that by explicitly specifying a dependency on the ``viettelidc_networking_router_interface_v2``
resource that connects the next hop to the router, as in the example above.

## Import

Routing entries can be imported using a combined ID using the following format: ``<router_id>-route-<destination_cidr>-<next_hop>``

```
$ terraform import viettelidc_networking_router_route_v2.router_route_1 686fe248-386c-4f70-9f6c-281607dad079-route-10.0.1.0/24-192.168.199.25
```