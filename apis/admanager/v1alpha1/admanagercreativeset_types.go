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

// API sources for AdmanagerCreativeSet, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/ads/admanager/v1/creative_set_messages.proto
// +kcc:source:service-docs=https://developers.google.com/ad-manager/api/beta
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/ad-manager/docs/reference/rest/v1/networks.creativeSets

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AdmanagerCreativeSetGVK = GroupVersion.WithKind("AdmanagerCreativeSet")

// AdmanagerCreativeSetSpec defines the desired state of AdmanagerCreativeSet
// +kcc:spec:proto=google.ads.admanager.v1.CreativeSet
// +kcc:required-from-proto
type AdmanagerCreativeSetSpec struct {

	// The AdmanagerCreativeSet name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. The name of the `CreativeSet`. This attribute has a maximum
	//  length of 255 characters.
	// +kcc:proto:field=google.ads.admanager.v1.CreativeSet.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Required. Immutable. The master
	//  [Creative][google.ads.admanager.v1.Creative] to which the `CreativeSet` is
	//  associated.
	// +kcc:proto:field=google.ads.admanager.v1.CreativeSet.master_creative
	// +required
	MasterCreative *string `json:"masterCreative,omitempty"`

	// Required. The resource names of the companion `Creative`s associated with
	//  this `CreativeSet`. Format: `networks/{network_code}/creatives/{creative}`
	// +kcc:proto:field=google.ads.admanager.v1.CreativeSet.companion_creatives
	// +required
	CompanionCreatives []string `json:"companionCreatives,omitempty"`
}

// AdmanagerCreativeSetStatus defines the config connector machine state of AdmanagerCreativeSet
type AdmanagerCreativeSetStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AdmanagerCreativeSet resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *AdmanagerCreativeSetObservedState `json:"observedState,omitempty"`
}

// AdmanagerCreativeSetObservedState is the state of the AdmanagerCreativeSet resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.ads.admanager.v1.CreativeSet
type AdmanagerCreativeSetObservedState struct {
	// Output only. The time this `CreativeSet` was last modified.
	// +kcc:proto:field=google.ads.admanager.v1.CreativeSet.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpadmanagercreativeset;gcpadmanagercreativesets
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AdmanagerCreativeSet is the Schema for the AdmanagerCreativeSet API
// +k8s:openapi-gen=true
type AdmanagerCreativeSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AdmanagerCreativeSetSpec   `json:"spec,omitempty"`
	Status AdmanagerCreativeSetStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AdmanagerCreativeSetList contains a list of AdmanagerCreativeSet
type AdmanagerCreativeSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AdmanagerCreativeSet `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AdmanagerCreativeSet{}, &AdmanagerCreativeSetList{})
}
