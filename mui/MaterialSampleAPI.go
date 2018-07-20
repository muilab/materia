package mui

// Save saves the sample in the database.
func (sample *MaterialSample) Save() {
	DB.Set("MaterialSample", sample.ID, sample)
}
