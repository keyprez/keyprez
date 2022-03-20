package orders

type retrieveSessionRequest struct {
	SessionID string `json:"sessionid"`
}

type createCheckoutSessionRequest struct {
	PriceID          string `json:"priceid"`
	CustomerStripeID string `json:"customerstripeid"`
}
