/*
Cifra SSO REST API

SSO REST API for Cifra app

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package resources

import (
	"encoding/json"
)

// checks if the ResetPasswordInitiateDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResetPasswordInitiateDataAttributes{}

// ResetPasswordInitiateDataAttributes struct for ResetPasswordInitiateDataAttributes
type ResetPasswordInitiateDataAttributes struct {
	// User email (required if username is not provided).
	Email *string `json:"email,omitempty"`
	// User username (required if email is not provided).
	Username *string `json:"username,omitempty"`
}

// NewResetPasswordInitiateDataAttributes instantiates a new ResetPasswordInitiateDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResetPasswordInitiateDataAttributes() *ResetPasswordInitiateDataAttributes {
	this := ResetPasswordInitiateDataAttributes{}
	return &this
}

// NewResetPasswordInitiateDataAttributesWithDefaults instantiates a new ResetPasswordInitiateDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResetPasswordInitiateDataAttributesWithDefaults() *ResetPasswordInitiateDataAttributes {
	this := ResetPasswordInitiateDataAttributes{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *ResetPasswordInitiateDataAttributes) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResetPasswordInitiateDataAttributes) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *ResetPasswordInitiateDataAttributes) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *ResetPasswordInitiateDataAttributes) SetEmail(v string) {
	o.Email = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ResetPasswordInitiateDataAttributes) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResetPasswordInitiateDataAttributes) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ResetPasswordInitiateDataAttributes) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ResetPasswordInitiateDataAttributes) SetUsername(v string) {
	o.Username = &v
}

func (o ResetPasswordInitiateDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResetPasswordInitiateDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

type NullableResetPasswordInitiateDataAttributes struct {
	value *ResetPasswordInitiateDataAttributes
	isSet bool
}

func (v NullableResetPasswordInitiateDataAttributes) Get() *ResetPasswordInitiateDataAttributes {
	return v.value
}

func (v *NullableResetPasswordInitiateDataAttributes) Set(val *ResetPasswordInitiateDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableResetPasswordInitiateDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableResetPasswordInitiateDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResetPasswordInitiateDataAttributes(val *ResetPasswordInitiateDataAttributes) *NullableResetPasswordInitiateDataAttributes {
	return &NullableResetPasswordInitiateDataAttributes{value: val, isSet: true}
}

func (v NullableResetPasswordInitiateDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResetPasswordInitiateDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


