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

// checks if the LoginRespData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoginRespData{}

// LoginRespData struct for LoginRespData
type LoginRespData struct {
	Type string `json:"type"`
	// User ID (UUID)
	Id string `json:"id"`
	Attributes LoginRespDataAttributes `json:"attributes"`
}

type _LoginRespData LoginRespData

// NewLoginRespData instantiates a new LoginRespData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoginRespData(type_ string, id string, attributes LoginRespDataAttributes) *LoginRespData {
	this := LoginRespData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewLoginRespDataWithDefaults instantiates a new LoginRespData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoginRespDataWithDefaults() *LoginRespData {
	this := LoginRespData{}
	return &this
}

// GetType returns the Type field value
func (o *LoginRespData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LoginRespData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LoginRespData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *LoginRespData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LoginRespData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LoginRespData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *LoginRespData) GetAttributes() LoginRespDataAttributes {
	if o == nil {
		var ret LoginRespDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *LoginRespData) GetAttributesOk() (*LoginRespDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *LoginRespData) SetAttributes(v LoginRespDataAttributes) {
	o.Attributes = v
}

func (o LoginRespData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoginRespData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["attributes"] = o.Attributes
	return toSerialize, nil
}

func (o *LoginRespData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
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

	varLoginRespData := _LoginRespData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLoginRespData)

	if err != nil {
		return err
	}

	*o = LoginRespData(varLoginRespData)

	return err
}

type NullableLoginRespData struct {
	value *LoginRespData
	isSet bool
}

func (v NullableLoginRespData) Get() *LoginRespData {
	return v.value
}

func (v *NullableLoginRespData) Set(val *LoginRespData) {
	v.value = val
	v.isSet = true
}

func (v NullableLoginRespData) IsSet() bool {
	return v.isSet
}

func (v *NullableLoginRespData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoginRespData(val *LoginRespData) *NullableLoginRespData {
	return &NullableLoginRespData{value: val, isSet: true}
}

func (v NullableLoginRespData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoginRespData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


