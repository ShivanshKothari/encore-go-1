package users

import "context"

// Create adds a new user to the user database.
// 
// Errors:
//  - AlreadyExists:      Returned if the user email already exists in the system.
//  - FailedPrecondition: Returned if the input data violates constraints.
// 
//encore:api auth method=POST path=/users
func Create(ctx context.Context, req *CreateUserRequest) (*User, error) {
    panic("not yet implemented")
}

// CreateUserRequest contains the information needed to create a user.
type CreateUserRequest struct {
  // Name must be unique and consist only of alphabetic characters.
  Name string `json:"name"`

  // Email must be a valid email address.
  Email string `json:"email"`

  // Password must be at least 8 characters long.
  Password string `json:"password"`
}