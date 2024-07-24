# DirectDepositResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **decimal.Decimal** |  | [optional] 
**BusinessToken** | Pointer to **string** |  | [optional] 
**CompanyDiscretionaryData** | Pointer to **string** |  | [optional] 
**CompanyEntryDescription** | Pointer to **string** |  | [optional] 
**CompanyIdentification** | Pointer to **string** |  | [optional] 
**CompanyName** | Pointer to **string** |  | [optional] 
**CreatedTime** | Pointer to **time.Time** |  | [optional] 
**DirectDepositAccountToken** | Pointer to **string** |  | [optional] 
**EarlyDirectDeposit** | Pointer to **bool** |  | [optional] [default to false]
**IndividualIdentificationNumber** | Pointer to **string** |  | [optional] 
**IndividualName** | Pointer to **string** |  | [optional] 
**LastModifiedTime** | Pointer to **time.Time** |  | [optional] 
**OriginatorStatusCode** | Pointer to **string** |  | [optional] 
**SettlementDate** | Pointer to **time.Time** |  | [optional] 
**StandardEntryClassCode** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**StateReason** | Pointer to **string** |  | [optional] 
**StateReasonCode** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**TraceNumber** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UserToken** | Pointer to **string** |  | [optional] 

## Methods

### NewDirectDepositResponse

`func NewDirectDepositResponse() *DirectDepositResponse`

NewDirectDepositResponse instantiates a new DirectDepositResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectDepositResponseWithDefaults

`func NewDirectDepositResponseWithDefaults() *DirectDepositResponse`

NewDirectDepositResponseWithDefaults instantiates a new DirectDepositResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *DirectDepositResponse) GetAmount() decimal.Decimal`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DirectDepositResponse) GetAmountOk() (*decimal.Decimal, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DirectDepositResponse) SetAmount(v decimal.Decimal)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *DirectDepositResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetBusinessToken

`func (o *DirectDepositResponse) GetBusinessToken() string`

GetBusinessToken returns the BusinessToken field if non-nil, zero value otherwise.

### GetBusinessTokenOk

`func (o *DirectDepositResponse) GetBusinessTokenOk() (*string, bool)`

GetBusinessTokenOk returns a tuple with the BusinessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessToken

`func (o *DirectDepositResponse) SetBusinessToken(v string)`

SetBusinessToken sets BusinessToken field to given value.

### HasBusinessToken

`func (o *DirectDepositResponse) HasBusinessToken() bool`

HasBusinessToken returns a boolean if a field has been set.

### GetCompanyDiscretionaryData

`func (o *DirectDepositResponse) GetCompanyDiscretionaryData() string`

GetCompanyDiscretionaryData returns the CompanyDiscretionaryData field if non-nil, zero value otherwise.

### GetCompanyDiscretionaryDataOk

`func (o *DirectDepositResponse) GetCompanyDiscretionaryDataOk() (*string, bool)`

GetCompanyDiscretionaryDataOk returns a tuple with the CompanyDiscretionaryData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyDiscretionaryData

`func (o *DirectDepositResponse) SetCompanyDiscretionaryData(v string)`

SetCompanyDiscretionaryData sets CompanyDiscretionaryData field to given value.

### HasCompanyDiscretionaryData

`func (o *DirectDepositResponse) HasCompanyDiscretionaryData() bool`

HasCompanyDiscretionaryData returns a boolean if a field has been set.

### GetCompanyEntryDescription

`func (o *DirectDepositResponse) GetCompanyEntryDescription() string`

GetCompanyEntryDescription returns the CompanyEntryDescription field if non-nil, zero value otherwise.

### GetCompanyEntryDescriptionOk

`func (o *DirectDepositResponse) GetCompanyEntryDescriptionOk() (*string, bool)`

GetCompanyEntryDescriptionOk returns a tuple with the CompanyEntryDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyEntryDescription

`func (o *DirectDepositResponse) SetCompanyEntryDescription(v string)`

SetCompanyEntryDescription sets CompanyEntryDescription field to given value.

### HasCompanyEntryDescription

`func (o *DirectDepositResponse) HasCompanyEntryDescription() bool`

HasCompanyEntryDescription returns a boolean if a field has been set.

### GetCompanyIdentification

`func (o *DirectDepositResponse) GetCompanyIdentification() string`

GetCompanyIdentification returns the CompanyIdentification field if non-nil, zero value otherwise.

### GetCompanyIdentificationOk

`func (o *DirectDepositResponse) GetCompanyIdentificationOk() (*string, bool)`

GetCompanyIdentificationOk returns a tuple with the CompanyIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyIdentification

`func (o *DirectDepositResponse) SetCompanyIdentification(v string)`

SetCompanyIdentification sets CompanyIdentification field to given value.

### HasCompanyIdentification

`func (o *DirectDepositResponse) HasCompanyIdentification() bool`

HasCompanyIdentification returns a boolean if a field has been set.

### GetCompanyName

`func (o *DirectDepositResponse) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *DirectDepositResponse) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *DirectDepositResponse) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *DirectDepositResponse) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCreatedTime

`func (o *DirectDepositResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *DirectDepositResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *DirectDepositResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *DirectDepositResponse) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDirectDepositAccountToken

`func (o *DirectDepositResponse) GetDirectDepositAccountToken() string`

GetDirectDepositAccountToken returns the DirectDepositAccountToken field if non-nil, zero value otherwise.

### GetDirectDepositAccountTokenOk

`func (o *DirectDepositResponse) GetDirectDepositAccountTokenOk() (*string, bool)`

GetDirectDepositAccountTokenOk returns a tuple with the DirectDepositAccountToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectDepositAccountToken

`func (o *DirectDepositResponse) SetDirectDepositAccountToken(v string)`

SetDirectDepositAccountToken sets DirectDepositAccountToken field to given value.

### HasDirectDepositAccountToken

`func (o *DirectDepositResponse) HasDirectDepositAccountToken() bool`

HasDirectDepositAccountToken returns a boolean if a field has been set.

### GetEarlyDirectDeposit

`func (o *DirectDepositResponse) GetEarlyDirectDeposit() bool`

GetEarlyDirectDeposit returns the EarlyDirectDeposit field if non-nil, zero value otherwise.

### GetEarlyDirectDepositOk

`func (o *DirectDepositResponse) GetEarlyDirectDepositOk() (*bool, bool)`

GetEarlyDirectDepositOk returns a tuple with the EarlyDirectDeposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarlyDirectDeposit

`func (o *DirectDepositResponse) SetEarlyDirectDeposit(v bool)`

SetEarlyDirectDeposit sets EarlyDirectDeposit field to given value.

### HasEarlyDirectDeposit

`func (o *DirectDepositResponse) HasEarlyDirectDeposit() bool`

HasEarlyDirectDeposit returns a boolean if a field has been set.

### GetIndividualIdentificationNumber

`func (o *DirectDepositResponse) GetIndividualIdentificationNumber() string`

GetIndividualIdentificationNumber returns the IndividualIdentificationNumber field if non-nil, zero value otherwise.

### GetIndividualIdentificationNumberOk

`func (o *DirectDepositResponse) GetIndividualIdentificationNumberOk() (*string, bool)`

GetIndividualIdentificationNumberOk returns a tuple with the IndividualIdentificationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualIdentificationNumber

`func (o *DirectDepositResponse) SetIndividualIdentificationNumber(v string)`

SetIndividualIdentificationNumber sets IndividualIdentificationNumber field to given value.

### HasIndividualIdentificationNumber

`func (o *DirectDepositResponse) HasIndividualIdentificationNumber() bool`

HasIndividualIdentificationNumber returns a boolean if a field has been set.

### GetIndividualName

`func (o *DirectDepositResponse) GetIndividualName() string`

GetIndividualName returns the IndividualName field if non-nil, zero value otherwise.

### GetIndividualNameOk

`func (o *DirectDepositResponse) GetIndividualNameOk() (*string, bool)`

GetIndividualNameOk returns a tuple with the IndividualName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualName

`func (o *DirectDepositResponse) SetIndividualName(v string)`

SetIndividualName sets IndividualName field to given value.

### HasIndividualName

`func (o *DirectDepositResponse) HasIndividualName() bool`

HasIndividualName returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *DirectDepositResponse) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *DirectDepositResponse) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *DirectDepositResponse) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *DirectDepositResponse) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.

### GetOriginatorStatusCode

`func (o *DirectDepositResponse) GetOriginatorStatusCode() string`

GetOriginatorStatusCode returns the OriginatorStatusCode field if non-nil, zero value otherwise.

### GetOriginatorStatusCodeOk

`func (o *DirectDepositResponse) GetOriginatorStatusCodeOk() (*string, bool)`

GetOriginatorStatusCodeOk returns a tuple with the OriginatorStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatorStatusCode

`func (o *DirectDepositResponse) SetOriginatorStatusCode(v string)`

SetOriginatorStatusCode sets OriginatorStatusCode field to given value.

### HasOriginatorStatusCode

`func (o *DirectDepositResponse) HasOriginatorStatusCode() bool`

HasOriginatorStatusCode returns a boolean if a field has been set.

### GetSettlementDate

`func (o *DirectDepositResponse) GetSettlementDate() time.Time`

GetSettlementDate returns the SettlementDate field if non-nil, zero value otherwise.

### GetSettlementDateOk

`func (o *DirectDepositResponse) GetSettlementDateOk() (*time.Time, bool)`

GetSettlementDateOk returns a tuple with the SettlementDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementDate

`func (o *DirectDepositResponse) SetSettlementDate(v time.Time)`

SetSettlementDate sets SettlementDate field to given value.

### HasSettlementDate

`func (o *DirectDepositResponse) HasSettlementDate() bool`

HasSettlementDate returns a boolean if a field has been set.

### GetStandardEntryClassCode

`func (o *DirectDepositResponse) GetStandardEntryClassCode() string`

GetStandardEntryClassCode returns the StandardEntryClassCode field if non-nil, zero value otherwise.

### GetStandardEntryClassCodeOk

`func (o *DirectDepositResponse) GetStandardEntryClassCodeOk() (*string, bool)`

GetStandardEntryClassCodeOk returns a tuple with the StandardEntryClassCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardEntryClassCode

`func (o *DirectDepositResponse) SetStandardEntryClassCode(v string)`

SetStandardEntryClassCode sets StandardEntryClassCode field to given value.

### HasStandardEntryClassCode

`func (o *DirectDepositResponse) HasStandardEntryClassCode() bool`

HasStandardEntryClassCode returns a boolean if a field has been set.

### GetState

`func (o *DirectDepositResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DirectDepositResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DirectDepositResponse) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DirectDepositResponse) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateReason

`func (o *DirectDepositResponse) GetStateReason() string`

GetStateReason returns the StateReason field if non-nil, zero value otherwise.

### GetStateReasonOk

`func (o *DirectDepositResponse) GetStateReasonOk() (*string, bool)`

GetStateReasonOk returns a tuple with the StateReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateReason

`func (o *DirectDepositResponse) SetStateReason(v string)`

SetStateReason sets StateReason field to given value.

### HasStateReason

`func (o *DirectDepositResponse) HasStateReason() bool`

HasStateReason returns a boolean if a field has been set.

### GetStateReasonCode

`func (o *DirectDepositResponse) GetStateReasonCode() string`

GetStateReasonCode returns the StateReasonCode field if non-nil, zero value otherwise.

### GetStateReasonCodeOk

`func (o *DirectDepositResponse) GetStateReasonCodeOk() (*string, bool)`

GetStateReasonCodeOk returns a tuple with the StateReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateReasonCode

`func (o *DirectDepositResponse) SetStateReasonCode(v string)`

SetStateReasonCode sets StateReasonCode field to given value.

### HasStateReasonCode

`func (o *DirectDepositResponse) HasStateReasonCode() bool`

HasStateReasonCode returns a boolean if a field has been set.

### GetToken

`func (o *DirectDepositResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DirectDepositResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DirectDepositResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DirectDepositResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTraceNumber

`func (o *DirectDepositResponse) GetTraceNumber() string`

GetTraceNumber returns the TraceNumber field if non-nil, zero value otherwise.

### GetTraceNumberOk

`func (o *DirectDepositResponse) GetTraceNumberOk() (*string, bool)`

GetTraceNumberOk returns a tuple with the TraceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceNumber

`func (o *DirectDepositResponse) SetTraceNumber(v string)`

SetTraceNumber sets TraceNumber field to given value.

### HasTraceNumber

`func (o *DirectDepositResponse) HasTraceNumber() bool`

HasTraceNumber returns a boolean if a field has been set.

### GetType

`func (o *DirectDepositResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DirectDepositResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DirectDepositResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DirectDepositResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserToken

`func (o *DirectDepositResponse) GetUserToken() string`

GetUserToken returns the UserToken field if non-nil, zero value otherwise.

### GetUserTokenOk

`func (o *DirectDepositResponse) GetUserTokenOk() (*string, bool)`

GetUserTokenOk returns a tuple with the UserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserToken

`func (o *DirectDepositResponse) SetUserToken(v string)`

SetUserToken sets UserToken field to given value.

### HasUserToken

`func (o *DirectDepositResponse) HasUserToken() bool`

HasUserToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


