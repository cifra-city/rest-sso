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

// checks if the LoginCompleteReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoginCompleteReq{}

// LoginCompleteReq struct for LoginCompleteReq
type LoginCompleteReq struct {
	Data LoginCompleteReqData `json:"data"`
}

type _LoginCompleteReq LoginCompleteReq

// NewLoginCompleteReq instantiates a new LoginCompleteReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoginCompleteReq(data LoginCompleteReqData) *LoginCompleteReq {
	this := LoginCompleteReq{}
	this.Data = data
	return &this
}

// NewLoginCompleteReqWithDefaults instantiates a new LoginCompleteReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoginCompleteReqWithDefaults() *LoginCompleteReq {
	this := LoginCompleteReq{}
	return &this
}

// GetData returns the Data field value
func (o *LoginCompleteReq) GetData() LoginCompleteReqData {
	if o == nil {
		var ret LoginCompleteReqData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *LoginCompleteReq) GetDataOk() (*LoginCompleteReqData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *LoginCompleteReq) SetData(v LoginCompleteReqData) {
	o.Data = v
}

func (o LoginCompleteReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoginCompleteReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *LoginCompleteReq) UnmarshalJSON(data []byte) (err error) {
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

	varLoginCompleteReq := _LoginCompleteReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLoginCompleteReq)

	if err != nil {
		return err
	}

	*o = LoginCompleteReq(varLoginCompleteReq)

	return err
}

type NullableLoginCompleteReq struct {
	value *LoginCompleteReq
	isSet bool
}

func (v NullableLoginCompleteReq) Get() *LoginCompleteReq {
	return v.value
}

func (v *NullableLoginCompleteReq) Set(val *LoginCompleteReq) {
	v.value = val
	v.isSet = true
}

func (v NullableLoginCompleteReq) IsSet() bool {
	return v.isSet
}

func (v *NullableLoginCompleteReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoginCompleteReq(val *LoginCompleteReq) *NullableLoginCompleteReq {
	return &NullableLoginCompleteReq{value: val, isSet: true}
}

func (v NullableLoginCompleteReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoginCompleteReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

