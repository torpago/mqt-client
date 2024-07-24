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

// checks if the NetworkMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkMetadata{}

// NetworkMetadata Contains network-related metadata for the transaction, including details about the card program and the card product. Returned if provided by the card network
type NetworkMetadata struct {
	AccountIdentification1 *string `json:"account_identification_1,omitempty"`
	// Visa Risk Management esponse code `59`, indicating suspected fraud.
	IncomingResponseCode *string `json:"incoming_response_code,omitempty"`
	// Product identification value assigned by the card network to each card product. Can be used to track card-level activity by individual account number for premium card products.
	ProductId *string `json:"product_id,omitempty"`
	// Program identification number used with `product_id` that identifies the programs associated with a card within a program registered by the issuer with the card network.
	ProgramId *string `json:"program_id,omitempty"`
	// Indicates whether or not the base spend-assessment threshold defined by the card network has been met.
	SpendQualifier *string `json:"spend_qualifier,omitempty"`
	// Name of the surcharge-free ATM network used to complete the transaction.
	SurchargeFreeAtmNetwork *string `json:"surcharge_free_atm_network,omitempty"`
}

// NewNetworkMetadata instantiates a new NetworkMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkMetadata() *NetworkMetadata {
	this := NetworkMetadata{}
	return &this
}

// NewNetworkMetadataWithDefaults instantiates a new NetworkMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkMetadataWithDefaults() *NetworkMetadata {
	this := NetworkMetadata{}
	return &this
}

// GetAccountIdentification1 returns the AccountIdentification1 field value if set, zero value otherwise.
func (o *NetworkMetadata) GetAccountIdentification1() string {
	if o == nil || IsNil(o.AccountIdentification1) {
		var ret string
		return ret
	}
	return *o.AccountIdentification1
}

// GetAccountIdentification1Ok returns a tuple with the AccountIdentification1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkMetadata) GetAccountIdentification1Ok() (*string, bool) {
	if o == nil || IsNil(o.AccountIdentification1) {
		return nil, false
	}
	return o.AccountIdentification1, true
}

// HasAccountIdentification1 returns a boolean if a field has been set.
func (o *NetworkMetadata) HasAccountIdentification1() bool {
	if o != nil && !IsNil(o.AccountIdentification1) {
		return true
	}

	return false
}

// SetAccountIdentification1 gets a reference to the given string and assigns it to the AccountIdentification1 field.
func (o *NetworkMetadata) SetAccountIdentification1(v string) {
	o.AccountIdentification1 = &v
}

// GetIncomingResponseCode returns the IncomingResponseCode field value if set, zero value otherwise.
func (o *NetworkMetadata) GetIncomingResponseCode() string {
	if o == nil || IsNil(o.IncomingResponseCode) {
		var ret string
		return ret
	}
	return *o.IncomingResponseCode
}

// GetIncomingResponseCodeOk returns a tuple with the IncomingResponseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkMetadata) GetIncomingResponseCodeOk() (*string, bool) {
	if o == nil || IsNil(o.IncomingResponseCode) {
		return nil, false
	}
	return o.IncomingResponseCode, true
}

// HasIncomingResponseCode returns a boolean if a field has been set.
func (o *NetworkMetadata) HasIncomingResponseCode() bool {
	if o != nil && !IsNil(o.IncomingResponseCode) {
		return true
	}

	return false
}

// SetIncomingResponseCode gets a reference to the given string and assigns it to the IncomingResponseCode field.
func (o *NetworkMetadata) SetIncomingResponseCode(v string) {
	o.IncomingResponseCode = &v
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *NetworkMetadata) GetProductId() string {
	if o == nil || IsNil(o.ProductId) {
		var ret string
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkMetadata) GetProductIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *NetworkMetadata) HasProductId() bool {
	if o != nil && !IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given string and assigns it to the ProductId field.
func (o *NetworkMetadata) SetProductId(v string) {
	o.ProductId = &v
}

// GetProgramId returns the ProgramId field value if set, zero value otherwise.
func (o *NetworkMetadata) GetProgramId() string {
	if o == nil || IsNil(o.ProgramId) {
		var ret string
		return ret
	}
	return *o.ProgramId
}

// GetProgramIdOk returns a tuple with the ProgramId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkMetadata) GetProgramIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProgramId) {
		return nil, false
	}
	return o.ProgramId, true
}

// HasProgramId returns a boolean if a field has been set.
func (o *NetworkMetadata) HasProgramId() bool {
	if o != nil && !IsNil(o.ProgramId) {
		return true
	}

	return false
}

// SetProgramId gets a reference to the given string and assigns it to the ProgramId field.
func (o *NetworkMetadata) SetProgramId(v string) {
	o.ProgramId = &v
}

// GetSpendQualifier returns the SpendQualifier field value if set, zero value otherwise.
func (o *NetworkMetadata) GetSpendQualifier() string {
	if o == nil || IsNil(o.SpendQualifier) {
		var ret string
		return ret
	}
	return *o.SpendQualifier
}

// GetSpendQualifierOk returns a tuple with the SpendQualifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkMetadata) GetSpendQualifierOk() (*string, bool) {
	if o == nil || IsNil(o.SpendQualifier) {
		return nil, false
	}
	return o.SpendQualifier, true
}

// HasSpendQualifier returns a boolean if a field has been set.
func (o *NetworkMetadata) HasSpendQualifier() bool {
	if o != nil && !IsNil(o.SpendQualifier) {
		return true
	}

	return false
}

// SetSpendQualifier gets a reference to the given string and assigns it to the SpendQualifier field.
func (o *NetworkMetadata) SetSpendQualifier(v string) {
	o.SpendQualifier = &v
}

// GetSurchargeFreeAtmNetwork returns the SurchargeFreeAtmNetwork field value if set, zero value otherwise.
func (o *NetworkMetadata) GetSurchargeFreeAtmNetwork() string {
	if o == nil || IsNil(o.SurchargeFreeAtmNetwork) {
		var ret string
		return ret
	}
	return *o.SurchargeFreeAtmNetwork
}

// GetSurchargeFreeAtmNetworkOk returns a tuple with the SurchargeFreeAtmNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkMetadata) GetSurchargeFreeAtmNetworkOk() (*string, bool) {
	if o == nil || IsNil(o.SurchargeFreeAtmNetwork) {
		return nil, false
	}
	return o.SurchargeFreeAtmNetwork, true
}

// HasSurchargeFreeAtmNetwork returns a boolean if a field has been set.
func (o *NetworkMetadata) HasSurchargeFreeAtmNetwork() bool {
	if o != nil && !IsNil(o.SurchargeFreeAtmNetwork) {
		return true
	}

	return false
}

// SetSurchargeFreeAtmNetwork gets a reference to the given string and assigns it to the SurchargeFreeAtmNetwork field.
func (o *NetworkMetadata) SetSurchargeFreeAtmNetwork(v string) {
	o.SurchargeFreeAtmNetwork = &v
}

func (o NetworkMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountIdentification1) {
		toSerialize["account_identification_1"] = o.AccountIdentification1
	}
	if !IsNil(o.IncomingResponseCode) {
		toSerialize["incoming_response_code"] = o.IncomingResponseCode
	}
	if !IsNil(o.ProductId) {
		toSerialize["product_id"] = o.ProductId
	}
	if !IsNil(o.ProgramId) {
		toSerialize["program_id"] = o.ProgramId
	}
	if !IsNil(o.SpendQualifier) {
		toSerialize["spend_qualifier"] = o.SpendQualifier
	}
	if !IsNil(o.SurchargeFreeAtmNetwork) {
		toSerialize["surcharge_free_atm_network"] = o.SurchargeFreeAtmNetwork
	}
	return toSerialize, nil
}

type NullableNetworkMetadata struct {
	value *NetworkMetadata
	isSet bool
}

func (v NullableNetworkMetadata) Get() *NetworkMetadata {
	return v.value
}

func (v *NullableNetworkMetadata) Set(val *NetworkMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkMetadata(val *NetworkMetadata) *NullableNetworkMetadata {
	return &NullableNetworkMetadata{value: val, isSet: true}
}

func (v NullableNetworkMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


