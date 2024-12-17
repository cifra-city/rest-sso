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

// checks if the ActivateEmail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActivateEmail{}

// ActivateEmail struct for ActivateEmail
type ActivateEmail struct {
	Data ActivateEmailData `json:"data"`
}

type _ActivateEmail ActivateEmail

// NewActivateEmail instantiates a new ActivateEmail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivateEmail(data ActivateEmailData) *ActivateEmail {
	this := ActivateEmail{}
	this.Data = data
	return &this
}

// NewActivateEmailWithDefaults instantiates a new ActivateEmail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivateEmailWithDefaults() *ActivateEmail {
	this := ActivateEmail{}
	return &this
}

// GetData returns the Data field value
func (o *ActivateEmail) GetData() ActivateEmailData {
	if o == nil {
		var ret ActivateEmailData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ActivateEmail) GetDataOk() (*ActivateEmailData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ActivateEmail) SetData(v ActivateEmailData) {
	o.Data = v
}

func (o ActivateEmail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActivateEmail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *ActivateEmail) UnmarshalJSON(data []byte) (err error) {
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

	varActivateEmail := _ActivateEmail{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varActivateEmail)

	if err != nil {
		return err
	}

	*o = ActivateEmail(varActivateEmail)

	return err
}

type NullableActivateEmail struct {
	value *ActivateEmail
	isSet bool
}

func (v NullableActivateEmail) Get() *ActivateEmail {
	return v.value
}

func (v *NullableActivateEmail) Set(val *ActivateEmail) {
	v.value = val
	v.isSet = true
}

func (v NullableActivateEmail) IsSet() bool {
	return v.isSet
}

func (v *NullableActivateEmail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivateEmail(val *ActivateEmail) *NullableActivateEmail {
	return &NullableActivateEmail{value: val, isSet: true}
}

func (v NullableActivateEmail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivateEmail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


