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

// API sources for MerchantapiOmnichannelSetting, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/shopping/merchant/accounts/v1/omnichannelsettings.proto
// +kcc:source:service-docs=https://developers.google.com/merchant/api
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/merchant/api/docs/reference/rest/v1/accounts.omnichannelSettings

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var MerchantapiOmnichannelSettingGVK = GroupVersion.WithKind("MerchantapiOmnichannelSetting")

// MerchantapiOmnichannelSettingSpec defines the desired state of MerchantapiOmnichannelSetting
// +kcc:spec:proto=google.shopping.merchant.accounts.v1.OmnichannelSetting
// +kcc:required-from-proto
type MerchantapiOmnichannelSettingSpec struct {

	// The MerchantapiOmnichannelSetting name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. Immutable. Region code defined by
	//  [CLDR](https://cldr.unicode.org/). Must be provided in the Create method,
	//  and is immutable.
	// +kcc:proto:field=google.shopping.merchant.accounts.v1.OmnichannelSetting.region_code
	// +required
	RegionCode *string `json:"regionCode,omitempty"`

	// Required. The Local Store Front type for this country.
	// +kcc:proto:field=google.shopping.merchant.accounts.v1.OmnichannelSetting.lsf_type
	// +required
	LsfType *string `json:"lsfType,omitempty"`

	// Optional. The InStock URI and state for this country.
	// +kcc:proto:field=google.shopping.merchant.accounts.v1.OmnichannelSetting.in_stock
	InStock *InStock `json:"inStock,omitempty"`

	// Optional. The Pickup URI and state for this country.
	// +kcc:proto:field=google.shopping.merchant.accounts.v1.OmnichannelSetting.pickup
	Pickup *Pickup `json:"pickup,omitempty"`

	// Optional. The On Display to Order (ODO) policy URI and state for this
	//  country.
	// +kcc:proto:field=google.shopping.merchant.accounts.v1.OmnichannelSetting.odo
	Odo *OnDisplayToOrder `json:"odo,omitempty"`

	// Optional. The about page URI and state for this country.
	// +kcc:proto:field=google.shopping.merchant.accounts.v1.OmnichannelSetting.about
	About *About `json:"about,omitempty"`

	// Optional. The inventory verification contact and state for this country.
	// +kcc:proto:field=google.shopping.merchant.accounts.v1.OmnichannelSetting.inventory_verification
	InventoryVerification *InventoryVerification `json:"inventoryVerification,omitempty"`
}

// MerchantapiOmnichannelSettingStatus defines the config connector machine state of MerchantapiOmnichannelSetting
type MerchantapiOmnichannelSettingStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the MerchantapiOmnichannelSetting resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *MerchantapiOmnichannelSettingObservedState `json:"observedState,omitempty"`
}

// MerchantapiOmnichannelSettingObservedState is the state of the MerchantapiOmnichannelSetting resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.shopping.merchant.accounts.v1.OmnichannelSetting
type MerchantapiOmnichannelSettingObservedState struct {
	// Optional. The InStock URI and state for this country.
	// +kcc:proto:field=google.shopping.merchant.accounts.v1.OmnichannelSetting.in_stock
	InStock *InStockObservedState `json:"inStock,omitempty"`

	// Optional. The Pickup URI and state for this country.
	// +kcc:proto:field=google.shopping.merchant.accounts.v1.OmnichannelSetting.pickup
	Pickup *PickupObservedState `json:"pickup,omitempty"`

	// Output only. The established link to a LFP provider.
	// +kcc:proto:field=google.shopping.merchant.accounts.v1.OmnichannelSetting.lfp_link
	LfpLink *LfpLinkObservedState `json:"lfpLink,omitempty"`

	// Optional. The On Display to Order (ODO) policy URI and state for this
	//  country.
	// +kcc:proto:field=google.shopping.merchant.accounts.v1.OmnichannelSetting.odo
	Odo *OnDisplayToOrderObservedState `json:"odo,omitempty"`

	// Optional. The about page URI and state for this country.
	// +kcc:proto:field=google.shopping.merchant.accounts.v1.OmnichannelSetting.about
	About *AboutObservedState `json:"about,omitempty"`

	// Optional. The inventory verification contact and state for this country.
	// +kcc:proto:field=google.shopping.merchant.accounts.v1.OmnichannelSetting.inventory_verification
	InventoryVerification *InventoryVerificationObservedState `json:"inventoryVerification,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpmerchantapiomnichannelsetting;gcpmerchantapiomnichannelsettings
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// MerchantapiOmnichannelSetting is the Schema for the MerchantapiOmnichannelSetting API
// +k8s:openapi-gen=true
type MerchantapiOmnichannelSetting struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   MerchantapiOmnichannelSettingSpec   `json:"spec,omitempty"`
	Status MerchantapiOmnichannelSettingStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// MerchantapiOmnichannelSettingList contains a list of MerchantapiOmnichannelSetting
type MerchantapiOmnichannelSettingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MerchantapiOmnichannelSetting `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MerchantapiOmnichannelSetting{}, &MerchantapiOmnichannelSettingList{})
}
