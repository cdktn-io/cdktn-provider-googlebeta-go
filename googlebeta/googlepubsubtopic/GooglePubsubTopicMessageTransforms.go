// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlepubsubtopic


type GooglePubsubTopicMessageTransforms struct {
	// ai_inference block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_pubsub_topic#ai_inference GooglePubsubTopic#ai_inference}
	AiInference *GooglePubsubTopicMessageTransformsAiInference `field:"optional" json:"aiInference" yaml:"aiInference"`
	// Controls whether or not to use this transform.
	//
	// If not set or 'false',
	// the transform will be applied to messages. Default: 'true'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_pubsub_topic#disabled GooglePubsubTopic#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// javascript_udf block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_pubsub_topic#javascript_udf GooglePubsubTopic#javascript_udf}
	JavascriptUdf *GooglePubsubTopicMessageTransformsJavascriptUdf `field:"optional" json:"javascriptUdf" yaml:"javascriptUdf"`
}

