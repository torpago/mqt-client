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
	"fmt"
)

// ProductClassification Specifies for whom the credit product is intended.  * `CONSUMER` - The credit product is intended for individual consumers. * `SMALL_AND_MEDIUM_BUSINESS` - The credit product is intended for small and medium business.
type ProductClassification string

// List of ProductClassification
const (
	PRODUCTCLASSIFICATION_CONSUMER ProductClassification = "CONSUMER"
	PRODUCTCLASSIFICATION_SMALL_AND_MEDIUM_BUSINESS ProductClassification = "SMALL_AND_MEDIUM_BUSINESS"
)

// All allowed values of ProductClassification enum
var AllowedProductClassificationEnumValues = []ProductClassification{
	"CONSUMER",
	"SMALL_AND_MEDIUM_BUSINESS",
}

func (v *ProductClassification) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProductClassification(value)
	for _, existing := range AllowedProductClassificationEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProductClassification", value)
}

// NewProductClassificationFromValue returns a pointer to a valid ProductClassification
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProductClassificationFromValue(v string) (*ProductClassification, error) {
	ev := ProductClassification(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProductClassification: valid values are %v", v, AllowedProductClassificationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProductClassification) IsValid() bool {
	for _, existing := range AllowedProductClassificationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductClassification value
func (v ProductClassification) Ptr() *ProductClassification {
	return &v
}

type NullableProductClassification struct {
	value *ProductClassification
	isSet bool
}

func (v NullableProductClassification) Get() *ProductClassification {
	return v.value
}

func (v *NullableProductClassification) Set(val *ProductClassification) {
	v.value = val
	v.isSet = true
}

func (v NullableProductClassification) IsSet() bool {
	return v.isSet
}

func (v *NullableProductClassification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductClassification(val *ProductClassification) *NullableProductClassification {
	return &NullableProductClassification{value: val, isSet: true}
}

func (v NullableProductClassification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductClassification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

