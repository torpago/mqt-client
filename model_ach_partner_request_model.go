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

// checks if the AchPartnerRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AchPartnerRequestModel{}

// AchPartnerRequestModel struct for AchPartnerRequestModel
type AchPartnerRequestModel struct {
	// Required if 'user_token' is null
	BusinessToken *string `json:"business_token,omitempty"`
	IdempotentHash *string `json:"idempotentHash,omitempty"`
	// If there are multiple funding sources, this field specifies which source is used by default in funding calls. If there is only one funding source, the system ignores this field and always uses that source.
	IsDefaultAccount *bool `json:"is_default_account,omitempty"`
	// Name of the partner who validated the account holder. Returned when a third-party partner was used for account validation.
	Partner string `json:"partner"`
	// Supplied by the account validation partner, this value is a reference to the account holder's details, such as the account number and routing number. Returned when a third-party partner was used for account validation.
	PartnerAccountLinkReferenceToken string `json:"partner_account_link_reference_token"`
	// Unique identifier of the funding source. If you do not include a token, the system will generate one automatically. This token is necessary for use in other calls, so we recommend that rather than let the system generate one, you use a simple string that is easy to remember. This value cannot be updated.
	Token *string `json:"token,omitempty"`
	// Supplied by the account validation partner, this value is a reference to the account holder's details, such as the account number and routing number. This token is required if a `business_token` is not specified.  Send a `GET` request to `/users` to retrieve user tokens.
	UserToken *string `json:"user_token,omitempty"`
}

type _AchPartnerRequestModel AchPartnerRequestModel

// NewAchPartnerRequestModel instantiates a new AchPartnerRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAchPartnerRequestModel(partner string, partnerAccountLinkReferenceToken string) *AchPartnerRequestModel {
	this := AchPartnerRequestModel{}
	var isDefaultAccount bool = false
	this.IsDefaultAccount = &isDefaultAccount
	this.Partner = partner
	this.PartnerAccountLinkReferenceToken = partnerAccountLinkReferenceToken
	return &this
}

// NewAchPartnerRequestModelWithDefaults instantiates a new AchPartnerRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAchPartnerRequestModelWithDefaults() *AchPartnerRequestModel {
	this := AchPartnerRequestModel{}
	var isDefaultAccount bool = false
	this.IsDefaultAccount = &isDefaultAccount
	return &this
}

// GetBusinessToken returns the BusinessToken field value if set, zero value otherwise.
func (o *AchPartnerRequestModel) GetBusinessToken() string {
	if o == nil || IsNil(o.BusinessToken) {
		var ret string
		return ret
	}
	return *o.BusinessToken
}

// GetBusinessTokenOk returns a tuple with the BusinessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AchPartnerRequestModel) GetBusinessTokenOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessToken) {
		return nil, false
	}
	return o.BusinessToken, true
}

// HasBusinessToken returns a boolean if a field has been set.
func (o *AchPartnerRequestModel) HasBusinessToken() bool {
	if o != nil && !IsNil(o.BusinessToken) {
		return true
	}

	return false
}

// SetBusinessToken gets a reference to the given string and assigns it to the BusinessToken field.
func (o *AchPartnerRequestModel) SetBusinessToken(v string) {
	o.BusinessToken = &v
}

// GetIdempotentHash returns the IdempotentHash field value if set, zero value otherwise.
func (o *AchPartnerRequestModel) GetIdempotentHash() string {
	if o == nil || IsNil(o.IdempotentHash) {
		var ret string
		return ret
	}
	return *o.IdempotentHash
}

// GetIdempotentHashOk returns a tuple with the IdempotentHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AchPartnerRequestModel) GetIdempotentHashOk() (*string, bool) {
	if o == nil || IsNil(o.IdempotentHash) {
		return nil, false
	}
	return o.IdempotentHash, true
}

// HasIdempotentHash returns a boolean if a field has been set.
func (o *AchPartnerRequestModel) HasIdempotentHash() bool {
	if o != nil && !IsNil(o.IdempotentHash) {
		return true
	}

	return false
}

// SetIdempotentHash gets a reference to the given string and assigns it to the IdempotentHash field.
func (o *AchPartnerRequestModel) SetIdempotentHash(v string) {
	o.IdempotentHash = &v
}

// GetIsDefaultAccount returns the IsDefaultAccount field value if set, zero value otherwise.
func (o *AchPartnerRequestModel) GetIsDefaultAccount() bool {
	if o == nil || IsNil(o.IsDefaultAccount) {
		var ret bool
		return ret
	}
	return *o.IsDefaultAccount
}

// GetIsDefaultAccountOk returns a tuple with the IsDefaultAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AchPartnerRequestModel) GetIsDefaultAccountOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefaultAccount) {
		return nil, false
	}
	return o.IsDefaultAccount, true
}

// HasIsDefaultAccount returns a boolean if a field has been set.
func (o *AchPartnerRequestModel) HasIsDefaultAccount() bool {
	if o != nil && !IsNil(o.IsDefaultAccount) {
		return true
	}

	return false
}

// SetIsDefaultAccount gets a reference to the given bool and assigns it to the IsDefaultAccount field.
func (o *AchPartnerRequestModel) SetIsDefaultAccount(v bool) {
	o.IsDefaultAccount = &v
}

// GetPartner returns the Partner field value
func (o *AchPartnerRequestModel) GetPartner() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Partner
}

// GetPartnerOk returns a tuple with the Partner field value
// and a boolean to check if the value has been set.
func (o *AchPartnerRequestModel) GetPartnerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Partner, true
}

// SetPartner sets field value
func (o *AchPartnerRequestModel) SetPartner(v string) {
	o.Partner = v
}

// GetPartnerAccountLinkReferenceToken returns the PartnerAccountLinkReferenceToken field value
func (o *AchPartnerRequestModel) GetPartnerAccountLinkReferenceToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PartnerAccountLinkReferenceToken
}

// GetPartnerAccountLinkReferenceTokenOk returns a tuple with the PartnerAccountLinkReferenceToken field value
// and a boolean to check if the value has been set.
func (o *AchPartnerRequestModel) GetPartnerAccountLinkReferenceTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PartnerAccountLinkReferenceToken, true
}

// SetPartnerAccountLinkReferenceToken sets field value
func (o *AchPartnerRequestModel) SetPartnerAccountLinkReferenceToken(v string) {
	o.PartnerAccountLinkReferenceToken = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *AchPartnerRequestModel) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AchPartnerRequestModel) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *AchPartnerRequestModel) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *AchPartnerRequestModel) SetToken(v string) {
	o.Token = &v
}

// GetUserToken returns the UserToken field value if set, zero value otherwise.
func (o *AchPartnerRequestModel) GetUserToken() string {
	if o == nil || IsNil(o.UserToken) {
		var ret string
		return ret
	}
	return *o.UserToken
}

// GetUserTokenOk returns a tuple with the UserToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AchPartnerRequestModel) GetUserTokenOk() (*string, bool) {
	if o == nil || IsNil(o.UserToken) {
		return nil, false
	}
	return o.UserToken, true
}

// HasUserToken returns a boolean if a field has been set.
func (o *AchPartnerRequestModel) HasUserToken() bool {
	if o != nil && !IsNil(o.UserToken) {
		return true
	}

	return false
}

// SetUserToken gets a reference to the given string and assigns it to the UserToken field.
func (o *AchPartnerRequestModel) SetUserToken(v string) {
	o.UserToken = &v
}

func (o AchPartnerRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AchPartnerRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BusinessToken) {
		toSerialize["business_token"] = o.BusinessToken
	}
	if !IsNil(o.IdempotentHash) {
		toSerialize["idempotentHash"] = o.IdempotentHash
	}
	if !IsNil(o.IsDefaultAccount) {
		toSerialize["is_default_account"] = o.IsDefaultAccount
	}
	toSerialize["partner"] = o.Partner
	toSerialize["partner_account_link_reference_token"] = o.PartnerAccountLinkReferenceToken
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.UserToken) {
		toSerialize["user_token"] = o.UserToken
	}
	return toSerialize, nil
}

func (o *AchPartnerRequestModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"partner",
		"partner_account_link_reference_token",
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

	varAchPartnerRequestModel := _AchPartnerRequestModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAchPartnerRequestModel)

	if err != nil {
		return err
	}

	*o = AchPartnerRequestModel(varAchPartnerRequestModel)

	return err
}

type NullableAchPartnerRequestModel struct {
	value *AchPartnerRequestModel
	isSet bool
}

func (v NullableAchPartnerRequestModel) Get() *AchPartnerRequestModel {
	return v.value
}

func (v *NullableAchPartnerRequestModel) Set(val *AchPartnerRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAchPartnerRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAchPartnerRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAchPartnerRequestModel(val *AchPartnerRequestModel) *NullableAchPartnerRequestModel {
	return &NullableAchPartnerRequestModel{value: val, isSet: true}
}

func (v NullableAchPartnerRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAchPartnerRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


