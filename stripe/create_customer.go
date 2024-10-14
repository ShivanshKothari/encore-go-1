package stripe

import (
	"context"
	"time"
)

// CreateCustomer creates a new customer in the Stripe system.
// 
//encore:api auth method=POST path=/stripe/customers
func CreateCustomer(ctx context.Context, req *CreateCustomerRequest) (*Customer, error) {
    panic("not yet implemented")
}

// CreateCustomerRequest contains the information necessary to create a new
// Stripe customer
type CreateCustomerRequest struct {
  // CustomerName must be longer than 2 characters
  CustomerName string `json:"customerName"`

  // Email must be a valid email address format
  Email string `json:"email"`
}

// Customer represents the Stripe customer entity
type Customer struct {
  // CustomerID is a unique identifier for a Stripe customer
  CustomerID string `json:"customerID"`

  // CustomerName of the Stripe customer
  CustomerName string `json:"customerName"`

  // Email of the Stripe customer
  Email string `json:"email"`

  CreatedAt time.Time `json:"createdAt"`
}