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
)

// checks if the AccountConfigResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountConfigResponse{}

// AccountConfigResponse Contains information returned when configuring an account's billing cycle day, payment due day, fees, and more.
type AccountConfigResponse struct {
	// Day of month the billing cycle starts.  If an override value is not provided, the default value is derived from the bundle.
	BillingCycleDay int32 `json:"billing_cycle_day"`
	// Level of the credit card.
	CardLevel string `json:"card_level"`
	// A value of `true` indicates that the account holder consents to receiving disclosures and statements electronically.
	EDisclosureActive bool `json:"e_disclosure_active"`
	// Contains one or more fees associated with the credit account.
	Fees []ConfigFeeScheduleResponse `json:"fees,omitempty"`
	MinPayment *AccountConfigMinPayment `json:"min_payment,omitempty"`
	// Day of month the payment for the previous billing cycle is due.
	// Deprecated
	PaymentDueDay *int32 `json:"payment_due_day,omitempty"`
	PaymentHolds AccountConfigPaymentHolds `json:"payment_holds"`
	// Contains one or more rewards associated with the credit account.
	Rewards []AccountReward `json:"rewards,omitempty"`
}

type _AccountConfigResponse AccountConfigResponse

// NewAccountConfigResponse instantiates a new AccountConfigResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountConfigResponse(billingCycleDay int32, cardLevel string, eDisclosureActive bool, paymentHolds AccountConfigPaymentHolds) *AccountConfigResponse {
	this := AccountConfigResponse{}
	this.BillingCycleDay = billingCycleDay
	this.CardLevel = cardLevel
	this.EDisclosureActive = eDisclosureActive
	this.PaymentHolds = paymentHolds
	return &this
}

// NewAccountConfigResponseWithDefaults instantiates a new AccountConfigResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountConfigResponseWithDefaults() *AccountConfigResponse {
	this := AccountConfigResponse{}
	var cardLevel string = "NA"
	this.CardLevel = cardLevel
	var eDisclosureActive bool = false
	this.EDisclosureActive = eDisclosureActive
	return &this
}

// GetBillingCycleDay returns the BillingCycleDay field value
func (o *AccountConfigResponse) GetBillingCycleDay() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BillingCycleDay
}

// GetBillingCycleDayOk returns a tuple with the BillingCycleDay field value
// and a boolean to check if the value has been set.
func (o *AccountConfigResponse) GetBillingCycleDayOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BillingCycleDay, true
}

// SetBillingCycleDay sets field value
func (o *AccountConfigResponse) SetBillingCycleDay(v int32) {
	o.BillingCycleDay = v
}

// GetCardLevel returns the CardLevel field value
func (o *AccountConfigResponse) GetCardLevel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CardLevel
}

// GetCardLevelOk returns a tuple with the CardLevel field value
// and a boolean to check if the value has been set.
func (o *AccountConfigResponse) GetCardLevelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardLevel, true
}

// SetCardLevel sets field value
func (o *AccountConfigResponse) SetCardLevel(v string) {
	o.CardLevel = v
}

// GetEDisclosureActive returns the EDisclosureActive field value
func (o *AccountConfigResponse) GetEDisclosureActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EDisclosureActive
}

// GetEDisclosureActiveOk returns a tuple with the EDisclosureActive field value
// and a boolean to check if the value has been set.
func (o *AccountConfigResponse) GetEDisclosureActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EDisclosureActive, true
}

// SetEDisclosureActive sets field value
func (o *AccountConfigResponse) SetEDisclosureActive(v bool) {
	o.EDisclosureActive = v
}

// GetFees returns the Fees field value if set, zero value otherwise.
func (o *AccountConfigResponse) GetFees() []ConfigFeeScheduleResponse {
	if o == nil || IsNil(o.Fees) {
		var ret []ConfigFeeScheduleResponse
		return ret
	}
	return o.Fees
}

// GetFeesOk returns a tuple with the Fees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountConfigResponse) GetFeesOk() ([]ConfigFeeScheduleResponse, bool) {
	if o == nil || IsNil(o.Fees) {
		return nil, false
	}
	return o.Fees, true
}

// HasFees returns a boolean if a field has been set.
func (o *AccountConfigResponse) HasFees() bool {
	if o != nil && !IsNil(o.Fees) {
		return true
	}

	return false
}

// SetFees gets a reference to the given []ConfigFeeScheduleResponse and assigns it to the Fees field.
func (o *AccountConfigResponse) SetFees(v []ConfigFeeScheduleResponse) {
	o.Fees = v
}

// GetMinPayment returns the MinPayment field value if set, zero value otherwise.
func (o *AccountConfigResponse) GetMinPayment() AccountConfigMinPayment {
	if o == nil || IsNil(o.MinPayment) {
		var ret AccountConfigMinPayment
		return ret
	}
	return *o.MinPayment
}

// GetMinPaymentOk returns a tuple with the MinPayment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountConfigResponse) GetMinPaymentOk() (*AccountConfigMinPayment, bool) {
	if o == nil || IsNil(o.MinPayment) {
		return nil, false
	}
	return o.MinPayment, true
}

// HasMinPayment returns a boolean if a field has been set.
func (o *AccountConfigResponse) HasMinPayment() bool {
	if o != nil && !IsNil(o.MinPayment) {
		return true
	}

	return false
}

// SetMinPayment gets a reference to the given AccountConfigMinPayment and assigns it to the MinPayment field.
func (o *AccountConfigResponse) SetMinPayment(v AccountConfigMinPayment) {
	o.MinPayment = &v
}

// GetPaymentDueDay returns the PaymentDueDay field value if set, zero value otherwise.
// Deprecated
func (o *AccountConfigResponse) GetPaymentDueDay() int32 {
	if o == nil || IsNil(o.PaymentDueDay) {
		var ret int32
		return ret
	}
	return *o.PaymentDueDay
}

// GetPaymentDueDayOk returns a tuple with the PaymentDueDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *AccountConfigResponse) GetPaymentDueDayOk() (*int32, bool) {
	if o == nil || IsNil(o.PaymentDueDay) {
		return nil, false
	}
	return o.PaymentDueDay, true
}

// HasPaymentDueDay returns a boolean if a field has been set.
func (o *AccountConfigResponse) HasPaymentDueDay() bool {
	if o != nil && !IsNil(o.PaymentDueDay) {
		return true
	}

	return false
}

// SetPaymentDueDay gets a reference to the given int32 and assigns it to the PaymentDueDay field.
// Deprecated
func (o *AccountConfigResponse) SetPaymentDueDay(v int32) {
	o.PaymentDueDay = &v
}

// GetPaymentHolds returns the PaymentHolds field value
func (o *AccountConfigResponse) GetPaymentHolds() AccountConfigPaymentHolds {
	if o == nil {
		var ret AccountConfigPaymentHolds
		return ret
	}

	return o.PaymentHolds
}

// GetPaymentHoldsOk returns a tuple with the PaymentHolds field value
// and a boolean to check if the value has been set.
func (o *AccountConfigResponse) GetPaymentHoldsOk() (*AccountConfigPaymentHolds, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentHolds, true
}

// SetPaymentHolds sets field value
func (o *AccountConfigResponse) SetPaymentHolds(v AccountConfigPaymentHolds) {
	o.PaymentHolds = v
}

// GetRewards returns the Rewards field value if set, zero value otherwise.
func (o *AccountConfigResponse) GetRewards() []AccountReward {
	if o == nil || IsNil(o.Rewards) {
		var ret []AccountReward
		return ret
	}
	return o.Rewards
}

// GetRewardsOk returns a tuple with the Rewards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountConfigResponse) GetRewardsOk() ([]AccountReward, bool) {
	if o == nil || IsNil(o.Rewards) {
		return nil, false
	}
	return o.Rewards, true
}

// HasRewards returns a boolean if a field has been set.
func (o *AccountConfigResponse) HasRewards() bool {
	if o != nil && !IsNil(o.Rewards) {
		return true
	}

	return false
}

// SetRewards gets a reference to the given []AccountReward and assigns it to the Rewards field.
func (o *AccountConfigResponse) SetRewards(v []AccountReward) {
	o.Rewards = v
}

func (o AccountConfigResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountConfigResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["billing_cycle_day"] = o.BillingCycleDay
	toSerialize["card_level"] = o.CardLevel
	toSerialize["e_disclosure_active"] = o.EDisclosureActive
	if !IsNil(o.Fees) {
		toSerialize["fees"] = o.Fees
	}
	if !IsNil(o.MinPayment) {
		toSerialize["min_payment"] = o.MinPayment
	}
	if !IsNil(o.PaymentDueDay) {
		toSerialize["payment_due_day"] = o.PaymentDueDay
	}
	toSerialize["payment_holds"] = o.PaymentHolds
	if !IsNil(o.Rewards) {
		toSerialize["rewards"] = o.Rewards
	}
	return toSerialize, nil
}

func (o *AccountConfigResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"billing_cycle_day",
		"card_level",
		"e_disclosure_active",
		"payment_holds",
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

	varAccountConfigResponse := _AccountConfigResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAccountConfigResponse)

	if err != nil {
		return err
	}

	*o = AccountConfigResponse(varAccountConfigResponse)

	return err
}

type NullableAccountConfigResponse struct {
	value *AccountConfigResponse
	isSet bool
}

func (v NullableAccountConfigResponse) Get() *AccountConfigResponse {
	return v.value
}

func (v *NullableAccountConfigResponse) Set(val *AccountConfigResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountConfigResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountConfigResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountConfigResponse(val *AccountConfigResponse) *NullableAccountConfigResponse {
	return &NullableAccountConfigResponse{value: val, isSet: true}
}

func (v NullableAccountConfigResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountConfigResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


