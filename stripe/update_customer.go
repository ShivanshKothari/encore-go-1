package stripe

import "context"

// UpdateCustomer updates customer details in the Stripe system.
// 
// Path Parameters:
//  - customerId: customerId is the unique identifier of the customer to update
// 
// Errors:
//  - NotFound: Returned if no customer is found with the given customerId
// 
//encore:api auth method=PUT path=/stripe/customers/:customerId
func UpdateCustomer(ctx context.Context, customerId string, req *UpdateCustomerRequest) (*Customer, error) {
    panic("not yet implemented")
}

// UpdateCustomerRequest contains the fields that can be updated for a Stripe
// customer
type UpdateCustomerRequest struct {
  // Email may be updated and must be a valid email address format
  Email string `json:"email"`
}