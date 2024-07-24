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

// checks if the FundingAccountResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FundingAccountResponseModel{}

// FundingAccountResponseModel struct for FundingAccountResponseModel
type FundingAccountResponseModel struct {
	// Account identifier appended to the bank account number. This field is returned if it exists in the resource.
	AccountSuffix *string `json:"account_suffix,omitempty"`
	// Type of account. This field is returned if it exists in the resource.
	AccountType *string `json:"account_type,omitempty"`
	// Specifies whether the account is active. This field is returned if it exists in the resource.
	Active *bool `json:"active,omitempty"`
	// Unique identifier of the business account holder. This token is returned if a `user_token` is not specified.
	BusinessToken *string `json:"business_token,omitempty"`
	// Date and time when the resource was created, in UTC.
	CreatedTime time.Time `json:"created_time"`
	// Date and time in UTC when either the request for account validation was sent to the third-party partner, or when the funding source was verified by microdeposits.  This field is returned if it exists in the resource.
	DateSentForVerification *time.Time `json:"date_sent_for_verification,omitempty"`
	// Date and time when the account was verified, in UTC. This field is returned if it exists in the resource.
	DateVerified *time.Time `json:"date_verified,omitempty"`
	// Payment card expiration date. This field is returned if it exists in the resource.
	ExpDate *string `json:"exp_date,omitempty"`
	// If there are multiple funding sources, this field specifies which source is used by default in funding calls. If there is only one funding source, the system ignores this field and always uses that source.  This field is returned if it exists in the resource.
	IsDefaultAccount *bool `json:"is_default_account,omitempty"`
	// Date and time when the resource was last modified, in UTC.
	LastModifiedTime time.Time `json:"last_modified_time"`
	LinkPartnerAccountReferenceToken *string `json:"link_partner_account_reference_token,omitempty"`
	// Name on the account. This field is returned if it exists in the resource.
	NameOnAccount *string `json:"name_on_account,omitempty"`
	// Name of the partner who validated the account holder. Returned when a third-party partner was used for account validation.
	Partner *string `json:"partner,omitempty"`
	// Supplied by the account validation partner, this value is a reference to the account holder's details, such as the account number and routing number. Returned when a third-party partner was used for account validation.
	PartnerAccountLinkReferenceToken *string `json:"partner_account_link_reference_token,omitempty"`
	// Unique identifier of the funding source. This field is returned if it exists in the resource.
	Token *string `json:"token,omitempty"`
	// Funding source type.
	Type *string `json:"type,omitempty"`
	// Unique identifier of the user account holder. This token is returned if a `business_token` is not specified.
	UserToken *string `json:"user_token,omitempty"`
	// Free-form text field for holding notes about verification. This field is returned only if `verification_override = true`.
	VerificationNotes *string `json:"verification_notes,omitempty"`
	// Allows the ACH funding source to be used regardless of its verification status.  *NOTE:* When using `PLAID` to validate a funding source, this field is always set to `true`.
	VerificationOverride *bool `json:"verification_override,omitempty"`
	// Account verification status. This field is returned if it exists in the resource.
	VerificationStatus *string `json:"verification_status,omitempty"`
}

type _FundingAccountResponseModel FundingAccountResponseModel

// NewFundingAccountResponseModel instantiates a new FundingAccountResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFundingAccountResponseModel(createdTime time.Time, lastModifiedTime time.Time) *FundingAccountResponseModel {
	this := FundingAccountResponseModel{}
	var active bool = false
	this.Active = &active
	this.CreatedTime = createdTime
	var isDefaultAccount bool = false
	this.IsDefaultAccount = &isDefaultAccount
	this.LastModifiedTime = lastModifiedTime
	var verificationOverride bool = false
	this.VerificationOverride = &verificationOverride
	return &this
}

// NewFundingAccountResponseModelWithDefaults instantiates a new FundingAccountResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFundingAccountResponseModelWithDefaults() *FundingAccountResponseModel {
	this := FundingAccountResponseModel{}
	var active bool = false
	this.Active = &active
	var isDefaultAccount bool = false
	this.IsDefaultAccount = &isDefaultAccount
	var verificationOverride bool = false
	this.VerificationOverride = &verificationOverride
	return &this
}

// GetAccountSuffix returns the AccountSuffix field value if set, zero value otherwise.
func (o *FundingAccountResponseModel) GetAccountSuffix() string {
	if o == nil || IsNil(o.AccountSuffix) {
		var ret string
		return ret
	}
	return *o.AccountSuffix
}

// GetAccountSuffixOk returns a tuple with the AccountSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingAccountResponseModel) GetAccountSuffixOk() (*string, bool) {
	if o == nil || IsNil(o.AccountSuffix) {
		return nil, false
	}
	return o.AccountSuffix, true
}

// HasAccountSuffix returns a boolean if a field has been set.
func (o *FundingAccountResponseModel) HasAccountSuffix() bool {
	if o != nil && !IsNil(o.AccountSuffix) {
		return true
	}

	return false
}

// SetAccountSuffix gets a reference to the given string and assigns it to the AccountSuffix field.
func (o *FundingAccountResponseModel) SetAccountSuffix(v string) {
	o.AccountSuffix = &v
}

// GetAccountType returns the AccountType field value if set, zero value otherwise.
func (o *FundingAccountResponseModel) GetAccountType() string {
	if o == nil || IsNil(o.AccountType) {
		var ret string
		return ret
	}
	return *o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingAccountResponseModel) GetAccountTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AccountType) {
		return nil, false
	}
	return o.AccountType, true
}

// HasAccountType returns a boolean if a field has been set.
func (o *FundingAccountResponseModel) HasAccountType() bool {
	if o != nil && !IsNil(o.AccountType) {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given string and assigns it to the AccountType field.
func (o *FundingAccountResponseModel) SetAccountType(v string) {
	o.AccountType = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *FundingAccountResponseModel) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingAccountResponseModel) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *FundingAccountResponseModel) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *FundingAccountResponseModel) SetActive(v bool) {
	o.Active = &v
}

// GetBusinessToken returns the BusinessToken field value if set, zero value otherwise.
func (o *FundingAccountResponseModel) GetBusinessToken() string {
	if o == nil || IsNil(o.BusinessToken) {
		var ret string
		return ret
	}
	return *o.BusinessToken
}

// GetBusinessTokenOk returns a tuple with the BusinessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingAccountResponseModel) GetBusinessTokenOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessToken) {
		return nil, false
	}
	return o.BusinessToken, true
}

// HasBusinessToken returns a boolean if a field has been set.
func (o *FundingAccountResponseModel) HasBusinessToken() bool {
	if o != nil && !IsNil(o.BusinessToken) {
		return true
	}

	return false
}

// SetBusinessToken gets a reference to the given string and assigns it to the BusinessToken field.
func (o *FundingAccountResponseModel) SetBusinessToken(v string) {
	o.BusinessToken = &v
}

// GetCreatedTime returns the CreatedTime field value
func (o *FundingAccountResponseModel) GetCreatedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value
// and a boolean to check if the value has been set.
func (o *FundingAccountResponseModel) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTime, true
}

// SetCreatedTime sets field value
func (o *FundingAccountResponseModel) SetCreatedTime(v time.Time) {
	o.CreatedTime = v
}

// GetDateSentForVerification returns the DateSentForVerification field value if set, zero value otherwise.
func (o *FundingAccountResponseModel) GetDateSentForVerification() time.Time {
	if o == nil || IsNil(o.DateSentForVerification) {
		var ret time.Time
		return ret
	}
	return *o.DateSentForVerification
}

// GetDateSentForVerificationOk returns a tuple with the DateSentForVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingAccountResponseModel) GetDateSentForVerificationOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateSentForVerification) {
		return nil, false
	}
	return o.DateSentForVerification, true
}

// HasDateSentForVerification returns a boolean if a field has been set.
func (o *FundingAccountResponseModel) HasDateSentForVerification() bool {
	if o != nil && !IsNil(o.DateSentForVerification) {
		return true
	}

	return false
}

// SetDateSentForVerification gets a reference to the given time.Time and assigns it to the DateSentForVerification field.
func (o *FundingAccountResponseModel) SetDateSentForVerification(v time.Time) {
	o.DateSentForVerification = &v
}

// GetDateVerified returns the DateVerified field value if set, zero value otherwise.
func (o *FundingAccountResponseModel) GetDateVerified() time.Time {
	if o == nil || IsNil(o.DateVerified) {
		var ret time.Time
		return ret
	}
	return *o.DateVerified
}

// GetDateVerifiedOk returns a tuple with the DateVerified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingAccountResponseModel) GetDateVerifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateVerified) {
		return nil, false
	}
	return o.DateVerified, true
}

// HasDateVerified returns a boolean if a field has been set.
func (o *FundingAccountResponseModel) HasDateVerified() bool {
	if o != nil && !IsNil(o.DateVerified) {
		return true
	}

	return false
}

// SetDateVerified gets a reference to the given time.Time and assigns it to the DateVerified field.
func (o *FundingAccountResponseModel) SetDateVerified(v time.Time) {
	o.DateVerified = &v
}

// GetExpDate returns the ExpDate field value if set, zero value otherwise.
func (o *FundingAccountResponseModel) GetExpDate() string {
	if o == nil || IsNil(o.ExpDate) {
		var ret string
		return ret
	}
	return *o.ExpDate
}

// GetExpDateOk returns a tuple with the ExpDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingAccountResponseModel) GetExpDateOk() (*string, bool) {
	if o == nil || IsNil(o.ExpDate) {
		return nil, false
	}
	return o.ExpDate, true
}

// HasExpDate returns a boolean if a field has been set.
func (o *FundingAccountResponseModel) HasExpDate() bool {
	if o != nil && !IsNil(o.ExpDate) {
		return true
	}

	return false
}

// SetExpDate gets a reference to the given string and assigns it to the ExpDate field.
func (o *FundingAccountResponseModel) SetExpDate(v string) {
	o.ExpDate = &v
}

// GetIsDefaultAccount returns the IsDefaultAccount field value if set, zero value otherwise.
func (o *FundingAccountResponseModel) GetIsDefaultAccount() bool {
	if o == nil || IsNil(o.IsDefaultAccount) {
		var ret bool
		return ret
	}
	return *o.IsDefaultAccount
}

// GetIsDefaultAccountOk returns a tuple with the IsDefaultAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingAccountResponseModel) GetIsDefaultAccountOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefaultAccount) {
		return nil, false
	}
	return o.IsDefaultAccount, true
}

// HasIsDefaultAccount returns a boolean if a field has been set.
func (o *FundingAccountResponseModel) HasIsDefaultAccount() bool {
	if o != nil && !IsNil(o.IsDefaultAccount) {
		return true
	}

	return false
}

// SetIsDefaultAccount gets a reference to the given bool and assigns it to the IsDefaultAccount field.
func (o *FundingAccountResponseModel) SetIsDefaultAccount(v bool) {
	o.IsDefaultAccount = &v
}

// GetLastModifiedTime returns the LastModifiedTime field value
func (o *FundingAccountResponseModel) GetLastModifiedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value
// and a boolean to check if the value has been set.
func (o *FundingAccountResponseModel) GetLastModifiedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastModifiedTime, true
}

// SetLastModifiedTime sets field value
func (o *FundingAccountResponseModel) SetLastModifiedTime(v time.Time) {
	o.LastModifiedTime = v
}

// GetLinkPartnerAccountReferenceToken returns the LinkPartnerAccountReferenceToken field value if set, zero value otherwise.
func (o *FundingAccountResponseModel) GetLinkPartnerAccountReferenceToken() string {
	if o == nil || IsNil(o.LinkPartnerAccountReferenceToken) {
		var ret string
		return ret
	}
	return *o.LinkPartnerAccountReferenceToken
}

// GetLinkPartnerAccountReferenceTokenOk returns a tuple with the LinkPartnerAccountReferenceToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingAccountResponseModel) GetLinkPartnerAccountReferenceTokenOk() (*string, bool) {
	if o == nil || IsNil(o.LinkPartnerAccountReferenceToken) {
		return nil, false
	}
	return o.LinkPartnerAccountReferenceToken, true
}

// HasLinkPartnerAccountReferenceToken returns a boolean if a field has been set.
func (o *FundingAccountResponseModel) HasLinkPartnerAccountReferenceToken() bool {
	if o != nil && !IsNil(o.LinkPartnerAccountReferenceToken) {
		return true
	}

	return false
}

// SetLinkPartnerAccountReferenceToken gets a reference to the given string and assigns it to the LinkPartnerAccountReferenceToken field.
func (o *FundingAccountResponseModel) SetLinkPartnerAccountReferenceToken(v string) {
	o.LinkPartnerAccountReferenceToken = &v
}

// GetNameOnAccount returns the NameOnAccount field value if set, zero value otherwise.
func (o *FundingAccountResponseModel) GetNameOnAccount() string {
	if o == nil || IsNil(o.NameOnAccount) {
		var ret string
		return ret
	}
	return *o.NameOnAccount
}

// GetNameOnAccountOk returns a tuple with the NameOnAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingAccountResponseModel) GetNameOnAccountOk() (*string, bool) {
	if o == nil || IsNil(o.NameOnAccount) {
		return nil, false
	}
	return o.NameOnAccount, true
}

// HasNameOnAccount returns a boolean if a field has been set.
func (o *FundingAccountResponseModel) HasNameOnAccount() bool {
	if o != nil && !IsNil(o.NameOnAccount) {
		return true
	}

	return false
}

// SetNameOnAccount gets a reference to the given string and assigns it to the NameOnAccount field.
func (o *FundingAccountResponseModel) SetNameOnAccount(v string) {
	o.NameOnAccount = &v
}

// GetPartner returns the Partner field value if set, zero value otherwise.
func (o *FundingAccountResponseModel) GetPartner() string {
	if o == nil || IsNil(o.Partner) {
		var ret string
		return ret
	}
	return *o.Partner
}

// GetPartnerOk returns a tuple with the Partner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingAccountResponseModel) GetPartnerOk() (*string, bool) {
	if o == nil || IsNil(o.Partner) {
		return nil, false
	}
	return o.Partner, true
}

// HasPartner returns a boolean if a field has been set.
func (o *FundingAccountResponseModel) HasPartner() bool {
	if o != nil && !IsNil(o.Partner) {
		return true
	}

	return false
}

// SetPartner gets a reference to the given string and assigns it to the Partner field.
func (o *FundingAccountResponseModel) SetPartner(v string) {
	o.Partner = &v
}

// GetPartnerAccountLinkReferenceToken returns the PartnerAccountLinkReferenceToken field value if set, zero value otherwise.
func (o *FundingAccountResponseModel) GetPartnerAccountLinkReferenceToken() string {
	if o == nil || IsNil(o.PartnerAccountLinkReferenceToken) {
		var ret string
		return ret
	}
	return *o.PartnerAccountLinkReferenceToken
}

// GetPartnerAccountLinkReferenceTokenOk returns a tuple with the PartnerAccountLinkReferenceToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingAccountResponseModel) GetPartnerAccountLinkReferenceTokenOk() (*string, bool) {
	if o == nil || IsNil(o.PartnerAccountLinkReferenceToken) {
		return nil, false
	}
	return o.PartnerAccountLinkReferenceToken, true
}

// HasPartnerAccountLinkReferenceToken returns a boolean if a field has been set.
func (o *FundingAccountResponseModel) HasPartnerAccountLinkReferenceToken() bool {
	if o != nil && !IsNil(o.PartnerAccountLinkReferenceToken) {
		return true
	}

	return false
}

// SetPartnerAccountLinkReferenceToken gets a reference to the given string and assigns it to the PartnerAccountLinkReferenceToken field.
func (o *FundingAccountResponseModel) SetPartnerAccountLinkReferenceToken(v string) {
	o.PartnerAccountLinkReferenceToken = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *FundingAccountResponseModel) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingAccountResponseModel) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *FundingAccountResponseModel) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *FundingAccountResponseModel) SetToken(v string) {
	o.Token = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FundingAccountResponseModel) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingAccountResponseModel) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FundingAccountResponseModel) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FundingAccountResponseModel) SetType(v string) {
	o.Type = &v
}

// GetUserToken returns the UserToken field value if set, zero value otherwise.
func (o *FundingAccountResponseModel) GetUserToken() string {
	if o == nil || IsNil(o.UserToken) {
		var ret string
		return ret
	}
	return *o.UserToken
}

// GetUserTokenOk returns a tuple with the UserToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingAccountResponseModel) GetUserTokenOk() (*string, bool) {
	if o == nil || IsNil(o.UserToken) {
		return nil, false
	}
	return o.UserToken, true
}

// HasUserToken returns a boolean if a field has been set.
func (o *FundingAccountResponseModel) HasUserToken() bool {
	if o != nil && !IsNil(o.UserToken) {
		return true
	}

	return false
}

// SetUserToken gets a reference to the given string and assigns it to the UserToken field.
func (o *FundingAccountResponseModel) SetUserToken(v string) {
	o.UserToken = &v
}

// GetVerificationNotes returns the VerificationNotes field value if set, zero value otherwise.
func (o *FundingAccountResponseModel) GetVerificationNotes() string {
	if o == nil || IsNil(o.VerificationNotes) {
		var ret string
		return ret
	}
	return *o.VerificationNotes
}

// GetVerificationNotesOk returns a tuple with the VerificationNotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingAccountResponseModel) GetVerificationNotesOk() (*string, bool) {
	if o == nil || IsNil(o.VerificationNotes) {
		return nil, false
	}
	return o.VerificationNotes, true
}

// HasVerificationNotes returns a boolean if a field has been set.
func (o *FundingAccountResponseModel) HasVerificationNotes() bool {
	if o != nil && !IsNil(o.VerificationNotes) {
		return true
	}

	return false
}

// SetVerificationNotes gets a reference to the given string and assigns it to the VerificationNotes field.
func (o *FundingAccountResponseModel) SetVerificationNotes(v string) {
	o.VerificationNotes = &v
}

// GetVerificationOverride returns the VerificationOverride field value if set, zero value otherwise.
func (o *FundingAccountResponseModel) GetVerificationOverride() bool {
	if o == nil || IsNil(o.VerificationOverride) {
		var ret bool
		return ret
	}
	return *o.VerificationOverride
}

// GetVerificationOverrideOk returns a tuple with the VerificationOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingAccountResponseModel) GetVerificationOverrideOk() (*bool, bool) {
	if o == nil || IsNil(o.VerificationOverride) {
		return nil, false
	}
	return o.VerificationOverride, true
}

// HasVerificationOverride returns a boolean if a field has been set.
func (o *FundingAccountResponseModel) HasVerificationOverride() bool {
	if o != nil && !IsNil(o.VerificationOverride) {
		return true
	}

	return false
}

// SetVerificationOverride gets a reference to the given bool and assigns it to the VerificationOverride field.
func (o *FundingAccountResponseModel) SetVerificationOverride(v bool) {
	o.VerificationOverride = &v
}

// GetVerificationStatus returns the VerificationStatus field value if set, zero value otherwise.
func (o *FundingAccountResponseModel) GetVerificationStatus() string {
	if o == nil || IsNil(o.VerificationStatus) {
		var ret string
		return ret
	}
	return *o.VerificationStatus
}

// GetVerificationStatusOk returns a tuple with the VerificationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingAccountResponseModel) GetVerificationStatusOk() (*string, bool) {
	if o == nil || IsNil(o.VerificationStatus) {
		return nil, false
	}
	return o.VerificationStatus, true
}

// HasVerificationStatus returns a boolean if a field has been set.
func (o *FundingAccountResponseModel) HasVerificationStatus() bool {
	if o != nil && !IsNil(o.VerificationStatus) {
		return true
	}

	return false
}

// SetVerificationStatus gets a reference to the given string and assigns it to the VerificationStatus field.
func (o *FundingAccountResponseModel) SetVerificationStatus(v string) {
	o.VerificationStatus = &v
}

func (o FundingAccountResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FundingAccountResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountSuffix) {
		toSerialize["account_suffix"] = o.AccountSuffix
	}
	if !IsNil(o.AccountType) {
		toSerialize["account_type"] = o.AccountType
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.BusinessToken) {
		toSerialize["business_token"] = o.BusinessToken
	}
	toSerialize["created_time"] = o.CreatedTime
	if !IsNil(o.DateSentForVerification) {
		toSerialize["date_sent_for_verification"] = o.DateSentForVerification
	}
	if !IsNil(o.DateVerified) {
		toSerialize["date_verified"] = o.DateVerified
	}
	if !IsNil(o.ExpDate) {
		toSerialize["exp_date"] = o.ExpDate
	}
	if !IsNil(o.IsDefaultAccount) {
		toSerialize["is_default_account"] = o.IsDefaultAccount
	}
	toSerialize["last_modified_time"] = o.LastModifiedTime
	if !IsNil(o.LinkPartnerAccountReferenceToken) {
		toSerialize["link_partner_account_reference_token"] = o.LinkPartnerAccountReferenceToken
	}
	if !IsNil(o.NameOnAccount) {
		toSerialize["name_on_account"] = o.NameOnAccount
	}
	if !IsNil(o.Partner) {
		toSerialize["partner"] = o.Partner
	}
	if !IsNil(o.PartnerAccountLinkReferenceToken) {
		toSerialize["partner_account_link_reference_token"] = o.PartnerAccountLinkReferenceToken
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
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
	if !IsNil(o.VerificationStatus) {
		toSerialize["verification_status"] = o.VerificationStatus
	}
	return toSerialize, nil
}

func (o *FundingAccountResponseModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"created_time",
		"last_modified_time",
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

	varFundingAccountResponseModel := _FundingAccountResponseModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFundingAccountResponseModel)

	if err != nil {
		return err
	}

	*o = FundingAccountResponseModel(varFundingAccountResponseModel)

	return err
}

type NullableFundingAccountResponseModel struct {
	value *FundingAccountResponseModel
	isSet bool
}

func (v NullableFundingAccountResponseModel) Get() *FundingAccountResponseModel {
	return v.value
}

func (v *NullableFundingAccountResponseModel) Set(val *FundingAccountResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableFundingAccountResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableFundingAccountResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFundingAccountResponseModel(val *FundingAccountResponseModel) *NullableFundingAccountResponseModel {
	return &NullableFundingAccountResponseModel{value: val, isSet: true}
}

func (v NullableFundingAccountResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFundingAccountResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


