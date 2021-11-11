export const endpoint = process.env.NODE_ENV === 'production' ? '/.netlify/functions' : 'http://localhost:8090/.netlify/functions';
// The key needs to be replaced with a production one once we go live.
export const stripePublishableKey =
  'pk_test_51IzFveLwcEKoBoonzTPmtwFhbwqix42XQtnxeHPtDb2IutljsV75FQI9jVi9JOCxBHYxQLvy3eVwpTWXYpmVDplN00ZeSz3HSi';
