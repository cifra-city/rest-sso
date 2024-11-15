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

// checks if the Registration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Registration{}

// Registration struct for Registration
type Registration struct {
	Data RegistrationData `json:"data"`
}

type _Registration Registration

// NewRegistration instantiates a new Registration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistration(data RegistrationData) *Registration {
	this := Registration{}
	this.Data = data
	return &this
}

// NewRegistrationWithDefaults instantiates a new Registration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistrationWithDefaults() *Registration {
	this := Registration{}
	return &this
}

// GetData returns the Data field value
func (o *Registration) GetData() RegistrationData {
	if o == nil {
		var ret RegistrationData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *Registration) GetDataOk() (*RegistrationData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *Registration) SetData(v RegistrationData) {
	o.Data = v
}

func (o Registration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Registration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *Registration) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varRegistration := _Registration{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRegistration)

	if err != nil {
		return err
	}

	*o = Registration(varRegistration)

	return err
}

type NullableRegistration struct {
	value *Registration
	isSet bool
}

func (v NullableRegistration) Get() *Registration {
	return v.value
}

func (v *NullableRegistration) Set(val *Registration) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistration) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistration(val *Registration) *NullableRegistration {
	return &NullableRegistration{value: val, isSet: true}
}

func (v NullableRegistration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


