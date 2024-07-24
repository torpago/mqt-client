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
	"time"
	"bytes"
	"fmt"
)

// checks if the BankTransferTransitionResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BankTransferTransitionResponseModel{}

// BankTransferTransitionResponseModel struct for BankTransferTransitionResponseModel
type BankTransferTransitionResponseModel struct {
	Amount *float32 `json:"amount,omitempty"`
	// Unique identifier of the ACH transfer being transitioned.
	BankTransferToken string `json:"bank_transfer_token"`
	// Field required in older versions of the API, but no longer used.
	BatchNumber *string `json:"batch_number,omitempty"`
	// Mechanism by which the transaction was initiated.
	Channel string `json:"channel"`
	// Date and time when the ACH transfer was created, in UTC.
	CreatedTime *time.Time `json:"created_time,omitempty"`
	EffectiveDate *time.Time `json:"effective_date,omitempty"`
	// Date and time when the ACH transfer was last modified.
	LastModifiedTime *time.Time `json:"last_modified_time,omitempty"`
	// Not currently used.
	ProgramReserveToken *string `json:"program_reserve_token,omitempty"`
	// Explanation of why the transition is being made, for example \"Created by POST call on API\". Returned if this information was supplied when the transition was originally created.
	Reason *string `json:"reason,omitempty"`
	// Standardized ACH return code for a returned transaction, generally sent by the RDFI.  Transactions can be returned for any of the reasons listed in the <</developer-guides/ach-origination#_nacha_ach_return_codes, NACHA ACH return codes table>> of the ACH Origination Guide.
	ReturnCode *string `json:"return_code,omitempty"`
	// Human-readable description correlating to the `return_code`, such as `Insufficient funds`, if a return code is present in the response.
	ReturnReason *string `json:"return_reason,omitempty"`
	ReversalAfter45Days *bool `json:"reversal_after_45_days,omitempty"`
	// New state of the ACH transfer.
	Status string `json:"status"`
	// Unique identifier of the bank transfer transition.
	Token *string `json:"token,omitempty"`
	// Transaction token for JIT-related ledger entries for the ACH transfer.
	TransactionJitToken *string `json:"transaction_jit_token,omitempty"`
	// Transaction token for *non*-JIT-related ledger entries for the ACH transfer.
	TransactionToken *string `json:"transaction_token,omitempty"`
}

type _BankTransferTransitionResponseModel BankTransferTransitionResponseModel

// NewBankTransferTransitionResponseModel instantiates a new BankTransferTransitionResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankTransferTransitionResponseModel(bankTransferToken string, channel string, status string) *BankTransferTransitionResponseModel {
	this := BankTransferTransitionResponseModel{}
	this.BankTransferToken = bankTransferToken
	this.Channel = channel
	this.Status = status
	return &this
}

// NewBankTransferTransitionResponseModelWithDefaults instantiates a new BankTransferTransitionResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankTransferTransitionResponseModelWithDefaults() *BankTransferTransitionResponseModel {
	this := BankTransferTransitionResponseModel{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *BankTransferTransitionResponseModel) GetAmount() float32 {
	if o == nil || IsNil(o.Amount) {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransferTransitionResponseModel) GetAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *BankTransferTransitionResponseModel) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *BankTransferTransitionResponseModel) SetAmount(v float32) {
	o.Amount = &v
}

// GetBankTransferToken returns the BankTransferToken field value
func (o *BankTransferTransitionResponseModel) GetBankTransferToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BankTransferToken
}

// GetBankTransferTokenOk returns a tuple with the BankTransferToken field value
// and a boolean to check if the value has been set.
func (o *BankTransferTransitionResponseModel) GetBankTransferTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BankTransferToken, true
}

// SetBankTransferToken sets field value
func (o *BankTransferTransitionResponseModel) SetBankTransferToken(v string) {
	o.BankTransferToken = v
}

// GetBatchNumber returns the BatchNumber field value if set, zero value otherwise.
func (o *BankTransferTransitionResponseModel) GetBatchNumber() string {
	if o == nil || IsNil(o.BatchNumber) {
		var ret string
		return ret
	}
	return *o.BatchNumber
}

// GetBatchNumberOk returns a tuple with the BatchNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransferTransitionResponseModel) GetBatchNumberOk() (*string, bool) {
	if o == nil || IsNil(o.BatchNumber) {
		return nil, false
	}
	return o.BatchNumber, true
}

// HasBatchNumber returns a boolean if a field has been set.
func (o *BankTransferTransitionResponseModel) HasBatchNumber() bool {
	if o != nil && !IsNil(o.BatchNumber) {
		return true
	}

	return false
}

// SetBatchNumber gets a reference to the given string and assigns it to the BatchNumber field.
func (o *BankTransferTransitionResponseModel) SetBatchNumber(v string) {
	o.BatchNumber = &v
}

// GetChannel returns the Channel field value
func (o *BankTransferTransitionResponseModel) GetChannel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Channel
}

// GetChannelOk returns a tuple with the Channel field value
// and a boolean to check if the value has been set.
func (o *BankTransferTransitionResponseModel) GetChannelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Channel, true
}

// SetChannel sets field value
func (o *BankTransferTransitionResponseModel) SetChannel(v string) {
	o.Channel = v
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *BankTransferTransitionResponseModel) GetCreatedTime() time.Time {
	if o == nil || IsNil(o.CreatedTime) {
		var ret time.Time
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransferTransitionResponseModel) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedTime) {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *BankTransferTransitionResponseModel) HasCreatedTime() bool {
	if o != nil && !IsNil(o.CreatedTime) {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given time.Time and assigns it to the CreatedTime field.
func (o *BankTransferTransitionResponseModel) SetCreatedTime(v time.Time) {
	o.CreatedTime = &v
}

// GetEffectiveDate returns the EffectiveDate field value if set, zero value otherwise.
func (o *BankTransferTransitionResponseModel) GetEffectiveDate() time.Time {
	if o == nil || IsNil(o.EffectiveDate) {
		var ret time.Time
		return ret
	}
	return *o.EffectiveDate
}

// GetEffectiveDateOk returns a tuple with the EffectiveDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransferTransitionResponseModel) GetEffectiveDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EffectiveDate) {
		return nil, false
	}
	return o.EffectiveDate, true
}

// HasEffectiveDate returns a boolean if a field has been set.
func (o *BankTransferTransitionResponseModel) HasEffectiveDate() bool {
	if o != nil && !IsNil(o.EffectiveDate) {
		return true
	}

	return false
}

// SetEffectiveDate gets a reference to the given time.Time and assigns it to the EffectiveDate field.
func (o *BankTransferTransitionResponseModel) SetEffectiveDate(v time.Time) {
	o.EffectiveDate = &v
}

// GetLastModifiedTime returns the LastModifiedTime field value if set, zero value otherwise.
func (o *BankTransferTransitionResponseModel) GetLastModifiedTime() time.Time {
	if o == nil || IsNil(o.LastModifiedTime) {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransferTransitionResponseModel) GetLastModifiedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModifiedTime) {
		return nil, false
	}
	return o.LastModifiedTime, true
}

// HasLastModifiedTime returns a boolean if a field has been set.
func (o *BankTransferTransitionResponseModel) HasLastModifiedTime() bool {
	if o != nil && !IsNil(o.LastModifiedTime) {
		return true
	}

	return false
}

// SetLastModifiedTime gets a reference to the given time.Time and assigns it to the LastModifiedTime field.
func (o *BankTransferTransitionResponseModel) SetLastModifiedTime(v time.Time) {
	o.LastModifiedTime = &v
}

// GetProgramReserveToken returns the ProgramReserveToken field value if set, zero value otherwise.
func (o *BankTransferTransitionResponseModel) GetProgramReserveToken() string {
	if o == nil || IsNil(o.ProgramReserveToken) {
		var ret string
		return ret
	}
	return *o.ProgramReserveToken
}

// GetProgramReserveTokenOk returns a tuple with the ProgramReserveToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransferTransitionResponseModel) GetProgramReserveTokenOk() (*string, bool) {
	if o == nil || IsNil(o.ProgramReserveToken) {
		return nil, false
	}
	return o.ProgramReserveToken, true
}

// HasProgramReserveToken returns a boolean if a field has been set.
func (o *BankTransferTransitionResponseModel) HasProgramReserveToken() bool {
	if o != nil && !IsNil(o.ProgramReserveToken) {
		return true
	}

	return false
}

// SetProgramReserveToken gets a reference to the given string and assigns it to the ProgramReserveToken field.
func (o *BankTransferTransitionResponseModel) SetProgramReserveToken(v string) {
	o.ProgramReserveToken = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *BankTransferTransitionResponseModel) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransferTransitionResponseModel) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *BankTransferTransitionResponseModel) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *BankTransferTransitionResponseModel) SetReason(v string) {
	o.Reason = &v
}

// GetReturnCode returns the ReturnCode field value if set, zero value otherwise.
func (o *BankTransferTransitionResponseModel) GetReturnCode() string {
	if o == nil || IsNil(o.ReturnCode) {
		var ret string
		return ret
	}
	return *o.ReturnCode
}

// GetReturnCodeOk returns a tuple with the ReturnCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransferTransitionResponseModel) GetReturnCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ReturnCode) {
		return nil, false
	}
	return o.ReturnCode, true
}

// HasReturnCode returns a boolean if a field has been set.
func (o *BankTransferTransitionResponseModel) HasReturnCode() bool {
	if o != nil && !IsNil(o.ReturnCode) {
		return true
	}

	return false
}

// SetReturnCode gets a reference to the given string and assigns it to the ReturnCode field.
func (o *BankTransferTransitionResponseModel) SetReturnCode(v string) {
	o.ReturnCode = &v
}

// GetReturnReason returns the ReturnReason field value if set, zero value otherwise.
func (o *BankTransferTransitionResponseModel) GetReturnReason() string {
	if o == nil || IsNil(o.ReturnReason) {
		var ret string
		return ret
	}
	return *o.ReturnReason
}

// GetReturnReasonOk returns a tuple with the ReturnReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransferTransitionResponseModel) GetReturnReasonOk() (*string, bool) {
	if o == nil || IsNil(o.ReturnReason) {
		return nil, false
	}
	return o.ReturnReason, true
}

// HasReturnReason returns a boolean if a field has been set.
func (o *BankTransferTransitionResponseModel) HasReturnReason() bool {
	if o != nil && !IsNil(o.ReturnReason) {
		return true
	}

	return false
}

// SetReturnReason gets a reference to the given string and assigns it to the ReturnReason field.
func (o *BankTransferTransitionResponseModel) SetReturnReason(v string) {
	o.ReturnReason = &v
}

// GetReversalAfter45Days returns the ReversalAfter45Days field value if set, zero value otherwise.
func (o *BankTransferTransitionResponseModel) GetReversalAfter45Days() bool {
	if o == nil || IsNil(o.ReversalAfter45Days) {
		var ret bool
		return ret
	}
	return *o.ReversalAfter45Days
}

// GetReversalAfter45DaysOk returns a tuple with the ReversalAfter45Days field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransferTransitionResponseModel) GetReversalAfter45DaysOk() (*bool, bool) {
	if o == nil || IsNil(o.ReversalAfter45Days) {
		return nil, false
	}
	return o.ReversalAfter45Days, true
}

// HasReversalAfter45Days returns a boolean if a field has been set.
func (o *BankTransferTransitionResponseModel) HasReversalAfter45Days() bool {
	if o != nil && !IsNil(o.ReversalAfter45Days) {
		return true
	}

	return false
}

// SetReversalAfter45Days gets a reference to the given bool and assigns it to the ReversalAfter45Days field.
func (o *BankTransferTransitionResponseModel) SetReversalAfter45Days(v bool) {
	o.ReversalAfter45Days = &v
}

// GetStatus returns the Status field value
func (o *BankTransferTransitionResponseModel) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *BankTransferTransitionResponseModel) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *BankTransferTransitionResponseModel) SetStatus(v string) {
	o.Status = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *BankTransferTransitionResponseModel) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransferTransitionResponseModel) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *BankTransferTransitionResponseModel) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *BankTransferTransitionResponseModel) SetToken(v string) {
	o.Token = &v
}

// GetTransactionJitToken returns the TransactionJitToken field value if set, zero value otherwise.
func (o *BankTransferTransitionResponseModel) GetTransactionJitToken() string {
	if o == nil || IsNil(o.TransactionJitToken) {
		var ret string
		return ret
	}
	return *o.TransactionJitToken
}

// GetTransactionJitTokenOk returns a tuple with the TransactionJitToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransferTransitionResponseModel) GetTransactionJitTokenOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionJitToken) {
		return nil, false
	}
	return o.TransactionJitToken, true
}

// HasTransactionJitToken returns a boolean if a field has been set.
func (o *BankTransferTransitionResponseModel) HasTransactionJitToken() bool {
	if o != nil && !IsNil(o.TransactionJitToken) {
		return true
	}

	return false
}

// SetTransactionJitToken gets a reference to the given string and assigns it to the TransactionJitToken field.
func (o *BankTransferTransitionResponseModel) SetTransactionJitToken(v string) {
	o.TransactionJitToken = &v
}

// GetTransactionToken returns the TransactionToken field value if set, zero value otherwise.
func (o *BankTransferTransitionResponseModel) GetTransactionToken() string {
	if o == nil || IsNil(o.TransactionToken) {
		var ret string
		return ret
	}
	return *o.TransactionToken
}

// GetTransactionTokenOk returns a tuple with the TransactionToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransferTransitionResponseModel) GetTransactionTokenOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionToken) {
		return nil, false
	}
	return o.TransactionToken, true
}

// HasTransactionToken returns a boolean if a field has been set.
func (o *BankTransferTransitionResponseModel) HasTransactionToken() bool {
	if o != nil && !IsNil(o.TransactionToken) {
		return true
	}

	return false
}

// SetTransactionToken gets a reference to the given string and assigns it to the TransactionToken field.
func (o *BankTransferTransitionResponseModel) SetTransactionToken(v string) {
	o.TransactionToken = &v
}

func (o BankTransferTransitionResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BankTransferTransitionResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	toSerialize["bank_transfer_token"] = o.BankTransferToken
	if !IsNil(o.BatchNumber) {
		toSerialize["batch_number"] = o.BatchNumber
	}
	toSerialize["channel"] = o.Channel
	if !IsNil(o.CreatedTime) {
		toSerialize["created_time"] = o.CreatedTime
	}
	if !IsNil(o.EffectiveDate) {
		toSerialize["effective_date"] = o.EffectiveDate
	}
	if !IsNil(o.LastModifiedTime) {
		toSerialize["last_modified_time"] = o.LastModifiedTime
	}
	if !IsNil(o.ProgramReserveToken) {
		toSerialize["program_reserve_token"] = o.ProgramReserveToken
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !IsNil(o.ReturnCode) {
		toSerialize["return_code"] = o.ReturnCode
	}
	if !IsNil(o.ReturnReason) {
		toSerialize["return_reason"] = o.ReturnReason
	}
	if !IsNil(o.ReversalAfter45Days) {
		toSerialize["reversal_after_45_days"] = o.ReversalAfter45Days
	}
	toSerialize["status"] = o.Status
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.TransactionJitToken) {
		toSerialize["transaction_jit_token"] = o.TransactionJitToken
	}
	if !IsNil(o.TransactionToken) {
		toSerialize["transaction_token"] = o.TransactionToken
	}
	return toSerialize, nil
}

func (o *BankTransferTransitionResponseModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"bank_transfer_token",
		"channel",
		"status",
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

	varBankTransferTransitionResponseModel := _BankTransferTransitionResponseModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBankTransferTransitionResponseModel)

	if err != nil {
		return err
	}

	*o = BankTransferTransitionResponseModel(varBankTransferTransitionResponseModel)

	return err
}

type NullableBankTransferTransitionResponseModel struct {
	value *BankTransferTransitionResponseModel
	isSet bool
}

func (v NullableBankTransferTransitionResponseModel) Get() *BankTransferTransitionResponseModel {
	return v.value
}

func (v *NullableBankTransferTransitionResponseModel) Set(val *BankTransferTransitionResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableBankTransferTransitionResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableBankTransferTransitionResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankTransferTransitionResponseModel(val *BankTransferTransitionResponseModel) *NullableBankTransferTransitionResponseModel {
	return &NullableBankTransferTransitionResponseModel{value: val, isSet: true}
}

func (v NullableBankTransferTransitionResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankTransferTransitionResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


