# ResetPasswordCompleteDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | User email (required if username is not provided). | [optional] 
**Username** | Pointer to **string** | User username (required if email is not provided). | [optional] 
**FirstPassword** | **string** | User password. | 
**SecondPassword** | **string** | User password. | 
**FactoryId** | **string** | Unique identifier for the user&#39;s device. | 
**DeviceName** | **string** | Human-readable name for the user&#39;s device (e.g., &#39;iPhone 14&#39;). | 
**OsVersion** | **string** | Operating system version of the user&#39;s device. | 
**IpAddress** | **string** | IP address of the user&#39;s device. | 

## Methods

### NewResetPasswordCompleteDataAttributes

`func NewResetPasswordCompleteDataAttributes(firstPassword string, secondPassword string, factoryId string, deviceName string, osVersion string, ipAddress string, ) *ResetPasswordCompleteDataAttributes`

NewResetPasswordCompleteDataAttributes instantiates a new ResetPasswordCompleteDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResetPasswordCompleteDataAttributesWithDefaults

`func NewResetPasswordCompleteDataAttributesWithDefaults() *ResetPasswordCompleteDataAttributes`

NewResetPasswordCompleteDataAttributesWithDefaults instantiates a new ResetPasswordCompleteDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *ResetPasswordCompleteDataAttributes) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ResetPasswordCompleteDataAttributes) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ResetPasswordCompleteDataAttributes) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ResetPasswordCompleteDataAttributes) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetUsername

`func (o *ResetPasswordCompleteDataAttributes) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ResetPasswordCompleteDataAttributes) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ResetPasswordCompleteDataAttributes) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ResetPasswordCompleteDataAttributes) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetFirstPassword

`func (o *ResetPasswordCompleteDataAttributes) GetFirstPassword() string`

GetFirstPassword returns the FirstPassword field if non-nil, zero value otherwise.

### GetFirstPasswordOk

`func (o *ResetPasswordCompleteDataAttributes) GetFirstPasswordOk() (*string, bool)`

GetFirstPasswordOk returns a tuple with the FirstPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPassword

`func (o *ResetPasswordCompleteDataAttributes) SetFirstPassword(v string)`

SetFirstPassword sets FirstPassword field to given value.


### GetSecondPassword

`func (o *ResetPasswordCompleteDataAttributes) GetSecondPassword() string`

GetSecondPassword returns the SecondPassword field if non-nil, zero value otherwise.

### GetSecondPasswordOk

`func (o *ResetPasswordCompleteDataAttributes) GetSecondPasswordOk() (*string, bool)`

GetSecondPasswordOk returns a tuple with the SecondPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondPassword

`func (o *ResetPasswordCompleteDataAttributes) SetSecondPassword(v string)`

SetSecondPassword sets SecondPassword field to given value.


### GetFactoryId

`func (o *ResetPasswordCompleteDataAttributes) GetFactoryId() string`

GetFactoryId returns the FactoryId field if non-nil, zero value otherwise.

### GetFactoryIdOk

`func (o *ResetPasswordCompleteDataAttributes) GetFactoryIdOk() (*string, bool)`

GetFactoryIdOk returns a tuple with the FactoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactoryId

`func (o *ResetPasswordCompleteDataAttributes) SetFactoryId(v string)`

SetFactoryId sets FactoryId field to given value.


### GetDeviceName

`func (o *ResetPasswordCompleteDataAttributes) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *ResetPasswordCompleteDataAttributes) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *ResetPasswordCompleteDataAttributes) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.


### GetOsVersion

`func (o *ResetPasswordCompleteDataAttributes) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *ResetPasswordCompleteDataAttributes) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *ResetPasswordCompleteDataAttributes) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.


### GetIpAddress

`func (o *ResetPasswordCompleteDataAttributes) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *ResetPasswordCompleteDataAttributes) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *ResetPasswordCompleteDataAttributes) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


