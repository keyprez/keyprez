<script>
  import { endpoint } from '../../config';

  let email = '';
  let loading = false;
  let showError = false;
  let errorMessage = '';
  let showSuccess = false;

  const isValidEmail = (email) => /^\w+([-+.']\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$/.test(email);

  const handleSubmit = async (e) => {
    loading = true;
    email = e.target.email.value;

    if (!isValidEmail(email)) {
      loading = false;
      errorMessage = 'Please provide a valid email';
      showError = true;
      return setTimeout(() => (showError = false), 5000);
    }

    const res = await fetch(`${endpoint}/newsletter`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ email }),
    });

    loading = false;

    if (res.status === 200) {
      showSuccess = true;
      e.target.email.value = '';
      return setTimeout(() => (showSuccess = false), 5000);
    } else {
      errorMessage = `${res.statusText}❌ Please try again later`;
      showError = true;
      return setTimeout(() => (showError = false), 5000);
    }
  };
</script>

<footer>
  {#if loading}
    <h2>Sending subscription request...</h2>
  {:else if showError}
    <h2>{errorMessage}</h2>
  {:else if showSuccess}
    <h2>Your email <strong class="subscribe">{email}</strong> has been subscribed to our newsletter 🎉</h2>
  {:else}
    <h2><strong class="subscribe">SUBSCRIBE</strong> to latest news and updates</h2>
  {/if}
  <form class="form" on:submit|preventDefault={handleSubmit}>
    <input type="text" name="email" placeholder="Your email" required />
    <button>SUBSCRIBE</button>
  </form>
  <div class="links">
    <a href="https://github.com/keyprez/keyprez" target="blank"><img src="/github_logo.svg" alt="GitHub" /></a>
    <a href="https://github.com/keyprez/keyprez" target="blank"><img src="/instagram_logo.svg" alt="Instagram" /></a>
    <a href="https://github.com/keyprez/keyprez" target="blank"><img src="/reddit_logo.svg" alt="Reddit" /></a>
    <a href="https://github.com/keyprez/keyprez" target="blank"><img src="/meetup_logo.svg" alt="Meetup" /></a>
  </div>
</footer>

<style type="text/scss">
  @import 'src/variables';

  footer {
    display: flex;
    flex-direction: column;
    gap: 2rem;
    justify-content: center;
    letter-spacing: 2px;
    padding: 10rem 1rem;
  }

  h2,
  strong {
    color: inherit;
    font-weight: 1000;
    margin: 0;
  }

  .subscribe {
    color: $color-tertiary;
  }

  .form {
    display: flex;
    flex-direction: row;
    gap: 0.5rem;
    justify-content: center;
  }

  input {
    width: 12rem;
  }

  .links {
    display: flex;
    flex-direction: row;
    gap: 1rem;
    justify-content: center;
  }

  img {
    width: 1.5rem;
  }

  img:hover {
    transform: scale(1.3);
  }

  @media (min-width: 768px) {
    input {
      width: 20rem;
    }
  }
</style>
