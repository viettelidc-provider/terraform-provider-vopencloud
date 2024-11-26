package viettelidc

import (
	"context"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/layer3/routers"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/ports"
)

func resourceNetworkingRouterInterfaceV2() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceNetworkingRouterInterfaceV2Create,
		ReadContext:   resourceNetworkingRouterInterfaceV2Read,
		UpdateContext: resourceNetworkingRouterInterfaceV2Update,
		DeleteContext: resourceNetworkingRouterInterfaceV2Delete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"region": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			"router_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"subnet_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			"port_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			"force_destroy": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func resourceNetworkingRouterInterfaceV2Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*Config)
	networkingClient, err := config.NetworkingV2Client(GetRegion(d, config))
	if err != nil {
		return diag.Errorf("Error creating OpenStack networking client: %s", err)
	}

	createOpts := routers.AddInterfaceOpts{
		SubnetID: d.Get("subnet_id").(string),
		PortID:   d.Get("port_id").(string),
	}

	log.Printf("[DEBUG] openstack_networking_router_interface_v2 create options: %#v", createOpts)
	r, err := routers.AddInterface(networkingClient, d.Get("router_id").(string), createOpts).Extract()
	if err != nil {
		return diag.Errorf("Error creating openstack_networking_router_interface_v2: %s", err)
	}

	log.Printf("[DEBUG] Waiting for openstack_networking_router_interface_v2 %s to become available", r.PortID)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"BUILD", "PENDING_CREATE", "PENDING_UPDATE"},
		Target:     []string{"ACTIVE", "DOWN"},
		Refresh:    resourceNetworkingRouterInterfaceV2StateRefreshFunc(networkingClient, r.PortID),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      5 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	_, err = stateConf.WaitForStateContext(ctx)
	if err != nil {
		return diag.Errorf("Error waiting for openstack_networking_router_interface_v2 %s to become available: %s", r.ID, err)
	}

	d.SetId(r.PortID)

	log.Printf("[DEBUG] Created openstack_networking_router_interface_v2 %s: %#v", r.ID, r)
	return resourceNetworkingRouterInterfaceV2Read(ctx, d, meta)
}

func resourceNetworkingRouterInterfaceV2Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*Config)
	networkingClient, err := config.NetworkingV2Client(GetRegion(d, config))
	if err != nil {
		return diag.Errorf("Error creating OpenStack networking client: %s", err)
	}

	r, err := ports.Get(networkingClient, d.Id()).Extract()
	if err != nil {
		if _, ok := err.(gophercloud.ErrDefault404); ok {
			d.SetId("")
			return nil
		}

		return diag.Errorf("Error retrieving openstack_networking_router_interface_v2: %s", err)
	}

	log.Printf("[DEBUG] Retrieved openstack_networking_router_interface_v2 %s: %#v", d.Id(), r)

	d.Set("router_id", r.DeviceID)
	d.Set("port_id", r.ID)
	d.Set("region", GetRegion(d, config))

	// Set the subnet ID by looking at the port's FixedIPs.
	// If there's more than one FixedIP, do not set the subnet
	// as it's not possible to confidently determine which subnet
	// belongs to this interface. However, that situation should
	// not happen.
	if len(r.FixedIPs) != 1 {
		log.Printf("[DEBUG] Unable to set openstack_networking_router_interface_v2 %s subnet_id", d.Id())
	} else {
		d.Set("subnet_id", r.FixedIPs[0].SubnetID)
	}

	return nil
}

func resourceNetworkingRouterInterfaceV2Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceNetworkingRouterInterfaceV2Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*Config)
	networkingClient, err := config.NetworkingV2Client(GetRegion(d, config))
	if err != nil {
		return diag.Errorf("Error creating OpenStack networking client: %s", err)
	}

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"ACTIVE"},
		Target:     []string{"DELETED"},
		Refresh:    resourceNetworkingRouterInterfaceV2DeleteRefreshFunc(networkingClient, d),
		Timeout:    d.Timeout(schema.TimeoutDelete),
		Delay:      5 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	_, err = stateConf.WaitForStateContext(ctx)
	if err != nil {
		return diag.Errorf("Error waiting for openstack_networking_router_interface_v2 %s to Delete:  %s", d.Id(), err)
	}

	d.SetId("")
	return nil
}
