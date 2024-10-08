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

// checks if the IssuerFraudView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IssuerFraudView{}

// IssuerFraudView Contains one or more fraud determinations by the card network that apply to either the transaction or the cardholder's account.
type IssuerFraudView struct {
	FraudScoreReasons []string `json:"fraud_score_reasons,omitempty"`
	// The action recommended based on the fraud score.
	RecommendedAction *string `json:"recommended_action,omitempty"`
	// The fraud rating, or level of risk associated with the transaction.
	RiskLevel *string `json:"risk_level,omitempty"`
	// The RiskControl tags that were triggered by the transaction.
	RiskcontrolTags []RiskcontrolTags `json:"riskcontrol_tags,omitempty"`
	// The rules violated by the transaction.
	RuleViolations []string `json:"rule_violations,omitempty"`
	// The risk score generated by RiskControl. This is either the Mastercard Decision Intelligence or Visa Advance Authorization transaction risk score.
	Score *int32 `json:"score,omitempty"`
	// Provides a list of rules triggered by a fraud event, along with the information on tags and rule characteristics.
	TriggeredRules []TriggeredRule `json:"triggered_rules,omitempty"`
}

// NewIssuerFraudView instantiates a new IssuerFraudView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssuerFraudView() *IssuerFraudView {
	this := IssuerFraudView{}
	return &this
}

// NewIssuerFraudViewWithDefaults instantiates a new IssuerFraudView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssuerFraudViewWithDefaults() *IssuerFraudView {
	this := IssuerFraudView{}
	return &this
}

// GetFraudScoreReasons returns the FraudScoreReasons field value if set, zero value otherwise.
func (o *IssuerFraudView) GetFraudScoreReasons() []string {
	if o == nil || IsNil(o.FraudScoreReasons) {
		var ret []string
		return ret
	}
	return o.FraudScoreReasons
}

// GetFraudScoreReasonsOk returns a tuple with the FraudScoreReasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuerFraudView) GetFraudScoreReasonsOk() ([]string, bool) {
	if o == nil || IsNil(o.FraudScoreReasons) {
		return nil, false
	}
	return o.FraudScoreReasons, true
}

// HasFraudScoreReasons returns a boolean if a field has been set.
func (o *IssuerFraudView) HasFraudScoreReasons() bool {
	if o != nil && !IsNil(o.FraudScoreReasons) {
		return true
	}

	return false
}

// SetFraudScoreReasons gets a reference to the given []string and assigns it to the FraudScoreReasons field.
func (o *IssuerFraudView) SetFraudScoreReasons(v []string) {
	o.FraudScoreReasons = v
}

// GetRecommendedAction returns the RecommendedAction field value if set, zero value otherwise.
func (o *IssuerFraudView) GetRecommendedAction() string {
	if o == nil || IsNil(o.RecommendedAction) {
		var ret string
		return ret
	}
	return *o.RecommendedAction
}

// GetRecommendedActionOk returns a tuple with the RecommendedAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuerFraudView) GetRecommendedActionOk() (*string, bool) {
	if o == nil || IsNil(o.RecommendedAction) {
		return nil, false
	}
	return o.RecommendedAction, true
}

// HasRecommendedAction returns a boolean if a field has been set.
func (o *IssuerFraudView) HasRecommendedAction() bool {
	if o != nil && !IsNil(o.RecommendedAction) {
		return true
	}

	return false
}

// SetRecommendedAction gets a reference to the given string and assigns it to the RecommendedAction field.
func (o *IssuerFraudView) SetRecommendedAction(v string) {
	o.RecommendedAction = &v
}

// GetRiskLevel returns the RiskLevel field value if set, zero value otherwise.
func (o *IssuerFraudView) GetRiskLevel() string {
	if o == nil || IsNil(o.RiskLevel) {
		var ret string
		return ret
	}
	return *o.RiskLevel
}

// GetRiskLevelOk returns a tuple with the RiskLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuerFraudView) GetRiskLevelOk() (*string, bool) {
	if o == nil || IsNil(o.RiskLevel) {
		return nil, false
	}
	return o.RiskLevel, true
}

// HasRiskLevel returns a boolean if a field has been set.
func (o *IssuerFraudView) HasRiskLevel() bool {
	if o != nil && !IsNil(o.RiskLevel) {
		return true
	}

	return false
}

// SetRiskLevel gets a reference to the given string and assigns it to the RiskLevel field.
func (o *IssuerFraudView) SetRiskLevel(v string) {
	o.RiskLevel = &v
}

// GetRiskcontrolTags returns the RiskcontrolTags field value if set, zero value otherwise.
func (o *IssuerFraudView) GetRiskcontrolTags() []RiskcontrolTags {
	if o == nil || IsNil(o.RiskcontrolTags) {
		var ret []RiskcontrolTags
		return ret
	}
	return o.RiskcontrolTags
}

// GetRiskcontrolTagsOk returns a tuple with the RiskcontrolTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuerFraudView) GetRiskcontrolTagsOk() ([]RiskcontrolTags, bool) {
	if o == nil || IsNil(o.RiskcontrolTags) {
		return nil, false
	}
	return o.RiskcontrolTags, true
}

// HasRiskcontrolTags returns a boolean if a field has been set.
func (o *IssuerFraudView) HasRiskcontrolTags() bool {
	if o != nil && !IsNil(o.RiskcontrolTags) {
		return true
	}

	return false
}

// SetRiskcontrolTags gets a reference to the given []RiskcontrolTags and assigns it to the RiskcontrolTags field.
func (o *IssuerFraudView) SetRiskcontrolTags(v []RiskcontrolTags) {
	o.RiskcontrolTags = v
}

// GetRuleViolations returns the RuleViolations field value if set, zero value otherwise.
func (o *IssuerFraudView) GetRuleViolations() []string {
	if o == nil || IsNil(o.RuleViolations) {
		var ret []string
		return ret
	}
	return o.RuleViolations
}

// GetRuleViolationsOk returns a tuple with the RuleViolations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuerFraudView) GetRuleViolationsOk() ([]string, bool) {
	if o == nil || IsNil(o.RuleViolations) {
		return nil, false
	}
	return o.RuleViolations, true
}

// HasRuleViolations returns a boolean if a field has been set.
func (o *IssuerFraudView) HasRuleViolations() bool {
	if o != nil && !IsNil(o.RuleViolations) {
		return true
	}

	return false
}

// SetRuleViolations gets a reference to the given []string and assigns it to the RuleViolations field.
func (o *IssuerFraudView) SetRuleViolations(v []string) {
	o.RuleViolations = v
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *IssuerFraudView) GetScore() int32 {
	if o == nil || IsNil(o.Score) {
		var ret int32
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuerFraudView) GetScoreOk() (*int32, bool) {
	if o == nil || IsNil(o.Score) {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *IssuerFraudView) HasScore() bool {
	if o != nil && !IsNil(o.Score) {
		return true
	}

	return false
}

// SetScore gets a reference to the given int32 and assigns it to the Score field.
func (o *IssuerFraudView) SetScore(v int32) {
	o.Score = &v
}

// GetTriggeredRules returns the TriggeredRules field value if set, zero value otherwise.
func (o *IssuerFraudView) GetTriggeredRules() []TriggeredRule {
	if o == nil || IsNil(o.TriggeredRules) {
		var ret []TriggeredRule
		return ret
	}
	return o.TriggeredRules
}

// GetTriggeredRulesOk returns a tuple with the TriggeredRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuerFraudView) GetTriggeredRulesOk() ([]TriggeredRule, bool) {
	if o == nil || IsNil(o.TriggeredRules) {
		return nil, false
	}
	return o.TriggeredRules, true
}

// HasTriggeredRules returns a boolean if a field has been set.
func (o *IssuerFraudView) HasTriggeredRules() bool {
	if o != nil && !IsNil(o.TriggeredRules) {
		return true
	}

	return false
}

// SetTriggeredRules gets a reference to the given []TriggeredRule and assigns it to the TriggeredRules field.
func (o *IssuerFraudView) SetTriggeredRules(v []TriggeredRule) {
	o.TriggeredRules = v
}

func (o IssuerFraudView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IssuerFraudView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FraudScoreReasons) {
		toSerialize["fraud_score_reasons"] = o.FraudScoreReasons
	}
	if !IsNil(o.RecommendedAction) {
		toSerialize["recommended_action"] = o.RecommendedAction
	}
	if !IsNil(o.RiskLevel) {
		toSerialize["risk_level"] = o.RiskLevel
	}
	if !IsNil(o.RiskcontrolTags) {
		toSerialize["riskcontrol_tags"] = o.RiskcontrolTags
	}
	if !IsNil(o.RuleViolations) {
		toSerialize["rule_violations"] = o.RuleViolations
	}
	if !IsNil(o.Score) {
		toSerialize["score"] = o.Score
	}
	if !IsNil(o.TriggeredRules) {
		toSerialize["triggered_rules"] = o.TriggeredRules
	}
	return toSerialize, nil
}

type NullableIssuerFraudView struct {
	value *IssuerFraudView
	isSet bool
}

func (v NullableIssuerFraudView) Get() *IssuerFraudView {
	return v.value
}

func (v *NullableIssuerFraudView) Set(val *IssuerFraudView) {
	v.value = val
	v.isSet = true
}

func (v NullableIssuerFraudView) IsSet() bool {
	return v.isSet
}

func (v *NullableIssuerFraudView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssuerFraudView(val *IssuerFraudView) *NullableIssuerFraudView {
	return &NullableIssuerFraudView{value: val, isSet: true}
}

func (v NullableIssuerFraudView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssuerFraudView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


