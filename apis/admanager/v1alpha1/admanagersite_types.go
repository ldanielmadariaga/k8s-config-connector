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

// API sources for AdmanagerSite, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/ads/admanager/v1/site_messages.proto
// +kcc:source:service-docs=https://developers.google.com/ad-manager/api/beta
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/ad-manager/docs/reference/rest/v1/networks.sites

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AdmanagerSiteGVK = GroupVersion.WithKind("AdmanagerSite")

// AdmanagerSiteSpec defines the desired state of AdmanagerSite
// +kcc:spec:proto=google.ads.admanager.v1.Site
// +kcc:required-from-proto
type AdmanagerSiteSpec struct {

	// The AdmanagerSite name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. The URL of the Site.
	// +kcc:proto:field=google.ads.admanager.v1.Site.url
	// +required
	URL *string `json:"url,omitempty"`

	// Optional. The network code of the child if the Site is being managed for an
	//  MCM child network, or null if owned by this network.
	// +kcc:proto:field=google.ads.admanager.v1.Site.child_network_code
	ChildNetworkCode *string `json:"childNetworkCode,omitempty"`
}

// AdmanagerSiteStatus defines the config connector machine state of AdmanagerSite
type AdmanagerSiteStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AdmanagerSite resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *AdmanagerSiteObservedState `json:"observedState,omitempty"`
}

// AdmanagerSiteObservedState is the state of the AdmanagerSite resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.ads.admanager.v1.Site
type AdmanagerSiteObservedState struct {
	// Output only. Status of the review performed on the Site by Google.
	// +kcc:proto:field=google.ads.admanager.v1.Site.approval_status
	ApprovalStatus *string `json:"approvalStatus,omitempty"`

	// Output only. The latest Site approval status change time.
	// +kcc:proto:field=google.ads.admanager.v1.Site.approval_status_update_time
	ApprovalStatusUpdateTime *string `json:"approvalStatusUpdateTime,omitempty"`

	// Output only. Provides reasons for disapproving the Site. It is null when
	//  the Site is not disapproved.
	// +kcc:proto:field=google.ads.admanager.v1.Site.disapproval_reasons
	DisapprovalReasons []DisapprovalReasonObservedState `json:"disapprovalReasons,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpadmanagersite;gcpadmanagersites
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AdmanagerSite is the Schema for the AdmanagerSite API
// +k8s:openapi-gen=true
type AdmanagerSite struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AdmanagerSiteSpec   `json:"spec,omitempty"`
	Status AdmanagerSiteStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AdmanagerSiteList contains a list of AdmanagerSite
type AdmanagerSiteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AdmanagerSite `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AdmanagerSite{}, &AdmanagerSiteList{})
}
