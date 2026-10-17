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

// API sources for MerchantapiOnlineReturnPolicy, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/shopping/merchant/accounts/v1/online_return_policy.proto
// +kcc:source:service-docs=https://developers.google.com/merchant/api
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/merchant/api/docs/reference/rest/v1/accounts.onlineReturnPolicies

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var MerchantapiOnlineReturnPolicyGVK = GroupVersion.WithKind("MerchantapiOnlineReturnPolicy")

// MerchantapiOnlineReturnPolicySpec defines the desired state of MerchantapiOnlineReturnPolicy
// +kcc:spec:proto=google.shopping.merchant.accounts.v1.OnlineReturnPolicy
// +kcc:required-from-proto
type MerchantapiOnlineReturnPolicySpec struct {

	// The MerchantapiOnlineReturnPolicy name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. Immutable. This field represents the unique user-defined label of
	//  the return policy for the given country. It is important to note that the
	//  same label cannot be used in different return policies for the same
	//  country. If not given, policies will be automatically treated as the
	//  'default' for the country. When using label, you are creating an exception
	//  policy in that country to assign a custom return policy to certain product
	//  groups, follow the instructions provided in the [Return policy label]
	//  (https://support.google.com/merchants/answer/9445425). The label can
	//  contain up to 50 characters.
	// +kcc:proto:field=google.shopping.merchant.accounts.v1.OnlineReturnPolicy.label
	Label *string `json:"label,omitempty"`

	// Required. Immutable. The countries of sale where the return policy applies.
	//  The values must be a valid 2 letter ISO 3166 code.
	// +kcc:proto:field=google.shopping.merchant.accounts.v1.OnlineReturnPolicy.countries
	// +required
	Countries []string `json:"countries,omitempty"`

	// Optional. The return policy.
	// +kcc:proto:field=google.shopping.merchant.accounts.v1.OnlineReturnPolicy.policy
	Policy *OnlineReturnPolicy_Policy `json:"policy,omitempty"`

	// Optional. Overrides to the general policy for orders placed during a
	//  specific set of time intervals.
	// +kcc:proto:field=google.shopping.merchant.accounts.v1.OnlineReturnPolicy.seasonal_overrides
	SeasonalOverrides []OnlineReturnPolicy_SeasonalOverride `json:"seasonalOverrides,omitempty"`

	// Optional. The restocking fee that applies to all return reason categories.
	//  This would be treated as a free restocking fee if the value is not set.
	// +kcc:proto:field=google.shopping.merchant.accounts.v1.OnlineReturnPolicy.restocking_fee
	RestockingFee *OnlineReturnPolicy_RestockingFee `json:"restockingFee,omitempty"`

	// Optional. The return methods of how customers can return an item. This
	//  value is required to not be empty unless the type of return policy is
	//  noReturns.
	// +kcc:proto:field=google.shopping.merchant.accounts.v1.OnlineReturnPolicy.return_methods
	ReturnMethods []string `json:"returnMethods,omitempty"`

	// Optional. The item conditions accepted for returns must not be empty unless
	//  the type of return policy is 'noReturns'.
	// +kcc:proto:field=google.shopping.merchant.accounts.v1.OnlineReturnPolicy.item_conditions
	ItemConditions []string `json:"itemConditions,omitempty"`

	// Optional. The return shipping fee. Should be set only when customer need to
	//  download and print the return label.
	// +kcc:proto:field=google.shopping.merchant.accounts.v1.OnlineReturnPolicy.return_shipping_fee
	ReturnShippingFee *OnlineReturnPolicy_ReturnShippingFee `json:"returnShippingFee,omitempty"`

	// Required. The return policy uri. This can used by Google to do a sanity
	//  check for the policy. It must be a valid URL.
	// +kcc:proto:field=google.shopping.merchant.accounts.v1.OnlineReturnPolicy.return_policy_uri
	// +required
	ReturnPolicyURI *string `json:"returnPolicyURI,omitempty"`

	// Optional. This field specifies if business
	//  only accepts defective products for returns.
	// +kcc:proto:field=google.shopping.merchant.accounts.v1.OnlineReturnPolicy.accept_defective_only
	AcceptDefectiveOnly *bool `json:"acceptDefectiveOnly,omitempty"`

	// Optional. The field specifies the number of
	//  days it takes for business to process refunds.
	// +kcc:proto:field=google.shopping.merchant.accounts.v1.OnlineReturnPolicy.process_refund_days
	ProcessRefundDays *int32 `json:"processRefundDays,omitempty"`

	// Optional. This field specifies if business
	//  allows customers to exchange products.
	// +kcc:proto:field=google.shopping.merchant.accounts.v1.OnlineReturnPolicy.accept_exchange
	AcceptExchange *bool `json:"acceptExchange,omitempty"`

	// Optional. The field specifies the return
	//  label source.
	// +kcc:proto:field=google.shopping.merchant.accounts.v1.OnlineReturnPolicy.return_label_source
	ReturnLabelSource *string `json:"returnLabelSource,omitempty"`
}

// MerchantapiOnlineReturnPolicyStatus defines the config connector machine state of MerchantapiOnlineReturnPolicy
type MerchantapiOnlineReturnPolicyStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the MerchantapiOnlineReturnPolicy resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *MerchantapiOnlineReturnPolicyObservedState `json:"observedState,omitempty"`
}

// MerchantapiOnlineReturnPolicyObservedState is the state of the MerchantapiOnlineReturnPolicy resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.shopping.merchant.accounts.v1.OnlineReturnPolicy
type MerchantapiOnlineReturnPolicyObservedState struct {
	// Output only. Return policy ID generated by Google.
	// +kcc:proto:field=google.shopping.merchant.accounts.v1.OnlineReturnPolicy.return_policy_id
	ReturnPolicyID *string `json:"returnPolicyID,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpmerchantapionlinereturnpolicy;gcpmerchantapionlinereturnpolicies
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// MerchantapiOnlineReturnPolicy is the Schema for the MerchantapiOnlineReturnPolicy API
// +k8s:openapi-gen=true
type MerchantapiOnlineReturnPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   MerchantapiOnlineReturnPolicySpec   `json:"spec,omitempty"`
	Status MerchantapiOnlineReturnPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// MerchantapiOnlineReturnPolicyList contains a list of MerchantapiOnlineReturnPolicy
type MerchantapiOnlineReturnPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MerchantapiOnlineReturnPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MerchantapiOnlineReturnPolicy{}, &MerchantapiOnlineReturnPolicyList{})
}
