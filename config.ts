export default {
  endpoint: process.env.NODE_ENV === 'production' ? 'https://keyprez.com' : 'http://localhost:3000',
  // The key needs to be replaced with a production one once we go live.
  stripePublishableKey:
    'pk_test_51IzFveLwcEKoBoonzTPmtwFhbwqix42XQtnxeHPtDb2IutljsV75FQI9jVi9JOCxBHYxQLvy3eVwpTWXYpmVDplN00ZeSz3HSi',
};
