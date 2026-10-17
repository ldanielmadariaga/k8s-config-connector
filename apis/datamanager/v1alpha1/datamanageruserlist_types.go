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

// API sources for DatamanagerUserList, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/ads/datamanager/v1/user_list.proto
// +kcc:source:service-docs=https://developers.google.com/data-manager
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/data-manager/docs/reference/rest/v1/accountTypes.accounts.userLists

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DatamanagerUserListGVK = GroupVersion.WithKind("DatamanagerUserList")

// DatamanagerUserListSpec defines the desired state of DatamanagerUserList
// +kcc:spec:proto=google.ads.datamanager.v1.UserList
// +kcc:required-from-proto
type DatamanagerUserListSpec struct {

	// The DatamanagerUserList name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. The display name of the user list.
	// +kcc:proto:field=google.ads.datamanager.v1.UserList.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. A description of the user list.
	// +kcc:proto:field=google.ads.datamanager.v1.UserList.description
	Description *string `json:"description,omitempty"`

	// Optional. Membership status of this user list.
	// +kcc:proto:field=google.ads.datamanager.v1.UserList.membership_status
	MembershipStatus *string `json:"membershipStatus,omitempty"`

	// Optional. An ID from external system. It is used by user list sellers to
	//  correlate IDs on their systems.
	// +kcc:proto:field=google.ads.datamanager.v1.UserList.integration_code
	IntegrationCode *string `json:"integrationCode,omitempty"`

	// Optional. The duration a user remains in the user list.
	//  Valid durations are exact multiples of 24 hours (86400 seconds).
	//  Providing a value that is not an exact multiple of 24 hours will result in
	//  an INVALID_ARGUMENT error.
	// +kcc:proto:field=google.ads.datamanager.v1.UserList.membership_duration
	MembershipDuration *string `json:"membershipDuration,omitempty"`

	// Optional. Indicates if this share is still enabled. When a user list is
	//  shared with the account this field is set to `ENABLED`. Later the user list
	//  owner can decide to revoke the share and make it `DISABLED`.
	// +kcc:proto:field=google.ads.datamanager.v1.UserList.account_access_status
	AccountAccessStatus *string `json:"accountAccessStatus,omitempty"`

	// Optional. Eligibility information for different target networks.
	// +kcc:proto:field=google.ads.datamanager.v1.UserList.target_network_info
	TargetNetworkInfo *TargetNetworkInfo `json:"targetNetworkInfo,omitempty"`

	// Optional. Represents a user list that is populated by user ingested data.
	// +kcc:proto:field=google.ads.datamanager.v1.UserList.ingested_user_list_info
	IngestedUserListInfo *IngestedUserListInfo `json:"ingestedUserListInfo,omitempty"`
}

// DatamanagerUserListStatus defines the config connector machine state of DatamanagerUserList
type DatamanagerUserListStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DatamanagerUserList resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DatamanagerUserListObservedState `json:"observedState,omitempty"`
}

// DatamanagerUserListObservedState is the state of the DatamanagerUserList resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.ads.datamanager.v1.UserList
type DatamanagerUserListObservedState struct {
	// Output only. The unique ID of the user list.
	// +kcc:proto:field=google.ads.datamanager.v1.UserList.id
	ID *int64 `json:"id,omitempty"`

	// Output only. An option that indicates if a user may edit a list.
	// +kcc:proto:field=google.ads.datamanager.v1.UserList.read_only
	ReadOnly *bool `json:"readOnly,omitempty"`

	// Output only. The reason why this user list membership status is closed.
	// +kcc:proto:field=google.ads.datamanager.v1.UserList.closing_reason
	ClosingReason *string `json:"closingReason,omitempty"`

	// Output only. The reason this account has been granted access to the list.
	// +kcc:proto:field=google.ads.datamanager.v1.UserList.access_reason
	AccessReason *string `json:"accessReason,omitempty"`

	// Output only. Estimated number of members in this user list in different
	//  target networks.
	// +kcc:proto:field=google.ads.datamanager.v1.UserList.size_info
	SizeInfo *SizeInfoObservedState `json:"sizeInfo,omitempty"`

	// Optional. Eligibility information for different target networks.
	// +kcc:proto:field=google.ads.datamanager.v1.UserList.target_network_info
	TargetNetworkInfo *TargetNetworkInfoObservedState `json:"targetNetworkInfo,omitempty"`

	// Optional. Represents a user list that is populated by user ingested data.
	// +kcc:proto:field=google.ads.datamanager.v1.UserList.ingested_user_list_info
	IngestedUserListInfo *IngestedUserListInfoObservedState `json:"ingestedUserListInfo,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdatamanageruserlist;gcpdatamanageruserlists
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DatamanagerUserList is the Schema for the DatamanagerUserList API
// +k8s:openapi-gen=true
type DatamanagerUserList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DatamanagerUserListSpec   `json:"spec,omitempty"`
	Status DatamanagerUserListStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DatamanagerUserListList contains a list of DatamanagerUserList
type DatamanagerUserListList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatamanagerUserList `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DatamanagerUserList{}, &DatamanagerUserListList{})
}
