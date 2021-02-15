/*
 * DFIRTrack
 *
 * OpenAPI 3 - Documentation of DFIRTrack API
 *
 * API version: v0.4.7
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dfirtrackapi

import (
	"encoding/json"
)

// Tagcolor struct for Tagcolor
type Tagcolor struct {
	TagcolorId *int32 `json:"tagcolor_id,omitempty"`
	TagcolorName string `json:"tagcolor_name"`
}

// NewTagcolor instantiates a new Tagcolor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagcolor(tagcolorName string, ) *Tagcolor {
	this := Tagcolor{}
	this.TagcolorName = tagcolorName
	return &this
}

// NewTagcolorWithDefaults instantiates a new Tagcolor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagcolorWithDefaults() *Tagcolor {
	this := Tagcolor{}
	return &this
}

// GetTagcolorId returns the TagcolorId field value if set, zero value otherwise.
func (o *Tagcolor) GetTagcolorId() int32 {
	if o == nil || o.TagcolorId == nil {
		var ret int32
		return ret
	}
	return *o.TagcolorId
}

// GetTagcolorIdOk returns a tuple with the TagcolorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tagcolor) GetTagcolorIdOk() (*int32, bool) {
	if o == nil || o.TagcolorId == nil {
		return nil, false
	}
	return o.TagcolorId, true
}

// HasTagcolorId returns a boolean if a field has been set.
func (o *Tagcolor) HasTagcolorId() bool {
	if o != nil && o.TagcolorId != nil {
		return true
	}

	return false
}

// SetTagcolorId gets a reference to the given int32 and assigns it to the TagcolorId field.
func (o *Tagcolor) SetTagcolorId(v int32) {
	o.TagcolorId = &v
}

// GetTagcolorName returns the TagcolorName field value
func (o *Tagcolor) GetTagcolorName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.TagcolorName
}

// GetTagcolorNameOk returns a tuple with the TagcolorName field value
// and a boolean to check if the value has been set.
func (o *Tagcolor) GetTagcolorNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TagcolorName, true
}

// SetTagcolorName sets field value
func (o *Tagcolor) SetTagcolorName(v string) {
	o.TagcolorName = v
}

func (o Tagcolor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TagcolorId != nil {
		toSerialize["tagcolor_id"] = o.TagcolorId
	}
	if true {
		toSerialize["tagcolor_name"] = o.TagcolorName
	}
	return json.Marshal(toSerialize)
}

type NullableTagcolor struct {
	value *Tagcolor
	isSet bool
}

func (v NullableTagcolor) Get() *Tagcolor {
	return v.value
}

func (v *NullableTagcolor) Set(val *Tagcolor) {
	v.value = val
	v.isSet = true
}

func (v NullableTagcolor) IsSet() bool {
	return v.isSet
}

func (v *NullableTagcolor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagcolor(val *Tagcolor) *NullableTagcolor {
	return &NullableTagcolor{value: val, isSet: true}
}

func (v NullableTagcolor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagcolor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


