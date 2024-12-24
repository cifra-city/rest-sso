# SessionsDataAttributesDevicesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Уникальный идентификатор устройства | 
**FactoryId** | **string** | Серийный номер устройства | 
**DeviceName** | **string** | Имя устройства, заданное пользователем | 
**OsVersion** | **string** | Версия операционной системы устройства | 
**LastUsed** | **time.Time** | Дата и время последнего использования устройства | 

## Methods

### NewSessionsDataAttributesDevicesInner

`func NewSessionsDataAttributesDevicesInner(id string, factoryId string, deviceName string, osVersion string, lastUsed time.Time, ) *SessionsDataAttributesDevicesInner`

NewSessionsDataAttributesDevicesInner instantiates a new SessionsDataAttributesDevicesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionsDataAttributesDevicesInnerWithDefaults

`func NewSessionsDataAttributesDevicesInnerWithDefaults() *SessionsDataAttributesDevicesInner`

NewSessionsDataAttributesDevicesInnerWithDefaults instantiates a new SessionsDataAttributesDevicesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SessionsDataAttributesDevicesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SessionsDataAttributesDevicesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SessionsDataAttributesDevicesInner) SetId(v string)`

SetId sets Id field to given value.


### GetFactoryId

`func (o *SessionsDataAttributesDevicesInner) GetFactoryId() string`

GetFactoryId returns the FactoryId field if non-nil, zero value otherwise.

### GetFactoryIdOk

`func (o *SessionsDataAttributesDevicesInner) GetFactoryIdOk() (*string, bool)`

GetFactoryIdOk returns a tuple with the FactoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactoryId

`func (o *SessionsDataAttributesDevicesInner) SetFactoryId(v string)`

SetFactoryId sets FactoryId field to given value.


### GetDeviceName

`func (o *SessionsDataAttributesDevicesInner) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *SessionsDataAttributesDevicesInner) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *SessionsDataAttributesDevicesInner) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.


### GetOsVersion

`func (o *SessionsDataAttributesDevicesInner) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *SessionsDataAttributesDevicesInner) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *SessionsDataAttributesDevicesInner) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.


### GetLastUsed

`func (o *SessionsDataAttributesDevicesInner) GetLastUsed() time.Time`

GetLastUsed returns the LastUsed field if non-nil, zero value otherwise.

### GetLastUsedOk

`func (o *SessionsDataAttributesDevicesInner) GetLastUsedOk() (*time.Time, bool)`

GetLastUsedOk returns a tuple with the LastUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsed

`func (o *SessionsDataAttributesDevicesInner) SetLastUsed(v time.Time)`

SetLastUsed sets LastUsed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


