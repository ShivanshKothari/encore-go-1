package users

import "context"

// Update modifies an existing user's information.
// 
// Path Parameters:
//  - userId: the UUID of the user to be updated.
// 
// Errors:
//  - NotFound:           Returned if the user with specified userID is not found.
//  - FailedPrecondition: Returned if the update data violates constraints.
// 
//encore:api auth method=PUT path=/users/:userId
func Update(ctx context.Context, userId string, req *UpdateUserRequest) (*User, error) {
    panic("not yet implemented")
}

// UpdateUserRequest contains the fields that can be updated for a user.
type UpdateUserRequest struct {
  // Name must be unique and consist only of alphabetic characters.
  Name string `json:"name"`

  // Email must be a valid email address.
  Email string `json:"email"`
}