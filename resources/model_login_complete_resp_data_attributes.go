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

// checks if the LoginCompleteRespDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoginCompleteRespDataAttributes{}

// LoginCompleteRespDataAttributes struct for LoginCompleteRespDataAttributes
type LoginCompleteRespDataAttributes struct {
	// Access Token для авторизации запросов.
	AccessToken string `json:"access_token"`
	// Refresh Token для обновления Access Token.
	RefreshToken string `json:"refresh_token"`
	// Время жизни Access Token в секундах.
	ExpiresIn int32 `json:"expires_in"`
}

type _LoginCompleteRespDataAttributes LoginCompleteRespDataAttributes

// NewLoginCompleteRespDataAttributes instantiates a new LoginCompleteRespDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoginCompleteRespDataAttributes(accessToken string, refreshToken string, expiresIn int32) *LoginCompleteRespDataAttributes {
	this := LoginCompleteRespDataAttributes{}
	this.AccessToken = accessToken
	this.RefreshToken = refreshToken
	this.ExpiresIn = expiresIn
	return &this
}

// NewLoginCompleteRespDataAttributesWithDefaults instantiates a new LoginCompleteRespDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoginCompleteRespDataAttributesWithDefaults() *LoginCompleteRespDataAttributes {
	this := LoginCompleteRespDataAttributes{}
	return &this
}

// GetAccessToken returns the AccessToken field value
func (o *LoginCompleteRespDataAttributes) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *LoginCompleteRespDataAttributes) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *LoginCompleteRespDataAttributes) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetRefreshToken returns the RefreshToken field value
func (o *LoginCompleteRespDataAttributes) GetRefreshToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RefreshToken
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value
// and a boolean to check if the value has been set.
func (o *LoginCompleteRespDataAttributes) GetRefreshTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RefreshToken, true
}

// SetRefreshToken sets field value
func (o *LoginCompleteRespDataAttributes) SetRefreshToken(v string) {
	o.RefreshToken = v
}

// GetExpiresIn returns the ExpiresIn field value
func (o *LoginCompleteRespDataAttributes) GetExpiresIn() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExpiresIn
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value
// and a boolean to check if the value has been set.
func (o *LoginCompleteRespDataAttributes) GetExpiresInOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresIn, true
}

// SetExpiresIn sets field value
func (o *LoginCompleteRespDataAttributes) SetExpiresIn(v int32) {
	o.ExpiresIn = v
}

func (o LoginCompleteRespDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoginCompleteRespDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["access_token"] = o.AccessToken
	toSerialize["refresh_token"] = o.RefreshToken
	toSerialize["expires_in"] = o.ExpiresIn
	return toSerialize, nil
}

func (o *LoginCompleteRespDataAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"access_token",
		"refresh_token",
		"expires_in",
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

	varLoginCompleteRespDataAttributes := _LoginCompleteRespDataAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLoginCompleteRespDataAttributes)

	if err != nil {
		return err
	}

	*o = LoginCompleteRespDataAttributes(varLoginCompleteRespDataAttributes)

	return err
}

type NullableLoginCompleteRespDataAttributes struct {
	value *LoginCompleteRespDataAttributes
	isSet bool
}

func (v NullableLoginCompleteRespDataAttributes) Get() *LoginCompleteRespDataAttributes {
	return v.value
}

func (v *NullableLoginCompleteRespDataAttributes) Set(val *LoginCompleteRespDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableLoginCompleteRespDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableLoginCompleteRespDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoginCompleteRespDataAttributes(val *LoginCompleteRespDataAttributes) *NullableLoginCompleteRespDataAttributes {
	return &NullableLoginCompleteRespDataAttributes{value: val, isSet: true}
}

func (v NullableLoginCompleteRespDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoginCompleteRespDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


