// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package providerfunctions

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktn/provider-google-beta.providerFunctions.GoogleBetaProviderFunctions",
		reflect.TypeOf((*GoogleBetaProviderFunctions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "locationFromId", GoMethod: "LocationFromId"},
			_jsii_.MemberMethod{JsiiMethod: "nameFromId", GoMethod: "NameFromId"},
			_jsii_.MemberMethod{JsiiMethod: "projectFromId", GoMethod: "ProjectFromId"},
			_jsii_.MemberMethod{JsiiMethod: "regionFromId", GoMethod: "RegionFromId"},
			_jsii_.MemberMethod{JsiiMethod: "regionFromZone", GoMethod: "RegionFromZone"},
			_jsii_.MemberMethod{JsiiMethod: "zoneFromId", GoMethod: "ZoneFromId"},
		},
		func() interface{} {
			return &jsiiProxy_GoogleBetaProviderFunctions{}
		},
	)
}
