// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecomputebackendservice


type GoogleComputeBackendServiceTlsSettings struct {
	// Reference to the BackendAuthenticationConfig resource from the networksecurity.googleapis.com namespace. Can be used in authenticating TLS connections to the backend, as specified by the authenticationMode field. Can only be specified if authenticationMode is not NONE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_compute_backend_service#authentication_config GoogleComputeBackendService#authentication_config}
	AuthenticationConfig *string `field:"optional" json:"authenticationConfig" yaml:"authenticationConfig"`
	// The fully-specified SPIFFE ID without the spiffe:// scheme.
	//
	// Must be in the format //<trust_domain>/ns/<namespace>/sa/<subject>.
	// The load balancer uses certificates and roots of trust provisioned by the Managed Workload Identity system for this identity.
	// The Trust Domain within the identity must refer to a valid Workload Identity Pool, from which the TrustConfig and CertificateIssuanceConfig are inherited.
	// If set, you cannot configure sni, subjectAltNames, or authenticationConfig manually.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_compute_backend_service#identity GoogleComputeBackendService#identity}
	Identity *string `field:"optional" json:"identity" yaml:"identity"`
	// Server Name Indication - see RFC3546 section 3.1. If set, the load balancer sends this string as the SNI hostname in the TLS connection to the backend, and requires that this string match a Subject Alternative Name (SAN) in the backend's server certificate. With a Regional Internet NEG backend, if the SNI is specified here, the load balancer uses it regardless of whether the Regional Internet NEG is specified with FQDN or IP address and port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_compute_backend_service#sni GoogleComputeBackendService#sni}
	Sni *string `field:"optional" json:"sni" yaml:"sni"`
	// subject_alt_names block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_compute_backend_service#subject_alt_names GoogleComputeBackendService#subject_alt_names}
	SubjectAltNames interface{} `field:"optional" json:"subjectAltNames" yaml:"subjectAltNames"`
}

