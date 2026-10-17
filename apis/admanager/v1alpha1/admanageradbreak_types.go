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

// API sources for AdmanagerAdBreak, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/ads/admanager/v1/ad_break_messages.proto
// +kcc:source:service-docs=https://developers.google.com/ad-manager/api/beta
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/ad-manager/docs/reference/rest/v1/networks.liveStreamEventsByAssetKey.adBreaks

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AdmanagerAdBreakGVK = GroupVersion.WithKind("AdmanagerAdBreak")

// AdmanagerAdBreakSpec defines the desired state of AdmanagerAdBreak
// +kcc:spec:proto=google.ads.admanager.v1.AdBreak
// +kcc:required-from-proto
type AdmanagerAdBreakSpec struct {

	// The AdmanagerAdBreak name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. Immutable. `AdBreak` ID.
	//  Must consist only of lowercase letters, digits, and hyphens. Ad break IDs
	//  have a maximum length of 63 characters.
	//  If not set, an ad break ID is generated as a UUID string.
	// +kcc:proto:field=google.ads.admanager.v1.AdBreak.ad_break_id
	AdBreakID *string `json:"adBreakID,omitempty"`

	// Optional. Immutable. The asset key of the `LiveStreamEvent` that the ad
	//  break belongs to. Either an asset key or a custom asset key must be
	//  provided for creation.
	// +kcc:proto:field=google.ads.admanager.v1.AdBreak.asset_key
	AssetKey *string `json:"assetKey,omitempty"`

	// Optional. Immutable. The custom asset key of the `LiveStreamEvent` that the
	//  ad break belongs to. Either an asset key or a custom asset key must be
	//  provided for creation.
	// +kcc:proto:field=google.ads.admanager.v1.AdBreak.custom_asset_key
	CustomAssetKey *string `json:"customAssetKey,omitempty"`

	// Optional. The expected start time of the ad break.
	//  When set, the scheduled ad break will be decisioned shortly before the
	//  expected start time.
	//  If this field is not set, ad decisioning begins immediately.
	//  An ad break can be created up to six hours before the expected start time.
	// +kcc:proto:field=google.ads.admanager.v1.AdBreak.expected_start_time
	ExpectedStartTime *string `json:"expectedStartTime,omitempty"`

	// Required. The duration of the `AdBreak`.
	//  An expected duration in seconds with up to nine fractional digits, ending
	//  with `s`, for example, `3.5s`. This value will be updated to match the
	//  actual duration from the manifest or pod request after the ad break's state
	//  is
	//  [`COMPLETE`][google.ads.admanager.v1.AdBreakStateEnum.AdBreakState.COMPLETE].
	// +kcc:proto:field=google.ads.admanager.v1.AdBreak.duration
	// +required
	Duration *string `json:"duration,omitempty"`

	// Optional. The pod template name of the `AdBreak`.
	//  This field is the required unique name across all pod templates in the
	//  network, not the display name.
	// +kcc:proto:field=google.ads.admanager.v1.AdBreak.pod_template_name
	PodTemplateName *string `json:"podTemplateName,omitempty"`

	// Optional. The key-value pairs to be included on the ad requests for this
	//  `AdBreak`. Key-value pairs to include on ad requests for this break for
	//  custom criteria targeting in Google Ad Manager, separated by `=` and joined
	//  by
	//  `&`.
	//
	//  Format:"key1=value&key2=value"
	// +kcc:proto:field=google.ads.admanager.v1.AdBreak.custom_params
	CustomParams *string `json:"customParams,omitempty"`

	// Optional. The Base64-encoded SCTE-35 command associated with the `AdBreak`.
	//  This field can include the `splice_insert()` or `time_signal()` command.
	//
	//  **Examples**
	//
	//  * `time_signal()`
	//    ```
	//    /DA0AAAAAAAA///wBQb+cr0AUAAeAhxDVUVJSAAAjn/PAAGlmbAICAAAAAAsoKGKNAIAmsnRfg==
	//    ```
	//  * `splice_insert()`
	//    ```
	//    /DAvAAAAAAAA///wFAVIAACPf+/+c2nALv4AUsz1AAAAAAAKAAhDVUVJAAABNWLbowo=
	//    ```
	// +kcc:proto:field=google.ads.admanager.v1.AdBreak.scte_35_cue_out
	Scte35CueOut *string `json:"scte35CueOut,omitempty"`
}

// AdmanagerAdBreakStatus defines the config connector machine state of AdmanagerAdBreak
type AdmanagerAdBreakStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AdmanagerAdBreak resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *AdmanagerAdBreakObservedState `json:"observedState,omitempty"`
}

// AdmanagerAdBreakObservedState is the state of the AdmanagerAdBreak resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.ads.admanager.v1.AdBreak
type AdmanagerAdBreakObservedState struct {
	// Output only. The state of the `AdBreak`.
	// +kcc:proto:field=google.ads.admanager.v1.AdBreak.break_state
	BreakState *string `json:"breakState,omitempty"`

	// Output only. The sequence id of the `AdBreak`.
	//  The unique sequence number of the created ad break. This value is only set
	//  after the `AdBreak` starts decisioning indicated by the ad break state's
	//  being
	//  [`DECISIONED`][google.ads.admanager.v1.AdBreakStateEnum.AdBreakState.DECISIONED]
	//  .
	// +kcc:proto:field=google.ads.admanager.v1.AdBreak.break_sequence
	BreakSequence *int64 `json:"breakSequence,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpadmanageradbreak;gcpadmanageradbreaks
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AdmanagerAdBreak is the Schema for the AdmanagerAdBreak API
// +k8s:openapi-gen=true
type AdmanagerAdBreak struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AdmanagerAdBreakSpec   `json:"spec,omitempty"`
	Status AdmanagerAdBreakStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AdmanagerAdBreakList contains a list of AdmanagerAdBreak
type AdmanagerAdBreakList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AdmanagerAdBreak `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AdmanagerAdBreak{}, &AdmanagerAdBreakList{})
}
