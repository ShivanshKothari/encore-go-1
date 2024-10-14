package stripe

import (
	"context"
	"time"
)

// ProcessPayment handles the transaction processing through Stripe for
// payments.
// 
// Errors:
//  - PermissionDenied:  Returned if the caller does not have rights to process the
//                       payment
//  - ResourceExhausted: Returned if there are insufficient resources to process the
//                       payment
//  - Aborted:           Returned if the payment process is aborted due to an error
// 
//encore:api auth method=POST path=/stripe/payments
func ProcessPayment(ctx context.Context, req *PaymentRequest) (*PaymentResponse, error) {
    panic("not yet implemented")
}

// PaymentRequest contains all data required to process a payment
type PaymentRequest struct {
  // Amount is the amount to be charged and must be greater than 0
  Amount float64 `json:"amount"`

  // Currency is the ISO 4217 currency code for the payment
  Currency string `json:"currency"`

  // CustomerID is the unique identifier of the customer making the payment
  CustomerID string `json:"customerID"`
}

// PaymentResponse provides details of the processed payment
type PaymentResponse struct {
  // PaymentID is the unique identifier of the transaction
  PaymentID string `json:"paymentID"`

  // Status indicates the result of the payment process
  Status string `json:"status"`

  // Timestamp is the time at which the payment was processed
  Timestamp time.Time `json:"timestamp"`
}