/*
HPC Portal - API

An interface for working with HPC clusters.

API version: 1.9.10-9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the PVCs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PVCs{}

// PVCs struct for PVCs
type PVCs struct {
	Name *string `json:"name,omitempty"`
	Path NullableString `json:"path,omitempty"`
	MountPath NullableString `json:"mount_path,omitempty"`
}

// NewPVCs instantiates a new PVCs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPVCs() *PVCs {
	this := PVCs{}
	var name string = "non-existing-pvc"
	this.Name = &name
	return &this
}

// NewPVCsWithDefaults instantiates a new PVCs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPVCsWithDefaults() *PVCs {
	this := PVCs{}
	var name string = "non-existing-pvc"
	this.Name = &name
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PVCs) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PVCs) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PVCs) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PVCs) SetName(v string) {
	o.Name = &v
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PVCs) GetPath() string {
	if o == nil || IsNil(o.Path.Get()) {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PVCs) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *PVCs) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *PVCs) SetPath(v string) {
	o.Path.Set(&v)
}
// SetPathNil sets the value for Path to be an explicit nil
func (o *PVCs) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *PVCs) UnsetPath() {
	o.Path.Unset()
}

// GetMountPath returns the MountPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PVCs) GetMountPath() string {
	if o == nil || IsNil(o.MountPath.Get()) {
		var ret string
		return ret
	}
	return *o.MountPath.Get()
}

// GetMountPathOk returns a tuple with the MountPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PVCs) GetMountPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MountPath.Get(), o.MountPath.IsSet()
}

// HasMountPath returns a boolean if a field has been set.
func (o *PVCs) HasMountPath() bool {
	if o != nil && o.MountPath.IsSet() {
		return true
	}

	return false
}

// SetMountPath gets a reference to the given NullableString and assigns it to the MountPath field.
func (o *PVCs) SetMountPath(v string) {
	o.MountPath.Set(&v)
}
// SetMountPathNil sets the value for MountPath to be an explicit nil
func (o *PVCs) SetMountPathNil() {
	o.MountPath.Set(nil)
}

// UnsetMountPath ensures that no value is present for MountPath, not even an explicit nil
func (o *PVCs) UnsetMountPath() {
	o.MountPath.Unset()
}

func (o PVCs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PVCs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.Path.IsSet() {
		toSerialize["path"] = o.Path.Get()
	}
	if o.MountPath.IsSet() {
		toSerialize["mount_path"] = o.MountPath.Get()
	}
	return toSerialize, nil
}

type NullablePVCs struct {
	value *PVCs
	isSet bool
}

func (v NullablePVCs) Get() *PVCs {
	return v.value
}

func (v *NullablePVCs) Set(val *PVCs) {
	v.value = val
	v.isSet = true
}

func (v NullablePVCs) IsSet() bool {
	return v.isSet
}

func (v *NullablePVCs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePVCs(val *PVCs) *NullablePVCs {
	return &NullablePVCs{value: val, isSet: true}
}

func (v NullablePVCs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePVCs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


