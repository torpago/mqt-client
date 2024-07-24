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

// checks if the PolicyDocumentTemplateReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicyDocumentTemplateReq{}

// PolicyDocumentTemplateReq Request details for a template.
type PolicyDocumentTemplateReq struct {
	// Unique identifier of a template, which is a document that serves as an initial disclosure but does not contain finalized values.
	TemplateToken string `json:"template_token" validate:"regexp=(?!^ +$)^.+$"`
}

type _PolicyDocumentTemplateReq PolicyDocumentTemplateReq

// NewPolicyDocumentTemplateReq instantiates a new PolicyDocumentTemplateReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyDocumentTemplateReq(templateToken string) *PolicyDocumentTemplateReq {
	this := PolicyDocumentTemplateReq{}
	this.TemplateToken = templateToken
	return &this
}

// NewPolicyDocumentTemplateReqWithDefaults instantiates a new PolicyDocumentTemplateReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyDocumentTemplateReqWithDefaults() *PolicyDocumentTemplateReq {
	this := PolicyDocumentTemplateReq{}
	return &this
}

// GetTemplateToken returns the TemplateToken field value
func (o *PolicyDocumentTemplateReq) GetTemplateToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TemplateToken
}

// GetTemplateTokenOk returns a tuple with the TemplateToken field value
// and a boolean to check if the value has been set.
func (o *PolicyDocumentTemplateReq) GetTemplateTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateToken, true
}

// SetTemplateToken sets field value
func (o *PolicyDocumentTemplateReq) SetTemplateToken(v string) {
	o.TemplateToken = v
}

func (o PolicyDocumentTemplateReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyDocumentTemplateReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["template_token"] = o.TemplateToken
	return toSerialize, nil
}

func (o *PolicyDocumentTemplateReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"template_token",
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

	varPolicyDocumentTemplateReq := _PolicyDocumentTemplateReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPolicyDocumentTemplateReq)

	if err != nil {
		return err
	}

	*o = PolicyDocumentTemplateReq(varPolicyDocumentTemplateReq)

	return err
}

type NullablePolicyDocumentTemplateReq struct {
	value *PolicyDocumentTemplateReq
	isSet bool
}

func (v NullablePolicyDocumentTemplateReq) Get() *PolicyDocumentTemplateReq {
	return v.value
}

func (v *NullablePolicyDocumentTemplateReq) Set(val *PolicyDocumentTemplateReq) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyDocumentTemplateReq) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyDocumentTemplateReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyDocumentTemplateReq(val *PolicyDocumentTemplateReq) *NullablePolicyDocumentTemplateReq {
	return &NullablePolicyDocumentTemplateReq{value: val, isSet: true}
}

func (v NullablePolicyDocumentTemplateReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyDocumentTemplateReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


