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

// checks if the WebPushProvisioningApplePayJWTResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebPushProvisioningApplePayJWTResponse{}

// WebPushProvisioningApplePayJWTResponse struct for WebPushProvisioningApplePayJWTResponse
type WebPushProvisioningApplePayJWTResponse struct {
	Jws WebPushProvisioningApplePayJWSModel `json:"jws"`
	// Unique state associated with the digital wallet token. The Marqeta platform returns a universally unique identifier (UUID) in this field.
	State string `json:"state"`
}

type _WebPushProvisioningApplePayJWTResponse WebPushProvisioningApplePayJWTResponse

// NewWebPushProvisioningApplePayJWTResponse instantiates a new WebPushProvisioningApplePayJWTResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebPushProvisioningApplePayJWTResponse(jws WebPushProvisioningApplePayJWSModel, state string) *WebPushProvisioningApplePayJWTResponse {
	this := WebPushProvisioningApplePayJWTResponse{}
	this.Jws = jws
	this.State = state
	return &this
}

// NewWebPushProvisioningApplePayJWTResponseWithDefaults instantiates a new WebPushProvisioningApplePayJWTResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebPushProvisioningApplePayJWTResponseWithDefaults() *WebPushProvisioningApplePayJWTResponse {
	this := WebPushProvisioningApplePayJWTResponse{}
	return &this
}

// GetJws returns the Jws field value
func (o *WebPushProvisioningApplePayJWTResponse) GetJws() WebPushProvisioningApplePayJWSModel {
	if o == nil {
		var ret WebPushProvisioningApplePayJWSModel
		return ret
	}

	return o.Jws
}

// GetJwsOk returns a tuple with the Jws field value
// and a boolean to check if the value has been set.
func (o *WebPushProvisioningApplePayJWTResponse) GetJwsOk() (*WebPushProvisioningApplePayJWSModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Jws, true
}

// SetJws sets field value
func (o *WebPushProvisioningApplePayJWTResponse) SetJws(v WebPushProvisioningApplePayJWSModel) {
	o.Jws = v
}

// GetState returns the State field value
func (o *WebPushProvisioningApplePayJWTResponse) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *WebPushProvisioningApplePayJWTResponse) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *WebPushProvisioningApplePayJWTResponse) SetState(v string) {
	o.State = v
}

func (o WebPushProvisioningApplePayJWTResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebPushProvisioningApplePayJWTResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["jws"] = o.Jws
	toSerialize["state"] = o.State
	return toSerialize, nil
}

func (o *WebPushProvisioningApplePayJWTResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"jws",
		"state",
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

	varWebPushProvisioningApplePayJWTResponse := _WebPushProvisioningApplePayJWTResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWebPushProvisioningApplePayJWTResponse)

	if err != nil {
		return err
	}

	*o = WebPushProvisioningApplePayJWTResponse(varWebPushProvisioningApplePayJWTResponse)

	return err
}

type NullableWebPushProvisioningApplePayJWTResponse struct {
	value *WebPushProvisioningApplePayJWTResponse
	isSet bool
}

func (v NullableWebPushProvisioningApplePayJWTResponse) Get() *WebPushProvisioningApplePayJWTResponse {
	return v.value
}

func (v *NullableWebPushProvisioningApplePayJWTResponse) Set(val *WebPushProvisioningApplePayJWTResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWebPushProvisioningApplePayJWTResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWebPushProvisioningApplePayJWTResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebPushProvisioningApplePayJWTResponse(val *WebPushProvisioningApplePayJWTResponse) *NullableWebPushProvisioningApplePayJWTResponse {
	return &NullableWebPushProvisioningApplePayJWTResponse{value: val, isSet: true}
}

func (v NullableWebPushProvisioningApplePayJWTResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebPushProvisioningApplePayJWTResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


