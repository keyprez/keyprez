package utils

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/checkout/session"
	"github.com/stripe/stripe-go/v72/customer"
)

type CreateCustomerRequest struct {
	Email      string `bson:"email" validate:"required,email"`
	Name       string `bson:"name" validate:"required,min=2"`
	City       string `bson:"city" validate:"required"`
	Country    string `bson:"country" validate:"required,eq=NO|eq=SE|eq=DK"`
	Line1      string `bson:"line1" validate:"required"`
	Line2      string `bson:"line2" validate:"omitempty,min=2"`
	PostalCode string `bson:"postalCode" validate:"required"`
	State      string `bson:"state"`
}

func CreateCheckoutSession(priceId string, customerStripeId string) (*stripe.CheckoutSession, error) {
	stripe.Key = GetEnvVar("STRIPE_SECRET_KEY")
	stripeCallbackUrl := GetEnvVar("STRIPE_CALLBACK_URL")
	successUrl := fmt.Sprintf("%s/success?session_id={CHECKOUT_SESSION_ID}", stripeCallbackUrl)

	params := &stripe.CheckoutSessionParams{
		SuccessURL: stripe.String(successUrl),
		CancelURL:  stripe.String(stripeCallbackUrl),
		Customer:   stripe.String(customerStripeId),
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
			"klarna",
		}),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			&stripe.CheckoutSessionLineItemParams{
				Price:    stripe.String(priceId),
				Quantity: stripe.Int64(1),
				AdjustableQuantity: &stripe.CheckoutSessionLineItemAdjustableQuantityParams{
					Enabled: stripe.Bool(true),
					Minimum: stripe.Int64(1),
					Maximum: stripe.Int64(5),
				},
			},
		},
		Mode: stripe.String(string(stripe.CheckoutSessionModePayment)),
	}

	return session.New(params)
}

func CreateCustomer(createRequest CreateCustomerRequest) (*stripe.Customer, error) {
	stripe.Key = GetEnvVar("STRIPE_SECRET_KEY")

	params := &stripe.CustomerParams{
		Email: stripe.String(createRequest.Email),
		Name:  stripe.String(createRequest.Name),
		Address: &stripe.AddressParams{
			City:       stripe.String(createRequest.City),
			Country:    stripe.String(createRequest.Country),
			Line1:      stripe.String(createRequest.Line1),
			Line2:      stripe.String(createRequest.Line2),
			PostalCode: stripe.String(createRequest.PostalCode),
			State:      stripe.String(createRequest.State),
		},
	}

	return customer.New(params)
}

func RetrieveSession(SessionID string) (*stripe.CheckoutSession, error) {
	stripe.Key = GetEnvVar("STRIPE_SECRET_KEY")

	params := &stripe.CheckoutSessionParams{}
	params.AddExpand("line_items")
	params.AddExpand("customer")

	return session.Get(SessionID, params)
}

func VerifyWebhookHeader(header string) (int64, error) {
	var timestamp string
	var signature string

	pairs := strings.Split(header, ",")
	for _, pair := range pairs {
		parts := strings.Split(pair, "=")
		switch parts[0] {
		case "t":
			timestamp = parts[1]
		case "v1":
			signature = parts[1]
		}
	}

	if timestamp == "" {
		return 0, errors.New("missing webhook timestamp")
	}

	if signature == "" {
		return 0, errors.New("missing webhook signature")
	}

	res, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return 0, errors.New("invalid timestamp")
	}

	return res, nil
}

func VerifyTimestamp(eventTime int64, webhookTime int64) error {
	if webhookTime-eventTime > 10 {
		return errors.New("webhook too old")
	}
	return nil
}
