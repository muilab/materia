package mui

// HasEditor includes user ID and date for the last edit of this object.
type HasEditor struct {
	Edited   DateTime `json:"edited"`
	EditedBy ID       `json:"editedBy"`
}

// Editor returns the user who last edited this object.
func (obj *HasEditor) Editor() *User {
	user, _ := GetUser(obj.EditedBy)
	return user
}
