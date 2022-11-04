
      _
     | | _____ _   _ _ __  _ __ ___ ____
     | |/ / _ \ | | | '_ \| '__/ _ \_  /
     |   <  __/ |_| | |_) | | |  __// /
     |_|\_\___|\__, | .__/|_|  \___/___|
               |___/|_|

🚧 WIP

Application built with SvelteKit (currently in public beta), TypeScript and Vite. Deployed on S3.

Update necessary variables. `STRIPE_PUBLISHABLE_KEY` and `STRIPE_SECRET_KEY` can be found in Stripe Dashboard. `STRIPE_WEBHOOK_SECRET` is returned from running the command `npm run listen`

For local development run:

```bash
nvm use
npm install
npm start
```

With `stripe-cli` you can listen for completion webhook:

```bash
npm run listen
```

Instead of going through the whole completion process run:

```bash
npm run trigger
```

To build the application run:

```bash
npm run build
```

To run the production build preview locally run:

```bash
npm run preview
```
