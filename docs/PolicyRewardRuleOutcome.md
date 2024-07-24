# PolicyRewardRuleOutcome

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxAmount** | Pointer to **decimal.Decimal** | Max amount of the reward. | [optional] 
**Percentage** | **decimal.Decimal** | Reward percentage applied when the balance for a billing cycle is within the range specified in the &#x60;filters.amount.greater_than&#x60; and &#x60;filters.amount.less_than&#x60; fields. For example, if the percentage is &#x60;1&#x60;, the account holder earns 1% of the account balance if they spend between the &#x60;greater_than&#x60; and &#x60;less_than&#x60; amounts during a billing cycle. | 

## Methods

### NewPolicyRewardRuleOutcome

`func NewPolicyRewardRuleOutcome(percentage decimal.Decimal, ) *PolicyRewardRuleOutcome`

NewPolicyRewardRuleOutcome instantiates a new PolicyRewardRuleOutcome object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyRewardRuleOutcomeWithDefaults

`func NewPolicyRewardRuleOutcomeWithDefaults() *PolicyRewardRuleOutcome`

NewPolicyRewardRuleOutcomeWithDefaults instantiates a new PolicyRewardRuleOutcome object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxAmount

`func (o *PolicyRewardRuleOutcome) GetMaxAmount() decimal.Decimal`

GetMaxAmount returns the MaxAmount field if non-nil, zero value otherwise.

### GetMaxAmountOk

`func (o *PolicyRewardRuleOutcome) GetMaxAmountOk() (*decimal.Decimal, bool)`

GetMaxAmountOk returns a tuple with the MaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAmount

`func (o *PolicyRewardRuleOutcome) SetMaxAmount(v decimal.Decimal)`

SetMaxAmount sets MaxAmount field to given value.

### HasMaxAmount

`func (o *PolicyRewardRuleOutcome) HasMaxAmount() bool`

HasMaxAmount returns a boolean if a field has been set.

### GetPercentage

`func (o *PolicyRewardRuleOutcome) GetPercentage() decimal.Decimal`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *PolicyRewardRuleOutcome) GetPercentageOk() (*decimal.Decimal, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *PolicyRewardRuleOutcome) SetPercentage(v decimal.Decimal)`

SetPercentage sets Percentage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


