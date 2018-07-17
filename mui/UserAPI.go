package mui

// Save saves the user object in the database.
func (user *User) Save() {
	DB.Set("User", user.ID, user)
}
