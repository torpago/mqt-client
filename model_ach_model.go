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

// checks if the AchModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AchModel{}

// AchModel struct for AchModel
type AchModel struct {
	// ACH account number.
	AccountNumber string `json:"account_number"`
	// Type of account.
	AccountType string `json:"account_type"`
	// Name of the financial institution where the ACH account is held.
	BankName *string `json:"bank_name,omitempty"`
	// Unique identifier of the business account holder. This token is required if a `user_token` is not specified.
	BusinessToken *string `json:"business_token,omitempty"`
	// If there are multiple funding sources, this field specifies which source is used by default in funding calls. If there is only one funding source, the system ignores this field and always uses that source.
	IsDefaultAccount *bool `json:"is_default_account,omitempty"`
	// Name on the ACH account.
	NameOnAccount string `json:"name_on_account"`
	// Routing number for the ACH account.
	RoutingNumber string `json:"routing_number"`
	// Unique identifier of the funding source. If you do not include a token, the system will generate one automatically. This token is necessary for use in other calls, so we recommend that rather than let the system generate one, you use a simple string that is easy to remember. This value cannot be updated.
	Token *string `json:"token,omitempty"`
	// Unique identifier of the user account holder. This token is required if a `business_token` is not specified.
	UserToken *string `json:"user_token,omitempty"`
	// Free-form text field for holding notes about verification. This field is returned only if `verification_override = true`.
	VerificationNotes *string `json:"verification_notes,omitempty"`
	// Allows the ACH funding source to be used, regardless of its verification status. Set this field to `true` if you can attest that you have verified the account on your own and that it will not be returned by the Federal Reserve.  *NOTE:* When using `PLAID` to validate a funding source, this field is always set to `true`.
	VerificationOverride *bool `json:"verification_override,omitempty"`
}

type _AchModel AchModel

// NewAchModel instantiates a new AchModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAchModel(accountNumber string, accountType string, nameOnAccount string, routingNumber string) *AchModel {
	this := AchModel{}
	this.AccountNumber = accountNumber
	this.AccountType = accountType
	var isDefaultAccount bool = false
	this.IsDefaultAccount = &isDefaultAccount
	this.NameOnAccount = nameOnAccount
	this.RoutingNumber = routingNumber
	var verificationOverride bool = false
	this.VerificationOverride = &verificationOverride
	return &this
}

// NewAchModelWithDefaults instantiates a new AchModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAchModelWithDefaults() *AchModel {
	this := AchModel{}
	var isDefaultAccount bool = false
	this.IsDefaultAccount = &isDefaultAccount
	var verificationOverride bool = false
	this.VerificationOverride = &verificationOverride
	return &this
}

// GetAccountNumber returns the AccountNumber field value
func (o *AchModel) GetAccountNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value
// and a boolean to check if the value has been set.
func (o *AchModel) GetAccountNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountNumber, true
}

// SetAccountNumber sets field value
func (o *AchModel) SetAccountNumber(v string) {
	o.AccountNumber = v
}

// GetAccountType returns the AccountType field value
func (o *AchModel) GetAccountType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value
// and a boolean to check if the value has been set.
func (o *AchModel) GetAccountTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountType, true
}

// SetAccountType sets field value
func (o *AchModel) SetAccountType(v string) {
	o.AccountType = v
}

// GetBankName returns the BankName field value if set, zero value otherwise.
func (o *AchModel) GetBankName() string {
	if o == nil || IsNil(o.BankName) {
		var ret string
		return ret
	}
	return *o.BankName
}

// GetBankNameOk returns a tuple with the BankName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AchModel) GetBankNameOk() (*string, bool) {
	if o == nil || IsNil(o.BankName) {
		return nil, false
	}
	return o.BankName, true
}

// HasBankName returns a boolean if a field has been set.
func (o *AchModel) HasBankName() bool {
	if o != nil && !IsNil(o.BankName) {
		return true
	}

	return false
}

// SetBankName gets a reference to the given string and assigns it to the BankName field.
func (o *AchModel) SetBankName(v string) {
	o.BankName = &v
}

// GetBusinessToken returns the BusinessToken field value if set, zero value otherwise.
func (o *AchModel) GetBusinessToken() string {
	if o == nil || IsNil(o.BusinessToken) {
		var ret string
		return ret
	}
	return *o.BusinessToken
}

// GetBusinessTokenOk returns a tuple with the BusinessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AchModel) GetBusinessTokenOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessToken) {
		return nil, false
	}
	return o.BusinessToken, true
}

// HasBusinessToken returns a boolean if a field has been set.
func (o *AchModel) HasBusinessToken() bool {
	if o != nil && !IsNil(o.BusinessToken) {
		return true
	}

	return false
}

// SetBusinessToken gets a reference to the given string and assigns it to the BusinessToken field.
func (o *AchModel) SetBusinessToken(v string) {
	o.BusinessToken = &v
}

// GetIsDefaultAccount returns the IsDefaultAccount field value if set, zero value otherwise.
func (o *AchModel) GetIsDefaultAccount() bool {
	if o == nil || IsNil(o.IsDefaultAccount) {
		var ret bool
		return ret
	}
	return *o.IsDefaultAccount
}

// GetIsDefaultAccountOk returns a tuple with the IsDefaultAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AchModel) GetIsDefaultAccountOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefaultAccount) {
		return nil, false
	}
	return o.IsDefaultAccount, true
}

// HasIsDefaultAccount returns a boolean if a field has been set.
func (o *AchModel) HasIsDefaultAccount() bool {
	if o != nil && !IsNil(o.IsDefaultAccount) {
		return true
	}

	return false
}

// SetIsDefaultAccount gets a reference to the given bool and assigns it to the IsDefaultAccount field.
func (o *AchModel) SetIsDefaultAccount(v bool) {
	o.IsDefaultAccount = &v
}

// GetNameOnAccount returns the NameOnAccount field value
func (o *AchModel) GetNameOnAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NameOnAccount
}

// GetNameOnAccountOk returns a tuple with the NameOnAccount field value
// and a boolean to check if the value has been set.
func (o *AchModel) GetNameOnAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NameOnAccount, true
}

// SetNameOnAccount sets field value
func (o *AchModel) SetNameOnAccount(v string) {
	o.NameOnAccount = v
}

// GetRoutingNumber returns the RoutingNumber field value
func (o *AchModel) GetRoutingNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoutingNumber
}

// GetRoutingNumberOk returns a tuple with the RoutingNumber field value
// and a boolean to check if the value has been set.
func (o *AchModel) GetRoutingNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoutingNumber, true
}

// SetRoutingNumber sets field value
func (o *AchModel) SetRoutingNumber(v string) {
	o.RoutingNumber = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *AchModel) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AchModel) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *AchModel) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *AchModel) SetToken(v string) {
	o.Token = &v
}

// GetUserToken returns the UserToken field value if set, zero value otherwise.
func (o *AchModel) GetUserToken() string {
	if o == nil || IsNil(o.UserToken) {
		var ret string
		return ret
	}
	return *o.UserToken
}

// GetUserTokenOk returns a tuple with the UserToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AchModel) GetUserTokenOk() (*string, bool) {
	if o == nil || IsNil(o.UserToken) {
		return nil, false
	}
	return o.UserToken, true
}

// HasUserToken returns a boolean if a field has been set.
func (o *AchModel) HasUserToken() bool {
	if o != nil && !IsNil(o.UserToken) {
		return true
	}

	return false
}

// SetUserToken gets a reference to the given string and assigns it to the UserToken field.
func (o *AchModel) SetUserToken(v string) {
	o.UserToken = &v
}

// GetVerificationNotes returns the VerificationNotes field value if set, zero value otherwise.
func (o *AchModel) GetVerificationNotes() string {
	if o == nil || IsNil(o.VerificationNotes) {
		var ret string
		return ret
	}
	return *o.VerificationNotes
}

// GetVerificationNotesOk returns a tuple with the VerificationNotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AchModel) GetVerificationNotesOk() (*string, bool) {
	if o == nil || IsNil(o.VerificationNotes) {
		return nil, false
	}
	return o.VerificationNotes, true
}

// HasVerificationNotes returns a boolean if a field has been set.
func (o *AchModel) HasVerificationNotes() bool {
	if o != nil && !IsNil(o.VerificationNotes) {
		return true
	}

	return false
}

// SetVerificationNotes gets a reference to the given string and assigns it to the VerificationNotes field.
func (o *AchModel) SetVerificationNotes(v string) {
	o.VerificationNotes = &v
}

// GetVerificationOverride returns the VerificationOverride field value if set, zero value otherwise.
func (o *AchModel) GetVerificationOverride() bool {
	if o == nil || IsNil(o.VerificationOverride) {
		var ret bool
		return ret
	}
	return *o.VerificationOverride
}

// GetVerificationOverrideOk returns a tuple with the VerificationOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AchModel) GetVerificationOverrideOk() (*bool, bool) {
	if o == nil || IsNil(o.VerificationOverride) {
		return nil, false
	}
	return o.VerificationOverride, true
}

// HasVerificationOverride returns a boolean if a field has been set.
func (o *AchModel) HasVerificationOverride() bool {
	if o != nil && !IsNil(o.VerificationOverride) {
		return true
	}

	return false
}

// SetVerificationOverride gets a reference to the given bool and assigns it to the VerificationOverride field.
func (o *AchModel) SetVerificationOverride(v bool) {
	o.VerificationOverride = &v
}

func (o AchModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AchModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account_number"] = o.AccountNumber
	toSerialize["account_type"] = o.AccountType
	if !IsNil(o.BankName) {
		toSerialize["bank_name"] = o.BankName
	}
	if !IsNil(o.BusinessToken) {
		toSerialize["business_token"] = o.BusinessToken
	}
	if !IsNil(o.IsDefaultAccount) {
		toSerialize["is_default_account"] = o.IsDefaultAccount
	}
	toSerialize["name_on_account"] = o.NameOnAccount
	toSerialize["routing_number"] = o.RoutingNumber
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.UserToken) {
		toSerialize["user_token"] = o.UserToken
	}
	if !IsNil(o.VerificationNotes) {
		toSerialize["verification_notes"] = o.VerificationNotes
	}
	if !IsNil(o.VerificationOverride) {
		toSerialize["verification_override"] = o.VerificationOverride
	}
	return toSerialize, nil
}

func (o *AchModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"account_number",
		"account_type",
		"name_on_account",
		"routing_number",
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

	varAchModel := _AchModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAchModel)

	if err != nil {
		return err
	}

	*o = AchModel(varAchModel)

	return err
}

type NullableAchModel struct {
	value *AchModel
	isSet bool
}

func (v NullableAchModel) Get() *AchModel {
	return v.value
}

func (v *NullableAchModel) Set(val *AchModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAchModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAchModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAchModel(val *AchModel) *NullableAchModel {
	return &NullableAchModel{value: val, isSet: true}
}

func (v NullableAchModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAchModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


