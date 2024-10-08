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
	"github.com/shopspring/decimal"
)

// checks if the ProductResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductResponse{}

// ProductResponse Specifies shared details for a credit program.  Once set to `ACTIVE`, cannot be edited or deleted.
type ProductResponse struct {
	// One or more associated card product tokens.
	CardProductTokens []string `json:"card_product_tokens,omitempty"`
	Classification *ProductClassification `json:"classification,omitempty"`
	Config *ProductConfig `json:"config,omitempty"`
	// Date and time when the credit product was created on Marqeta's credit platform, in UTC.
	CreatedTime *time.Time `json:"created_time,omitempty"`
	CreditLine *ProductCreditLine `json:"credit_line,omitempty"`
	CurrencyCode *CurrencyCode `json:"currency_code,omitempty"`
	// Description of the credit product.
	Description *string `json:"description,omitempty"`
	InterestCalculation *InterestCalculation `json:"interest_calculation,omitempty"`
	MinPaymentCalculation *ProductMinPaymentCalculation `json:"min_payment_calculation,omitempty"`
	// Minimum payment, expressed as a flat amount, due on the payment due day.
	MinPaymentFlatAmount *decimal.Decimal `json:"min_payment_flat_amount,omitempty"`
	// Minimum payment, expressed as a percentage of the total statement balance, due on the payment due day.
	MinPaymentPercentage *decimal.Decimal `json:"min_payment_percentage,omitempty"`
	// Name of the credit product.
	Name *string `json:"name,omitempty"`
	// Unique identifier of the parent credit product.
	ParentProductToken *string `json:"parent_product_token,omitempty"`
	// Ordered list of balance types to which payments are allocated, from first to last.
	PaymentAllocationOrder []PaymentAllocationOrderEnum `json:"payment_allocation_order,omitempty"`
	ProductSubType *ProductSubType `json:"product_sub_type,omitempty"`
	ProductType *ProductType `json:"product_type,omitempty"`
	Status *ResourceStatus `json:"status,omitempty"`
	// Unique identifier of the credit product.
	Token *string `json:"token,omitempty"`
	// Date and time when the credit product was last updated on Marqeta's credit platform, in UTC.
	UpdatedTime *time.Time `json:"updated_time,omitempty"`
	// One or more usage types for the credit product.
	Usage []BalanceType `json:"usage,omitempty"`
}

// NewProductResponse instantiates a new ProductResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductResponse() *ProductResponse {
	this := ProductResponse{}
	var classification ProductClassification = PRODUCTCLASSIFICATION_CONSUMER
	this.Classification = &classification
	var currencyCode CurrencyCode = CURRENCYCODE_USD
	this.CurrencyCode = &currencyCode
	var productSubType ProductSubType = PRODUCTSUBTYPE_CREDIT_CARD
	this.ProductSubType = &productSubType
	var productType ProductType = PRODUCTTYPE_REVOLVING
	this.ProductType = &productType
	return &this
}

// NewProductResponseWithDefaults instantiates a new ProductResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductResponseWithDefaults() *ProductResponse {
	this := ProductResponse{}
	var classification ProductClassification = PRODUCTCLASSIFICATION_CONSUMER
	this.Classification = &classification
	var currencyCode CurrencyCode = CURRENCYCODE_USD
	this.CurrencyCode = &currencyCode
	var productSubType ProductSubType = PRODUCTSUBTYPE_CREDIT_CARD
	this.ProductSubType = &productSubType
	var productType ProductType = PRODUCTTYPE_REVOLVING
	this.ProductType = &productType
	return &this
}

// GetCardProductTokens returns the CardProductTokens field value if set, zero value otherwise.
func (o *ProductResponse) GetCardProductTokens() []string {
	if o == nil || IsNil(o.CardProductTokens) {
		var ret []string
		return ret
	}
	return o.CardProductTokens
}

// GetCardProductTokensOk returns a tuple with the CardProductTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductResponse) GetCardProductTokensOk() ([]string, bool) {
	if o == nil || IsNil(o.CardProductTokens) {
		return nil, false
	}
	return o.CardProductTokens, true
}

// HasCardProductTokens returns a boolean if a field has been set.
func (o *ProductResponse) HasCardProductTokens() bool {
	if o != nil && !IsNil(o.CardProductTokens) {
		return true
	}

	return false
}

// SetCardProductTokens gets a reference to the given []string and assigns it to the CardProductTokens field.
func (o *ProductResponse) SetCardProductTokens(v []string) {
	o.CardProductTokens = v
}

// GetClassification returns the Classification field value if set, zero value otherwise.
func (o *ProductResponse) GetClassification() ProductClassification {
	if o == nil || IsNil(o.Classification) {
		var ret ProductClassification
		return ret
	}
	return *o.Classification
}

// GetClassificationOk returns a tuple with the Classification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductResponse) GetClassificationOk() (*ProductClassification, bool) {
	if o == nil || IsNil(o.Classification) {
		return nil, false
	}
	return o.Classification, true
}

// HasClassification returns a boolean if a field has been set.
func (o *ProductResponse) HasClassification() bool {
	if o != nil && !IsNil(o.Classification) {
		return true
	}

	return false
}

// SetClassification gets a reference to the given ProductClassification and assigns it to the Classification field.
func (o *ProductResponse) SetClassification(v ProductClassification) {
	o.Classification = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *ProductResponse) GetConfig() ProductConfig {
	if o == nil || IsNil(o.Config) {
		var ret ProductConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductResponse) GetConfigOk() (*ProductConfig, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *ProductResponse) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given ProductConfig and assigns it to the Config field.
func (o *ProductResponse) SetConfig(v ProductConfig) {
	o.Config = &v
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *ProductResponse) GetCreatedTime() time.Time {
	if o == nil || IsNil(o.CreatedTime) {
		var ret time.Time
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductResponse) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedTime) {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *ProductResponse) HasCreatedTime() bool {
	if o != nil && !IsNil(o.CreatedTime) {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given time.Time and assigns it to the CreatedTime field.
func (o *ProductResponse) SetCreatedTime(v time.Time) {
	o.CreatedTime = &v
}

// GetCreditLine returns the CreditLine field value if set, zero value otherwise.
func (o *ProductResponse) GetCreditLine() ProductCreditLine {
	if o == nil || IsNil(o.CreditLine) {
		var ret ProductCreditLine
		return ret
	}
	return *o.CreditLine
}

// GetCreditLineOk returns a tuple with the CreditLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductResponse) GetCreditLineOk() (*ProductCreditLine, bool) {
	if o == nil || IsNil(o.CreditLine) {
		return nil, false
	}
	return o.CreditLine, true
}

// HasCreditLine returns a boolean if a field has been set.
func (o *ProductResponse) HasCreditLine() bool {
	if o != nil && !IsNil(o.CreditLine) {
		return true
	}

	return false
}

// SetCreditLine gets a reference to the given ProductCreditLine and assigns it to the CreditLine field.
func (o *ProductResponse) SetCreditLine(v ProductCreditLine) {
	o.CreditLine = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *ProductResponse) GetCurrencyCode() CurrencyCode {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret CurrencyCode
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductResponse) GetCurrencyCodeOk() (*CurrencyCode, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *ProductResponse) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given CurrencyCode and assigns it to the CurrencyCode field.
func (o *ProductResponse) SetCurrencyCode(v CurrencyCode) {
	o.CurrencyCode = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ProductResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ProductResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ProductResponse) SetDescription(v string) {
	o.Description = &v
}

// GetInterestCalculation returns the InterestCalculation field value if set, zero value otherwise.
func (o *ProductResponse) GetInterestCalculation() InterestCalculation {
	if o == nil || IsNil(o.InterestCalculation) {
		var ret InterestCalculation
		return ret
	}
	return *o.InterestCalculation
}

// GetInterestCalculationOk returns a tuple with the InterestCalculation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductResponse) GetInterestCalculationOk() (*InterestCalculation, bool) {
	if o == nil || IsNil(o.InterestCalculation) {
		return nil, false
	}
	return o.InterestCalculation, true
}

// HasInterestCalculation returns a boolean if a field has been set.
func (o *ProductResponse) HasInterestCalculation() bool {
	if o != nil && !IsNil(o.InterestCalculation) {
		return true
	}

	return false
}

// SetInterestCalculation gets a reference to the given InterestCalculation and assigns it to the InterestCalculation field.
func (o *ProductResponse) SetInterestCalculation(v InterestCalculation) {
	o.InterestCalculation = &v
}

// GetMinPaymentCalculation returns the MinPaymentCalculation field value if set, zero value otherwise.
func (o *ProductResponse) GetMinPaymentCalculation() ProductMinPaymentCalculation {
	if o == nil || IsNil(o.MinPaymentCalculation) {
		var ret ProductMinPaymentCalculation
		return ret
	}
	return *o.MinPaymentCalculation
}

// GetMinPaymentCalculationOk returns a tuple with the MinPaymentCalculation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductResponse) GetMinPaymentCalculationOk() (*ProductMinPaymentCalculation, bool) {
	if o == nil || IsNil(o.MinPaymentCalculation) {
		return nil, false
	}
	return o.MinPaymentCalculation, true
}

// HasMinPaymentCalculation returns a boolean if a field has been set.
func (o *ProductResponse) HasMinPaymentCalculation() bool {
	if o != nil && !IsNil(o.MinPaymentCalculation) {
		return true
	}

	return false
}

// SetMinPaymentCalculation gets a reference to the given ProductMinPaymentCalculation and assigns it to the MinPaymentCalculation field.
func (o *ProductResponse) SetMinPaymentCalculation(v ProductMinPaymentCalculation) {
	o.MinPaymentCalculation = &v
}

// GetMinPaymentFlatAmount returns the MinPaymentFlatAmount field value if set, zero value otherwise.
func (o *ProductResponse) GetMinPaymentFlatAmount() decimal.Decimal {
	if o == nil || IsNil(o.MinPaymentFlatAmount) {
		var ret decimal.Decimal
		return ret
	}
	return *o.MinPaymentFlatAmount
}

// GetMinPaymentFlatAmountOk returns a tuple with the MinPaymentFlatAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductResponse) GetMinPaymentFlatAmountOk() (*decimal.Decimal, bool) {
	if o == nil || IsNil(o.MinPaymentFlatAmount) {
		return nil, false
	}
	return o.MinPaymentFlatAmount, true
}

// HasMinPaymentFlatAmount returns a boolean if a field has been set.
func (o *ProductResponse) HasMinPaymentFlatAmount() bool {
	if o != nil && !IsNil(o.MinPaymentFlatAmount) {
		return true
	}

	return false
}

// SetMinPaymentFlatAmount gets a reference to the given decimal.Decimal and assigns it to the MinPaymentFlatAmount field.
func (o *ProductResponse) SetMinPaymentFlatAmount(v decimal.Decimal) {
	o.MinPaymentFlatAmount = &v
}

// GetMinPaymentPercentage returns the MinPaymentPercentage field value if set, zero value otherwise.
func (o *ProductResponse) GetMinPaymentPercentage() decimal.Decimal {
	if o == nil || IsNil(o.MinPaymentPercentage) {
		var ret decimal.Decimal
		return ret
	}
	return *o.MinPaymentPercentage
}

// GetMinPaymentPercentageOk returns a tuple with the MinPaymentPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductResponse) GetMinPaymentPercentageOk() (*decimal.Decimal, bool) {
	if o == nil || IsNil(o.MinPaymentPercentage) {
		return nil, false
	}
	return o.MinPaymentPercentage, true
}

// HasMinPaymentPercentage returns a boolean if a field has been set.
func (o *ProductResponse) HasMinPaymentPercentage() bool {
	if o != nil && !IsNil(o.MinPaymentPercentage) {
		return true
	}

	return false
}

// SetMinPaymentPercentage gets a reference to the given decimal.Decimal and assigns it to the MinPaymentPercentage field.
func (o *ProductResponse) SetMinPaymentPercentage(v decimal.Decimal) {
	o.MinPaymentPercentage = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProductResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProductResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProductResponse) SetName(v string) {
	o.Name = &v
}

// GetParentProductToken returns the ParentProductToken field value if set, zero value otherwise.
func (o *ProductResponse) GetParentProductToken() string {
	if o == nil || IsNil(o.ParentProductToken) {
		var ret string
		return ret
	}
	return *o.ParentProductToken
}

// GetParentProductTokenOk returns a tuple with the ParentProductToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductResponse) GetParentProductTokenOk() (*string, bool) {
	if o == nil || IsNil(o.ParentProductToken) {
		return nil, false
	}
	return o.ParentProductToken, true
}

// HasParentProductToken returns a boolean if a field has been set.
func (o *ProductResponse) HasParentProductToken() bool {
	if o != nil && !IsNil(o.ParentProductToken) {
		return true
	}

	return false
}

// SetParentProductToken gets a reference to the given string and assigns it to the ParentProductToken field.
func (o *ProductResponse) SetParentProductToken(v string) {
	o.ParentProductToken = &v
}

// GetPaymentAllocationOrder returns the PaymentAllocationOrder field value if set, zero value otherwise.
func (o *ProductResponse) GetPaymentAllocationOrder() []PaymentAllocationOrderEnum {
	if o == nil || IsNil(o.PaymentAllocationOrder) {
		var ret []PaymentAllocationOrderEnum
		return ret
	}
	return o.PaymentAllocationOrder
}

// GetPaymentAllocationOrderOk returns a tuple with the PaymentAllocationOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductResponse) GetPaymentAllocationOrderOk() ([]PaymentAllocationOrderEnum, bool) {
	if o == nil || IsNil(o.PaymentAllocationOrder) {
		return nil, false
	}
	return o.PaymentAllocationOrder, true
}

// HasPaymentAllocationOrder returns a boolean if a field has been set.
func (o *ProductResponse) HasPaymentAllocationOrder() bool {
	if o != nil && !IsNil(o.PaymentAllocationOrder) {
		return true
	}

	return false
}

// SetPaymentAllocationOrder gets a reference to the given []PaymentAllocationOrderEnum and assigns it to the PaymentAllocationOrder field.
func (o *ProductResponse) SetPaymentAllocationOrder(v []PaymentAllocationOrderEnum) {
	o.PaymentAllocationOrder = v
}

// GetProductSubType returns the ProductSubType field value if set, zero value otherwise.
func (o *ProductResponse) GetProductSubType() ProductSubType {
	if o == nil || IsNil(o.ProductSubType) {
		var ret ProductSubType
		return ret
	}
	return *o.ProductSubType
}

// GetProductSubTypeOk returns a tuple with the ProductSubType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductResponse) GetProductSubTypeOk() (*ProductSubType, bool) {
	if o == nil || IsNil(o.ProductSubType) {
		return nil, false
	}
	return o.ProductSubType, true
}

// HasProductSubType returns a boolean if a field has been set.
func (o *ProductResponse) HasProductSubType() bool {
	if o != nil && !IsNil(o.ProductSubType) {
		return true
	}

	return false
}

// SetProductSubType gets a reference to the given ProductSubType and assigns it to the ProductSubType field.
func (o *ProductResponse) SetProductSubType(v ProductSubType) {
	o.ProductSubType = &v
}

// GetProductType returns the ProductType field value if set, zero value otherwise.
func (o *ProductResponse) GetProductType() ProductType {
	if o == nil || IsNil(o.ProductType) {
		var ret ProductType
		return ret
	}
	return *o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductResponse) GetProductTypeOk() (*ProductType, bool) {
	if o == nil || IsNil(o.ProductType) {
		return nil, false
	}
	return o.ProductType, true
}

// HasProductType returns a boolean if a field has been set.
func (o *ProductResponse) HasProductType() bool {
	if o != nil && !IsNil(o.ProductType) {
		return true
	}

	return false
}

// SetProductType gets a reference to the given ProductType and assigns it to the ProductType field.
func (o *ProductResponse) SetProductType(v ProductType) {
	o.ProductType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ProductResponse) GetStatus() ResourceStatus {
	if o == nil || IsNil(o.Status) {
		var ret ResourceStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductResponse) GetStatusOk() (*ResourceStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ProductResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ResourceStatus and assigns it to the Status field.
func (o *ProductResponse) SetStatus(v ResourceStatus) {
	o.Status = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *ProductResponse) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductResponse) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *ProductResponse) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *ProductResponse) SetToken(v string) {
	o.Token = &v
}

// GetUpdatedTime returns the UpdatedTime field value if set, zero value otherwise.
func (o *ProductResponse) GetUpdatedTime() time.Time {
	if o == nil || IsNil(o.UpdatedTime) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedTime
}

// GetUpdatedTimeOk returns a tuple with the UpdatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductResponse) GetUpdatedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedTime) {
		return nil, false
	}
	return o.UpdatedTime, true
}

// HasUpdatedTime returns a boolean if a field has been set.
func (o *ProductResponse) HasUpdatedTime() bool {
	if o != nil && !IsNil(o.UpdatedTime) {
		return true
	}

	return false
}

// SetUpdatedTime gets a reference to the given time.Time and assigns it to the UpdatedTime field.
func (o *ProductResponse) SetUpdatedTime(v time.Time) {
	o.UpdatedTime = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *ProductResponse) GetUsage() []BalanceType {
	if o == nil || IsNil(o.Usage) {
		var ret []BalanceType
		return ret
	}
	return o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductResponse) GetUsageOk() ([]BalanceType, bool) {
	if o == nil || IsNil(o.Usage) {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *ProductResponse) HasUsage() bool {
	if o != nil && !IsNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given []BalanceType and assigns it to the Usage field.
func (o *ProductResponse) SetUsage(v []BalanceType) {
	o.Usage = v
}

func (o ProductResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CardProductTokens) {
		toSerialize["card_product_tokens"] = o.CardProductTokens
	}
	if !IsNil(o.Classification) {
		toSerialize["classification"] = o.Classification
	}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.CreatedTime) {
		toSerialize["created_time"] = o.CreatedTime
	}
	if !IsNil(o.CreditLine) {
		toSerialize["credit_line"] = o.CreditLine
	}
	if !IsNil(o.CurrencyCode) {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.InterestCalculation) {
		toSerialize["interest_calculation"] = o.InterestCalculation
	}
	if !IsNil(o.MinPaymentCalculation) {
		toSerialize["min_payment_calculation"] = o.MinPaymentCalculation
	}
	if !IsNil(o.MinPaymentFlatAmount) {
		toSerialize["min_payment_flat_amount"] = o.MinPaymentFlatAmount
	}
	if !IsNil(o.MinPaymentPercentage) {
		toSerialize["min_payment_percentage"] = o.MinPaymentPercentage
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ParentProductToken) {
		toSerialize["parent_product_token"] = o.ParentProductToken
	}
	if !IsNil(o.PaymentAllocationOrder) {
		toSerialize["payment_allocation_order"] = o.PaymentAllocationOrder
	}
	if !IsNil(o.ProductSubType) {
		toSerialize["product_sub_type"] = o.ProductSubType
	}
	if !IsNil(o.ProductType) {
		toSerialize["product_type"] = o.ProductType
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.UpdatedTime) {
		toSerialize["updated_time"] = o.UpdatedTime
	}
	if !IsNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}
	return toSerialize, nil
}

type NullableProductResponse struct {
	value *ProductResponse
	isSet bool
}

func (v NullableProductResponse) Get() *ProductResponse {
	return v.value
}

func (v *NullableProductResponse) Set(val *ProductResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProductResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProductResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductResponse(val *ProductResponse) *NullableProductResponse {
	return &NullableProductResponse{value: val, isSet: true}
}

func (v NullableProductResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


