package stripe

import "context"

// RetrieveCustomer retrieves customer details from Stripe using customerId.
// 
// Path Parameters:
//  - customerId: customerId is the unique identifier of the customer to retrieve
// 
// Errors:
//  - NotFound: Returned if no customer is found with the given customerId
// 
//encore:api auth method=GET path=/stripe/customers/:customerId
func RetrieveCustomer(ctx context.Context, customerId string) (*Customer, error) {
    panic("not yet implemented")
}

