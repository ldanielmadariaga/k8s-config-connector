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

// API sources for AdmanagerTeam, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/ads/admanager/v1/team_messages.proto
// +kcc:source:service-docs=https://developers.google.com/ad-manager/api/beta
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/ad-manager/docs/reference/rest/v1/networks.teams

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AdmanagerTeamGVK = GroupVersion.WithKind("AdmanagerTeam")

// AdmanagerTeamSpec defines the desired state of AdmanagerTeam
// +kcc:spec:proto=google.ads.admanager.v1.Team
// +kcc:required-from-proto
type AdmanagerTeamSpec struct {

	// The AdmanagerTeam name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. The name of the Team. This value has a maximum length of 127
	//  characters.
	// +kcc:proto:field=google.ads.admanager.v1.Team.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. The description of the Team. This value has a maximum length of
	//  255 characters.
	// +kcc:proto:field=google.ads.admanager.v1.Team.description
	Description *string `json:"description,omitempty"`

	// Optional. Whether or not users on this team have access to all companies.
	//  If this value is true, then an error will be thrown if an attempt is made
	//  to associate this team with a Company.
	// +kcc:proto:field=google.ads.admanager.v1.Team.all_companies_access
	AllCompaniesAccess *bool `json:"allCompaniesAccess,omitempty"`

	// Optional. Whether or not users on this team have access to all inventory.
	//  If this value is true, then an error will be thrown if an attempt is made
	//  to associate this team with an AdUnit.
	// +kcc:proto:field=google.ads.admanager.v1.Team.all_inventory_access
	AllInventoryAccess *bool `json:"allInventoryAccess,omitempty"`

	// Optional. The default access to orders for users on this team.
	// +kcc:proto:field=google.ads.admanager.v1.Team.access_type
	AccessType *string `json:"accessType,omitempty"`
}

// AdmanagerTeamStatus defines the config connector machine state of AdmanagerTeam
type AdmanagerTeamStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AdmanagerTeam resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *AdmanagerTeamObservedState `json:"observedState,omitempty"`
}

// AdmanagerTeamObservedState is the state of the AdmanagerTeam resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.ads.admanager.v1.Team
type AdmanagerTeamObservedState struct {
	// Output only. The status of the Team. This value determines the visibility
	//  of the team in the UI.
	// +kcc:proto:field=google.ads.admanager.v1.Team.status
	Status *string `json:"status,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpadmanagerteam;gcpadmanagerteams
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AdmanagerTeam is the Schema for the AdmanagerTeam API
// +k8s:openapi-gen=true
type AdmanagerTeam struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AdmanagerTeamSpec   `json:"spec,omitempty"`
	Status AdmanagerTeamStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AdmanagerTeamList contains a list of AdmanagerTeam
type AdmanagerTeamList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AdmanagerTeam `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AdmanagerTeam{}, &AdmanagerTeamList{})
}
