# BadRequestError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | **string** |  | 
**ErrorMessage** | **string** |  | 

## Methods

### NewBadRequestError

`func NewBadRequestError(errorCode string, errorMessage string, ) *BadRequestError`

NewBadRequestError instantiates a new BadRequestError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBadRequestErrorWithDefaults

`func NewBadRequestErrorWithDefaults() *BadRequestError`

NewBadRequestErrorWithDefaults instantiates a new BadRequestError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *BadRequestError) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *BadRequestError) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *BadRequestError) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.


### GetErrorMessage

`func (o *BadRequestError) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *BadRequestError) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *BadRequestError) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


