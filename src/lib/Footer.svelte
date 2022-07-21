<script>
  import { endpoint } from '../config';
  import { createForm } from 'svelte-forms-lib';
  import * as yup from 'yup';
  import { Button } from '$lib';

  let displayedEmail = '';
  let loading = false;
  let showError = false;
  let showSuccess = false;
  let hasSubscription = false;
  let inputRef;
  let focus = () => inputRef.focus();

  const onSubmit = async ({ email }) => {
    loading = true;

    const res = await fetch(`${endpoint}/newsletters`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ email }),
    });

    loading = false;

    if (res.status === 200 || res.status === 201) {
      showSuccess = true;
      displayedEmail = email;
      hasSubscription = res.status === 200;
      return setTimeout(() => (showSuccess = false), 5000);
    } else {
      showError = true;
      return setTimeout(() => (showError = false), 5000);
    }
  };

  const { form, errors, handleChange, handleSubmit } = createForm({
    initialValues: { email: '' },
    validationSchema: yup.object().shape({
      email: yup.string().email(),
    }),
    onSubmit,
  });
</script>

<footer>
  <h2>
    {#if showError}
      <h2>There was an error ❌ Please try again later</h2>
    {:else if hasSubscription}
      <h2><strong>{displayedEmail}</strong> has been already subscribed to our newsletter 😊</h2>
    {:else if showSuccess}
      <h2>Thank you for subscribing to our newsletter 🎉</h2>
    {:else}
      <h2><strong>SUBSCRIBE</strong> to latest news and updates</h2>
    {/if}
  </h2>
  <form on:submit={handleSubmit}>
    <div>
      <input
        bind:this={inputRef}
        class:errorInput={$errors.email}
        type="text"
        name="email"
        placeholder="Type your email"
        bind:value={$form.email}
        on:keyup={handleChange}
        on:blur={handleChange}
      />
      <small class="error {$errors.email ? 'errorVisible' : ''}">{$errors.email}</small>
    </div>
    <Button type="submit" text="SUBSCRIBE" {loading} onClick={focus} />
  </form>
  <div class="links">
    <a href="https://github.com/keyprez/keyprez" target="blank"><img src="/github_logo.svg" alt="GitHub" /></a>
    <a href="https://github.com/keyprez/keyprez" target="blank"><img src="/instagram_logo.svg" alt="Instagram" /></a>
    <a href="https://github.com/keyprez/keyprez" target="blank"><img src="/reddit_logo.svg" alt="Reddit" /></a>
    <a href="https://github.com/keyprez/keyprez" target="blank"><img src="/meetup_logo.svg" alt="Meetup" /></a>
  </div>
</footer>

<style lang="scss">
  @import 'src/variables';

  footer {
    display: flex;
    flex-direction: column;
    gap: 2rem;
    justify-content: center;
    padding: 10rem 1rem 5rem;
  }

  h2 {
    letter-spacing: 2px;
    color: inherit;
    font-weight: 1000;
    margin: 0;
  }

  form {
    display: flex;
    flex-direction: row;
    gap: 0.5rem;
    justify-content: center;

    div {
      position: relative;
      display: flex;
      width: 100%;
      max-width: 400px;
    }
  }

  .error {
    display: none;
    position: absolute;
    right: 0.3em;
    bottom: 0.3em;
    color: $color-warning;

    &Visible {
      display: block;
    }

    &Input {
      border-color: $color-warning;
    }
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

  @media (max-width: 600px) {
    footer {
      padding: 5rem 1rem 3rem;
    }

    form {
      flex-direction: column;

      div {
        max-width: 100%;
      }
    }
  }
</style>
