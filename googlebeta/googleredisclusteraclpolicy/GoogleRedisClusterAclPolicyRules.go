// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleredisclusteraclpolicy


type GoogleRedisClusterAclPolicyRules struct {
	// The rule to be applied to the username.
	//
	// Ex: "on >password123 ~* +@all"
	// The format of the rule is defined by Redis OSS:
	// https://redis.io/docs/latest/operate/oss_and_stack/management/security/acl/
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_redis_cluster_acl_policy#rule GoogleRedisClusterAclPolicy#rule}
	Rule *string `field:"required" json:"rule" yaml:"rule"`
	// Specifies the IAM user or service account to be added to the ACL policy.
	//
	// This username will be directly set on the Redis OSS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_redis_cluster_acl_policy#username GoogleRedisClusterAclPolicy#username}
	Username *string `field:"required" json:"username" yaml:"username"`
}

