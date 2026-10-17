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

// API sources for AdmanagerEntitySignalsMapping, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/ads/admanager/v1/entity_signals_mapping_messages.proto
// +kcc:source:service-docs=https://developers.google.com/ad-manager/api/beta
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/ad-manager/docs/reference/rest/v1/networks.entitySignalsMappings

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AdmanagerEntitySignalsMappingGVK = GroupVersion.WithKind("AdmanagerEntitySignalsMapping")

// AdmanagerEntitySignalsMappingSpec defines the desired state of AdmanagerEntitySignalsMapping
// +kcc:spec:proto=google.ads.admanager.v1.EntitySignalsMapping
// +kcc:required-from-proto
type AdmanagerEntitySignalsMappingSpec struct {

	// The AdmanagerEntitySignalsMapping name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// ID of an AudienceSegment that this mapping belongs to.
	// +kcc:proto:field=google.ads.admanager.v1.EntitySignalsMapping.audience_segment_id
	AudienceSegmentID *int64 `json:"audienceSegmentID,omitempty"`

	// ID of a ContentBundle that this mapping belongs to.
	// +kcc:proto:field=google.ads.admanager.v1.EntitySignalsMapping.content_bundle_id
	ContentBundleID *int64 `json:"contentBundleID,omitempty"`

	// ID of a CustomValue that this mapping belongs to.
	// +kcc:proto:field=google.ads.admanager.v1.EntitySignalsMapping.custom_targeting_value_id
	CustomTargetingValueID *int64 `json:"customTargetingValueID,omitempty"`

	// Optional. The IDs of the categories that are associated with the
	//  referencing entity.
	// +kcc:proto:field=google.ads.admanager.v1.EntitySignalsMapping.taxonomy_category_ids
	TaxonomyCategoryIDs []int64 `json:"taxonomyCategoryIDs,omitempty"`
}

// AdmanagerEntitySignalsMappingStatus defines the config connector machine state of AdmanagerEntitySignalsMapping
type AdmanagerEntitySignalsMappingStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AdmanagerEntitySignalsMapping resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *AdmanagerEntitySignalsMappingObservedState `json:"observedState,omitempty"`
}

// AdmanagerEntitySignalsMappingObservedState is the state of the AdmanagerEntitySignalsMapping resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.ads.admanager.v1.EntitySignalsMapping
type AdmanagerEntitySignalsMappingObservedState struct {
	// Output only. Deprecated: `EntitySignalsMapping` ID.
	// +kcc:proto:field=google.ads.admanager.v1.EntitySignalsMapping.entity_signals_mapping_id
	EntitySignalsMappingID *int64 `json:"entitySignalsMappingID,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpadmanagerentitysignalsmapping;gcpadmanagerentitysignalsmappings
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AdmanagerEntitySignalsMapping is the Schema for the AdmanagerEntitySignalsMapping API
// +k8s:openapi-gen=true
type AdmanagerEntitySignalsMapping struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AdmanagerEntitySignalsMappingSpec   `json:"spec,omitempty"`
	Status AdmanagerEntitySignalsMappingStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AdmanagerEntitySignalsMappingList contains a list of AdmanagerEntitySignalsMapping
type AdmanagerEntitySignalsMappingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AdmanagerEntitySignalsMapping `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AdmanagerEntitySignalsMapping{}, &AdmanagerEntitySignalsMappingList{})
}
