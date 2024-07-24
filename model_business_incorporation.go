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

// checks if the BusinessIncorporation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BusinessIncorporation{}

// BusinessIncorporation Contains information about the organizational structure of the business.  This object is required for KYC verification (US-based accounts only).
type BusinessIncorporation struct {
	AddressRegisteredUnder *AddressRequestModel `json:"address_registered_under,omitempty"`
	// Organizational structure of the business, such as corporation or sole proprietorship.  This field is required for KYC verification (US-based accounts only).
	IncorporationType *string `json:"incorporation_type,omitempty"`
	// A value of `true` indicates that the business is publicly held.
	IsPublic *bool `json:"is_public,omitempty"`
	// Name under which the business is registered.
	NameRegisteredUnder *string `json:"name_registered_under,omitempty"`
	// State where the business is incorporated.  This field is required for KYC verification (US-based accounts only).
	StateOfIncorporation *string `json:"state_of_incorporation,omitempty"`
	// Business stock symbol.
	StockSymbol *string `json:"stock_symbol,omitempty"`
}

// NewBusinessIncorporation instantiates a new BusinessIncorporation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBusinessIncorporation() *BusinessIncorporation {
	this := BusinessIncorporation{}
	var isPublic bool = false
	this.IsPublic = &isPublic
	return &this
}

// NewBusinessIncorporationWithDefaults instantiates a new BusinessIncorporation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBusinessIncorporationWithDefaults() *BusinessIncorporation {
	this := BusinessIncorporation{}
	var isPublic bool = false
	this.IsPublic = &isPublic
	return &this
}

// GetAddressRegisteredUnder returns the AddressRegisteredUnder field value if set, zero value otherwise.
func (o *BusinessIncorporation) GetAddressRegisteredUnder() AddressRequestModel {
	if o == nil || IsNil(o.AddressRegisteredUnder) {
		var ret AddressRequestModel
		return ret
	}
	return *o.AddressRegisteredUnder
}

// GetAddressRegisteredUnderOk returns a tuple with the AddressRegisteredUnder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessIncorporation) GetAddressRegisteredUnderOk() (*AddressRequestModel, bool) {
	if o == nil || IsNil(o.AddressRegisteredUnder) {
		return nil, false
	}
	return o.AddressRegisteredUnder, true
}

// HasAddressRegisteredUnder returns a boolean if a field has been set.
func (o *BusinessIncorporation) HasAddressRegisteredUnder() bool {
	if o != nil && !IsNil(o.AddressRegisteredUnder) {
		return true
	}

	return false
}

// SetAddressRegisteredUnder gets a reference to the given AddressRequestModel and assigns it to the AddressRegisteredUnder field.
func (o *BusinessIncorporation) SetAddressRegisteredUnder(v AddressRequestModel) {
	o.AddressRegisteredUnder = &v
}

// GetIncorporationType returns the IncorporationType field value if set, zero value otherwise.
func (o *BusinessIncorporation) GetIncorporationType() string {
	if o == nil || IsNil(o.IncorporationType) {
		var ret string
		return ret
	}
	return *o.IncorporationType
}

// GetIncorporationTypeOk returns a tuple with the IncorporationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessIncorporation) GetIncorporationTypeOk() (*string, bool) {
	if o == nil || IsNil(o.IncorporationType) {
		return nil, false
	}
	return o.IncorporationType, true
}

// HasIncorporationType returns a boolean if a field has been set.
func (o *BusinessIncorporation) HasIncorporationType() bool {
	if o != nil && !IsNil(o.IncorporationType) {
		return true
	}

	return false
}

// SetIncorporationType gets a reference to the given string and assigns it to the IncorporationType field.
func (o *BusinessIncorporation) SetIncorporationType(v string) {
	o.IncorporationType = &v
}

// GetIsPublic returns the IsPublic field value if set, zero value otherwise.
func (o *BusinessIncorporation) GetIsPublic() bool {
	if o == nil || IsNil(o.IsPublic) {
		var ret bool
		return ret
	}
	return *o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessIncorporation) GetIsPublicOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPublic) {
		return nil, false
	}
	return o.IsPublic, true
}

// HasIsPublic returns a boolean if a field has been set.
func (o *BusinessIncorporation) HasIsPublic() bool {
	if o != nil && !IsNil(o.IsPublic) {
		return true
	}

	return false
}

// SetIsPublic gets a reference to the given bool and assigns it to the IsPublic field.
func (o *BusinessIncorporation) SetIsPublic(v bool) {
	o.IsPublic = &v
}

// GetNameRegisteredUnder returns the NameRegisteredUnder field value if set, zero value otherwise.
func (o *BusinessIncorporation) GetNameRegisteredUnder() string {
	if o == nil || IsNil(o.NameRegisteredUnder) {
		var ret string
		return ret
	}
	return *o.NameRegisteredUnder
}

// GetNameRegisteredUnderOk returns a tuple with the NameRegisteredUnder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessIncorporation) GetNameRegisteredUnderOk() (*string, bool) {
	if o == nil || IsNil(o.NameRegisteredUnder) {
		return nil, false
	}
	return o.NameRegisteredUnder, true
}

// HasNameRegisteredUnder returns a boolean if a field has been set.
func (o *BusinessIncorporation) HasNameRegisteredUnder() bool {
	if o != nil && !IsNil(o.NameRegisteredUnder) {
		return true
	}

	return false
}

// SetNameRegisteredUnder gets a reference to the given string and assigns it to the NameRegisteredUnder field.
func (o *BusinessIncorporation) SetNameRegisteredUnder(v string) {
	o.NameRegisteredUnder = &v
}

// GetStateOfIncorporation returns the StateOfIncorporation field value if set, zero value otherwise.
func (o *BusinessIncorporation) GetStateOfIncorporation() string {
	if o == nil || IsNil(o.StateOfIncorporation) {
		var ret string
		return ret
	}
	return *o.StateOfIncorporation
}

// GetStateOfIncorporationOk returns a tuple with the StateOfIncorporation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessIncorporation) GetStateOfIncorporationOk() (*string, bool) {
	if o == nil || IsNil(o.StateOfIncorporation) {
		return nil, false
	}
	return o.StateOfIncorporation, true
}

// HasStateOfIncorporation returns a boolean if a field has been set.
func (o *BusinessIncorporation) HasStateOfIncorporation() bool {
	if o != nil && !IsNil(o.StateOfIncorporation) {
		return true
	}

	return false
}

// SetStateOfIncorporation gets a reference to the given string and assigns it to the StateOfIncorporation field.
func (o *BusinessIncorporation) SetStateOfIncorporation(v string) {
	o.StateOfIncorporation = &v
}

// GetStockSymbol returns the StockSymbol field value if set, zero value otherwise.
func (o *BusinessIncorporation) GetStockSymbol() string {
	if o == nil || IsNil(o.StockSymbol) {
		var ret string
		return ret
	}
	return *o.StockSymbol
}

// GetStockSymbolOk returns a tuple with the StockSymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessIncorporation) GetStockSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.StockSymbol) {
		return nil, false
	}
	return o.StockSymbol, true
}

// HasStockSymbol returns a boolean if a field has been set.
func (o *BusinessIncorporation) HasStockSymbol() bool {
	if o != nil && !IsNil(o.StockSymbol) {
		return true
	}

	return false
}

// SetStockSymbol gets a reference to the given string and assigns it to the StockSymbol field.
func (o *BusinessIncorporation) SetStockSymbol(v string) {
	o.StockSymbol = &v
}

func (o BusinessIncorporation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BusinessIncorporation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddressRegisteredUnder) {
		toSerialize["address_registered_under"] = o.AddressRegisteredUnder
	}
	if !IsNil(o.IncorporationType) {
		toSerialize["incorporation_type"] = o.IncorporationType
	}
	if !IsNil(o.IsPublic) {
		toSerialize["is_public"] = o.IsPublic
	}
	if !IsNil(o.NameRegisteredUnder) {
		toSerialize["name_registered_under"] = o.NameRegisteredUnder
	}
	if !IsNil(o.StateOfIncorporation) {
		toSerialize["state_of_incorporation"] = o.StateOfIncorporation
	}
	if !IsNil(o.StockSymbol) {
		toSerialize["stock_symbol"] = o.StockSymbol
	}
	return toSerialize, nil
}

type NullableBusinessIncorporation struct {
	value *BusinessIncorporation
	isSet bool
}

func (v NullableBusinessIncorporation) Get() *BusinessIncorporation {
	return v.value
}

func (v *NullableBusinessIncorporation) Set(val *BusinessIncorporation) {
	v.value = val
	v.isSet = true
}

func (v NullableBusinessIncorporation) IsSet() bool {
	return v.isSet
}

func (v *NullableBusinessIncorporation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBusinessIncorporation(val *BusinessIncorporation) *NullableBusinessIncorporation {
	return &NullableBusinessIncorporation{value: val, isSet: true}
}

func (v NullableBusinessIncorporation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBusinessIncorporation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


