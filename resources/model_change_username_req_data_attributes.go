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

// checks if the ChangeUsernameReqDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChangeUsernameReqDataAttributes{}

// ChangeUsernameReqDataAttributes struct for ChangeUsernameReqDataAttributes
type ChangeUsernameReqDataAttributes struct {
	// New username
	NewUsername *string `json:"new_username,omitempty"`
	// User password
	Password string `json:"password"`
}

type _ChangeUsernameReqDataAttributes ChangeUsernameReqDataAttributes

// NewChangeUsernameReqDataAttributes instantiates a new ChangeUsernameReqDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangeUsernameReqDataAttributes(password string) *ChangeUsernameReqDataAttributes {
	this := ChangeUsernameReqDataAttributes{}
	this.Password = password
	return &this
}

// NewChangeUsernameReqDataAttributesWithDefaults instantiates a new ChangeUsernameReqDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangeUsernameReqDataAttributesWithDefaults() *ChangeUsernameReqDataAttributes {
	this := ChangeUsernameReqDataAttributes{}
	return &this
}

// GetNewUsername returns the NewUsername field value if set, zero value otherwise.
func (o *ChangeUsernameReqDataAttributes) GetNewUsername() string {
	if o == nil || IsNil(o.NewUsername) {
		var ret string
		return ret
	}
	return *o.NewUsername
}

// GetNewUsernameOk returns a tuple with the NewUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeUsernameReqDataAttributes) GetNewUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.NewUsername) {
		return nil, false
	}
	return o.NewUsername, true
}

// HasNewUsername returns a boolean if a field has been set.
func (o *ChangeUsernameReqDataAttributes) HasNewUsername() bool {
	if o != nil && !IsNil(o.NewUsername) {
		return true
	}

	return false
}

// SetNewUsername gets a reference to the given string and assigns it to the NewUsername field.
func (o *ChangeUsernameReqDataAttributes) SetNewUsername(v string) {
	o.NewUsername = &v
}

// GetPassword returns the Password field value
func (o *ChangeUsernameReqDataAttributes) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *ChangeUsernameReqDataAttributes) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *ChangeUsernameReqDataAttributes) SetPassword(v string) {
	o.Password = v
}

func (o ChangeUsernameReqDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChangeUsernameReqDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NewUsername) {
		toSerialize["new_username"] = o.NewUsername
	}
	toSerialize["password"] = o.Password
	return toSerialize, nil
}

func (o *ChangeUsernameReqDataAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varChangeUsernameReqDataAttributes := _ChangeUsernameReqDataAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChangeUsernameReqDataAttributes)

	if err != nil {
		return err
	}

	*o = ChangeUsernameReqDataAttributes(varChangeUsernameReqDataAttributes)

	return err
}

type NullableChangeUsernameReqDataAttributes struct {
	value *ChangeUsernameReqDataAttributes
	isSet bool
}

func (v NullableChangeUsernameReqDataAttributes) Get() *ChangeUsernameReqDataAttributes {
	return v.value
}

func (v *NullableChangeUsernameReqDataAttributes) Set(val *ChangeUsernameReqDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeUsernameReqDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeUsernameReqDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeUsernameReqDataAttributes(val *ChangeUsernameReqDataAttributes) *NullableChangeUsernameReqDataAttributes {
	return &NullableChangeUsernameReqDataAttributes{value: val, isSet: true}
}

func (v NullableChangeUsernameReqDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeUsernameReqDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


