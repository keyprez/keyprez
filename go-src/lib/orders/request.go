package orders

type retrieveSessionRequest struct {
	SessionID string `json:"sessionid"`
}

type createCheckoutSessionRequest struct {
	PriceID          string `json:"priceid"`
	CustomerStripeID string `json:"customerstripeid"`
}

type fetchShippingRateParams struct {
	Country    string `json:"country" validate:"required,eq=NO|eq=SE|eq=DK"`
	PostalCode string `json:"postalCode" validate:"required"`
}

type fetchShippingRateRequest struct {
	Language                         string        `json:"language"`
	WithPrice                        bool          `json:"withPrice"`
	WithExpectedDelivery             bool          `json:"withExpectedDelivery"`
	WithGuiInformation               bool          `json:"withGuiInformation"`
	NumberOfAlternativeDeliveryDates int16         `json:"numberOfAlternativeDeliveryDates"`
	WithUniqueAlternateDeliveryDates bool          `json:"withUniqueAlternateDeliveryDates"`
	Edi                              bool          `json:"edi"`
	PostingAtPostOffice              bool          `json:"postingAtPostOffice"`
	Trace                            bool          `json:"trace"`
	Consignments                     []consignment `json:"consignments"`
}

type consignment struct {
	ID                 int                 `json:"id"`
	Products           []bringProduct      `json:"products"`
	FromCountryCode    string              `json:"fromCountryCode"`
	ToCountryCode      string              `json:"toCountryCode"`
	FromPostalCode     string              `json:"fromPostalCode"`
	ToPostalCode       string              `json:"toPostalCode"`
	AddressLine        string              `json:"addressLine"`
	ShippingDate       shippingDate        `json:"shippingDate"`
	Packages           []packages          `json:"packages"`
	AdditionalServices []additionalService `json:"additionalServices"`
}

type shippingDate struct {
	Day    string `json:"day"`
	Hour   string `json:"hour"`
	Minute string `json:"minute"`
	Month  string `json:"month"`
	Year   string `json:"year"`
}

type packages struct {
	ID          string `json:"id"`
	Length      int    `json:"length"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	GrossWeight int    `json:"grossWeight"`
}

type bringProduct struct {
	ID string `json:"id"`
}
type additionalService struct {
	ID string `json:"id"`
}
