package vopencloud

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/logging"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/meta"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/utils/terraform/auth"
	"github.com/gophercloud/utils/terraform/mutexkv"
)

// Use openstackbase.Config as the base/foundation of this provider's
// Config struct.
type Config struct {
	auth.Config
}

// Provider returns a schema.Provider for OpenStack.
func Provider() *schema.Provider {
	provider := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"auth_url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_AUTH_URL", ""),
				Description: descriptions["auth_url"],
			},

			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["region"],
				DefaultFunc: schema.EnvDefaultFunc("OS_REGION_NAME", ""),
			},

			"user_name": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_USERNAME", ""),
				Description: descriptions["user_name"],
			},

			"user_id": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_USER_ID", ""),
				Description: descriptions["user_id"],
			},

			"application_credential_id": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_APPLICATION_CREDENTIAL_ID", ""),
				Description: descriptions["application_credential_id"],
			},

			"application_credential_name": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_APPLICATION_CREDENTIAL_NAME", ""),
				Description: descriptions["application_credential_name"],
			},

			"application_credential_secret": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_APPLICATION_CREDENTIAL_SECRET", ""),
				Description: descriptions["application_credential_secret"],
			},

			"tenant_id": {
				Type:     schema.TypeString,
				Optional: true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"OS_TENANT_ID",
					"OS_PROJECT_ID",
				}, ""),
				Description: descriptions["tenant_id"],
			},

			"tenant_name": {
				Type:     schema.TypeString,
				Optional: true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"OS_TENANT_NAME",
					"OS_PROJECT_NAME",
				}, ""),
				Description: descriptions["tenant_name"],
			},

			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("OS_PASSWORD", ""),
				Description: descriptions["password"],
			},

			"token": {
				Type:     schema.TypeString,
				Optional: true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"OS_TOKEN",
					"OS_AUTH_TOKEN",
				}, ""),
				Description: descriptions["token"],
			},

			"user_domain_name": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_USER_DOMAIN_NAME", ""),
				Description: descriptions["user_domain_name"],
			},

			"user_domain_id": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_USER_DOMAIN_ID", ""),
				Description: descriptions["user_domain_id"],
			},

			"project_domain_name": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_PROJECT_DOMAIN_NAME", ""),
				Description: descriptions["project_domain_name"],
			},

			"project_domain_id": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_PROJECT_DOMAIN_ID", ""),
				Description: descriptions["project_domain_id"],
			},

			"domain_id": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_DOMAIN_ID", ""),
				Description: descriptions["domain_id"],
			},

			"domain_name": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_DOMAIN_NAME", ""),
				Description: descriptions["domain_name"],
			},

			"default_domain": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_DEFAULT_DOMAIN", "default"),
				Description: descriptions["default_domain"],
			},

			"system_scope": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_SYSTEM_SCOPE", false),
				Description: descriptions["system_scope"],
			},

			"insecure": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_INSECURE", nil),
				Description: descriptions["insecure"],
			},

			"endpoint_type": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_ENDPOINT_TYPE", ""),
			},

			"cacert_file": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_CACERT", ""),
				Description: descriptions["cacert_file"],
			},

			"cert": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_CERT", ""),
				Description: descriptions["cert"],
			},

			"key": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_KEY", ""),
				Description: descriptions["key"],
			},

			"swauth": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_SWAUTH", false),
				Description: descriptions["swauth"],
			},

			"use_octavia": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_USE_OCTAVIA", true),
				Description: descriptions["use_octavia"],
			},

			"delayed_auth": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_DELAYED_AUTH", true),
				Description: descriptions["delayed_auth"],
			},

			"allow_reauth": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_ALLOW_REAUTH", true),
				Description: descriptions["allow_reauth"],
			},

			"cloud": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_CLOUD", ""),
				Description: descriptions["cloud"],
			},

			"max_retries": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: descriptions["max_retries"],
			},

			"endpoint_overrides": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: descriptions["endpoint_overrides"],
			},

			"disable_no_cache_header": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: descriptions["disable_no_cache_header"],
			},

			"enable_logging": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: descriptions["enable_logging"],
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"vopencloud_blockstorage_availability_zones_v3":       dataSourceBlockStorageAvailabilityZonesV3(),
			"vopencloud_blockstorage_snapshot_v2":                 dataSourceBlockStorageSnapshotV2(),
			"vopencloud_blockstorage_snapshot_v3":                 dataSourceBlockStorageSnapshotV3(),
			"vopencloud_blockstorage_volume_v2":                   dataSourceBlockStorageVolumeV2(),
			"vopencloud_blockstorage_volume_v3":                   dataSourceBlockStorageVolumeV3(),
			"vopencloud_blockstorage_quotaset_v3":                 dataSourceBlockStorageQuotasetV3(),
			"vopencloud_compute_aggregate_v2":                     dataSourceComputeAggregateV2(),
			"vopencloud_compute_availability_zones_v2":            dataSourceComputeAvailabilityZonesV2(),
			"vopencloud_compute_instance_v2":                      dataSourceComputeInstanceV2(),
			"vopencloud_compute_flavor_v2":                        dataSourceComputeFlavorV2(),
			"vopencloud_compute_hypervisor_v2":                    dataSourceComputeHypervisorV2(),
			"vopencloud_compute_keypair_v2":                       dataSourceComputeKeypairV2(),
			"vopencloud_compute_quotaset_v2":                      dataSourceComputeQuotasetV2(),
			"vopencloud_compute_limits_v2":                        dataSourceComputeLimitsV2(),
			"vopencloud_containerinfra_nodegroup_v1":              dataSourceContainerInfraNodeGroupV1(),
			"vopencloud_containerinfra_clustertemplate_v1":        dataSourceContainerInfraClusterTemplateV1(),
			"vopencloud_containerinfra_cluster_v1":                dataSourceContainerInfraCluster(),
			"vopencloud_dns_zone_v2":                              dataSourceDNSZoneV2(),
			"vopencloud_fw_group_v2":                              dataSourceFWGroupV2(),
			"vopencloud_fw_policy_v1":                             dataSourceFWPolicyV1(),
			"vopencloud_fw_policy_v2":                             dataSourceFWPolicyV2(),
			"vopencloud_fw_rule_v2":                               dataSourceFWRuleV2(),
			"vopencloud_identity_role_v3":                         dataSourceIdentityRoleV3(),
			"vopencloud_identity_project_v3":                      dataSourceIdentityProjectV3(),
			"vopencloud_identity_user_v3":                         dataSourceIdentityUserV3(),
			"vopencloud_identity_auth_scope_v3":                   dataSourceIdentityAuthScopeV3(),
			"vopencloud_identity_endpoint_v3":                     dataSourceIdentityEndpointV3(),
			"vopencloud_identity_service_v3":                      dataSourceIdentityServiceV3(),
			"vopencloud_identity_group_v3":                        dataSourceIdentityGroupV3(),
			"vopencloud_images_image_v2":                          dataSourceImagesImageV2(),
			"vopencloud_images_image_ids_v2":                      dataSourceImagesImageIDsV2(),
			"vopencloud_networking_addressscope_v2":               dataSourceNetworkingAddressScopeV2(),
			"vopencloud_networking_network_v2":                    dataSourceNetworkingNetworkV2(),
			"vopencloud_networking_qos_bandwidth_limit_rule_v2":   dataSourceNetworkingQoSBandwidthLimitRuleV2(),
			"vopencloud_networking_qos_dscp_marking_rule_v2":      dataSourceNetworkingQoSDSCPMarkingRuleV2(),
			"vopencloud_networking_qos_minimum_bandwidth_rule_v2": dataSourceNetworkingQoSMinimumBandwidthRuleV2(),
			"vopencloud_networking_qos_policy_v2":                 dataSourceNetworkingQoSPolicyV2(),
			"vopencloud_networking_quota_v2":                      dataSourceNetworkingQuotaV2(),
			"vopencloud_networking_subnet_v2":                     dataSourceNetworkingSubnetV2(),
			"vopencloud_networking_subnet_ids_v2":                 dataSourceNetworkingSubnetIDsV2(),
			"vopencloud_networking_secgroup_v2":                   dataSourceNetworkingSecGroupV2(),
			"vopencloud_networking_subnetpool_v2":                 dataSourceNetworkingSubnetPoolV2(),
			"vopencloud_networking_floatingip_v2":                 dataSourceNetworkingFloatingIPV2(),
			"vopencloud_networking_router_v2":                     dataSourceNetworkingRouterV2(),
			"vopencloud_networking_port_v2":                       dataSourceNetworkingPortV2(),
			"vopencloud_networking_port_ids_v2":                   dataSourceNetworkingPortIDsV2(),
			"vopencloud_networking_trunk_v2":                      dataSourceNetworkingTrunkV2(),
			"vopencloud_sharedfilesystem_availability_zones_v2":   dataSourceSharedFilesystemAvailabilityZonesV2(),
			"vopencloud_sharedfilesystem_sharenetwork_v2":         dataSourceSharedFilesystemShareNetworkV2(),
			"vopencloud_sharedfilesystem_share_v2":                dataSourceSharedFilesystemShareV2(),
			"vopencloud_sharedfilesystem_snapshot_v2":             dataSourceSharedFilesystemSnapshotV2(),
			"vopencloud_keymanager_secret_v1":                     dataSourceKeyManagerSecretV1(),
			"vopencloud_keymanager_container_v1":                  dataSourceKeyManagerContainerV1(),
			"vopencloud_kubernetes_v1":                            dataSourceKubernetesV1(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"vopencloud_blockstorage_qos_association_v3":          resourceBlockStorageQosAssociationV3(),
			"vopencloud_blockstorage_qos_v3":                      resourceBlockStorageQosV3(),
			"vopencloud_blockstorage_quotaset_v2":                 resourceBlockStorageQuotasetV2(),
			"vopencloud_blockstorage_quotaset_v3":                 resourceBlockStorageQuotasetV3(),
			"vopencloud_blockstorage_volume_v1":                   resourceBlockStorageVolumeV1(),
			"vopencloud_blockstorage_volume_v2":                   resourceBlockStorageVolumeV2(),
			"vopencloud_blockstorage_volume_v3":                   resourceBlockStorageVolumeV3(),
			"vopencloud_blockstorage_volume_attach_v2":            resourceBlockStorageVolumeAttachV2(),
			"vopencloud_blockstorage_volume_attach_v3":            resourceBlockStorageVolumeAttachV3(),
			"vopencloud_blockstorage_volume_type_access_v3":       resourceBlockstorageVolumeTypeAccessV3(),
			"vopencloud_blockstorage_volume_type_v3":              resourceBlockStorageVolumeTypeV3(),
			"vopencloud_compute_aggregate_v2":                     resourceComputeAggregateV2(),
			"vopencloud_compute_flavor_v2":                        resourceComputeFlavorV2(),
			"vopencloud_compute_flavor_access_v2":                 resourceComputeFlavorAccessV2(),
			"vopencloud_compute_instance_v2":                      resourceComputeInstanceV2(),
			"vopencloud_compute_interface_attach_v2":              resourceComputeInterfaceAttachV2(),
			"vopencloud_compute_keypair_v2":                       resourceComputeKeypairV2(),
			"vopencloud_compute_secgroup_v2":                      resourceComputeSecGroupV2(),
			"vopencloud_compute_servergroup_v2":                   resourceComputeServerGroupV2(),
			"vopencloud_compute_quotaset_v2":                      resourceComputeQuotasetV2(),
			"vopencloud_compute_floatingip_v2":                    resourceComputeFloatingIPV2(),
			"vopencloud_compute_floatingip_associate_v2":          resourceComputeFloatingIPAssociateV2(),
			"vopencloud_compute_volume_attach_v2":                 resourceComputeVolumeAttachV2(),
			"vopencloud_containerinfra_nodegroup_v1":              resourceContainerInfraNodeGroupV1(),
			"vopencloud_containerinfra_clustertemplate_v1":        resourceContainerInfraClusterTemplateV1(),
			"vopencloud_containerinfra_cluster_v1":                resourceContainerInfraClusterV1(),
			"vopencloud_db_instance_v1":                           resourceDatabaseInstanceV1(),
			"vopencloud_db_user_v1":                               resourceDatabaseUserV1(),
			"vopencloud_db_configuration_v1":                      resourceDatabaseConfigurationV1(),
			"vopencloud_db_database_v1":                           resourceDatabaseDatabaseV1(),
			"vopencloud_dns_recordset_v2":                         resourceDNSRecordSetV2(),
			"vopencloud_dns_zone_v2":                              resourceDNSZoneV2(),
			"vopencloud_dns_transfer_request_v2":                  resourceDNSTransferRequestV2(),
			"vopencloud_dns_transfer_accept_v2":                   resourceDNSTransferAcceptV2(),
			"vopencloud_fw_firewall_v1":                           resourceFWFirewallV1(),
			"vopencloud_fw_group_v2":                              resourceFWGroupV2(),
			"vopencloud_fw_policy_v1":                             resourceFWPolicyV1(),
			"vopencloud_fw_policy_v2":                             resourceFWPolicyV2(),
			"vopencloud_fw_rule_v1":                               resourceFWRuleV1(),
			"vopencloud_fw_rule_v2":                               resourceFWRuleV2(),
			"vopencloud_identity_endpoint_v3":                     resourceIdentityEndpointV3(),
			"vopencloud_identity_project_v3":                      resourceIdentityProjectV3(),
			"vopencloud_identity_role_v3":                         resourceIdentityRoleV3(),
			"vopencloud_identity_role_assignment_v3":              resourceIdentityRoleAssignmentV3(),
			"vopencloud_identity_inherit_role_assignment_v3":      resourceIdentityInheritRoleAssignmentV3(),
			"vopencloud_identity_service_v3":                      resourceIdentityServiceV3(),
			"vopencloud_identity_user_v3":                         resourceIdentityUserV3(),
			"vopencloud_identity_user_membership_v3":              resourceIdentityUserMembershipV3(),
			"vopencloud_identity_group_v3":                        resourceIdentityGroupV3(),
			"vopencloud_identity_application_credential_v3":       resourceIdentityApplicationCredentialV3(),
			"vopencloud_identity_ec2_credential_v3":               resourceIdentityEc2CredentialV3(),
			"vopencloud_images_image_v2":                          resourceImagesImageV2(),
			"vopencloud_images_image_access_v2":                   resourceImagesImageAccessV2(),
			"vopencloud_images_image_access_accept_v2":            resourceImagesImageAccessAcceptV2(),
			"vopencloud_lb_member_v1":                             resourceLBMemberV1(),
			"vopencloud_lb_monitor_v1":                            resourceLBMonitorV1(),
			"vopencloud_lb_pool_v1":                               resourceLBPoolV1(),
			"vopencloud_lb_vip_v1":                                resourceLBVipV1(),
			"vopencloud_lb_loadbalancer_v2":                       resourceLoadBalancerV2(),
			"vopencloud_lb_listener_v2":                           resourceListenerV2(),
			"vopencloud_lb_pool_v2":                               resourcePoolV2(),
			"vopencloud_lb_member_v2":                             resourceMemberV2(),
			"vopencloud_lb_members_v2":                            resourceMembersV2(),
			"vopencloud_lb_monitor_v2":                            resourceMonitorV2(),
			"vopencloud_lb_l7policy_v2":                           resourceL7PolicyV2(),
			"vopencloud_lb_l7rule_v2":                             resourceL7RuleV2(),
			"vopencloud_lb_quota_v2":                              resourceLoadBalancerQuotaV2(),
			"vopencloud_networking_floatingip_v2":                 resourceNetworkingFloatingIPV2(),
			"vopencloud_networking_floatingip_associate_v2":       resourceNetworkingFloatingIPAssociateV2(),
			"vopencloud_networking_network_v2":                    resourceNetworkingNetworkV2(),
			"vopencloud_networking_port_v2":                       resourceNetworkingPortV2(),
			"vopencloud_networking_rbac_policy_v2":                resourceNetworkingRBACPolicyV2(),
			"vopencloud_networking_port_secgroup_associate_v2":    resourceNetworkingPortSecGroupAssociateV2(),
			"vopencloud_networking_qos_bandwidth_limit_rule_v2":   resourceNetworkingQoSBandwidthLimitRuleV2(),
			"vopencloud_networking_qos_dscp_marking_rule_v2":      resourceNetworkingQoSDSCPMarkingRuleV2(),
			"vopencloud_networking_qos_minimum_bandwidth_rule_v2": resourceNetworkingQoSMinimumBandwidthRuleV2(),
			"vopencloud_networking_qos_policy_v2":                 resourceNetworkingQoSPolicyV2(),
			"vopencloud_networking_quota_v2":                      resourceNetworkingQuotaV2(),
			"vopencloud_networking_router_v2":                     resourceNetworkingRouterV2(),
			"vopencloud_networking_router_interface_v2":           resourceNetworkingRouterInterfaceV2(),
			"vopencloud_networking_router_route_v2":               resourceNetworkingRouterRouteV2(),
			"vopencloud_networking_secgroup_v2":                   resourceNetworkingSecGroupV2(),
			"vopencloud_networking_secgroup_rule_v2":              resourceNetworkingSecGroupRuleV2(),
			"vopencloud_networking_subnet_v2":                     resourceNetworkingSubnetV2(),
			"vopencloud_networking_subnet_route_v2":               resourceNetworkingSubnetRouteV2(),
			"vopencloud_networking_subnetpool_v2":                 resourceNetworkingSubnetPoolV2(),
			"vopencloud_networking_addressscope_v2":               resourceNetworkingAddressScopeV2(),
			"vopencloud_networking_trunk_v2":                      resourceNetworkingTrunkV2(),
			"vopencloud_networking_portforwarding_v2":             resourceNetworkingPortForwardingV2(),
			"vopencloud_objectstorage_container_v1":               resourceObjectStorageContainerV1(),
			"vopencloud_objectstorage_object_v1":                  resourceObjectStorageObjectV1(),
			"vopencloud_objectstorage_tempurl_v1":                 resourceObjectstorageTempurlV1(),
			"vopencloud_orchestration_stack_v1":                   resourceOrchestrationStackV1(),
			"vopencloud_vpnaas_ipsec_policy_v2":                   resourceIPSecPolicyV2(),
			"vopencloud_vpnaas_service_v2":                        resourceServiceV2(),
			"vopencloud_vpnaas_ike_policy_v2":                     resourceIKEPolicyV2(),
			"vopencloud_vpnaas_endpoint_group_v2":                 resourceEndpointGroupV2(),
			"vopencloud_vpnaas_site_connection_v2":                resourceSiteConnectionV2(),
			"vopencloud_sharedfilesystem_securityservice_v2":      resourceSharedFilesystemSecurityServiceV2(),
			"vopencloud_sharedfilesystem_sharenetwork_v2":         resourceSharedFilesystemShareNetworkV2(),
			"vopencloud_sharedfilesystem_share_v2":                resourceSharedFilesystemShareV2(),
			"vopencloud_sharedfilesystem_share_access_v2":         resourceSharedFilesystemShareAccessV2(),
			"vopencloud_keymanager_secret_v1":                     resourceKeyManagerSecretV1(),
			"vopencloud_keymanager_container_v1":                  resourceKeyManagerContainerV1(),
			"vopencloud_keymanager_order_v1":                      resourceKeyManagerOrderV1(),
			"vopencloud_kubernetes_v1":                            resourceKubernetesV1(),
		},
	}

	provider.ConfigureContextFunc = func(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		terraformVersion := provider.TerraformVersion
		if terraformVersion == "" {
			// Terraform 0.12 introduced this field to the protocol
			// We can therefore assume that if it's missing it's 0.10 or 0.11
			terraformVersion = "0.11+compatible"
		}
		return configureProvider(d, terraformVersion)
	}

	return provider
}

var descriptions map[string]string

func init() {
	descriptions = map[string]string{
		"auth_url": "The Identity authentication URL.",

		"cloud": "An entry in a `clouds.yaml` file to use.",

		"region": "The OpenStack region to connect to.",

		"user_name": "Username to login with.",

		"user_id": "User ID to login with.",

		"application_credential_id": "Application Credential ID to login with.",

		"application_credential_name": "Application Credential name to login with.",

		"application_credential_secret": "Application Credential secret to login with.",

		"tenant_id": "The ID of the Tenant (Identity v2) or Project (Identity v3)\n" +
			"to login with.",

		"tenant_name": "The name of the Tenant (Identity v2) or Project (Identity v3)\n" +
			"to login with.",

		"password": "Password to login with.",

		"token": "Authentication token to use as an alternative to username/password.",

		"user_domain_name": "The name of the domain where the user resides (Identity v3).",

		"user_domain_id": "The ID of the domain where the user resides (Identity v3).",

		"project_domain_name": "The name of the domain where the project resides (Identity v3).",

		"project_domain_id": "The ID of the domain where the proejct resides (Identity v3).",

		"domain_id": "The ID of the Domain to scope to (Identity v3).",

		"domain_name": "The name of the Domain to scope to (Identity v3).",

		"default_domain": "The name of the Domain ID to scope to if no other domain is specified. Defaults to `default` (Identity v3).",

		"system_scope": "If set to `true`, system scoped authorization will be enabled. Defaults to `false` (Identity v3).",

		"insecure": "Trust self-signed certificates.",

		"cacert_file": "A Custom CA certificate.",

		"cert": "A client certificate to authenticate with.",

		"key": "A client private key to authenticate with.",

		"endpoint_type": "The catalog endpoint type to use.",

		"endpoint_overrides": "A map of services with an endpoint to override what was\n" +
			"from the Keystone catalog",

		"swauth": "Use Swift's authentication system instead of Keystone. Only used for\n" +
			"interaction with Swift.",

		"use_octavia": "If set to `true`, API requests will go the Load Balancer\n" +
			"service (Octavia) instead of the Networking service (Neutron).",

		"disable_no_cache_header": "If set to `true`, the HTTP `Cache-Control: no-cache` header will not be added by default to all API requests.",

		"delayed_auth": "If set to `false`, OpenStack authorization will be perfomed,\n" +
			"every time the service provider client is called. Defaults to `true`.",

		"allow_reauth": "If set to `false`, OpenStack authorization won't be perfomed\n" +
			"automatically, if the initial auth token get expired. Defaults to `true`",

		"max_retries": "How many times HTTP connection should be retried until giving up.",

		"enable_logging": "Outputs very verbose logs with all calls made to and responses from OpenStack",
	}
}

func configureProvider(d *schema.ResourceData, terraformVersion string) (interface{}, diag.Diagnostics) {
	enableLogging := d.Get("enable_logging").(bool)
	if !enableLogging {
		// enforce logging (similar to OS_DEBUG) when TF_LOG is 'DEBUG' or 'TRACE'
		if logLevel := logging.LogLevel(); logLevel != "" && os.Getenv("OS_DEBUG") == "" {
			if logLevel == "DEBUG" || logLevel == "TRACE" {
				enableLogging = true
			}
		}
	}

	authOpts := &gophercloud.AuthOptions{
		Scope: &gophercloud.AuthScope{System: d.Get("system_scope").(bool)},
	}

	config := Config{
		auth.Config{
			CACertFile:                  d.Get("cacert_file").(string),
			ClientCertFile:              d.Get("cert").(string),
			ClientKeyFile:               d.Get("key").(string),
			Cloud:                       d.Get("cloud").(string),
			DefaultDomain:               d.Get("default_domain").(string),
			DomainID:                    d.Get("domain_id").(string),
			DomainName:                  d.Get("domain_name").(string),
			EndpointOverrides:           d.Get("endpoint_overrides").(map[string]interface{}),
			EndpointType:                d.Get("endpoint_type").(string),
			IdentityEndpoint:            d.Get("auth_url").(string),
			Password:                    d.Get("password").(string),
			ProjectDomainID:             d.Get("project_domain_id").(string),
			ProjectDomainName:           d.Get("project_domain_name").(string),
			Region:                      d.Get("region").(string),
			Swauth:                      d.Get("swauth").(bool),
			Token:                       d.Get("token").(string),
			TenantID:                    d.Get("tenant_id").(string),
			TenantName:                  d.Get("tenant_name").(string),
			UserDomainID:                d.Get("user_domain_id").(string),
			UserDomainName:              d.Get("user_domain_name").(string),
			Username:                    d.Get("user_name").(string),
			UserID:                      d.Get("user_id").(string),
			ApplicationCredentialID:     d.Get("application_credential_id").(string),
			ApplicationCredentialName:   d.Get("application_credential_name").(string),
			ApplicationCredentialSecret: d.Get("application_credential_secret").(string),
			UseOctavia:                  d.Get("use_octavia").(bool),
			DelayedAuth:                 d.Get("delayed_auth").(bool),
			AllowReauth:                 d.Get("allow_reauth").(bool),
			AuthOpts:                    authOpts,
			MaxRetries:                  d.Get("max_retries").(int),
			DisableNoCacheHeader:        d.Get("disable_no_cache_header").(bool),
			TerraformVersion:            terraformVersion,
			SDKVersion:                  meta.SDKVersionString(),
			MutexKV:                     mutexkv.NewMutexKV(),
			EnableLogger:                enableLogging,
		},
	}

	v, ok := d.GetOkExists("insecure")
	if ok {
		insecure := v.(bool)
		config.Insecure = &insecure
	}

	if err := config.LoadAndValidate(); err != nil {
		return nil, diag.FromErr(err)
	}

	return &config, nil
}
