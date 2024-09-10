/*
HPC Portal - API

An interface for working with HPC clusters.

API version: 1.9.10-9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// SortBy the model 'SortBy'
type SortBy string

// List of SortBy
const (
	NONE SortBy = "None"
	JOB_NAME SortBy = "JobName"
	OWNER SortBy = "Owner"
	CREATION_TIME SortBy = "CreationTime"
	GPU_UTILIZATION SortBy = "GpuUtilization"
	GPU_COUNT SortBy = "GpuCount"
	STATUS SortBy = "Status"
	CLUSTER SortBy = "Cluster"
)

// All allowed values of SortBy enum
var AllowedSortByEnumValues = []SortBy{
	"None",
	"JobName",
	"Owner",
	"CreationTime",
	"GpuUtilization",
	"GpuCount",
	"Status",
	"Cluster",
}

func (v *SortBy) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SortBy(value)
	for _, existing := range AllowedSortByEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SortBy", value)
}

// NewSortByFromValue returns a pointer to a valid SortBy
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSortByFromValue(v string) (*SortBy, error) {
	ev := SortBy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SortBy: valid values are %v", v, AllowedSortByEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SortBy) IsValid() bool {
	for _, existing := range AllowedSortByEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SortBy value
func (v SortBy) Ptr() *SortBy {
	return &v
}

type NullableSortBy struct {
	value *SortBy
	isSet bool
}

func (v NullableSortBy) Get() *SortBy {
	return v.value
}

func (v *NullableSortBy) Set(val *SortBy) {
	v.value = val
	v.isSet = true
}

func (v NullableSortBy) IsSet() bool {
	return v.isSet
}

func (v *NullableSortBy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSortBy(val *SortBy) *NullableSortBy {
	return &NullableSortBy{value: val, isSet: true}
}

func (v NullableSortBy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSortBy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

