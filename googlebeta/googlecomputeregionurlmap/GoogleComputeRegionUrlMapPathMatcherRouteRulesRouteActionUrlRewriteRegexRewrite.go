// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecomputeregionurlmap


type GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionUrlRewriteRegexRewrite struct {
	// The regular expression used to match against the URL path. It uses RE2 syntax with the following constraints:.
	//
	// * Any single character operators are allowed.
	// * Groups may only contain a submatch operator, and may not
	//   contain character repetition (for example, '.*').
	// * Character repetition (for example, '.*') may only be used in
	//   a regex together with empty string operators, other
	//   repetitions, ranges, and repetitions of ranges.
	// * Ranges may only contain character ranges, digit ranges, and
	//   symbols allowed for ranges.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_compute_region_url_map#path_pattern GoogleComputeRegionUrlMap#path_pattern}
	PathPattern *string `field:"required" json:"pathPattern" yaml:"pathPattern"`
	// The substitution used to rewrite the parts of the URL path matched by pathPattern. May reference capture groups from pathPattern.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_compute_region_url_map#path_substitution GoogleComputeRegionUrlMap#path_substitution}
	PathSubstitution *string `field:"required" json:"pathSubstitution" yaml:"pathSubstitution"`
}

