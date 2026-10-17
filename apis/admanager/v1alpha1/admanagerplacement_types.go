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

// API sources for AdmanagerPlacement, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/ads/admanager/v1/placement_messages.proto
// +kcc:source:service-docs=https://developers.google.com/ad-manager/api/beta
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/ad-manager/docs/reference/rest/v1/networks.placements

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AdmanagerPlacementGVK = GroupVersion.WithKind("AdmanagerPlacement")

// AdmanagerPlacementSpec defines the desired state of AdmanagerPlacement
// +kcc:spec:proto=google.ads.admanager.v1.Placement
// +kcc:required-from-proto
type AdmanagerPlacementSpec struct {

	// The AdmanagerPlacement name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. The display name of the placement. This attribute has a maximum
	//  length of 255 characters.
	// +kcc:proto:field=google.ads.admanager.v1.Placement.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. A description of the Placement. This attribute has a maximum
	//  length of 65,535 characters.
	// +kcc:proto:field=google.ads.admanager.v1.Placement.description
	Description *string `json:"description,omitempty"`

	// Optional. The resource names of AdUnits that constitute the Placement.
	//  Format: "networks/{network_code}/adUnits/{ad_unit}"
	// +kcc:proto:field=google.ads.admanager.v1.Placement.targeted_ad_units
	TargetedAdUnits []string `json:"targetedAdUnits,omitempty"`
}

// AdmanagerPlacementStatus defines the config connector machine state of AdmanagerPlacement
type AdmanagerPlacementStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AdmanagerPlacement resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *AdmanagerPlacementObservedState `json:"observedState,omitempty"`
}

// AdmanagerPlacementObservedState is the state of the AdmanagerPlacement resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.ads.admanager.v1.Placement
type AdmanagerPlacementObservedState struct {
	// Output only. Deprecated: `Placement` ID.
	// +kcc:proto:field=google.ads.admanager.v1.Placement.placement_id
	PlacementID *int64 `json:"placementID,omitempty"`

	// Output only. A string used to uniquely identify the Placement for purposes
	//  of serving the ad. This attribute is assigned by Google.
	// +kcc:proto:field=google.ads.admanager.v1.Placement.placement_code
	PlacementCode *string `json:"placementCode,omitempty"`

	// Output only. The status of the Placement.
	// +kcc:proto:field=google.ads.admanager.v1.Placement.status
	Status *string `json:"status,omitempty"`

	// Output only. The instant this Placement was last modified.
	// +kcc:proto:field=google.ads.admanager.v1.Placement.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpadmanagerplacement;gcpadmanagerplacements
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AdmanagerPlacement is the Schema for the AdmanagerPlacement API
// +k8s:openapi-gen=true
type AdmanagerPlacement struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AdmanagerPlacementSpec   `json:"spec,omitempty"`
	Status AdmanagerPlacementStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AdmanagerPlacementList contains a list of AdmanagerPlacement
type AdmanagerPlacementList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AdmanagerPlacement `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AdmanagerPlacement{}, &AdmanagerPlacementList{})
}
