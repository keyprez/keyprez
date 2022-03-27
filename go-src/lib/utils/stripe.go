package utils

import (
	"errors"
	"strconv"
	"strings"

	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/checkout/session"
	"github.com/stripe/stripe-go/v72/customer"
)

func CreateCheckoutSession(priceId string, customerStripeId string) (*stripe.CheckoutSession, error) {
	stripe.Key = GetEnvVar("STRIPE_SECRET_KEY")

	params := &stripe.CheckoutSessionParams{
		SuccessURL: stripe.String("http://localhost:8080/success?session_id={CHECKOUT_SESSION_ID}"),
		CancelURL:  stripe.String("http://localhost:8080"),
		Customer:   stripe.String(customerStripeId),
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
			"klarna",
		}),
		ShippingAddressCollection: &stripe.CheckoutSessionShippingAddressCollectionParams{
			AllowedCountries: stripe.StringSlice([]string{
				"NO",
			}),
		},
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

func CreateCustomer(email string) (*stripe.Customer, error) {
	stripe.Key = GetEnvVar("STRIPE_SECRET_KEY")

	params := &stripe.CustomerParams{
		Email: stripe.String(email),
	}

	return customer.New(params)
}

func RetrieveSession(SessionID string) (*stripe.CheckoutSession, error) {
	stripe.Key = GetEnvVar("STRIPE_SECRET_KEY")

	params := &stripe.CheckoutSessionParams{}
	params.AddExpand("line_items")

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
