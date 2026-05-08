// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googletranscoderjob


type GoogleTranscoderJobConfigElementaryStreamsVideoStream struct {
	// h264 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_transcoder_job#h264 GoogleTranscoderJob#h264}
	H264 *GoogleTranscoderJobConfigElementaryStreamsVideoStreamH264 `field:"optional" json:"h264" yaml:"h264"`
}

