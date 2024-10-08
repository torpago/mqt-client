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

// checks if the DigitalServiceProvider type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DigitalServiceProvider{}

// DigitalServiceProvider struct for DigitalServiceProvider
type DigitalServiceProvider struct {
	CvmAuthenticationIndicator *string `json:"cvm_authentication_indicator,omitempty"`
	CvmAuthenticationMethod *string `json:"cvm_authentication_method,omitempty"`
	FirstAuthenticationFactor *string `json:"first_authentication_factor,omitempty"`
	SecondAuthenticationFactor *string `json:"second_authentication_factor,omitempty"`
}

// NewDigitalServiceProvider instantiates a new DigitalServiceProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDigitalServiceProvider() *DigitalServiceProvider {
	this := DigitalServiceProvider{}
	return &this
}

// NewDigitalServiceProviderWithDefaults instantiates a new DigitalServiceProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDigitalServiceProviderWithDefaults() *DigitalServiceProvider {
	this := DigitalServiceProvider{}
	return &this
}

// GetCvmAuthenticationIndicator returns the CvmAuthenticationIndicator field value if set, zero value otherwise.
func (o *DigitalServiceProvider) GetCvmAuthenticationIndicator() string {
	if o == nil || IsNil(o.CvmAuthenticationIndicator) {
		var ret string
		return ret
	}
	return *o.CvmAuthenticationIndicator
}

// GetCvmAuthenticationIndicatorOk returns a tuple with the CvmAuthenticationIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigitalServiceProvider) GetCvmAuthenticationIndicatorOk() (*string, bool) {
	if o == nil || IsNil(o.CvmAuthenticationIndicator) {
		return nil, false
	}
	return o.CvmAuthenticationIndicator, true
}

// HasCvmAuthenticationIndicator returns a boolean if a field has been set.
func (o *DigitalServiceProvider) HasCvmAuthenticationIndicator() bool {
	if o != nil && !IsNil(o.CvmAuthenticationIndicator) {
		return true
	}

	return false
}

// SetCvmAuthenticationIndicator gets a reference to the given string and assigns it to the CvmAuthenticationIndicator field.
func (o *DigitalServiceProvider) SetCvmAuthenticationIndicator(v string) {
	o.CvmAuthenticationIndicator = &v
}

// GetCvmAuthenticationMethod returns the CvmAuthenticationMethod field value if set, zero value otherwise.
func (o *DigitalServiceProvider) GetCvmAuthenticationMethod() string {
	if o == nil || IsNil(o.CvmAuthenticationMethod) {
		var ret string
		return ret
	}
	return *o.CvmAuthenticationMethod
}

// GetCvmAuthenticationMethodOk returns a tuple with the CvmAuthenticationMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigitalServiceProvider) GetCvmAuthenticationMethodOk() (*string, bool) {
	if o == nil || IsNil(o.CvmAuthenticationMethod) {
		return nil, false
	}
	return o.CvmAuthenticationMethod, true
}

// HasCvmAuthenticationMethod returns a boolean if a field has been set.
func (o *DigitalServiceProvider) HasCvmAuthenticationMethod() bool {
	if o != nil && !IsNil(o.CvmAuthenticationMethod) {
		return true
	}

	return false
}

// SetCvmAuthenticationMethod gets a reference to the given string and assigns it to the CvmAuthenticationMethod field.
func (o *DigitalServiceProvider) SetCvmAuthenticationMethod(v string) {
	o.CvmAuthenticationMethod = &v
}

// GetFirstAuthenticationFactor returns the FirstAuthenticationFactor field value if set, zero value otherwise.
func (o *DigitalServiceProvider) GetFirstAuthenticationFactor() string {
	if o == nil || IsNil(o.FirstAuthenticationFactor) {
		var ret string
		return ret
	}
	return *o.FirstAuthenticationFactor
}

// GetFirstAuthenticationFactorOk returns a tuple with the FirstAuthenticationFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigitalServiceProvider) GetFirstAuthenticationFactorOk() (*string, bool) {
	if o == nil || IsNil(o.FirstAuthenticationFactor) {
		return nil, false
	}
	return o.FirstAuthenticationFactor, true
}

// HasFirstAuthenticationFactor returns a boolean if a field has been set.
func (o *DigitalServiceProvider) HasFirstAuthenticationFactor() bool {
	if o != nil && !IsNil(o.FirstAuthenticationFactor) {
		return true
	}

	return false
}

// SetFirstAuthenticationFactor gets a reference to the given string and assigns it to the FirstAuthenticationFactor field.
func (o *DigitalServiceProvider) SetFirstAuthenticationFactor(v string) {
	o.FirstAuthenticationFactor = &v
}

// GetSecondAuthenticationFactor returns the SecondAuthenticationFactor field value if set, zero value otherwise.
func (o *DigitalServiceProvider) GetSecondAuthenticationFactor() string {
	if o == nil || IsNil(o.SecondAuthenticationFactor) {
		var ret string
		return ret
	}
	return *o.SecondAuthenticationFactor
}

// GetSecondAuthenticationFactorOk returns a tuple with the SecondAuthenticationFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigitalServiceProvider) GetSecondAuthenticationFactorOk() (*string, bool) {
	if o == nil || IsNil(o.SecondAuthenticationFactor) {
		return nil, false
	}
	return o.SecondAuthenticationFactor, true
}

// HasSecondAuthenticationFactor returns a boolean if a field has been set.
func (o *DigitalServiceProvider) HasSecondAuthenticationFactor() bool {
	if o != nil && !IsNil(o.SecondAuthenticationFactor) {
		return true
	}

	return false
}

// SetSecondAuthenticationFactor gets a reference to the given string and assigns it to the SecondAuthenticationFactor field.
func (o *DigitalServiceProvider) SetSecondAuthenticationFactor(v string) {
	o.SecondAuthenticationFactor = &v
}

func (o DigitalServiceProvider) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DigitalServiceProvider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CvmAuthenticationIndicator) {
		toSerialize["cvm_authentication_indicator"] = o.CvmAuthenticationIndicator
	}
	if !IsNil(o.CvmAuthenticationMethod) {
		toSerialize["cvm_authentication_method"] = o.CvmAuthenticationMethod
	}
	if !IsNil(o.FirstAuthenticationFactor) {
		toSerialize["first_authentication_factor"] = o.FirstAuthenticationFactor
	}
	if !IsNil(o.SecondAuthenticationFactor) {
		toSerialize["second_authentication_factor"] = o.SecondAuthenticationFactor
	}
	return toSerialize, nil
}

type NullableDigitalServiceProvider struct {
	value *DigitalServiceProvider
	isSet bool
}

func (v NullableDigitalServiceProvider) Get() *DigitalServiceProvider {
	return v.value
}

func (v *NullableDigitalServiceProvider) Set(val *DigitalServiceProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableDigitalServiceProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableDigitalServiceProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDigitalServiceProvider(val *DigitalServiceProvider) *NullableDigitalServiceProvider {
	return &NullableDigitalServiceProvider{value: val, isSet: true}
}

func (v NullableDigitalServiceProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDigitalServiceProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


