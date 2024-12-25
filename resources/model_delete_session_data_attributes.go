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

// checks if the DeleteSessionDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteSessionDataAttributes{}

// DeleteSessionDataAttributes struct for DeleteSessionDataAttributes
type DeleteSessionDataAttributes struct {
	// Device ID.
	SessionId string `json:"session_id"`
}

type _DeleteSessionDataAttributes DeleteSessionDataAttributes

// NewDeleteSessionDataAttributes instantiates a new DeleteSessionDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteSessionDataAttributes(sessionId string) *DeleteSessionDataAttributes {
	this := DeleteSessionDataAttributes{}
	this.SessionId = sessionId
	return &this
}

// NewDeleteSessionDataAttributesWithDefaults instantiates a new DeleteSessionDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteSessionDataAttributesWithDefaults() *DeleteSessionDataAttributes {
	this := DeleteSessionDataAttributes{}
	return &this
}

// GetSessionId returns the SessionId field value
func (o *DeleteSessionDataAttributes) GetSessionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value
// and a boolean to check if the value has been set.
func (o *DeleteSessionDataAttributes) GetSessionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionId, true
}

// SetSessionId sets field value
func (o *DeleteSessionDataAttributes) SetSessionId(v string) {
	o.SessionId = v
}

func (o DeleteSessionDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteSessionDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["session_id"] = o.SessionId
	return toSerialize, nil
}

func (o *DeleteSessionDataAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"session_id",
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

	varDeleteSessionDataAttributes := _DeleteSessionDataAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeleteSessionDataAttributes)

	if err != nil {
		return err
	}

	*o = DeleteSessionDataAttributes(varDeleteSessionDataAttributes)

	return err
}

type NullableDeleteSessionDataAttributes struct {
	value *DeleteSessionDataAttributes
	isSet bool
}

func (v NullableDeleteSessionDataAttributes) Get() *DeleteSessionDataAttributes {
	return v.value
}

func (v *NullableDeleteSessionDataAttributes) Set(val *DeleteSessionDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteSessionDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteSessionDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteSessionDataAttributes(val *DeleteSessionDataAttributes) *NullableDeleteSessionDataAttributes {
	return &NullableDeleteSessionDataAttributes{value: val, isSet: true}
}

func (v NullableDeleteSessionDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteSessionDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


