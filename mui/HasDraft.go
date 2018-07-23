package mui

// HasDraft includes a boolean indicating whether the object is a draft.
type HasDraft struct {
	Public bool `json:"public" editable:"true"`
}

// GetPublic tells you whether the object is a draft or not.
func (obj *HasDraft) GetPublic() bool {
	return obj.Public
}

// SetPublic sets the draft state for this object.
func (obj *HasDraft) SetPublic(Public bool) {
	obj.Public = Public
}
