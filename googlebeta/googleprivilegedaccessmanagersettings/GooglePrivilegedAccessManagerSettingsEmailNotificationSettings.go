// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleprivilegedaccessmanagersettings


type GooglePrivilegedAccessManagerSettingsEmailNotificationSettings struct {
	// custom_notification_behavior block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_privileged_access_manager_settings#custom_notification_behavior GooglePrivilegedAccessManagerSettings#custom_notification_behavior}
	CustomNotificationBehavior *GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehavior `field:"optional" json:"customNotificationBehavior" yaml:"customNotificationBehavior"`
	// disable_all_notifications block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_privileged_access_manager_settings#disable_all_notifications GooglePrivilegedAccessManagerSettings#disable_all_notifications}
	DisableAllNotifications *GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsDisableAllNotifications `field:"optional" json:"disableAllNotifications" yaml:"disableAllNotifications"`
}

