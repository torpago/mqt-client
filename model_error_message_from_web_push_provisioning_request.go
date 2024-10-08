/*
Core API

Marqeta's Core API endpoints, conveniently annotated to enable code generation (including SDKs), test cases, and documentation. Currently in beta.

API version: 3.0.19
Contact: support@marqeta.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ErrorMessageFromWebPushProvisioningRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorMessageFromWebPushProvisioningRequest{}

// ErrorMessageFromWebPushProvisioningRequest struct for ErrorMessageFromWebPushProvisioningRequest
type ErrorMessageFromWebPushProvisioningRequest struct {
	// Code for the error that occurred.
	ErrorCode string `json:"error_code"`
	// Descriptive error message.
	ErrorMessage string `json:"error_message"`
}

type _ErrorMessageFromWebPushProvisioningRequest ErrorMessageFromWebPushProvisioningRequest

// NewErrorMessageFromWebPushProvisioningRequest instantiates a new ErrorMessageFromWebPushProvisioningRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorMessageFromWebPushProvisioningRequest(errorCode string, errorMessage string) *ErrorMessageFromWebPushProvisioningRequest {
	this := ErrorMessageFromWebPushProvisioningRequest{}
	this.ErrorCode = errorCode
	this.ErrorMessage = errorMessage
	return &this
}

// NewErrorMessageFromWebPushProvisioningRequestWithDefaults instantiates a new ErrorMessageFromWebPushProvisioningRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorMessageFromWebPushProvisioningRequestWithDefaults() *ErrorMessageFromWebPushProvisioningRequest {
	this := ErrorMessageFromWebPushProvisioningRequest{}
	return &this
}

// GetErrorCode returns the ErrorCode field value
func (o *ErrorMessageFromWebPushProvisioningRequest) GetErrorCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value
// and a boolean to check if the value has been set.
func (o *ErrorMessageFromWebPushProvisioningRequest) GetErrorCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorCode, true
}

// SetErrorCode sets field value
func (o *ErrorMessageFromWebPushProvisioningRequest) SetErrorCode(v string) {
	o.ErrorCode = v
}

// GetErrorMessage returns the ErrorMessage field value
func (o *ErrorMessageFromWebPushProvisioningRequest) GetErrorMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value
// and a boolean to check if the value has been set.
func (o *ErrorMessageFromWebPushProvisioningRequest) GetErrorMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorMessage, true
}

// SetErrorMessage sets field value
func (o *ErrorMessageFromWebPushProvisioningRequest) SetErrorMessage(v string) {
	o.ErrorMessage = v
}

func (o ErrorMessageFromWebPushProvisioningRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorMessageFromWebPushProvisioningRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["error_code"] = o.ErrorCode
	toSerialize["error_message"] = o.ErrorMessage
	return toSerialize, nil
}

func (o *ErrorMessageFromWebPushProvisioningRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"error_code",
		"error_message",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varErrorMessageFromWebPushProvisioningRequest := _ErrorMessageFromWebPushProvisioningRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varErrorMessageFromWebPushProvisioningRequest)

	if err != nil {
		return err
	}

	*o = ErrorMessageFromWebPushProvisioningRequest(varErrorMessageFromWebPushProvisioningRequest)

	return err
}

type NullableErrorMessageFromWebPushProvisioningRequest struct {
	value *ErrorMessageFromWebPushProvisioningRequest
	isSet bool
}

func (v NullableErrorMessageFromWebPushProvisioningRequest) Get() *ErrorMessageFromWebPushProvisioningRequest {
	return v.value
}

func (v *NullableErrorMessageFromWebPushProvisioningRequest) Set(val *ErrorMessageFromWebPushProvisioningRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorMessageFromWebPushProvisioningRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorMessageFromWebPushProvisioningRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorMessageFromWebPushProvisioningRequest(val *ErrorMessageFromWebPushProvisioningRequest) *NullableErrorMessageFromWebPushProvisioningRequest {
	return &NullableErrorMessageFromWebPushProvisioningRequest{value: val, isSet: true}
}

func (v NullableErrorMessageFromWebPushProvisioningRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorMessageFromWebPushProvisioningRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


