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

// API sources for AdmanagerReport, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/ads/admanager/v1/report_messages.proto
// +kcc:source:service-docs=https://developers.google.com/ad-manager/api/beta
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/ad-manager/docs/reference/rest/v1/networks.reports

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AdmanagerReportGVK = GroupVersion.WithKind("AdmanagerReport")

// AdmanagerReportSpec defines the desired state of AdmanagerReport
// +kcc:spec:proto=google.ads.admanager.v1.Report
// +kcc:required-from-proto
type AdmanagerReportSpec struct {

	// The AdmanagerReport name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. The visibility of a report.
	// +kcc:proto:field=google.ads.admanager.v1.Report.visibility
	Visibility *string `json:"visibility,omitempty"`

	// Required. The report definition of the report.
	// +kcc:proto:field=google.ads.admanager.v1.Report.report_definition
	// +required
	ReportDefinition *ReportDefinition `json:"reportDefinition,omitempty"`

	// Optional. The draft report definition of the report.
	// +kcc:proto:field=google.ads.admanager.v1.Report.draft_report_definition
	DraftReportDefinition *ReportDefinition `json:"draftReportDefinition,omitempty"`

	// Optional. Display name for the report.
	// +kcc:proto:field=google.ads.admanager.v1.Report.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. The schedule options of this report.
	// +kcc:proto:field=google.ads.admanager.v1.Report.schedule_options
	ScheduleOptions *ScheduleOptions `json:"scheduleOptions,omitempty"`
}

// AdmanagerReportStatus defines the config connector machine state of AdmanagerReport
type AdmanagerReportStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AdmanagerReport resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *AdmanagerReportObservedState `json:"observedState,omitempty"`
}

// AdmanagerReportObservedState is the state of the AdmanagerReport resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.ads.admanager.v1.Report
type AdmanagerReportObservedState struct {
	// Output only. Report ID.
	// +kcc:proto:field=google.ads.admanager.v1.Report.report_id
	ReportID *int64 `json:"reportID,omitempty"`

	// Output only. The instant this report was last modified.
	// +kcc:proto:field=google.ads.admanager.v1.Report.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The instant this report was created.
	// +kcc:proto:field=google.ads.admanager.v1.Report.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The locale of this report. Locale is set from the user's
	//  locale at the time of the request. Locale can't be modified.
	// +kcc:proto:field=google.ads.admanager.v1.Report.locale
	Locale *string `json:"locale,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpadmanagerreport;gcpadmanagerreports
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AdmanagerReport is the Schema for the AdmanagerReport API
// +k8s:openapi-gen=true
type AdmanagerReport struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AdmanagerReportSpec   `json:"spec,omitempty"`
	Status AdmanagerReportStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AdmanagerReportList contains a list of AdmanagerReport
type AdmanagerReportList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AdmanagerReport `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AdmanagerReport{}, &AdmanagerReportList{})
}
