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

// checks if the TransactionControls type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionControls{}

// TransactionControls Controls transactional characteristics of card usage.
type TransactionControls struct {
	// Set to `accept_us_only` to allow transactions only within the US.  Set to `decline_ofac_countries` to allow international transactions except with countries that the Financial Action Task Force (FATF) and Office of Foreign Assets Control (OFAC) have identified as high risk.  Users with the Admin role can create and update additional lists of accepted countries for transactions at the `/acceptedcountries` endpoint. See <</core-api/accepted-countries, Accepted Countries>>.
	AcceptedCountriesToken *string `json:"accepted_countries_token,omitempty"`
	AddressVerification *AvsControls `json:"address_verification,omitempty"`
	// Indicates whether to allow transactions where a Europay Mastercard and Visa (EMV) chip-enabled card was processed using the magstripe as fallback.
	AllowChipFallback *bool `json:"allow_chip_fallback,omitempty"`
	// *WARNING:* This field is deprecated and will be unsupported in a future release.  Allows cardholders to define a personal identification number (PIN) as they complete their first PIN-debit transaction.
	AllowFirstPinSetViaFinancialTransaction *bool `json:"allow_first_pin_set_via_financial_transaction,omitempty"`
	// If set to `true`, transactions can be authorized using GPA funds.  *NOTE:* For most programs, this field should be set to `true`.
	AllowGpaAuth *bool `json:"allow_gpa_auth,omitempty"`
	// The <</core-api/mcc-groups, MCC group>> `authorization_controls` object allows you to automatically increase authorization holds and to specify authorization expiration times based on merchant type. By default, these settings apply to all cards in your program. You can, however, exempt cards associated with a particular card product by setting this field to `false`.  *NOTE:* Partial authorizations are disallowed if this field is set to `true`.
	AllowMccGroupAuthorizationControls *bool `json:"allow_mcc_group_authorization_controls,omitempty"`
	// Indicates whether card network loads are allowed. The associated card's state must be `ACTIVE` or the load will be rejected.
	AllowNetworkLoad *bool `json:"allow_network_load,omitempty"`
	// Indicates whether card network loads are allowed. Sets the associated card's state to `ACTIVE` if its current state is `INACTIVE`.
	AllowNetworkLoadCardActivation *bool `json:"allow_network_load_card_activation,omitempty"`
	// Indicates whether quasi-cash transactions are allowed. In a quasi-cash transaction, the cardholder purchases an item that can be directly converted to cash, such as traveler's checks, money orders, casino chips, or lottery tickets.
	AllowQuasiCash *bool `json:"allow_quasi_cash,omitempty"`
	// If set to `true`, cards of this card product type require an Integrated Circuit Card.
	AlwaysRequireIcc *bool `json:"always_require_icc,omitempty"`
	// If set to `true`, cards of this card product type require a personal identification number (PIN).
	AlwaysRequirePin *bool `json:"always_require_pin,omitempty"`
	EnableCreditService *bool `json:"enable_credit_service,omitempty"`
	// Set to `true` to enable partial authorizations.  When this setting is `false` and the requested authorization amount exceeds available funds, the transaction is declined. When this setting is `true` and the requested authorization amount exceeds available funds, the transaction is authorized for the amount of available funds.
	EnablePartialAuthApproval *bool `json:"enable_partial_auth_approval,omitempty"`
	// Allows transactions to be approved even if the card's `state = SUSPENDED`. When this field is set to `true`, the card behaves as if its `state = ACTIVE`.
	IgnoreCardSuspendedState *bool `json:"ignore_card_suspended_state,omitempty"`
	// Specifies the language for 3D Secure and digital wallet token notifications sent to cardholders under this card program.  You can send notifications to your cardholders in the following languages:  * *ces* – Czech * *deu* – German * *eng* – English * *fra* – French * *grc* – Greek * *ita* – Italian * *nld* – Dutch * *pol* – Polish * *prt* – Portuguese * *rou* – Romanian * *spa* – Spanish * *swe* – Swedish  By default, notifications are sent in English.  To specify the language for OTP notifications at the user level, see <</core-api/users, Users>>. Languages set at the user level take precedence over the language set at the card product level.
	NotificationLanguage *string `json:"notification_language,omitempty"`
	// The token of the merchant group that you want to exempt from quasi-cash transaction authorization control, allowing your cardholders to conduct quasi-cash transactions. In a quasi-cash transaction, the cardholder purchases an item that can be directly converted to cash, such as traveler's checks, money orders, casino chips, or lottery tickets.  You can specify a merchant group token in addition to whatever merchant identifiers you listed in the `quasi_cash_exempt_mids` field, if any. For more information, see <</core-api/merchant-groups, Merchant Groups>>.
	QuasiCashExemptMerchantGroupToken *string `json:"quasi_cash_exempt_merchant_group_token,omitempty"`
	// Comma-separated list of merchant identifiers that you want to exempt from quasi-cash transaction authorization control, allowing your cardholders to conduct quasi-cash transactions. In a quasi-cash transaction, the cardholder purchases an item that can be directly converted to cash, such as traveler's checks, money orders, casino chips, or lottery tickets.
	QuasiCashExemptMids *string `json:"quasi_cash_exempt_mids,omitempty"`
	// A value of `true` indicates that if `card_presence_required` is `true`, the card's security code is required.
	RequireCardNotPresentCardSecurityCode *bool `json:"require_card_not_present_card_security_code,omitempty"`
	StrongCustomerAuthenticationLimits *StrongCustomerAuthenticationLimits `json:"strong_customer_authentication_limits,omitempty"`
}

// NewTransactionControls instantiates a new TransactionControls object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionControls() *TransactionControls {
	this := TransactionControls{}
	var allowChipFallback bool = false
	this.AllowChipFallback = &allowChipFallback
	var allowFirstPinSetViaFinancialTransaction bool = false
	this.AllowFirstPinSetViaFinancialTransaction = &allowFirstPinSetViaFinancialTransaction
	var allowGpaAuth bool = false
	this.AllowGpaAuth = &allowGpaAuth
	var allowMccGroupAuthorizationControls bool = false
	this.AllowMccGroupAuthorizationControls = &allowMccGroupAuthorizationControls
	var allowNetworkLoad bool = false
	this.AllowNetworkLoad = &allowNetworkLoad
	var allowNetworkLoadCardActivation bool = false
	this.AllowNetworkLoadCardActivation = &allowNetworkLoadCardActivation
	var allowQuasiCash bool = false
	this.AllowQuasiCash = &allowQuasiCash
	var alwaysRequireIcc bool = false
	this.AlwaysRequireIcc = &alwaysRequireIcc
	var alwaysRequirePin bool = false
	this.AlwaysRequirePin = &alwaysRequirePin
	var enableCreditService bool = false
	this.EnableCreditService = &enableCreditService
	var enablePartialAuthApproval bool = false
	this.EnablePartialAuthApproval = &enablePartialAuthApproval
	var ignoreCardSuspendedState bool = false
	this.IgnoreCardSuspendedState = &ignoreCardSuspendedState
	var requireCardNotPresentCardSecurityCode bool = false
	this.RequireCardNotPresentCardSecurityCode = &requireCardNotPresentCardSecurityCode
	return &this
}

// NewTransactionControlsWithDefaults instantiates a new TransactionControls object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionControlsWithDefaults() *TransactionControls {
	this := TransactionControls{}
	var allowChipFallback bool = false
	this.AllowChipFallback = &allowChipFallback
	var allowFirstPinSetViaFinancialTransaction bool = false
	this.AllowFirstPinSetViaFinancialTransaction = &allowFirstPinSetViaFinancialTransaction
	var allowGpaAuth bool = false
	this.AllowGpaAuth = &allowGpaAuth
	var allowMccGroupAuthorizationControls bool = false
	this.AllowMccGroupAuthorizationControls = &allowMccGroupAuthorizationControls
	var allowNetworkLoad bool = false
	this.AllowNetworkLoad = &allowNetworkLoad
	var allowNetworkLoadCardActivation bool = false
	this.AllowNetworkLoadCardActivation = &allowNetworkLoadCardActivation
	var allowQuasiCash bool = false
	this.AllowQuasiCash = &allowQuasiCash
	var alwaysRequireIcc bool = false
	this.AlwaysRequireIcc = &alwaysRequireIcc
	var alwaysRequirePin bool = false
	this.AlwaysRequirePin = &alwaysRequirePin
	var enableCreditService bool = false
	this.EnableCreditService = &enableCreditService
	var enablePartialAuthApproval bool = false
	this.EnablePartialAuthApproval = &enablePartialAuthApproval
	var ignoreCardSuspendedState bool = false
	this.IgnoreCardSuspendedState = &ignoreCardSuspendedState
	var requireCardNotPresentCardSecurityCode bool = false
	this.RequireCardNotPresentCardSecurityCode = &requireCardNotPresentCardSecurityCode
	return &this
}

// GetAcceptedCountriesToken returns the AcceptedCountriesToken field value if set, zero value otherwise.
func (o *TransactionControls) GetAcceptedCountriesToken() string {
	if o == nil || IsNil(o.AcceptedCountriesToken) {
		var ret string
		return ret
	}
	return *o.AcceptedCountriesToken
}

// GetAcceptedCountriesTokenOk returns a tuple with the AcceptedCountriesToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionControls) GetAcceptedCountriesTokenOk() (*string, bool) {
	if o == nil || IsNil(o.AcceptedCountriesToken) {
		return nil, false
	}
	return o.AcceptedCountriesToken, true
}

// HasAcceptedCountriesToken returns a boolean if a field has been set.
func (o *TransactionControls) HasAcceptedCountriesToken() bool {
	if o != nil && !IsNil(o.AcceptedCountriesToken) {
		return true
	}

	return false
}

// SetAcceptedCountriesToken gets a reference to the given string and assigns it to the AcceptedCountriesToken field.
func (o *TransactionControls) SetAcceptedCountriesToken(v string) {
	o.AcceptedCountriesToken = &v
}

// GetAddressVerification returns the AddressVerification field value if set, zero value otherwise.
func (o *TransactionControls) GetAddressVerification() AvsControls {
	if o == nil || IsNil(o.AddressVerification) {
		var ret AvsControls
		return ret
	}
	return *o.AddressVerification
}

// GetAddressVerificationOk returns a tuple with the AddressVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionControls) GetAddressVerificationOk() (*AvsControls, bool) {
	if o == nil || IsNil(o.AddressVerification) {
		return nil, false
	}
	return o.AddressVerification, true
}

// HasAddressVerification returns a boolean if a field has been set.
func (o *TransactionControls) HasAddressVerification() bool {
	if o != nil && !IsNil(o.AddressVerification) {
		return true
	}

	return false
}

// SetAddressVerification gets a reference to the given AvsControls and assigns it to the AddressVerification field.
func (o *TransactionControls) SetAddressVerification(v AvsControls) {
	o.AddressVerification = &v
}

// GetAllowChipFallback returns the AllowChipFallback field value if set, zero value otherwise.
func (o *TransactionControls) GetAllowChipFallback() bool {
	if o == nil || IsNil(o.AllowChipFallback) {
		var ret bool
		return ret
	}
	return *o.AllowChipFallback
}

// GetAllowChipFallbackOk returns a tuple with the AllowChipFallback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionControls) GetAllowChipFallbackOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowChipFallback) {
		return nil, false
	}
	return o.AllowChipFallback, true
}

// HasAllowChipFallback returns a boolean if a field has been set.
func (o *TransactionControls) HasAllowChipFallback() bool {
	if o != nil && !IsNil(o.AllowChipFallback) {
		return true
	}

	return false
}

// SetAllowChipFallback gets a reference to the given bool and assigns it to the AllowChipFallback field.
func (o *TransactionControls) SetAllowChipFallback(v bool) {
	o.AllowChipFallback = &v
}

// GetAllowFirstPinSetViaFinancialTransaction returns the AllowFirstPinSetViaFinancialTransaction field value if set, zero value otherwise.
func (o *TransactionControls) GetAllowFirstPinSetViaFinancialTransaction() bool {
	if o == nil || IsNil(o.AllowFirstPinSetViaFinancialTransaction) {
		var ret bool
		return ret
	}
	return *o.AllowFirstPinSetViaFinancialTransaction
}

// GetAllowFirstPinSetViaFinancialTransactionOk returns a tuple with the AllowFirstPinSetViaFinancialTransaction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionControls) GetAllowFirstPinSetViaFinancialTransactionOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowFirstPinSetViaFinancialTransaction) {
		return nil, false
	}
	return o.AllowFirstPinSetViaFinancialTransaction, true
}

// HasAllowFirstPinSetViaFinancialTransaction returns a boolean if a field has been set.
func (o *TransactionControls) HasAllowFirstPinSetViaFinancialTransaction() bool {
	if o != nil && !IsNil(o.AllowFirstPinSetViaFinancialTransaction) {
		return true
	}

	return false
}

// SetAllowFirstPinSetViaFinancialTransaction gets a reference to the given bool and assigns it to the AllowFirstPinSetViaFinancialTransaction field.
func (o *TransactionControls) SetAllowFirstPinSetViaFinancialTransaction(v bool) {
	o.AllowFirstPinSetViaFinancialTransaction = &v
}

// GetAllowGpaAuth returns the AllowGpaAuth field value if set, zero value otherwise.
func (o *TransactionControls) GetAllowGpaAuth() bool {
	if o == nil || IsNil(o.AllowGpaAuth) {
		var ret bool
		return ret
	}
	return *o.AllowGpaAuth
}

// GetAllowGpaAuthOk returns a tuple with the AllowGpaAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionControls) GetAllowGpaAuthOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowGpaAuth) {
		return nil, false
	}
	return o.AllowGpaAuth, true
}

// HasAllowGpaAuth returns a boolean if a field has been set.
func (o *TransactionControls) HasAllowGpaAuth() bool {
	if o != nil && !IsNil(o.AllowGpaAuth) {
		return true
	}

	return false
}

// SetAllowGpaAuth gets a reference to the given bool and assigns it to the AllowGpaAuth field.
func (o *TransactionControls) SetAllowGpaAuth(v bool) {
	o.AllowGpaAuth = &v
}

// GetAllowMccGroupAuthorizationControls returns the AllowMccGroupAuthorizationControls field value if set, zero value otherwise.
func (o *TransactionControls) GetAllowMccGroupAuthorizationControls() bool {
	if o == nil || IsNil(o.AllowMccGroupAuthorizationControls) {
		var ret bool
		return ret
	}
	return *o.AllowMccGroupAuthorizationControls
}

// GetAllowMccGroupAuthorizationControlsOk returns a tuple with the AllowMccGroupAuthorizationControls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionControls) GetAllowMccGroupAuthorizationControlsOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowMccGroupAuthorizationControls) {
		return nil, false
	}
	return o.AllowMccGroupAuthorizationControls, true
}

// HasAllowMccGroupAuthorizationControls returns a boolean if a field has been set.
func (o *TransactionControls) HasAllowMccGroupAuthorizationControls() bool {
	if o != nil && !IsNil(o.AllowMccGroupAuthorizationControls) {
		return true
	}

	return false
}

// SetAllowMccGroupAuthorizationControls gets a reference to the given bool and assigns it to the AllowMccGroupAuthorizationControls field.
func (o *TransactionControls) SetAllowMccGroupAuthorizationControls(v bool) {
	o.AllowMccGroupAuthorizationControls = &v
}

// GetAllowNetworkLoad returns the AllowNetworkLoad field value if set, zero value otherwise.
func (o *TransactionControls) GetAllowNetworkLoad() bool {
	if o == nil || IsNil(o.AllowNetworkLoad) {
		var ret bool
		return ret
	}
	return *o.AllowNetworkLoad
}

// GetAllowNetworkLoadOk returns a tuple with the AllowNetworkLoad field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionControls) GetAllowNetworkLoadOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowNetworkLoad) {
		return nil, false
	}
	return o.AllowNetworkLoad, true
}

// HasAllowNetworkLoad returns a boolean if a field has been set.
func (o *TransactionControls) HasAllowNetworkLoad() bool {
	if o != nil && !IsNil(o.AllowNetworkLoad) {
		return true
	}

	return false
}

// SetAllowNetworkLoad gets a reference to the given bool and assigns it to the AllowNetworkLoad field.
func (o *TransactionControls) SetAllowNetworkLoad(v bool) {
	o.AllowNetworkLoad = &v
}

// GetAllowNetworkLoadCardActivation returns the AllowNetworkLoadCardActivation field value if set, zero value otherwise.
func (o *TransactionControls) GetAllowNetworkLoadCardActivation() bool {
	if o == nil || IsNil(o.AllowNetworkLoadCardActivation) {
		var ret bool
		return ret
	}
	return *o.AllowNetworkLoadCardActivation
}

// GetAllowNetworkLoadCardActivationOk returns a tuple with the AllowNetworkLoadCardActivation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionControls) GetAllowNetworkLoadCardActivationOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowNetworkLoadCardActivation) {
		return nil, false
	}
	return o.AllowNetworkLoadCardActivation, true
}

// HasAllowNetworkLoadCardActivation returns a boolean if a field has been set.
func (o *TransactionControls) HasAllowNetworkLoadCardActivation() bool {
	if o != nil && !IsNil(o.AllowNetworkLoadCardActivation) {
		return true
	}

	return false
}

// SetAllowNetworkLoadCardActivation gets a reference to the given bool and assigns it to the AllowNetworkLoadCardActivation field.
func (o *TransactionControls) SetAllowNetworkLoadCardActivation(v bool) {
	o.AllowNetworkLoadCardActivation = &v
}

// GetAllowQuasiCash returns the AllowQuasiCash field value if set, zero value otherwise.
func (o *TransactionControls) GetAllowQuasiCash() bool {
	if o == nil || IsNil(o.AllowQuasiCash) {
		var ret bool
		return ret
	}
	return *o.AllowQuasiCash
}

// GetAllowQuasiCashOk returns a tuple with the AllowQuasiCash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionControls) GetAllowQuasiCashOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowQuasiCash) {
		return nil, false
	}
	return o.AllowQuasiCash, true
}

// HasAllowQuasiCash returns a boolean if a field has been set.
func (o *TransactionControls) HasAllowQuasiCash() bool {
	if o != nil && !IsNil(o.AllowQuasiCash) {
		return true
	}

	return false
}

// SetAllowQuasiCash gets a reference to the given bool and assigns it to the AllowQuasiCash field.
func (o *TransactionControls) SetAllowQuasiCash(v bool) {
	o.AllowQuasiCash = &v
}

// GetAlwaysRequireIcc returns the AlwaysRequireIcc field value if set, zero value otherwise.
func (o *TransactionControls) GetAlwaysRequireIcc() bool {
	if o == nil || IsNil(o.AlwaysRequireIcc) {
		var ret bool
		return ret
	}
	return *o.AlwaysRequireIcc
}

// GetAlwaysRequireIccOk returns a tuple with the AlwaysRequireIcc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionControls) GetAlwaysRequireIccOk() (*bool, bool) {
	if o == nil || IsNil(o.AlwaysRequireIcc) {
		return nil, false
	}
	return o.AlwaysRequireIcc, true
}

// HasAlwaysRequireIcc returns a boolean if a field has been set.
func (o *TransactionControls) HasAlwaysRequireIcc() bool {
	if o != nil && !IsNil(o.AlwaysRequireIcc) {
		return true
	}

	return false
}

// SetAlwaysRequireIcc gets a reference to the given bool and assigns it to the AlwaysRequireIcc field.
func (o *TransactionControls) SetAlwaysRequireIcc(v bool) {
	o.AlwaysRequireIcc = &v
}

// GetAlwaysRequirePin returns the AlwaysRequirePin field value if set, zero value otherwise.
func (o *TransactionControls) GetAlwaysRequirePin() bool {
	if o == nil || IsNil(o.AlwaysRequirePin) {
		var ret bool
		return ret
	}
	return *o.AlwaysRequirePin
}

// GetAlwaysRequirePinOk returns a tuple with the AlwaysRequirePin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionControls) GetAlwaysRequirePinOk() (*bool, bool) {
	if o == nil || IsNil(o.AlwaysRequirePin) {
		return nil, false
	}
	return o.AlwaysRequirePin, true
}

// HasAlwaysRequirePin returns a boolean if a field has been set.
func (o *TransactionControls) HasAlwaysRequirePin() bool {
	if o != nil && !IsNil(o.AlwaysRequirePin) {
		return true
	}

	return false
}

// SetAlwaysRequirePin gets a reference to the given bool and assigns it to the AlwaysRequirePin field.
func (o *TransactionControls) SetAlwaysRequirePin(v bool) {
	o.AlwaysRequirePin = &v
}

// GetEnableCreditService returns the EnableCreditService field value if set, zero value otherwise.
func (o *TransactionControls) GetEnableCreditService() bool {
	if o == nil || IsNil(o.EnableCreditService) {
		var ret bool
		return ret
	}
	return *o.EnableCreditService
}

// GetEnableCreditServiceOk returns a tuple with the EnableCreditService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionControls) GetEnableCreditServiceOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableCreditService) {
		return nil, false
	}
	return o.EnableCreditService, true
}

// HasEnableCreditService returns a boolean if a field has been set.
func (o *TransactionControls) HasEnableCreditService() bool {
	if o != nil && !IsNil(o.EnableCreditService) {
		return true
	}

	return false
}

// SetEnableCreditService gets a reference to the given bool and assigns it to the EnableCreditService field.
func (o *TransactionControls) SetEnableCreditService(v bool) {
	o.EnableCreditService = &v
}

// GetEnablePartialAuthApproval returns the EnablePartialAuthApproval field value if set, zero value otherwise.
func (o *TransactionControls) GetEnablePartialAuthApproval() bool {
	if o == nil || IsNil(o.EnablePartialAuthApproval) {
		var ret bool
		return ret
	}
	return *o.EnablePartialAuthApproval
}

// GetEnablePartialAuthApprovalOk returns a tuple with the EnablePartialAuthApproval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionControls) GetEnablePartialAuthApprovalOk() (*bool, bool) {
	if o == nil || IsNil(o.EnablePartialAuthApproval) {
		return nil, false
	}
	return o.EnablePartialAuthApproval, true
}

// HasEnablePartialAuthApproval returns a boolean if a field has been set.
func (o *TransactionControls) HasEnablePartialAuthApproval() bool {
	if o != nil && !IsNil(o.EnablePartialAuthApproval) {
		return true
	}

	return false
}

// SetEnablePartialAuthApproval gets a reference to the given bool and assigns it to the EnablePartialAuthApproval field.
func (o *TransactionControls) SetEnablePartialAuthApproval(v bool) {
	o.EnablePartialAuthApproval = &v
}

// GetIgnoreCardSuspendedState returns the IgnoreCardSuspendedState field value if set, zero value otherwise.
func (o *TransactionControls) GetIgnoreCardSuspendedState() bool {
	if o == nil || IsNil(o.IgnoreCardSuspendedState) {
		var ret bool
		return ret
	}
	return *o.IgnoreCardSuspendedState
}

// GetIgnoreCardSuspendedStateOk returns a tuple with the IgnoreCardSuspendedState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionControls) GetIgnoreCardSuspendedStateOk() (*bool, bool) {
	if o == nil || IsNil(o.IgnoreCardSuspendedState) {
		return nil, false
	}
	return o.IgnoreCardSuspendedState, true
}

// HasIgnoreCardSuspendedState returns a boolean if a field has been set.
func (o *TransactionControls) HasIgnoreCardSuspendedState() bool {
	if o != nil && !IsNil(o.IgnoreCardSuspendedState) {
		return true
	}

	return false
}

// SetIgnoreCardSuspendedState gets a reference to the given bool and assigns it to the IgnoreCardSuspendedState field.
func (o *TransactionControls) SetIgnoreCardSuspendedState(v bool) {
	o.IgnoreCardSuspendedState = &v
}

// GetNotificationLanguage returns the NotificationLanguage field value if set, zero value otherwise.
func (o *TransactionControls) GetNotificationLanguage() string {
	if o == nil || IsNil(o.NotificationLanguage) {
		var ret string
		return ret
	}
	return *o.NotificationLanguage
}

// GetNotificationLanguageOk returns a tuple with the NotificationLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionControls) GetNotificationLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.NotificationLanguage) {
		return nil, false
	}
	return o.NotificationLanguage, true
}

// HasNotificationLanguage returns a boolean if a field has been set.
func (o *TransactionControls) HasNotificationLanguage() bool {
	if o != nil && !IsNil(o.NotificationLanguage) {
		return true
	}

	return false
}

// SetNotificationLanguage gets a reference to the given string and assigns it to the NotificationLanguage field.
func (o *TransactionControls) SetNotificationLanguage(v string) {
	o.NotificationLanguage = &v
}

// GetQuasiCashExemptMerchantGroupToken returns the QuasiCashExemptMerchantGroupToken field value if set, zero value otherwise.
func (o *TransactionControls) GetQuasiCashExemptMerchantGroupToken() string {
	if o == nil || IsNil(o.QuasiCashExemptMerchantGroupToken) {
		var ret string
		return ret
	}
	return *o.QuasiCashExemptMerchantGroupToken
}

// GetQuasiCashExemptMerchantGroupTokenOk returns a tuple with the QuasiCashExemptMerchantGroupToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionControls) GetQuasiCashExemptMerchantGroupTokenOk() (*string, bool) {
	if o == nil || IsNil(o.QuasiCashExemptMerchantGroupToken) {
		return nil, false
	}
	return o.QuasiCashExemptMerchantGroupToken, true
}

// HasQuasiCashExemptMerchantGroupToken returns a boolean if a field has been set.
func (o *TransactionControls) HasQuasiCashExemptMerchantGroupToken() bool {
	if o != nil && !IsNil(o.QuasiCashExemptMerchantGroupToken) {
		return true
	}

	return false
}

// SetQuasiCashExemptMerchantGroupToken gets a reference to the given string and assigns it to the QuasiCashExemptMerchantGroupToken field.
func (o *TransactionControls) SetQuasiCashExemptMerchantGroupToken(v string) {
	o.QuasiCashExemptMerchantGroupToken = &v
}

// GetQuasiCashExemptMids returns the QuasiCashExemptMids field value if set, zero value otherwise.
func (o *TransactionControls) GetQuasiCashExemptMids() string {
	if o == nil || IsNil(o.QuasiCashExemptMids) {
		var ret string
		return ret
	}
	return *o.QuasiCashExemptMids
}

// GetQuasiCashExemptMidsOk returns a tuple with the QuasiCashExemptMids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionControls) GetQuasiCashExemptMidsOk() (*string, bool) {
	if o == nil || IsNil(o.QuasiCashExemptMids) {
		return nil, false
	}
	return o.QuasiCashExemptMids, true
}

// HasQuasiCashExemptMids returns a boolean if a field has been set.
func (o *TransactionControls) HasQuasiCashExemptMids() bool {
	if o != nil && !IsNil(o.QuasiCashExemptMids) {
		return true
	}

	return false
}

// SetQuasiCashExemptMids gets a reference to the given string and assigns it to the QuasiCashExemptMids field.
func (o *TransactionControls) SetQuasiCashExemptMids(v string) {
	o.QuasiCashExemptMids = &v
}

// GetRequireCardNotPresentCardSecurityCode returns the RequireCardNotPresentCardSecurityCode field value if set, zero value otherwise.
func (o *TransactionControls) GetRequireCardNotPresentCardSecurityCode() bool {
	if o == nil || IsNil(o.RequireCardNotPresentCardSecurityCode) {
		var ret bool
		return ret
	}
	return *o.RequireCardNotPresentCardSecurityCode
}

// GetRequireCardNotPresentCardSecurityCodeOk returns a tuple with the RequireCardNotPresentCardSecurityCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionControls) GetRequireCardNotPresentCardSecurityCodeOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireCardNotPresentCardSecurityCode) {
		return nil, false
	}
	return o.RequireCardNotPresentCardSecurityCode, true
}

// HasRequireCardNotPresentCardSecurityCode returns a boolean if a field has been set.
func (o *TransactionControls) HasRequireCardNotPresentCardSecurityCode() bool {
	if o != nil && !IsNil(o.RequireCardNotPresentCardSecurityCode) {
		return true
	}

	return false
}

// SetRequireCardNotPresentCardSecurityCode gets a reference to the given bool and assigns it to the RequireCardNotPresentCardSecurityCode field.
func (o *TransactionControls) SetRequireCardNotPresentCardSecurityCode(v bool) {
	o.RequireCardNotPresentCardSecurityCode = &v
}

// GetStrongCustomerAuthenticationLimits returns the StrongCustomerAuthenticationLimits field value if set, zero value otherwise.
func (o *TransactionControls) GetStrongCustomerAuthenticationLimits() StrongCustomerAuthenticationLimits {
	if o == nil || IsNil(o.StrongCustomerAuthenticationLimits) {
		var ret StrongCustomerAuthenticationLimits
		return ret
	}
	return *o.StrongCustomerAuthenticationLimits
}

// GetStrongCustomerAuthenticationLimitsOk returns a tuple with the StrongCustomerAuthenticationLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionControls) GetStrongCustomerAuthenticationLimitsOk() (*StrongCustomerAuthenticationLimits, bool) {
	if o == nil || IsNil(o.StrongCustomerAuthenticationLimits) {
		return nil, false
	}
	return o.StrongCustomerAuthenticationLimits, true
}

// HasStrongCustomerAuthenticationLimits returns a boolean if a field has been set.
func (o *TransactionControls) HasStrongCustomerAuthenticationLimits() bool {
	if o != nil && !IsNil(o.StrongCustomerAuthenticationLimits) {
		return true
	}

	return false
}

// SetStrongCustomerAuthenticationLimits gets a reference to the given StrongCustomerAuthenticationLimits and assigns it to the StrongCustomerAuthenticationLimits field.
func (o *TransactionControls) SetStrongCustomerAuthenticationLimits(v StrongCustomerAuthenticationLimits) {
	o.StrongCustomerAuthenticationLimits = &v
}

func (o TransactionControls) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionControls) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AcceptedCountriesToken) {
		toSerialize["accepted_countries_token"] = o.AcceptedCountriesToken
	}
	if !IsNil(o.AddressVerification) {
		toSerialize["address_verification"] = o.AddressVerification
	}
	if !IsNil(o.AllowChipFallback) {
		toSerialize["allow_chip_fallback"] = o.AllowChipFallback
	}
	if !IsNil(o.AllowFirstPinSetViaFinancialTransaction) {
		toSerialize["allow_first_pin_set_via_financial_transaction"] = o.AllowFirstPinSetViaFinancialTransaction
	}
	if !IsNil(o.AllowGpaAuth) {
		toSerialize["allow_gpa_auth"] = o.AllowGpaAuth
	}
	if !IsNil(o.AllowMccGroupAuthorizationControls) {
		toSerialize["allow_mcc_group_authorization_controls"] = o.AllowMccGroupAuthorizationControls
	}
	if !IsNil(o.AllowNetworkLoad) {
		toSerialize["allow_network_load"] = o.AllowNetworkLoad
	}
	if !IsNil(o.AllowNetworkLoadCardActivation) {
		toSerialize["allow_network_load_card_activation"] = o.AllowNetworkLoadCardActivation
	}
	if !IsNil(o.AllowQuasiCash) {
		toSerialize["allow_quasi_cash"] = o.AllowQuasiCash
	}
	if !IsNil(o.AlwaysRequireIcc) {
		toSerialize["always_require_icc"] = o.AlwaysRequireIcc
	}
	if !IsNil(o.AlwaysRequirePin) {
		toSerialize["always_require_pin"] = o.AlwaysRequirePin
	}
	if !IsNil(o.EnableCreditService) {
		toSerialize["enable_credit_service"] = o.EnableCreditService
	}
	if !IsNil(o.EnablePartialAuthApproval) {
		toSerialize["enable_partial_auth_approval"] = o.EnablePartialAuthApproval
	}
	if !IsNil(o.IgnoreCardSuspendedState) {
		toSerialize["ignore_card_suspended_state"] = o.IgnoreCardSuspendedState
	}
	if !IsNil(o.NotificationLanguage) {
		toSerialize["notification_language"] = o.NotificationLanguage
	}
	if !IsNil(o.QuasiCashExemptMerchantGroupToken) {
		toSerialize["quasi_cash_exempt_merchant_group_token"] = o.QuasiCashExemptMerchantGroupToken
	}
	if !IsNil(o.QuasiCashExemptMids) {
		toSerialize["quasi_cash_exempt_mids"] = o.QuasiCashExemptMids
	}
	if !IsNil(o.RequireCardNotPresentCardSecurityCode) {
		toSerialize["require_card_not_present_card_security_code"] = o.RequireCardNotPresentCardSecurityCode
	}
	if !IsNil(o.StrongCustomerAuthenticationLimits) {
		toSerialize["strong_customer_authentication_limits"] = o.StrongCustomerAuthenticationLimits
	}
	return toSerialize, nil
}

type NullableTransactionControls struct {
	value *TransactionControls
	isSet bool
}

func (v NullableTransactionControls) Get() *TransactionControls {
	return v.value
}

func (v *NullableTransactionControls) Set(val *TransactionControls) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionControls) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionControls) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionControls(val *TransactionControls) *NullableTransactionControls {
	return &NullableTransactionControls{value: val, isSet: true}
}

func (v NullableTransactionControls) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionControls) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


