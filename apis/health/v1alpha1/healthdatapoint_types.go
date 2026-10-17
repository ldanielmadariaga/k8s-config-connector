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

// API sources for HealthDataPoint, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/devicesandservices/health/v4/data_points.proto
// +kcc:source:service-docs=https://developers.google.com/health
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/health/docs/reference/rest/v4/users.dataTypes.dataPoints

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// HealthDataPointSpec defines the desired state of HealthDataPoint
// +kcc:spec:proto=google.devicesandservices.health.v4.DataPoint
// +kcc:required-from-proto
type HealthDataPointSpec struct {

	// The HealthDataPoint name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. Data for points in the `steps` interval data type collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.steps
	Steps *Steps `json:"steps,omitempty"`

	// Optional. Data for points in the `floors` interval data type collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.floors
	Floors *Floors `json:"floors,omitempty"`

	// Optional. Data for points in the `heart-rate` sample data type
	//  collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.heart_rate
	HeartRate *HeartRate `json:"heartRate,omitempty"`

	// Optional. Data for points in the `sleep` session data type collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.sleep
	Sleep *Sleep `json:"sleep,omitempty"`

	// Optional. Data for points in the `daily-resting-heart-rate` daily data
	//  type collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.daily_resting_heart_rate
	DailyRestingHeartRate *DailyRestingHeartRate `json:"dailyRestingHeartRate,omitempty"`

	// Optional. Data for points in the `daily-heart-rate-variability` daily
	//  data type collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.daily_heart_rate_variability
	DailyHeartRateVariability *DailyHeartRateVariability `json:"dailyHeartRateVariability,omitempty"`

	// Optional. Data for points in the `exercise` session data type collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.exercise
	Exercise *Exercise `json:"exercise,omitempty"`

	// Optional. Data for points in the `weight` sample data type collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.weight
	Weight *Weight `json:"weight,omitempty"`

	// Optional. Data for points in the `altitude` interval data type
	//  collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.altitude
	Altitude *Altitude `json:"altitude,omitempty"`

	// Optional. Data for points in the `distance` interval data type
	//  collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.distance
	Distance *Distance `json:"distance,omitempty"`

	// Optional. Data for points in the `body-fat` sample data type collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.body_fat
	BodyFat *BodyFat `json:"bodyFat,omitempty"`

	// Optional. Data for points in the `active-zone-minutes` interval data type
	//  collection, measured in minutes.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.active_zone_minutes
	ActiveZoneMinutes *ActiveZoneMinutes `json:"activeZoneMinutes,omitempty"`

	// Optional. Data for points in the `heart-rate-variability` sample data
	//  type collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.heart_rate_variability
	HeartRateVariability *HeartRateVariability `json:"heartRateVariability,omitempty"`

	// Optional. Data for points in the `daily-sleep-temperature-derivations`
	//  daily data type collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.daily_sleep_temperature_derivations
	DailySleepTemperatureDerivations *DailySleepTemperatureDerivations `json:"dailySleepTemperatureDerivations,omitempty"`

	// Optional. Data for points in the `sedentary-period` interval data type
	//  collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.sedentary_period
	SedentaryPeriod *SedentaryPeriod `json:"sedentaryPeriod,omitempty"`

	// Optional. Data for points in the `run-vo2-max` sample data type
	//  collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.run_vo2_max
	RunVo2Max *RunVo2Max `json:"runVo2Max,omitempty"`

	// Optional. Data for points in the `oxygen-saturation` sample data type
	//  collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.oxygen_saturation
	OxygenSaturation *OxygenSaturation `json:"oxygenSaturation,omitempty"`

	// Optional. Data for points in the `daily-oxygen-saturation` daily data
	//  type collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.daily_oxygen_saturation
	DailyOxygenSaturation *DailyOxygenSaturation `json:"dailyOxygenSaturation,omitempty"`

	// Optional. Data for points in the `activity-level` daily data type
	//  collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.activity_level
	ActivityLevel *ActivityLevel `json:"activityLevel,omitempty"`

	// Optional. Data for points in the `vo2-max` sample data type collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.vo2_max
	Vo2Max *Vo2Max `json:"vo2Max,omitempty"`

	// Optional. Data for points in the `daily-vo2-max` daily data type
	//  collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.daily_vo2_max
	DailyVo2Max *DailyVo2Max `json:"dailyVo2Max,omitempty"`

	// Optional. Data for points in the `nutrition-log` session data type
	//  collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.nutrition_log
	NutritionLog *NutritionLog `json:"nutritionLog,omitempty"`

	// Optional. Data for points in the `irregular-rhythm-notification` session
	//  data type collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.irregular_rhythm_notification
	IrregularRhythmNotification *IrregularRhythmNotification `json:"irregularRhythmNotification,omitempty"`

	// Optional. Data for points in the `electrocardiogram` session data type
	//  collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.electrocardiogram
	Electrocardiogram *Electrocardiogram `json:"electrocardiogram,omitempty"`

	// Optional. Data for points in the `daily-heart-rate-zones` daily data type
	//  collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.daily_heart_rate_zones
	DailyHeartRateZones *DailyHeartRateZones `json:"dailyHeartRateZones,omitempty"`

	// Optional. Data for points in the `hydration-log` session data type
	//  collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.hydration_log
	HydrationLog *HydrationLog `json:"hydrationLog,omitempty"`

	// Optional. The food details.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.food
	Food *Food `json:"food,omitempty"`

	// Optional. Data for points in the `time-in-heart-rate-zone` interval data
	//  type collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.time_in_heart_rate_zone
	TimeInHeartRateZone *TimeInHeartRateZone `json:"timeInHeartRateZone,omitempty"`

	// Optional. Data for points in the `active-minutes` interval data type
	//  collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.active_minutes
	ActiveMinutes *ActiveMinutes `json:"activeMinutes,omitempty"`

	// Optional. Data for points in the `respiratory-rate-sleep-summary` sample
	//  data type collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.respiratory_rate_sleep_summary
	RespiratoryRateSleepSummary *RespiratoryRateSleepSummary `json:"respiratoryRateSleepSummary,omitempty"`

	// Optional. Data for points in the `daily-respiratory-rate` daily data type
	//  collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.daily_respiratory_rate
	DailyRespiratoryRate *DailyRespiratoryRate `json:"dailyRespiratoryRate,omitempty"`

	// Optional. Data for points in the `swim-lengths-data` interval data type
	//  collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.swim_lengths_data
	SwimLengthsData *SwimLengthsData `json:"swimLengthsData,omitempty"`

	// Optional. Data for points in the `height` sample data type collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.height
	Height *Height `json:"height,omitempty"`

	// Optional. Data for points in the `basal-energy-burned` interval data type
	//  collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.basal_energy_burned
	BasalEnergyBurned *BasalEnergyBurned `json:"basalEnergyBurned,omitempty"`

	// Optional. Data for points in the `core-body-temperature` sample data type
	//  collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.core_body_temperature
	CoreBodyTemperature *CoreBodyTemperature `json:"coreBodyTemperature,omitempty"`

	// Optional. Data for points in the `active-energy-burned` interval data
	//  type collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.active_energy_burned
	ActiveEnergyBurned *ActiveEnergyBurned `json:"activeEnergyBurned,omitempty"`

	// Optional. The food measurement unit details.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.food_measurement_unit
	FoodMeasurementUnit *FoodMeasurementUnit `json:"foodMeasurementUnit,omitempty"`

	// Optional. Data for points in the `blood-glucose` sample data type
	//  collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.blood_glucose
	BloodGlucose *BloodGlucose `json:"bloodGlucose,omitempty"`

	// Optional. Data for points in the `menstrual-period` interval data type
	//  collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.menstrual_period
	MenstrualPeriod *MenstrualPeriod `json:"menstrualPeriod,omitempty"`

	// Optional. Data for points in the `ovulation-test` sample data type
	//  collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.ovulation_test
	OvulationTest *OvulationTest `json:"ovulationTest,omitempty"`

	// Optional. Data for points in the `symptoms` sample data type
	//  collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.symptoms
	Symptoms *Symptoms `json:"symptoms,omitempty"`

	// Optional. Data for points in the `moods` sample data type collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.moods
	Moods *Moods `json:"moods,omitempty"`

	// Optional. Data source information for the metric
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.data_source
	DataSource *DataSource `json:"dataSource,omitempty"`
}

// HealthDataPointStatus defines the config connector machine state of HealthDataPoint
type HealthDataPointStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the HealthDataPoint resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *HealthDataPointObservedState `json:"observedState,omitempty"`
}

// HealthDataPointObservedState is the state of the HealthDataPoint resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.devicesandservices.health.v4.DataPoint
type HealthDataPointObservedState struct {
	// Optional. Data for points in the `steps` interval data type collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.steps
	Steps *StepsObservedState `json:"steps,omitempty"`

	// Optional. Data for points in the `floors` interval data type collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.floors
	Floors *FloorsObservedState `json:"floors,omitempty"`

	// Optional. Data for points in the `heart-rate` sample data type
	//  collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.heart_rate
	HeartRate *HeartRateObservedState `json:"heartRate,omitempty"`

	// Optional. Data for points in the `sleep` session data type collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.sleep
	Sleep *SleepObservedState `json:"sleep,omitempty"`

	// Optional. Data for points in the `exercise` session data type collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.exercise
	Exercise *ExerciseObservedState `json:"exercise,omitempty"`

	// Optional. Data for points in the `weight` sample data type collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.weight
	Weight *WeightObservedState `json:"weight,omitempty"`

	// Optional. Data for points in the `altitude` interval data type
	//  collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.altitude
	Altitude *AltitudeObservedState `json:"altitude,omitempty"`

	// Optional. Data for points in the `distance` interval data type
	//  collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.distance
	Distance *DistanceObservedState `json:"distance,omitempty"`

	// Optional. Data for points in the `body-fat` sample data type collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.body_fat
	BodyFat *BodyFatObservedState `json:"bodyFat,omitempty"`

	// Optional. Data for points in the `active-zone-minutes` interval data type
	//  collection, measured in minutes.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.active_zone_minutes
	ActiveZoneMinutes *ActiveZoneMinutesObservedState `json:"activeZoneMinutes,omitempty"`

	// Optional. Data for points in the `heart-rate-variability` sample data
	//  type collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.heart_rate_variability
	HeartRateVariability *HeartRateVariabilityObservedState `json:"heartRateVariability,omitempty"`

	// Optional. Data for points in the `sedentary-period` interval data type
	//  collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.sedentary_period
	SedentaryPeriod *SedentaryPeriodObservedState `json:"sedentaryPeriod,omitempty"`

	// Optional. Data for points in the `run-vo2-max` sample data type
	//  collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.run_vo2_max
	RunVo2Max *RunVo2MaxObservedState `json:"runVo2Max,omitempty"`

	// Optional. Data for points in the `oxygen-saturation` sample data type
	//  collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.oxygen_saturation
	OxygenSaturation *OxygenSaturationObservedState `json:"oxygenSaturation,omitempty"`

	// Optional. Data for points in the `activity-level` daily data type
	//  collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.activity_level
	ActivityLevel *ActivityLevelObservedState `json:"activityLevel,omitempty"`

	// Optional. Data for points in the `vo2-max` sample data type collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.vo2_max
	Vo2Max *Vo2MaxObservedState `json:"vo2Max,omitempty"`

	// Optional. Data for points in the `nutrition-log` session data type
	//  collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.nutrition_log
	NutritionLog *NutritionLogObservedState `json:"nutritionLog,omitempty"`

	// Optional. Data for points in the `irregular-rhythm-notification` session
	//  data type collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.irregular_rhythm_notification
	IrregularRhythmNotification *IrregularRhythmNotificationObservedState `json:"irregularRhythmNotification,omitempty"`

	// Optional. Data for points in the `electrocardiogram` session data type
	//  collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.electrocardiogram
	Electrocardiogram *ElectrocardiogramObservedState `json:"electrocardiogram,omitempty"`

	// Optional. Data for points in the `hydration-log` session data type
	//  collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.hydration_log
	HydrationLog *HydrationLogObservedState `json:"hydrationLog,omitempty"`

	// Optional. The food details.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.food
	Food *FoodObservedState `json:"food,omitempty"`

	// Optional. Data for points in the `time-in-heart-rate-zone` interval data
	//  type collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.time_in_heart_rate_zone
	TimeInHeartRateZone *TimeInHeartRateZoneObservedState `json:"timeInHeartRateZone,omitempty"`

	// Optional. Data for points in the `active-minutes` interval data type
	//  collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.active_minutes
	ActiveMinutes *ActiveMinutesObservedState `json:"activeMinutes,omitempty"`

	// Optional. Data for points in the `respiratory-rate-sleep-summary` sample
	//  data type collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.respiratory_rate_sleep_summary
	RespiratoryRateSleepSummary *RespiratoryRateSleepSummaryObservedState `json:"respiratoryRateSleepSummary,omitempty"`

	// Optional. Data for points in the `swim-lengths-data` interval data type
	//  collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.swim_lengths_data
	SwimLengthsData *SwimLengthsDataObservedState `json:"swimLengthsData,omitempty"`

	// Optional. Data for points in the `height` sample data type collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.height
	Height *HeightObservedState `json:"height,omitempty"`

	// Optional. Data for points in the `basal-energy-burned` interval data type
	//  collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.basal_energy_burned
	BasalEnergyBurned *BasalEnergyBurnedObservedState `json:"basalEnergyBurned,omitempty"`

	// Optional. Data for points in the `core-body-temperature` sample data type
	//  collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.core_body_temperature
	CoreBodyTemperature *CoreBodyTemperatureObservedState `json:"coreBodyTemperature,omitempty"`

	// Optional. Data for points in the `active-energy-burned` interval data
	//  type collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.active_energy_burned
	ActiveEnergyBurned *ActiveEnergyBurnedObservedState `json:"activeEnergyBurned,omitempty"`

	// Optional. Data for points in the `blood-glucose` sample data type
	//  collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.blood_glucose
	BloodGlucose *BloodGlucoseObservedState `json:"bloodGlucose,omitempty"`

	// Optional. Data for points in the `menstrual-period` interval data type
	//  collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.menstrual_period
	MenstrualPeriod *MenstrualPeriodObservedState `json:"menstrualPeriod,omitempty"`

	// Optional. Data for points in the `ovulation-test` sample data type
	//  collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.ovulation_test
	OvulationTest *OvulationTestObservedState `json:"ovulationTest,omitempty"`

	// Optional. Data for points in the `symptoms` sample data type
	//  collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.symptoms
	Symptoms *SymptomsObservedState `json:"symptoms,omitempty"`

	// Optional. Data for points in the `moods` sample data type collection.
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.moods
	Moods *MoodsObservedState `json:"moods,omitempty"`

	// Optional. Data source information for the metric
	// +kcc:proto:field=google.devicesandservices.health.v4.DataPoint.data_source
	DataSource *DataSourceObservedState `json:"dataSource,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcphealthdatapoint;gcphealthdatapoints
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// HealthDataPoint is the Schema for the HealthDataPoint API
// +k8s:openapi-gen=true
type HealthDataPoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   HealthDataPointSpec   `json:"spec,omitempty"`
	Status HealthDataPointStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// HealthDataPointList contains a list of HealthDataPoint
type HealthDataPointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HealthDataPoint `json:"items"`
}

func init() {
	SchemeBuilder.Register(&HealthDataPoint{}, &HealthDataPointList{})
}
