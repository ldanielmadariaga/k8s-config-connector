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

// API sources for AdmanagerAdUnit, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/ads/admanager/v1/ad_unit_messages.proto
// +kcc:source:service-docs=https://developers.google.com/ad-manager/api/beta
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/ad-manager/docs/reference/rest/v1/networks.adUnits

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AdmanagerAdUnitGVK = GroupVersion.WithKind("AdmanagerAdUnit")

// AdmanagerAdUnitSpec defines the desired state of AdmanagerAdUnit
// +kcc:spec:proto=google.ads.admanager.v1.AdUnit
// +kcc:required-from-proto
type AdmanagerAdUnitSpec struct {

	// The AdmanagerAdUnit name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. Immutable. The AdUnit's parent. Every ad unit has a parent except
	//  for the root ad unit, which is created by Google. Format:
	//  "networks/{network_code}/adUnits/{ad_unit_id}"
	// +kcc:proto:field=google.ads.admanager.v1.AdUnit.parent_ad_unit
	// +required
	ParentAdUnit *string `json:"parentAdUnit,omitempty"`

	// Required. The display name of the ad unit. Its maximum length is 255
	//  characters.
	// +kcc:proto:field=google.ads.admanager.v1.AdUnit.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Immutable. A string used to uniquely identify the ad unit for the
	//  purposes of serving the ad. This attribute can be set during ad unit
	//  creation. If it is not provided, it will be assigned by Google based on the
	//  ad unit ID.
	// +kcc:proto:field=google.ads.admanager.v1.AdUnit.ad_unit_code
	AdUnitCode *string `json:"adUnitCode,omitempty"`

	// Optional. The target window directly applied to this AdUnit.
	//  If this field is not set, this AdUnit uses the target window specified in
	//  effectiveTargetWindow.
	// +kcc:proto:field=google.ads.admanager.v1.AdUnit.applied_target_window
	AppliedTargetWindow *string `json:"appliedTargetWindow,omitempty"`

	// Optional. The resource names of Teams directly applied to this AdUnit.
	//  Format: "networks/{network_code}/teams/{team_id}"
	// +kcc:proto:field=google.ads.admanager.v1.AdUnit.applied_teams
	AppliedTeams []string `json:"appliedTeams,omitempty"`

	// Optional. A description of the ad unit. The maximum length is 65,535
	//  characters.
	// +kcc:proto:field=google.ads.admanager.v1.AdUnit.description
	Description *string `json:"description,omitempty"`

	// Optional. If this field is set to true, then the AdUnit will not be
	//  implicitly targeted when its parent is. Traffickers must explicitly
	//  target such an AdUnit or else no line items will serve to it. This
	//  feature is only available for Ad Manager 360 accounts.
	// +kcc:proto:field=google.ads.admanager.v1.AdUnit.explicitly_targeted
	ExplicitlyTargeted *bool `json:"explicitlyTargeted,omitempty"`

	// Optional. The sizes that can be served inside this ad unit.
	// +kcc:proto:field=google.ads.admanager.v1.AdUnit.ad_unit_sizes
	AdUnitSizes []AdUnitSize `json:"adUnitSizes,omitempty"`

	// Optional. The duration after which an Ad Unit will automatically refresh.
	//  This is only valid for ad units in mobile apps. If not set, the ad unit
	//  will not refresh.
	// +kcc:proto:field=google.ads.admanager.v1.AdUnit.refresh_delay
	RefreshDelay *string `json:"refreshDelay,omitempty"`

	// Optional. The set of labels applied directly to this ad unit.
	// +kcc:proto:field=google.ads.admanager.v1.AdUnit.applied_labels
	AppliedLabels []AppliedLabel `json:"appliedLabels,omitempty"`

	// Optional. The set of label frequency caps applied directly to this ad unit.
	//  There is a limit of 10 label frequency caps per ad unit.
	// +kcc:proto:field=google.ads.admanager.v1.AdUnit.applied_label_frequency_caps
	AppliedLabelFrequencyCaps []LabelFrequencyCap `json:"appliedLabelFrequencyCaps,omitempty"`

	// Optional. Non-empty default. The smart size mode for this ad unit. This
	//  attribute defaults to SmartSizeMode.NONE for fixed sizes.
	// +kcc:proto:field=google.ads.admanager.v1.AdUnit.smart_size_mode
	SmartSizeMode *string `json:"smartSizeMode,omitempty"`

	// Optional. The value of AdSense enabled directly applied to this ad unit. If
	//  not specified this ad unit will inherit the value of
	//  effectiveAdsenseEnabled from its ancestors.
	// +kcc:proto:field=google.ads.admanager.v1.AdUnit.applied_adsense_enabled
	AppliedAdsenseEnabled *bool `json:"appliedAdsenseEnabled,omitempty"`

	// Optional. Non-empty default. Defines the type of refresh rate control for
	//  this ad unit. This field defaults to `DISABLED`.
	// +kcc:proto:field=google.ads.admanager.v1.AdUnit.refresh_rate_type
	RefreshRateType *string `json:"refreshRateType,omitempty"`
}

// AdmanagerAdUnitStatus defines the config connector machine state of AdmanagerAdUnit
type AdmanagerAdUnitStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AdmanagerAdUnit resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *AdmanagerAdUnitObservedState `json:"observedState,omitempty"`
}

// AdmanagerAdUnitObservedState is the state of the AdmanagerAdUnit resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.ads.admanager.v1.AdUnit
type AdmanagerAdUnitObservedState struct {
	// Output only. Deprecated: AdUnit ID.
	// +kcc:proto:field=google.ads.admanager.v1.AdUnit.ad_unit_id
	AdUnitID *int64 `json:"adUnitID,omitempty"`

	// Output only. The path to this AdUnit in the ad unit hierarchy represented
	//  as a list from the root to this ad unit's parent. For root ad units, this
	//  list is empty.
	// +kcc:proto:field=google.ads.admanager.v1.AdUnit.parent_path
	ParentPath []AdUnitParentObservedState `json:"parentPath,omitempty"`

	// Output only. The status of this ad unit.  It defaults to ACTIVE.
	// +kcc:proto:field=google.ads.admanager.v1.AdUnit.status
	Status *string `json:"status,omitempty"`

	// Output only. Non-empty default. The target window of this AdUnit. This
	//  value is inherited from ancestor AdUnits and defaults to TOP if no AdUnit
	//  in the hierarchy specifies it.
	// +kcc:proto:field=google.ads.admanager.v1.AdUnit.effective_target_window
	EffectiveTargetWindow *string `json:"effectiveTargetWindow,omitempty"`

	// Output only. The resource names of all Teams that this AdUnit is on as well
	//  as those inherited from parent AdUnits. Format:
	//  "networks/{network_code}/teams/{team_id}"
	// +kcc:proto:field=google.ads.admanager.v1.AdUnit.teams
	Teams []string `json:"teams,omitempty"`

	// Output only. This field is set to true if the ad unit has any children.
	// +kcc:proto:field=google.ads.admanager.v1.AdUnit.has_children
	HasChildren *bool `json:"hasChildren,omitempty"`

	// Output only. The time this AdUnit was last modified.
	// +kcc:proto:field=google.ads.admanager.v1.AdUnit.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Contains the set of labels applied directly to the ad unit as
	//  well as those inherited from the parent ad units. If a label has been
	//  negated, only the negated label is returned. This attribute is assigned by
	//  Google.
	// +kcc:proto:field=google.ads.admanager.v1.AdUnit.effective_applied_labels
	EffectiveAppliedLabels []AppliedLabelObservedState `json:"effectiveAppliedLabels,omitempty"`

	// Output only. The label frequency caps applied directly to the ad unit as
	//  well as those inherited from parent ad units.
	// +kcc:proto:field=google.ads.admanager.v1.AdUnit.effective_label_frequency_caps
	EffectiveLabelFrequencyCaps []LabelFrequencyCapObservedState `json:"effectiveLabelFrequencyCaps,omitempty"`

	// Output only. Specifies whether or not the AdUnit is enabled for serving ads
	//  from the AdSense content network. This attribute defaults to the ad unit's
	//  parent or ancestor's setting if one has been set. If no ancestor of the ad
	//  unit has set appliedAdsenseEnabled, the attribute is defaulted to true.
	// +kcc:proto:field=google.ads.admanager.v1.AdUnit.effective_adsense_enabled
	EffectiveAdsenseEnabled *bool `json:"effectiveAdsenseEnabled,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpadmanageradunit;gcpadmanageradunits
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AdmanagerAdUnit is the Schema for the AdmanagerAdUnit API
// +k8s:openapi-gen=true
type AdmanagerAdUnit struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AdmanagerAdUnitSpec   `json:"spec,omitempty"`
	Status AdmanagerAdUnitStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AdmanagerAdUnitList contains a list of AdmanagerAdUnit
type AdmanagerAdUnitList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AdmanagerAdUnit `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AdmanagerAdUnit{}, &AdmanagerAdUnitList{})
}
