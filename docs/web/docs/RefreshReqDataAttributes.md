# RefreshReqDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RefreshToken** | **string** | The refresh token to generate a new access token. | 
**DeviceId** | **string** | The device ID of the device that the refresh token was issued to. (Not factory ID) | 
**FactoryId** | **string** | Unique identifier for the user&#39;s device. | 
**DeviceName** | **string** | Human-readable name for the user&#39;s device (e.g., &#39;iPhone 14&#39;). | 
**OsVersion** | **string** | Operating system version of the user&#39;s device. | 
**IpAddress** | **string** | IP address of the user&#39;s device. | 

## Methods

### NewRefreshReqDataAttributes

`func NewRefreshReqDataAttributes(refreshToken string, deviceId string, factoryId string, deviceName string, osVersion string, ipAddress string, ) *RefreshReqDataAttributes`

NewRefreshReqDataAttributes instantiates a new RefreshReqDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefreshReqDataAttributesWithDefaults

`func NewRefreshReqDataAttributesWithDefaults() *RefreshReqDataAttributes`

NewRefreshReqDataAttributesWithDefaults instantiates a new RefreshReqDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefreshToken

`func (o *RefreshReqDataAttributes) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *RefreshReqDataAttributes) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *RefreshReqDataAttributes) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.


### GetDeviceId

`func (o *RefreshReqDataAttributes) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *RefreshReqDataAttributes) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *RefreshReqDataAttributes) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.


### GetFactoryId

`func (o *RefreshReqDataAttributes) GetFactoryId() string`

GetFactoryId returns the FactoryId field if non-nil, zero value otherwise.

### GetFactoryIdOk

`func (o *RefreshReqDataAttributes) GetFactoryIdOk() (*string, bool)`

GetFactoryIdOk returns a tuple with the FactoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactoryId

`func (o *RefreshReqDataAttributes) SetFactoryId(v string)`

SetFactoryId sets FactoryId field to given value.


### GetDeviceName

`func (o *RefreshReqDataAttributes) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *RefreshReqDataAttributes) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *RefreshReqDataAttributes) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.


### GetOsVersion

`func (o *RefreshReqDataAttributes) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *RefreshReqDataAttributes) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *RefreshReqDataAttributes) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.


### GetIpAddress

`func (o *RefreshReqDataAttributes) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *RefreshReqDataAttributes) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *RefreshReqDataAttributes) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


