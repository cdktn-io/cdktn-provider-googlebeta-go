// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecestoolset


type GoogleCesToolsetConnectorToolsetConnectorActionsEntityOperation struct {
	// ID of the entity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_ces_toolset#entity_id GoogleCesToolset#entity_id}
	EntityId *string `field:"required" json:"entityId" yaml:"entityId"`
	// Operation to perform on the entity. Possible values: LIST GET CREATE UPDATE DELETE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_ces_toolset#operation GoogleCesToolset#operation}
	Operation *string `field:"required" json:"operation" yaml:"operation"`
}

