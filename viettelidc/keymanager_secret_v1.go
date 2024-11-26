package viettelidc

import (
	"encoding/base64"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/keymanager/v1/secrets"
)

func keyManagerSecretV1WaitForSecretDeletion(kmClient *gophercloud.ServiceClient, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		err := secrets.Delete(kmClient, id).Err
		if err == nil {
			return "", "DELETED", nil
		}

		if _, ok := err.(gophercloud.ErrDefault404); ok {
			return "", "DELETED", nil
		}

		return nil, "ACTIVE", err
	}
}

func keyManagerSecretV1SecretType(v string) secrets.SecretType {
	var stype secrets.SecretType
	switch v {
	case "symmetric":
		stype = secrets.SymmetricSecret
	case "public":
		stype = secrets.PublicSecret
	case "private":
		stype = secrets.PrivateSecret
	case "passphrase":
		stype = secrets.PassphraseSecret
	case "certificate":
		stype = secrets.CertificateSecret
	case "opaque":
		stype = secrets.OpaqueSecret
	}

	return stype
}

func keyManagerSecretV1WaitForSecretCreation(kmClient *gophercloud.ServiceClient, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		secret, err := secrets.Get(kmClient, id).Extract()
		if err != nil {
			if _, ok := err.(gophercloud.ErrDefault404); ok {
				return "", "NOT_CREATED", nil
			}

			return "", "NOT_CREATED", err
		}

		if secret.Status == "ERROR" {
			return "", secret.Status, fmt.Errorf("Error creating secret")
		}

		return secret, secret.Status, nil
	}
}

func keyManagerSecretV1GetUUIDfromSecretRef(ref string) string {
	// secret ref has form https://{barbican_host}/v1/secrets/{secret_uuid}
	// so we are only interested in the last part
	refSplit := strings.Split(ref, "/")
	uuid := refSplit[len(refSplit)-1]
	return uuid
}

func flattenKeyManagerSecretV1Metadata(d *schema.ResourceData) map[string]string {
	m := make(map[string]string)
	for key, val := range d.Get("metadata").(map[string]interface{}) {
		m[key] = val.(string)
	}
	return m
}

func keyManagerSecretMetadataV1WaitForSecretMetadataCreation(kmClient *gophercloud.ServiceClient, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		metadata, err := secrets.GetMetadata(kmClient, id).Extract()
		if err != nil {
			if _, ok := err.(gophercloud.ErrDefault404); ok {
				return "", "NOT_CREATED", nil
			}

			return "", "NOT_CREATED", err
		}
		return metadata, "ACTIVE", nil
	}
}

func keyManagerSecretV1GetPayload(kmClient *gophercloud.ServiceClient, id, contentType string) string {
	opts := secrets.GetPayloadOpts{
		PayloadContentType: contentType,
	}
	payload, err := secrets.GetPayload(kmClient, id, opts).Extract()
	if err != nil {
		log.Printf("[DEBUG] Could not retrieve payload for secret with id %s: %s", id, err)
	}

	if !strings.HasPrefix(contentType, "text/") {
		return base64.StdEncoding.EncodeToString(payload)
	}

	return string(payload)
}
