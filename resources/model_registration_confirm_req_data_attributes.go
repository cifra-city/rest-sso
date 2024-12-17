/*
Example API

This is a sample API.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package resources

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the RegistrationConfirmReqDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegistrationConfirmReqDataAttributes{}

// RegistrationConfirmReqDataAttributes struct for RegistrationConfirmReqDataAttributes
type RegistrationConfirmReqDataAttributes struct {
	// First password
	FirstPassword string `json:"first_password"`
	// Second password
	SecondPassword string `json:"second_password"`
	// User email
	Email string `json:"email"`
	// User username
	Username *string `json:"username,omitempty"`
}

type _RegistrationConfirmReqDataAttributes RegistrationConfirmReqDataAttributes

// NewRegistrationConfirmReqDataAttributes instantiates a new RegistrationConfirmReqDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistrationConfirmReqDataAttributes(firstPassword string, secondPassword string, email string) *RegistrationConfirmReqDataAttributes {
	this := RegistrationConfirmReqDataAttributes{}
	this.FirstPassword = firstPassword
	this.SecondPassword = secondPassword
	this.Email = email
	return &this
}

// NewRegistrationConfirmReqDataAttributesWithDefaults instantiates a new RegistrationConfirmReqDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistrationConfirmReqDataAttributesWithDefaults() *RegistrationConfirmReqDataAttributes {
	this := RegistrationConfirmReqDataAttributes{}
	return &this
}

// GetFirstPassword returns the FirstPassword field value
func (o *RegistrationConfirmReqDataAttributes) GetFirstPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstPassword
}

// GetFirstPasswordOk returns a tuple with the FirstPassword field value
// and a boolean to check if the value has been set.
func (o *RegistrationConfirmReqDataAttributes) GetFirstPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstPassword, true
}

// SetFirstPassword sets field value
func (o *RegistrationConfirmReqDataAttributes) SetFirstPassword(v string) {
	o.FirstPassword = v
}

// GetSecondPassword returns the SecondPassword field value
func (o *RegistrationConfirmReqDataAttributes) GetSecondPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecondPassword
}

// GetSecondPasswordOk returns a tuple with the SecondPassword field value
// and a boolean to check if the value has been set.
func (o *RegistrationConfirmReqDataAttributes) GetSecondPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecondPassword, true
}

// SetSecondPassword sets field value
func (o *RegistrationConfirmReqDataAttributes) SetSecondPassword(v string) {
	o.SecondPassword = v
}

// GetEmail returns the Email field value
func (o *RegistrationConfirmReqDataAttributes) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *RegistrationConfirmReqDataAttributes) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *RegistrationConfirmReqDataAttributes) SetEmail(v string) {
	o.Email = v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *RegistrationConfirmReqDataAttributes) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationConfirmReqDataAttributes) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *RegistrationConfirmReqDataAttributes) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *RegistrationConfirmReqDataAttributes) SetUsername(v string) {
	o.Username = &v
}

func (o RegistrationConfirmReqDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegistrationConfirmReqDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["first_password"] = o.FirstPassword
	toSerialize["second_password"] = o.SecondPassword
	toSerialize["email"] = o.Email
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

func (o *RegistrationConfirmReqDataAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"first_password",
		"second_password",
		"email",
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

	varRegistrationConfirmReqDataAttributes := _RegistrationConfirmReqDataAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRegistrationConfirmReqDataAttributes)

	if err != nil {
		return err
	}

	*o = RegistrationConfirmReqDataAttributes(varRegistrationConfirmReqDataAttributes)

	return err
}

type NullableRegistrationConfirmReqDataAttributes struct {
	value *RegistrationConfirmReqDataAttributes
	isSet bool
}

func (v NullableRegistrationConfirmReqDataAttributes) Get() *RegistrationConfirmReqDataAttributes {
	return v.value
}

func (v *NullableRegistrationConfirmReqDataAttributes) Set(val *RegistrationConfirmReqDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistrationConfirmReqDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistrationConfirmReqDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistrationConfirmReqDataAttributes(val *RegistrationConfirmReqDataAttributes) *NullableRegistrationConfirmReqDataAttributes {
	return &NullableRegistrationConfirmReqDataAttributes{value: val, isSet: true}
}

func (v NullableRegistrationConfirmReqDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistrationConfirmReqDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

