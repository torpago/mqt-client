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

// checks if the WebhookEventResendContainerResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookEventResendContainerResponse{}

// WebhookEventResendContainerResponse Contains information about a webhook event.
type WebhookEventResendContainerResponse struct {
	// Event notification that was resent to your webhook endpoint.
	Unused *string `json:"unused,omitempty"`
}

// NewWebhookEventResendContainerResponse instantiates a new WebhookEventResendContainerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookEventResendContainerResponse() *WebhookEventResendContainerResponse {
	this := WebhookEventResendContainerResponse{}
	return &this
}

// NewWebhookEventResendContainerResponseWithDefaults instantiates a new WebhookEventResendContainerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookEventResendContainerResponseWithDefaults() *WebhookEventResendContainerResponse {
	this := WebhookEventResendContainerResponse{}
	return &this
}

// GetUnused returns the Unused field value if set, zero value otherwise.
func (o *WebhookEventResendContainerResponse) GetUnused() string {
	if o == nil || IsNil(o.Unused) {
		var ret string
		return ret
	}
	return *o.Unused
}

// GetUnusedOk returns a tuple with the Unused field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookEventResendContainerResponse) GetUnusedOk() (*string, bool) {
	if o == nil || IsNil(o.Unused) {
		return nil, false
	}
	return o.Unused, true
}

// HasUnused returns a boolean if a field has been set.
func (o *WebhookEventResendContainerResponse) HasUnused() bool {
	if o != nil && !IsNil(o.Unused) {
		return true
	}

	return false
}

// SetUnused gets a reference to the given string and assigns it to the Unused field.
func (o *WebhookEventResendContainerResponse) SetUnused(v string) {
	o.Unused = &v
}

func (o WebhookEventResendContainerResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookEventResendContainerResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Unused) {
		toSerialize["unused"] = o.Unused
	}
	return toSerialize, nil
}

type NullableWebhookEventResendContainerResponse struct {
	value *WebhookEventResendContainerResponse
	isSet bool
}

func (v NullableWebhookEventResendContainerResponse) Get() *WebhookEventResendContainerResponse {
	return v.value
}

func (v *NullableWebhookEventResendContainerResponse) Set(val *WebhookEventResendContainerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookEventResendContainerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookEventResendContainerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookEventResendContainerResponse(val *WebhookEventResendContainerResponse) *NullableWebhookEventResendContainerResponse {
	return &NullableWebhookEventResendContainerResponse{value: val, isSet: true}
}

func (v NullableWebhookEventResendContainerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookEventResendContainerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


