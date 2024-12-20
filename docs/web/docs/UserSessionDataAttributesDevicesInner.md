# UserSessionDataAttributesDevicesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Уникальный идентификатор устройства | 
**FactoryId** | **string** | Серийный номер устройства | 
**DeviceName** | **string** | Имя устройства, заданное пользователем | 
**OsVersion** | **string** | Версия операционной системы устройства | 
**LastUsed** | **time.Time** | Дата и время последнего использования устройства | 

## Methods

### NewUserSessionDataAttributesDevicesInner

`func NewUserSessionDataAttributesDevicesInner(id string, factoryId string, deviceName string, osVersion string, lastUsed time.Time, ) *UserSessionDataAttributesDevicesInner`

NewUserSessionDataAttributesDevicesInner instantiates a new UserSessionDataAttributesDevicesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSessionDataAttributesDevicesInnerWithDefaults

`func NewUserSessionDataAttributesDevicesInnerWithDefaults() *UserSessionDataAttributesDevicesInner`

NewUserSessionDataAttributesDevicesInnerWithDefaults instantiates a new UserSessionDataAttributesDevicesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserSessionDataAttributesDevicesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserSessionDataAttributesDevicesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserSessionDataAttributesDevicesInner) SetId(v string)`

SetId sets Id field to given value.


### GetFactoryId

`func (o *UserSessionDataAttributesDevicesInner) GetFactoryId() string`

GetFactoryId returns the FactoryId field if non-nil, zero value otherwise.

### GetFactoryIdOk

`func (o *UserSessionDataAttributesDevicesInner) GetFactoryIdOk() (*string, bool)`

GetFactoryIdOk returns a tuple with the FactoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactoryId

`func (o *UserSessionDataAttributesDevicesInner) SetFactoryId(v string)`

SetFactoryId sets FactoryId field to given value.


### GetDeviceName

`func (o *UserSessionDataAttributesDevicesInner) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *UserSessionDataAttributesDevicesInner) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *UserSessionDataAttributesDevicesInner) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.


### GetOsVersion

`func (o *UserSessionDataAttributesDevicesInner) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *UserSessionDataAttributesDevicesInner) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *UserSessionDataAttributesDevicesInner) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.


### GetLastUsed

`func (o *UserSessionDataAttributesDevicesInner) GetLastUsed() time.Time`

GetLastUsed returns the LastUsed field if non-nil, zero value otherwise.

### GetLastUsedOk

`func (o *UserSessionDataAttributesDevicesInner) GetLastUsedOk() (*time.Time, bool)`

GetLastUsedOk returns a tuple with the LastUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsed

`func (o *UserSessionDataAttributesDevicesInner) SetLastUsed(v time.Time)`

SetLastUsed sets LastUsed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


