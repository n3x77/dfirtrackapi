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

// AuthToken struct for AuthToken
type AuthToken struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token *string `json:"token,omitempty"`
}

// NewAuthToken instantiates a new AuthToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthToken(username string, password string, ) *AuthToken {
	this := AuthToken{}
	this.Username = username
	this.Password = password
	return &this
}

// NewAuthTokenWithDefaults instantiates a new AuthToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthTokenWithDefaults() *AuthToken {
	this := AuthToken{}
	return &this
}

// GetUsername returns the Username field value
func (o *AuthToken) GetUsername() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *AuthToken) GetUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *AuthToken) SetUsername(v string) {
	o.Username = v
}

// GetPassword returns the Password field value
func (o *AuthToken) GetPassword() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *AuthToken) GetPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *AuthToken) SetPassword(v string) {
	o.Password = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *AuthToken) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthToken) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *AuthToken) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *AuthToken) SetToken(v string) {
	o.Token = &v
}

func (o AuthToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["username"] = o.Username
	}
	if true {
		toSerialize["password"] = o.Password
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	return json.Marshal(toSerialize)
}

type NullableAuthToken struct {
	value *AuthToken
	isSet bool
}

func (v NullableAuthToken) Get() *AuthToken {
	return v.value
}

func (v *NullableAuthToken) Set(val *AuthToken) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthToken) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthToken(val *AuthToken) *NullableAuthToken {
	return &NullableAuthToken{value: val, isSet: true}
}

func (v NullableAuthToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

