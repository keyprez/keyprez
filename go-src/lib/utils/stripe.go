package utils

import (
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/checkout/session"
)

func CreateCheckoutSession() (*stripe.CheckoutSession, error) {
	stripe.Key = GetEnvVar("STRIPE_SECRET_KEY")

	params := &stripe.CheckoutSessionParams{
		SuccessURL: stripe.String("http://localhost:8080/success?session_id={CHECKOUT_SESSION_ID}"),
		CancelURL:  stripe.String("http://localhost:8080"),
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
				Name:     stripe.String("Eagle"),
				Amount:   stripe.Int64(200),
				Quantity: stripe.Int64(2),
				Currency: stripe.String("NOK"),
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

	return session.Get(SessionID, nil)
}
