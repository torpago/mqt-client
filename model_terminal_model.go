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

// checks if the TerminalModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TerminalModel{}

// TerminalModel Contains information about the point of sale, including details on how the card was presented.  Returned if provided by the card network, and the request uses Transaction Model v1 of the Marqeta Core API. Not returned for Transaction Model v2 requests.
type TerminalModel struct {
	// Indicates whether the card was present during the transaction.
	CardPresence *string `json:"card_presence,omitempty"`
	// Indicates whether the cardholder was present during the transaction.
	CardholderPresence *string `json:"cardholder_presence,omitempty"`
	// Indicates whether the card acceptor or terminal supports partial-approval transactions.
	PartialApprovalCapable *string `json:"partial_approval_capable,omitempty"`
	// Indicates whether the cardholder entered a PIN during the transaction.
	PinPresent *string `json:"pin_present,omitempty"`
	// Indicates a higher-risk operation, such as a quasi-cash or cryptocurrency transaction.  These transactions typically involve non-financial institutions.
	SpecialConditionIndicator *string `json:"special_condition_indicator,omitempty"`
	// Card acceptor or terminal identification number.
	Tid *string `json:"tid,omitempty"`
	// Specifies whether the transaction was initiated by a cardholder or a merchant.
	TransactionInitiatedBy *string `json:"transaction_initiated_by,omitempty"`
}

// NewTerminalModel instantiates a new TerminalModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerminalModel() *TerminalModel {
	this := TerminalModel{}
	return &this
}

// NewTerminalModelWithDefaults instantiates a new TerminalModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerminalModelWithDefaults() *TerminalModel {
	this := TerminalModel{}
	return &this
}

// GetCardPresence returns the CardPresence field value if set, zero value otherwise.
func (o *TerminalModel) GetCardPresence() string {
	if o == nil || IsNil(o.CardPresence) {
		var ret string
		return ret
	}
	return *o.CardPresence
}

// GetCardPresenceOk returns a tuple with the CardPresence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalModel) GetCardPresenceOk() (*string, bool) {
	if o == nil || IsNil(o.CardPresence) {
		return nil, false
	}
	return o.CardPresence, true
}

// HasCardPresence returns a boolean if a field has been set.
func (o *TerminalModel) HasCardPresence() bool {
	if o != nil && !IsNil(o.CardPresence) {
		return true
	}

	return false
}

// SetCardPresence gets a reference to the given string and assigns it to the CardPresence field.
func (o *TerminalModel) SetCardPresence(v string) {
	o.CardPresence = &v
}

// GetCardholderPresence returns the CardholderPresence field value if set, zero value otherwise.
func (o *TerminalModel) GetCardholderPresence() string {
	if o == nil || IsNil(o.CardholderPresence) {
		var ret string
		return ret
	}
	return *o.CardholderPresence
}

// GetCardholderPresenceOk returns a tuple with the CardholderPresence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalModel) GetCardholderPresenceOk() (*string, bool) {
	if o == nil || IsNil(o.CardholderPresence) {
		return nil, false
	}
	return o.CardholderPresence, true
}

// HasCardholderPresence returns a boolean if a field has been set.
func (o *TerminalModel) HasCardholderPresence() bool {
	if o != nil && !IsNil(o.CardholderPresence) {
		return true
	}

	return false
}

// SetCardholderPresence gets a reference to the given string and assigns it to the CardholderPresence field.
func (o *TerminalModel) SetCardholderPresence(v string) {
	o.CardholderPresence = &v
}

// GetPartialApprovalCapable returns the PartialApprovalCapable field value if set, zero value otherwise.
func (o *TerminalModel) GetPartialApprovalCapable() string {
	if o == nil || IsNil(o.PartialApprovalCapable) {
		var ret string
		return ret
	}
	return *o.PartialApprovalCapable
}

// GetPartialApprovalCapableOk returns a tuple with the PartialApprovalCapable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalModel) GetPartialApprovalCapableOk() (*string, bool) {
	if o == nil || IsNil(o.PartialApprovalCapable) {
		return nil, false
	}
	return o.PartialApprovalCapable, true
}

// HasPartialApprovalCapable returns a boolean if a field has been set.
func (o *TerminalModel) HasPartialApprovalCapable() bool {
	if o != nil && !IsNil(o.PartialApprovalCapable) {
		return true
	}

	return false
}

// SetPartialApprovalCapable gets a reference to the given string and assigns it to the PartialApprovalCapable field.
func (o *TerminalModel) SetPartialApprovalCapable(v string) {
	o.PartialApprovalCapable = &v
}

// GetPinPresent returns the PinPresent field value if set, zero value otherwise.
func (o *TerminalModel) GetPinPresent() string {
	if o == nil || IsNil(o.PinPresent) {
		var ret string
		return ret
	}
	return *o.PinPresent
}

// GetPinPresentOk returns a tuple with the PinPresent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalModel) GetPinPresentOk() (*string, bool) {
	if o == nil || IsNil(o.PinPresent) {
		return nil, false
	}
	return o.PinPresent, true
}

// HasPinPresent returns a boolean if a field has been set.
func (o *TerminalModel) HasPinPresent() bool {
	if o != nil && !IsNil(o.PinPresent) {
		return true
	}

	return false
}

// SetPinPresent gets a reference to the given string and assigns it to the PinPresent field.
func (o *TerminalModel) SetPinPresent(v string) {
	o.PinPresent = &v
}

// GetSpecialConditionIndicator returns the SpecialConditionIndicator field value if set, zero value otherwise.
func (o *TerminalModel) GetSpecialConditionIndicator() string {
	if o == nil || IsNil(o.SpecialConditionIndicator) {
		var ret string
		return ret
	}
	return *o.SpecialConditionIndicator
}

// GetSpecialConditionIndicatorOk returns a tuple with the SpecialConditionIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalModel) GetSpecialConditionIndicatorOk() (*string, bool) {
	if o == nil || IsNil(o.SpecialConditionIndicator) {
		return nil, false
	}
	return o.SpecialConditionIndicator, true
}

// HasSpecialConditionIndicator returns a boolean if a field has been set.
func (o *TerminalModel) HasSpecialConditionIndicator() bool {
	if o != nil && !IsNil(o.SpecialConditionIndicator) {
		return true
	}

	return false
}

// SetSpecialConditionIndicator gets a reference to the given string and assigns it to the SpecialConditionIndicator field.
func (o *TerminalModel) SetSpecialConditionIndicator(v string) {
	o.SpecialConditionIndicator = &v
}

// GetTid returns the Tid field value if set, zero value otherwise.
func (o *TerminalModel) GetTid() string {
	if o == nil || IsNil(o.Tid) {
		var ret string
		return ret
	}
	return *o.Tid
}

// GetTidOk returns a tuple with the Tid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalModel) GetTidOk() (*string, bool) {
	if o == nil || IsNil(o.Tid) {
		return nil, false
	}
	return o.Tid, true
}

// HasTid returns a boolean if a field has been set.
func (o *TerminalModel) HasTid() bool {
	if o != nil && !IsNil(o.Tid) {
		return true
	}

	return false
}

// SetTid gets a reference to the given string and assigns it to the Tid field.
func (o *TerminalModel) SetTid(v string) {
	o.Tid = &v
}

// GetTransactionInitiatedBy returns the TransactionInitiatedBy field value if set, zero value otherwise.
func (o *TerminalModel) GetTransactionInitiatedBy() string {
	if o == nil || IsNil(o.TransactionInitiatedBy) {
		var ret string
		return ret
	}
	return *o.TransactionInitiatedBy
}

// GetTransactionInitiatedByOk returns a tuple with the TransactionInitiatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalModel) GetTransactionInitiatedByOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionInitiatedBy) {
		return nil, false
	}
	return o.TransactionInitiatedBy, true
}

// HasTransactionInitiatedBy returns a boolean if a field has been set.
func (o *TerminalModel) HasTransactionInitiatedBy() bool {
	if o != nil && !IsNil(o.TransactionInitiatedBy) {
		return true
	}

	return false
}

// SetTransactionInitiatedBy gets a reference to the given string and assigns it to the TransactionInitiatedBy field.
func (o *TerminalModel) SetTransactionInitiatedBy(v string) {
	o.TransactionInitiatedBy = &v
}

func (o TerminalModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TerminalModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CardPresence) {
		toSerialize["card_presence"] = o.CardPresence
	}
	if !IsNil(o.CardholderPresence) {
		toSerialize["cardholder_presence"] = o.CardholderPresence
	}
	if !IsNil(o.PartialApprovalCapable) {
		toSerialize["partial_approval_capable"] = o.PartialApprovalCapable
	}
	if !IsNil(o.PinPresent) {
		toSerialize["pin_present"] = o.PinPresent
	}
	if !IsNil(o.SpecialConditionIndicator) {
		toSerialize["special_condition_indicator"] = o.SpecialConditionIndicator
	}
	if !IsNil(o.Tid) {
		toSerialize["tid"] = o.Tid
	}
	if !IsNil(o.TransactionInitiatedBy) {
		toSerialize["transaction_initiated_by"] = o.TransactionInitiatedBy
	}
	return toSerialize, nil
}

type NullableTerminalModel struct {
	value *TerminalModel
	isSet bool
}

func (v NullableTerminalModel) Get() *TerminalModel {
	return v.value
}

func (v *NullableTerminalModel) Set(val *TerminalModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTerminalModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTerminalModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerminalModel(val *TerminalModel) *NullableTerminalModel {
	return &NullableTerminalModel{value: val, isSet: true}
}

func (v NullableTerminalModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerminalModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


