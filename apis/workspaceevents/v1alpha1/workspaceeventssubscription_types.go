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

// API sources for WorkspaceeventsSubscription, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/apps/events/subscriptions/v1/subscription_resource.proto
// +kcc:source:service-docs=https://developers.google.com/workspace/events
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/workspace/events/reference/rest/v1/subscriptions

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var WorkspaceeventsSubscriptionGVK = GroupVersion.WithKind("WorkspaceeventsSubscription")

// WorkspaceeventsSubscriptionSpec defines the desired state of WorkspaceeventsSubscription
// +kcc:spec:proto=google.apps.events.subscriptions.v1.Subscription
// +kcc:required-from-proto
type WorkspaceeventsSubscriptionSpec struct {

	// The WorkspaceeventsSubscription name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Non-empty default. The timestamp in UTC when the subscription expires.
	//  Always displayed on output, regardless of what was used on input.
	// +kcc:proto:field=google.apps.events.subscriptions.v1.Subscription.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`

	// Input only. The time-to-live (TTL) or duration for the subscription. If
	//  unspecified or set to `0`, uses the maximum possible duration.
	// +kcc:proto:field=google.apps.events.subscriptions.v1.Subscription.ttl
	TTL *string `json:"ttl,omitempty"`

	// Required. Immutable. The Google Workspace resource that's monitored for
	//  events, formatted as the [full resource
	//  name](https://google.aip.dev/122#full-resource-names). To learn about
	//  target resources and the events that they support, see [Supported Google
	//  Workspace
	//  events](https://developers.google.com/workspace/events#supported-events).
	//
	//  A user can only authorize your app to create one subscription for a given
	//  target resource. If your app tries to create another subscription with the
	//  same user credentials, the request returns an `ALREADY_EXISTS` error.
	// +kcc:proto:field=google.apps.events.subscriptions.v1.Subscription.target_resource
	// +required
	TargetResource *string `json:"targetResource,omitempty"`

	// Required. Immutable. Unordered list. Input for creating a subscription.
	//  Otherwise, output only. One or more types of events to receive about the
	//  target resource. Formatted according to the CloudEvents specification.
	//
	//  The supported event types depend on the target resource of your
	//  subscription. For details, see [Supported Google Workspace
	//  events](https://developers.google.com/workspace/events/guides#supported-events).
	//
	//  By default, you also receive events about the [lifecycle of your
	//  subscription](https://developers.google.com/workspace/events/guides/events-lifecycle).
	//  You don't need to specify lifecycle events for this field.
	//
	//  If you specify an event type that doesn't exist for the target resource,
	//  the request returns an HTTP `400 Bad Request` status code.
	// +kcc:proto:field=google.apps.events.subscriptions.v1.Subscription.event_types
	// +required
	EventTypes []string `json:"eventTypes,omitempty"`

	// Optional. Options about what data to include in the event payload. Only
	//  supported for Google Chat events.
	// +kcc:proto:field=google.apps.events.subscriptions.v1.Subscription.payload_options
	PayloadOptions *PayloadOptions `json:"payloadOptions,omitempty"`

	// Required. Immutable. The endpoint where the subscription delivers events,
	//  such as a Pub/Sub topic.
	// +kcc:proto:field=google.apps.events.subscriptions.v1.Subscription.notification_endpoint
	// +required
	NotificationEndpoint *NotificationEndpoint `json:"notificationEndpoint,omitempty"`

	// Optional. This checksum is computed by the server based on the value of
	//  other fields, and might be sent on update requests to ensure the client has
	//  an up-to-date value before proceeding.
	// +kcc:proto:field=google.apps.events.subscriptions.v1.Subscription.etag
	Etag *string `json:"etag,omitempty"`
}

// WorkspaceeventsSubscriptionStatus defines the config connector machine state of WorkspaceeventsSubscription
type WorkspaceeventsSubscriptionStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the WorkspaceeventsSubscription resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *WorkspaceeventsSubscriptionObservedState `json:"observedState,omitempty"`
}

// WorkspaceeventsSubscriptionObservedState is the state of the WorkspaceeventsSubscription resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.apps.events.subscriptions.v1.Subscription
type WorkspaceeventsSubscriptionObservedState struct {
	// Output only. System-assigned unique identifier for the subscription.
	// +kcc:proto:field=google.apps.events.subscriptions.v1.Subscription.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The state of the subscription. Determines whether the
	//  subscription can receive events and deliver them to the notification
	//  endpoint.
	// +kcc:proto:field=google.apps.events.subscriptions.v1.Subscription.state
	State *string `json:"state,omitempty"`

	// Output only. The error that suspended the subscription.
	//
	//  To reactivate the subscription, resolve the error and call the
	//  [`ReactivateSubscription`][google.apps.events.subscriptions.v1.SubscriptionsService.ReactivateSubscription]
	//  method.
	// +kcc:proto:field=google.apps.events.subscriptions.v1.Subscription.suspension_reason
	SuspensionReason *string `json:"suspensionReason,omitempty"`

	// Output only. The user who authorized the creation of the
	//  subscription.
	//
	//  Format: `users/{user}`
	//
	//  For Google Workspace users, the `{user}` value is the
	//  [`user.id`](https://developers.google.com/admin-sdk/directory/reference/rest/v1/users#User.FIELDS.ids)
	//  field from the Directory API.
	// +kcc:proto:field=google.apps.events.subscriptions.v1.Subscription.authority
	Authority *string `json:"authority,omitempty"`

	// Output only. The time when the subscription is created.
	// +kcc:proto:field=google.apps.events.subscriptions.v1.Subscription.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The last time that the subscription is updated.
	// +kcc:proto:field=google.apps.events.subscriptions.v1.Subscription.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. If `true`, the subscription is in the process of being
	//  updated.
	// +kcc:proto:field=google.apps.events.subscriptions.v1.Subscription.reconciling
	Reconciling *bool `json:"reconciling,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpworkspaceeventssubscription;gcpworkspaceeventssubscriptions
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// WorkspaceeventsSubscription is the Schema for the WorkspaceeventsSubscription API
// +k8s:openapi-gen=true
type WorkspaceeventsSubscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   WorkspaceeventsSubscriptionSpec   `json:"spec,omitempty"`
	Status WorkspaceeventsSubscriptionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// WorkspaceeventsSubscriptionList contains a list of WorkspaceeventsSubscription
type WorkspaceeventsSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WorkspaceeventsSubscription `json:"items"`
}

func init() {
	SchemeBuilder.Register(&WorkspaceeventsSubscription{}, &WorkspaceeventsSubscriptionList{})
}
