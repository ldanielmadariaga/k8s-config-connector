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

// API sources for AdmanagerAdRule, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/ads/admanager/v1/ad_rule_messages.proto
// +kcc:source:service-docs=https://developers.google.com/ad-manager/api/beta
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/ad-manager/docs/reference/rest/v1/networks.adRules

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AdmanagerAdRuleGVK = GroupVersion.WithKind("AdmanagerAdRule")

// AdmanagerAdRuleSpec defines the desired state of AdmanagerAdRule
// +kcc:spec:proto=google.ads.admanager.v1.AdRule
// +kcc:required-from-proto
type AdmanagerAdRuleSpec struct {

	// The AdmanagerAdRule name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. The unique name of the AdRule. This attribute is required to
	//  create an ad rule and has a maximum length of 255 characters.
	// +kcc:proto:field=google.ads.admanager.v1.AdRule.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. The priority of the AdRule. This attribute is required and can
	//  range from 1 to 1000, with 1 being the highest possible priority.
	//
	//  Changing an ad rule's priority can affect the priorities of other ad rules.
	//  For example,
	//  increasing an ad rule's priority from 5 to 1 will shift the ad rules that
	//  were previously in
	//  priority positions 1 through 4 down one.
	// +kcc:proto:field=google.ads.admanager.v1.AdRule.priority
	Priority *int32 `json:"priority,omitempty"`

	// Required. The start time of the AdRule. This attribute is required and must
	//  be a date in the future for new ad rules.
	// +kcc:proto:field=google.ads.admanager.v1.AdRule.start_time
	// +required
	StartTime *string `json:"startTime,omitempty"`

	// Optional. This end time of the AdRule. This attribute is required
	//  unless
	//  [endTimeUnlimited][google.ads.admanager.v1.AdRule.end_time_unlimited] is
	//  set to true. If specified, it must be after the
	//  [startTime][google.ads.admanager.v1.AdRule.start_time].
	// +kcc:proto:field=google.ads.admanager.v1.AdRule.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Optional. Specifies whether or not the AdRule has an end time. This
	//  attribute is optional and defaults to false.
	// +kcc:proto:field=google.ads.admanager.v1.AdRule.end_time_unlimited
	EndTimeUnlimited *bool `json:"endTimeUnlimited,omitempty"`

	// Optional. The FrequencyCapBehavior of the AdRule. This attribute is
	//  optional and defaults to
	//  [AdRuleFrequencyCapBehaviorEnum.AdRuleFrequencyCapBehavior.DEFER][google.ads.admanager.v1.AdRuleFrequencyCapBehaviorEnum.AdRuleFrequencyCapBehavior.DEFER].
	// +kcc:proto:field=google.ads.admanager.v1.AdRule.frequency_cap_behavior
	FrequencyCapBehavior *string `json:"frequencyCapBehavior,omitempty"`

	// Optional. This AdRule object's frequency cap for the maximum impressions
	//  per stream. This attribute is optional and defaults to 0.
	// +kcc:proto:field=google.ads.admanager.v1.AdRule.max_impressions_per_line_item_per_stream
	MaxImpressionsPerLineItemPerStream *int32 `json:"maxImpressionsPerLineItemPerStream,omitempty"`

	// Optional. This AdRule object's frequency cap for the maximum impressions
	//  per pod. This attribute is optional and defaults to 0.
	// +kcc:proto:field=google.ads.admanager.v1.AdRule.max_impressions_per_line_item_per_pod
	MaxImpressionsPerLineItemPerPod *int32 `json:"maxImpressionsPerLineItemPerPod,omitempty"`

	// Required. This AdRule object's pre-roll slot. This attribute is required.
	// +kcc:proto:field=google.ads.admanager.v1.AdRule.preroll
	// +required
	Preroll *AdRuleSlot `json:"preroll,omitempty"`

	// Required. This AdRule object's video mid-roll slots. This attribute is
	//  required, and there must be at least one mid-roll.
	// +kcc:proto:field=google.ads.admanager.v1.AdRule.midrolls
	// +required
	Midrolls []AdRuleSlot `json:"midrolls,omitempty"`

	// Required. This AdRule object's post-roll slot. This attribute is required.
	// +kcc:proto:field=google.ads.admanager.v1.AdRule.postroll
	// +required
	Postroll *AdRuleSlot `json:"postroll,omitempty"`

	// Required. The targeting of the `AdRule`.
	// +kcc:proto:field=google.ads.admanager.v1.AdRule.targeting
	// +required
	Targeting *Targeting `json:"targeting,omitempty"`
}

// AdmanagerAdRuleStatus defines the config connector machine state of AdmanagerAdRule
type AdmanagerAdRuleStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AdmanagerAdRule resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *AdmanagerAdRuleObservedState `json:"observedState,omitempty"`
}

// AdmanagerAdRuleObservedState is the state of the AdmanagerAdRule resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.ads.admanager.v1.AdRule
type AdmanagerAdRuleObservedState struct {
	// Output only. The AdRuleStatus of the AdRule. This attribute is read-only
	//  and defaults to
	//  [AdRuleStatusEnum.AdRuleStatus.INACTIVE][google.ads.admanager.v1.AdRuleStatusEnum.AdRuleStatus.INACTIVE].
	// +kcc:proto:field=google.ads.admanager.v1.AdRule.status
	Status *string `json:"status,omitempty"`

	// Required. The targeting of the `AdRule`.
	// +kcc:proto:field=google.ads.admanager.v1.AdRule.targeting
	Targeting *TargetingObservedState `json:"targeting,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpadmanageradrule;gcpadmanageradrules
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AdmanagerAdRule is the Schema for the AdmanagerAdRule API
// +k8s:openapi-gen=true
type AdmanagerAdRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AdmanagerAdRuleSpec   `json:"spec,omitempty"`
	Status AdmanagerAdRuleStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AdmanagerAdRuleList contains a list of AdmanagerAdRule
type AdmanagerAdRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AdmanagerAdRule `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AdmanagerAdRule{}, &AdmanagerAdRuleList{})
}
