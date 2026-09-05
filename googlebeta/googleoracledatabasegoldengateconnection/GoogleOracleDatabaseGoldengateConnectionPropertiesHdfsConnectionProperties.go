// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleoracledatabasegoldengateconnection


type GoogleOracleDatabaseGoldengateConnectionPropertiesHdfsConnectionProperties struct {
	// The content of the Hadoop Distributed File System configuration file (core-site.xml).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_oracle_database_goldengate_connection#core_site_xml GoogleOracleDatabaseGoldengateConnection#core_site_xml}
	CoreSiteXml *string `field:"optional" json:"coreSiteXml" yaml:"coreSiteXml"`
	// The technology type of HdfsConnection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_oracle_database_goldengate_connection#technology_type GoogleOracleDatabaseGoldengateConnection#technology_type}
	TechnologyType *string `field:"optional" json:"technologyType" yaml:"technologyType"`
}

