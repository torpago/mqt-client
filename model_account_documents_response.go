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

// checks if the AccountDocumentsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountDocumentsResponse{}

// AccountDocumentsResponse Account documents response.
type AccountDocumentsResponse struct {
	AccountStatement *AccountDocumentResponse `json:"account_statement,omitempty"`
	BenefitsDisclosure *AccountDocumentResponse `json:"benefits_disclosure,omitempty"`
	CardMemberAgreement *AccountDocumentResponse `json:"card_member_agreement,omitempty"`
	EDisclosure *AccountDocumentResponse `json:"e_disclosure,omitempty"`
	PrivacyPolicy *AccountDocumentResponse `json:"privacy_policy,omitempty"`
	RewardsDisclosure *AccountDocumentResponse `json:"rewards_disclosure,omitempty"`
	SummaryOfCreditTerms *AccountDocumentResponse `json:"summary_of_credit_terms,omitempty"`
	TermsSchedule *AccountDocumentResponse `json:"terms_schedule,omitempty"`
}

// NewAccountDocumentsResponse instantiates a new AccountDocumentsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountDocumentsResponse() *AccountDocumentsResponse {
	this := AccountDocumentsResponse{}
	return &this
}

// NewAccountDocumentsResponseWithDefaults instantiates a new AccountDocumentsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountDocumentsResponseWithDefaults() *AccountDocumentsResponse {
	this := AccountDocumentsResponse{}
	return &this
}

// GetAccountStatement returns the AccountStatement field value if set, zero value otherwise.
func (o *AccountDocumentsResponse) GetAccountStatement() AccountDocumentResponse {
	if o == nil || IsNil(o.AccountStatement) {
		var ret AccountDocumentResponse
		return ret
	}
	return *o.AccountStatement
}

// GetAccountStatementOk returns a tuple with the AccountStatement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDocumentsResponse) GetAccountStatementOk() (*AccountDocumentResponse, bool) {
	if o == nil || IsNil(o.AccountStatement) {
		return nil, false
	}
	return o.AccountStatement, true
}

// HasAccountStatement returns a boolean if a field has been set.
func (o *AccountDocumentsResponse) HasAccountStatement() bool {
	if o != nil && !IsNil(o.AccountStatement) {
		return true
	}

	return false
}

// SetAccountStatement gets a reference to the given AccountDocumentResponse and assigns it to the AccountStatement field.
func (o *AccountDocumentsResponse) SetAccountStatement(v AccountDocumentResponse) {
	o.AccountStatement = &v
}

// GetBenefitsDisclosure returns the BenefitsDisclosure field value if set, zero value otherwise.
func (o *AccountDocumentsResponse) GetBenefitsDisclosure() AccountDocumentResponse {
	if o == nil || IsNil(o.BenefitsDisclosure) {
		var ret AccountDocumentResponse
		return ret
	}
	return *o.BenefitsDisclosure
}

// GetBenefitsDisclosureOk returns a tuple with the BenefitsDisclosure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDocumentsResponse) GetBenefitsDisclosureOk() (*AccountDocumentResponse, bool) {
	if o == nil || IsNil(o.BenefitsDisclosure) {
		return nil, false
	}
	return o.BenefitsDisclosure, true
}

// HasBenefitsDisclosure returns a boolean if a field has been set.
func (o *AccountDocumentsResponse) HasBenefitsDisclosure() bool {
	if o != nil && !IsNil(o.BenefitsDisclosure) {
		return true
	}

	return false
}

// SetBenefitsDisclosure gets a reference to the given AccountDocumentResponse and assigns it to the BenefitsDisclosure field.
func (o *AccountDocumentsResponse) SetBenefitsDisclosure(v AccountDocumentResponse) {
	o.BenefitsDisclosure = &v
}

// GetCardMemberAgreement returns the CardMemberAgreement field value if set, zero value otherwise.
func (o *AccountDocumentsResponse) GetCardMemberAgreement() AccountDocumentResponse {
	if o == nil || IsNil(o.CardMemberAgreement) {
		var ret AccountDocumentResponse
		return ret
	}
	return *o.CardMemberAgreement
}

// GetCardMemberAgreementOk returns a tuple with the CardMemberAgreement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDocumentsResponse) GetCardMemberAgreementOk() (*AccountDocumentResponse, bool) {
	if o == nil || IsNil(o.CardMemberAgreement) {
		return nil, false
	}
	return o.CardMemberAgreement, true
}

// HasCardMemberAgreement returns a boolean if a field has been set.
func (o *AccountDocumentsResponse) HasCardMemberAgreement() bool {
	if o != nil && !IsNil(o.CardMemberAgreement) {
		return true
	}

	return false
}

// SetCardMemberAgreement gets a reference to the given AccountDocumentResponse and assigns it to the CardMemberAgreement field.
func (o *AccountDocumentsResponse) SetCardMemberAgreement(v AccountDocumentResponse) {
	o.CardMemberAgreement = &v
}

// GetEDisclosure returns the EDisclosure field value if set, zero value otherwise.
func (o *AccountDocumentsResponse) GetEDisclosure() AccountDocumentResponse {
	if o == nil || IsNil(o.EDisclosure) {
		var ret AccountDocumentResponse
		return ret
	}
	return *o.EDisclosure
}

// GetEDisclosureOk returns a tuple with the EDisclosure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDocumentsResponse) GetEDisclosureOk() (*AccountDocumentResponse, bool) {
	if o == nil || IsNil(o.EDisclosure) {
		return nil, false
	}
	return o.EDisclosure, true
}

// HasEDisclosure returns a boolean if a field has been set.
func (o *AccountDocumentsResponse) HasEDisclosure() bool {
	if o != nil && !IsNil(o.EDisclosure) {
		return true
	}

	return false
}

// SetEDisclosure gets a reference to the given AccountDocumentResponse and assigns it to the EDisclosure field.
func (o *AccountDocumentsResponse) SetEDisclosure(v AccountDocumentResponse) {
	o.EDisclosure = &v
}

// GetPrivacyPolicy returns the PrivacyPolicy field value if set, zero value otherwise.
func (o *AccountDocumentsResponse) GetPrivacyPolicy() AccountDocumentResponse {
	if o == nil || IsNil(o.PrivacyPolicy) {
		var ret AccountDocumentResponse
		return ret
	}
	return *o.PrivacyPolicy
}

// GetPrivacyPolicyOk returns a tuple with the PrivacyPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDocumentsResponse) GetPrivacyPolicyOk() (*AccountDocumentResponse, bool) {
	if o == nil || IsNil(o.PrivacyPolicy) {
		return nil, false
	}
	return o.PrivacyPolicy, true
}

// HasPrivacyPolicy returns a boolean if a field has been set.
func (o *AccountDocumentsResponse) HasPrivacyPolicy() bool {
	if o != nil && !IsNil(o.PrivacyPolicy) {
		return true
	}

	return false
}

// SetPrivacyPolicy gets a reference to the given AccountDocumentResponse and assigns it to the PrivacyPolicy field.
func (o *AccountDocumentsResponse) SetPrivacyPolicy(v AccountDocumentResponse) {
	o.PrivacyPolicy = &v
}

// GetRewardsDisclosure returns the RewardsDisclosure field value if set, zero value otherwise.
func (o *AccountDocumentsResponse) GetRewardsDisclosure() AccountDocumentResponse {
	if o == nil || IsNil(o.RewardsDisclosure) {
		var ret AccountDocumentResponse
		return ret
	}
	return *o.RewardsDisclosure
}

// GetRewardsDisclosureOk returns a tuple with the RewardsDisclosure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDocumentsResponse) GetRewardsDisclosureOk() (*AccountDocumentResponse, bool) {
	if o == nil || IsNil(o.RewardsDisclosure) {
		return nil, false
	}
	return o.RewardsDisclosure, true
}

// HasRewardsDisclosure returns a boolean if a field has been set.
func (o *AccountDocumentsResponse) HasRewardsDisclosure() bool {
	if o != nil && !IsNil(o.RewardsDisclosure) {
		return true
	}

	return false
}

// SetRewardsDisclosure gets a reference to the given AccountDocumentResponse and assigns it to the RewardsDisclosure field.
func (o *AccountDocumentsResponse) SetRewardsDisclosure(v AccountDocumentResponse) {
	o.RewardsDisclosure = &v
}

// GetSummaryOfCreditTerms returns the SummaryOfCreditTerms field value if set, zero value otherwise.
func (o *AccountDocumentsResponse) GetSummaryOfCreditTerms() AccountDocumentResponse {
	if o == nil || IsNil(o.SummaryOfCreditTerms) {
		var ret AccountDocumentResponse
		return ret
	}
	return *o.SummaryOfCreditTerms
}

// GetSummaryOfCreditTermsOk returns a tuple with the SummaryOfCreditTerms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDocumentsResponse) GetSummaryOfCreditTermsOk() (*AccountDocumentResponse, bool) {
	if o == nil || IsNil(o.SummaryOfCreditTerms) {
		return nil, false
	}
	return o.SummaryOfCreditTerms, true
}

// HasSummaryOfCreditTerms returns a boolean if a field has been set.
func (o *AccountDocumentsResponse) HasSummaryOfCreditTerms() bool {
	if o != nil && !IsNil(o.SummaryOfCreditTerms) {
		return true
	}

	return false
}

// SetSummaryOfCreditTerms gets a reference to the given AccountDocumentResponse and assigns it to the SummaryOfCreditTerms field.
func (o *AccountDocumentsResponse) SetSummaryOfCreditTerms(v AccountDocumentResponse) {
	o.SummaryOfCreditTerms = &v
}

// GetTermsSchedule returns the TermsSchedule field value if set, zero value otherwise.
func (o *AccountDocumentsResponse) GetTermsSchedule() AccountDocumentResponse {
	if o == nil || IsNil(o.TermsSchedule) {
		var ret AccountDocumentResponse
		return ret
	}
	return *o.TermsSchedule
}

// GetTermsScheduleOk returns a tuple with the TermsSchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDocumentsResponse) GetTermsScheduleOk() (*AccountDocumentResponse, bool) {
	if o == nil || IsNil(o.TermsSchedule) {
		return nil, false
	}
	return o.TermsSchedule, true
}

// HasTermsSchedule returns a boolean if a field has been set.
func (o *AccountDocumentsResponse) HasTermsSchedule() bool {
	if o != nil && !IsNil(o.TermsSchedule) {
		return true
	}

	return false
}

// SetTermsSchedule gets a reference to the given AccountDocumentResponse and assigns it to the TermsSchedule field.
func (o *AccountDocumentsResponse) SetTermsSchedule(v AccountDocumentResponse) {
	o.TermsSchedule = &v
}

func (o AccountDocumentsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountDocumentsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountStatement) {
		toSerialize["account_statement"] = o.AccountStatement
	}
	if !IsNil(o.BenefitsDisclosure) {
		toSerialize["benefits_disclosure"] = o.BenefitsDisclosure
	}
	if !IsNil(o.CardMemberAgreement) {
		toSerialize["card_member_agreement"] = o.CardMemberAgreement
	}
	if !IsNil(o.EDisclosure) {
		toSerialize["e_disclosure"] = o.EDisclosure
	}
	if !IsNil(o.PrivacyPolicy) {
		toSerialize["privacy_policy"] = o.PrivacyPolicy
	}
	if !IsNil(o.RewardsDisclosure) {
		toSerialize["rewards_disclosure"] = o.RewardsDisclosure
	}
	if !IsNil(o.SummaryOfCreditTerms) {
		toSerialize["summary_of_credit_terms"] = o.SummaryOfCreditTerms
	}
	if !IsNil(o.TermsSchedule) {
		toSerialize["terms_schedule"] = o.TermsSchedule
	}
	return toSerialize, nil
}

type NullableAccountDocumentsResponse struct {
	value *AccountDocumentsResponse
	isSet bool
}

func (v NullableAccountDocumentsResponse) Get() *AccountDocumentsResponse {
	return v.value
}

func (v *NullableAccountDocumentsResponse) Set(val *AccountDocumentsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountDocumentsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountDocumentsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountDocumentsResponse(val *AccountDocumentsResponse) *NullableAccountDocumentsResponse {
	return &NullableAccountDocumentsResponse{value: val, isSet: true}
}

func (v NullableAccountDocumentsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountDocumentsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


