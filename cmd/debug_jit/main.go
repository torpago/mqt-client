package main

import (
	_ "errors"
	"log/slog"
	"strings"

	"github.com/go-chi/render"
	mq "github.com/torpago/mqt-client"
)

func main() {
	body := `{
		"acquirer": {
		  "institution_country": "840",
		  "institution_id_code": "411641",
		  "retrieval_reference_number": "302317150303",
		  "system_trace_audit_number": "150303"
		},
		"acquirer_fee_amount": 0,
		"acquirer_reference_data": "411641",
		"acting_user_token": "4a3456aa-55ac-43f8-86b9-bc44bdd98c2b",
		"amount": 1224077.63,
		"business": {
		  "metadata": {}
		},
		"business_token": "c7ba9563-f4ff-4561-a0e9-f63a6c505f56",
		"card": {
		  "last_four": "4015",
		  "metadata": {}
		},
		"card_acceptor": {
		  "city": "MIAMI",
		  "country_code": "USA",
		  "mcc": "5814",
		  "mid": "000690555",
		  "name": "WINGSTOP 1891",
		  "postal_code": "33175",
		  "state": "12"
		},
		"card_product_token": "aaa4fb2d-b522-4490-b114-b07ba84ece4b",
		"card_token": "58a64185-79c7-44a5-863a-77869dacd553",
		"cardholder_authentication_data": {
		  "verification_result": "not_present"
		},
		"created_time": "2023-01-23T17:42:58Z",
		"currency_code": "USD",
		"currency_conversion": {
		  "network": {
			"conversion_rate": 1,
			"original_amount": 1224077.63,
			"original_currency_code": "840"
		  }
		},
		"fraud": {
		  "issuer_processor": {
			"fraud_score_reasons": [],
			"recommended_action": "APPROVE",
			"risk_level": "low",
			"riskcontrol_tags": [],
			"rule_violations": [],
			"score": 1
		  },
		  "network": {
			"account_risk_score": "00",
			"transaction_risk_score": 1
		  }
		},
		"gpa": {
		  "available_balance": 0,
		  "balances": {
			"USD": {
			  "available_balance": 0,
			  "credit_balance": 0,
			  "currency_code": "USD",
			  "ledger_balance": 0,
			  "pending_credits": 0
			}
		  },
		  "credit_balance": 0,
		  "currency_code": "USD",
		  "ledger_balance": 0,
		  "pending_credits": 0
		},
		"gpa_order": {
		  "amount": 1224077.63,
		  "business_token": "c7ba9563-f4ff-4561-a0e9-f63a6c505f56",
		  "currency_code": "USD",
		  "funding": {
			"amount": 1224077.63,
			"source": {
			  "active": true,
			  "created_time": "2020-02-20T22:06:58Z",
			  "is_default_account": false,
			  "last_modified_time": "2020-02-20T22:06:58Z",
			  "token": "**********93de",
			  "type": "programgateway"
			}
		  },
		  "funding_source_token": "**********93de",
		  
		  "state": "PENDING",
		  "token": "a42985c1-f14c-480e-abc5-03addf290c11",
		  "transaction_token": "87ea06b6-157e-4b7c-a57a-e8ca8356aa99"
		},
		"issuer_payment_node": "ba6d271cbbfbc3af7ee4f41d18665df5",
		"issuer_received_time": "2023-01-23T17:42:58.534Z",
		"network": "VISA",
		"network_metadata": {
		  "product_id": "VISA_PURCHASING"
		},
		"network_reference_id": "0002303023637780665",
		"pos": {
		  "card_data_input_capability": "CHIP",
		  "card_holder_presence": true,
		  "card_presence": true,
		  "cardholder_authentication_method": "UNSPECIFIED",
		  "is_installment": false,
		  "is_recurring": false,
		  "pan_entry_mode": "CHIP_CONTACTLESS",
		  "partial_approval_capable": true,
		  "pin_entry_mode": "TRUE",
		  "pin_present": false,
		  "purchase_amount_only": true,
		  "terminal_attendance": "NO_TERMINAL",
		  "terminal_id": "99999999"
		},
		"request_amount": 1224077.63,
		"settlement_date": "2023-01-24T00:00:00Z",
		"state": "PENDING",
		"subnetwork": "VISANET",
		"token": "603084ee-af44-48ff-be7e-3b53819a7d95",
		"transaction_metadata": {
		  "payment_channel": "OTHER"
		},
		"type": "authorization",
		"user_transaction_time": "2023-01-23T17:42:58Z"
	  }`
	data := &mq.TransactionModel{}

	errb := render.DecodeJSON(strings.NewReader(body), &data)
	slog.Info("body", "amount", data.Amount, "token", data.Token, "error", errb)
}
