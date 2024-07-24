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
	"time"
)

// checks if the BankAccountFundingSourceModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BankAccountFundingSourceModel{}

// BankAccountFundingSourceModel struct for BankAccountFundingSourceModel
type BankAccountFundingSourceModel struct {
	FundingSourceModel
	AccountSuffix string `json:"account_suffix"`
	AccountType string `json:"account_type"`
	// Required if 'user_token' is null
	BusinessToken *string `json:"business_token,omitempty"`
	NameOnAccount string `json:"name_on_account"`
	RoutingNumber string `json:"routing_number"`
	// Required if 'business_token' is null
	UserToken *string `json:"user_token,omitempty"`
	VerificationStatus string `json:"verification_status"`
}

type _BankAccountFundingSourceModel BankAccountFundingSourceModel

// NewBankAccountFundingSourceModel instantiates a new BankAccountFundingSourceModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankAccountFundingSourceModel(accountSuffix string, accountType string, nameOnAccount string, routingNumber string, verificationStatus string, active bool, createdTime time.Time, isDefaultAccount bool, lastModifiedTime time.Time, token string, type_ string) *BankAccountFundingSourceModel {
	this := BankAccountFundingSourceModel{}
	this.Active = active
	this.CreatedTime = createdTime
	this.IsDefaultAccount = isDefaultAccount
	this.LastModifiedTime = lastModifiedTime
	this.Token = token
	this.Type = type_
	this.AccountSuffix = accountSuffix
	this.AccountType = accountType
	this.NameOnAccount = nameOnAccount
	this.RoutingNumber = routingNumber
	this.VerificationStatus = verificationStatus
	return &this
}

// NewBankAccountFundingSourceModelWithDefaults instantiates a new BankAccountFundingSourceModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankAccountFundingSourceModelWithDefaults() *BankAccountFundingSourceModel {
	this := BankAccountFundingSourceModel{}
	return &this
}

// GetAccountSuffix returns the AccountSuffix field value
func (o *BankAccountFundingSourceModel) GetAccountSuffix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountSuffix
}

// GetAccountSuffixOk returns a tuple with the AccountSuffix field value
// and a boolean to check if the value has been set.
func (o *BankAccountFundingSourceModel) GetAccountSuffixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountSuffix, true
}

// SetAccountSuffix sets field value
func (o *BankAccountFundingSourceModel) SetAccountSuffix(v string) {
	o.AccountSuffix = v
}

// GetAccountType returns the AccountType field value
func (o *BankAccountFundingSourceModel) GetAccountType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value
// and a boolean to check if the value has been set.
func (o *BankAccountFundingSourceModel) GetAccountTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountType, true
}

// SetAccountType sets field value
func (o *BankAccountFundingSourceModel) SetAccountType(v string) {
	o.AccountType = v
}

// GetBusinessToken returns the BusinessToken field value if set, zero value otherwise.
func (o *BankAccountFundingSourceModel) GetBusinessToken() string {
	if o == nil || IsNil(o.BusinessToken) {
		var ret string
		return ret
	}
	return *o.BusinessToken
}

// GetBusinessTokenOk returns a tuple with the BusinessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccountFundingSourceModel) GetBusinessTokenOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessToken) {
		return nil, false
	}
	return o.BusinessToken, true
}

// HasBusinessToken returns a boolean if a field has been set.
func (o *BankAccountFundingSourceModel) HasBusinessToken() bool {
	if o != nil && !IsNil(o.BusinessToken) {
		return true
	}

	return false
}

// SetBusinessToken gets a reference to the given string and assigns it to the BusinessToken field.
func (o *BankAccountFundingSourceModel) SetBusinessToken(v string) {
	o.BusinessToken = &v
}

// GetNameOnAccount returns the NameOnAccount field value
func (o *BankAccountFundingSourceModel) GetNameOnAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NameOnAccount
}

// GetNameOnAccountOk returns a tuple with the NameOnAccount field value
// and a boolean to check if the value has been set.
func (o *BankAccountFundingSourceModel) GetNameOnAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NameOnAccount, true
}

// SetNameOnAccount sets field value
func (o *BankAccountFundingSourceModel) SetNameOnAccount(v string) {
	o.NameOnAccount = v
}

// GetRoutingNumber returns the RoutingNumber field value
func (o *BankAccountFundingSourceModel) GetRoutingNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoutingNumber
}

// GetRoutingNumberOk returns a tuple with the RoutingNumber field value
// and a boolean to check if the value has been set.
func (o *BankAccountFundingSourceModel) GetRoutingNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoutingNumber, true
}

// SetRoutingNumber sets field value
func (o *BankAccountFundingSourceModel) SetRoutingNumber(v string) {
	o.RoutingNumber = v
}

// GetUserToken returns the UserToken field value if set, zero value otherwise.
func (o *BankAccountFundingSourceModel) GetUserToken() string {
	if o == nil || IsNil(o.UserToken) {
		var ret string
		return ret
	}
	return *o.UserToken
}

// GetUserTokenOk returns a tuple with the UserToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccountFundingSourceModel) GetUserTokenOk() (*string, bool) {
	if o == nil || IsNil(o.UserToken) {
		return nil, false
	}
	return o.UserToken, true
}

// HasUserToken returns a boolean if a field has been set.
func (o *BankAccountFundingSourceModel) HasUserToken() bool {
	if o != nil && !IsNil(o.UserToken) {
		return true
	}

	return false
}

// SetUserToken gets a reference to the given string and assigns it to the UserToken field.
func (o *BankAccountFundingSourceModel) SetUserToken(v string) {
	o.UserToken = &v
}

// GetVerificationStatus returns the VerificationStatus field value
func (o *BankAccountFundingSourceModel) GetVerificationStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerificationStatus
}

// GetVerificationStatusOk returns a tuple with the VerificationStatus field value
// and a boolean to check if the value has been set.
func (o *BankAccountFundingSourceModel) GetVerificationStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerificationStatus, true
}

// SetVerificationStatus sets field value
func (o *BankAccountFundingSourceModel) SetVerificationStatus(v string) {
	o.VerificationStatus = v
}

func (o BankAccountFundingSourceModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BankAccountFundingSourceModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedFundingSourceModel, errFundingSourceModel := json.Marshal(o.FundingSourceModel)
	if errFundingSourceModel != nil {
		return map[string]interface{}{}, errFundingSourceModel
	}
	errFundingSourceModel = json.Unmarshal([]byte(serializedFundingSourceModel), &toSerialize)
	if errFundingSourceModel != nil {
		return map[string]interface{}{}, errFundingSourceModel
	}
	toSerialize["account_suffix"] = o.AccountSuffix
	toSerialize["account_type"] = o.AccountType
	if !IsNil(o.BusinessToken) {
		toSerialize["business_token"] = o.BusinessToken
	}
	toSerialize["name_on_account"] = o.NameOnAccount
	toSerialize["routing_number"] = o.RoutingNumber
	if !IsNil(o.UserToken) {
		toSerialize["user_token"] = o.UserToken
	}
	toSerialize["verification_status"] = o.VerificationStatus
	return toSerialize, nil
}

func (o *BankAccountFundingSourceModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"account_suffix",
		"account_type",
		"name_on_account",
		"routing_number",
		"verification_status",
		"active",
		"created_time",
		"is_default_account",
		"last_modified_time",
		"token",
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

	varBankAccountFundingSourceModel := _BankAccountFundingSourceModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBankAccountFundingSourceModel)

	if err != nil {
		return err
	}

	*o = BankAccountFundingSourceModel(varBankAccountFundingSourceModel)

	return err
}

type NullableBankAccountFundingSourceModel struct {
	value *BankAccountFundingSourceModel
	isSet bool
}

func (v NullableBankAccountFundingSourceModel) Get() *BankAccountFundingSourceModel {
	return v.value
}

func (v *NullableBankAccountFundingSourceModel) Set(val *BankAccountFundingSourceModel) {
	v.value = val
	v.isSet = true
}

func (v NullableBankAccountFundingSourceModel) IsSet() bool {
	return v.isSet
}

func (v *NullableBankAccountFundingSourceModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankAccountFundingSourceModel(val *BankAccountFundingSourceModel) *NullableBankAccountFundingSourceModel {
	return &NullableBankAccountFundingSourceModel{value: val, isSet: true}
}

func (v NullableBankAccountFundingSourceModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankAccountFundingSourceModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


