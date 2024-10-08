/*
HPC Portal - API

An interface for working with HPC clusters.

API version: 1.9.10-9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the Capacity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Capacity{}

// Capacity struct for Capacity
type Capacity struct {
	Storage string `json:"storage"`
}

type _Capacity Capacity

// NewCapacity instantiates a new Capacity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapacity(storage string) *Capacity {
	this := Capacity{}
	this.Storage = storage
	return &this
}

// NewCapacityWithDefaults instantiates a new Capacity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapacityWithDefaults() *Capacity {
	this := Capacity{}
	return &this
}

// GetStorage returns the Storage field value
func (o *Capacity) GetStorage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value
// and a boolean to check if the value has been set.
func (o *Capacity) GetStorageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Storage, true
}

// SetStorage sets field value
func (o *Capacity) SetStorage(v string) {
	o.Storage = v
}

func (o Capacity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Capacity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["storage"] = o.Storage
	return toSerialize, nil
}

func (o *Capacity) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"storage",
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

	varCapacity := _Capacity{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCapacity)

	if err != nil {
		return err
	}

	*o = Capacity(varCapacity)

	return err
}

type NullableCapacity struct {
	value *Capacity
	isSet bool
}

func (v NullableCapacity) Get() *Capacity {
	return v.value
}

func (v *NullableCapacity) Set(val *Capacity) {
	v.value = val
	v.isSet = true
}

func (v NullableCapacity) IsSet() bool {
	return v.isSet
}

func (v *NullableCapacity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapacity(val *Capacity) *NullableCapacity {
	return &NullableCapacity{value: val, isSet: true}
}

func (v NullableCapacity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapacity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


