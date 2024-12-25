# RefreshTokenDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** | Access Token для авторизации запросов. | 
**RefreshToken** | **string** | Refresh Token для обновления Access Token. | 
**ExpiresIn** | **int32** | Время жизни Access Token в секундах. | 

## Methods

### NewRefreshTokenDataAttributes

`func NewRefreshTokenDataAttributes(accessToken string, refreshToken string, expiresIn int32, ) *RefreshTokenDataAttributes`

NewRefreshTokenDataAttributes instantiates a new RefreshTokenDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefreshTokenDataAttributesWithDefaults

`func NewRefreshTokenDataAttributesWithDefaults() *RefreshTokenDataAttributes`

NewRefreshTokenDataAttributesWithDefaults instantiates a new RefreshTokenDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *RefreshTokenDataAttributes) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *RefreshTokenDataAttributes) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *RefreshTokenDataAttributes) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetRefreshToken

`func (o *RefreshTokenDataAttributes) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *RefreshTokenDataAttributes) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *RefreshTokenDataAttributes) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.


### GetExpiresIn

`func (o *RefreshTokenDataAttributes) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *RefreshTokenDataAttributes) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *RefreshTokenDataAttributes) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


