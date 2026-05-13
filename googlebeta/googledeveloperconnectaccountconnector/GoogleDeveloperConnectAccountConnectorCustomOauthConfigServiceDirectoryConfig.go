// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledeveloperconnectaccountconnector


type GoogleDeveloperConnectAccountConnectorCustomOauthConfigServiceDirectoryConfig struct {
	// The Service Directory service name. Format: projects/{project}/locations/{location}/namespaces/{namespace}/services/{service}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_developer_connect_account_connector#service GoogleDeveloperConnectAccountConnector#service}
	Service *string `field:"required" json:"service" yaml:"service"`
}

