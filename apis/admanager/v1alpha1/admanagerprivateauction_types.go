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

// API sources for AdmanagerPrivateAuction, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/ads/admanager/v1/private_auction_messages.proto
// +kcc:source:service-docs=https://developers.google.com/ad-manager/api/beta
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/ad-manager/docs/reference/rest/v1/networks.privateAuctions

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AdmanagerPrivateAuctionGVK = GroupVersion.WithKind("AdmanagerPrivateAuction")

// AdmanagerPrivateAuctionSpec defines the desired state of AdmanagerPrivateAuction
// +kcc:spec:proto=google.ads.admanager.v1.PrivateAuction
// +kcc:required-from-proto
type AdmanagerPrivateAuctionSpec struct {

	// The AdmanagerPrivateAuction name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. Display name of the `PrivateAuction`. This attribute has a
	//  maximum length of 255 bytes.
	// +kcc:proto:field=google.ads.admanager.v1.PrivateAuction.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Description of the `PrivateAuction`. This attribute has a
	//  maximum length of 4096 bytes.
	// +kcc:proto:field=google.ads.admanager.v1.PrivateAuction.description
	Description *string `json:"description,omitempty"`

	// Optional. The resource names of the seller contact users associated with
	//  this `PrivateAuction`. Format: `networks/{network_code}/users/{user_id}`
	// +kcc:proto:field=google.ads.admanager.v1.PrivateAuction.seller_contact_users
	SellerContactUsers []string `json:"sellerContactUsers,omitempty"`
}

// AdmanagerPrivateAuctionStatus defines the config connector machine state of AdmanagerPrivateAuction
type AdmanagerPrivateAuctionStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AdmanagerPrivateAuction resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *AdmanagerPrivateAuctionObservedState `json:"observedState,omitempty"`
}

// AdmanagerPrivateAuctionObservedState is the state of the AdmanagerPrivateAuction resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.ads.admanager.v1.PrivateAuction
type AdmanagerPrivateAuctionObservedState struct {
	// Output only. Deprecated: `PrivateAuction` ID.
	// +kcc:proto:field=google.ads.admanager.v1.PrivateAuction.private_auction_id
	PrivateAuctionID *int64 `json:"privateAuctionID,omitempty"`

	// Output only. Whether the `PrivateAuction` is archived.
	// +kcc:proto:field=google.ads.admanager.v1.PrivateAuction.archived
	Archived *bool `json:"archived,omitempty"`

	// Output only. The instant at which the PrivateAuction was created.
	// +kcc:proto:field=google.ads.admanager.v1.PrivateAuction.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The instant at which the PrivateAuction was last updated.
	// +kcc:proto:field=google.ads.admanager.v1.PrivateAuction.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpadmanagerprivateauction;gcpadmanagerprivateauctions
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AdmanagerPrivateAuction is the Schema for the AdmanagerPrivateAuction API
// +k8s:openapi-gen=true
type AdmanagerPrivateAuction struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AdmanagerPrivateAuctionSpec   `json:"spec,omitempty"`
	Status AdmanagerPrivateAuctionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AdmanagerPrivateAuctionList contains a list of AdmanagerPrivateAuction
type AdmanagerPrivateAuctionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AdmanagerPrivateAuction `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AdmanagerPrivateAuction{}, &AdmanagerPrivateAuctionList{})
}
