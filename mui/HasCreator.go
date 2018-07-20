package mui

// HasCreator includes user ID and date for the creation of this object.
type HasCreator struct {
	Created   DateTime `json:"created"`
	CreatedBy ID       `json:"createdBy"`
}

// Creator returns the user who created this object.
func (obj *HasCreator) Creator() *User {
	user, _ := GetUser(obj.CreatedBy)
	return user
}

// GetCreatedBy returns the ID of the user who created this object.
func (obj *HasCreator) GetCreatedBy() string {
	return obj.CreatedBy
}
