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
	"time"
)

// checks if the ChargebackFundingSourceModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChargebackFundingSourceModel{}

// ChargebackFundingSourceModel struct for ChargebackFundingSourceModel
type ChargebackFundingSourceModel struct {
	FundingSourceModel
	Credit bool `json:"credit"`
	Name string `json:"name"`
}

type _ChargebackFundingSourceModel ChargebackFundingSourceModel

// NewChargebackFundingSourceModel instantiates a new ChargebackFundingSourceModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChargebackFundingSourceModel(credit bool, name string, active bool, createdTime time.Time, isDefaultAccount bool, lastModifiedTime time.Time, token string, type_ string) *ChargebackFundingSourceModel {
	this := ChargebackFundingSourceModel{}
	this.Active = active
	this.CreatedTime = createdTime
	this.IsDefaultAccount = isDefaultAccount
	this.LastModifiedTime = lastModifiedTime
	this.Token = token
	this.Type = type_
	this.Credit = credit
	this.Name = name
	return &this
}

// NewChargebackFundingSourceModelWithDefaults instantiates a new ChargebackFundingSourceModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChargebackFundingSourceModelWithDefaults() *ChargebackFundingSourceModel {
	this := ChargebackFundingSourceModel{}
	return &this
}

// GetCredit returns the Credit field value
func (o *ChargebackFundingSourceModel) GetCredit() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Credit
}

// GetCreditOk returns a tuple with the Credit field value
// and a boolean to check if the value has been set.
func (o *ChargebackFundingSourceModel) GetCreditOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Credit, true
}

// SetCredit sets field value
func (o *ChargebackFundingSourceModel) SetCredit(v bool) {
	o.Credit = v
}

// GetName returns the Name field value
func (o *ChargebackFundingSourceModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ChargebackFundingSourceModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ChargebackFundingSourceModel) SetName(v string) {
	o.Name = v
}

func (o ChargebackFundingSourceModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChargebackFundingSourceModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedFundingSourceModel, errFundingSourceModel := json.Marshal(o.FundingSourceModel)
	if errFundingSourceModel != nil {
		return map[string]interface{}{}, errFundingSourceModel
	}
	errFundingSourceModel = json.Unmarshal([]byte(serializedFundingSourceModel), &toSerialize)
	if errFundingSourceModel != nil {
		return map[string]interface{}{}, errFundingSourceModel
	}
	toSerialize["credit"] = o.Credit
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *ChargebackFundingSourceModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"credit",
		"name",
		"active",
		"created_time",
		"is_default_account",
		"last_modified_time",
		"token",
		"type",
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

	varChargebackFundingSourceModel := _ChargebackFundingSourceModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChargebackFundingSourceModel)

	if err != nil {
		return err
	}

	*o = ChargebackFundingSourceModel(varChargebackFundingSourceModel)

	return err
}

type NullableChargebackFundingSourceModel struct {
	value *ChargebackFundingSourceModel
	isSet bool
}

func (v NullableChargebackFundingSourceModel) Get() *ChargebackFundingSourceModel {
	return v.value
}

func (v *NullableChargebackFundingSourceModel) Set(val *ChargebackFundingSourceModel) {
	v.value = val
	v.isSet = true
}

func (v NullableChargebackFundingSourceModel) IsSet() bool {
	return v.isSet
}

func (v *NullableChargebackFundingSourceModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChargebackFundingSourceModel(val *ChargebackFundingSourceModel) *NullableChargebackFundingSourceModel {
	return &NullableChargebackFundingSourceModel{value: val, isSet: true}
}

func (v NullableChargebackFundingSourceModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChargebackFundingSourceModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


