# EmailData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Attributes** | [**EmailDataAttributes**](EmailDataAttributes.md) |  | 

## Methods

### NewEmailData

`func NewEmailData(type_ string, attributes EmailDataAttributes, ) *EmailData`

NewEmailData instantiates a new EmailData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailDataWithDefaults

`func NewEmailDataWithDefaults() *EmailData`

NewEmailDataWithDefaults instantiates a new EmailData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EmailData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EmailData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EmailData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *EmailData) GetAttributes() EmailDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *EmailData) GetAttributesOk() (*EmailDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *EmailData) SetAttributes(v EmailDataAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

