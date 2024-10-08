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
	"time"
	"bytes"
	"fmt"
)

// checks if the AuthControlUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthControlUpdateRequest{}

// AuthControlUpdateRequest struct for AuthControlUpdateRequest
type AuthControlUpdateRequest struct {
	// Indicates whether the authorization control is active.
	Active *bool `json:"active,omitempty"`
	Association *SpendControlAssociation `json:"association,omitempty"`
	// Date and time when the exception ends, in UTC.
	EndTime *time.Time `json:"end_time,omitempty"`
	MerchantScope *MerchantScope `json:"merchant_scope,omitempty"`
	// Name of the authorization control.
	Name *string `json:"name,omitempty"`
	// Date and time when the exception goes into effect, in UTC.
	StartTime *time.Time `json:"start_time,omitempty"`
	// Unique identifier of the authorization control.
	Token string `json:"token"`
}

type _AuthControlUpdateRequest AuthControlUpdateRequest

// NewAuthControlUpdateRequest instantiates a new AuthControlUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthControlUpdateRequest(token string) *AuthControlUpdateRequest {
	this := AuthControlUpdateRequest{}
	var active bool = true
	this.Active = &active
	this.Token = token
	return &this
}

// NewAuthControlUpdateRequestWithDefaults instantiates a new AuthControlUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthControlUpdateRequestWithDefaults() *AuthControlUpdateRequest {
	this := AuthControlUpdateRequest{}
	var active bool = true
	this.Active = &active
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *AuthControlUpdateRequest) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthControlUpdateRequest) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *AuthControlUpdateRequest) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *AuthControlUpdateRequest) SetActive(v bool) {
	o.Active = &v
}

// GetAssociation returns the Association field value if set, zero value otherwise.
func (o *AuthControlUpdateRequest) GetAssociation() SpendControlAssociation {
	if o == nil || IsNil(o.Association) {
		var ret SpendControlAssociation
		return ret
	}
	return *o.Association
}

// GetAssociationOk returns a tuple with the Association field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthControlUpdateRequest) GetAssociationOk() (*SpendControlAssociation, bool) {
	if o == nil || IsNil(o.Association) {
		return nil, false
	}
	return o.Association, true
}

// HasAssociation returns a boolean if a field has been set.
func (o *AuthControlUpdateRequest) HasAssociation() bool {
	if o != nil && !IsNil(o.Association) {
		return true
	}

	return false
}

// SetAssociation gets a reference to the given SpendControlAssociation and assigns it to the Association field.
func (o *AuthControlUpdateRequest) SetAssociation(v SpendControlAssociation) {
	o.Association = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *AuthControlUpdateRequest) GetEndTime() time.Time {
	if o == nil || IsNil(o.EndTime) {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthControlUpdateRequest) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *AuthControlUpdateRequest) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *AuthControlUpdateRequest) SetEndTime(v time.Time) {
	o.EndTime = &v
}

// GetMerchantScope returns the MerchantScope field value if set, zero value otherwise.
func (o *AuthControlUpdateRequest) GetMerchantScope() MerchantScope {
	if o == nil || IsNil(o.MerchantScope) {
		var ret MerchantScope
		return ret
	}
	return *o.MerchantScope
}

// GetMerchantScopeOk returns a tuple with the MerchantScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthControlUpdateRequest) GetMerchantScopeOk() (*MerchantScope, bool) {
	if o == nil || IsNil(o.MerchantScope) {
		return nil, false
	}
	return o.MerchantScope, true
}

// HasMerchantScope returns a boolean if a field has been set.
func (o *AuthControlUpdateRequest) HasMerchantScope() bool {
	if o != nil && !IsNil(o.MerchantScope) {
		return true
	}

	return false
}

// SetMerchantScope gets a reference to the given MerchantScope and assigns it to the MerchantScope field.
func (o *AuthControlUpdateRequest) SetMerchantScope(v MerchantScope) {
	o.MerchantScope = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AuthControlUpdateRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthControlUpdateRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AuthControlUpdateRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AuthControlUpdateRequest) SetName(v string) {
	o.Name = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *AuthControlUpdateRequest) GetStartTime() time.Time {
	if o == nil || IsNil(o.StartTime) {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthControlUpdateRequest) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *AuthControlUpdateRequest) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *AuthControlUpdateRequest) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetToken returns the Token field value
func (o *AuthControlUpdateRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *AuthControlUpdateRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *AuthControlUpdateRequest) SetToken(v string) {
	o.Token = v
}

func (o AuthControlUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthControlUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.Association) {
		toSerialize["association"] = o.Association
	}
	if !IsNil(o.EndTime) {
		toSerialize["end_time"] = o.EndTime
	}
	if !IsNil(o.MerchantScope) {
		toSerialize["merchant_scope"] = o.MerchantScope
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.StartTime) {
		toSerialize["start_time"] = o.StartTime
	}
	toSerialize["token"] = o.Token
	return toSerialize, nil
}

func (o *AuthControlUpdateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"token",
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

	varAuthControlUpdateRequest := _AuthControlUpdateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAuthControlUpdateRequest)

	if err != nil {
		return err
	}

	*o = AuthControlUpdateRequest(varAuthControlUpdateRequest)

	return err
}

type NullableAuthControlUpdateRequest struct {
	value *AuthControlUpdateRequest
	isSet bool
}

func (v NullableAuthControlUpdateRequest) Get() *AuthControlUpdateRequest {
	return v.value
}

func (v *NullableAuthControlUpdateRequest) Set(val *AuthControlUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthControlUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthControlUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthControlUpdateRequest(val *AuthControlUpdateRequest) *NullableAuthControlUpdateRequest {
	return &NullableAuthControlUpdateRequest{value: val, isSet: true}
}

func (v NullableAuthControlUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthControlUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


