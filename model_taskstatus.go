/*
 * DFIRTrack
 *
 * OpenAPI 3 - Documentation of DFIRTrack API
 *
 * API version: v1.5.4
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dfirtrackapi

import (
	"encoding/json"
)

// Taskstatus struct for Taskstatus
type Taskstatus struct {
	TaskstatusId *int32 `json:"taskstatus_id,omitempty"`
	TaskstatusName string `json:"taskstatus_name"`
}

// NewTaskstatus instantiates a new Taskstatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskstatus(taskstatusName string, ) *Taskstatus {
	this := Taskstatus{}
	this.TaskstatusName = taskstatusName
	return &this
}

// NewTaskstatusWithDefaults instantiates a new Taskstatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskstatusWithDefaults() *Taskstatus {
	this := Taskstatus{}
	return &this
}

// GetTaskstatusId returns the TaskstatusId field value if set, zero value otherwise.
func (o *Taskstatus) GetTaskstatusId() int32 {
	if o == nil || o.TaskstatusId == nil {
		var ret int32
		return ret
	}
	return *o.TaskstatusId
}

// GetTaskstatusIdOk returns a tuple with the TaskstatusId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Taskstatus) GetTaskstatusIdOk() (*int32, bool) {
	if o == nil || o.TaskstatusId == nil {
		return nil, false
	}
	return o.TaskstatusId, true
}

// HasTaskstatusId returns a boolean if a field has been set.
func (o *Taskstatus) HasTaskstatusId() bool {
	if o != nil && o.TaskstatusId != nil {
		return true
	}

	return false
}

// SetTaskstatusId gets a reference to the given int32 and assigns it to the TaskstatusId field.
func (o *Taskstatus) SetTaskstatusId(v int32) {
	o.TaskstatusId = &v
}

// GetTaskstatusName returns the TaskstatusName field value
func (o *Taskstatus) GetTaskstatusName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.TaskstatusName
}

// GetTaskstatusNameOk returns a tuple with the TaskstatusName field value
// and a boolean to check if the value has been set.
func (o *Taskstatus) GetTaskstatusNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TaskstatusName, true
}

// SetTaskstatusName sets field value
func (o *Taskstatus) SetTaskstatusName(v string) {
	o.TaskstatusName = v
}

func (o Taskstatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TaskstatusId != nil {
		toSerialize["taskstatus_id"] = o.TaskstatusId
	}
	if true {
		toSerialize["taskstatus_name"] = o.TaskstatusName
	}
	return json.Marshal(toSerialize)
}

type NullableTaskstatus struct {
	value *Taskstatus
	isSet bool
}

func (v NullableTaskstatus) Get() *Taskstatus {
	return v.value
}

func (v *NullableTaskstatus) Set(val *Taskstatus) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskstatus) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskstatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskstatus(val *Taskstatus) *NullableTaskstatus {
	return &NullableTaskstatus{value: val, isSet: true}
}

func (v NullableTaskstatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskstatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


