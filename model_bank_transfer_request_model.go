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

// checks if the BankTransferRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BankTransferRequestModel{}

// BankTransferRequestModel struct for BankTransferRequestModel
type BankTransferRequestModel struct {
	// Amount to push or pull.
	Amount decimal.Decimal `json:"amount"`
	// default = API
	Channel *string `json:"channel,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	// Currency of the push or pull.
	CurrencyCode *string `json:"currency_code,omitempty"`
	// ACH funding source token for the external account.
	FundingSourceToken string `json:"funding_source_token"`
	// Additional text describing the ACH transfer.
	Memo *string `json:"memo,omitempty"`
	// Three-letter code identifying the type of entry.  * *WEB* — An internet-initiated entry * *PPD* — Prearranged Payment and Deposit * *CCD* — Cash Concentration and Disbursement
	StandardEntryClassCode *string `json:"standard_entry_class_code,omitempty"`
	// Description of the transaction, as it will appear on the receiver's bank statement.
	StatementDescriptor *string `json:"statement_descriptor,omitempty"`
	// Unique identifier of the ACH transfer to retrieve.
	Token *string `json:"token,omitempty"`
	// Specifies how quickly to initiate the ACH transfer.  *NOTE:* Same-day transfers are limited to a maximum amount of $100,000.
	TransferSpeed *string `json:"transfer_speed,omitempty"`
	// Specifies whether the ACH transfer is a push (credit) or pull (debit).
	Type string `json:"type"`
}

type _BankTransferRequestModel BankTransferRequestModel

// NewBankTransferRequestModel instantiates a new BankTransferRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankTransferRequestModel(amount decimal.Decimal, fundingSourceToken string, type_ string) *BankTransferRequestModel {
	this := BankTransferRequestModel{}
	this.Amount = amount
	this.FundingSourceToken = fundingSourceToken
	this.Type = type_
	return &this
}

// NewBankTransferRequestModelWithDefaults instantiates a new BankTransferRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankTransferRequestModelWithDefaults() *BankTransferRequestModel {
	this := BankTransferRequestModel{}
	return &this
}

// GetAmount returns the Amount field value
func (o *BankTransferRequestModel) GetAmount() decimal.Decimal {
	if o == nil {
		var ret decimal.Decimal
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *BankTransferRequestModel) GetAmountOk() (*decimal.Decimal, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *BankTransferRequestModel) SetAmount(v decimal.Decimal) {
	o.Amount = v
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *BankTransferRequestModel) GetChannel() string {
	if o == nil || IsNil(o.Channel) {
		var ret string
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransferRequestModel) GetChannelOk() (*string, bool) {
	if o == nil || IsNil(o.Channel) {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *BankTransferRequestModel) HasChannel() bool {
	if o != nil && !IsNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given string and assigns it to the Channel field.
func (o *BankTransferRequestModel) SetChannel(v string) {
	o.Channel = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *BankTransferRequestModel) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransferRequestModel) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *BankTransferRequestModel) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *BankTransferRequestModel) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *BankTransferRequestModel) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransferRequestModel) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *BankTransferRequestModel) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *BankTransferRequestModel) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetFundingSourceToken returns the FundingSourceToken field value
func (o *BankTransferRequestModel) GetFundingSourceToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FundingSourceToken
}

// GetFundingSourceTokenOk returns a tuple with the FundingSourceToken field value
// and a boolean to check if the value has been set.
func (o *BankTransferRequestModel) GetFundingSourceTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FundingSourceToken, true
}

// SetFundingSourceToken sets field value
func (o *BankTransferRequestModel) SetFundingSourceToken(v string) {
	o.FundingSourceToken = v
}

// GetMemo returns the Memo field value if set, zero value otherwise.
func (o *BankTransferRequestModel) GetMemo() string {
	if o == nil || IsNil(o.Memo) {
		var ret string
		return ret
	}
	return *o.Memo
}

// GetMemoOk returns a tuple with the Memo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransferRequestModel) GetMemoOk() (*string, bool) {
	if o == nil || IsNil(o.Memo) {
		return nil, false
	}
	return o.Memo, true
}

// HasMemo returns a boolean if a field has been set.
func (o *BankTransferRequestModel) HasMemo() bool {
	if o != nil && !IsNil(o.Memo) {
		return true
	}

	return false
}

// SetMemo gets a reference to the given string and assigns it to the Memo field.
func (o *BankTransferRequestModel) SetMemo(v string) {
	o.Memo = &v
}

// GetStandardEntryClassCode returns the StandardEntryClassCode field value if set, zero value otherwise.
func (o *BankTransferRequestModel) GetStandardEntryClassCode() string {
	if o == nil || IsNil(o.StandardEntryClassCode) {
		var ret string
		return ret
	}
	return *o.StandardEntryClassCode
}

// GetStandardEntryClassCodeOk returns a tuple with the StandardEntryClassCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransferRequestModel) GetStandardEntryClassCodeOk() (*string, bool) {
	if o == nil || IsNil(o.StandardEntryClassCode) {
		return nil, false
	}
	return o.StandardEntryClassCode, true
}

// HasStandardEntryClassCode returns a boolean if a field has been set.
func (o *BankTransferRequestModel) HasStandardEntryClassCode() bool {
	if o != nil && !IsNil(o.StandardEntryClassCode) {
		return true
	}

	return false
}

// SetStandardEntryClassCode gets a reference to the given string and assigns it to the StandardEntryClassCode field.
func (o *BankTransferRequestModel) SetStandardEntryClassCode(v string) {
	o.StandardEntryClassCode = &v
}

// GetStatementDescriptor returns the StatementDescriptor field value if set, zero value otherwise.
func (o *BankTransferRequestModel) GetStatementDescriptor() string {
	if o == nil || IsNil(o.StatementDescriptor) {
		var ret string
		return ret
	}
	return *o.StatementDescriptor
}

// GetStatementDescriptorOk returns a tuple with the StatementDescriptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransferRequestModel) GetStatementDescriptorOk() (*string, bool) {
	if o == nil || IsNil(o.StatementDescriptor) {
		return nil, false
	}
	return o.StatementDescriptor, true
}

// HasStatementDescriptor returns a boolean if a field has been set.
func (o *BankTransferRequestModel) HasStatementDescriptor() bool {
	if o != nil && !IsNil(o.StatementDescriptor) {
		return true
	}

	return false
}

// SetStatementDescriptor gets a reference to the given string and assigns it to the StatementDescriptor field.
func (o *BankTransferRequestModel) SetStatementDescriptor(v string) {
	o.StatementDescriptor = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *BankTransferRequestModel) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransferRequestModel) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *BankTransferRequestModel) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *BankTransferRequestModel) SetToken(v string) {
	o.Token = &v
}

// GetTransferSpeed returns the TransferSpeed field value if set, zero value otherwise.
func (o *BankTransferRequestModel) GetTransferSpeed() string {
	if o == nil || IsNil(o.TransferSpeed) {
		var ret string
		return ret
	}
	return *o.TransferSpeed
}

// GetTransferSpeedOk returns a tuple with the TransferSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransferRequestModel) GetTransferSpeedOk() (*string, bool) {
	if o == nil || IsNil(o.TransferSpeed) {
		return nil, false
	}
	return o.TransferSpeed, true
}

// HasTransferSpeed returns a boolean if a field has been set.
func (o *BankTransferRequestModel) HasTransferSpeed() bool {
	if o != nil && !IsNil(o.TransferSpeed) {
		return true
	}

	return false
}

// SetTransferSpeed gets a reference to the given string and assigns it to the TransferSpeed field.
func (o *BankTransferRequestModel) SetTransferSpeed(v string) {
	o.TransferSpeed = &v
}

// GetType returns the Type field value
func (o *BankTransferRequestModel) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BankTransferRequestModel) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BankTransferRequestModel) SetType(v string) {
	o.Type = v
}

func (o BankTransferRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BankTransferRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	if !IsNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["created_by"] = o.CreatedBy
	}
	if !IsNil(o.CurrencyCode) {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	toSerialize["funding_source_token"] = o.FundingSourceToken
	if !IsNil(o.Memo) {
		toSerialize["memo"] = o.Memo
	}
	if !IsNil(o.StandardEntryClassCode) {
		toSerialize["standard_entry_class_code"] = o.StandardEntryClassCode
	}
	if !IsNil(o.StatementDescriptor) {
		toSerialize["statement_descriptor"] = o.StatementDescriptor
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.TransferSpeed) {
		toSerialize["transfer_speed"] = o.TransferSpeed
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *BankTransferRequestModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amount",
		"funding_source_token",
		"type",
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

	varBankTransferRequestModel := _BankTransferRequestModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBankTransferRequestModel)

	if err != nil {
		return err
	}

	*o = BankTransferRequestModel(varBankTransferRequestModel)

	return err
}

type NullableBankTransferRequestModel struct {
	value *BankTransferRequestModel
	isSet bool
}

func (v NullableBankTransferRequestModel) Get() *BankTransferRequestModel {
	return v.value
}

func (v *NullableBankTransferRequestModel) Set(val *BankTransferRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableBankTransferRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableBankTransferRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankTransferRequestModel(val *BankTransferRequestModel) *NullableBankTransferRequestModel {
	return &NullableBankTransferRequestModel{value: val, isSet: true}
}

func (v NullableBankTransferRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankTransferRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


