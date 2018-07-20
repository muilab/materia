package mui

import (
	"errors"

	"github.com/aerogo/aero"
	"github.com/aerogo/api"
)

// Force interface implementations
var (
	_ api.Deletable = (*MaterialSample)(nil)
)

// Authorize returns an error if the given API POST request is not authorized.
func (sample *MaterialSample) Authorize(ctx *aero.Context, action string) error {
	user := GetUserFromContext(ctx)

	if user == nil {
		return errors.New("Not logged in")
	}

	if user.ID != sample.Material().CreatedBy {
		return errors.New("Can't edit data from other users")
	}

	return nil
}

// DeleteInContext deletes the sample in the given context.
func (sample *MaterialSample) DeleteInContext(ctx *aero.Context) error {
	return sample.Delete()
}

// Delete deletes the sample from the database.
func (sample *MaterialSample) Delete() error {
	// Delete the reference in the material file
	material := sample.Material()
	material.RemoveSample(sample.ID)
	material.Save()

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
