package mui

import "github.com/aerogo/aero"

// User represents a single authenticated user.
type User struct {
	ID        string
	Email     string
	FirstName string
	LastName  string
}

// RealName returns the full name of the user.
func (user *User) RealName() string {
	if user.LastName == "" {
		return user.FirstName
	}

	if user.FirstName == "" {
		return user.LastName
	}

	return user.FirstName + " " + user.LastName
}

// GetUser fetches the user with the given ID from the database.
func GetUser(id string) (*User, error) {
	obj, err := DB.Get("User", id)

	if err != nil {
		return nil, err
	}

	return obj.(*User), nil
}

// GetUserFromContext returns the logged in user for the given context.
func GetUserFromContext(ctx *aero.Context) *User {
	if !ctx.HasSession() {
		return nil
	}

	userID := ctx.Session().GetString("userId")

	if userID == "" {
		return nil
	}

	user, err := GetUser(userID)

	if err != nil {
		return nil
	}

	return user
}
