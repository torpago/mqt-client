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

// checks if the SettlementData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SettlementData{}

// SettlementData Contains information from the card network about currency conversion at the time of settlement, including the original currency of the transaction, the amount of the transaction in the original currency, and the conversion rate.
type SettlementData struct {
	// The settled amount.
	Amount *decimal.Decimal `json:"amount,omitempty"`
	// Returned when the transaction currency is different from the origination currency.  Conversion rate between the origination currency and the settlement currency.
	ConversionRate *decimal.Decimal `json:"conversion_rate,omitempty"`
	// The ISO 4217 code of the currency used in the transaction.
	CurrencyCode *string `json:"currency_code,omitempty"`
}

// NewSettlementData instantiates a new SettlementData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettlementData() *SettlementData {
	this := SettlementData{}
	return &this
}

// NewSettlementDataWithDefaults instantiates a new SettlementData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettlementDataWithDefaults() *SettlementData {
	this := SettlementData{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *SettlementData) GetAmount() decimal.Decimal {
	if o == nil || IsNil(o.Amount) {
		var ret decimal.Decimal
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettlementData) GetAmountOk() (*decimal.Decimal, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *SettlementData) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given decimal.Decimal and assigns it to the Amount field.
func (o *SettlementData) SetAmount(v decimal.Decimal) {
	o.Amount = &v
}

// GetConversionRate returns the ConversionRate field value if set, zero value otherwise.
func (o *SettlementData) GetConversionRate() decimal.Decimal {
	if o == nil || IsNil(o.ConversionRate) {
		var ret decimal.Decimal
		return ret
	}
	return *o.ConversionRate
}

// GetConversionRateOk returns a tuple with the ConversionRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettlementData) GetConversionRateOk() (*decimal.Decimal, bool) {
	if o == nil || IsNil(o.ConversionRate) {
		return nil, false
	}
	return o.ConversionRate, true
}

// HasConversionRate returns a boolean if a field has been set.
func (o *SettlementData) HasConversionRate() bool {
	if o != nil && !IsNil(o.ConversionRate) {
		return true
	}

	return false
}

// SetConversionRate gets a reference to the given decimal.Decimal and assigns it to the ConversionRate field.
func (o *SettlementData) SetConversionRate(v decimal.Decimal) {
	o.ConversionRate = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *SettlementData) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettlementData) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *SettlementData) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *SettlementData) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

func (o SettlementData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SettlementData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.ConversionRate) {
		toSerialize["conversion_rate"] = o.ConversionRate
	}
	if !IsNil(o.CurrencyCode) {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	return toSerialize, nil
}

type NullableSettlementData struct {
	value *SettlementData
	isSet bool
}

func (v NullableSettlementData) Get() *SettlementData {
	return v.value
}

func (v *NullableSettlementData) Set(val *SettlementData) {
	v.value = val
	v.isSet = true
}

func (v NullableSettlementData) IsSet() bool {
	return v.isSet
}

func (v *NullableSettlementData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettlementData(val *SettlementData) *NullableSettlementData {
	return &NullableSettlementData{value: val, isSet: true}
}

func (v NullableSettlementData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettlementData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


