/*
Cifra SSO REST API

SSO REST API for Cifra app

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package resources

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ResetPasswordCompleteDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResetPasswordCompleteDataAttributes{}

// ResetPasswordCompleteDataAttributes struct for ResetPasswordCompleteDataAttributes
type ResetPasswordCompleteDataAttributes struct {
	// User password.
	FirstPassword string `json:"first_password"`
	// User password.
	SecondPassword string `json:"second_password"`
	// User email (required if username is not provided).
	Email *string `json:"email,omitempty"`
	// User username (required if email is not provided).
	Username *string `json:"username,omitempty"`
}

type _ResetPasswordCompleteDataAttributes ResetPasswordCompleteDataAttributes

// NewResetPasswordCompleteDataAttributes instantiates a new ResetPasswordCompleteDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResetPasswordCompleteDataAttributes(firstPassword string, secondPassword string) *ResetPasswordCompleteDataAttributes {
	this := ResetPasswordCompleteDataAttributes{}
	this.FirstPassword = firstPassword
	this.SecondPassword = secondPassword
	return &this
}

// NewResetPasswordCompleteDataAttributesWithDefaults instantiates a new ResetPasswordCompleteDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResetPasswordCompleteDataAttributesWithDefaults() *ResetPasswordCompleteDataAttributes {
	this := ResetPasswordCompleteDataAttributes{}
	return &this
}

// GetFirstPassword returns the FirstPassword field value
func (o *ResetPasswordCompleteDataAttributes) GetFirstPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstPassword
}

// GetFirstPasswordOk returns a tuple with the FirstPassword field value
// and a boolean to check if the value has been set.
func (o *ResetPasswordCompleteDataAttributes) GetFirstPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstPassword, true
}

// SetFirstPassword sets field value
func (o *ResetPasswordCompleteDataAttributes) SetFirstPassword(v string) {
	o.FirstPassword = v
}

// GetSecondPassword returns the SecondPassword field value
func (o *ResetPasswordCompleteDataAttributes) GetSecondPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecondPassword
}

// GetSecondPasswordOk returns a tuple with the SecondPassword field value
// and a boolean to check if the value has been set.
func (o *ResetPasswordCompleteDataAttributes) GetSecondPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecondPassword, true
}

// SetSecondPassword sets field value
func (o *ResetPasswordCompleteDataAttributes) SetSecondPassword(v string) {
	o.SecondPassword = v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *ResetPasswordCompleteDataAttributes) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResetPasswordCompleteDataAttributes) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *ResetPasswordCompleteDataAttributes) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *ResetPasswordCompleteDataAttributes) SetEmail(v string) {
	o.Email = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ResetPasswordCompleteDataAttributes) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResetPasswordCompleteDataAttributes) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ResetPasswordCompleteDataAttributes) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ResetPasswordCompleteDataAttributes) SetUsername(v string) {
	o.Username = &v
}

func (o ResetPasswordCompleteDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResetPasswordCompleteDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["first_password"] = o.FirstPassword
	toSerialize["second_password"] = o.SecondPassword
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

func (o *ResetPasswordCompleteDataAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"first_password",
		"second_password",
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

	varResetPasswordCompleteDataAttributes := _ResetPasswordCompleteDataAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResetPasswordCompleteDataAttributes)

	if err != nil {
		return err
	}

	*o = ResetPasswordCompleteDataAttributes(varResetPasswordCompleteDataAttributes)

	return err
}

type NullableResetPasswordCompleteDataAttributes struct {
	value *ResetPasswordCompleteDataAttributes
	isSet bool
}

func (v NullableResetPasswordCompleteDataAttributes) Get() *ResetPasswordCompleteDataAttributes {
	return v.value
}

func (v *NullableResetPasswordCompleteDataAttributes) Set(val *ResetPasswordCompleteDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableResetPasswordCompleteDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableResetPasswordCompleteDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResetPasswordCompleteDataAttributes(val *ResetPasswordCompleteDataAttributes) *NullableResetPasswordCompleteDataAttributes {
	return &NullableResetPasswordCompleteDataAttributes{value: val, isSet: true}
}

func (v NullableResetPasswordCompleteDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResetPasswordCompleteDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


