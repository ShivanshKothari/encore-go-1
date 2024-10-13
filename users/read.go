package users

import (
	"context"
	"time"
)

// Read fetches a user by userId.
// 
// Path Parameters:
//  - userId: the UUID of the user to be retrieved.
// 
// Errors:
//  - NotFound: Returned if the specified userID does not exist.
// 
//encore:api auth method=GET path=/users/:userId
func Read(ctx context.Context, userId string) (*User, error) {
    panic("not yet implemented")
}

// User represents the user entity in the system.
type User struct {
  // UserID is a unique identifier for each user.
  UserID string `json:"userID"`

  // Name of the user.
  Name string `json:"name"`

  // Email is used for user communication.
  Email string `json:"email"`

  // CreatedAt is the date and time when the user was created.
  CreatedAt time.Time `json:"createdAt"`
}