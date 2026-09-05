// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleiamworkforcepoolprovider


type GoogleIamWorkforcePoolProviderOidcClientSecretValue struct {
	// The plain text of the client secret value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_iam_workforce_pool_provider#plain_text GoogleIamWorkforcePoolProvider#plain_text}
	PlainText *string `field:"optional" json:"plainText" yaml:"plainText"`
	// The plain text of the client secret value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_iam_workforce_pool_provider#plain_text_wo GoogleIamWorkforcePoolProvider#plain_text_wo}
	PlainTextWo *string `field:"optional" json:"plainTextWo" yaml:"plainTextWo"`
	// Triggers update of 'plain_text_wo' write-only.
	//
	// Increment this value when an update to 'plain_text_wo' is needed. For more info see [updating write-only arguments](/docs/providers/google/guides/using_write_only_arguments.html#updating-write-only-arguments)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_iam_workforce_pool_provider#plain_text_wo_version GoogleIamWorkforcePoolProvider#plain_text_wo_version}
	PlainTextWoVersion *string `field:"optional" json:"plainTextWoVersion" yaml:"plainTextWoVersion"`
}

