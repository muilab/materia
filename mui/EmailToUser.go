package mui

// EmailToUser stores the user ID for an email address.
type EmailToUser struct {
	Email  string `json:"email"`
	UserID ID     `json:"userId"`
}
