package utils

import (
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
