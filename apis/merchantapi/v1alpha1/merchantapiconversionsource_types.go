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

// API sources for MerchantapiConversionSource, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/shopping/merchant/conversions/v1/conversionsources.proto
// +kcc:source:service-docs=https://developers.google.com/merchant/api
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/merchant/api/docs/reference/rest/v1/accounts.conversionSources

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var MerchantapiConversionSourceGVK = GroupVersion.WithKind("MerchantapiConversionSource")

// MerchantapiConversionSourceSpec defines the desired state of MerchantapiConversionSource
// +kcc:spec:proto=google.shopping.merchant.conversions.v1.ConversionSource
// +kcc:required-from-proto
type MerchantapiConversionSourceSpec struct {

	// The MerchantapiConversionSource name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Immutable. Conversion Source of type "Link to Google Analytics Property".
	// +kcc:proto:field=google.shopping.merchant.conversions.v1.ConversionSource.google_analytics_link
	GoogleAnalyticsLink *GoogleAnalyticsLink `json:"googleAnalyticsLink,omitempty"`

	// Conversion Source of type "Merchant Center Tag Destination".
	// +kcc:proto:field=google.shopping.merchant.conversions.v1.ConversionSource.merchant_center_destination
	MerchantCenterDestination *MerchantCenterDestination `json:"merchantCenterDestination,omitempty"`
}

// MerchantapiConversionSourceStatus defines the config connector machine state of MerchantapiConversionSource
type MerchantapiConversionSourceStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the MerchantapiConversionSource resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *MerchantapiConversionSourceObservedState `json:"observedState,omitempty"`
}

// MerchantapiConversionSourceObservedState is the state of the MerchantapiConversionSource resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.shopping.merchant.conversions.v1.ConversionSource
type MerchantapiConversionSourceObservedState struct {
	// Immutable. Conversion Source of type "Link to Google Analytics Property".
	// +kcc:proto:field=google.shopping.merchant.conversions.v1.ConversionSource.google_analytics_link
	GoogleAnalyticsLink *GoogleAnalyticsLinkObservedState `json:"googleAnalyticsLink,omitempty"`

	// Conversion Source of type "Merchant Center Tag Destination".
	// +kcc:proto:field=google.shopping.merchant.conversions.v1.ConversionSource.merchant_center_destination
	MerchantCenterDestination *MerchantCenterDestinationObservedState `json:"merchantCenterDestination,omitempty"`

	// Output only. Current state of this conversion source. Can't be edited
	//  through the API.
	// +kcc:proto:field=google.shopping.merchant.conversions.v1.ConversionSource.state
	State *string `json:"state,omitempty"`

	// Output only. The time when an archived conversion source becomes
	//  permanently deleted and is no longer available to undelete.
	// +kcc:proto:field=google.shopping.merchant.conversions.v1.ConversionSource.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`

	// Output only. Controller of the conversion source.
	// +kcc:proto:field=google.shopping.merchant.conversions.v1.ConversionSource.controller
	Controller *string `json:"controller,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpmerchantapiconversionsource;gcpmerchantapiconversionsources
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// MerchantapiConversionSource is the Schema for the MerchantapiConversionSource API
// +k8s:openapi-gen=true
type MerchantapiConversionSource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   MerchantapiConversionSourceSpec   `json:"spec,omitempty"`
	Status MerchantapiConversionSourceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// MerchantapiConversionSourceList contains a list of MerchantapiConversionSource
type MerchantapiConversionSourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MerchantapiConversionSource `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MerchantapiConversionSource{}, &MerchantapiConversionSourceList{})
}
