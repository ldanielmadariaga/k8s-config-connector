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

// API sources for AdmanagerPrivateAuctionDeal, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/ads/admanager/v1/private_auction_deal_messages.proto
// +kcc:source:service-docs=https://developers.google.com/ad-manager/api/beta
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/ad-manager/docs/reference/rest/v1/networks.privateAuctionDeals

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AdmanagerPrivateAuctionDealGVK = GroupVersion.WithKind("AdmanagerPrivateAuctionDeal")

// AdmanagerPrivateAuctionDealSpec defines the desired state of AdmanagerPrivateAuctionDeal
// +kcc:spec:proto=google.ads.admanager.v1.PrivateAuctionDeal
// +kcc:required-from-proto
type AdmanagerPrivateAuctionDealSpec struct {

	// The AdmanagerPrivateAuctionDeal name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Immutable. The ID of the
	//  [PrivateAuction][google.ads.admanager.v1.PrivateAuction].
	// +kcc:proto:field=google.ads.admanager.v1.PrivateAuctionDeal.private_auction_id
	PrivateAuctionID *int64 `json:"privateAuctionID,omitempty"`

	// Immutable. The account ID of the buyer of the `PrivateAuctionDeal`.
	// +kcc:proto:field=google.ads.admanager.v1.PrivateAuctionDeal.buyer_account_id
	BuyerAccountID *int64 `json:"buyerAccountID,omitempty"`

	// Optional. The targeting of the `PrivateAuctionDeal`.
	// +kcc:proto:field=google.ads.admanager.v1.PrivateAuctionDeal.targeting
	Targeting *Targeting `json:"targeting,omitempty"`

	// Optional. Non-empty default. The start time of the `PrivateAuctionDeal`.
	//  If unset, the `startTime` will default as follows:
	//  On create: To the deal's `create_time`.
	//  On update: To the deal's existing `start_time`, which can be null.
	// +kcc:proto:field=google.ads.admanager.v1.PrivateAuctionDeal.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Optional. The end time of the `PrivateAuctionDeal`.
	// +kcc:proto:field=google.ads.admanager.v1.PrivateAuctionDeal.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Required. The floor price of the `PrivateAuctionDeal`.
	// +kcc:proto:field=google.ads.admanager.v1.PrivateAuctionDeal.floor_price
	// +required
	FloorPrice *Money `json:"floorPrice,omitempty"`

	// Optional. The expected creative sizes of the `PrivateAuctionDeal`.
	// +kcc:proto:field=google.ads.admanager.v1.PrivateAuctionDeal.creative_sizes
	CreativeSizes []Size `json:"creativeSizes,omitempty"`

	// Optional. Whether the deal is enabled with priority over open auction.
	// +kcc:proto:field=google.ads.admanager.v1.PrivateAuctionDeal.auction_priority_enabled
	AuctionPriorityEnabled *bool `json:"auctionPriorityEnabled,omitempty"`

	// Optional. Whether the deal has block override enabled.
	// +kcc:proto:field=google.ads.admanager.v1.PrivateAuctionDeal.block_override_enabled
	BlockOverrideEnabled *bool `json:"blockOverrideEnabled,omitempty"`

	// Optional. Whether the deal is exempt from publisher floor price.
	// +kcc:proto:field=google.ads.admanager.v1.PrivateAuctionDeal.publisher_floor_exempt
	PublisherFloorExempt *bool `json:"publisherFloorExempt,omitempty"`

	// Optional. The buyer permission model defining how the deal would transact
	//  among all buyers under the same bidder.
	// +kcc:proto:field=google.ads.admanager.v1.PrivateAuctionDeal.buyer_permission_type
	BuyerPermissionType *string `json:"buyerPermissionType,omitempty"`

	// Optional. The buyer data required by the Marketplace API.
	// +kcc:proto:field=google.ads.admanager.v1.PrivateAuctionDeal.buyer_data
	BuyerData *PrivateAuctionDeal_BuyerData `json:"buyerData,omitempty"`

	// Optional. The priority of the deal across all non-guaranteed deals.
	// +kcc:proto:field=google.ads.admanager.v1.PrivateAuctionDeal.deal_priority
	DealPriority *NonGuaranteedDealPriority `json:"dealPriority,omitempty"`
}

// AdmanagerPrivateAuctionDealStatus defines the config connector machine state of AdmanagerPrivateAuctionDeal
type AdmanagerPrivateAuctionDealStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AdmanagerPrivateAuctionDeal resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *AdmanagerPrivateAuctionDealObservedState `json:"observedState,omitempty"`
}

// AdmanagerPrivateAuctionDealObservedState is the state of the AdmanagerPrivateAuctionDeal resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.ads.admanager.v1.PrivateAuctionDeal
type AdmanagerPrivateAuctionDealObservedState struct {
	// Output only. Deprecated: `PrivateAuctionDeal` ID.
	// +kcc:proto:field=google.ads.admanager.v1.PrivateAuctionDeal.private_auction_deal_id
	PrivateAuctionDealID *int64 `json:"privateAuctionDealID,omitempty"`

	// Output only. The display name of the
	//  [PrivateAuction][google.ads.admanager.v1.PrivateAuction].
	// +kcc:proto:field=google.ads.admanager.v1.PrivateAuctionDeal.private_auction_display_name
	PrivateAuctionDisplayName *string `json:"privateAuctionDisplayName,omitempty"`

	// Output only. The external ID of the `PrivateAuctionDeal`.
	// +kcc:proto:field=google.ads.admanager.v1.PrivateAuctionDeal.external_deal_id
	ExternalDealID *int64 `json:"externalDealID,omitempty"`

	// Optional. The targeting of the `PrivateAuctionDeal`.
	// +kcc:proto:field=google.ads.admanager.v1.PrivateAuctionDeal.targeting
	Targeting *TargetingObservedState `json:"targeting,omitempty"`

	// Output only. The status of the `PrivateAuctionDeal`.
	// +kcc:proto:field=google.ads.admanager.v1.PrivateAuctionDeal.status
	Status *string `json:"status,omitempty"`

	// Output only. Whether the deal is archived.
	// +kcc:proto:field=google.ads.admanager.v1.PrivateAuctionDeal.archived
	Archived *bool `json:"archived,omitempty"`

	// Output only. The instant at which the `PrivateAuctionDeal` was created.
	// +kcc:proto:field=google.ads.admanager.v1.PrivateAuctionDeal.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The instant at which the `PrivateAuctionDeal` was last
	//  updated.
	// +kcc:proto:field=google.ads.admanager.v1.PrivateAuctionDeal.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpadmanagerprivateauctiondeal;gcpadmanagerprivateauctiondeals
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AdmanagerPrivateAuctionDeal is the Schema for the AdmanagerPrivateAuctionDeal API
// +k8s:openapi-gen=true
type AdmanagerPrivateAuctionDeal struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AdmanagerPrivateAuctionDealSpec   `json:"spec,omitempty"`
	Status AdmanagerPrivateAuctionDealStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AdmanagerPrivateAuctionDealList contains a list of AdmanagerPrivateAuctionDeal
type AdmanagerPrivateAuctionDealList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AdmanagerPrivateAuctionDeal `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AdmanagerPrivateAuctionDeal{}, &AdmanagerPrivateAuctionDealList{})
}
