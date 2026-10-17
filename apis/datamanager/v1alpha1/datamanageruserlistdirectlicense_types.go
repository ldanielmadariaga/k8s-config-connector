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

// API sources for DatamanagerUserListDirectLicense, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/ads/datamanager/v1/user_list_direct_license.proto
// +kcc:source:service-docs=https://developers.google.com/data-manager
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/data-manager/docs/reference/rest/v1/accountTypes.accounts.userListDirectLicenses

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DatamanagerUserListDirectLicenseGVK = GroupVersion.WithKind("DatamanagerUserListDirectLicense")

// DatamanagerUserListDirectLicenseSpec defines the desired state of DatamanagerUserListDirectLicense
// +kcc:spec:proto=google.ads.datamanager.v1.UserListDirectLicense
// +kcc:required-from-proto
type DatamanagerUserListDirectLicenseSpec struct {

	// The DatamanagerUserListDirectLicense name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Immutable. ID of the user list being licensed.
	// +kcc:proto:field=google.ads.datamanager.v1.UserListDirectLicense.user_list_id
	UserListID *int64 `json:"userListID,omitempty"`

	// Immutable. Account type of client customer which the user list is being
	//  licensed to.
	// +kcc:proto:field=google.ads.datamanager.v1.UserListDirectLicense.client_account_type
	ClientAccountType *string `json:"clientAccountType,omitempty"`

	// Immutable. ID of client customer which the user list is being licensed to.
	// +kcc:proto:field=google.ads.datamanager.v1.UserListDirectLicense.client_account_id
	ClientAccountID *int64 `json:"clientAccountID,omitempty"`

	// Optional. Status of UserListDirectLicense - ENABLED or DISABLED.
	// +kcc:proto:field=google.ads.datamanager.v1.UserListDirectLicense.status
	Status *string `json:"status,omitempty"`

	// Optional. UserListDirectLicense pricing.
	// +kcc:proto:field=google.ads.datamanager.v1.UserListDirectLicense.pricing
	Pricing *UserListLicensePricing `json:"pricing,omitempty"`
}

// DatamanagerUserListDirectLicenseStatus defines the config connector machine state of DatamanagerUserListDirectLicense
type DatamanagerUserListDirectLicenseStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DatamanagerUserListDirectLicense resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DatamanagerUserListDirectLicenseObservedState `json:"observedState,omitempty"`
}

// DatamanagerUserListDirectLicenseObservedState is the state of the DatamanagerUserListDirectLicense resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.ads.datamanager.v1.UserListDirectLicense
type DatamanagerUserListDirectLicenseObservedState struct {
	// Output only. Name of the user list being licensed.
	//
	//  This field is read-only.
	// +kcc:proto:field=google.ads.datamanager.v1.UserListDirectLicense.user_list_display_name
	UserListDisplayName *string `json:"userListDisplayName,omitempty"`

	// Output only. Name of client customer which the user list is being licensed
	//  to.
	//
	//  This field is read-only.
	// +kcc:proto:field=google.ads.datamanager.v1.UserListDirectLicense.client_account_display_name
	ClientAccountDisplayName *string `json:"clientAccountDisplayName,omitempty"`

	// Optional. UserListDirectLicense pricing.
	// +kcc:proto:field=google.ads.datamanager.v1.UserListDirectLicense.pricing
	Pricing *UserListLicensePricingObservedState `json:"pricing,omitempty"`

	// Output only. Pricing history of this user list license.
	//
	//  This field is read-only.
	// +kcc:proto:field=google.ads.datamanager.v1.UserListDirectLicense.historical_pricings
	HistoricalPricings []UserListLicensePricingObservedState `json:"historicalPricings,omitempty"`

	// Output only. Metrics related to this license
	//
	//  This field is read-only and only populated if the start and end dates are
	//  set in the ListUserListDirectLicenses call
	// +kcc:proto:field=google.ads.datamanager.v1.UserListDirectLicense.metrics
	Metrics *UserListLicenseMetricsObservedState `json:"metrics,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdatamanageruserlistdirectlicense;gcpdatamanageruserlistdirectlicenses
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DatamanagerUserListDirectLicense is the Schema for the DatamanagerUserListDirectLicense API
// +k8s:openapi-gen=true
type DatamanagerUserListDirectLicense struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DatamanagerUserListDirectLicenseSpec   `json:"spec,omitempty"`
	Status DatamanagerUserListDirectLicenseStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DatamanagerUserListDirectLicenseList contains a list of DatamanagerUserListDirectLicense
type DatamanagerUserListDirectLicenseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatamanagerUserListDirectLicense `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DatamanagerUserListDirectLicense{}, &DatamanagerUserListDirectLicenseList{})
}
