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

// API sources for AdmanagerCustomField, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/ads/admanager/v1/custom_field_messages.proto
// +kcc:source:service-docs=https://developers.google.com/ad-manager/api/beta
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/ad-manager/docs/reference/rest/v1/networks.customFields

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AdmanagerCustomFieldGVK = GroupVersion.WithKind("AdmanagerCustomField")

// AdmanagerCustomFieldSpec defines the desired state of AdmanagerCustomField
// +kcc:spec:proto=google.ads.admanager.v1.CustomField
// +kcc:required-from-proto
type AdmanagerCustomFieldSpec struct {

	// The AdmanagerCustomField name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. Name of the CustomField. The max length is 127 characters.
	// +kcc:proto:field=google.ads.admanager.v1.CustomField.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. A description of the custom field. The maximum length is 511
	//  characters.
	// +kcc:proto:field=google.ads.admanager.v1.CustomField.description
	Description *string `json:"description,omitempty"`

	// Required. The type of entity the `CustomField` can be applied to.
	// +kcc:proto:field=google.ads.admanager.v1.CustomField.entity_type
	// +required
	EntityType *string `json:"entityType,omitempty"`

	// Required. The data type of the `CustomField`.
	// +kcc:proto:field=google.ads.admanager.v1.CustomField.data_type
	// +required
	DataType *string `json:"dataType,omitempty"`

	// Required. The visibility of the `CustomField`.
	// +kcc:proto:field=google.ads.admanager.v1.CustomField.visibility
	// +required
	Visibility *string `json:"visibility,omitempty"`

	// Optional. The drop-down options for the `CustomField`.
	//
	//  Only applicable for `CustomField` with the drop-down data type.
	// +kcc:proto:field=google.ads.admanager.v1.CustomField.options
	Options []CustomFieldOption `json:"options,omitempty"`
}

// AdmanagerCustomFieldStatus defines the config connector machine state of AdmanagerCustomField
type AdmanagerCustomFieldStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AdmanagerCustomField resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *AdmanagerCustomFieldObservedState `json:"observedState,omitempty"`
}

// AdmanagerCustomFieldObservedState is the state of the AdmanagerCustomField resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.ads.admanager.v1.CustomField
type AdmanagerCustomFieldObservedState struct {
	// Output only. Deprecated: Unique ID of the CustomField.
	// +kcc:proto:field=google.ads.admanager.v1.CustomField.custom_field_id
	CustomFieldID *int64 `json:"customFieldID,omitempty"`

	// Output only. The status of the `CustomField`.
	// +kcc:proto:field=google.ads.admanager.v1.CustomField.status
	Status *string `json:"status,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpadmanagercustomfield;gcpadmanagercustomfields
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AdmanagerCustomField is the Schema for the AdmanagerCustomField API
// +k8s:openapi-gen=true
type AdmanagerCustomField struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AdmanagerCustomFieldSpec   `json:"spec,omitempty"`
	Status AdmanagerCustomFieldStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AdmanagerCustomFieldList contains a list of AdmanagerCustomField
type AdmanagerCustomFieldList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AdmanagerCustomField `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AdmanagerCustomField{}, &AdmanagerCustomFieldList{})
}
