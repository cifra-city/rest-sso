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

// checks if the LoginReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoginReq{}

// LoginReq struct for LoginReq
type LoginReq struct {
	Data LoginReqData `json:"data"`
}

type _LoginReq LoginReq

// NewLoginReq instantiates a new LoginReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoginReq(data LoginReqData) *LoginReq {
	this := LoginReq{}
	this.Data = data
	return &this
}

// NewLoginReqWithDefaults instantiates a new LoginReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoginReqWithDefaults() *LoginReq {
	this := LoginReq{}
	return &this
}

// GetData returns the Data field value
func (o *LoginReq) GetData() LoginReqData {
	if o == nil {
		var ret LoginReqData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *LoginReq) GetDataOk() (*LoginReqData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *LoginReq) SetData(v LoginReqData) {
	o.Data = v
}

func (o LoginReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoginReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *LoginReq) UnmarshalJSON(data []byte) (err error) {
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

	varLoginReq := _LoginReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLoginReq)

	if err != nil {
		return err
	}

	*o = LoginReq(varLoginReq)

	return err
}

type NullableLoginReq struct {
	value *LoginReq
	isSet bool
}

func (v NullableLoginReq) Get() *LoginReq {
	return v.value
}

func (v *NullableLoginReq) Set(val *LoginReq) {
	v.value = val
	v.isSet = true
}

func (v NullableLoginReq) IsSet() bool {
	return v.isSet
}

func (v *NullableLoginReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoginReq(val *LoginReq) *NullableLoginReq {
	return &NullableLoginReq{value: val, isSet: true}
}

func (v NullableLoginReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoginReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


