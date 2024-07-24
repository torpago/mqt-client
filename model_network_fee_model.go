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
	"github.com/shopspring/decimal"
)

// checks if the NetworkFeeModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkFeeModel{}

// NetworkFeeModel Contains card network fees assessed against the cardholder.
type NetworkFeeModel struct {
	// The amount of the network fee.
	Amount *decimal.Decimal `json:"amount,omitempty"`
	// Indicates whether the fee is a credit or a debit.  * *C* indicates a credit * *D* indicates a debit
	CreditDebit *string `json:"credit_debit,omitempty"`
	// The type of fee assessed by the card network.
	Type *string `json:"type,omitempty"`
}

// NewNetworkFeeModel instantiates a new NetworkFeeModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkFeeModel() *NetworkFeeModel {
	this := NetworkFeeModel{}
	return &this
}

// NewNetworkFeeModelWithDefaults instantiates a new NetworkFeeModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkFeeModelWithDefaults() *NetworkFeeModel {
	this := NetworkFeeModel{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *NetworkFeeModel) GetAmount() decimal.Decimal {
	if o == nil || IsNil(o.Amount) {
		var ret decimal.Decimal
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkFeeModel) GetAmountOk() (*decimal.Decimal, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *NetworkFeeModel) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given decimal.Decimal and assigns it to the Amount field.
func (o *NetworkFeeModel) SetAmount(v decimal.Decimal) {
	o.Amount = &v
}

// GetCreditDebit returns the CreditDebit field value if set, zero value otherwise.
func (o *NetworkFeeModel) GetCreditDebit() string {
	if o == nil || IsNil(o.CreditDebit) {
		var ret string
		return ret
	}
	return *o.CreditDebit
}

// GetCreditDebitOk returns a tuple with the CreditDebit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkFeeModel) GetCreditDebitOk() (*string, bool) {
	if o == nil || IsNil(o.CreditDebit) {
		return nil, false
	}
	return o.CreditDebit, true
}

// HasCreditDebit returns a boolean if a field has been set.
func (o *NetworkFeeModel) HasCreditDebit() bool {
	if o != nil && !IsNil(o.CreditDebit) {
		return true
	}

	return false
}

// SetCreditDebit gets a reference to the given string and assigns it to the CreditDebit field.
func (o *NetworkFeeModel) SetCreditDebit(v string) {
	o.CreditDebit = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NetworkFeeModel) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkFeeModel) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NetworkFeeModel) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *NetworkFeeModel) SetType(v string) {
	o.Type = &v
}

func (o NetworkFeeModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkFeeModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.CreditDebit) {
		toSerialize["credit_debit"] = o.CreditDebit
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableNetworkFeeModel struct {
	value *NetworkFeeModel
	isSet bool
}

func (v NullableNetworkFeeModel) Get() *NetworkFeeModel {
	return v.value
}

func (v *NullableNetworkFeeModel) Set(val *NetworkFeeModel) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkFeeModel) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkFeeModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkFeeModel(val *NetworkFeeModel) *NullableNetworkFeeModel {
	return &NullableNetworkFeeModel{value: val, isSet: true}
}

func (v NullableNetworkFeeModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkFeeModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


