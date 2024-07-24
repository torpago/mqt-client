# DirectDepositTransitionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **decimal.Decimal** |  | [optional] 
**BusinessToken** | Pointer to **string** |  | [optional] 
**Channel** | Pointer to **string** |  | [optional] 
**CompanyDiscretionaryData** | Pointer to **string** |  | [optional] 
**CompanyEntryDescription** | Pointer to **string** |  | [optional] 
**CompanyIdentification** | Pointer to **string** |  | [optional] 
**CompanyName** | Pointer to **string** |  | [optional] 
**CreatedTime** | Pointer to **time.Time** |  | [optional] 
**DirectDepositAccountToken** | Pointer to **string** |  | [optional] 
**DirectDepositToken** | Pointer to **string** |  | [optional] 
**EarlyDirectDeposit** | Pointer to **bool** |  | [optional] [default to false]
**IndividualIdentificationNumber** | Pointer to **string** |  | [optional] 
**IndividualName** | Pointer to **string** |  | [optional] 
**LastModifiedTime** | Pointer to **time.Time** |  | [optional] 
**OriginatorStatusCode** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**ReasonCode** | Pointer to **string** |  | [optional] 
**SettlementDate** | Pointer to **time.Time** |  | [optional] 
**StandardEntryClassCode** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**TraceNumber** | Pointer to **string** |  | [optional] 
**TransactionToken** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UserToken** | Pointer to **string** |  | [optional] 

## Methods

### NewDirectDepositTransitionResponse

`func NewDirectDepositTransitionResponse() *DirectDepositTransitionResponse`

NewDirectDepositTransitionResponse instantiates a new DirectDepositTransitionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectDepositTransitionResponseWithDefaults

`func NewDirectDepositTransitionResponseWithDefaults() *DirectDepositTransitionResponse`

NewDirectDepositTransitionResponseWithDefaults instantiates a new DirectDepositTransitionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *DirectDepositTransitionResponse) GetAmount() decimal.Decimal`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DirectDepositTransitionResponse) GetAmountOk() (*decimal.Decimal, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DirectDepositTransitionResponse) SetAmount(v decimal.Decimal)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *DirectDepositTransitionResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetBusinessToken

`func (o *DirectDepositTransitionResponse) GetBusinessToken() string`

GetBusinessToken returns the BusinessToken field if non-nil, zero value otherwise.

### GetBusinessTokenOk

`func (o *DirectDepositTransitionResponse) GetBusinessTokenOk() (*string, bool)`

GetBusinessTokenOk returns a tuple with the BusinessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessToken

`func (o *DirectDepositTransitionResponse) SetBusinessToken(v string)`

SetBusinessToken sets BusinessToken field to given value.

### HasBusinessToken

`func (o *DirectDepositTransitionResponse) HasBusinessToken() bool`

HasBusinessToken returns a boolean if a field has been set.

### GetChannel

`func (o *DirectDepositTransitionResponse) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *DirectDepositTransitionResponse) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *DirectDepositTransitionResponse) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *DirectDepositTransitionResponse) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetCompanyDiscretionaryData

`func (o *DirectDepositTransitionResponse) GetCompanyDiscretionaryData() string`

GetCompanyDiscretionaryData returns the CompanyDiscretionaryData field if non-nil, zero value otherwise.

### GetCompanyDiscretionaryDataOk

`func (o *DirectDepositTransitionResponse) GetCompanyDiscretionaryDataOk() (*string, bool)`

GetCompanyDiscretionaryDataOk returns a tuple with the CompanyDiscretionaryData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyDiscretionaryData

`func (o *DirectDepositTransitionResponse) SetCompanyDiscretionaryData(v string)`

SetCompanyDiscretionaryData sets CompanyDiscretionaryData field to given value.

### HasCompanyDiscretionaryData

`func (o *DirectDepositTransitionResponse) HasCompanyDiscretionaryData() bool`

HasCompanyDiscretionaryData returns a boolean if a field has been set.

### GetCompanyEntryDescription

`func (o *DirectDepositTransitionResponse) GetCompanyEntryDescription() string`

GetCompanyEntryDescription returns the CompanyEntryDescription field if non-nil, zero value otherwise.

### GetCompanyEntryDescriptionOk

`func (o *DirectDepositTransitionResponse) GetCompanyEntryDescriptionOk() (*string, bool)`

GetCompanyEntryDescriptionOk returns a tuple with the CompanyEntryDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyEntryDescription

`func (o *DirectDepositTransitionResponse) SetCompanyEntryDescription(v string)`

SetCompanyEntryDescription sets CompanyEntryDescription field to given value.

### HasCompanyEntryDescription

`func (o *DirectDepositTransitionResponse) HasCompanyEntryDescription() bool`

HasCompanyEntryDescription returns a boolean if a field has been set.

### GetCompanyIdentification

`func (o *DirectDepositTransitionResponse) GetCompanyIdentification() string`

GetCompanyIdentification returns the CompanyIdentification field if non-nil, zero value otherwise.

### GetCompanyIdentificationOk

`func (o *DirectDepositTransitionResponse) GetCompanyIdentificationOk() (*string, bool)`

GetCompanyIdentificationOk returns a tuple with the CompanyIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyIdentification

`func (o *DirectDepositTransitionResponse) SetCompanyIdentification(v string)`

SetCompanyIdentification sets CompanyIdentification field to given value.

### HasCompanyIdentification

`func (o *DirectDepositTransitionResponse) HasCompanyIdentification() bool`

HasCompanyIdentification returns a boolean if a field has been set.

### GetCompanyName

`func (o *DirectDepositTransitionResponse) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *DirectDepositTransitionResponse) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *DirectDepositTransitionResponse) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *DirectDepositTransitionResponse) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCreatedTime

`func (o *DirectDepositTransitionResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *DirectDepositTransitionResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *DirectDepositTransitionResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *DirectDepositTransitionResponse) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDirectDepositAccountToken

`func (o *DirectDepositTransitionResponse) GetDirectDepositAccountToken() string`

GetDirectDepositAccountToken returns the DirectDepositAccountToken field if non-nil, zero value otherwise.

### GetDirectDepositAccountTokenOk

`func (o *DirectDepositTransitionResponse) GetDirectDepositAccountTokenOk() (*string, bool)`

GetDirectDepositAccountTokenOk returns a tuple with the DirectDepositAccountToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectDepositAccountToken

`func (o *DirectDepositTransitionResponse) SetDirectDepositAccountToken(v string)`

SetDirectDepositAccountToken sets DirectDepositAccountToken field to given value.

### HasDirectDepositAccountToken

`func (o *DirectDepositTransitionResponse) HasDirectDepositAccountToken() bool`

HasDirectDepositAccountToken returns a boolean if a field has been set.

### GetDirectDepositToken

`func (o *DirectDepositTransitionResponse) GetDirectDepositToken() string`

GetDirectDepositToken returns the DirectDepositToken field if non-nil, zero value otherwise.

### GetDirectDepositTokenOk

`func (o *DirectDepositTransitionResponse) GetDirectDepositTokenOk() (*string, bool)`

GetDirectDepositTokenOk returns a tuple with the DirectDepositToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectDepositToken

`func (o *DirectDepositTransitionResponse) SetDirectDepositToken(v string)`

SetDirectDepositToken sets DirectDepositToken field to given value.

### HasDirectDepositToken

`func (o *DirectDepositTransitionResponse) HasDirectDepositToken() bool`

HasDirectDepositToken returns a boolean if a field has been set.

### GetEarlyDirectDeposit

`func (o *DirectDepositTransitionResponse) GetEarlyDirectDeposit() bool`

GetEarlyDirectDeposit returns the EarlyDirectDeposit field if non-nil, zero value otherwise.

### GetEarlyDirectDepositOk

`func (o *DirectDepositTransitionResponse) GetEarlyDirectDepositOk() (*bool, bool)`

GetEarlyDirectDepositOk returns a tuple with the EarlyDirectDeposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarlyDirectDeposit

`func (o *DirectDepositTransitionResponse) SetEarlyDirectDeposit(v bool)`

SetEarlyDirectDeposit sets EarlyDirectDeposit field to given value.

### HasEarlyDirectDeposit

`func (o *DirectDepositTransitionResponse) HasEarlyDirectDeposit() bool`

HasEarlyDirectDeposit returns a boolean if a field has been set.

### GetIndividualIdentificationNumber

`func (o *DirectDepositTransitionResponse) GetIndividualIdentificationNumber() string`

GetIndividualIdentificationNumber returns the IndividualIdentificationNumber field if non-nil, zero value otherwise.

### GetIndividualIdentificationNumberOk

`func (o *DirectDepositTransitionResponse) GetIndividualIdentificationNumberOk() (*string, bool)`

GetIndividualIdentificationNumberOk returns a tuple with the IndividualIdentificationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualIdentificationNumber

`func (o *DirectDepositTransitionResponse) SetIndividualIdentificationNumber(v string)`

SetIndividualIdentificationNumber sets IndividualIdentificationNumber field to given value.

### HasIndividualIdentificationNumber

`func (o *DirectDepositTransitionResponse) HasIndividualIdentificationNumber() bool`

HasIndividualIdentificationNumber returns a boolean if a field has been set.

### GetIndividualName

`func (o *DirectDepositTransitionResponse) GetIndividualName() string`

GetIndividualName returns the IndividualName field if non-nil, zero value otherwise.

### GetIndividualNameOk

`func (o *DirectDepositTransitionResponse) GetIndividualNameOk() (*string, bool)`

GetIndividualNameOk returns a tuple with the IndividualName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualName

`func (o *DirectDepositTransitionResponse) SetIndividualName(v string)`

SetIndividualName sets IndividualName field to given value.

### HasIndividualName

`func (o *DirectDepositTransitionResponse) HasIndividualName() bool`

HasIndividualName returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *DirectDepositTransitionResponse) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *DirectDepositTransitionResponse) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *DirectDepositTransitionResponse) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *DirectDepositTransitionResponse) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.

### GetOriginatorStatusCode

`func (o *DirectDepositTransitionResponse) GetOriginatorStatusCode() string`

GetOriginatorStatusCode returns the OriginatorStatusCode field if non-nil, zero value otherwise.

### GetOriginatorStatusCodeOk

`func (o *DirectDepositTransitionResponse) GetOriginatorStatusCodeOk() (*string, bool)`

GetOriginatorStatusCodeOk returns a tuple with the OriginatorStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatorStatusCode

`func (o *DirectDepositTransitionResponse) SetOriginatorStatusCode(v string)`

SetOriginatorStatusCode sets OriginatorStatusCode field to given value.

### HasOriginatorStatusCode

`func (o *DirectDepositTransitionResponse) HasOriginatorStatusCode() bool`

HasOriginatorStatusCode returns a boolean if a field has been set.

### GetReason

`func (o *DirectDepositTransitionResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *DirectDepositTransitionResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *DirectDepositTransitionResponse) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *DirectDepositTransitionResponse) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetReasonCode

`func (o *DirectDepositTransitionResponse) GetReasonCode() string`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *DirectDepositTransitionResponse) GetReasonCodeOk() (*string, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *DirectDepositTransitionResponse) SetReasonCode(v string)`

SetReasonCode sets ReasonCode field to given value.

### HasReasonCode

`func (o *DirectDepositTransitionResponse) HasReasonCode() bool`

HasReasonCode returns a boolean if a field has been set.

### GetSettlementDate

`func (o *DirectDepositTransitionResponse) GetSettlementDate() time.Time`

GetSettlementDate returns the SettlementDate field if non-nil, zero value otherwise.

### GetSettlementDateOk

`func (o *DirectDepositTransitionResponse) GetSettlementDateOk() (*time.Time, bool)`

GetSettlementDateOk returns a tuple with the SettlementDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementDate

`func (o *DirectDepositTransitionResponse) SetSettlementDate(v time.Time)`

SetSettlementDate sets SettlementDate field to given value.

### HasSettlementDate

`func (o *DirectDepositTransitionResponse) HasSettlementDate() bool`

HasSettlementDate returns a boolean if a field has been set.

### GetStandardEntryClassCode

`func (o *DirectDepositTransitionResponse) GetStandardEntryClassCode() string`

GetStandardEntryClassCode returns the StandardEntryClassCode field if non-nil, zero value otherwise.

### GetStandardEntryClassCodeOk

`func (o *DirectDepositTransitionResponse) GetStandardEntryClassCodeOk() (*string, bool)`

GetStandardEntryClassCodeOk returns a tuple with the StandardEntryClassCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardEntryClassCode

`func (o *DirectDepositTransitionResponse) SetStandardEntryClassCode(v string)`

SetStandardEntryClassCode sets StandardEntryClassCode field to given value.

### HasStandardEntryClassCode

`func (o *DirectDepositTransitionResponse) HasStandardEntryClassCode() bool`

HasStandardEntryClassCode returns a boolean if a field has been set.

### GetState

`func (o *DirectDepositTransitionResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DirectDepositTransitionResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DirectDepositTransitionResponse) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DirectDepositTransitionResponse) HasState() bool`

HasState returns a boolean if a field has been set.

### GetToken

`func (o *DirectDepositTransitionResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DirectDepositTransitionResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DirectDepositTransitionResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DirectDepositTransitionResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTraceNumber

`func (o *DirectDepositTransitionResponse) GetTraceNumber() string`

GetTraceNumber returns the TraceNumber field if non-nil, zero value otherwise.

### GetTraceNumberOk

`func (o *DirectDepositTransitionResponse) GetTraceNumberOk() (*string, bool)`

GetTraceNumberOk returns a tuple with the TraceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceNumber

`func (o *DirectDepositTransitionResponse) SetTraceNumber(v string)`

SetTraceNumber sets TraceNumber field to given value.

### HasTraceNumber

`func (o *DirectDepositTransitionResponse) HasTraceNumber() bool`

HasTraceNumber returns a boolean if a field has been set.

### GetTransactionToken

`func (o *DirectDepositTransitionResponse) GetTransactionToken() string`

GetTransactionToken returns the TransactionToken field if non-nil, zero value otherwise.

### GetTransactionTokenOk

`func (o *DirectDepositTransitionResponse) GetTransactionTokenOk() (*string, bool)`

GetTransactionTokenOk returns a tuple with the TransactionToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionToken

`func (o *DirectDepositTransitionResponse) SetTransactionToken(v string)`

SetTransactionToken sets TransactionToken field to given value.

### HasTransactionToken

`func (o *DirectDepositTransitionResponse) HasTransactionToken() bool`

HasTransactionToken returns a boolean if a field has been set.

### GetType

`func (o *DirectDepositTransitionResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DirectDepositTransitionResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DirectDepositTransitionResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DirectDepositTransitionResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserToken

`func (o *DirectDepositTransitionResponse) GetUserToken() string`

GetUserToken returns the UserToken field if non-nil, zero value otherwise.

### GetUserTokenOk

`func (o *DirectDepositTransitionResponse) GetUserTokenOk() (*string, bool)`

GetUserTokenOk returns a tuple with the UserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserToken

`func (o *DirectDepositTransitionResponse) SetUserToken(v string)`

SetUserToken sets UserToken field to given value.

### HasUserToken

`func (o *DirectDepositTransitionResponse) HasUserToken() bool`

HasUserToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


