/*
Cifra SSO REST API

SSO REST API for Cifra app

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package resources

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the RefreshReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RefreshReq{}

// RefreshReq struct for RefreshReq
type RefreshReq struct {
	Data RefreshReqData `json:"data"`
}

type _RefreshReq RefreshReq

// NewRefreshReq instantiates a new RefreshReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefreshReq(data RefreshReqData) *RefreshReq {
	this := RefreshReq{}
	this.Data = data
	return &this
}

// NewRefreshReqWithDefaults instantiates a new RefreshReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefreshReqWithDefaults() *RefreshReq {
	this := RefreshReq{}
	return &this
}

// GetData returns the Data field value
func (o *RefreshReq) GetData() RefreshReqData {
	if o == nil {
		var ret RefreshReqData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *RefreshReq) GetDataOk() (*RefreshReqData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *RefreshReq) SetData(v RefreshReqData) {
	o.Data = v
}

func (o RefreshReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RefreshReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *RefreshReq) UnmarshalJSON(data []byte) (err error) {
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

	varRefreshReq := _RefreshReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRefreshReq)

	if err != nil {
		return err
	}

	*o = RefreshReq(varRefreshReq)

	return err
}

type NullableRefreshReq struct {
	value *RefreshReq
	isSet bool
}

func (v NullableRefreshReq) Get() *RefreshReq {
	return v.value
}

func (v *NullableRefreshReq) Set(val *RefreshReq) {
	v.value = val
	v.isSet = true
}

func (v NullableRefreshReq) IsSet() bool {
	return v.isSet
}

func (v *NullableRefreshReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefreshReq(val *RefreshReq) *NullableRefreshReq {
	return &NullableRefreshReq{value: val, isSet: true}
}

func (v NullableRefreshReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefreshReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


