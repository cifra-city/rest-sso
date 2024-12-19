# ResetPasswordInitiateDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | User email (required if username is not provided). | [optional] 
**Username** | Pointer to **string** | User username (required if email is not provided). | [optional] 
**DeviceName** | **string** | Human-readable name for the user&#39;s device (e.g., &#39;iPhone 14&#39;). | 
**OsVersion** | **string** | Operating system version of the user&#39;s device. | 

## Methods

### NewResetPasswordInitiateDataAttributes

`func NewResetPasswordInitiateDataAttributes(deviceName string, osVersion string, ) *ResetPasswordInitiateDataAttributes`

NewResetPasswordInitiateDataAttributes instantiates a new ResetPasswordInitiateDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResetPasswordInitiateDataAttributesWithDefaults

`func NewResetPasswordInitiateDataAttributesWithDefaults() *ResetPasswordInitiateDataAttributes`

NewResetPasswordInitiateDataAttributesWithDefaults instantiates a new ResetPasswordInitiateDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *ResetPasswordInitiateDataAttributes) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ResetPasswordInitiateDataAttributes) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ResetPasswordInitiateDataAttributes) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ResetPasswordInitiateDataAttributes) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetUsername

`func (o *ResetPasswordInitiateDataAttributes) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ResetPasswordInitiateDataAttributes) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ResetPasswordInitiateDataAttributes) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ResetPasswordInitiateDataAttributes) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetDeviceName

`func (o *ResetPasswordInitiateDataAttributes) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *ResetPasswordInitiateDataAttributes) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *ResetPasswordInitiateDataAttributes) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.


### GetOsVersion

`func (o *ResetPasswordInitiateDataAttributes) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *ResetPasswordInitiateDataAttributes) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *ResetPasswordInitiateDataAttributes) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


