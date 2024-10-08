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
)

// checks if the WebhookUpdateCustomHeaderRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookUpdateCustomHeaderRequest{}

// WebhookUpdateCustomHeaderRequest struct for WebhookUpdateCustomHeaderRequest
type WebhookUpdateCustomHeaderRequest struct {
	// Additional custom information included in the HTTP header. For example, this might contain security information, along with Basic Authentication, when making a JIT Funding request.
	CustomHeader *map[string]string `json:"custom_header,omitempty"`
}

// NewWebhookUpdateCustomHeaderRequest instantiates a new WebhookUpdateCustomHeaderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookUpdateCustomHeaderRequest() *WebhookUpdateCustomHeaderRequest {
	this := WebhookUpdateCustomHeaderRequest{}
	return &this
}

// NewWebhookUpdateCustomHeaderRequestWithDefaults instantiates a new WebhookUpdateCustomHeaderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookUpdateCustomHeaderRequestWithDefaults() *WebhookUpdateCustomHeaderRequest {
	this := WebhookUpdateCustomHeaderRequest{}
	return &this
}

// GetCustomHeader returns the CustomHeader field value if set, zero value otherwise.
func (o *WebhookUpdateCustomHeaderRequest) GetCustomHeader() map[string]string {
	if o == nil || IsNil(o.CustomHeader) {
		var ret map[string]string
		return ret
	}
	return *o.CustomHeader
}

// GetCustomHeaderOk returns a tuple with the CustomHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookUpdateCustomHeaderRequest) GetCustomHeaderOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.CustomHeader) {
		return nil, false
	}
	return o.CustomHeader, true
}

// HasCustomHeader returns a boolean if a field has been set.
func (o *WebhookUpdateCustomHeaderRequest) HasCustomHeader() bool {
	if o != nil && !IsNil(o.CustomHeader) {
		return true
	}

	return false
}

// SetCustomHeader gets a reference to the given map[string]string and assigns it to the CustomHeader field.
func (o *WebhookUpdateCustomHeaderRequest) SetCustomHeader(v map[string]string) {
	o.CustomHeader = &v
}

func (o WebhookUpdateCustomHeaderRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookUpdateCustomHeaderRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomHeader) {
		toSerialize["custom_header"] = o.CustomHeader
	}
	return toSerialize, nil
}

type NullableWebhookUpdateCustomHeaderRequest struct {
	value *WebhookUpdateCustomHeaderRequest
	isSet bool
}

func (v NullableWebhookUpdateCustomHeaderRequest) Get() *WebhookUpdateCustomHeaderRequest {
	return v.value
}

func (v *NullableWebhookUpdateCustomHeaderRequest) Set(val *WebhookUpdateCustomHeaderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookUpdateCustomHeaderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookUpdateCustomHeaderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookUpdateCustomHeaderRequest(val *WebhookUpdateCustomHeaderRequest) *NullableWebhookUpdateCustomHeaderRequest {
	return &NullableWebhookUpdateCustomHeaderRequest{value: val, isSet: true}
}

func (v NullableWebhookUpdateCustomHeaderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookUpdateCustomHeaderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


