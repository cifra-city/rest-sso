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

// checks if the RefreshReqData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RefreshReqData{}

// RefreshReqData struct for RefreshReqData
type RefreshReqData struct {
	Type string `json:"type"`
	Attributes RefreshReqDataAttributes `json:"attributes"`
}

type _RefreshReqData RefreshReqData

// NewRefreshReqData instantiates a new RefreshReqData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefreshReqData(type_ string, attributes RefreshReqDataAttributes) *RefreshReqData {
	this := RefreshReqData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewRefreshReqDataWithDefaults instantiates a new RefreshReqData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefreshReqDataWithDefaults() *RefreshReqData {
	this := RefreshReqData{}
	return &this
}

// GetType returns the Type field value
func (o *RefreshReqData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RefreshReqData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RefreshReqData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *RefreshReqData) GetAttributes() RefreshReqDataAttributes {
	if o == nil {
		var ret RefreshReqDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *RefreshReqData) GetAttributesOk() (*RefreshReqDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *RefreshReqData) SetAttributes(v RefreshReqDataAttributes) {
	o.Attributes = v
}

func (o RefreshReqData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RefreshReqData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	return toSerialize, nil
}

func (o *RefreshReqData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"attributes",
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

	varRefreshReqData := _RefreshReqData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRefreshReqData)

	if err != nil {
		return err
	}

	*o = RefreshReqData(varRefreshReqData)

	return err
}

type NullableRefreshReqData struct {
	value *RefreshReqData
	isSet bool
}

func (v NullableRefreshReqData) Get() *RefreshReqData {
	return v.value
}

func (v *NullableRefreshReqData) Set(val *RefreshReqData) {
	v.value = val
	v.isSet = true
}

func (v NullableRefreshReqData) IsSet() bool {
	return v.isSet
}

func (v *NullableRefreshReqData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefreshReqData(val *RefreshReqData) *NullableRefreshReqData {
	return &NullableRefreshReqData{value: val, isSet: true}
}

func (v NullableRefreshReqData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefreshReqData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


