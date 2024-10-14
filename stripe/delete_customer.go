package stripe

import "context"

// DeleteCustomer deletes a customer in the Stripe system.
// 
// Path Parameters:
//  - customerId: customerId is the unique identifier of the customer to delete
// 
// Errors:
//  - NotFound:           Returned if no customer is found with the given customerId
//  - FailedPrecondition: Returned if the customer cannot be deleted due to existing
//                        dependencies
// 
//encore:api auth method=DELETE path=/stripe/customers/:customerId
func DeleteCustomer(ctx context.Context, customerId string) (*DeletionConfirmation, error) {
    panic("not yet implemented")
}

// DeletionConfirmation provides information on the success of a delete
// operation
type DeletionConfirmation struct {
  // Success indicates if the deletion was successful
  Success bool `json:"success"`

  // Message provides additional information on the deletion operation
  Message string `json:"message"`
}