package models

type Product struct {
	UniqueID             string   `json:"uniqueID"`
	PartNumber           string   `json:"partNumber"`
	Name                 string   `json:"name"`
	FullImage            string   `json:"fullImage"`
	Images               []string `json:"images"`
	ShortDescription     string   `json:"shortDescription"`
	LongDescription      string   `json:"longDescription"`
	XCatEntryCategory    string   `json:"xCatEntryCategory"`
	ProductString        string   `json:"productString"`
	IsMarketplaceProduct bool     `json:"isMarketplaceProduct"`
	URL                  string   `json:"url"`
	ThumbnailImage       string   `json:"thumbnailImage"`
	Single               bool     `json:"single"`
}
