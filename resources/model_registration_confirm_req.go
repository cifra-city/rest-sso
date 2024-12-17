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

// checks if the RegistrationConfirmReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegistrationConfirmReq{}

// RegistrationConfirmReq struct for RegistrationConfirmReq
type RegistrationConfirmReq struct {
	Data RegistrationConfirmData `json:"data"`
}

type _RegistrationConfirmReq RegistrationConfirmReq

// NewRegistrationConfirmReq instantiates a new RegistrationConfirmReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistrationConfirmReq(data RegistrationConfirmData) *RegistrationConfirmReq {
	this := RegistrationConfirmReq{}
	this.Data = data
	return &this
}

// NewRegistrationConfirmReqWithDefaults instantiates a new RegistrationConfirmReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistrationConfirmReqWithDefaults() *RegistrationConfirmReq {
	this := RegistrationConfirmReq{}
	return &this
}

// GetData returns the Data field value
func (o *RegistrationConfirmReq) GetData() RegistrationConfirmData {
	if o == nil {
		var ret RegistrationConfirmData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *RegistrationConfirmReq) GetDataOk() (*RegistrationConfirmData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *RegistrationConfirmReq) SetData(v RegistrationConfirmData) {
	o.Data = v
}

func (o RegistrationConfirmReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegistrationConfirmReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *RegistrationConfirmReq) UnmarshalJSON(data []byte) (err error) {
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

	varRegistrationConfirmReq := _RegistrationConfirmReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRegistrationConfirmReq)

	if err != nil {
		return err
	}

	*o = RegistrationConfirmReq(varRegistrationConfirmReq)

	return err
}

type NullableRegistrationConfirmReq struct {
	value *RegistrationConfirmReq
	isSet bool
}

func (v NullableRegistrationConfirmReq) Get() *RegistrationConfirmReq {
	return v.value
}

func (v *NullableRegistrationConfirmReq) Set(val *RegistrationConfirmReq) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistrationConfirmReq) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistrationConfirmReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistrationConfirmReq(val *RegistrationConfirmReq) *NullableRegistrationConfirmReq {
	return &NullableRegistrationConfirmReq{value: val, isSet: true}
}

func (v NullableRegistrationConfirmReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistrationConfirmReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


