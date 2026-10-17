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

// API sources for AdmanagerApplication, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/ads/admanager/v1/application_messages.proto
// +kcc:source:service-docs=https://developers.google.com/ad-manager/api/beta
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/ad-manager/docs/reference/rest/v1/networks.applications

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AdmanagerApplicationGVK = GroupVersion.WithKind("AdmanagerApplication")

// AdmanagerApplicationSpec defines the desired state of AdmanagerApplication
// +kcc:spec:proto=google.ads.admanager.v1.Application
// +kcc:required-from-proto
type AdmanagerApplicationSpec struct {

	// The AdmanagerApplication name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. The display name of the application.
	// +kcc:proto:field=google.ads.admanager.v1.Application.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. The app store ID of the app to claim.
	// +kcc:proto:field=google.ads.admanager.v1.Application.app_store_id
	AppStoreID *string `json:"appStoreID,omitempty"`

	// Optional. The app stores the application belongs to. This attribute is
	//  mutable to allow for third party app store linking.
	// +kcc:proto:field=google.ads.admanager.v1.Application.app_stores
	AppStores []string `json:"appStores,omitempty"`
}

// AdmanagerApplicationStatus defines the config connector machine state of AdmanagerApplication
type AdmanagerApplicationStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AdmanagerApplication resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *AdmanagerApplicationObservedState `json:"observedState,omitempty"`
}

// AdmanagerApplicationObservedState is the state of the AdmanagerApplication resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.ads.admanager.v1.Application
type AdmanagerApplicationObservedState struct {
	// Output only. The archival status of the application.
	//
	//  When true, an application cannot be targeted and will not serve ads,
	//  regardless of its `status`.
	// +kcc:proto:field=google.ads.admanager.v1.Application.archived
	Archived *bool `json:"archived,omitempty"`

	// Output only. The name of the application on the app store.
	// +kcc:proto:field=google.ads.admanager.v1.Application.app_store_display_name
	AppStoreDisplayName *string `json:"appStoreDisplayName,omitempty"`

	// Output only. The application code used to identify the app in the SDK.
	//
	//  Note that the UI refers to this as "App ID".
	// +kcc:proto:field=google.ads.admanager.v1.Application.application_code
	ApplicationCode *string `json:"applicationCode,omitempty"`

	// Output only. The name of the developer of the application.
	// +kcc:proto:field=google.ads.admanager.v1.Application.developer
	Developer *string `json:"developer,omitempty"`

	// Output only. The platform the application runs on.
	// +kcc:proto:field=google.ads.admanager.v1.Application.platform
	Platform *string `json:"platform,omitempty"`

	// Output only. Whether the application is free on the app store it belongs
	//  to.
	// +kcc:proto:field=google.ads.admanager.v1.Application.free
	Free *bool `json:"free,omitempty"`

	// Output only. The download URL of the application on the app store it
	//  belongs to.
	// +kcc:proto:field=google.ads.admanager.v1.Application.download_url
	DownloadURL *string `json:"downloadURL,omitempty"`

	// Output only. The approval status for the application.
	// +kcc:proto:field=google.ads.admanager.v1.Application.approval_status
	ApprovalStatus *string `json:"approvalStatus,omitempty"`

	// Output only. The webview claiming status for the application.
	// +kcc:proto:field=google.ads.admanager.v1.Application.webview_claiming_status
	WebviewClaimingStatus *string `json:"webviewClaimingStatus,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpadmanagerapplication;gcpadmanagerapplications
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AdmanagerApplication is the Schema for the AdmanagerApplication API
// +k8s:openapi-gen=true
type AdmanagerApplication struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AdmanagerApplicationSpec   `json:"spec,omitempty"`
	Status AdmanagerApplicationStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AdmanagerApplicationList contains a list of AdmanagerApplication
type AdmanagerApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AdmanagerApplication `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AdmanagerApplication{}, &AdmanagerApplicationList{})
}
