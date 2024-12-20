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

// checks if the TerminateSessionsData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TerminateSessionsData{}

// TerminateSessionsData struct for TerminateSessionsData
type TerminateSessionsData struct {
	Type string `json:"type"`
	Attributes TerminateSessionsDataAttributes `json:"attributes"`
}

type _TerminateSessionsData TerminateSessionsData

// NewTerminateSessionsData instantiates a new TerminateSessionsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerminateSessionsData(type_ string, attributes TerminateSessionsDataAttributes) *TerminateSessionsData {
	this := TerminateSessionsData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewTerminateSessionsDataWithDefaults instantiates a new TerminateSessionsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerminateSessionsDataWithDefaults() *TerminateSessionsData {
	this := TerminateSessionsData{}
	return &this
}

// GetType returns the Type field value
func (o *TerminateSessionsData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TerminateSessionsData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TerminateSessionsData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *TerminateSessionsData) GetAttributes() TerminateSessionsDataAttributes {
	if o == nil {
		var ret TerminateSessionsDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *TerminateSessionsData) GetAttributesOk() (*TerminateSessionsDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *TerminateSessionsData) SetAttributes(v TerminateSessionsDataAttributes) {
	o.Attributes = v
}

func (o TerminateSessionsData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TerminateSessionsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	return toSerialize, nil
}

func (o *TerminateSessionsData) UnmarshalJSON(data []byte) (err error) {
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

	varTerminateSessionsData := _TerminateSessionsData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTerminateSessionsData)

	if err != nil {
		return err
	}

	*o = TerminateSessionsData(varTerminateSessionsData)

	return err
}

type NullableTerminateSessionsData struct {
	value *TerminateSessionsData
	isSet bool
}

func (v NullableTerminateSessionsData) Get() *TerminateSessionsData {
	return v.value
}

func (v *NullableTerminateSessionsData) Set(val *TerminateSessionsData) {
	v.value = val
	v.isSet = true
}

func (v NullableTerminateSessionsData) IsSet() bool {
	return v.isSet
}

func (v *NullableTerminateSessionsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerminateSessionsData(val *TerminateSessionsData) *NullableTerminateSessionsData {
	return &NullableTerminateSessionsData{value: val, isSet: true}
}

func (v NullableTerminateSessionsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerminateSessionsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


