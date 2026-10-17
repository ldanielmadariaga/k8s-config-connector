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

// API sources for DatamanagerUserListGlobalLicense, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/ads/datamanager/v1/user_list_global_license.proto
// +kcc:source:service-docs=https://developers.google.com/data-manager
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/data-manager/docs/reference/rest/v1/accountTypes.accounts.userListGlobalLicenses

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DatamanagerUserListGlobalLicenseGVK = GroupVersion.WithKind("DatamanagerUserListGlobalLicense")

// DatamanagerUserListGlobalLicenseSpec defines the desired state of DatamanagerUserListGlobalLicense
// +kcc:spec:proto=google.ads.datamanager.v1.UserListGlobalLicense
// +kcc:required-from-proto
type DatamanagerUserListGlobalLicenseSpec struct {

	// The DatamanagerUserListGlobalLicense name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Immutable. ID of the user list being licensed.
	// +kcc:proto:field=google.ads.datamanager.v1.UserListGlobalLicense.user_list_id
	UserListID *int64 `json:"userListID,omitempty"`

	// Immutable. Product type of client customer which the user list is being
	//  licensed to.
	// +kcc:proto:field=google.ads.datamanager.v1.UserListGlobalLicense.license_type
	LicenseType *string `json:"licenseType,omitempty"`

	// Optional. Status of UserListGlobalLicense - ENABLED or DISABLED.
	// +kcc:proto:field=google.ads.datamanager.v1.UserListGlobalLicense.status
	Status *string `json:"status,omitempty"`

	// Optional. UserListGlobalLicense pricing.
	// +kcc:proto:field=google.ads.datamanager.v1.UserListGlobalLicense.pricing
	Pricing *UserListLicensePricing `json:"pricing,omitempty"`
}

// DatamanagerUserListGlobalLicenseStatus defines the config connector machine state of DatamanagerUserListGlobalLicense
type DatamanagerUserListGlobalLicenseStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DatamanagerUserListGlobalLicense resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DatamanagerUserListGlobalLicenseObservedState `json:"observedState,omitempty"`
}

// DatamanagerUserListGlobalLicenseObservedState is the state of the DatamanagerUserListGlobalLicense resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.ads.datamanager.v1.UserListGlobalLicense
type DatamanagerUserListGlobalLicenseObservedState struct {
	// Output only. Name of the user list being licensed.
	//
	//  This field is read-only.
	// +kcc:proto:field=google.ads.datamanager.v1.UserListGlobalLicense.user_list_display_name
	UserListDisplayName *string `json:"userListDisplayName,omitempty"`

	// Optional. UserListGlobalLicense pricing.
	// +kcc:proto:field=google.ads.datamanager.v1.UserListGlobalLicense.pricing
	Pricing *UserListLicensePricingObservedState `json:"pricing,omitempty"`

	// Output only. Pricing history of this user list license.
	//
	//  This field is read-only.
	// +kcc:proto:field=google.ads.datamanager.v1.UserListGlobalLicense.historical_pricings
	HistoricalPricings []UserListLicensePricingObservedState `json:"historicalPricings,omitempty"`

	// Output only. Metrics related to this license
	//
	//  This field is read-only and only populated if the start and end dates are
	//  set in the ListUserListGlobalLicenses call
	// +kcc:proto:field=google.ads.datamanager.v1.UserListGlobalLicense.metrics
	Metrics *UserListLicenseMetricsObservedState `json:"metrics,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdatamanageruserlistgloballicense;gcpdatamanageruserlistgloballicenses
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DatamanagerUserListGlobalLicense is the Schema for the DatamanagerUserListGlobalLicense API
// +k8s:openapi-gen=true
type DatamanagerUserListGlobalLicense struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DatamanagerUserListGlobalLicenseSpec   `json:"spec,omitempty"`
	Status DatamanagerUserListGlobalLicenseStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DatamanagerUserListGlobalLicenseList contains a list of DatamanagerUserListGlobalLicense
type DatamanagerUserListGlobalLicenseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatamanagerUserListGlobalLicense `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DatamanagerUserListGlobalLicense{}, &DatamanagerUserListGlobalLicenseList{})
}
