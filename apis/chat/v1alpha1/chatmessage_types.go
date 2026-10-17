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

// API sources for ChatMessage, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/chat/v1/message.proto
// +kcc:source:service-docs=https://developers.google.com/workspace/chat
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/chat/api/reference/rest/v1/spaces.messages

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ChatMessageGVK = GroupVersion.WithKind("ChatMessage")

// ChatMessageSpec defines the desired state of ChatMessage
// +kcc:spec:proto=google.chat.v1.Message
// +kcc:required-from-proto
type ChatMessageSpec struct {

	// The ChatMessage name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. Immutable. For spaces created in Chat, the time at which the
	//  message was created. This field is output only, except when used in import
	//  mode spaces.
	//
	//  For import mode spaces, set this field to the historical timestamp at which
	//  the message was created in the source in order to preserve the original
	//  creation time.
	// +kcc:proto:field=google.chat.v1.Message.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Optional. Plain-text body of the message. The first link to an image,
	//  video, or web page generates a [preview
	//  chip](https://developers.google.com/workspace/chat/preview-links). You can
	//  also [@mention a Google Chat
	//  user](https://developers.google.com/workspace/chat/format-messages#messages-@mention),
	//  or everyone in the space.
	//
	//  To learn about creating text messages, see [Send a
	//  message](https://developers.google.com/workspace/chat/create-messages).
	// +kcc:proto:field=google.chat.v1.Message.text
	Text *string `json:"text,omitempty"`

	// Optional. An array of
	//  [cards](https://developers.google.com/workspace/chat/api/reference/rest/v1/cards).
	//
	//  Only Chat apps can create cards. If your Chat app [authenticates as a
	//  user](https://developers.google.com/workspace/chat/authenticate-authorize-chat-user),
	//  the messages can't contain cards.
	//
	//  To learn how to create a message that contains cards, see [Send a
	//  message](https://developers.google.com/workspace/chat/create-messages).
	//
	//  [Card builder](https://addons.gsuite.google.com/uikit/builder)
	// +kcc:proto:field=google.chat.v1.Message.cards_v2
	CardsV2 []CardWithID `json:"cardsV2,omitempty"`

	// The thread the message belongs to. For example usage, see
	//  [Start or reply to a message
	//  thread](https://developers.google.com/workspace/chat/create-messages#create-message-thread).
	// +kcc:proto:field=google.chat.v1.Message.thread
	Thread *Thread `json:"thread,omitempty"`

	// Optional. A plain-text description of the message's cards, used when the
	//  actual cards can't be displayed—for example, mobile notifications.
	// +kcc:proto:field=google.chat.v1.Message.fallback_text
	FallbackText *string `json:"fallbackText,omitempty"`

	// Input only. Parameters that a Chat app can use to configure how its
	//  response is posted.
	// +kcc:proto:field=google.chat.v1.Message.action_response
	ActionResponse *ActionResponse `json:"actionResponse,omitempty"`

	// Optional. User-uploaded attachment.
	// +kcc:proto:field=google.chat.v1.Message.attachment
	Attachment []Attachment `json:"attachment,omitempty"`

	// Optional. A custom ID for the message. You can use field to identify a
	//  message, or to get, delete, or update a message. To set a custom ID,
	//  specify the
	//  [`messageId`](https://developers.google.com/workspace/chat/api/reference/rest/v1/spaces.messages/create#body.QUERY_PARAMETERS.message_id)
	//  field when you create the message. For details, see [Name a
	//  message](https://developers.google.com/workspace/chat/create-messages#name_a_created_message).
	// +kcc:proto:field=google.chat.v1.Message.client_assigned_message_id
	ClientAssignedMessageID *string `json:"clientAssignedMessageID,omitempty"`

	// Optional. Immutable. Input for creating a message, otherwise output only.
	//  The user that can view the message. When set, the message is private and
	//  only visible to the specified user and the Chat app. To include this field
	//  in your request, you must call the Chat API using [app
	//  authentication](https://developers.google.com/workspace/chat/authenticate-authorize-chat-app)
	//  and omit the following:
	//
	//  * [Attachments](https://developers.google.com/workspace/chat/api/reference/rest/v1/spaces.messages.attachments)
	//  * [Accessory
	//  widgets](https://developers.google.com/workspace/chat/api/reference/rest/v1/spaces.messages#Message.AccessoryWidget)
	//
	//  For details, see [Send a message
	//  privately](https://developers.google.com/workspace/chat/create-messages#private).
	// +kcc:proto:field=google.chat.v1.Message.private_message_viewer
	PrivateMessageViewer *User `json:"privateMessageViewer,omitempty"`

	// Optional. One or more interactive widgets that appear at the bottom of a
	//  message. You can add accessory widgets to messages that contain text,
	//  cards, or both text and cards. Not supported for messages that contain
	//  dialogs. For details, see [Add interactive widgets at the bottom of a
	//  message](https://developers.google.com/workspace/chat/create-messages#add-accessory-widgets).
	//
	//  Creating a message with accessory widgets requires [app
	//  authentication]
	//  (https://developers.google.com/workspace/chat/authenticate-authorize-chat-app).
	// +kcc:proto:field=google.chat.v1.Message.accessory_widgets
	AccessoryWidgets []AccessoryWidget `json:"accessoryWidgets,omitempty"`
}

// ChatMessageStatus defines the config connector machine state of ChatMessage
type ChatMessageStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ChatMessage resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ChatMessageObservedState `json:"observedState,omitempty"`
}

// ChatMessageObservedState is the state of the ChatMessage resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.chat.v1.Message
type ChatMessageObservedState struct {
	// Output only. The user who created the message.
	//  If your Chat app [authenticates as a
	//  user](https://developers.google.com/workspace/chat/authenticate-authorize-chat-user),
	//  the output populates the
	//  [user](https://developers.google.com/workspace/chat/api/reference/rest/v1/User)
	//  `name` and `type`.
	// +kcc:proto:field=google.chat.v1.Message.sender
	Sender *UserObservedState `json:"sender,omitempty"`

	// Output only. The time at which the message was last edited by a user. If
	//  the message has never been edited, this field is empty.
	// +kcc:proto:field=google.chat.v1.Message.last_update_time
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`

	// Output only. The time at which the message was deleted in
	//  Google Chat. If the message is never deleted, this field is empty.
	// +kcc:proto:field=google.chat.v1.Message.delete_time
	DeleteTime *string `json:"deleteTime,omitempty"`

	// Output only. Contains the message `text` with markups added to communicate
	//  formatting. This field might not capture all formatting visible in the UI,
	//  but includes the following:
	//
	//  * [Markup
	//  syntax](https://developers.google.com/workspace/chat/format-messages)
	//  for bold, italic, strikethrough, monospace, monospace block, and bulleted
	//  list.
	//
	//  * [User
	//  mentions](https://developers.google.com/workspace/chat/format-messages#messages-@mention)
	//  using the format `<users/{user}>`.
	//
	//  * Custom hyperlinks using the format `<{url}|{rendered_text}>` where the
	//  first string is the URL and the second is the rendered text—for example,
	//  `<http://example.com|custom text>`.
	//
	//  * Custom emoji using the format `:{emoji_name}:`—for example, `:smile:`.
	//  This doesn't apply to Unicode emoji, such as `U+1F600` for a grinning
	//  face emoji.
	//
	//  For more information, see [View text formatting sent in a
	//  message](https://developers.google.com/workspace/chat/format-messages#view_text_formatting_sent_in_a_message)
	// +kcc:proto:field=google.chat.v1.Message.formatted_text
	FormattedText *string `json:"formattedText,omitempty"`

	// Output only. Annotations can be associated with the plain-text body of the
	//  message or with chips that link to Google Workspace resources like Google
	//  Docs or Sheets with `start_index` and `length` of 0.
	// +kcc:proto:field=google.chat.v1.Message.annotations
	Annotations []AnnotationObservedState `json:"annotations,omitempty"`

	// Output only. If your Chat app [authenticates as a
	//  user](https://developers.google.com/workspace/chat/authenticate-authorize-chat-user),
	//  the output only populates the
	//  [space](https://developers.google.com/workspace/chat/api/reference/rest/v1/spaces)
	//  `name`.
	// +kcc:proto:field=google.chat.v1.Message.space
	Space *SpaceObservedState `json:"space,omitempty"`

	// Output only. Plain-text body of the message with all Chat app mentions
	//  stripped out.
	// +kcc:proto:field=google.chat.v1.Message.argument_text
	ArgumentText *string `json:"argumentText,omitempty"`

	// Output only. Slash command information, if applicable.
	// +kcc:proto:field=google.chat.v1.Message.slash_command
	SlashCommand *SlashCommand `json:"slashCommand,omitempty"`

	// Optional. User-uploaded attachment.
	// +kcc:proto:field=google.chat.v1.Message.attachment
	Attachment []AttachmentObservedState `json:"attachment,omitempty"`

	// Output only. A URL in `spaces.messages.text` that matches a link preview
	//  pattern. For more information, see [Preview
	//  links](https://developers.google.com/workspace/chat/preview-links).
	// +kcc:proto:field=google.chat.v1.Message.matched_url
	MatchedURL *MatchedURLObservedState `json:"matchedURL,omitempty"`

	// Output only. When `true`, the message is a response in a reply thread. When
	//  `false`, the message is visible in the space's top-level conversation as
	//  either the first message of a thread or a message with no threaded replies.
	//
	//  If the space doesn't support reply in threads, this field is always
	//  `false`.
	// +kcc:proto:field=google.chat.v1.Message.thread_reply
	ThreadReply *bool `json:"threadReply,omitempty"`

	// Output only. The list of emoji reaction summaries on the message.
	// +kcc:proto:field=google.chat.v1.Message.emoji_reaction_summaries
	EmojiReactionSummaries []EmojiReactionSummaryObservedState `json:"emojiReactionSummaries,omitempty"`

	// Optional. Immutable. Input for creating a message, otherwise output only.
	//  The user that can view the message. When set, the message is private and
	//  only visible to the specified user and the Chat app. To include this field
	//  in your request, you must call the Chat API using [app
	//  authentication](https://developers.google.com/workspace/chat/authenticate-authorize-chat-app)
	//  and omit the following:
	//
	//  * [Attachments](https://developers.google.com/workspace/chat/api/reference/rest/v1/spaces.messages.attachments)
	//  * [Accessory
	//  widgets](https://developers.google.com/workspace/chat/api/reference/rest/v1/spaces.messages#Message.AccessoryWidget)
	//
	//  For details, see [Send a message
	//  privately](https://developers.google.com/workspace/chat/create-messages#private).
	// +kcc:proto:field=google.chat.v1.Message.private_message_viewer
	PrivateMessageViewer *UserObservedState `json:"privateMessageViewer,omitempty"`

	// Output only. Information about a deleted message. A message is deleted when
	//  `delete_time` is set.
	// +kcc:proto:field=google.chat.v1.Message.deletion_metadata
	DeletionMetadata *DeletionMetadata `json:"deletionMetadata,omitempty"`

	// Output only. Information about a message that's quoted by a Google Chat
	//  user in a space. Google Chat users can quote a message to reply to it.
	// +kcc:proto:field=google.chat.v1.Message.quoted_message_metadata
	QuotedMessageMetadata *QuotedMessageMetadataObservedState `json:"quotedMessageMetadata,omitempty"`

	// Output only. GIF images that are attached to the message.
	// +kcc:proto:field=google.chat.v1.Message.attached_gifs
	AttachedGifs []AttachedGifObservedState `json:"attachedGifs,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpchatmessage;gcpchatmessages
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ChatMessage is the Schema for the ChatMessage API
// +k8s:openapi-gen=true
type ChatMessage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ChatMessageSpec   `json:"spec,omitempty"`
	Status ChatMessageStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ChatMessageList contains a list of ChatMessage
type ChatMessageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ChatMessage `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ChatMessage{}, &ChatMessageList{})
}
