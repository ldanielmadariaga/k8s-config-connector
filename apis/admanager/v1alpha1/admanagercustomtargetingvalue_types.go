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

// API sources for AdmanagerCustomTargetingValue, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/ads/admanager/v1/custom_targeting_value_messages.proto
// +kcc:source:service-docs=https://developers.google.com/ad-manager/api/beta
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/ad-manager/docs/reference/rest/v1/networks.customTargetingValues

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AdmanagerCustomTargetingValueGVK = GroupVersion.WithKind("AdmanagerCustomTargetingValue")

// AdmanagerCustomTargetingValueSpec defines the desired state of AdmanagerCustomTargetingValue
// +kcc:spec:proto=google.ads.admanager.v1.CustomTargetingValue
// +kcc:required-from-proto
type AdmanagerCustomTargetingValueSpec struct {

	// The AdmanagerCustomTargetingValue name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. Immutable. The resource name of the `CustomTargetingKey`.
	//  Format:
	//  `networks/{network_code}/customTargetingKeys/{custom_targeting_key_id}`
	// +kcc:guess=possible-reference target=AdmanagerCustomTargetingKey
	// +kcc:proto:field=google.ads.admanager.v1.CustomTargetingValue.custom_targeting_key
	// +required
	CustomTargetingKey *string `json:"customTargetingKey,omitempty"`

	// Immutable. Name of the `CustomTargetingValue`. Values can contain up to 40
	//  characters each. You can use alphanumeric characters and symbols other than
	//  the following: ", ', =, !, +, #, *, ~, ;, ^, (, ), <, >, [, ]. Values are
	//  not data-specific; all values are treated as strings. For example, instead
	//  of using "age>=18 AND <=34", try "18-34"
	// +kcc:proto:field=google.ads.admanager.v1.CustomTargetingValue.ad_tag_name
	AdTagName *string `json:"adTagName,omitempty"`

	// Optional. Descriptive name for the `CustomTargetingValue`.
	// +kcc:proto:field=google.ads.admanager.v1.CustomTargetingValue.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Required. Immutable. The way in which the CustomTargetingValue.name strings
	//  will be matched.
	// +kcc:proto:field=google.ads.admanager.v1.CustomTargetingValue.match_type
	// +required
	MatchType *string `json:"matchType,omitempty"`
}

// AdmanagerCustomTargetingValueStatus defines the config connector machine state of AdmanagerCustomTargetingValue
type AdmanagerCustomTargetingValueStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AdmanagerCustomTargetingValue resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *AdmanagerCustomTargetingValueObservedState `json:"observedState,omitempty"`
}

// AdmanagerCustomTargetingValueObservedState is the state of the AdmanagerCustomTargetingValue resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.ads.admanager.v1.CustomTargetingValue
type AdmanagerCustomTargetingValueObservedState struct {
	// Output only. Status of the `CustomTargetingValue`.
	// +kcc:proto:field=google.ads.admanager.v1.CustomTargetingValue.status
	Status *string `json:"status,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpadmanagercustomtargetingvalue;gcpadmanagercustomtargetingvalues
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AdmanagerCustomTargetingValue is the Schema for the AdmanagerCustomTargetingValue API
// +k8s:openapi-gen=true
type AdmanagerCustomTargetingValue struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AdmanagerCustomTargetingValueSpec   `json:"spec,omitempty"`
	Status AdmanagerCustomTargetingValueStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AdmanagerCustomTargetingValueList contains a list of AdmanagerCustomTargetingValue
type AdmanagerCustomTargetingValueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AdmanagerCustomTargetingValue `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AdmanagerCustomTargetingValue{}, &AdmanagerCustomTargetingValueList{})
}
