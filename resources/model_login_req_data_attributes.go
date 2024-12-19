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

// checks if the LoginReqDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoginReqDataAttributes{}

// LoginReqDataAttributes struct for LoginReqDataAttributes
type LoginReqDataAttributes struct {
	// User email (required if username is not provided).
	Email *string `json:"email,omitempty"`
	// User username (required if email is not provided).
	Username *string `json:"username,omitempty"`
	// User password.
	Password string `json:"password"`
	// Unique identifier for the user's device.
	FactoryId string `json:"factory_id"`
	// Human-readable name for the user's device (e.g., 'iPhone 14').
	DeviceName string `json:"device_name"`
	// Operating system version of the user's device.
	OsVersion string `json:"os_version"`
}

type _LoginReqDataAttributes LoginReqDataAttributes

// NewLoginReqDataAttributes instantiates a new LoginReqDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoginReqDataAttributes(password string, factoryId string, deviceName string, osVersion string) *LoginReqDataAttributes {
	this := LoginReqDataAttributes{}
	this.Password = password
	this.FactoryId = factoryId
	this.DeviceName = deviceName
	this.OsVersion = osVersion
	return &this
}

// NewLoginReqDataAttributesWithDefaults instantiates a new LoginReqDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoginReqDataAttributesWithDefaults() *LoginReqDataAttributes {
	this := LoginReqDataAttributes{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *LoginReqDataAttributes) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginReqDataAttributes) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *LoginReqDataAttributes) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *LoginReqDataAttributes) SetEmail(v string) {
	o.Email = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *LoginReqDataAttributes) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginReqDataAttributes) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *LoginReqDataAttributes) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *LoginReqDataAttributes) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value
func (o *LoginReqDataAttributes) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *LoginReqDataAttributes) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *LoginReqDataAttributes) SetPassword(v string) {
	o.Password = v
}

// GetFactoryId returns the FactoryId field value
func (o *LoginReqDataAttributes) GetFactoryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FactoryId
}

// GetFactoryIdOk returns a tuple with the FactoryId field value
// and a boolean to check if the value has been set.
func (o *LoginReqDataAttributes) GetFactoryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FactoryId, true
}

// SetFactoryId sets field value
func (o *LoginReqDataAttributes) SetFactoryId(v string) {
	o.FactoryId = v
}

// GetDeviceName returns the DeviceName field value
func (o *LoginReqDataAttributes) GetDeviceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeviceName
}

// GetDeviceNameOk returns a tuple with the DeviceName field value
// and a boolean to check if the value has been set.
func (o *LoginReqDataAttributes) GetDeviceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceName, true
}

// SetDeviceName sets field value
func (o *LoginReqDataAttributes) SetDeviceName(v string) {
	o.DeviceName = v
}

// GetOsVersion returns the OsVersion field value
func (o *LoginReqDataAttributes) GetOsVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OsVersion
}

// GetOsVersionOk returns a tuple with the OsVersion field value
// and a boolean to check if the value has been set.
func (o *LoginReqDataAttributes) GetOsVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OsVersion, true
}

// SetOsVersion sets field value
func (o *LoginReqDataAttributes) SetOsVersion(v string) {
	o.OsVersion = v
}

func (o LoginReqDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoginReqDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	toSerialize["password"] = o.Password
	toSerialize["factory_id"] = o.FactoryId
	toSerialize["device_name"] = o.DeviceName
	toSerialize["os_version"] = o.OsVersion
	return toSerialize, nil
}

func (o *LoginReqDataAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"password",
		"factory_id",
		"device_name",
		"os_version",
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

	varLoginReqDataAttributes := _LoginReqDataAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLoginReqDataAttributes)

	if err != nil {
		return err
	}

	*o = LoginReqDataAttributes(varLoginReqDataAttributes)

	return err
}

type NullableLoginReqDataAttributes struct {
	value *LoginReqDataAttributes
	isSet bool
}

func (v NullableLoginReqDataAttributes) Get() *LoginReqDataAttributes {
	return v.value
}

func (v *NullableLoginReqDataAttributes) Set(val *LoginReqDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableLoginReqDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableLoginReqDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoginReqDataAttributes(val *LoginReqDataAttributes) *NullableLoginReqDataAttributes {
	return &NullableLoginReqDataAttributes{value: val, isSet: true}
}

func (v NullableLoginReqDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoginReqDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


