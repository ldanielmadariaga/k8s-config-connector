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

// API sources for ChatMembership, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/chat/v1/membership.proto
// +kcc:source:service-docs=https://developers.google.com/workspace/chat
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/chat/api/reference/rest/v1/spaces.members

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ChatMembershipGVK = GroupVersion.WithKind("ChatMembership")

// ChatMembershipSpec defines the desired state of ChatMembership
// +kcc:spec:proto=google.chat.v1.Membership
// +kcc:required-from-proto
type ChatMembershipSpec struct {

	// The ChatMembership name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. User's role within a Chat space, which determines their permitted
	//  actions in the space.
	//
	//  This field can only be used as input in `UpdateMembership`.
	// +kcc:proto:field=google.chat.v1.Membership.role
	Role *string `json:"role,omitempty"`

	// Optional. The Google Chat user or app the membership corresponds to.
	//  If your Chat app [authenticates as a
	//  user](https://developers.google.com/workspace/chat/authenticate-authorize-chat-user),
	//  the output populates the
	//  [user](https://developers.google.com/workspace/chat/api/reference/rest/v1/User)
	//  `name` and `type`.
	// +kcc:proto:field=google.chat.v1.Membership.member
	Member *User `json:"member,omitempty"`

	// Optional. The Google Group the membership corresponds to.
	//
	//  Reading or mutating memberships for Google Groups requires [user
	//  authentication](https://developers.google.com/workspace/chat/authenticate-authorize-chat-user).
	// +kcc:proto:field=google.chat.v1.Membership.group_member
	GroupMember *Group `json:"groupMember,omitempty"`

	// Optional. Immutable. The creation time of the membership, such as when a
	//  member joined or was invited to join a space. This field is output only,
	//  except when used to import historical memberships in import mode spaces.
	// +kcc:proto:field=google.chat.v1.Membership.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Optional. Immutable. The deletion time of the membership, such as when a
	//  member left or was removed from a space. This field is output only, except
	//  when used to import historical memberships in import mode spaces.
	// +kcc:proto:field=google.chat.v1.Membership.delete_time
	DeleteTime *string `json:"deleteTime,omitempty"`
}

// ChatMembershipStatus defines the config connector machine state of ChatMembership
type ChatMembershipStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ChatMembership resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ChatMembershipObservedState `json:"observedState,omitempty"`
}

// ChatMembershipObservedState is the state of the ChatMembership resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.chat.v1.Membership
type ChatMembershipObservedState struct {
	// Output only. State of the membership.
	// +kcc:proto:field=google.chat.v1.Membership.state
	State *string `json:"state,omitempty"`

	// Optional. The Google Chat user or app the membership corresponds to.
	//  If your Chat app [authenticates as a
	//  user](https://developers.google.com/workspace/chat/authenticate-authorize-chat-user),
	//  the output populates the
	//  [user](https://developers.google.com/workspace/chat/api/reference/rest/v1/User)
	//  `name` and `type`.
	// +kcc:proto:field=google.chat.v1.Membership.member
	Member *UserObservedState `json:"member,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpchatmembership;gcpchatmemberships
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ChatMembership is the Schema for the ChatMembership API
// +k8s:openapi-gen=true
type ChatMembership struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ChatMembershipSpec   `json:"spec,omitempty"`
	Status ChatMembershipStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ChatMembershipList contains a list of ChatMembership
type ChatMembershipList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ChatMembership `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ChatMembership{}, &ChatMembershipList{})
}
