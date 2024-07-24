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

// checks if the ProductMinPaymentCalculation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductMinPaymentCalculation{}

// ProductMinPaymentCalculation Contains information used to calculate the minimum payment amount.
type ProductMinPaymentCalculation struct {
	// Whether to include the past due amount when calculating the minimum payment.
	IncludePastDueAmount bool `json:"include_past_due_amount"`
	// Minimum payment, expressed as a flat amount, due on the payment due day.
	MinPaymentFlatAmount decimal.Decimal `json:"min_payment_flat_amount"`
	MinPaymentPercentage ProductMinPaymentPercentage `json:"min_payment_percentage"`
}

type _ProductMinPaymentCalculation ProductMinPaymentCalculation

// NewProductMinPaymentCalculation instantiates a new ProductMinPaymentCalculation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductMinPaymentCalculation(includePastDueAmount bool, minPaymentFlatAmount decimal.Decimal, minPaymentPercentage ProductMinPaymentPercentage) *ProductMinPaymentCalculation {
	this := ProductMinPaymentCalculation{}
	this.IncludePastDueAmount = includePastDueAmount
	this.MinPaymentFlatAmount = minPaymentFlatAmount
	this.MinPaymentPercentage = minPaymentPercentage
	return &this
}

// NewProductMinPaymentCalculationWithDefaults instantiates a new ProductMinPaymentCalculation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductMinPaymentCalculationWithDefaults() *ProductMinPaymentCalculation {
	this := ProductMinPaymentCalculation{}
	return &this
}

// GetIncludePastDueAmount returns the IncludePastDueAmount field value
func (o *ProductMinPaymentCalculation) GetIncludePastDueAmount() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IncludePastDueAmount
}

// GetIncludePastDueAmountOk returns a tuple with the IncludePastDueAmount field value
// and a boolean to check if the value has been set.
func (o *ProductMinPaymentCalculation) GetIncludePastDueAmountOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncludePastDueAmount, true
}

// SetIncludePastDueAmount sets field value
func (o *ProductMinPaymentCalculation) SetIncludePastDueAmount(v bool) {
	o.IncludePastDueAmount = v
}

// GetMinPaymentFlatAmount returns the MinPaymentFlatAmount field value
func (o *ProductMinPaymentCalculation) GetMinPaymentFlatAmount() decimal.Decimal {
	if o == nil {
		var ret decimal.Decimal
		return ret
	}

	return o.MinPaymentFlatAmount
}

// GetMinPaymentFlatAmountOk returns a tuple with the MinPaymentFlatAmount field value
// and a boolean to check if the value has been set.
func (o *ProductMinPaymentCalculation) GetMinPaymentFlatAmountOk() (*decimal.Decimal, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinPaymentFlatAmount, true
}

// SetMinPaymentFlatAmount sets field value
func (o *ProductMinPaymentCalculation) SetMinPaymentFlatAmount(v decimal.Decimal) {
	o.MinPaymentFlatAmount = v
}

// GetMinPaymentPercentage returns the MinPaymentPercentage field value
func (o *ProductMinPaymentCalculation) GetMinPaymentPercentage() ProductMinPaymentPercentage {
	if o == nil {
		var ret ProductMinPaymentPercentage
		return ret
	}

	return o.MinPaymentPercentage
}

// GetMinPaymentPercentageOk returns a tuple with the MinPaymentPercentage field value
// and a boolean to check if the value has been set.
func (o *ProductMinPaymentCalculation) GetMinPaymentPercentageOk() (*ProductMinPaymentPercentage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinPaymentPercentage, true
}

// SetMinPaymentPercentage sets field value
func (o *ProductMinPaymentCalculation) SetMinPaymentPercentage(v ProductMinPaymentPercentage) {
	o.MinPaymentPercentage = v
}

func (o ProductMinPaymentCalculation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductMinPaymentCalculation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["include_past_due_amount"] = o.IncludePastDueAmount
	toSerialize["min_payment_flat_amount"] = o.MinPaymentFlatAmount
	toSerialize["min_payment_percentage"] = o.MinPaymentPercentage
	return toSerialize, nil
}

func (o *ProductMinPaymentCalculation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"include_past_due_amount",
		"min_payment_flat_amount",
		"min_payment_percentage",
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

	varProductMinPaymentCalculation := _ProductMinPaymentCalculation{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProductMinPaymentCalculation)

	if err != nil {
		return err
	}

	*o = ProductMinPaymentCalculation(varProductMinPaymentCalculation)

	return err
}

type NullableProductMinPaymentCalculation struct {
	value *ProductMinPaymentCalculation
	isSet bool
}

func (v NullableProductMinPaymentCalculation) Get() *ProductMinPaymentCalculation {
	return v.value
}

func (v *NullableProductMinPaymentCalculation) Set(val *ProductMinPaymentCalculation) {
	v.value = val
	v.isSet = true
}

func (v NullableProductMinPaymentCalculation) IsSet() bool {
	return v.isSet
}

func (v *NullableProductMinPaymentCalculation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductMinPaymentCalculation(val *ProductMinPaymentCalculation) *NullableProductMinPaymentCalculation {
	return &NullableProductMinPaymentCalculation{value: val, isSet: true}
}

func (v NullableProductMinPaymentCalculation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductMinPaymentCalculation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


