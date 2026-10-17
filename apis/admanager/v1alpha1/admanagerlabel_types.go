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

// API sources for AdmanagerLabel, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/ads/admanager/v1/label_messages.proto
// +kcc:source:service-docs=https://developers.google.com/ad-manager/api/beta
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/ad-manager/docs/reference/rest/v1/networks.labels

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AdmanagerLabelGVK = GroupVersion.WithKind("AdmanagerLabel")

// AdmanagerLabelSpec defines the desired state of AdmanagerLabel
// +kcc:spec:proto=google.ads.admanager.v1.Label
// +kcc:required-from-proto
type AdmanagerLabelSpec struct {

	// The AdmanagerLabel name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. Display name of the Label. This attribute has a maximum length of
	//  127 characters.
	// +kcc:proto:field=google.ads.admanager.v1.Label.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. A description of the label. This attribute has a maximum length
	//  of 255 characters.
	// +kcc:proto:field=google.ads.admanager.v1.Label.description
	Description *string `json:"description,omitempty"`

	// Required. Unordered list. The types of the Label.
	// +kcc:proto:field=google.ads.admanager.v1.Label.types
	// +required
	Types []string `json:"types,omitempty"`
}

// AdmanagerLabelStatus defines the config connector machine state of AdmanagerLabel
type AdmanagerLabelStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AdmanagerLabel resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *AdmanagerLabelObservedState `json:"observedState,omitempty"`
}

// AdmanagerLabelObservedState is the state of the AdmanagerLabel resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.ads.admanager.v1.Label
type AdmanagerLabelObservedState struct {
	// Output only. Specifies whether or not the label is active.
	// +kcc:proto:field=google.ads.admanager.v1.Label.active
	Active *bool `json:"active,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpadmanagerlabel;gcpadmanagerlabels
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AdmanagerLabel is the Schema for the AdmanagerLabel API
// +k8s:openapi-gen=true
type AdmanagerLabel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AdmanagerLabelSpec   `json:"spec,omitempty"`
	Status AdmanagerLabelStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AdmanagerLabelList contains a list of AdmanagerLabel
type AdmanagerLabelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AdmanagerLabel `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AdmanagerLabel{}, &AdmanagerLabelList{})
}
