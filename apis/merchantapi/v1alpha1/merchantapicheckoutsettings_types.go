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

// API sources for MerchantapiCheckoutSettings, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/shopping/merchant/accounts/v1/checkoutsettings.proto
// +kcc:source:service-docs=https://developers.google.com/merchant/api
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/merchant/api/docs/reference/rest/v1/accounts.programs.checkoutSettings

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var MerchantapiCheckoutSettingsGVK = GroupVersion.WithKind("MerchantapiCheckoutSettings")

// MerchantapiCheckoutSettingsSpec defines the desired state of MerchantapiCheckoutSettings
// +kcc:spec:proto=google.shopping.merchant.accounts.v1.CheckoutSettings
// +kcc:required-from-proto
type MerchantapiCheckoutSettingsSpec struct {

	// The MerchantapiCheckoutSettings name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// URI settings for cart or checkout URL.
	// +kcc:proto:field=google.shopping.merchant.accounts.v1.CheckoutSettings.uri_settings
	URISettings *URISettings `json:"uriSettings,omitempty"`

	// Optional. The destinations (also known as [Marketing
	//  methods](https://support.google.com/merchants/answer/15130232)) to which
	//  the checkout program applies, valid destination values are `SHOPPING_ADS`,
	//  `FREE_LISTINGS`
	// +kcc:proto:field=google.shopping.merchant.accounts.v1.CheckoutSettings.eligible_destinations
	EligibleDestinations []string `json:"eligibleDestinations,omitempty"`
}

// MerchantapiCheckoutSettingsStatus defines the config connector machine state of MerchantapiCheckoutSettings
type MerchantapiCheckoutSettingsStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the MerchantapiCheckoutSettings resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *MerchantapiCheckoutSettingsObservedState `json:"observedState,omitempty"`
}

// MerchantapiCheckoutSettingsObservedState is the state of the MerchantapiCheckoutSettings resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.shopping.merchant.accounts.v1.CheckoutSettings
type MerchantapiCheckoutSettingsObservedState struct {
	// Output only. Reflects the merchant enrollment state in `Checkout` program.
	// +kcc:proto:field=google.shopping.merchant.accounts.v1.CheckoutSettings.enrollment_state
	EnrollmentState *string `json:"enrollmentState,omitempty"`

	// Output only. Reflects the merchant review state in `Checkout` program.
	//  This is set based on the data quality reviews of the URL provided by
	//  the merchant.
	//  A merchant with enrollment state
	//  as `ENROLLED` can be in the following review states: `IN_REVIEW`,
	//  `APPROVED` or `DISAPPROVED`. A merchant must be in an `enrollment_state` of
	//  `ENROLLED` before a review can begin for the merchant.For more details,
	//  check the help center doc.
	// +kcc:proto:field=google.shopping.merchant.accounts.v1.CheckoutSettings.review_state
	ReviewState *string `json:"reviewState,omitempty"`

	// Output only. The effective value of `uri_settings` for a given merchant. If
	//  account level settings are present then this value will be a copy of url
	//  settings. Otherwise, it will have the value of the parent account (for only
	//  marketplace sellers).
	// +kcc:proto:field=google.shopping.merchant.accounts.v1.CheckoutSettings.effective_uri_settings
	EffectiveURISettings *URISettings `json:"effectiveURISettings,omitempty"`

	// Output only. The effective value of enrollment_state for a given merchant
	//  ID. If account level settings are present then this value will be a copy of
	//  the account level settings. Otherwise, it will have the value of the parent
	//  account (for only marketplace sellers).
	// +kcc:proto:field=google.shopping.merchant.accounts.v1.CheckoutSettings.effective_enrollment_state
	EffectiveEnrollmentState *string `json:"effectiveEnrollmentState,omitempty"`

	// Output only. The effective value of `review_state` for a given merchant ID.
	//  If account level settings are present then this value will be a copy of the
	//  account level settings. Otherwise, it will have the value of the parent
	//  account (for only marketplace sellers).
	// +kcc:proto:field=google.shopping.merchant.accounts.v1.CheckoutSettings.effective_review_state
	EffectiveReviewState *string `json:"effectiveReviewState,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpmerchantapicheckoutsettings;gcpmerchantapicheckoutsettings
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// MerchantapiCheckoutSettings is the Schema for the MerchantapiCheckoutSettings API
// +k8s:openapi-gen=true
type MerchantapiCheckoutSettings struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   MerchantapiCheckoutSettingsSpec   `json:"spec,omitempty"`
	Status MerchantapiCheckoutSettingsStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// MerchantapiCheckoutSettingsList contains a list of MerchantapiCheckoutSettings
type MerchantapiCheckoutSettingsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MerchantapiCheckoutSettings `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MerchantapiCheckoutSettings{}, &MerchantapiCheckoutSettingsList{})
}
