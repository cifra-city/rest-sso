# LoginCompleteReqDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | User email | [optional] 
**Username** | Pointer to **string** | User username | [optional] 
**DeviceName** | **string** | Device name | 

## Methods

### NewLoginCompleteReqDataAttributes

`func NewLoginCompleteReqDataAttributes(deviceName string, ) *LoginCompleteReqDataAttributes`

NewLoginCompleteReqDataAttributes instantiates a new LoginCompleteReqDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginCompleteReqDataAttributesWithDefaults

`func NewLoginCompleteReqDataAttributesWithDefaults() *LoginCompleteReqDataAttributes`

NewLoginCompleteReqDataAttributesWithDefaults instantiates a new LoginCompleteReqDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *LoginCompleteReqDataAttributes) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *LoginCompleteReqDataAttributes) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *LoginCompleteReqDataAttributes) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *LoginCompleteReqDataAttributes) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetUsername

`func (o *LoginCompleteReqDataAttributes) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *LoginCompleteReqDataAttributes) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *LoginCompleteReqDataAttributes) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *LoginCompleteReqDataAttributes) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetDeviceName

`func (o *LoginCompleteReqDataAttributes) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *LoginCompleteReqDataAttributes) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *LoginCompleteReqDataAttributes) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


