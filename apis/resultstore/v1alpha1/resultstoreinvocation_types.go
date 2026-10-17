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

// API sources for ResultstoreInvocation, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/devtools/resultstore/v2/invocation.proto
// +kcc:guess=source-link reason=verify-service-docs-link
// +kcc:source:service-docs=
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ResultstoreInvocationGVK = GroupVersion.WithKind("ResultstoreInvocation")

// ResultstoreInvocationSpec defines the desired state of ResultstoreInvocation
// +kcc:spec:proto=google.devtools.resultstore.v2.Invocation
// +kcc:required-from-proto
type ResultstoreInvocationSpec struct {

	// The ResultstoreInvocation name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// The aggregate status of the invocation.
	// +kcc:proto:field=google.devtools.resultstore.v2.Invocation.status_attributes
	StatusAttributes *StatusAttributes `json:"statusAttributes,omitempty"`

	// When this invocation started and its duration.
	// +kcc:proto:field=google.devtools.resultstore.v2.Invocation.timing
	Timing *Timing `json:"timing,omitempty"`

	// Attributes of this invocation.
	// +kcc:proto:field=google.devtools.resultstore.v2.Invocation.invocation_attributes
	InvocationAttributes *InvocationAttributes `json:"invocationAttributes,omitempty"`

	// The workspace the tool was run in.
	// +kcc:proto:field=google.devtools.resultstore.v2.Invocation.workspace_info
	WorkspaceInfo *WorkspaceInfo `json:"workspaceInfo,omitempty"`

	// Arbitrary name-value pairs.
	//  This is implemented as a multi-map. Multiple properties are allowed with
	//  the same key. Properties will be returned in lexicographical order by key.
	// +kcc:proto:field=google.devtools.resultstore.v2.Invocation.properties
	Properties []Property `json:"properties,omitempty"`

	// A list of file references for invocation level files.
	//  The file IDs must be unique within this list. Duplicate file IDs will
	//  result in an error. Files will be returned in lexicographical order by ID.
	//  Use this field to specify build logs, and other invocation level logs.
	//
	//  Files with the following reserved file IDs cause specific post-processing
	//  or have special handling. These files must be immediately available to
	//  ResultStore for processing when the reference is uploaded.
	//
	//  build.log: The primary log for the Invocation.
	//  coverage_report.lcov: Aggregate coverage report for the invocation.
	// +kcc:proto:field=google.devtools.resultstore.v2.Invocation.files
	Files []File `json:"files,omitempty"`

	// Summary of aggregate coverage across all Actions in this Invocation.
	//  If missing, this data will be populated by the server from the
	//  coverage_report.lcov file or the union of all ActionCoverages under this
	//  invocation (in that order).
	// +kcc:proto:field=google.devtools.resultstore.v2.Invocation.coverage_summaries
	CoverageSummaries []LanguageCoverageSummary `json:"coverageSummaries,omitempty"`

	// Aggregate code coverage for all build and test Actions within this
	//  Invocation. If missing, this data will be populated by the server
	//  from the coverage_report.lcov file or the union of all ActionCoverages
	//  under this invocation (in that order).
	// +kcc:proto:field=google.devtools.resultstore.v2.Invocation.aggregate_coverage
	AggregateCoverage *AggregateCoverage `json:"aggregateCoverage,omitempty"`

	// NOT IMPLEMENTED.
	//  ResultStore will read and parse Files with reserved IDs listed above. Read
	//  and parse errors for all these Files are reported here.
	//  This is implemented as a map, with one FileProcessingErrors for each file.
	//  Typically produced when parsing Files, but may also be provided directly
	//  by clients.
	// +kcc:proto:field=google.devtools.resultstore.v2.Invocation.file_processing_errors
	FileProcessingErrors []FileProcessingErrors `json:"fileProcessingErrors,omitempty"`
}

// ResultstoreInvocationStatus defines the config connector machine state of ResultstoreInvocation
type ResultstoreInvocationStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ResultstoreInvocation resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ResultstoreInvocationObservedState `json:"observedState,omitempty"`
}

// ResultstoreInvocationObservedState is the state of the ResultstoreInvocation resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.devtools.resultstore.v2.Invocation
type ResultstoreInvocationObservedState struct {
	// The resource ID components that identify the Invocation. They must match
	//  the resource name after proper encoding.
	// +kcc:guess=placement reason=no-field-behavior-on-message
	// +kcc:proto:field=google.devtools.resultstore.v2.Invocation.id
	ID *Invocation_ID `json:"id,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpresultstoreinvocation;gcpresultstoreinvocations
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ResultstoreInvocation is the Schema for the ResultstoreInvocation API
// +k8s:openapi-gen=true
type ResultstoreInvocation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ResultstoreInvocationSpec   `json:"spec,omitempty"`
	Status ResultstoreInvocationStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ResultstoreInvocationList contains a list of ResultstoreInvocation
type ResultstoreInvocationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResultstoreInvocation `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ResultstoreInvocation{}, &ResultstoreInvocationList{})
}
