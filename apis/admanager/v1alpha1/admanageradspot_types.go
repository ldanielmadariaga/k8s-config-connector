// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// API sources for AdmanagerAdSpot, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/ads/admanager/v1/ad_spot_messages.proto
// +kcc:source:service-docs=https://developers.google.com/ad-manager/api/beta
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/ad-manager/docs/reference/rest/v1/networks.adSpots

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AdmanagerAdSpotGVK = GroupVersion.WithKind("AdmanagerAdSpot")

// AdmanagerAdSpotSpec defines the desired state of AdmanagerAdSpot
// +kcc:spec:proto=google.ads.admanager.v1.AdSpot
// +kcc:required-from-proto
type AdmanagerAdSpotSpec struct {

	// The AdmanagerAdSpot name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. Name of the AdSpot. The name is case insensitive and can be
	//  referenced in ad tags. This value is required if `customSpot` is true, and
	//  cannot be set otherwise. You can use alphanumeric characters and symbols
	//  other than the following: ", ', =, !, +, #, , ~, ;, ^, (, ), <, >, [, ],
	//  the white space character.
	// +kcc:proto:field=google.ads.admanager.v1.AdSpot.canonical_display_name
	CanonicalDisplayName *string `json:"canonicalDisplayName,omitempty"`

	// Optional. Descriptive name for the `AdSpot`.This value is optional if
	//  `customSpot` is true, and cannot be set otherwise.
	// +kcc:proto:field=google.ads.admanager.v1.AdSpot.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Whether this ad spot is a custom spot. This field is optional and
	//  defaults to false. Custom spots can be reused and targeted in the targeting
	//  picker.
	// +kcc:proto:field=google.ads.admanager.v1.AdSpot.custom_spot
	CustomSpot *bool `json:"customSpot,omitempty"`

	// Optional. Whether this ad spot is a flexible spot. This field is optional
	//  and defaults to false. Flexible spots are allowed to have no max number of
	//  ads.
	// +kcc:proto:field=google.ads.admanager.v1.AdSpot.flexible
	Flexible *bool `json:"flexible,omitempty"`

	// Optional. The maximum total duration for this AdSpot. This field is
	//  optional, defaults to 0, and supports precision to the nearest second.
	// +kcc:proto:field=google.ads.admanager.v1.AdSpot.max_duration
	MaxDuration *string `json:"maxDuration,omitempty"`

	// Optional. The minimum allowed duration for ads in the AdSpot. This field is
	//  optional and defaults to 0.
	// +kcc:proto:field=google.ads.admanager.v1.AdSpot.min_ad_duration
	MinAdDuration *string `json:"minAdDuration,omitempty"`

	// Required. The maximum allowed duration for ads in the `AdSpot`. This field
	//  is required and must be greater than
	//  [minAdDuration][google.ads.admanager.v1.AdSpot.min_ad_duration].
	// +kcc:proto:field=google.ads.admanager.v1.AdSpot.max_ad_duration
	// +required
	MaxAdDuration *string `json:"maxAdDuration,omitempty"`

	// Optional. The maximum number of ads allowed in the AdSpot. This field is
	//  optional and defaults to 0. A value of 0 means that there is no maximum for
	//  the number of ads in the ad spot. No max ads is only supported for ad spots
	//  that have flexible set to true.
	// +kcc:proto:field=google.ads.admanager.v1.AdSpot.max_ads
	MaxAds *int32 `json:"maxAds,omitempty"`

	// Optional. The AdSpot TargetingType determines how this ad spot can be
	//  targeted. This field is required.
	// +kcc:proto:field=google.ads.admanager.v1.AdSpot.targeting_type
	TargetingType *string `json:"targetingType,omitempty"`

	// Optional. Whether backfill is blocked in this ad spot. This field is
	//  optional and defaults to false.
	// +kcc:proto:field=google.ads.admanager.v1.AdSpot.backfill_blocked
	BackfillBlocked *bool `json:"backfillBlocked,omitempty"`

	// Optional. The set of line item types that may appear in the ad spot. This
	//  field is optional and defaults to an empty set, which means that all types
	//  are allowed. Note, backfill reservation types are controlled by the
	//  `backfillBlocked` field.
	// +kcc:proto:field=google.ads.admanager.v1.AdSpot.allowed_line_item_types
	AllowedLineItemTypes []string `json:"allowedLineItemTypes,omitempty"`

	// Optional. Whether inventory sharing is blocked in this ad spot. This field
	//  is optional and defaults to false.
	// +kcc:proto:field=google.ads.admanager.v1.AdSpot.inventory_sharing_blocked
	InventorySharingBlocked *bool `json:"inventorySharingBlocked,omitempty"`
}

// AdmanagerAdSpotStatus defines the config connector machine state of AdmanagerAdSpot
type AdmanagerAdSpotStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AdmanagerAdSpot resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	// ObservedState *AdmanagerAdSpotObservedState `json:"observedState,omitempty"`
}

// AdmanagerAdSpotObservedState is the state of the AdmanagerAdSpot resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.ads.admanager.v1.AdSpot
// type AdmanagerAdSpotObservedState struct {
// }

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpadmanageradspot;gcpadmanageradspots
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AdmanagerAdSpot is the Schema for the AdmanagerAdSpot API
// +k8s:openapi-gen=true
type AdmanagerAdSpot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AdmanagerAdSpotSpec   `json:"spec,omitempty"`
	Status AdmanagerAdSpotStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AdmanagerAdSpotList contains a list of AdmanagerAdSpot
type AdmanagerAdSpotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AdmanagerAdSpot `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AdmanagerAdSpot{}, &AdmanagerAdSpotList{})
}
