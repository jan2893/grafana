package v0alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type TimeIntervals struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec TimeIntervalsSpec `json:"spec,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type TimeIntervalsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []TimeIntervals `json:"items,omitempty"`
}

type TimeIntervalsSpec struct {
	Intervals []TimeInterval `json:"time_intervals,omitempty"`
}

type TimeInterval struct {
	Times       []TimeIntervalTimeRange `json:"times,omitempty" hcl:"times,block"`
	Weekdays    *[]string               `json:"weekdays,omitempty" hcl:"weekdays"`
	DaysOfMonth *[]string               `json:"days_of_month,omitempty" hcl:"days_of_month"`
	Months      *[]string               `json:"months,omitempty" hcl:"months"`
	Years       *[]string               `json:"years,omitempty" hcl:"years"`
	Location    *string                 `json:"location,omitempty" hcl:"location"`
}

type TimeIntervalTimeRange struct {
	StartMinute string `json:"start_time" hcl:"start"`
	EndMinute   string `json:"end_time" hcl:"end"`
}
