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

// checks if the DigitalWalletTokenTransitionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DigitalWalletTokenTransitionResponse{}

// DigitalWalletTokenTransitionResponse struct for DigitalWalletTokenTransitionResponse
type DigitalWalletTokenTransitionResponse struct {
	CardSwap *CardSwapHash `json:"card_swap,omitempty"`
	// Mechanism by which the transition was initiated.
	Channel string `json:"channel"`
	// Date and time when the transition was created, in UTC.
	CreatedTime *time.Time `json:"created_time,omitempty"`
	DigitalWalletToken DigitalWalletTokenHash `json:"digital_wallet_token"`
	// Provisioning status of the digital wallet token.
	FulfillmentStatus string `json:"fulfillment_status"`
	// Reason for the transition.
	Reason *string `json:"reason,omitempty"`
	// Standard code describing the reason for the transition:  * *00:* Object activated for the first time * *01:* Requested by you * *02:* Inactivity over time * *03:* This address cannot accept mail or the addressee is unknown * *04:* Negative account balance * *05:* Account under review * *06:* Suspicious activity was identified * *07:* Activity outside the program parameters was identified * *08:* Confirmed fraud was identified * *09:* Matched with an Office of Foreign Assets Control list * *10:* Card was reported lost * *11:* Card information was cloned * *12:* Account or card information was compromised * *13:* Temporary status change while on hold/leave * *14:* Initiated by Marqeta * *15:* Initiated by issuer * *16:* Card expired * *17:* Failed KYC * *18:* Changed to `ACTIVE` because information was properly validated * *19:* Changed to `ACTIVE` because account activity was properly validated * *20:* Change occurred prior to the normalization of reason codes * *21:* Initiated by a third party, often a digital wallet provider * *22:* PIN retry limit reached * *23:* Card was reported stolen * *24:* Address issue * *25:* Name issue * *26:* SSN issue * *27:* DOB issue * *28:* Email issue * *29:* Phone issue * *30:* Account/fulfillment mismatch * *31:* Other reason
	ReasonCode *string `json:"reason_code,omitempty"`
	// Specifies the state to which the digital wallet token is transitioning.
	State string `json:"state"`
	// Unique identifier of the digital wallet token transition, and not the identifier of the digital wallet token itself.
	Token string `json:"token"`
	// Type of digital wallet token transition. `state.activated`, for example.
	Type string `json:"type"`
}

type _DigitalWalletTokenTransitionResponse DigitalWalletTokenTransitionResponse

// NewDigitalWalletTokenTransitionResponse instantiates a new DigitalWalletTokenTransitionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDigitalWalletTokenTransitionResponse(channel string, digitalWalletToken DigitalWalletTokenHash, fulfillmentStatus string, state string, token string, type_ string) *DigitalWalletTokenTransitionResponse {
	this := DigitalWalletTokenTransitionResponse{}
	this.Channel = channel
	this.DigitalWalletToken = digitalWalletToken
	this.FulfillmentStatus = fulfillmentStatus
	this.State = state
	this.Token = token
	this.Type = type_
	return &this
}

// NewDigitalWalletTokenTransitionResponseWithDefaults instantiates a new DigitalWalletTokenTransitionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDigitalWalletTokenTransitionResponseWithDefaults() *DigitalWalletTokenTransitionResponse {
	this := DigitalWalletTokenTransitionResponse{}
	return &this
}

// GetCardSwap returns the CardSwap field value if set, zero value otherwise.
func (o *DigitalWalletTokenTransitionResponse) GetCardSwap() CardSwapHash {
	if o == nil || IsNil(o.CardSwap) {
		var ret CardSwapHash
		return ret
	}
	return *o.CardSwap
}

// GetCardSwapOk returns a tuple with the CardSwap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigitalWalletTokenTransitionResponse) GetCardSwapOk() (*CardSwapHash, bool) {
	if o == nil || IsNil(o.CardSwap) {
		return nil, false
	}
	return o.CardSwap, true
}

// HasCardSwap returns a boolean if a field has been set.
func (o *DigitalWalletTokenTransitionResponse) HasCardSwap() bool {
	if o != nil && !IsNil(o.CardSwap) {
		return true
	}

	return false
}

// SetCardSwap gets a reference to the given CardSwapHash and assigns it to the CardSwap field.
func (o *DigitalWalletTokenTransitionResponse) SetCardSwap(v CardSwapHash) {
	o.CardSwap = &v
}

// GetChannel returns the Channel field value
func (o *DigitalWalletTokenTransitionResponse) GetChannel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Channel
}

// GetChannelOk returns a tuple with the Channel field value
// and a boolean to check if the value has been set.
func (o *DigitalWalletTokenTransitionResponse) GetChannelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Channel, true
}

// SetChannel sets field value
func (o *DigitalWalletTokenTransitionResponse) SetChannel(v string) {
	o.Channel = v
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *DigitalWalletTokenTransitionResponse) GetCreatedTime() time.Time {
	if o == nil || IsNil(o.CreatedTime) {
		var ret time.Time
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigitalWalletTokenTransitionResponse) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedTime) {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *DigitalWalletTokenTransitionResponse) HasCreatedTime() bool {
	if o != nil && !IsNil(o.CreatedTime) {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given time.Time and assigns it to the CreatedTime field.
func (o *DigitalWalletTokenTransitionResponse) SetCreatedTime(v time.Time) {
	o.CreatedTime = &v
}

// GetDigitalWalletToken returns the DigitalWalletToken field value
func (o *DigitalWalletTokenTransitionResponse) GetDigitalWalletToken() DigitalWalletTokenHash {
	if o == nil {
		var ret DigitalWalletTokenHash
		return ret
	}

	return o.DigitalWalletToken
}

// GetDigitalWalletTokenOk returns a tuple with the DigitalWalletToken field value
// and a boolean to check if the value has been set.
func (o *DigitalWalletTokenTransitionResponse) GetDigitalWalletTokenOk() (*DigitalWalletTokenHash, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DigitalWalletToken, true
}

// SetDigitalWalletToken sets field value
func (o *DigitalWalletTokenTransitionResponse) SetDigitalWalletToken(v DigitalWalletTokenHash) {
	o.DigitalWalletToken = v
}

// GetFulfillmentStatus returns the FulfillmentStatus field value
func (o *DigitalWalletTokenTransitionResponse) GetFulfillmentStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FulfillmentStatus
}

// GetFulfillmentStatusOk returns a tuple with the FulfillmentStatus field value
// and a boolean to check if the value has been set.
func (o *DigitalWalletTokenTransitionResponse) GetFulfillmentStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FulfillmentStatus, true
}

// SetFulfillmentStatus sets field value
func (o *DigitalWalletTokenTransitionResponse) SetFulfillmentStatus(v string) {
	o.FulfillmentStatus = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *DigitalWalletTokenTransitionResponse) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigitalWalletTokenTransitionResponse) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *DigitalWalletTokenTransitionResponse) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *DigitalWalletTokenTransitionResponse) SetReason(v string) {
	o.Reason = &v
}

// GetReasonCode returns the ReasonCode field value if set, zero value otherwise.
func (o *DigitalWalletTokenTransitionResponse) GetReasonCode() string {
	if o == nil || IsNil(o.ReasonCode) {
		var ret string
		return ret
	}
	return *o.ReasonCode
}

// GetReasonCodeOk returns a tuple with the ReasonCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigitalWalletTokenTransitionResponse) GetReasonCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ReasonCode) {
		return nil, false
	}
	return o.ReasonCode, true
}

// HasReasonCode returns a boolean if a field has been set.
func (o *DigitalWalletTokenTransitionResponse) HasReasonCode() bool {
	if o != nil && !IsNil(o.ReasonCode) {
		return true
	}

	return false
}

// SetReasonCode gets a reference to the given string and assigns it to the ReasonCode field.
func (o *DigitalWalletTokenTransitionResponse) SetReasonCode(v string) {
	o.ReasonCode = &v
}

// GetState returns the State field value
func (o *DigitalWalletTokenTransitionResponse) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *DigitalWalletTokenTransitionResponse) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *DigitalWalletTokenTransitionResponse) SetState(v string) {
	o.State = v
}

// GetToken returns the Token field value
func (o *DigitalWalletTokenTransitionResponse) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *DigitalWalletTokenTransitionResponse) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *DigitalWalletTokenTransitionResponse) SetToken(v string) {
	o.Token = v
}

// GetType returns the Type field value
func (o *DigitalWalletTokenTransitionResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DigitalWalletTokenTransitionResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DigitalWalletTokenTransitionResponse) SetType(v string) {
	o.Type = v
}

func (o DigitalWalletTokenTransitionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DigitalWalletTokenTransitionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CardSwap) {
		toSerialize["card_swap"] = o.CardSwap
	}
	toSerialize["channel"] = o.Channel
	if !IsNil(o.CreatedTime) {
		toSerialize["created_time"] = o.CreatedTime
	}
	toSerialize["digital_wallet_token"] = o.DigitalWalletToken
	toSerialize["fulfillment_status"] = o.FulfillmentStatus
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !IsNil(o.ReasonCode) {
		toSerialize["reason_code"] = o.ReasonCode
	}
	toSerialize["state"] = o.State
	toSerialize["token"] = o.Token
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *DigitalWalletTokenTransitionResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"channel",
		"digital_wallet_token",
		"fulfillment_status",
		"state",
		"token",
		"type",
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

	varDigitalWalletTokenTransitionResponse := _DigitalWalletTokenTransitionResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDigitalWalletTokenTransitionResponse)

	if err != nil {
		return err
	}

	*o = DigitalWalletTokenTransitionResponse(varDigitalWalletTokenTransitionResponse)

	return err
}

type NullableDigitalWalletTokenTransitionResponse struct {
	value *DigitalWalletTokenTransitionResponse
	isSet bool
}

func (v NullableDigitalWalletTokenTransitionResponse) Get() *DigitalWalletTokenTransitionResponse {
	return v.value
}

func (v *NullableDigitalWalletTokenTransitionResponse) Set(val *DigitalWalletTokenTransitionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDigitalWalletTokenTransitionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDigitalWalletTokenTransitionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDigitalWalletTokenTransitionResponse(val *DigitalWalletTokenTransitionResponse) *NullableDigitalWalletTokenTransitionResponse {
	return &NullableDigitalWalletTokenTransitionResponse{value: val, isSet: true}
}

func (v NullableDigitalWalletTokenTransitionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDigitalWalletTokenTransitionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


