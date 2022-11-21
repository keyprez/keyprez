export const endpoint =
  process.env.NODE_ENV === 'production'
    ? 'https://qmw0sv0yz7.execute-api.eu-west-1.amazonaws.com'
    : 'http://localhost:3000';
// The key needs to be replaced with a production one once we go live.
export const stripePublishableKey =
  'pk_test_51IzFveLwcEKoBoonzTPmtwFhbwqix42XQtnxeHPtDb2IutljsV75FQI9jVi9JOCxBHYxQLvy3eVwpTWXYpmVDplN00ZeSz3HSi';
