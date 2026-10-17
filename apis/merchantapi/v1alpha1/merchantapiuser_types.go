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

// API sources for MerchantapiUser, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/shopping/merchant/accounts/v1/user.proto
// +kcc:source:service-docs=https://developers.google.com/merchant/api
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/merchant/api/docs/reference/rest/v1/accounts.users

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var MerchantapiUserGVK = GroupVersion.WithKind("MerchantapiUser")

// MerchantapiUserSpec defines the desired state of MerchantapiUser
// +kcc:spec:proto=google.shopping.merchant.accounts.v1.User
// +kcc:required-from-proto
type MerchantapiUserSpec struct {

	// The MerchantapiUser name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. The [access
	//  rights](https://support.google.com/merchants/answer/12160472?sjid=6789834943175119429-EU#accesstypes)
	//  the user has.
	// +kcc:proto:field=google.shopping.merchant.accounts.v1.User.access_rights
	// +required
	AccessRights []string `json:"accessRights,omitempty"`
}

// MerchantapiUserStatus defines the config connector machine state of MerchantapiUser
type MerchantapiUserStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the MerchantapiUser resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *MerchantapiUserObservedState `json:"observedState,omitempty"`
}

// MerchantapiUserObservedState is the state of the MerchantapiUser resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.shopping.merchant.accounts.v1.User
type MerchantapiUserObservedState struct {
	// Output only. The state of the user.
	// +kcc:proto:field=google.shopping.merchant.accounts.v1.User.state
	State *string `json:"state,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpmerchantapiuser;gcpmerchantapiusers
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// MerchantapiUser is the Schema for the MerchantapiUser API
// +k8s:openapi-gen=true
type MerchantapiUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   MerchantapiUserSpec   `json:"spec,omitempty"`
	Status MerchantapiUserStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// MerchantapiUserList contains a list of MerchantapiUser
type MerchantapiUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MerchantapiUser `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MerchantapiUser{}, &MerchantapiUserList{})
}
