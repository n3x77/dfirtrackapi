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

// Dnsname struct for Dnsname
type Dnsname struct {
	DnsnameId *int32 `json:"dnsname_id,omitempty"`
	DnsnameName string `json:"dnsname_name"`
	Domain NullableInt32 `json:"domain,omitempty"`
}

// NewDnsname instantiates a new Dnsname object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnsname(dnsnameName string, ) *Dnsname {
	this := Dnsname{}
	this.DnsnameName = dnsnameName
	return &this
}

// NewDnsnameWithDefaults instantiates a new Dnsname object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnsnameWithDefaults() *Dnsname {
	this := Dnsname{}
	return &this
}

// GetDnsnameId returns the DnsnameId field value if set, zero value otherwise.
func (o *Dnsname) GetDnsnameId() int32 {
	if o == nil || o.DnsnameId == nil {
		var ret int32
		return ret
	}
	return *o.DnsnameId
}

// GetDnsnameIdOk returns a tuple with the DnsnameId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dnsname) GetDnsnameIdOk() (*int32, bool) {
	if o == nil || o.DnsnameId == nil {
		return nil, false
	}
	return o.DnsnameId, true
}

// HasDnsnameId returns a boolean if a field has been set.
func (o *Dnsname) HasDnsnameId() bool {
	if o != nil && o.DnsnameId != nil {
		return true
	}

	return false
}

// SetDnsnameId gets a reference to the given int32 and assigns it to the DnsnameId field.
func (o *Dnsname) SetDnsnameId(v int32) {
	o.DnsnameId = &v
}

// GetDnsnameName returns the DnsnameName field value
func (o *Dnsname) GetDnsnameName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.DnsnameName
}

// GetDnsnameNameOk returns a tuple with the DnsnameName field value
// and a boolean to check if the value has been set.
func (o *Dnsname) GetDnsnameNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DnsnameName, true
}

// SetDnsnameName sets field value
func (o *Dnsname) SetDnsnameName(v string) {
	o.DnsnameName = v
}

// GetDomain returns the Domain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Dnsname) GetDomain() int32 {
	if o == nil || o.Domain.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Domain.Get()
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Dnsname) GetDomainOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Domain.Get(), o.Domain.IsSet()
}

// HasDomain returns a boolean if a field has been set.
func (o *Dnsname) HasDomain() bool {
	if o != nil && o.Domain.IsSet() {
		return true
	}

	return false
}

// SetDomain gets a reference to the given NullableInt32 and assigns it to the Domain field.
func (o *Dnsname) SetDomain(v int32) {
	o.Domain.Set(&v)
}
// SetDomainNil sets the value for Domain to be an explicit nil
func (o *Dnsname) SetDomainNil() {
	o.Domain.Set(nil)
}

// UnsetDomain ensures that no value is present for Domain, not even an explicit nil
func (o *Dnsname) UnsetDomain() {
	o.Domain.Unset()
}

func (o Dnsname) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DnsnameId != nil {
		toSerialize["dnsname_id"] = o.DnsnameId
	}
	if true {
		toSerialize["dnsname_name"] = o.DnsnameName
	}
	if o.Domain.IsSet() {
		toSerialize["domain"] = o.Domain.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableDnsname struct {
	value *Dnsname
	isSet bool
}

func (v NullableDnsname) Get() *Dnsname {
	return v.value
}

func (v *NullableDnsname) Set(val *Dnsname) {
	v.value = val
	v.isSet = true
}

func (v NullableDnsname) IsSet() bool {
	return v.isSet
}

func (v *NullableDnsname) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnsname(val *Dnsname) *NullableDnsname {
	return &NullableDnsname{value: val, isSet: true}
}

func (v NullableDnsname) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnsname) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


