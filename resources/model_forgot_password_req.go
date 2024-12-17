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

// checks if the ForgotPasswordReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ForgotPasswordReq{}

// ForgotPasswordReq struct for ForgotPasswordReq
type ForgotPasswordReq struct {
	Data ForgotPasswordReqData `json:"data"`
}

type _ForgotPasswordReq ForgotPasswordReq

// NewForgotPasswordReq instantiates a new ForgotPasswordReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewForgotPasswordReq(data ForgotPasswordReqData) *ForgotPasswordReq {
	this := ForgotPasswordReq{}
	this.Data = data
	return &this
}

// NewForgotPasswordReqWithDefaults instantiates a new ForgotPasswordReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewForgotPasswordReqWithDefaults() *ForgotPasswordReq {
	this := ForgotPasswordReq{}
	return &this
}

// GetData returns the Data field value
func (o *ForgotPasswordReq) GetData() ForgotPasswordReqData {
	if o == nil {
		var ret ForgotPasswordReqData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ForgotPasswordReq) GetDataOk() (*ForgotPasswordReqData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ForgotPasswordReq) SetData(v ForgotPasswordReqData) {
	o.Data = v
}

func (o ForgotPasswordReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ForgotPasswordReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *ForgotPasswordReq) UnmarshalJSON(data []byte) (err error) {
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

	varForgotPasswordReq := _ForgotPasswordReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varForgotPasswordReq)

	if err != nil {
		return err
	}

	*o = ForgotPasswordReq(varForgotPasswordReq)

	return err
}

type NullableForgotPasswordReq struct {
	value *ForgotPasswordReq
	isSet bool
}

func (v NullableForgotPasswordReq) Get() *ForgotPasswordReq {
	return v.value
}

func (v *NullableForgotPasswordReq) Set(val *ForgotPasswordReq) {
	v.value = val
	v.isSet = true
}

func (v NullableForgotPasswordReq) IsSet() bool {
	return v.isSet
}

func (v *NullableForgotPasswordReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForgotPasswordReq(val *ForgotPasswordReq) *NullableForgotPasswordReq {
	return &NullableForgotPasswordReq{value: val, isSet: true}
}

func (v NullableForgotPasswordReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForgotPasswordReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

