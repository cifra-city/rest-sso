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

// checks if the RegistrationDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegistrationDataAttributes{}

// RegistrationDataAttributes struct for RegistrationDataAttributes
type RegistrationDataAttributes struct {
	// User email
	Email string `json:"email"`
	// User password
	Password string `json:"password"`
}

type _RegistrationDataAttributes RegistrationDataAttributes

// NewRegistrationDataAttributes instantiates a new RegistrationDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistrationDataAttributes(email string, password string) *RegistrationDataAttributes {
	this := RegistrationDataAttributes{}
	this.Email = email
	this.Password = password
	return &this
}

// NewRegistrationDataAttributesWithDefaults instantiates a new RegistrationDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistrationDataAttributesWithDefaults() *RegistrationDataAttributes {
	this := RegistrationDataAttributes{}
	return &this
}

// GetEmail returns the Email field value
func (o *RegistrationDataAttributes) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *RegistrationDataAttributes) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *RegistrationDataAttributes) SetEmail(v string) {
	o.Email = v
}

// GetPassword returns the Password field value
func (o *RegistrationDataAttributes) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *RegistrationDataAttributes) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *RegistrationDataAttributes) SetPassword(v string) {
	o.Password = v
}

func (o RegistrationDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegistrationDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	toSerialize["password"] = o.Password
	return toSerialize, nil
}

func (o *RegistrationDataAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
		"password",
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

	varRegistrationDataAttributes := _RegistrationDataAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRegistrationDataAttributes)

	if err != nil {
		return err
	}

	*o = RegistrationDataAttributes(varRegistrationDataAttributes)

	return err
}

type NullableRegistrationDataAttributes struct {
	value *RegistrationDataAttributes
	isSet bool
}

func (v NullableRegistrationDataAttributes) Get() *RegistrationDataAttributes {
	return v.value
}

func (v *NullableRegistrationDataAttributes) Set(val *RegistrationDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistrationDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistrationDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistrationDataAttributes(val *RegistrationDataAttributes) *NullableRegistrationDataAttributes {
	return &NullableRegistrationDataAttributes{value: val, isSet: true}
}

func (v NullableRegistrationDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistrationDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


