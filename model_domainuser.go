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

// Domainuser struct for Domainuser
type Domainuser struct {
	DomainuserId *int32 `json:"domainuser_id,omitempty"`
	DomainuserName string `json:"domainuser_name"`
	Domain int32 `json:"domain"`
	DomainuserIsDomainadmin NullableBool `json:"domainuser_is_domainadmin,omitempty"`
}

// NewDomainuser instantiates a new Domainuser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainuser(domainuserName string, domain int32, ) *Domainuser {
	this := Domainuser{}
	this.DomainuserName = domainuserName
	this.Domain = domain
	return &this
}

// NewDomainuserWithDefaults instantiates a new Domainuser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainuserWithDefaults() *Domainuser {
	this := Domainuser{}
	return &this
}

// GetDomainuserId returns the DomainuserId field value if set, zero value otherwise.
func (o *Domainuser) GetDomainuserId() int32 {
	if o == nil || o.DomainuserId == nil {
		var ret int32
		return ret
	}
	return *o.DomainuserId
}

// GetDomainuserIdOk returns a tuple with the DomainuserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domainuser) GetDomainuserIdOk() (*int32, bool) {
	if o == nil || o.DomainuserId == nil {
		return nil, false
	}
	return o.DomainuserId, true
}

// HasDomainuserId returns a boolean if a field has been set.
func (o *Domainuser) HasDomainuserId() bool {
	if o != nil && o.DomainuserId != nil {
		return true
	}

	return false
}

// SetDomainuserId gets a reference to the given int32 and assigns it to the DomainuserId field.
func (o *Domainuser) SetDomainuserId(v int32) {
	o.DomainuserId = &v
}

// GetDomainuserName returns the DomainuserName field value
func (o *Domainuser) GetDomainuserName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.DomainuserName
}

// GetDomainuserNameOk returns a tuple with the DomainuserName field value
// and a boolean to check if the value has been set.
func (o *Domainuser) GetDomainuserNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DomainuserName, true
}

// SetDomainuserName sets field value
func (o *Domainuser) SetDomainuserName(v string) {
	o.DomainuserName = v
}

// GetDomain returns the Domain field value
func (o *Domainuser) GetDomain() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *Domainuser) GetDomainOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *Domainuser) SetDomain(v int32) {
	o.Domain = v
}

// GetDomainuserIsDomainadmin returns the DomainuserIsDomainadmin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Domainuser) GetDomainuserIsDomainadmin() bool {
	if o == nil || o.DomainuserIsDomainadmin.Get() == nil {
		var ret bool
		return ret
	}
	return *o.DomainuserIsDomainadmin.Get()
}

// GetDomainuserIsDomainadminOk returns a tuple with the DomainuserIsDomainadmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Domainuser) GetDomainuserIsDomainadminOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DomainuserIsDomainadmin.Get(), o.DomainuserIsDomainadmin.IsSet()
}

// HasDomainuserIsDomainadmin returns a boolean if a field has been set.
func (o *Domainuser) HasDomainuserIsDomainadmin() bool {
	if o != nil && o.DomainuserIsDomainadmin.IsSet() {
		return true
	}

	return false
}

// SetDomainuserIsDomainadmin gets a reference to the given NullableBool and assigns it to the DomainuserIsDomainadmin field.
func (o *Domainuser) SetDomainuserIsDomainadmin(v bool) {
	o.DomainuserIsDomainadmin.Set(&v)
}
// SetDomainuserIsDomainadminNil sets the value for DomainuserIsDomainadmin to be an explicit nil
func (o *Domainuser) SetDomainuserIsDomainadminNil() {
	o.DomainuserIsDomainadmin.Set(nil)
}

// UnsetDomainuserIsDomainadmin ensures that no value is present for DomainuserIsDomainadmin, not even an explicit nil
func (o *Domainuser) UnsetDomainuserIsDomainadmin() {
	o.DomainuserIsDomainadmin.Unset()
}

func (o Domainuser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DomainuserId != nil {
		toSerialize["domainuser_id"] = o.DomainuserId
	}
	if true {
		toSerialize["domainuser_name"] = o.DomainuserName
	}
	if true {
		toSerialize["domain"] = o.Domain
	}
	if o.DomainuserIsDomainadmin.IsSet() {
		toSerialize["domainuser_is_domainadmin"] = o.DomainuserIsDomainadmin.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableDomainuser struct {
	value *Domainuser
	isSet bool
}

func (v NullableDomainuser) Get() *Domainuser {
	return v.value
}

func (v *NullableDomainuser) Set(val *Domainuser) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainuser) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainuser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainuser(val *Domainuser) *NullableDomainuser {
	return &NullableDomainuser{value: val, isSet: true}
}

func (v NullableDomainuser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainuser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


