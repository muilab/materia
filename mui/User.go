package mui

import (
	"errors"

	"github.com/aerogo/aero"
	"github.com/muilab/materia/mui/utils"
)

// User represents a single authenticated user.
type User struct {
	ID         ID           `json:"id"`
	Nick       string       `json:"nick" editable:"true"`
	FirstName  string       `json:"firstName" private:"true"`
	LastName   string       `json:"lastName" private:"true"`
	Registered string       `json:"registered"`
	LastLogin  string       `json:"lastLogin" private:"true"`
	LastSeen   string       `json:"lastSeen" private:"true"`
	IP         string       `json:"ip" private:"true"`
	Agent      string       `json:"agent" private:"true"`
	Accounts   UserAccounts `json:"accounts" private:"true"`
}

// UserAccounts includes the external accounts of a user.
type UserAccounts struct {
	Email struct {
		Address string `json:"address" private:"true"`
	} `json:"email"`

	Facebook struct {
		ID string `json:"id" private:"true"`
	} `json:"facebook"`

	Google struct {
		ID string `json:"id" private:"true"`
	} `json:"google"`
}

// NewUser creates an empty user object with a unique ID.
func NewUser() *User {
	user := &User{
		ID: GenerateID("User"),
	}

	return user
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

// Link returns the permalink of the user.
func (user *User) Link() string {
	return "/+" + user.ID
}

// ConnectGoogle connects the user's account with a Google account.
func (user *User) ConnectGoogle(googleID string) {
	if googleID == "" {
		return
	}

	user.Accounts.Google.ID = googleID

	DB.Set("GoogleToUser", googleID, &GoogleToUser{
		ID:     googleID,
		UserID: user.ID,
	})
}

// ConnectEmail connects the user's account with an Email account.
func (user *User) ConnectEmail(email string) {
	if email == "" {
		return
	}

	user.Accounts.Email.Address = email

	DB.Set("EmailToUser", email, &EmailToUser{
		Email:  email,
		UserID: user.ID,
	})
}

// GetUser fetches the user with the given ID from the database.
func GetUser(id string) (*User, error) {
	obj, err := DB.Get("User", id)

	if err != nil {
		return nil, err
	}

	return obj.(*User), nil
}

// GetUserByEmail fetches the user with the given email from the database.
func GetUserByEmail(email string) (*User, error) {
	if email == "" {
		return nil, errors.New("Email is empty")
	}

	obj, err := DB.Get("EmailToUser", email)

	if err != nil {
		return nil, err
	}

	userID := obj.(*EmailToUser).UserID
	user, err := GetUser(userID)

	return user, err
}

// GetUserByGoogleID fetches the user with the given Google ID from the database.
func GetUserByGoogleID(googleID string) (*User, error) {
	obj, err := DB.Get("GoogleToUser", googleID)

	if err != nil {
		return nil, err
	}

	userID := obj.(*GoogleToUser).UserID
	user, err := GetUser(userID)

	return user, err
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

// RegisterUser registers a new user in the database and sets up all the required references.
func RegisterUser(user *User) {
	user.Registered = utils.DateTimeUTC()
	user.LastLogin = user.Registered
	user.LastSeen = user.Registered
}
