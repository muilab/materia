package mui

// ImageFile represents the information about an image file on disk.
type ImageFile struct {
	Extension    string   `json:"extension"`
	Width        int      `json:"width"`
	Height       int      `json:"height"`
	AverageColor HSLColor `json:"averageColor"`
	LastModified int64    `json:"lastModified"`
}
