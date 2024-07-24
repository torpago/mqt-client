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

// checks if the PolicyDocumentCreateReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicyDocumentCreateReq{}

// PolicyDocumentCreateReq Contains information on a document policy.
type PolicyDocumentCreateReq struct {
	AccountStatement PolicyDocumentTemplateReq `json:"account_statement"`
	BenefitsDisclosurePremium PolicyDocumentAssetReq `json:"benefits_disclosure_premium"`
	BenefitsDisclosureTraditional PolicyDocumentAssetReq `json:"benefits_disclosure_traditional"`
	CardMemberAgreement PolicyDocumentAssetReq `json:"card_member_agreement"`
	EDisclosure PolicyDocumentAssetReq `json:"e_disclosure"`
	// Name of the document policy.
	Name string `json:"name" validate:"regexp=(?!^ +$)^.+$"`
	NoaaMultipleReasonWithDoddFrank PolicyDocumentTemplateReq `json:"noaa_multiple_reason_with_dodd_frank"`
	NoaaSingleReason PolicyDocumentTemplateReq `json:"noaa_single_reason"`
	NoaaSingleReasonWithDoddFrank PolicyDocumentTemplateReq `json:"noaa_single_reason_with_dodd_frank"`
	PreQualificationDisclosure *PolicyDocumentAssetAndTemplateReq `json:"pre_qualification_disclosure,omitempty"`
	PrivacyPolicy PolicyDocumentAssetReq `json:"privacy_policy"`
	RewardsDisclosure *PolicyDocumentAssetAndTemplateReq `json:"rewards_disclosure,omitempty"`
	SummaryOfCreditTerms PolicyDocumentAssetAndTemplateReq `json:"summary_of_credit_terms"`
	TermsSchedule PolicyDocumentTemplateReq `json:"terms_schedule"`
	// Unique identifier of the document policy.
	Token *string `json:"token,omitempty" validate:"regexp=(?!^ +$)^.+$"`
}

type _PolicyDocumentCreateReq PolicyDocumentCreateReq

// NewPolicyDocumentCreateReq instantiates a new PolicyDocumentCreateReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyDocumentCreateReq(accountStatement PolicyDocumentTemplateReq, benefitsDisclosurePremium PolicyDocumentAssetReq, benefitsDisclosureTraditional PolicyDocumentAssetReq, cardMemberAgreement PolicyDocumentAssetReq, eDisclosure PolicyDocumentAssetReq, name string, noaaMultipleReasonWithDoddFrank PolicyDocumentTemplateReq, noaaSingleReason PolicyDocumentTemplateReq, noaaSingleReasonWithDoddFrank PolicyDocumentTemplateReq, privacyPolicy PolicyDocumentAssetReq, summaryOfCreditTerms PolicyDocumentAssetAndTemplateReq, termsSchedule PolicyDocumentTemplateReq) *PolicyDocumentCreateReq {
	this := PolicyDocumentCreateReq{}
	this.AccountStatement = accountStatement
	this.BenefitsDisclosurePremium = benefitsDisclosurePremium
	this.BenefitsDisclosureTraditional = benefitsDisclosureTraditional
	this.CardMemberAgreement = cardMemberAgreement
	this.EDisclosure = eDisclosure
	this.Name = name
	this.NoaaMultipleReasonWithDoddFrank = noaaMultipleReasonWithDoddFrank
	this.NoaaSingleReason = noaaSingleReason
	this.NoaaSingleReasonWithDoddFrank = noaaSingleReasonWithDoddFrank
	this.PrivacyPolicy = privacyPolicy
	this.SummaryOfCreditTerms = summaryOfCreditTerms
	this.TermsSchedule = termsSchedule
	return &this
}

// NewPolicyDocumentCreateReqWithDefaults instantiates a new PolicyDocumentCreateReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyDocumentCreateReqWithDefaults() *PolicyDocumentCreateReq {
	this := PolicyDocumentCreateReq{}
	return &this
}

// GetAccountStatement returns the AccountStatement field value
func (o *PolicyDocumentCreateReq) GetAccountStatement() PolicyDocumentTemplateReq {
	if o == nil {
		var ret PolicyDocumentTemplateReq
		return ret
	}

	return o.AccountStatement
}

// GetAccountStatementOk returns a tuple with the AccountStatement field value
// and a boolean to check if the value has been set.
func (o *PolicyDocumentCreateReq) GetAccountStatementOk() (*PolicyDocumentTemplateReq, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountStatement, true
}

// SetAccountStatement sets field value
func (o *PolicyDocumentCreateReq) SetAccountStatement(v PolicyDocumentTemplateReq) {
	o.AccountStatement = v
}

// GetBenefitsDisclosurePremium returns the BenefitsDisclosurePremium field value
func (o *PolicyDocumentCreateReq) GetBenefitsDisclosurePremium() PolicyDocumentAssetReq {
	if o == nil {
		var ret PolicyDocumentAssetReq
		return ret
	}

	return o.BenefitsDisclosurePremium
}

// GetBenefitsDisclosurePremiumOk returns a tuple with the BenefitsDisclosurePremium field value
// and a boolean to check if the value has been set.
func (o *PolicyDocumentCreateReq) GetBenefitsDisclosurePremiumOk() (*PolicyDocumentAssetReq, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BenefitsDisclosurePremium, true
}

// SetBenefitsDisclosurePremium sets field value
func (o *PolicyDocumentCreateReq) SetBenefitsDisclosurePremium(v PolicyDocumentAssetReq) {
	o.BenefitsDisclosurePremium = v
}

// GetBenefitsDisclosureTraditional returns the BenefitsDisclosureTraditional field value
func (o *PolicyDocumentCreateReq) GetBenefitsDisclosureTraditional() PolicyDocumentAssetReq {
	if o == nil {
		var ret PolicyDocumentAssetReq
		return ret
	}

	return o.BenefitsDisclosureTraditional
}

// GetBenefitsDisclosureTraditionalOk returns a tuple with the BenefitsDisclosureTraditional field value
// and a boolean to check if the value has been set.
func (o *PolicyDocumentCreateReq) GetBenefitsDisclosureTraditionalOk() (*PolicyDocumentAssetReq, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BenefitsDisclosureTraditional, true
}

// SetBenefitsDisclosureTraditional sets field value
func (o *PolicyDocumentCreateReq) SetBenefitsDisclosureTraditional(v PolicyDocumentAssetReq) {
	o.BenefitsDisclosureTraditional = v
}

// GetCardMemberAgreement returns the CardMemberAgreement field value
func (o *PolicyDocumentCreateReq) GetCardMemberAgreement() PolicyDocumentAssetReq {
	if o == nil {
		var ret PolicyDocumentAssetReq
		return ret
	}

	return o.CardMemberAgreement
}

// GetCardMemberAgreementOk returns a tuple with the CardMemberAgreement field value
// and a boolean to check if the value has been set.
func (o *PolicyDocumentCreateReq) GetCardMemberAgreementOk() (*PolicyDocumentAssetReq, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardMemberAgreement, true
}

// SetCardMemberAgreement sets field value
func (o *PolicyDocumentCreateReq) SetCardMemberAgreement(v PolicyDocumentAssetReq) {
	o.CardMemberAgreement = v
}

// GetEDisclosure returns the EDisclosure field value
func (o *PolicyDocumentCreateReq) GetEDisclosure() PolicyDocumentAssetReq {
	if o == nil {
		var ret PolicyDocumentAssetReq
		return ret
	}

	return o.EDisclosure
}

// GetEDisclosureOk returns a tuple with the EDisclosure field value
// and a boolean to check if the value has been set.
func (o *PolicyDocumentCreateReq) GetEDisclosureOk() (*PolicyDocumentAssetReq, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EDisclosure, true
}

// SetEDisclosure sets field value
func (o *PolicyDocumentCreateReq) SetEDisclosure(v PolicyDocumentAssetReq) {
	o.EDisclosure = v
}

// GetName returns the Name field value
func (o *PolicyDocumentCreateReq) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PolicyDocumentCreateReq) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PolicyDocumentCreateReq) SetName(v string) {
	o.Name = v
}

// GetNoaaMultipleReasonWithDoddFrank returns the NoaaMultipleReasonWithDoddFrank field value
func (o *PolicyDocumentCreateReq) GetNoaaMultipleReasonWithDoddFrank() PolicyDocumentTemplateReq {
	if o == nil {
		var ret PolicyDocumentTemplateReq
		return ret
	}

	return o.NoaaMultipleReasonWithDoddFrank
}

// GetNoaaMultipleReasonWithDoddFrankOk returns a tuple with the NoaaMultipleReasonWithDoddFrank field value
// and a boolean to check if the value has been set.
func (o *PolicyDocumentCreateReq) GetNoaaMultipleReasonWithDoddFrankOk() (*PolicyDocumentTemplateReq, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NoaaMultipleReasonWithDoddFrank, true
}

// SetNoaaMultipleReasonWithDoddFrank sets field value
func (o *PolicyDocumentCreateReq) SetNoaaMultipleReasonWithDoddFrank(v PolicyDocumentTemplateReq) {
	o.NoaaMultipleReasonWithDoddFrank = v
}

// GetNoaaSingleReason returns the NoaaSingleReason field value
func (o *PolicyDocumentCreateReq) GetNoaaSingleReason() PolicyDocumentTemplateReq {
	if o == nil {
		var ret PolicyDocumentTemplateReq
		return ret
	}

	return o.NoaaSingleReason
}

// GetNoaaSingleReasonOk returns a tuple with the NoaaSingleReason field value
// and a boolean to check if the value has been set.
func (o *PolicyDocumentCreateReq) GetNoaaSingleReasonOk() (*PolicyDocumentTemplateReq, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NoaaSingleReason, true
}

// SetNoaaSingleReason sets field value
func (o *PolicyDocumentCreateReq) SetNoaaSingleReason(v PolicyDocumentTemplateReq) {
	o.NoaaSingleReason = v
}

// GetNoaaSingleReasonWithDoddFrank returns the NoaaSingleReasonWithDoddFrank field value
func (o *PolicyDocumentCreateReq) GetNoaaSingleReasonWithDoddFrank() PolicyDocumentTemplateReq {
	if o == nil {
		var ret PolicyDocumentTemplateReq
		return ret
	}

	return o.NoaaSingleReasonWithDoddFrank
}

// GetNoaaSingleReasonWithDoddFrankOk returns a tuple with the NoaaSingleReasonWithDoddFrank field value
// and a boolean to check if the value has been set.
func (o *PolicyDocumentCreateReq) GetNoaaSingleReasonWithDoddFrankOk() (*PolicyDocumentTemplateReq, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NoaaSingleReasonWithDoddFrank, true
}

// SetNoaaSingleReasonWithDoddFrank sets field value
func (o *PolicyDocumentCreateReq) SetNoaaSingleReasonWithDoddFrank(v PolicyDocumentTemplateReq) {
	o.NoaaSingleReasonWithDoddFrank = v
}

// GetPreQualificationDisclosure returns the PreQualificationDisclosure field value if set, zero value otherwise.
func (o *PolicyDocumentCreateReq) GetPreQualificationDisclosure() PolicyDocumentAssetAndTemplateReq {
	if o == nil || IsNil(o.PreQualificationDisclosure) {
		var ret PolicyDocumentAssetAndTemplateReq
		return ret
	}
	return *o.PreQualificationDisclosure
}

// GetPreQualificationDisclosureOk returns a tuple with the PreQualificationDisclosure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyDocumentCreateReq) GetPreQualificationDisclosureOk() (*PolicyDocumentAssetAndTemplateReq, bool) {
	if o == nil || IsNil(o.PreQualificationDisclosure) {
		return nil, false
	}
	return o.PreQualificationDisclosure, true
}

// HasPreQualificationDisclosure returns a boolean if a field has been set.
func (o *PolicyDocumentCreateReq) HasPreQualificationDisclosure() bool {
	if o != nil && !IsNil(o.PreQualificationDisclosure) {
		return true
	}

	return false
}

// SetPreQualificationDisclosure gets a reference to the given PolicyDocumentAssetAndTemplateReq and assigns it to the PreQualificationDisclosure field.
func (o *PolicyDocumentCreateReq) SetPreQualificationDisclosure(v PolicyDocumentAssetAndTemplateReq) {
	o.PreQualificationDisclosure = &v
}

// GetPrivacyPolicy returns the PrivacyPolicy field value
func (o *PolicyDocumentCreateReq) GetPrivacyPolicy() PolicyDocumentAssetReq {
	if o == nil {
		var ret PolicyDocumentAssetReq
		return ret
	}

	return o.PrivacyPolicy
}

// GetPrivacyPolicyOk returns a tuple with the PrivacyPolicy field value
// and a boolean to check if the value has been set.
func (o *PolicyDocumentCreateReq) GetPrivacyPolicyOk() (*PolicyDocumentAssetReq, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivacyPolicy, true
}

// SetPrivacyPolicy sets field value
func (o *PolicyDocumentCreateReq) SetPrivacyPolicy(v PolicyDocumentAssetReq) {
	o.PrivacyPolicy = v
}

// GetRewardsDisclosure returns the RewardsDisclosure field value if set, zero value otherwise.
func (o *PolicyDocumentCreateReq) GetRewardsDisclosure() PolicyDocumentAssetAndTemplateReq {
	if o == nil || IsNil(o.RewardsDisclosure) {
		var ret PolicyDocumentAssetAndTemplateReq
		return ret
	}
	return *o.RewardsDisclosure
}

// GetRewardsDisclosureOk returns a tuple with the RewardsDisclosure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyDocumentCreateReq) GetRewardsDisclosureOk() (*PolicyDocumentAssetAndTemplateReq, bool) {
	if o == nil || IsNil(o.RewardsDisclosure) {
		return nil, false
	}
	return o.RewardsDisclosure, true
}

// HasRewardsDisclosure returns a boolean if a field has been set.
func (o *PolicyDocumentCreateReq) HasRewardsDisclosure() bool {
	if o != nil && !IsNil(o.RewardsDisclosure) {
		return true
	}

	return false
}

// SetRewardsDisclosure gets a reference to the given PolicyDocumentAssetAndTemplateReq and assigns it to the RewardsDisclosure field.
func (o *PolicyDocumentCreateReq) SetRewardsDisclosure(v PolicyDocumentAssetAndTemplateReq) {
	o.RewardsDisclosure = &v
}

// GetSummaryOfCreditTerms returns the SummaryOfCreditTerms field value
func (o *PolicyDocumentCreateReq) GetSummaryOfCreditTerms() PolicyDocumentAssetAndTemplateReq {
	if o == nil {
		var ret PolicyDocumentAssetAndTemplateReq
		return ret
	}

	return o.SummaryOfCreditTerms
}

// GetSummaryOfCreditTermsOk returns a tuple with the SummaryOfCreditTerms field value
// and a boolean to check if the value has been set.
func (o *PolicyDocumentCreateReq) GetSummaryOfCreditTermsOk() (*PolicyDocumentAssetAndTemplateReq, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SummaryOfCreditTerms, true
}

// SetSummaryOfCreditTerms sets field value
func (o *PolicyDocumentCreateReq) SetSummaryOfCreditTerms(v PolicyDocumentAssetAndTemplateReq) {
	o.SummaryOfCreditTerms = v
}

// GetTermsSchedule returns the TermsSchedule field value
func (o *PolicyDocumentCreateReq) GetTermsSchedule() PolicyDocumentTemplateReq {
	if o == nil {
		var ret PolicyDocumentTemplateReq
		return ret
	}

	return o.TermsSchedule
}

// GetTermsScheduleOk returns a tuple with the TermsSchedule field value
// and a boolean to check if the value has been set.
func (o *PolicyDocumentCreateReq) GetTermsScheduleOk() (*PolicyDocumentTemplateReq, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TermsSchedule, true
}

// SetTermsSchedule sets field value
func (o *PolicyDocumentCreateReq) SetTermsSchedule(v PolicyDocumentTemplateReq) {
	o.TermsSchedule = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *PolicyDocumentCreateReq) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyDocumentCreateReq) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *PolicyDocumentCreateReq) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *PolicyDocumentCreateReq) SetToken(v string) {
	o.Token = &v
}

func (o PolicyDocumentCreateReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyDocumentCreateReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account_statement"] = o.AccountStatement
	toSerialize["benefits_disclosure_premium"] = o.BenefitsDisclosurePremium
	toSerialize["benefits_disclosure_traditional"] = o.BenefitsDisclosureTraditional
	toSerialize["card_member_agreement"] = o.CardMemberAgreement
	toSerialize["e_disclosure"] = o.EDisclosure
	toSerialize["name"] = o.Name
	toSerialize["noaa_multiple_reason_with_dodd_frank"] = o.NoaaMultipleReasonWithDoddFrank
	toSerialize["noaa_single_reason"] = o.NoaaSingleReason
	toSerialize["noaa_single_reason_with_dodd_frank"] = o.NoaaSingleReasonWithDoddFrank
	if !IsNil(o.PreQualificationDisclosure) {
		toSerialize["pre_qualification_disclosure"] = o.PreQualificationDisclosure
	}
	toSerialize["privacy_policy"] = o.PrivacyPolicy
	if !IsNil(o.RewardsDisclosure) {
		toSerialize["rewards_disclosure"] = o.RewardsDisclosure
	}
	toSerialize["summary_of_credit_terms"] = o.SummaryOfCreditTerms
	toSerialize["terms_schedule"] = o.TermsSchedule
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	return toSerialize, nil
}

func (o *PolicyDocumentCreateReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"account_statement",
		"benefits_disclosure_premium",
		"benefits_disclosure_traditional",
		"card_member_agreement",
		"e_disclosure",
		"name",
		"noaa_multiple_reason_with_dodd_frank",
		"noaa_single_reason",
		"noaa_single_reason_with_dodd_frank",
		"privacy_policy",
		"summary_of_credit_terms",
		"terms_schedule",
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

	varPolicyDocumentCreateReq := _PolicyDocumentCreateReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPolicyDocumentCreateReq)

	if err != nil {
		return err
	}

	*o = PolicyDocumentCreateReq(varPolicyDocumentCreateReq)

	return err
}

type NullablePolicyDocumentCreateReq struct {
	value *PolicyDocumentCreateReq
	isSet bool
}

func (v NullablePolicyDocumentCreateReq) Get() *PolicyDocumentCreateReq {
	return v.value
}

func (v *NullablePolicyDocumentCreateReq) Set(val *PolicyDocumentCreateReq) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyDocumentCreateReq) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyDocumentCreateReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyDocumentCreateReq(val *PolicyDocumentCreateReq) *NullablePolicyDocumentCreateReq {
	return &NullablePolicyDocumentCreateReq{value: val, isSet: true}
}

func (v NullablePolicyDocumentCreateReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyDocumentCreateReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


