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

// checks if the RegistrationInitiateDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegistrationInitiateDataAttributes{}

// RegistrationInitiateDataAttributes struct for RegistrationInitiateDataAttributes
type RegistrationInitiateDataAttributes struct {
	// User email
	Email string `json:"email"`
	// User username
	Username *string `json:"username,omitempty"`
}

type _RegistrationInitiateDataAttributes RegistrationInitiateDataAttributes

// NewRegistrationInitiateDataAttributes instantiates a new RegistrationInitiateDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistrationInitiateDataAttributes(email string) *RegistrationInitiateDataAttributes {
	this := RegistrationInitiateDataAttributes{}
	this.Email = email
	return &this
}

// NewRegistrationInitiateDataAttributesWithDefaults instantiates a new RegistrationInitiateDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistrationInitiateDataAttributesWithDefaults() *RegistrationInitiateDataAttributes {
	this := RegistrationInitiateDataAttributes{}
	return &this
}

// GetEmail returns the Email field value
func (o *RegistrationInitiateDataAttributes) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *RegistrationInitiateDataAttributes) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *RegistrationInitiateDataAttributes) SetEmail(v string) {
	o.Email = v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *RegistrationInitiateDataAttributes) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationInitiateDataAttributes) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *RegistrationInitiateDataAttributes) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *RegistrationInitiateDataAttributes) SetUsername(v string) {
	o.Username = &v
}

func (o RegistrationInitiateDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegistrationInitiateDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

func (o *RegistrationInitiateDataAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varRegistrationInitiateDataAttributes := _RegistrationInitiateDataAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRegistrationInitiateDataAttributes)

	if err != nil {
		return err
	}

	*o = RegistrationInitiateDataAttributes(varRegistrationInitiateDataAttributes)

	return err
}

type NullableRegistrationInitiateDataAttributes struct {
	value *RegistrationInitiateDataAttributes
	isSet bool
}

func (v NullableRegistrationInitiateDataAttributes) Get() *RegistrationInitiateDataAttributes {
	return v.value
}

func (v *NullableRegistrationInitiateDataAttributes) Set(val *RegistrationInitiateDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistrationInitiateDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistrationInitiateDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistrationInitiateDataAttributes(val *RegistrationInitiateDataAttributes) *NullableRegistrationInitiateDataAttributes {
	return &NullableRegistrationInitiateDataAttributes{value: val, isSet: true}
}

func (v NullableRegistrationInitiateDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistrationInitiateDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

