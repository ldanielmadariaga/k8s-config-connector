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

// API sources for AdmanagerTargetingPreset, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/ads/admanager/v1/targeting_preset_messages.proto
// +kcc:source:service-docs=https://developers.google.com/ad-manager/api/beta
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/ad-manager/docs/reference/rest/v1/networks.targetingPresets

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AdmanagerTargetingPresetGVK = GroupVersion.WithKind("AdmanagerTargetingPreset")

// AdmanagerTargetingPresetSpec defines the desired state of AdmanagerTargetingPreset
// +kcc:spec:proto=google.ads.admanager.v1.TargetingPreset
// +kcc:required-from-proto
type AdmanagerTargetingPresetSpec struct {

	// The AdmanagerTargetingPreset name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. The name of the TargetingPreset. This attribute has a maximum
	//  length of 255 characters.
	// +kcc:proto:field=google.ads.admanager.v1.TargetingPreset.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Required. Contains the targeting criteria for the TargetingPreset. If not
	//  provided, empty targeting will be used as a default.
	// +kcc:proto:field=google.ads.admanager.v1.TargetingPreset.targeting
	// +required
	Targeting *Targeting `json:"targeting,omitempty"`
}

// AdmanagerTargetingPresetStatus defines the config connector machine state of AdmanagerTargetingPreset
type AdmanagerTargetingPresetStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AdmanagerTargetingPreset resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *AdmanagerTargetingPresetObservedState `json:"observedState,omitempty"`
}

// AdmanagerTargetingPresetObservedState is the state of the AdmanagerTargetingPreset resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.ads.admanager.v1.TargetingPreset
type AdmanagerTargetingPresetObservedState struct {
	// Required. Contains the targeting criteria for the TargetingPreset. If not
	//  provided, empty targeting will be used as a default.
	// +kcc:proto:field=google.ads.admanager.v1.TargetingPreset.targeting
	Targeting *TargetingObservedState `json:"targeting,omitempty"`

	// Output only. The status of the TargetingPreset.
	// +kcc:proto:field=google.ads.admanager.v1.TargetingPreset.status
	Status *string `json:"status,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpadmanagertargetingpreset;gcpadmanagertargetingpresets
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AdmanagerTargetingPreset is the Schema for the AdmanagerTargetingPreset API
// +k8s:openapi-gen=true
type AdmanagerTargetingPreset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AdmanagerTargetingPresetSpec   `json:"spec,omitempty"`
	Status AdmanagerTargetingPresetStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AdmanagerTargetingPresetList contains a list of AdmanagerTargetingPreset
type AdmanagerTargetingPresetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AdmanagerTargetingPreset `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AdmanagerTargetingPreset{}, &AdmanagerTargetingPresetList{})
}
