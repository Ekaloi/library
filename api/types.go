package api

type BookResponse struct {
	Kind       string     `json:"kind"`
    TotalItems int        `json:"totalItems"`
    Items      []BookItem `json:"items"`
}

type BookItem struct {
	Kind       string  `json:"kind"`
	ID         string  `json:"id"`
	ETag       string  `json:"etag"`
	SelfLink   string  `json:"selfLink"`
	VolumeInfo Volume  `json:"volumeInfo"`
	SaleInfo   Sale    `json:"saleInfo"`
	AccessInfo Access  `json:"accessInfo"`
	SearchInfo *Search `json:"searchInfo,omitempty"`
}

type Volume struct {
	Title               string       `json:"title"`
	Authors             []string     `json:"authors"`
	Publisher           string       `json:"publisher"`
	PublishedDate       string       `json:"publishedDate"`
	Description         string       `json:"description"`
	IndustryIdentifiers []IndustryID `json:"industryIdentifiers"`
	PageCount           int          `json:"pageCount"`
	PrintType           string       `json:"printType"`
	Categories          []string     `json:"categories"`
	MaturityRating      string       `json:"maturityRating"`
	Language            string       `json:"language"`
	PreviewLink         string       `json:"previewLink"`
	InfoLink            string       `json:"infoLink"`
	CanonicalVolumeLink string       `json:"canonicalVolumeLink"`
	ImageLinks          ImageLinks   `json:"imageLinks"`
}

type IndustryID struct {
	Type       string `json:"type"`
	Identifier string `json:"identifier"`
}

type ImageLinks struct {
	SmallThumbnail string `json:"smallThumbnail"`
	Thumbnail      string `json:"thumbnail"`
}

type Sale struct {
	Country     string `json:"country"`
	Saleability string `json:"saleability"`
	IsEbook     bool   `json:"isEbook"`
	ListPrice   *Price `json:"listPrice,omitempty"`
	RetailPrice *Price `json:"retailPrice,omitempty"`
	BuyLink     string `json:"buyLink,omitempty"`
}

type Price struct {
	Amount       float64 `json:"amount"`
	CurrencyCode string  `json:"currencyCode"`
}

type Access struct {
	Country             string `json:"country"`
	Viewability         string `json:"viewability"`
	Embeddable          bool   `json:"embeddable"`
	PublicDomain        bool   `json:"publicDomain"`
	TextToSpeechAllowed string `json:"textToSpeechPermission"`
	WebReaderLink       string `json:"webReaderLink"`
}

type Search struct {
	TextSnippet string `json:"textSnippet"`
}
