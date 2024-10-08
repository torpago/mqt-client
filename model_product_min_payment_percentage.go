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
	"github.com/shopspring/decimal"
)

// checks if the ProductMinPaymentPercentage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductMinPaymentPercentage{}

// ProductMinPaymentPercentage Contains information used to calculate the minimum payment percentage.
type ProductMinPaymentPercentage struct {
	// One or more fee types to include when calculating the minimum payment.
	IncludeFeesCharged []ProductFeeType `json:"include_fees_charged"`
	// Minimum payment, expressed as a percentage of the total statement balance, due on the payment due day.
	PercentageOfBalance decimal.Decimal `json:"percentage_of_balance"`
}

type _ProductMinPaymentPercentage ProductMinPaymentPercentage

// NewProductMinPaymentPercentage instantiates a new ProductMinPaymentPercentage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductMinPaymentPercentage(includeFeesCharged []ProductFeeType, percentageOfBalance decimal.Decimal) *ProductMinPaymentPercentage {
	this := ProductMinPaymentPercentage{}
	this.IncludeFeesCharged = includeFeesCharged
	this.PercentageOfBalance = percentageOfBalance
	return &this
}

// NewProductMinPaymentPercentageWithDefaults instantiates a new ProductMinPaymentPercentage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductMinPaymentPercentageWithDefaults() *ProductMinPaymentPercentage {
	this := ProductMinPaymentPercentage{}
	return &this
}

// GetIncludeFeesCharged returns the IncludeFeesCharged field value
func (o *ProductMinPaymentPercentage) GetIncludeFeesCharged() []ProductFeeType {
	if o == nil {
		var ret []ProductFeeType
		return ret
	}

	return o.IncludeFeesCharged
}

// GetIncludeFeesChargedOk returns a tuple with the IncludeFeesCharged field value
// and a boolean to check if the value has been set.
func (o *ProductMinPaymentPercentage) GetIncludeFeesChargedOk() ([]ProductFeeType, bool) {
	if o == nil {
		return nil, false
	}
	return o.IncludeFeesCharged, true
}

// SetIncludeFeesCharged sets field value
func (o *ProductMinPaymentPercentage) SetIncludeFeesCharged(v []ProductFeeType) {
	o.IncludeFeesCharged = v
}

// GetPercentageOfBalance returns the PercentageOfBalance field value
func (o *ProductMinPaymentPercentage) GetPercentageOfBalance() decimal.Decimal {
	if o == nil {
		var ret decimal.Decimal
		return ret
	}

	return o.PercentageOfBalance
}

// GetPercentageOfBalanceOk returns a tuple with the PercentageOfBalance field value
// and a boolean to check if the value has been set.
func (o *ProductMinPaymentPercentage) GetPercentageOfBalanceOk() (*decimal.Decimal, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PercentageOfBalance, true
}

// SetPercentageOfBalance sets field value
func (o *ProductMinPaymentPercentage) SetPercentageOfBalance(v decimal.Decimal) {
	o.PercentageOfBalance = v
}

func (o ProductMinPaymentPercentage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductMinPaymentPercentage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["include_fees_charged"] = o.IncludeFeesCharged
	toSerialize["percentage_of_balance"] = o.PercentageOfBalance
	return toSerialize, nil
}

func (o *ProductMinPaymentPercentage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"include_fees_charged",
		"percentage_of_balance",
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

	varProductMinPaymentPercentage := _ProductMinPaymentPercentage{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProductMinPaymentPercentage)

	if err != nil {
		return err
	}

	*o = ProductMinPaymentPercentage(varProductMinPaymentPercentage)

	return err
}

type NullableProductMinPaymentPercentage struct {
	value *ProductMinPaymentPercentage
	isSet bool
}

func (v NullableProductMinPaymentPercentage) Get() *ProductMinPaymentPercentage {
	return v.value
}

func (v *NullableProductMinPaymentPercentage) Set(val *ProductMinPaymentPercentage) {
	v.value = val
	v.isSet = true
}

func (v NullableProductMinPaymentPercentage) IsSet() bool {
	return v.isSet
}

func (v *NullableProductMinPaymentPercentage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductMinPaymentPercentage(val *ProductMinPaymentPercentage) *NullableProductMinPaymentPercentage {
	return &NullableProductMinPaymentPercentage{value: val, isSet: true}
}

func (v NullableProductMinPaymentPercentage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductMinPaymentPercentage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


