package mui

import "github.com/aerogo/aero"

// DeleteInContext deletes the sample in the given context.
func (sample *MaterialSample) DeleteInContext(ctx *aero.Context) error {
	return sample.Delete()
}

// Delete deletes the sample from the database.
func (sample *MaterialSample) Delete() error {
	// Delete all image files
	for _, output := range materialSampleImageOutputs {
		output.Delete(sample.ID)
	}

	DB.Delete("MaterialSample", sample.ID)
	return nil
}

// Save saves the sample in the database.
func (sample *MaterialSample) Save() {
	DB.Set("MaterialSample", sample.ID, sample)
}
