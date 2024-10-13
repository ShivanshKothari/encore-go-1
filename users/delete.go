package users

import "context"

// Delete removes a user from the database.
// 
// Path Parameters:
//  - userId: the UUID of the user to be deleted.
// 
// Errors:
//  - NotFound:           Returned if the user with specified userID is not found.
//  - FailedPrecondition: Returned if deletion prerequisites are not met.
// 
//encore:api auth method=DELETE path=/users/:userId
func Delete(ctx context.Context, userId string) error {
    panic("not yet implemented")
}

