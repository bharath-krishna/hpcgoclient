/*
HPC Portal - API

An interface for working with HPC clusters.

API version: 1.9.10-9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the Condition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Condition{}

// Condition struct for Condition
type Condition struct {
	LastTransitionTime NullableTime `json:"last_transition_time"`
	Message NullableString `json:"message"`
	Reason NullableString `json:"reason"`
	Status string `json:"status"`
	Type string `json:"type"`
}

type _Condition Condition

// NewCondition instantiates a new Condition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCondition(lastTransitionTime NullableTime, message NullableString, reason NullableString, status string, type_ string) *Condition {
	this := Condition{}
	this.LastTransitionTime = lastTransitionTime
	this.Message = message
	this.Reason = reason
	this.Status = status
	this.Type = type_
	return &this
}

// NewConditionWithDefaults instantiates a new Condition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConditionWithDefaults() *Condition {
	this := Condition{}
	return &this
}

// GetLastTransitionTime returns the LastTransitionTime field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Condition) GetLastTransitionTime() time.Time {
	if o == nil || o.LastTransitionTime.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastTransitionTime.Get()
}

// GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Condition) GetLastTransitionTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastTransitionTime.Get(), o.LastTransitionTime.IsSet()
}

// SetLastTransitionTime sets field value
func (o *Condition) SetLastTransitionTime(v time.Time) {
	o.LastTransitionTime.Set(&v)
}

// GetMessage returns the Message field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Condition) GetMessage() string {
	if o == nil || o.Message.Get() == nil {
		var ret string
		return ret
	}

	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Condition) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// SetMessage sets field value
func (o *Condition) SetMessage(v string) {
	o.Message.Set(&v)
}

// GetReason returns the Reason field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Condition) GetReason() string {
	if o == nil || o.Reason.Get() == nil {
		var ret string
		return ret
	}

	return *o.Reason.Get()
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Condition) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reason.Get(), o.Reason.IsSet()
}

// SetReason sets field value
func (o *Condition) SetReason(v string) {
	o.Reason.Set(&v)
}

// GetStatus returns the Status field value
func (o *Condition) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Condition) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Condition) SetStatus(v string) {
	o.Status = v
}

// GetType returns the Type field value
func (o *Condition) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Condition) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Condition) SetType(v string) {
	o.Type = v
}

func (o Condition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Condition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["last_transition_time"] = o.LastTransitionTime.Get()
	toSerialize["message"] = o.Message.Get()
	toSerialize["reason"] = o.Reason.Get()
	toSerialize["status"] = o.Status
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *Condition) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"last_transition_time",
		"message",
		"reason",
		"status",
		"type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCondition := _Condition{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCondition)

	if err != nil {
		return err
	}

	*o = Condition(varCondition)

	return err
}

type NullableCondition struct {
	value *Condition
	isSet bool
}

func (v NullableCondition) Get() *Condition {
	return v.value
}

func (v *NullableCondition) Set(val *Condition) {
	v.value = val
	v.isSet = true
}

func (v NullableCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCondition(val *Condition) *NullableCondition {
	return &NullableCondition{value: val, isSet: true}
}

func (v NullableCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


