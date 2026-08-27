// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecertificatemanagercertificate


type GoogleCertificateManagerCertificateSelfManaged struct {
	// The certificate chain in PEM-encoded form.
	//
	// Leaf certificate comes first, followed by intermediate ones if any.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_certificate_manager_certificate#certificate_pem GoogleCertificateManagerCertificate#certificate_pem}
	CertificatePem *string `field:"optional" json:"certificatePem" yaml:"certificatePem"`
	// The certificate chain in PEM-encoded form.
	//
	// Leaf certificate comes first, followed by intermediate ones if any.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_certificate_manager_certificate#pem_certificate GoogleCertificateManagerCertificate#pem_certificate}
	PemCertificate *string `field:"optional" json:"pemCertificate" yaml:"pemCertificate"`
	// The private key of the leaf certificate in PEM-encoded form.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_certificate_manager_certificate#pem_private_key GoogleCertificateManagerCertificate#pem_private_key}
	PemPrivateKey *string `field:"optional" json:"pemPrivateKey" yaml:"pemPrivateKey"`
	// The private key of the leaf certificate in PEM-encoded form.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_certificate_manager_certificate#pem_private_key_wo GoogleCertificateManagerCertificate#pem_private_key_wo}
	PemPrivateKeyWo *string `field:"optional" json:"pemPrivateKeyWo" yaml:"pemPrivateKeyWo"`
	// Triggers update of 'pem_private_key_wo' write-only.
	//
	// Increment this value when an update to 'pem_private_key_wo' is needed. For more info see [updating write-only arguments](/docs/providers/google/guides/using_write_only_arguments.html#updating-write-only-arguments)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_certificate_manager_certificate#pem_private_key_wo_version GoogleCertificateManagerCertificate#pem_private_key_wo_version}
	PemPrivateKeyWoVersion *string `field:"optional" json:"pemPrivateKeyWoVersion" yaml:"pemPrivateKeyWoVersion"`
	// The private key of the leaf certificate in PEM-encoded form.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_certificate_manager_certificate#private_key_pem GoogleCertificateManagerCertificate#private_key_pem}
	PrivateKeyPem *string `field:"optional" json:"privateKeyPem" yaml:"privateKeyPem"`
}

