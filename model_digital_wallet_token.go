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
)

// checks if the DigitalWalletToken type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DigitalWalletToken{}

// DigitalWalletToken Contains information about the digital wallet that funded the transaction.  Returned for all transactions funded by a digital wallet or related to digital wallet token provisioning.  For more on digital wallets, see the <</core-api/digital-wallets-management, Digital Wallets Management>> API reference and <</developer-guides/digital-wallets-and-tokenization, Digital Wallets and Tokenization>> developer guide.
type DigitalWalletToken struct {
	AddressVerification *AddressVerification `json:"address_verification,omitempty"`
	// Unique identifier of the card.
	CardToken *string `json:"card_token,omitempty"`
	// Date and time when the digital wallet token object was created, in UTC.
	CreatedTime *time.Time `json:"created_time,omitempty"`
	Device *Device `json:"device,omitempty"`
	// Digital wallet token's provisioning status.  For fulfillment status descriptions, see <</core-api/digital-wallets-management#postDigitalwallettokentransitions, Create digital wallet token transition>>.
	FulfillmentStatus *string `json:"fulfillment_status,omitempty"`
	// The Marqeta platform's decision as to whether the digital wallet token should be provisioned.  * *0000* – The token should be provisioned.  * *token.activation.verification.required* – Provisioning is pending; further action is required for completion.  For all other values, check the value of the `fulfillment_status` field to definitively ascertain the provisioning outcome.  *NOTE:* The value `invalid.cid` indicates an invalid CVV2 number.
	IssuerEligibilityDecision *string `json:"issuer_eligibility_decision,omitempty"`
	// Date and time when the digital wallet token object was last modified, in UTC.
	LastModifiedTime *time.Time `json:"last_modified_time,omitempty"`
	Metadata *DigitalWalletTokenMetadata `json:"metadata,omitempty"`
	// State of the digital wallet token.  For state descriptions, see <</developer-guides/managing-the-digital-wallet-token-lifecycle#_transitioning_token_states, Transitioning Token States>>.
	State *string `json:"state,omitempty"`
	// Reason why the digital wallet token transitioned to its current state.
	StateReason *string `json:"state_reason,omitempty"`
	// Unique identifier of the digital wallet token.
	Token *string `json:"token,omitempty"`
	TokenServiceProvider *TokenServiceProvider `json:"token_service_provider,omitempty"`
	TransactionDevice *TransactionDevice `json:"transaction_device,omitempty"`
	User *UserCardHolderResponse `json:"user,omitempty"`
	WalletProviderProfile *WalletProviderProfile `json:"wallet_provider_profile,omitempty"`
}

// NewDigitalWalletToken instantiates a new DigitalWalletToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDigitalWalletToken() *DigitalWalletToken {
	this := DigitalWalletToken{}
	return &this
}

// NewDigitalWalletTokenWithDefaults instantiates a new DigitalWalletToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDigitalWalletTokenWithDefaults() *DigitalWalletToken {
	this := DigitalWalletToken{}
	return &this
}

// GetAddressVerification returns the AddressVerification field value if set, zero value otherwise.
func (o *DigitalWalletToken) GetAddressVerification() AddressVerification {
	if o == nil || IsNil(o.AddressVerification) {
		var ret AddressVerification
		return ret
	}
	return *o.AddressVerification
}

// GetAddressVerificationOk returns a tuple with the AddressVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigitalWalletToken) GetAddressVerificationOk() (*AddressVerification, bool) {
	if o == nil || IsNil(o.AddressVerification) {
		return nil, false
	}
	return o.AddressVerification, true
}

// HasAddressVerification returns a boolean if a field has been set.
func (o *DigitalWalletToken) HasAddressVerification() bool {
	if o != nil && !IsNil(o.AddressVerification) {
		return true
	}

	return false
}

// SetAddressVerification gets a reference to the given AddressVerification and assigns it to the AddressVerification field.
func (o *DigitalWalletToken) SetAddressVerification(v AddressVerification) {
	o.AddressVerification = &v
}

// GetCardToken returns the CardToken field value if set, zero value otherwise.
func (o *DigitalWalletToken) GetCardToken() string {
	if o == nil || IsNil(o.CardToken) {
		var ret string
		return ret
	}
	return *o.CardToken
}

// GetCardTokenOk returns a tuple with the CardToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigitalWalletToken) GetCardTokenOk() (*string, bool) {
	if o == nil || IsNil(o.CardToken) {
		return nil, false
	}
	return o.CardToken, true
}

// HasCardToken returns a boolean if a field has been set.
func (o *DigitalWalletToken) HasCardToken() bool {
	if o != nil && !IsNil(o.CardToken) {
		return true
	}

	return false
}

// SetCardToken gets a reference to the given string and assigns it to the CardToken field.
func (o *DigitalWalletToken) SetCardToken(v string) {
	o.CardToken = &v
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *DigitalWalletToken) GetCreatedTime() time.Time {
	if o == nil || IsNil(o.CreatedTime) {
		var ret time.Time
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigitalWalletToken) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedTime) {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *DigitalWalletToken) HasCreatedTime() bool {
	if o != nil && !IsNil(o.CreatedTime) {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given time.Time and assigns it to the CreatedTime field.
func (o *DigitalWalletToken) SetCreatedTime(v time.Time) {
	o.CreatedTime = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *DigitalWalletToken) GetDevice() Device {
	if o == nil || IsNil(o.Device) {
		var ret Device
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigitalWalletToken) GetDeviceOk() (*Device, bool) {
	if o == nil || IsNil(o.Device) {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *DigitalWalletToken) HasDevice() bool {
	if o != nil && !IsNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given Device and assigns it to the Device field.
func (o *DigitalWalletToken) SetDevice(v Device) {
	o.Device = &v
}

// GetFulfillmentStatus returns the FulfillmentStatus field value if set, zero value otherwise.
func (o *DigitalWalletToken) GetFulfillmentStatus() string {
	if o == nil || IsNil(o.FulfillmentStatus) {
		var ret string
		return ret
	}
	return *o.FulfillmentStatus
}

// GetFulfillmentStatusOk returns a tuple with the FulfillmentStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigitalWalletToken) GetFulfillmentStatusOk() (*string, bool) {
	if o == nil || IsNil(o.FulfillmentStatus) {
		return nil, false
	}
	return o.FulfillmentStatus, true
}

// HasFulfillmentStatus returns a boolean if a field has been set.
func (o *DigitalWalletToken) HasFulfillmentStatus() bool {
	if o != nil && !IsNil(o.FulfillmentStatus) {
		return true
	}

	return false
}

// SetFulfillmentStatus gets a reference to the given string and assigns it to the FulfillmentStatus field.
func (o *DigitalWalletToken) SetFulfillmentStatus(v string) {
	o.FulfillmentStatus = &v
}

// GetIssuerEligibilityDecision returns the IssuerEligibilityDecision field value if set, zero value otherwise.
func (o *DigitalWalletToken) GetIssuerEligibilityDecision() string {
	if o == nil || IsNil(o.IssuerEligibilityDecision) {
		var ret string
		return ret
	}
	return *o.IssuerEligibilityDecision
}

// GetIssuerEligibilityDecisionOk returns a tuple with the IssuerEligibilityDecision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigitalWalletToken) GetIssuerEligibilityDecisionOk() (*string, bool) {
	if o == nil || IsNil(o.IssuerEligibilityDecision) {
		return nil, false
	}
	return o.IssuerEligibilityDecision, true
}

// HasIssuerEligibilityDecision returns a boolean if a field has been set.
func (o *DigitalWalletToken) HasIssuerEligibilityDecision() bool {
	if o != nil && !IsNil(o.IssuerEligibilityDecision) {
		return true
	}

	return false
}

// SetIssuerEligibilityDecision gets a reference to the given string and assigns it to the IssuerEligibilityDecision field.
func (o *DigitalWalletToken) SetIssuerEligibilityDecision(v string) {
	o.IssuerEligibilityDecision = &v
}

// GetLastModifiedTime returns the LastModifiedTime field value if set, zero value otherwise.
func (o *DigitalWalletToken) GetLastModifiedTime() time.Time {
	if o == nil || IsNil(o.LastModifiedTime) {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigitalWalletToken) GetLastModifiedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModifiedTime) {
		return nil, false
	}
	return o.LastModifiedTime, true
}

// HasLastModifiedTime returns a boolean if a field has been set.
func (o *DigitalWalletToken) HasLastModifiedTime() bool {
	if o != nil && !IsNil(o.LastModifiedTime) {
		return true
	}

	return false
}

// SetLastModifiedTime gets a reference to the given time.Time and assigns it to the LastModifiedTime field.
func (o *DigitalWalletToken) SetLastModifiedTime(v time.Time) {
	o.LastModifiedTime = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *DigitalWalletToken) GetMetadata() DigitalWalletTokenMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret DigitalWalletTokenMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigitalWalletToken) GetMetadataOk() (*DigitalWalletTokenMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *DigitalWalletToken) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given DigitalWalletTokenMetadata and assigns it to the Metadata field.
func (o *DigitalWalletToken) SetMetadata(v DigitalWalletTokenMetadata) {
	o.Metadata = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *DigitalWalletToken) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigitalWalletToken) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *DigitalWalletToken) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *DigitalWalletToken) SetState(v string) {
	o.State = &v
}

// GetStateReason returns the StateReason field value if set, zero value otherwise.
func (o *DigitalWalletToken) GetStateReason() string {
	if o == nil || IsNil(o.StateReason) {
		var ret string
		return ret
	}
	return *o.StateReason
}

// GetStateReasonOk returns a tuple with the StateReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigitalWalletToken) GetStateReasonOk() (*string, bool) {
	if o == nil || IsNil(o.StateReason) {
		return nil, false
	}
	return o.StateReason, true
}

// HasStateReason returns a boolean if a field has been set.
func (o *DigitalWalletToken) HasStateReason() bool {
	if o != nil && !IsNil(o.StateReason) {
		return true
	}

	return false
}

// SetStateReason gets a reference to the given string and assigns it to the StateReason field.
func (o *DigitalWalletToken) SetStateReason(v string) {
	o.StateReason = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *DigitalWalletToken) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigitalWalletToken) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *DigitalWalletToken) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *DigitalWalletToken) SetToken(v string) {
	o.Token = &v
}

// GetTokenServiceProvider returns the TokenServiceProvider field value if set, zero value otherwise.
func (o *DigitalWalletToken) GetTokenServiceProvider() TokenServiceProvider {
	if o == nil || IsNil(o.TokenServiceProvider) {
		var ret TokenServiceProvider
		return ret
	}
	return *o.TokenServiceProvider
}

// GetTokenServiceProviderOk returns a tuple with the TokenServiceProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigitalWalletToken) GetTokenServiceProviderOk() (*TokenServiceProvider, bool) {
	if o == nil || IsNil(o.TokenServiceProvider) {
		return nil, false
	}
	return o.TokenServiceProvider, true
}

// HasTokenServiceProvider returns a boolean if a field has been set.
func (o *DigitalWalletToken) HasTokenServiceProvider() bool {
	if o != nil && !IsNil(o.TokenServiceProvider) {
		return true
	}

	return false
}

// SetTokenServiceProvider gets a reference to the given TokenServiceProvider and assigns it to the TokenServiceProvider field.
func (o *DigitalWalletToken) SetTokenServiceProvider(v TokenServiceProvider) {
	o.TokenServiceProvider = &v
}

// GetTransactionDevice returns the TransactionDevice field value if set, zero value otherwise.
func (o *DigitalWalletToken) GetTransactionDevice() TransactionDevice {
	if o == nil || IsNil(o.TransactionDevice) {
		var ret TransactionDevice
		return ret
	}
	return *o.TransactionDevice
}

// GetTransactionDeviceOk returns a tuple with the TransactionDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigitalWalletToken) GetTransactionDeviceOk() (*TransactionDevice, bool) {
	if o == nil || IsNil(o.TransactionDevice) {
		return nil, false
	}
	return o.TransactionDevice, true
}

// HasTransactionDevice returns a boolean if a field has been set.
func (o *DigitalWalletToken) HasTransactionDevice() bool {
	if o != nil && !IsNil(o.TransactionDevice) {
		return true
	}

	return false
}

// SetTransactionDevice gets a reference to the given TransactionDevice and assigns it to the TransactionDevice field.
func (o *DigitalWalletToken) SetTransactionDevice(v TransactionDevice) {
	o.TransactionDevice = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *DigitalWalletToken) GetUser() UserCardHolderResponse {
	if o == nil || IsNil(o.User) {
		var ret UserCardHolderResponse
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigitalWalletToken) GetUserOk() (*UserCardHolderResponse, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *DigitalWalletToken) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given UserCardHolderResponse and assigns it to the User field.
func (o *DigitalWalletToken) SetUser(v UserCardHolderResponse) {
	o.User = &v
}

// GetWalletProviderProfile returns the WalletProviderProfile field value if set, zero value otherwise.
func (o *DigitalWalletToken) GetWalletProviderProfile() WalletProviderProfile {
	if o == nil || IsNil(o.WalletProviderProfile) {
		var ret WalletProviderProfile
		return ret
	}
	return *o.WalletProviderProfile
}

// GetWalletProviderProfileOk returns a tuple with the WalletProviderProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigitalWalletToken) GetWalletProviderProfileOk() (*WalletProviderProfile, bool) {
	if o == nil || IsNil(o.WalletProviderProfile) {
		return nil, false
	}
	return o.WalletProviderProfile, true
}

// HasWalletProviderProfile returns a boolean if a field has been set.
func (o *DigitalWalletToken) HasWalletProviderProfile() bool {
	if o != nil && !IsNil(o.WalletProviderProfile) {
		return true
	}

	return false
}

// SetWalletProviderProfile gets a reference to the given WalletProviderProfile and assigns it to the WalletProviderProfile field.
func (o *DigitalWalletToken) SetWalletProviderProfile(v WalletProviderProfile) {
	o.WalletProviderProfile = &v
}

func (o DigitalWalletToken) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DigitalWalletToken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddressVerification) {
		toSerialize["address_verification"] = o.AddressVerification
	}
	if !IsNil(o.CardToken) {
		toSerialize["card_token"] = o.CardToken
	}
	if !IsNil(o.CreatedTime) {
		toSerialize["created_time"] = o.CreatedTime
	}
	if !IsNil(o.Device) {
		toSerialize["device"] = o.Device
	}
	if !IsNil(o.FulfillmentStatus) {
		toSerialize["fulfillment_status"] = o.FulfillmentStatus
	}
	if !IsNil(o.IssuerEligibilityDecision) {
		toSerialize["issuer_eligibility_decision"] = o.IssuerEligibilityDecision
	}
	if !IsNil(o.LastModifiedTime) {
		toSerialize["last_modified_time"] = o.LastModifiedTime
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.StateReason) {
		toSerialize["state_reason"] = o.StateReason
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.TokenServiceProvider) {
		toSerialize["token_service_provider"] = o.TokenServiceProvider
	}
	if !IsNil(o.TransactionDevice) {
		toSerialize["transaction_device"] = o.TransactionDevice
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !IsNil(o.WalletProviderProfile) {
		toSerialize["wallet_provider_profile"] = o.WalletProviderProfile
	}
	return toSerialize, nil
}

type NullableDigitalWalletToken struct {
	value *DigitalWalletToken
	isSet bool
}

func (v NullableDigitalWalletToken) Get() *DigitalWalletToken {
	return v.value
}

func (v *NullableDigitalWalletToken) Set(val *DigitalWalletToken) {
	v.value = val
	v.isSet = true
}

func (v NullableDigitalWalletToken) IsSet() bool {
	return v.isSet
}

func (v *NullableDigitalWalletToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDigitalWalletToken(val *DigitalWalletToken) *NullableDigitalWalletToken {
	return &NullableDigitalWalletToken{value: val, isSet: true}
}

func (v NullableDigitalWalletToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDigitalWalletToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


