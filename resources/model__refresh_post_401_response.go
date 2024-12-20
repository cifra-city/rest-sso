/*
Cifra SSO REST API

SSO REST API for Cifra app

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package resources

import (
	"encoding/json"
)

// checks if the RefreshPost401Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RefreshPost401Response{}

// RefreshPost401Response struct for RefreshPost401Response
type RefreshPost401Response struct {
	Error *string `json:"error,omitempty"`
}

// NewRefreshPost401Response instantiates a new RefreshPost401Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefreshPost401Response() *RefreshPost401Response {
	this := RefreshPost401Response{}
	return &this
}

// NewRefreshPost401ResponseWithDefaults instantiates a new RefreshPost401Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefreshPost401ResponseWithDefaults() *RefreshPost401Response {
	this := RefreshPost401Response{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *RefreshPost401Response) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshPost401Response) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *RefreshPost401Response) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *RefreshPost401Response) SetError(v string) {
	o.Error = &v
}

func (o RefreshPost401Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RefreshPost401Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableRefreshPost401Response struct {
	value *RefreshPost401Response
	isSet bool
}

func (v NullableRefreshPost401Response) Get() *RefreshPost401Response {
	return v.value
}

func (v *NullableRefreshPost401Response) Set(val *RefreshPost401Response) {
	v.value = val
	v.isSet = true
}

func (v NullableRefreshPost401Response) IsSet() bool {
	return v.isSet
}

func (v *NullableRefreshPost401Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefreshPost401Response(val *RefreshPost401Response) *NullableRefreshPost401Response {
	return &NullableRefreshPost401Response{value: val, isSet: true}
}

func (v NullableRefreshPost401Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefreshPost401Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


