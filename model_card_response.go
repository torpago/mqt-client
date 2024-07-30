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
	"github.com/shopspring/decimal"
)

// checks if the CardResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CardResponse{}

// CardResponse Contains information about the card used in the transaction.
type CardResponse struct {
	ActivationActions *ActivationActions `json:"activation_actions,omitempty"`
	// Barcode printed on the card, expressed as numerals.
	Barcode *string `json:"barcode,omitempty"`
	// Unique identifier of the bulk card order.
	BulkIssuanceToken *string `json:"bulk_issuance_token,omitempty"`
	// Unique identifier of the card product.
	CardProductToken *string `json:"card_product_token,omitempty"`
	// Three-digit card verification value (ICVV) stored on the chip of the card.
	ChipCvvNumber *string `json:"chip_cvv_number,omitempty"`
	// Running count of the contactless transactions successfully completed since the last strong customer authentication (SCA) challenge was issued. You can limit the number of contactless transactions that can be performed without issuing an SCA challenge at the card product level.  For more information about strong customer authentication, see <</core-api/card-products, Card Products>>.
	ContactlessExemptionCounter *int32 `json:"contactless_exemption_counter,omitempty"`
	// Running total of the money spent in contactless transactions successfully completed since the last strong customer authentication (SCA) challenge was issued. You can limit the total amount that can be spent in contactless transactions without issuing an SCA challenge at the card product level.  For more information about strong customer authentication, see <</core-api/card-products, Card Products>>.
	ContactlessExemptionTotalAmount *decimal.Decimal `json:"contactless_exemption_total_amount,omitempty"`
	// Date and time when the resource was created, in UTC.
	CreatedTime *time.Time `json:"created_time,omitempty"`
	// Three-digit card verification value (CVV2 or CVC2) printed on the card.
	CvvNumber *string `json:"cvv_number,omitempty"`
	// A value of `true` indicates that you requested expedited processing of the card from your card fulfillment provider.
	Expedite *bool `json:"expedite,omitempty"`
	// Expiration date in `MMyy` format.
	Expiration *string `json:"expiration,omitempty"`
	// Expiration date and time, in UTC.
	ExpirationTime *time.Time `json:"expiration_time"`
	Fulfillment *CardFulfillmentResponse `json:"fulfillment,omitempty"`
	// Card fulfillment status:  * *ISSUED:* Initial state of all newly created/issued cards. * *ORDERED:* Card ordered through the card fulfillment provider. * *REORDERED:* Card reordered through the card fulfillment provider. * *REJECTED:* Card rejected by the card fulfillment provider. * *SHIPPED:* Card shipped by the card fulfillment provider. * *DELIVERED:* Card delivered by the card fulfillment provider. * *DIGITALLY_PRESENTED:* Card digitally presented using the `/cards/{token}/showpan` endpoint; does not affect the delivery of physical cards.
	FulfillmentStatus *string `json:"fulfillment_status,omitempty"`
	// Instrument type of the card:  * *PHYSICAL_MSR:* A physical card with a magnetic stripe. This is the default physical card type. * *PHYSICAL_ICC:* A physical card with an integrated circuit, or \"chip.\" * *PHYSICAL_CONTACTLESS:* A physical card that uses radio frequency identification (RFID) or near-field communication (NFC) to enable payment over a secure radio interface. * *PHYSICAL_COMBO:* A physical card with a chip that also supports contactless payments. * *VIRTUAL_PAN:* A virtual card with a primary account number (PAN).
	InstrumentType *string `json:"instrument_type,omitempty"`
	// Last four digits of the card primary account number (PAN).
	LastFour string `json:"last_four"`
	// Date and time when the resource was last modified, in UTC.
	LastModifiedTime *time.Time `json:"last_modified_time,omitempty"`
	// Associates customer-provided metadata with the card.
	Metadata *map[string]string `json:"metadata,omitempty"`
	// Reissues the specified card (known as the \"source\" card) with a new primary account number (PAN).
	NewPanFromCardToken *string `json:"new_pan_from_card_token,omitempty"`
	// Primary account number (PAN) of the card.
	Pan *string `json:"pan,omitempty"`
	// Specifies if the personal identification number (PIN) has been set for the card.
	PinIsSet *bool `json:"pin_is_set,omitempty"`
	// Reissues the specified card (known as the \"source\" card).
	ReissuePanFromCardToken *string `json:"reissue_pan_from_card_token,omitempty"`
	// Indicates the state of the card.
	State *string `json:"state,omitempty"`
	// Descriptive reason for why the card is in its current state. For example, \"Card activated by cardholder\".
	StateReason *string `json:"state_reason,omitempty"`
	// Unique identifier of the card.
	Token *string `json:"token,omitempty"`
	// Copies the personal identification number (PIN) from the specified source card to the newly created card.
	TranslatePinFromCardToken *string `json:"translate_pin_from_card_token,omitempty"`
	// Unique identifier of the cardholder.
	UserToken *string `json:"user_token,omitempty"`
}

type _CardResponse CardResponse

// NewCardResponse instantiates a new CardResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardResponse(barcode string, cardProductToken string, createdTime time.Time, expiration string, expirationTime time.Time, fulfillmentStatus string, lastFour string, lastModifiedTime time.Time, pan string, pinIsSet bool, state string, stateReason string, token string, userToken string) *CardResponse {
	this := CardResponse{}
	this.Barcode = &barcode
	this.CardProductToken = &cardProductToken
	this.CreatedTime = &createdTime
	var expedite bool = false
	this.Expedite = &expedite
	this.Expiration = &expiration
	this.ExpirationTime = &expirationTime
	this.FulfillmentStatus = &fulfillmentStatus
	this.LastFour = lastFour
	this.LastModifiedTime = &lastModifiedTime
	this.Pan = &pan
	this.PinIsSet = &pinIsSet
	this.State = &state
	this.StateReason = &stateReason
	this.Token = &token
	this.UserToken = &userToken
	return &this
}

// NewCardResponseWithDefaults instantiates a new CardResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardResponseWithDefaults() *CardResponse {
	this := CardResponse{}
	var expedite bool = false
	this.Expedite = &expedite
	var pinIsSet bool = false
	this.PinIsSet = pinIsSet
	return &this
}

// GetActivationActions returns the ActivationActions field value if set, zero value otherwise.
func (o *CardResponse) GetActivationActions() ActivationActions {
	if o == nil || IsNil(o.ActivationActions) {
		var ret ActivationActions
		return ret
	}
	return *o.ActivationActions
}

// GetActivationActionsOk returns a tuple with the ActivationActions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardResponse) GetActivationActionsOk() (*ActivationActions, bool) {
	if o == nil || IsNil(o.ActivationActions) {
		return nil, false
	}
	return o.ActivationActions, true
}

// HasActivationActions returns a boolean if a field has been set.
func (o *CardResponse) HasActivationActions() bool {
	if o != nil && !IsNil(o.ActivationActions) {
		return true
	}

	return false
}

// SetActivationActions gets a reference to the given ActivationActions and assigns it to the ActivationActions field.
func (o *CardResponse) SetActivationActions(v ActivationActions) {
	o.ActivationActions = &v
}

// GetBarcode returns the Barcode field value
func (o *CardResponse) GetBarcode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Barcode
}

// GetBarcodeOk returns a tuple with the Barcode field value
// and a boolean to check if the value has been set.
func (o *CardResponse) GetBarcodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Barcode, true
}

// SetBarcode sets field value
func (o *CardResponse) SetBarcode(v string) {
	o.Barcode = v
}

// GetBulkIssuanceToken returns the BulkIssuanceToken field value if set, zero value otherwise.
func (o *CardResponse) GetBulkIssuanceToken() string {
	if o == nil || IsNil(o.BulkIssuanceToken) {
		var ret string
		return ret
	}
	return *o.BulkIssuanceToken
}

// GetBulkIssuanceTokenOk returns a tuple with the BulkIssuanceToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardResponse) GetBulkIssuanceTokenOk() (*string, bool) {
	if o == nil || IsNil(o.BulkIssuanceToken) {
		return nil, false
	}
	return o.BulkIssuanceToken, true
}

// HasBulkIssuanceToken returns a boolean if a field has been set.
func (o *CardResponse) HasBulkIssuanceToken() bool {
	if o != nil && !IsNil(o.BulkIssuanceToken) {
		return true
	}

	return false
}

// SetBulkIssuanceToken gets a reference to the given string and assigns it to the BulkIssuanceToken field.
func (o *CardResponse) SetBulkIssuanceToken(v string) {
	o.BulkIssuanceToken = &v
}

// GetCardProductToken returns the CardProductToken field value
func (o *CardResponse) GetCardProductToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CardProductToken
}

// GetCardProductTokenOk returns a tuple with the CardProductToken field value
// and a boolean to check if the value has been set.
func (o *CardResponse) GetCardProductTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardProductToken, true
}

// SetCardProductToken sets field value
func (o *CardResponse) SetCardProductToken(v string) {
	o.CardProductToken = v
}

// GetChipCvvNumber returns the ChipCvvNumber field value if set, zero value otherwise.
func (o *CardResponse) GetChipCvvNumber() string {
	if o == nil || IsNil(o.ChipCvvNumber) {
		var ret string
		return ret
	}
	return *o.ChipCvvNumber
}

// GetChipCvvNumberOk returns a tuple with the ChipCvvNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardResponse) GetChipCvvNumberOk() (*string, bool) {
	if o == nil || IsNil(o.ChipCvvNumber) {
		return nil, false
	}
	return o.ChipCvvNumber, true
}

// HasChipCvvNumber returns a boolean if a field has been set.
func (o *CardResponse) HasChipCvvNumber() bool {
	if o != nil && !IsNil(o.ChipCvvNumber) {
		return true
	}

	return false
}

// SetChipCvvNumber gets a reference to the given string and assigns it to the ChipCvvNumber field.
func (o *CardResponse) SetChipCvvNumber(v string) {
	o.ChipCvvNumber = &v
}

// GetContactlessExemptionCounter returns the ContactlessExemptionCounter field value if set, zero value otherwise.
func (o *CardResponse) GetContactlessExemptionCounter() int32 {
	if o == nil || IsNil(o.ContactlessExemptionCounter) {
		var ret int32
		return ret
	}
	return *o.ContactlessExemptionCounter
}

// GetContactlessExemptionCounterOk returns a tuple with the ContactlessExemptionCounter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardResponse) GetContactlessExemptionCounterOk() (*int32, bool) {
	if o == nil || IsNil(o.ContactlessExemptionCounter) {
		return nil, false
	}
	return o.ContactlessExemptionCounter, true
}

// HasContactlessExemptionCounter returns a boolean if a field has been set.
func (o *CardResponse) HasContactlessExemptionCounter() bool {
	if o != nil && !IsNil(o.ContactlessExemptionCounter) {
		return true
	}

	return false
}

// SetContactlessExemptionCounter gets a reference to the given int32 and assigns it to the ContactlessExemptionCounter field.
func (o *CardResponse) SetContactlessExemptionCounter(v int32) {
	o.ContactlessExemptionCounter = &v
}

// GetContactlessExemptionTotalAmount returns the ContactlessExemptionTotalAmount field value if set, zero value otherwise.
func (o *CardResponse) GetContactlessExemptionTotalAmount() decimal.Decimal {
	if o == nil || IsNil(o.ContactlessExemptionTotalAmount) {
		var ret decimal.Decimal
		return ret
	}
	return *o.ContactlessExemptionTotalAmount
}

// GetContactlessExemptionTotalAmountOk returns a tuple with the ContactlessExemptionTotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardResponse) GetContactlessExemptionTotalAmountOk() (*decimal.Decimal, bool) {
	if o == nil || IsNil(o.ContactlessExemptionTotalAmount) {
		return nil, false
	}
	return o.ContactlessExemptionTotalAmount, true
}

// HasContactlessExemptionTotalAmount returns a boolean if a field has been set.
func (o *CardResponse) HasContactlessExemptionTotalAmount() bool {
	if o != nil && !IsNil(o.ContactlessExemptionTotalAmount) {
		return true
	}

	return false
}

// SetContactlessExemptionTotalAmount gets a reference to the given decimal.Decimal and assigns it to the ContactlessExemptionTotalAmount field.
func (o *CardResponse) SetContactlessExemptionTotalAmount(v decimal.Decimal) {
	o.ContactlessExemptionTotalAmount = &v
}

// GetCreatedTime returns the CreatedTime field value
func (o *CardResponse) GetCreatedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value
// and a boolean to check if the value has been set.
func (o *CardResponse) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTime, true
}

// SetCreatedTime sets field value
func (o *CardResponse) SetCreatedTime(v time.Time) {
	o.CreatedTime = v
}

// GetCvvNumber returns the CvvNumber field value if set, zero value otherwise.
func (o *CardResponse) GetCvvNumber() string {
	if o == nil || IsNil(o.CvvNumber) {
		var ret string
		return ret
	}
	return *o.CvvNumber
}

// GetCvvNumberOk returns a tuple with the CvvNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardResponse) GetCvvNumberOk() (*string, bool) {
	if o == nil || IsNil(o.CvvNumber) {
		return nil, false
	}
	return o.CvvNumber, true
}

// HasCvvNumber returns a boolean if a field has been set.
func (o *CardResponse) HasCvvNumber() bool {
	if o != nil && !IsNil(o.CvvNumber) {
		return true
	}

	return false
}

// SetCvvNumber gets a reference to the given string and assigns it to the CvvNumber field.
func (o *CardResponse) SetCvvNumber(v string) {
	o.CvvNumber = &v
}

// GetExpedite returns the Expedite field value if set, zero value otherwise.
func (o *CardResponse) GetExpedite() bool {
	if o == nil || IsNil(o.Expedite) {
		var ret bool
		return ret
	}
	return *o.Expedite
}

// GetExpediteOk returns a tuple with the Expedite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardResponse) GetExpediteOk() (*bool, bool) {
	if o == nil || IsNil(o.Expedite) {
		return nil, false
	}
	return o.Expedite, true
}

// HasExpedite returns a boolean if a field has been set.
func (o *CardResponse) HasExpedite() bool {
	if o != nil && !IsNil(o.Expedite) {
		return true
	}

	return false
}

// SetExpedite gets a reference to the given bool and assigns it to the Expedite field.
func (o *CardResponse) SetExpedite(v bool) {
	o.Expedite = &v
}

// GetExpiration returns the Expiration field value
func (o *CardResponse) GetExpiration() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value
// and a boolean to check if the value has been set.
func (o *CardResponse) GetExpirationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expiration, true
}

// SetExpiration sets field value
func (o *CardResponse) SetExpiration(v string) {
	o.Expiration = v
}

// GetExpirationTime returns the ExpirationTime field value
func (o *CardResponse) GetExpirationTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ExpirationTime
}

// GetExpirationTimeOk returns a tuple with the ExpirationTime field value
// and a boolean to check if the value has been set.
func (o *CardResponse) GetExpirationTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpirationTime, true
}

// SetExpirationTime sets field value
func (o *CardResponse) SetExpirationTime(v time.Time) {
	o.ExpirationTime = v
}

// GetFulfillment returns the Fulfillment field value if set, zero value otherwise.
func (o *CardResponse) GetFulfillment() CardFulfillmentResponse {
	if o == nil || IsNil(o.Fulfillment) {
		var ret CardFulfillmentResponse
		return ret
	}
	return *o.Fulfillment
}

// GetFulfillmentOk returns a tuple with the Fulfillment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardResponse) GetFulfillmentOk() (*CardFulfillmentResponse, bool) {
	if o == nil || IsNil(o.Fulfillment) {
		return nil, false
	}
	return o.Fulfillment, true
}

// HasFulfillment returns a boolean if a field has been set.
func (o *CardResponse) HasFulfillment() bool {
	if o != nil && !IsNil(o.Fulfillment) {
		return true
	}

	return false
}

// SetFulfillment gets a reference to the given CardFulfillmentResponse and assigns it to the Fulfillment field.
func (o *CardResponse) SetFulfillment(v CardFulfillmentResponse) {
	o.Fulfillment = &v
}

// GetFulfillmentStatus returns the FulfillmentStatus field value
func (o *CardResponse) GetFulfillmentStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FulfillmentStatus
}

// GetFulfillmentStatusOk returns a tuple with the FulfillmentStatus field value
// and a boolean to check if the value has been set.
func (o *CardResponse) GetFulfillmentStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FulfillmentStatus, true
}

// SetFulfillmentStatus sets field value
func (o *CardResponse) SetFulfillmentStatus(v string) {
	o.FulfillmentStatus = v
}

// GetInstrumentType returns the InstrumentType field value if set, zero value otherwise.
func (o *CardResponse) GetInstrumentType() string {
	if o == nil || IsNil(o.InstrumentType) {
		var ret string
		return ret
	}
	return *o.InstrumentType
}

// GetInstrumentTypeOk returns a tuple with the InstrumentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardResponse) GetInstrumentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.InstrumentType) {
		return nil, false
	}
	return o.InstrumentType, true
}

// HasInstrumentType returns a boolean if a field has been set.
func (o *CardResponse) HasInstrumentType() bool {
	if o != nil && !IsNil(o.InstrumentType) {
		return true
	}

	return false
}

// SetInstrumentType gets a reference to the given string and assigns it to the InstrumentType field.
func (o *CardResponse) SetInstrumentType(v string) {
	o.InstrumentType = &v
}

// GetLastFour returns the LastFour field value
func (o *CardResponse) GetLastFour() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastFour
}

// GetLastFourOk returns a tuple with the LastFour field value
// and a boolean to check if the value has been set.
func (o *CardResponse) GetLastFourOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastFour, true
}

// SetLastFour sets field value
func (o *CardResponse) SetLastFour(v string) {
	o.LastFour = v
}

// GetLastModifiedTime returns the LastModifiedTime field value
func (o *CardResponse) GetLastModifiedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value
// and a boolean to check if the value has been set.
func (o *CardResponse) GetLastModifiedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastModifiedTime, true
}

// SetLastModifiedTime sets field value
func (o *CardResponse) SetLastModifiedTime(v time.Time) {
	o.LastModifiedTime = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CardResponse) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardResponse) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CardResponse) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *CardResponse) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetNewPanFromCardToken returns the NewPanFromCardToken field value if set, zero value otherwise.
func (o *CardResponse) GetNewPanFromCardToken() string {
	if o == nil || IsNil(o.NewPanFromCardToken) {
		var ret string
		return ret
	}
	return *o.NewPanFromCardToken
}

// GetNewPanFromCardTokenOk returns a tuple with the NewPanFromCardToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardResponse) GetNewPanFromCardTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NewPanFromCardToken) {
		return nil, false
	}
	return o.NewPanFromCardToken, true
}

// HasNewPanFromCardToken returns a boolean if a field has been set.
func (o *CardResponse) HasNewPanFromCardToken() bool {
	if o != nil && !IsNil(o.NewPanFromCardToken) {
		return true
	}

	return false
}

// SetNewPanFromCardToken gets a reference to the given string and assigns it to the NewPanFromCardToken field.
func (o *CardResponse) SetNewPanFromCardToken(v string) {
	o.NewPanFromCardToken = &v
}

// GetPan returns the Pan field value
func (o *CardResponse) GetPan() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pan
}

// GetPanOk returns a tuple with the Pan field value
// and a boolean to check if the value has been set.
func (o *CardResponse) GetPanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pan, true
}

// SetPan sets field value
func (o *CardResponse) SetPan(v string) {
	o.Pan = v
}

// GetPinIsSet returns the PinIsSet field value
func (o *CardResponse) GetPinIsSet() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PinIsSet
}

// GetPinIsSetOk returns a tuple with the PinIsSet field value
// and a boolean to check if the value has been set.
func (o *CardResponse) GetPinIsSetOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PinIsSet, true
}

// SetPinIsSet sets field value
func (o *CardResponse) SetPinIsSet(v bool) {
	o.PinIsSet = v
}

// GetReissuePanFromCardToken returns the ReissuePanFromCardToken field value if set, zero value otherwise.
func (o *CardResponse) GetReissuePanFromCardToken() string {
	if o == nil || IsNil(o.ReissuePanFromCardToken) {
		var ret string
		return ret
	}
	return *o.ReissuePanFromCardToken
}

// GetReissuePanFromCardTokenOk returns a tuple with the ReissuePanFromCardToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardResponse) GetReissuePanFromCardTokenOk() (*string, bool) {
	if o == nil || IsNil(o.ReissuePanFromCardToken) {
		return nil, false
	}
	return o.ReissuePanFromCardToken, true
}

// HasReissuePanFromCardToken returns a boolean if a field has been set.
func (o *CardResponse) HasReissuePanFromCardToken() bool {
	if o != nil && !IsNil(o.ReissuePanFromCardToken) {
		return true
	}

	return false
}

// SetReissuePanFromCardToken gets a reference to the given string and assigns it to the ReissuePanFromCardToken field.
func (o *CardResponse) SetReissuePanFromCardToken(v string) {
	o.ReissuePanFromCardToken = &v
}

// GetState returns the State field value
func (o *CardResponse) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *CardResponse) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *CardResponse) SetState(v string) {
	o.State = v
}

// GetStateReason returns the StateReason field value
func (o *CardResponse) GetStateReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StateReason
}

// GetStateReasonOk returns a tuple with the StateReason field value
// and a boolean to check if the value has been set.
func (o *CardResponse) GetStateReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StateReason, true
}

// SetStateReason sets field value
func (o *CardResponse) SetStateReason(v string) {
	o.StateReason = v
}

// GetToken returns the Token field value
func (o *CardResponse) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *CardResponse) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *CardResponse) SetToken(v string) {
	o.Token = v
}

// GetTranslatePinFromCardToken returns the TranslatePinFromCardToken field value if set, zero value otherwise.
func (o *CardResponse) GetTranslatePinFromCardToken() string {
	if o == nil || IsNil(o.TranslatePinFromCardToken) {
		var ret string
		return ret
	}
	return *o.TranslatePinFromCardToken
}

// GetTranslatePinFromCardTokenOk returns a tuple with the TranslatePinFromCardToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardResponse) GetTranslatePinFromCardTokenOk() (*string, bool) {
	if o == nil || IsNil(o.TranslatePinFromCardToken) {
		return nil, false
	}
	return o.TranslatePinFromCardToken, true
}

// HasTranslatePinFromCardToken returns a boolean if a field has been set.
func (o *CardResponse) HasTranslatePinFromCardToken() bool {
	if o != nil && !IsNil(o.TranslatePinFromCardToken) {
		return true
	}

	return false
}

// SetTranslatePinFromCardToken gets a reference to the given string and assigns it to the TranslatePinFromCardToken field.
func (o *CardResponse) SetTranslatePinFromCardToken(v string) {
	o.TranslatePinFromCardToken = &v
}

// GetUserToken returns the UserToken field value
func (o *CardResponse) GetUserToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserToken
}

// GetUserTokenOk returns a tuple with the UserToken field value
// and a boolean to check if the value has been set.
func (o *CardResponse) GetUserTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserToken, true
}

// SetUserToken sets field value
func (o *CardResponse) SetUserToken(v string) {
	o.UserToken = v
}

func (o CardResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActivationActions) {
		toSerialize["activation_actions"] = o.ActivationActions
	}
	toSerialize["barcode"] = o.Barcode
	if !IsNil(o.BulkIssuanceToken) {
		toSerialize["bulk_issuance_token"] = o.BulkIssuanceToken
	}
	toSerialize["card_product_token"] = o.CardProductToken
	if !IsNil(o.ChipCvvNumber) {
		toSerialize["chip_cvv_number"] = o.ChipCvvNumber
	}
	if !IsNil(o.ContactlessExemptionCounter) {
		toSerialize["contactless_exemption_counter"] = o.ContactlessExemptionCounter
	}
	if !IsNil(o.ContactlessExemptionTotalAmount) {
		toSerialize["contactless_exemption_total_amount"] = o.ContactlessExemptionTotalAmount
	}
	toSerialize["created_time"] = o.CreatedTime
	if !IsNil(o.CvvNumber) {
		toSerialize["cvv_number"] = o.CvvNumber
	}
	if !IsNil(o.Expedite) {
		toSerialize["expedite"] = o.Expedite
	}
	toSerialize["expiration"] = o.Expiration
	toSerialize["expiration_time"] = o.ExpirationTime
	if !IsNil(o.Fulfillment) {
		toSerialize["fulfillment"] = o.Fulfillment
	}
	toSerialize["fulfillment_status"] = o.FulfillmentStatus
	if !IsNil(o.InstrumentType) {
		toSerialize["instrument_type"] = o.InstrumentType
	}
	toSerialize["last_four"] = o.LastFour
	toSerialize["last_modified_time"] = o.LastModifiedTime
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.NewPanFromCardToken) {
		toSerialize["new_pan_from_card_token"] = o.NewPanFromCardToken
	}
	toSerialize["pan"] = o.Pan
	toSerialize["pin_is_set"] = o.PinIsSet
	if !IsNil(o.ReissuePanFromCardToken) {
		toSerialize["reissue_pan_from_card_token"] = o.ReissuePanFromCardToken
	}
	toSerialize["state"] = o.State
	toSerialize["state_reason"] = o.StateReason
	toSerialize["token"] = o.Token
	if !IsNil(o.TranslatePinFromCardToken) {
		toSerialize["translate_pin_from_card_token"] = o.TranslatePinFromCardToken
	}
	toSerialize["user_token"] = o.UserToken
	return toSerialize, nil
}

func (o *CardResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"barcode",
		"card_product_token",
		"created_time",
		"expiration",
		"expiration_time",
		"fulfillment_status",
		"last_four",
		"last_modified_time",
		"pan",
		"pin_is_set",
		"state",
		"state_reason",
		"token",
		"user_token",
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

	varCardResponse := _CardResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCardResponse)

	if err != nil {
		return err
	}

	*o = CardResponse(varCardResponse)

	return err
}

type NullableCardResponse struct {
	value *CardResponse
	isSet bool
}

func (v NullableCardResponse) Get() *CardResponse {
	return v.value
}

func (v *NullableCardResponse) Set(val *CardResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCardResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCardResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardResponse(val *CardResponse) *NullableCardResponse {
	return &NullableCardResponse{value: val, isSet: true}
}

func (v NullableCardResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


