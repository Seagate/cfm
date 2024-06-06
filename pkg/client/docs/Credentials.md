# Credentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | The username credentials used to communicate with another service | 
**Password** | **string** | The password credentials used to communicate with another service | 
**IpAddress** | **string** | The IP Address in dot notation of the service | 
**Port** | **int32** |  | 
**Insecure** | Pointer to **bool** | Set to true if an insecure connection should be used | [optional] [default to false]
**Protocol** | Pointer to **string** | Examples of http vs https | [optional] [default to "https"]
**Backend** | Pointer to **string** | Examples of backend | [optional] [default to "httpfish"]
**CustomId** | Pointer to **string** | A user-defined string to uniquely identify an individual endpoint device. | [optional] 

## Methods

### NewCredentials

`func NewCredentials(username string, password string, ipAddress string, port int32, ) *Credentials`

NewCredentials instantiates a new Credentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialsWithDefaults

`func NewCredentialsWithDefaults() *Credentials`

NewCredentialsWithDefaults instantiates a new Credentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *Credentials) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Credentials) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Credentials) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *Credentials) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Credentials) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Credentials) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetIpAddress

`func (o *Credentials) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *Credentials) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *Credentials) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.


### GetPort

`func (o *Credentials) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *Credentials) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *Credentials) SetPort(v int32)`

SetPort sets Port field to given value.


### GetInsecure

`func (o *Credentials) GetInsecure() bool`

GetInsecure returns the Insecure field if non-nil, zero value otherwise.

### GetInsecureOk

`func (o *Credentials) GetInsecureOk() (*bool, bool)`

GetInsecureOk returns a tuple with the Insecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecure

`func (o *Credentials) SetInsecure(v bool)`

SetInsecure sets Insecure field to given value.

### HasInsecure

`func (o *Credentials) HasInsecure() bool`

HasInsecure returns a boolean if a field has been set.

### GetProtocol

`func (o *Credentials) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *Credentials) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *Credentials) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *Credentials) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetBackend

`func (o *Credentials) GetBackend() string`

GetBackend returns the Backend field if non-nil, zero value otherwise.

### GetBackendOk

`func (o *Credentials) GetBackendOk() (*string, bool)`

GetBackendOk returns a tuple with the Backend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackend

`func (o *Credentials) SetBackend(v string)`

SetBackend sets Backend field to given value.

### HasBackend

`func (o *Credentials) HasBackend() bool`

HasBackend returns a boolean if a field has been set.

### GetCustomId

`func (o *Credentials) GetCustomId() string`

GetCustomId returns the CustomId field if non-nil, zero value otherwise.

### GetCustomIdOk

`func (o *Credentials) GetCustomIdOk() (*string, bool)`

GetCustomIdOk returns a tuple with the CustomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomId

`func (o *Credentials) SetCustomId(v string)`

SetCustomId sets CustomId field to given value.

### HasCustomId

`func (o *Credentials) HasCustomId() bool`

HasCustomId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


