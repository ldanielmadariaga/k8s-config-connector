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

// API sources for AdmanagerContact, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/ads/admanager/v1/contact_messages.proto
// +kcc:source:service-docs=https://developers.google.com/ad-manager/api/beta
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/ad-manager/docs/reference/rest/v1/networks.contacts

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AdmanagerContactGVK = GroupVersion.WithKind("AdmanagerContact")

// AdmanagerContactSpec defines the desired state of AdmanagerContact
// +kcc:spec:proto=google.ads.admanager.v1.Contact
// +kcc:required-from-proto
type AdmanagerContactSpec struct {

	// The AdmanagerContact name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. The name of the contact. This attribute has a maximum length of
	//  127 characters.
	// +kcc:proto:field=google.ads.admanager.v1.Contact.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Required. Immutable. The resource name of the Company.
	//  Format: "networks/{network_code}/companies/{company_id}"
	// +kcc:proto:field=google.ads.admanager.v1.Contact.company
	// +required
	Company *string `json:"company,omitempty"`

	// Optional. The address of the contact. This attribute has a maximum length
	//  of 1024 characters.
	// +kcc:proto:field=google.ads.admanager.v1.Contact.address
	Address *string `json:"address,omitempty"`

	// Optional. The cell phone number where the contact can be reached.
	// +kcc:proto:field=google.ads.admanager.v1.Contact.cell_phone
	CellPhone *string `json:"cellPhone,omitempty"`

	// Optional. A free-form text comment for the contact. This attribute has a
	//  maximum length of 1024 characters.
	// +kcc:proto:field=google.ads.admanager.v1.Contact.comment
	Comment *string `json:"comment,omitempty"`

	// Optional. The e-mail address where the contact can be reached. This
	//  attribute has a maximum length of 128 characters.
	// +kcc:proto:field=google.ads.admanager.v1.Contact.email
	Email *string `json:"email,omitempty"`

	// Optional. The fax number where the contact can be reached. This attribute
	//  has a maximum length of 1024 characters.
	// +kcc:proto:field=google.ads.admanager.v1.Contact.fax
	Fax *string `json:"fax,omitempty"`

	// Optional. The job title of the contact. This attribute has a maximum length
	//  of 1024 characters.
	// +kcc:proto:field=google.ads.admanager.v1.Contact.title
	Title *string `json:"title,omitempty"`

	// Optional. The work phone number where the contact can be reached. This
	//  attribute has a maximum length of 1024 characters.
	// +kcc:proto:field=google.ads.admanager.v1.Contact.work_phone
	WorkPhone *string `json:"workPhone,omitempty"`
}

// AdmanagerContactStatus defines the config connector machine state of AdmanagerContact
type AdmanagerContactStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AdmanagerContact resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *AdmanagerContactObservedState `json:"observedState,omitempty"`
}

// AdmanagerContactObservedState is the state of the AdmanagerContact resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.ads.admanager.v1.Contact
type AdmanagerContactObservedState struct {
	// Output only. The status of the contact. This attribute is assigned by
	//  Google.
	// +kcc:proto:field=google.ads.admanager.v1.Contact.status
	Status *string `json:"status,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpadmanagercontact;gcpadmanagercontacts
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AdmanagerContact is the Schema for the AdmanagerContact API
// +k8s:openapi-gen=true
type AdmanagerContact struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AdmanagerContactSpec   `json:"spec,omitempty"`
	Status AdmanagerContactStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AdmanagerContactList contains a list of AdmanagerContact
type AdmanagerContactList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AdmanagerContact `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AdmanagerContact{}, &AdmanagerContactList{})
}
