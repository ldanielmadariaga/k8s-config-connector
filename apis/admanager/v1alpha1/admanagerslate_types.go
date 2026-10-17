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

// API sources for AdmanagerSlate, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/ads/admanager/v1/slate_messages.proto
// +kcc:source:service-docs=https://developers.google.com/ad-manager/api/beta
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/ad-manager/docs/reference/rest/v1/networks.slates

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AdmanagerSlateGVK = GroupVersion.WithKind("AdmanagerSlate")

// AdmanagerSlateSpec defines the desired state of AdmanagerSlate
// +kcc:spec:proto=google.ads.admanager.v1.Slate
// +kcc:required-from-proto
type AdmanagerSlateSpec struct {

	// The AdmanagerSlate name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. The display name of the Slate. It has a maximum length of 255
	//  characters.
	// +kcc:proto:field=google.ads.admanager.v1.Slate.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. The location of the original asset if publisher provided and
	//  slate is externally hosted.
	// +kcc:proto:field=google.ads.admanager.v1.Slate.video_source_url
	VideoSourceURL *string `json:"videoSourceURL,omitempty"`
}

// AdmanagerSlateStatus defines the config connector machine state of AdmanagerSlate
type AdmanagerSlateStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AdmanagerSlate resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *AdmanagerSlateObservedState `json:"observedState,omitempty"`
}

// AdmanagerSlateObservedState is the state of the AdmanagerSlate resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.ads.admanager.v1.Slate
type AdmanagerSlateObservedState struct {
	// Output only. The status of this Slate. Slates are created in the
	//  [SlateStatusEnum.SlateStatus.ACTIVE][google.ads.admanager.v1.SlateStatusEnum.SlateStatus.ACTIVE]
	//  state.
	// +kcc:proto:field=google.ads.admanager.v1.Slate.status
	Status *string `json:"status,omitempty"`

	// Output only. Server side transcoding status of the current slate.
	// +kcc:proto:field=google.ads.admanager.v1.Slate.transcode_status
	TranscodeStatus *string `json:"transcodeStatus,omitempty"`

	// Output only. The date and time this slate was last modified.
	// +kcc:proto:field=google.ads.admanager.v1.Slate.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpadmanagerslate;gcpadmanagerslates
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AdmanagerSlate is the Schema for the AdmanagerSlate API
// +k8s:openapi-gen=true
type AdmanagerSlate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AdmanagerSlateSpec   `json:"spec,omitempty"`
	Status AdmanagerSlateStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AdmanagerSlateList contains a list of AdmanagerSlate
type AdmanagerSlateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AdmanagerSlate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AdmanagerSlate{}, &AdmanagerSlateList{})
}
