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

// checks if the DeleteSessionData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteSessionData{}

// DeleteSessionData struct for DeleteSessionData
type DeleteSessionData struct {
	Type string `json:"type"`
	Attributes DeleteSessionDataAttributes `json:"attributes"`
}

type _DeleteSessionData DeleteSessionData

// NewDeleteSessionData instantiates a new DeleteSessionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteSessionData(type_ string, attributes DeleteSessionDataAttributes) *DeleteSessionData {
	this := DeleteSessionData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewDeleteSessionDataWithDefaults instantiates a new DeleteSessionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteSessionDataWithDefaults() *DeleteSessionData {
	this := DeleteSessionData{}
	return &this
}

// GetType returns the Type field value
func (o *DeleteSessionData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DeleteSessionData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DeleteSessionData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *DeleteSessionData) GetAttributes() DeleteSessionDataAttributes {
	if o == nil {
		var ret DeleteSessionDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *DeleteSessionData) GetAttributesOk() (*DeleteSessionDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *DeleteSessionData) SetAttributes(v DeleteSessionDataAttributes) {
	o.Attributes = v
}

func (o DeleteSessionData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteSessionData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	return toSerialize, nil
}

func (o *DeleteSessionData) UnmarshalJSON(data []byte) (err error) {
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

	varDeleteSessionData := _DeleteSessionData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeleteSessionData)

	if err != nil {
		return err
	}

	*o = DeleteSessionData(varDeleteSessionData)

	return err
}

type NullableDeleteSessionData struct {
	value *DeleteSessionData
	isSet bool
}

func (v NullableDeleteSessionData) Get() *DeleteSessionData {
	return v.value
}

func (v *NullableDeleteSessionData) Set(val *DeleteSessionData) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteSessionData) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteSessionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteSessionData(val *DeleteSessionData) *NullableDeleteSessionData {
	return &NullableDeleteSessionData{value: val, isSet: true}
}

func (v NullableDeleteSessionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteSessionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

